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

package componenttest

import (
	"context"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/component/componenthelper"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/consumer"
	"github.com/ikrijah2/nudge/consumer/consumertest"
)

// nopProcessorFactory is factory for nopProcessor.
type nopProcessorFactory struct {
	component.BaseProcessorFactory
}

var nopProcessorFactoryInstance = &nopProcessorFactory{}

// NewNopProcessorFactory returns a component.ProcessorFactory that constructs nop exporters.
func NewNopProcessorFactory() component.ProcessorFactory {
	return nopProcessorFactoryInstance
}

// Type gets the type of the Processor config created by this factory.
func (f *nopProcessorFactory) Type() config.Type {
	return "nop"
}

// CreateDefaultConfig creates the default configuration for the Processor.
func (f *nopProcessorFactory) CreateDefaultConfig() config.Processor {
	return &config.ProcessorSettings{
		TypeVal: f.Type(),
	}
}

// CreateTracesProcessor implements component.ProcessorFactory interface.
func (f *nopProcessorFactory) CreateTracesProcessor(
	_ context.Context,
	_ component.ProcessorCreateParams,
	_ config.Processor,
	_ consumer.Traces,
) (component.TracesProcessor, error) {
	return nopProcessorInstance, nil
}

// CreateMetricsProcessor implements component.ProcessorFactory interface.
func (f *nopProcessorFactory) CreateMetricsProcessor(
	_ context.Context,
	_ component.ProcessorCreateParams,
	_ config.Processor,
	_ consumer.Metrics,
) (component.MetricsProcessor, error) {
	return nopProcessorInstance, nil
}

// CreateLogsProcessor implements component.ProcessorFactory interface.
func (f *nopProcessorFactory) CreateLogsProcessor(
	_ context.Context,
	_ component.ProcessorCreateParams,
	_ config.Processor,
	_ consumer.Logs,
) (component.LogsProcessor, error) {
	return nopProcessorInstance, nil
}

var nopProcessorInstance = &nopProcessor{
	Component: componenthelper.New(),
	Consumer:  consumertest.NewNop(),
}

// nopProcessor stores consumed traces and metrics for testing purposes.
type nopProcessor struct {
	component.Component
	consumertest.Consumer
}

func (*nopProcessor) GetCapabilities() component.ProcessorCapabilities {
	return component.ProcessorCapabilities{MutatesConsumedData: false}
}
