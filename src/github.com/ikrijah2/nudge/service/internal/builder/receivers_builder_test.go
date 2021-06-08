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

package builder

import (
	"context"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/component/componenttest"
	"github.com/ikrijah2/nudge/config/configtest"
	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/internal/testcomponents"
	"github.com/ikrijah2/nudge/internal/testdata"
	"github.com/ikrijah2/nudge/processor/attributesprocessor"
)

type testCase struct {
	name                      string
	receiverName              string
	exporterNames             []string
	spanDuplicationByExporter map[string]int
	hasTraces                 bool
	hasMetrics                bool
}

func TestBuildReceivers(t *testing.T) {
	tests := []testCase{
		{
			name:          "one-exporter",
			receiverName:  "examplereceiver",
			exporterNames: []string{"exampleexporter"},
			hasTraces:     true,
			hasMetrics:    true,
		},
		{
			name:          "multi-exporter",
			receiverName:  "examplereceiver/2",
			exporterNames: []string{"exampleexporter", "exampleexporter/2"},
			hasTraces:     true,
		},
		{
			name:          "multi-metrics-receiver",
			receiverName:  "examplereceiver/3",
			exporterNames: []string{"exampleexporter", "exampleexporter/2"},
			hasTraces:     false,
			hasMetrics:    true,
		},
		{
			name:          "multi-receiver-multi-exporter",
			receiverName:  "examplereceiver/multi",
			exporterNames: []string{"exampleexporter", "exampleexporter/2"},

			// Check pipelines_builder.yaml to understand this case.
			// We have 2 pipelines, one exporting to one exporter, the other
			// exporting to both exporters, so we expect a duplication on
			// one of the exporters, but not on the other.
			spanDuplicationByExporter: map[string]int{
				"exampleexporter": 2, "exampleexporter/2": 1,
			},
			hasTraces: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testReceivers(t, test)
		})
	}
}

func testReceivers(
	t *testing.T,
	test testCase,
) {
	factories, err := testcomponents.ExampleComponents()
	assert.NoError(t, err)

	attrFactory := attributesprocessor.NewFactory()
	factories.Processors[attrFactory.Type()] = attrFactory
	cfg, err := configtest.LoadConfigFile(t, "testdata/pipelines_builder.yaml", factories)
	require.NoError(t, err)

	// Build the pipeline
	allExporters, err := BuildExporters(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, factories.Exporters)
	assert.NoError(t, err)
	pipelineProcessors, err := BuildPipelines(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, allExporters, factories.Processors)
	assert.NoError(t, err)
	receivers, err := BuildReceivers(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, pipelineProcessors, factories.Receivers)

	assert.NoError(t, err)
	require.NotNil(t, receivers)

	receiver := receivers[cfg.Receivers[test.receiverName]]

	// Ensure receiver has its fields correctly populated.
	require.NotNil(t, receiver)

	assert.NotNil(t, receiver.receiver)

	// Compose the list of created exporters.
	var exporters []*builtExporter
	for _, name := range test.exporterNames {
		// Ensure exporter is created.
		exp := allExporters[cfg.Exporters[name]]
		require.NotNil(t, exp)
		exporters = append(exporters, exp)
	}

	// Send TraceData via receiver and verify that all exporters of the pipeline receive it.

	// First check that there are no traces in the exporters yet.
	for _, exporter := range exporters {
		consumer := exporter.getTracesExporter().(*testcomponents.ExampleExporterConsumer)
		require.Equal(t, len(consumer.Traces), 0)
		require.Equal(t, len(consumer.Metrics), 0)
	}

	td := testdata.GenerateTraceDataOneSpan()
	if test.hasTraces {
		traceProducer := receiver.receiver.(*testcomponents.ExampleReceiverProducer)
		assert.NoError(t, traceProducer.ConsumeTraces(context.Background(), td))
	}

	md := testdata.GenerateMetricsOneMetric()
	if test.hasMetrics {
		metricsProducer := receiver.receiver.(*testcomponents.ExampleReceiverProducer)
		assert.NoError(t, metricsProducer.ConsumeMetrics(context.Background(), md))
	}

	// Now verify received data.
	for _, name := range test.exporterNames {
		// Check that the data is received by exporter.
		exporter := allExporters[cfg.Exporters[name]]

		// Validate traces.
		if test.hasTraces {
			var spanDuplicationCount int
			if test.spanDuplicationByExporter != nil {
				spanDuplicationCount = test.spanDuplicationByExporter[name]
			} else {
				spanDuplicationCount = 1
			}

			traceConsumer := exporter.getTracesExporter().(*testcomponents.ExampleExporterConsumer)
			require.Equal(t, spanDuplicationCount, len(traceConsumer.Traces))

			for i := 0; i < spanDuplicationCount; i++ {
				assert.EqualValues(t, td, traceConsumer.Traces[i])
			}
		}

		// Validate metrics.
		if test.hasMetrics {
			metricsConsumer := exporter.getMetricExporter().(*testcomponents.ExampleExporterConsumer)
			require.Equal(t, 1, len(metricsConsumer.Metrics))
			assert.EqualValues(t, md, metricsConsumer.Metrics[0])
		}
	}
}

func TestBuildReceivers_BuildCustom(t *testing.T) {
	factories := createTestFactories()

	tests := []struct {
		dataType   string
		shouldFail bool
	}{
		{
			dataType:   "logs",
			shouldFail: false,
		},
		{
			dataType:   "nosuchdatatype",
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.dataType, func(t *testing.T) {
			dataType := test.dataType

			cfg := createExampleConfig(dataType)

			// Build the pipeline
			allExporters, err := BuildExporters(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, factories.Exporters)
			if test.shouldFail {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			pipelineProcessors, err := BuildPipelines(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, allExporters, factories.Processors)
			assert.NoError(t, err)
			receivers, err := BuildReceivers(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, pipelineProcessors, factories.Receivers)

			assert.NoError(t, err)
			require.NotNil(t, receivers)

			receiver := receivers[cfg.Receivers["examplereceiver"]]

			// Ensure receiver has its fields correctly populated.
			require.NotNil(t, receiver)

			assert.NotNil(t, receiver.receiver)

			// Compose the list of created exporters.
			exporterNames := []string{"exampleexporter"}
			var exporters []*builtExporter
			for _, name := range exporterNames {
				// Ensure exporter is created.
				exp := allExporters[cfg.Exporters[name]]
				require.NotNil(t, exp)
				exporters = append(exporters, exp)
			}

			// Send Data via receiver and verify that all exporters of the pipeline receive it.

			// First check that there are no traces in the exporters yet.
			for _, exporter := range exporters {
				consumer := exporter.getLogExporter().(*testcomponents.ExampleExporterConsumer)
				require.Equal(t, len(consumer.Logs), 0)
			}

			// Send one data.
			log := pdata.Logs{}
			producer := receiver.receiver.(*testcomponents.ExampleReceiverProducer)
			require.NoError(t, producer.ConsumeLogs(context.Background(), log))

			// Now verify received data.
			for _, name := range exporterNames {
				// Check that the data is received by exporter.
				exporter := allExporters[cfg.Exporters[name]]

				// Validate exported data.
				consumer := exporter.getLogExporter().(*testcomponents.ExampleExporterConsumer)
				require.Equal(t, 1, len(consumer.Logs))
				assert.EqualValues(t, log, consumer.Logs[0])
			}
		})
	}
}

func TestBuildReceivers_StartAll(t *testing.T) {
	receivers := make(Receivers)
	rcvCfg := &testcomponents.ExampleReceiver{}
	receiver := &testcomponents.ExampleReceiverProducer{}

	receivers[rcvCfg] = &builtReceiver{
		logger:   zap.NewNop(),
		receiver: receiver,
	}

	assert.False(t, receiver.Started)

	err := receivers.StartAll(context.Background(), componenttest.NewNopHost())
	assert.NoError(t, err)

	assert.True(t, receiver.Started)
}

func TestBuildReceivers_StopAll(t *testing.T) {
	receivers := make(Receivers)
	rcvCfg := &testcomponents.ExampleReceiver{}
	receiver := &testcomponents.ExampleReceiverProducer{}

	receivers[rcvCfg] = &builtReceiver{
		logger:   zap.NewNop(),
		receiver: receiver,
	}

	assert.False(t, receiver.Stopped)

	assert.NoError(t, receivers.ShutdownAll(context.Background()))

	assert.True(t, receiver.Stopped)
}

func TestBuildReceivers_Unused(t *testing.T) {
	factories, err := testcomponents.ExampleComponents()
	assert.NoError(t, err)

	cfg, err := configtest.LoadConfigFile(t, "testdata/unused_receiver.yaml", factories)
	assert.NoError(t, err)

	// Build the pipeline
	allExporters, err := BuildExporters(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, factories.Exporters)
	assert.NoError(t, err)
	pipelineProcessors, err := BuildPipelines(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, allExporters, factories.Processors)
	assert.NoError(t, err)
	receivers, err := BuildReceivers(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, pipelineProcessors, factories.Receivers)
	assert.NoError(t, err)
	assert.NotNil(t, receivers)

	assert.NoError(t, receivers.StartAll(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, receivers.ShutdownAll(context.Background()))
}

func TestBuildReceivers_NotSupportedDataType(t *testing.T) {
	factories := createTestFactories()

	tests := []struct {
		configFile string
	}{
		{
			configFile: "not_supported_receiver_logs.yaml",
		},
		{
			configFile: "not_supported_receiver_metrics.yaml",
		},
		{
			configFile: "not_supported_receiver_traces.yaml",
		},
	}

	for _, test := range tests {
		t.Run(test.configFile, func(t *testing.T) {

			cfg, err := configtest.LoadConfigFile(t, path.Join("testdata", test.configFile), factories)
			assert.NoError(t, err)
			require.NotNil(t, cfg)

			allExporters, err := BuildExporters(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, factories.Exporters)
			assert.NoError(t, err)

			pipelineProcessors, err := BuildPipelines(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, allExporters, factories.Processors)
			assert.NoError(t, err)

			receivers, err := BuildReceivers(zap.NewNop(), component.DefaultApplicationStartInfo(), cfg, pipelineProcessors, factories.Receivers)
			assert.Error(t, err)
			assert.Zero(t, len(receivers))
		})
	}
}
