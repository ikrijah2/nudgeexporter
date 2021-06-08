// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trace

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/ikrijah2/nudge/consumer"
	"github.com/ikrijah2/nudge/consumer/consumertest"
	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/internal"
	"github.com/ikrijah2/nudge/internal/data"
	collectortrace "github.com/ikrijah2/nudge/internal/data/protogen/collector/trace/v1"
	otlptrace "github.com/ikrijah2/nudge/internal/data/protogen/trace/v1"
	"github.com/ikrijah2/nudge/obsreport"
)

var _ collectortrace.TraceServiceServer = (*Receiver)(nil)

func TestExport(t *testing.T) {
	// given

	traceSink := new(consumertest.TracesSink)

	port, doneFn := otlpReceiverOnGRPCServer(t, traceSink)
	defer doneFn()

	traceClient, traceClientDoneFn, err := makeTraceServiceClient(port)
	require.NoError(t, err, "Failed to create the TraceServiceClient: %v", err)
	defer traceClientDoneFn()

	// when

	unixnanos := uint64(12578940000000012345)
	traceID := [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}
	spanID := [8]byte{8, 7, 6, 5, 4, 3, 2, 1}
	req := &collectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{
							{
								TraceId:           data.NewTraceID(traceID),
								SpanId:            data.NewSpanID(spanID),
								Name:              "operationB",
								Kind:              otlptrace.Span_SPAN_KIND_SERVER,
								StartTimeUnixNano: unixnanos,
								EndTimeUnixNano:   unixnanos,
								Status:            otlptrace.Status{Message: "status-cancelled", Code: otlptrace.Status_STATUS_CODE_ERROR},
								TraceState:        "a=text,b=123",
							},
						},
					},
				},
			},
		},
	}

	// Keep trace data to compare the test result against it
	// Clone needed because OTLP proto XXX_ fields are altered in the GRPC downstream
	traceData := pdata.TracesFromInternalRep(internal.TracesFromOtlp(req)).Clone()

	resp, err := traceClient.Export(context.Background(), req)
	require.NoError(t, err, "Failed to export trace: %v", err)
	require.NotNil(t, resp, "The response is missing")

	// assert

	require.Equal(t, 1, len(traceSink.AllTraces()), "unexpected length: %v", len(traceSink.AllTraces()))

	assert.EqualValues(t, traceData, traceSink.AllTraces()[0])
}

func TestExport_EmptyRequest(t *testing.T) {
	traceSink := new(consumertest.TracesSink)

	addr, doneFn := otlpReceiverOnGRPCServer(t, traceSink)
	defer doneFn()

	traceClient, traceClientDoneFn, err := makeTraceServiceClient(addr)
	require.NoError(t, err, "Failed to create the TraceServiceClient: %v", err)
	defer traceClientDoneFn()

	resp, err := traceClient.Export(context.Background(), &collectortrace.ExportTraceServiceRequest{})
	assert.NoError(t, err, "Failed to export trace: %v", err)
	assert.NotNil(t, resp, "The response is missing")
}

func TestExport_ErrorConsumer(t *testing.T) {
	addr, doneFn := otlpReceiverOnGRPCServer(t, consumertest.NewErr(errors.New("my error")))
	defer doneFn()

	traceClient, traceClientDoneFn, err := makeTraceServiceClient(addr)
	require.NoError(t, err, "Failed to create the TraceServiceClient: %v", err)
	defer traceClientDoneFn()

	req := &collectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{
							{
								Name: "operationB",
							},
						},
					},
				},
			},
		},
	}

	resp, err := traceClient.Export(context.Background(), req)
	assert.EqualError(t, err, "rpc error: code = Unknown desc = my error")
	assert.Nil(t, resp)
}

func makeTraceServiceClient(addr net.Addr) (collectortrace.TraceServiceClient, func(), error) {
	cc, err := grpc.Dial(addr.String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}

	metricsClient := collectortrace.NewTraceServiceClient(cc)

	doneFn := func() { _ = cc.Close() }
	return metricsClient, doneFn, nil
}

func otlpReceiverOnGRPCServer(t *testing.T, tc consumer.Traces) (net.Addr, func()) {
	ln, err := net.Listen("tcp", "localhost:")
	require.NoError(t, err, "Failed to find an available address to run the gRPC server: %v", err)

	doneFnList := []func(){func() { ln.Close() }}
	done := func() {
		for _, doneFn := range doneFnList {
			doneFn()
		}
	}

	r := New(receiverTagValue, tc)
	require.NoError(t, err)

	// Now run it as a gRPC server
	srv := obsreport.GRPCServerWithObservabilityEnabled()
	collectortrace.RegisterTraceServiceServer(srv, r)
	go func() {
		_ = srv.Serve(ln)
	}()

	return ln.Addr(), done
}
