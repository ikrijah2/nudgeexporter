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

package pdata

import (
	"testing"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	goproto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ikrijah2/nudge/internal"
	otlpcollectortrace "github.com/ikrijah2/nudge/internal/data/protogen/collector/trace/v1"
	otlptrace "github.com/ikrijah2/nudge/internal/data/protogen/trace/v1"
)

func TestSpanCount(t *testing.T) {
	md := NewTraces()
	assert.EqualValues(t, 0, md.SpanCount())

	md.ResourceSpans().Resize(1)
	assert.EqualValues(t, 0, md.SpanCount())

	md.ResourceSpans().At(0).InstrumentationLibrarySpans().Resize(1)
	assert.EqualValues(t, 0, md.SpanCount())

	md.ResourceSpans().At(0).InstrumentationLibrarySpans().At(0).Spans().Resize(1)
	assert.EqualValues(t, 1, md.SpanCount())

	rms := md.ResourceSpans()
	rms.Resize(3)
	rms.At(0).InstrumentationLibrarySpans().Resize(1)
	rms.At(0).InstrumentationLibrarySpans().At(0).Spans().Resize(1)
	rms.At(1).InstrumentationLibrarySpans().Resize(1)
	rms.At(2).InstrumentationLibrarySpans().Resize(1)
	rms.At(2).InstrumentationLibrarySpans().At(0).Spans().Resize(5)
	assert.EqualValues(t, 6, md.SpanCount())
}

func TestSize(t *testing.T) {
	td := NewTraces()
	assert.Equal(t, 0, td.OtlpProtoSize())
	rms := td.ResourceSpans()
	rms.Resize(1)
	rms.At(0).InstrumentationLibrarySpans().Resize(1)
	rms.At(0).InstrumentationLibrarySpans().At(0).Spans().Resize(1)
	rms.At(0).InstrumentationLibrarySpans().At(0).Spans().At(0).SetName("foo")
	otlp := internal.TracesToOtlp(td.InternalRep())
	size := otlp.Size()
	bytes, err := otlp.Marshal()
	require.NoError(t, err)
	assert.Equal(t, size, td.OtlpProtoSize())
	assert.Equal(t, len(bytes), td.OtlpProtoSize())
}

func TestTracesSizeWithNil(t *testing.T) {
	assert.Equal(t, 0, NewTraces().OtlpProtoSize())
}

func TestSpanCountWithEmpty(t *testing.T) {
	assert.EqualValues(t, 0, Traces{orig: &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{{}},
	}}.SpanCount())
	assert.EqualValues(t, 0, Traces{orig: &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{{}},
			},
		},
	}}.SpanCount())
	assert.EqualValues(t, 1, Traces{orig: &otlpcollectortrace.ExportTraceServiceRequest{
		ResourceSpans: []*otlptrace.ResourceSpans{
			{
				InstrumentationLibrarySpans: []*otlptrace.InstrumentationLibrarySpans{
					{
						Spans: []*otlptrace.Span{{}},
					},
				},
			},
		},
	}}.SpanCount())
}

func TestSpanStatusCode(t *testing.T) {
	td := NewTraces()
	rss := td.ResourceSpans()
	rss.Resize(1)
	rss.At(0).InstrumentationLibrarySpans().Resize(1)
	rss.At(0).InstrumentationLibrarySpans().At(0).Spans().Resize(1)
	status := rss.At(0).InstrumentationLibrarySpans().At(0).Spans().At(0).Status()

	// Check handling of deprecated status code, see spec here:
	// https://github.com/open-telemetry/opentelemetry-proto/blob/59c488bfb8fb6d0458ad6425758b70259ff4a2bd/opentelemetry/proto/trace/v1/trace.proto#L231
	//
	// 2. New senders, which are aware of the `code` field MUST set both the
	// `deprecated_code` and `code` fields according to the following rules:
	//
	//   if code==STATUS_CODE_UNSET then `deprecated_code` MUST be
	//   set to DEPRECATED_STATUS_CODE_OK.
	status.orig.DeprecatedCode = otlptrace.Status_DEPRECATED_STATUS_CODE_UNKNOWN_ERROR
	status.SetCode(StatusCodeUnset)
	assert.EqualValues(t, otlptrace.Status_DEPRECATED_STATUS_CODE_OK, status.orig.DeprecatedCode)

	//   if code==STATUS_CODE_OK then `deprecated_code` MUST be
	//   set to DEPRECATED_STATUS_CODE_OK.
	status.orig.DeprecatedCode = otlptrace.Status_DEPRECATED_STATUS_CODE_UNKNOWN_ERROR
	status.SetCode(StatusCodeOk)
	assert.EqualValues(t, otlptrace.Status_DEPRECATED_STATUS_CODE_OK, status.orig.DeprecatedCode)

	//   if code==STATUS_CODE_ERROR then `deprecated_code` MUST be
	//   set to DEPRECATED_STATUS_CODE_UNKNOWN_ERROR.
	status.orig.DeprecatedCode = otlptrace.Status_DEPRECATED_STATUS_CODE_OK
	status.SetCode(StatusCodeError)
	assert.EqualValues(t, otlptrace.Status_DEPRECATED_STATUS_CODE_UNKNOWN_ERROR, status.orig.DeprecatedCode)
}

func TestToFromOtlp(t *testing.T) {
	otlp := &otlpcollectortrace.ExportTraceServiceRequest{}
	td := TracesFromInternalRep(internal.TracesFromOtlp(otlp))
	assert.EqualValues(t, NewTraces(), td)
	assert.EqualValues(t, otlp, internal.TracesToOtlp(td.InternalRep()))
	// More tests in ./tracedata/trace_test.go. Cannot have them here because of
	// circular dependency.
}

func TestResourceSpansWireCompatibility(t *testing.T) {
	// This test verifies that OTLP ProtoBufs generated using goproto lib in
	// opentelemetry-proto repository OTLP ProtoBufs generated using gogoproto lib in
	// this repository are wire compatible.

	// Generate ResourceSpans as pdata struct.
	pdataRS := generateTestResourceSpans()

	// Marshal its underlying ProtoBuf to wire.
	wire1, err := gogoproto.Marshal(pdataRS.orig)
	assert.NoError(t, err)
	assert.NotNil(t, wire1)

	// Unmarshal from the wire to OTLP Protobuf in goproto's representation.
	var goprotoMessage emptypb.Empty
	err = goproto.Unmarshal(wire1, &goprotoMessage)
	assert.NoError(t, err)

	// Marshal to the wire again.
	wire2, err := goproto.Marshal(&goprotoMessage)
	assert.NoError(t, err)
	assert.NotNil(t, wire2)

	// Unmarshal from the wire into gogoproto's representation.
	var gogoprotoRS2 otlptrace.ResourceSpans
	err = gogoproto.Unmarshal(wire2, &gogoprotoRS2)
	assert.NoError(t, err)

	// Now compare that the original and final ProtoBuf messages are the same.
	// This proves that goproto and gogoproto marshaling/unmarshaling are wire compatible.
	assert.EqualValues(t, pdataRS.orig, &gogoprotoRS2)
}

func TestTracesToFromOtlpProtoBytes(t *testing.T) {
	send := NewTraces()
	fillTestResourceSpansSlice(send.ResourceSpans())
	bytes, err := send.ToOtlpProtoBytes()
	assert.NoError(t, err)

	recv, err := TracesFromOtlpProtoBytes(bytes)
	assert.NoError(t, err)
	assert.EqualValues(t, send, recv)
}

func TestTracesFromInvalidOtlpProtoBytes(t *testing.T) {
	_, err := TracesFromOtlpProtoBytes([]byte{0xFF})
	assert.EqualError(t, err, "unexpected EOF")
}

func TestTracesClone(t *testing.T) {
	traces := NewTraces()
	fillTestResourceSpansSlice(traces.ResourceSpans())
	assert.EqualValues(t, traces, traces.Clone())
}

func BenchmarkTracesClone(b *testing.B) {
	traces := NewTraces()
	fillTestResourceSpansSlice(traces.ResourceSpans())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		clone := traces.Clone()
		if clone.ResourceSpans().Len() != traces.ResourceSpans().Len() {
			b.Fail()
		}
	}
}

func BenchmarkTracesToOtlp(b *testing.B) {
	traces := NewTraces()
	fillTestResourceSpansSlice(traces.ResourceSpans())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf, err := traces.ToOtlpProtoBytes()
		require.NoError(b, err)
		assert.NotEqual(b, 0, len(buf))
	}
}

func BenchmarkTracesFromOtlp(b *testing.B) {
	baseTraces := NewTraces()
	fillTestResourceSpansSlice(baseTraces.ResourceSpans())
	buf, err := baseTraces.ToOtlpProtoBytes()
	require.NoError(b, err)
	assert.NotEqual(b, 0, len(buf))
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		traces, err := TracesFromOtlpProtoBytes(buf)
		require.NoError(b, err)
		assert.Equal(b, baseTraces.ResourceSpans().Len(), traces.ResourceSpans().Len())
	}
}
