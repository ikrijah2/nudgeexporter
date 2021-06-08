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

package obsreporttest_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ikrijah2/nudge/config/configtelemetry"
	"github.com/ikrijah2/nudge/obsreport"
	"github.com/ikrijah2/nudge/obsreport/obsreporttest"
)

const (
	exporter  = "fakeExporter"
	receiver  = "fakeReicever"
	transport = "fakeTransport"
	format    = "fakeFormat"
)

func TestCheckReceiverTracesViews(t *testing.T) {
	doneFn, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer doneFn()

	receiverCtx := obsreport.ReceiverContext(context.Background(), receiver, transport)
	ctx := obsreport.StartTraceDataReceiveOp(receiverCtx, receiver, transport)
	assert.NotNil(t, ctx)
	obsreport.EndTraceDataReceiveOp(
		ctx,
		format,
		7,
		nil)

	obsreporttest.CheckReceiverTraces(t, receiver, transport, 7, 0)
}

func TestCheckReceiverMetricsViews(t *testing.T) {
	doneFn, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer doneFn()

	receiverCtx := obsreport.ReceiverContext(context.Background(), receiver, transport)
	ctx := obsreport.StartMetricsReceiveOp(receiverCtx, receiver, transport)
	assert.NotNil(t, ctx)
	obsreport.EndMetricsReceiveOp(ctx, format, 7, nil)

	obsreporttest.CheckReceiverMetrics(t, receiver, transport, 7, 0)
}

func TestCheckReceiverLogsViews(t *testing.T) {
	doneFn, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer doneFn()

	receiverCtx := obsreport.ReceiverContext(context.Background(), receiver, transport)
	ctx := obsreport.StartLogsReceiveOp(receiverCtx, receiver, transport)
	assert.NotNil(t, ctx)
	obsreport.EndLogsReceiveOp(ctx, format, 7, nil)

	obsreporttest.CheckReceiverLogs(t, receiver, transport, 7, 0)
}

func TestCheckExporterTracesViews(t *testing.T) {
	doneFn, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer doneFn()

	obsrep := obsreport.NewExporter(obsreport.ExporterSettings{
		Level:        configtelemetry.LevelNormal,
		ExporterName: exporter,
	})
	ctx := obsrep.StartTracesExportOp(context.Background())
	assert.NotNil(t, ctx)

	obsrep.EndTracesExportOp(ctx, 7, nil)

	obsreporttest.CheckExporterTraces(t, exporter, 7, 0)
}

func TestCheckExporterMetricsViews(t *testing.T) {
	doneFn, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer doneFn()

	obsrep := obsreport.NewExporter(obsreport.ExporterSettings{
		Level:        configtelemetry.LevelNormal,
		ExporterName: exporter,
	})
	ctx := obsrep.StartMetricsExportOp(context.Background())
	assert.NotNil(t, ctx)

	obsrep.EndMetricsExportOp(ctx, 7, nil)

	obsreporttest.CheckExporterMetrics(t, exporter, 7, 0)
}

func TestCheckExporterLogsViews(t *testing.T) {
	doneFn, err := obsreporttest.SetupRecordedMetricsTest()
	require.NoError(t, err)
	defer doneFn()

	obsrep := obsreport.NewExporter(obsreport.ExporterSettings{
		Level:        configtelemetry.LevelNormal,
		ExporterName: exporter,
	})
	ctx := obsrep.StartLogsExportOp(context.Background())
	assert.NotNil(t, ctx)
	obsrep.EndLogsExportOp(ctx, 7, nil)

	obsreporttest.CheckExporterLogs(t, exporter, 7, 0)
}
