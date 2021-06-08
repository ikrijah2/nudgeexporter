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

// Package defaultcomponents composes the default set of components used by the otel service
package defaultcomponents

import (
	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/consumer/consumererror"
	"github.com/ikrijah2/nudge/exporter/fileexporter"
	"github.com/ikrijah2/nudge/exporter/jaegerexporter"
	"github.com/ikrijah2/nudge/exporter/kafkaexporter"
	"github.com/ikrijah2/nudge/exporter/loggingexporter"
	"github.com/ikrijah2/nudge/exporter/opencensusexporter"
	"github.com/ikrijah2/nudge/exporter/otlpexporter"
	"github.com/ikrijah2/nudge/exporter/otlphttpexporter"
	"github.com/ikrijah2/nudge/exporter/prometheusexporter"
	"github.com/ikrijah2/nudge/exporter/prometheusremotewriteexporter"
	"github.com/ikrijah2/nudge/exporter/zipkinexporter"
	"github.com/ikrijah2/nudge/extension/healthcheckextension"
	"github.com/ikrijah2/nudge/extension/pprofextension"
	"github.com/ikrijah2/nudge/extension/zpagesextension"
	"github.com/ikrijah2/nudge/processor/attributesprocessor"
	"github.com/ikrijah2/nudge/processor/batchprocessor"
	"github.com/ikrijah2/nudge/processor/filterprocessor"
	"github.com/ikrijah2/nudge/processor/memorylimiter"
	"github.com/ikrijah2/nudge/processor/probabilisticsamplerprocessor"
	"github.com/ikrijah2/nudge/processor/resourceprocessor"
	"github.com/ikrijah2/nudge/processor/spanprocessor"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver"
	"github.com/ikrijah2/nudge/receiver/jaegerreceiver"
	"github.com/ikrijah2/nudge/receiver/kafkareceiver"
	"github.com/ikrijah2/nudge/receiver/opencensusreceiver"
	"github.com/ikrijah2/nudge/receiver/otlpreceiver"
	"github.com/ikrijah2/nudge/receiver/prometheusreceiver"
	"github.com/ikrijah2/nudge/receiver/zipkinreceiver"
)

// Components returns the default set of components used by the
// OpenTelemetry collector.
func Components() (
	component.Factories,
	error,
) {
	var errs []error

	extensions, err := component.MakeExtensionFactoryMap(
		healthcheckextension.NewFactory(),
		pprofextension.NewFactory(),
		zpagesextension.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	receivers, err := component.MakeReceiverFactoryMap(
		jaegerreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
		prometheusreceiver.NewFactory(),
		opencensusreceiver.NewFactory(),
		otlpreceiver.NewFactory(),
		hostmetricsreceiver.NewFactory(),
		kafkareceiver.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	exporters, err := component.MakeExporterFactoryMap(
		opencensusexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		prometheusremotewriteexporter.NewFactory(),
		loggingexporter.NewFactory(),
		zipkinexporter.NewFactory(),
		jaegerexporter.NewFactory(),
		fileexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		kafkaexporter.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	processors, err := component.MakeProcessorFactoryMap(
		attributesprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		batchprocessor.NewFactory(),
		memorylimiter.NewFactory(),
		probabilisticsamplerprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		filterprocessor.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	factories := component.Factories{
		Extensions: extensions,
		Receivers:  receivers,
		Processors: processors,
		Exporters:  exporters,
	}

	return factories, consumererror.Combine(errs)
}
