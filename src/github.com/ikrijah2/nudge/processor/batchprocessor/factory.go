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

package batchprocessor

import (
	"context"
	"time"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/configtelemetry"
	"github.com/ikrijah2/nudge/consumer"
	"github.com/ikrijah2/nudge/processor/processorhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "batch"

	defaultSendBatchSize = uint32(8192)
	defaultTimeout       = 200 * time.Millisecond
)

// NewFactory returns a new factory for the Batch processor.
func NewFactory() component.ProcessorFactory {
	return processorhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		processorhelper.WithTraces(createTracesProcessor),
		processorhelper.WithMetrics(createMetricsProcessor),
		processorhelper.WithLogs(createLogsProcessor))
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(typeStr),
		SendBatchSize:     defaultSendBatchSize,
		Timeout:           defaultTimeout,
	}
}

func createTracesProcessor(
	_ context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextConsumer consumer.Traces,
) (component.TracesProcessor, error) {
	oCfg := cfg.(*Config)
	level := configtelemetry.GetMetricsLevelFlagValue()
	return newBatchTracesProcessor(params, nextConsumer, oCfg, level), nil
}

func createMetricsProcessor(
	_ context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextConsumer consumer.Metrics,
) (component.MetricsProcessor, error) {
	oCfg := cfg.(*Config)
	level := configtelemetry.GetMetricsLevelFlagValue()
	return newBatchMetricsProcessor(params, nextConsumer, oCfg, level), nil
}

func createLogsProcessor(
	_ context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextConsumer consumer.Logs,
) (component.LogsProcessor, error) {
	oCfg := cfg.(*Config)
	level := configtelemetry.GetMetricsLevelFlagValue()
	return newBatchLogsProcessor(params, nextConsumer, oCfg, level), nil
}
