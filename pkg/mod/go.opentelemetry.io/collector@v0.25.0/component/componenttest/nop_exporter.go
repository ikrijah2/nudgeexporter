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

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenthelper"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer/consumertest"
)

// nopExporterFactory is factory for nopExporter.
type nopExporterFactory struct{}

var nopExporterFactoryInstance = &nopExporterFactory{}

// NewNopExporterFactory returns a component.ExporterFactory that constructs nop exporters.
func NewNopExporterFactory() component.ExporterFactory {
	return nopExporterFactoryInstance
}

// Type gets the type of the Exporter config created by this factory.
func (f *nopExporterFactory) Type() config.Type {
	return "nop"
}

// CreateDefaultConfig creates the default configuration for the Exporter.
func (f *nopExporterFactory) CreateDefaultConfig() config.Exporter {
	return &config.ExporterSettings{
		TypeVal: f.Type(),
	}
}

// CreateTracesExporter implements component.ExporterFactory interface.
func (f *nopExporterFactory) CreateTracesExporter(
	_ context.Context,
	_ component.ExporterCreateParams,
	_ config.Exporter,
) (component.TracesExporter, error) {
	return nopExporterInstance, nil
}

// CreateMetricsExporter implements component.ExporterFactory interface.
func (f *nopExporterFactory) CreateMetricsExporter(
	_ context.Context,
	_ component.ExporterCreateParams,
	_ config.Exporter,
) (component.MetricsExporter, error) {
	return nopExporterInstance, nil
}

// CreateLogsExporter implements component.ExporterFactory interface.
func (f *nopExporterFactory) CreateLogsExporter(
	_ context.Context,
	_ component.ExporterCreateParams,
	_ config.Exporter,
) (component.LogsExporter, error) {
	return nopExporterInstance, nil
}

var nopExporterInstance = &nopExporter{
	Component: componenthelper.New(),
	Consumer:  consumertest.NewNop(),
}

// nopExporter stores consumed traces and metrics for testing purposes.
type nopExporter struct {
	component.Component
	consumertest.Consumer
}
