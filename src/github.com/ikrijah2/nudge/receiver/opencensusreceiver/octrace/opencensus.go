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

package octrace

import (
	"context"
	"errors"
	"io"

	commonpb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	agenttracepb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/trace/v1"
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"

	"github.com/ikrijah2/nudge/client"
	"github.com/ikrijah2/nudge/component/componenterror"
	"github.com/ikrijah2/nudge/consumer"
	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/obsreport"
	"github.com/ikrijah2/nudge/translator/internaldata"
)

const (
	receiverTagValue   = "oc_trace"
	receiverTransport  = "grpc" // TODO: transport is being hard coded for now, investigate if info is available on context.
	receiverDataFormat = "protobuf"
)

// Receiver is the type used to handle spans from OpenCensus exporters.
type Receiver struct {
	agenttracepb.UnimplementedTraceServiceServer
	nextConsumer consumer.Traces
	instanceName string
}

// New creates a new opencensus.Receiver reference.
func New(instanceName string, nextConsumer consumer.Traces, opts ...Option) (*Receiver, error) {
	if nextConsumer == nil {
		return nil, componenterror.ErrNilNextConsumer
	}

	ocr := &Receiver{
		nextConsumer: nextConsumer,
		instanceName: instanceName,
	}
	for _, opt := range opts {
		opt(ocr)
	}

	return ocr, nil
}

var _ agenttracepb.TraceServiceServer = (*Receiver)(nil)

var errUnimplemented = errors.New("unimplemented")

// Config handles configuration messages.
func (ocr *Receiver) Config(agenttracepb.TraceService_ConfigServer) error {
	// TODO: Implement when we define the config receiver/sender.
	return errUnimplemented
}

var errTraceExportProtocolViolation = errors.New("protocol violation: Export's first message must have a Node")

// Export is the gRPC method that receives streamed traces from
// OpenCensus-traceproto compatible libraries/applications.
func (ocr *Receiver) Export(tes agenttracepb.TraceService_ExportServer) error {
	ctx := tes.Context()
	if c, ok := client.FromGRPC(ctx); ok {
		ctx = client.NewContext(ctx, c)
	}

	longLivedRPCCtx := obsreport.ReceiverContext(ctx, ocr.instanceName, receiverTransport)

	// The first message MUST have a non-nil Node.
	recv, err := tes.Recv()
	if err != nil {
		return err
	}

	// Check the condition that the first message has a non-nil Node.
	if recv.Node == nil {
		return errTraceExportProtocolViolation
	}

	var lastNonNilNode *commonpb.Node
	var resource *resourcepb.Resource
	// Now that we've got the first message with a Node, we can start to receive streamed up spans.
	for {
		lastNonNilNode, resource, err = ocr.processReceivedMsg(
			longLivedRPCCtx,
			lastNonNilNode,
			resource,
			recv)
		if err != nil {
			return err
		}

		recv, err = tes.Recv()
		if err != nil {
			if err == io.EOF {
				// Do not return EOF as an error so that grpc-gateway calls get an empty
				// response with HTTP status code 200 rather than a 500 error with EOF.
				return nil
			}
			return err
		}
	}
}

func (ocr *Receiver) processReceivedMsg(
	longLivedRPCCtx context.Context,
	lastNonNilNode *commonpb.Node,
	resource *resourcepb.Resource,
	recv *agenttracepb.ExportTraceServiceRequest,
) (*commonpb.Node, *resourcepb.Resource, error) {
	// If a Node has been sent from downstream, save and use it.
	if recv.Node != nil {
		lastNonNilNode = recv.Node
	}

	// TODO(songya): differentiate between unset and nil resource. See
	// https://github.com/census-instrumentation/opencensus-proto/issues/146.
	if recv.Resource != nil {
		resource = recv.Resource
	}

	td := internaldata.OCToTraces(lastNonNilNode, resource, recv.Spans)
	err := ocr.sendToNextConsumer(longLivedRPCCtx, td)
	return lastNonNilNode, resource, err
}

func (ocr *Receiver) sendToNextConsumer(longLivedRPCCtx context.Context, td pdata.Traces) error {
	ctx := obsreport.StartTraceDataReceiveOp(
		longLivedRPCCtx,
		ocr.instanceName,
		receiverTransport,
		obsreport.WithLongLivedCtx())

	err := ocr.nextConsumer.ConsumeTraces(ctx, td)
	obsreport.EndTraceDataReceiveOp(ctx, receiverDataFormat, td.SpanCount(), err)

	return err
}
