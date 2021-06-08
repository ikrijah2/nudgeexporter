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

package prometheusremotewriteexporter

import (
	"context"
	"errors"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/confighttp"
	"github.com/ikrijah2/nudge/exporter/exporterhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "prometheusremotewrite"
)

// NewFactory creates a new Prometheus Remote Write exporter.
func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		exporterhelper.WithMetrics(createMetricsExporter))
}

func createMetricsExporter(_ context.Context, params component.ExporterCreateParams,
	cfg config.Exporter) (component.MetricsExporter, error) {

	prwCfg, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("invalid configuration")
	}

	client, err := prwCfg.HTTPClientSettings.ToClient()
	if err != nil {
		return nil, err
	}

	prwe, err := NewPrwExporter(prwCfg.Namespace, prwCfg.HTTPClientSettings.Endpoint, client, prwCfg.ExternalLabels)
	if err != nil {
		return nil, err
	}

	// Don't support the queue.
	// See https://github.com/open-telemetry/opentelemetry-collector/issues/2949.
	// Prometheus remote write samples needs to be in chronological
	// order for each timeseries. If we shard the incoming metrics
	// without considering this limitation, we experience
	// "out of order samples" errors.
	prwexp, err := exporterhelper.NewMetricsExporter(
		cfg,
		params.Logger,
		prwe.PushMetrics,
		exporterhelper.WithTimeout(prwCfg.TimeoutSettings),
		exporterhelper.WithRetry(prwCfg.RetrySettings),
		exporterhelper.WithShutdown(prwe.Shutdown),
	)

	return prwexp, err
}

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings: config.NewExporterSettings(typeStr),
		Namespace:        "",
		ExternalLabels:   map[string]string{},
		TimeoutSettings:  exporterhelper.DefaultTimeoutSettings(),
		RetrySettings:    exporterhelper.DefaultRetrySettings(),
		HTTPClientSettings: confighttp.HTTPClientSettings{
			Endpoint: "http://some.url:9411/api/prom/push",
			// We almost read 0 bytes, so no need to tune ReadBufferSize.
			ReadBufferSize:  0,
			WriteBufferSize: 512 * 1024,
			Timeout:         exporterhelper.DefaultTimeoutSettings().Timeout,
			Headers:         map[string]string{},
		},
	}
}
