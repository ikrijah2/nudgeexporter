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

package zipkinexporter

import (
	"context"
	"errors"
	"time"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/confighttp"
	"github.com/ikrijah2/nudge/exporter/exporterhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "zipkin"

	defaultTimeout = time.Second * 5

	defaultFormat = "json"

	defaultServiceName string = "<missing service name>"
)

// NewFactory creates a factory for Zipkin exporter.
func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		exporterhelper.WithTraces(createTracesExporter))
}

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings: config.ExporterSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		RetrySettings: exporterhelper.DefaultRetrySettings(),
		QueueSettings: exporterhelper.DefaultQueueSettings(),
		HTTPClientSettings: confighttp.HTTPClientSettings{
			Timeout: defaultTimeout,
			// We almost read 0 bytes, so no need to tune ReadBufferSize.
			WriteBufferSize: 512 * 1024,
		},
		Format:             defaultFormat,
		DefaultServiceName: defaultServiceName,
	}
}

func createTracesExporter(
	_ context.Context,
	params component.ExporterCreateParams,
	cfg config.Exporter,
) (component.TracesExporter, error) {
	zc := cfg.(*Config)

	if zc.Endpoint == "" {
		// TODO https://github.com/open-telemetry/opentelemetry-collector/issues/215
		return nil, errors.New("exporter config requires a non-empty 'endpoint'")
	}

	ze, err := createZipkinExporter(zc)
	if err != nil {
		return nil, err
	}
	return exporterhelper.NewTracesExporter(
		zc,
		params.Logger,
		ze.pushTraceData,
		// explicitly disable since we rely on http.Client timeout logic.
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithQueue(zc.QueueSettings),
		exporterhelper.WithRetry(zc.RetrySettings))
}
