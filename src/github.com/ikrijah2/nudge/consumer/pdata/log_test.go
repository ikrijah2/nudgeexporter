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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ikrijah2/nudge/internal"
	otlpcollectorlog "github.com/ikrijah2/nudge/internal/data/protogen/collector/logs/v1"
	otlplogs "github.com/ikrijah2/nudge/internal/data/protogen/logs/v1"
)

func TestLogRecordCount(t *testing.T) {
	md := NewLogs()
	assert.EqualValues(t, 0, md.LogRecordCount())

	md.ResourceLogs().Resize(1)
	assert.EqualValues(t, 0, md.LogRecordCount())

	md.ResourceLogs().At(0).InstrumentationLibraryLogs().Resize(1)
	assert.EqualValues(t, 0, md.LogRecordCount())

	md.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().Resize(1)
	assert.EqualValues(t, 1, md.LogRecordCount())

	rms := md.ResourceLogs()
	rms.Resize(3)
	rms.At(0).InstrumentationLibraryLogs().Resize(1)
	rms.At(0).InstrumentationLibraryLogs().At(0).Logs().Resize(1)
	rms.At(1).InstrumentationLibraryLogs().Resize(1)
	rms.At(2).InstrumentationLibraryLogs().Resize(1)
	rms.At(2).InstrumentationLibraryLogs().At(0).Logs().Resize(5)
	assert.EqualValues(t, 6, md.LogRecordCount())
}

func TestLogRecordCountWithEmpty(t *testing.T) {
	assert.Zero(t, NewLogs().LogRecordCount())
	assert.Zero(t, LogsFromInternalRep(internal.LogsFromOtlp(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{{}},
	})).LogRecordCount())
	assert.Zero(t, LogsFromInternalRep(internal.LogsFromOtlp(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{
			{
				InstrumentationLibraryLogs: []*otlplogs.InstrumentationLibraryLogs{{}},
			},
		},
	})).LogRecordCount())
	assert.Equal(t, 1, LogsFromInternalRep(internal.LogsFromOtlp(&otlpcollectorlog.ExportLogsServiceRequest{
		ResourceLogs: []*otlplogs.ResourceLogs{
			{
				InstrumentationLibraryLogs: []*otlplogs.InstrumentationLibraryLogs{
					{
						Logs: []*otlplogs.LogRecord{{}},
					},
				},
			},
		},
	})).LogRecordCount())
}

func TestToFromLogProto(t *testing.T) {
	wrapper := internal.LogsFromOtlp(&otlpcollectorlog.ExportLogsServiceRequest{})
	ld := LogsFromInternalRep(wrapper)
	assert.EqualValues(t, NewLogs(), ld)
	assert.EqualValues(t, &otlpcollectorlog.ExportLogsServiceRequest{}, ld.orig)
}

func TestLogsToFromOtlpProtoBytes(t *testing.T) {
	send := NewLogs()
	fillTestResourceLogsSlice(send.ResourceLogs())
	bytes, err := send.ToOtlpProtoBytes()
	assert.NoError(t, err)

	recv, err := LogsFromOtlpProtoBytes(bytes)
	assert.NoError(t, err)
	assert.EqualValues(t, send, recv)
}

func TestLogsFromInvalidOtlpProtoBytes(t *testing.T) {
	_, err := LogsFromOtlpProtoBytes([]byte{0xFF})
	assert.EqualError(t, err, "unexpected EOF")
}

func TestLogsClone(t *testing.T) {
	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())
	assert.EqualValues(t, logs, logs.Clone())
}

func BenchmarkLogsClone(b *testing.B) {
	logs := NewLogs()
	fillTestResourceLogsSlice(logs.ResourceLogs())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		clone := logs.Clone()
		if clone.ResourceLogs().Len() != logs.ResourceLogs().Len() {
			b.Fail()
		}
	}
}

func BenchmarkLogsToOtlp(b *testing.B) {
	traces := NewLogs()
	fillTestResourceLogsSlice(traces.ResourceLogs())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buf, err := traces.ToOtlpProtoBytes()
		require.NoError(b, err)
		assert.NotEqual(b, 0, len(buf))
	}
}

func BenchmarkLogsFromOtlp(b *testing.B) {
	baseLogs := NewLogs()
	fillTestResourceLogsSlice(baseLogs.ResourceLogs())
	buf, err := baseLogs.ToOtlpProtoBytes()
	require.NoError(b, err)
	assert.NotEqual(b, 0, len(buf))
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		logs, err := LogsFromOtlpProtoBytes(buf)
		require.NoError(b, err)
		assert.Equal(b, baseLogs.ResourceLogs().Len(), logs.ResourceLogs().Len())
	}
}
