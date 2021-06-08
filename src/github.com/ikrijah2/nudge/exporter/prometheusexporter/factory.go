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

package prometheusexporter

import (
	"context"
	"time"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/exporter/exporterhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "prometheus"
)

// NewFactory creates a new Prometheus exporter factory.
func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		exporterhelper.WithMetrics(createMetricsExporter))
}

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings: config.NewExporterSettings(typeStr),
		ConstLabels:      map[string]string{},
		SendTimestamps:   false,
		MetricExpiration: time.Minute * 5,
	}
}

func createMetricsExporter(
	_ context.Context,
	params component.ExporterCreateParams,
	cfg config.Exporter,
) (component.MetricsExporter, error) {
	pcfg := cfg.(*Config)

	return newPrometheusExporter(pcfg, params.Logger)
}
