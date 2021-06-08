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

package exporterhelper

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/obsreport"
)

// PushTraces is a helper function that is similar to ConsumeTraces but also
// returns the number of dropped spans.
type PushTraces func(ctx context.Context, td pdata.Traces) error

type tracesRequest struct {
	baseRequest
	td     pdata.Traces
	pusher PushTraces
}

func newTracesRequest(ctx context.Context, td pdata.Traces, pusher PushTraces) request {
	return &tracesRequest{
		baseRequest: baseRequest{ctx: ctx},
		td:          td,
		pusher:      pusher,
	}
}

func (req *tracesRequest) onError(err error) request {
	var traceError consumererror.Traces
	if consumererror.AsTraces(err, &traceError) {
		return newTracesRequest(req.ctx, traceError.GetTraces(), req.pusher)
	}
	return req
}

func (req *tracesRequest) export(ctx context.Context) error {
	return req.pusher(ctx, req.td)
}

func (req *tracesRequest) count() int {
	return req.td.SpanCount()
}

type traceExporter struct {
	*baseExporter
	pusher PushTraces
}

func (texp *traceExporter) ConsumeTraces(ctx context.Context, td pdata.Traces) error {
	return texp.sender.send(newTracesRequest(ctx, td, texp.pusher))
}

// NewTracesExporter creates a TracesExporter that records observability metrics and wraps every request with a Span.
func NewTracesExporter(
	cfg config.Exporter,
	logger *zap.Logger,
	pusher PushTraces,
	options ...Option,
) (component.TracesExporter, error) {

	if cfg == nil {
		return nil, errNilConfig
	}

	if logger == nil {
		return nil, errNilLogger
	}

	if pusher == nil {
		return nil, errNilPushTraceData
	}

	be := newBaseExporter(cfg, logger, options...)
	be.wrapConsumerSender(func(nextSender requestSender) requestSender {
		return &tracesExporterWithObservability{
			obsrep: obsreport.NewExporter(
				obsreport.ExporterSettings{
					Level:        configtelemetry.GetMetricsLevelFlagValue(),
					ExporterName: cfg.Name(),
				}),
			nextSender: nextSender,
		}
	})

	return &traceExporter{
		baseExporter: be,
		pusher:       pusher,
	}, nil
}

type tracesExporterWithObservability struct {
	obsrep     *obsreport.Exporter
	nextSender requestSender
}

func (tewo *tracesExporterWithObservability) send(req request) error {
	req.setContext(tewo.obsrep.StartTracesExportOp(req.context()))
	// Forward the data to the next consumer (this pusher is the next).
	err := tewo.nextSender.send(req)
	tewo.obsrep.EndTracesExportOp(req.context(), req.count(), err)
	return err
}
