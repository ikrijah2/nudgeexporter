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
package fileexporter

import (
	"context"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/internal"
	collectorlogs "github.com/ikrijah2/nudge/internal/data/protogen/collector/logs/v1"
	collectormetrics "github.com/ikrijah2/nudge/internal/data/protogen/collector/metrics/v1"
	collectortrace "github.com/ikrijah2/nudge/internal/data/protogen/collector/trace/v1"
	otlpcommon "github.com/ikrijah2/nudge/internal/data/protogen/common/v1"
	logspb "github.com/ikrijah2/nudge/internal/data/protogen/logs/v1"
	otresourcepb "github.com/ikrijah2/nudge/internal/data/protogen/resource/v1"
	"github.com/ikrijah2/nudge/internal/testdata"
	"github.com/ikrijah2/nudge/testutil"
)

func TestFileTracesExporterNoErrors(t *testing.T) {
	mf := &testutil.LimitedWriter{}
	lte := &fileExporter{file: mf}
	require.NotNil(t, lte)

	td := testdata.GenerateTraceDataTwoSpansSameResource()

	assert.NoError(t, lte.ConsumeTraces(context.Background(), td))
	assert.NoError(t, lte.Shutdown(context.Background()))

	var unmarshaler = &jsonpb.Unmarshaler{}
	got := &collectortrace.ExportTraceServiceRequest{}
	assert.NoError(t, unmarshaler.Unmarshal(mf, got))

	assert.EqualValues(t, internal.TracesToOtlp(td.InternalRep()), got)
}

func TestFileMetricsExporterNoErrors(t *testing.T) {
	mf := &testutil.LimitedWriter{}
	lme := &fileExporter{file: mf}
	require.NotNil(t, lme)

	md := testdata.GenerateMetricsTwoMetrics()
	assert.NoError(t, lme.ConsumeMetrics(context.Background(), md))
	assert.NoError(t, lme.Shutdown(context.Background()))

	var unmarshaler = &jsonpb.Unmarshaler{}
	j := &collectormetrics.ExportMetricsServiceRequest{}
	assert.NoError(t, unmarshaler.Unmarshal(mf, j))

	assert.EqualValues(t, internal.MetricsToOtlp(md.InternalRep()), j)
}

func TestFileLogsExporterNoErrors(t *testing.T) {
	mf := &testutil.LimitedWriter{}
	exporter := &fileExporter{file: mf}
	require.NotNil(t, exporter)

	now := time.Now()
	otlp := &collectorlogs.ExportLogsServiceRequest{
		ResourceLogs: []*logspb.ResourceLogs{
			{
				Resource: otresourcepb.Resource{
					Attributes: []otlpcommon.KeyValue{
						{
							Key:   "attr1",
							Value: otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "value1"}},
						},
					},
				},
				InstrumentationLibraryLogs: []*logspb.InstrumentationLibraryLogs{
					{
						Logs: []*logspb.LogRecord{
							{
								TimeUnixNano: uint64(now.UnixNano()),
								Name:         "logA",
							},
							{
								TimeUnixNano: uint64(now.UnixNano()),
								Name:         "logB",
							},
						},
					},
				},
			},
			{
				Resource: otresourcepb.Resource{
					Attributes: []otlpcommon.KeyValue{
						{
							Key:   "attr2",
							Value: otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "value2"}},
						},
					},
				},
				InstrumentationLibraryLogs: []*logspb.InstrumentationLibraryLogs{
					{
						Logs: []*logspb.LogRecord{
							{
								TimeUnixNano: uint64(now.UnixNano()),
								Name:         "logC",
							},
						},
					},
				},
			},
		},
	}
	assert.NoError(t, exporter.ConsumeLogs(context.Background(), pdata.LogsFromInternalRep(internal.LogsFromOtlp(otlp))))
	assert.NoError(t, exporter.Shutdown(context.Background()))

	var unmarshaler = &jsonpb.Unmarshaler{}
	var j collectorlogs.ExportLogsServiceRequest

	assert.NoError(t, unmarshaler.Unmarshal(mf, &j))
	assert.EqualValues(t, otlp.ResourceLogs, j.ResourceLogs)
}

func TestFileLogsExporterErrors(t *testing.T) {

	now := time.Now()
	otlp := &collectorlogs.ExportLogsServiceRequest{
		ResourceLogs: []*logspb.ResourceLogs{
			{
				Resource: otresourcepb.Resource{
					Attributes: []otlpcommon.KeyValue{
						{
							Key:   "attr1",
							Value: otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "value1"}},
						},
					},
				},
				InstrumentationLibraryLogs: []*logspb.InstrumentationLibraryLogs{
					{
						Logs: []*logspb.LogRecord{
							{
								TimeUnixNano: uint64(now.UnixNano()),
								Name:         "logA",
							},
							{
								TimeUnixNano: uint64(now.UnixNano()),
								Name:         "logB",
							},
						},
					},
				},
			},
			{
				Resource: otresourcepb.Resource{
					Attributes: []otlpcommon.KeyValue{
						{
							Key:   "attr2",
							Value: otlpcommon.AnyValue{Value: &otlpcommon.AnyValue_StringValue{StringValue: "value2"}},
						},
					},
				},
				InstrumentationLibraryLogs: []*logspb.InstrumentationLibraryLogs{
					{
						Logs: []*logspb.LogRecord{
							{
								TimeUnixNano: uint64(now.UnixNano()),
								Name:         "logC",
							},
						},
					},
				},
			},
		},
	}

	cases := []struct {
		Name   string
		MaxLen int
	}{
		{
			Name:   "opening",
			MaxLen: 1,
		},
		{
			Name:   "resource",
			MaxLen: 16,
		},
		{
			Name:   "log_start",
			MaxLen: 78,
		},
		{
			Name:   "logs",
			MaxLen: 128,
		},
	}

	for i := range cases {
		maxLen := cases[i].MaxLen
		t.Run(cases[i].Name, func(t *testing.T) {
			mf := &testutil.LimitedWriter{
				MaxLen: maxLen,
			}
			exporter := &fileExporter{file: mf}
			require.NotNil(t, exporter)

			assert.Error(t, exporter.ConsumeLogs(context.Background(), pdata.LogsFromInternalRep(internal.LogsFromOtlp(otlp))))
			assert.NoError(t, exporter.Shutdown(context.Background()))
		})
	}
}
