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

package opencensusexporter

import (
	"context"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/configgrpc"
	"github.com/ikrijah2/nudge/exporter/exporterhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "opencensus"
)

// NewFactory creates a factory for OTLP exporter.
func NewFactory() component.ExporterFactory {
	return exporterhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		exporterhelper.WithTraces(createTracesExporter),
		exporterhelper.WithMetrics(createMetricsExporter))
}

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings: config.NewExporterSettings(typeStr),
		GRPCClientSettings: configgrpc.GRPCClientSettings{
			Headers: map[string]string{},
			// We almost read 0 bytes, so no need to tune ReadBufferSize.
			WriteBufferSize: 512 * 1024,
		},
		NumWorkers: 2,
	}
}

func createTracesExporter(ctx context.Context, params component.ExporterCreateParams, cfg config.Exporter) (component.TracesExporter, error) {
	oCfg := cfg.(*Config)
	oce, err := newTracesExporter(ctx, oCfg)
	if err != nil {
		return nil, err
	}

	return exporterhelper.NewTracesExporter(
		cfg,
		params.Logger,
		oce.pushTraceData,
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithShutdown(oce.shutdown))
}

func createMetricsExporter(ctx context.Context, params component.ExporterCreateParams, cfg config.Exporter) (component.MetricsExporter, error) {
	oCfg := cfg.(*Config)
	oce, err := newMetricsExporter(ctx, oCfg)
	if err != nil {
		return nil, err
	}

	return exporterhelper.NewMetricsExporter(
		cfg,
		params.Logger,
		oce.pushMetricsData,
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithShutdown(oce.shutdown))
}
