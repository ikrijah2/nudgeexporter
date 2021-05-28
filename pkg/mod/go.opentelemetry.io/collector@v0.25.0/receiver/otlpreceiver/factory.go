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

package otlpreceiver

import (
	"context"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "otlp"

	defaultGRPCEndpoint = "0.0.0.0:4317"
	defaultHTTPEndpoint = "0.0.0.0:55681"
	legacyGRPCEndpoint  = "0.0.0.0:55680"
)

// NewFactory creates a new OTLP receiver factory.
func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		receiverhelper.WithTraces(createTracesReceiver),
		receiverhelper.WithMetrics(createMetricsReceiver),
		receiverhelper.WithLogs(createLogReceiver))
}

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings: config.ReceiverSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		Protocols: Protocols{
			GRPC: &configgrpc.GRPCServerSettings{
				NetAddr: confignet.NetAddr{
					Endpoint:  defaultGRPCEndpoint,
					Transport: "tcp",
				},
				// We almost write 0 bytes, so no need to tune WriteBufferSize.
				ReadBufferSize: 512 * 1024,
			},
			HTTP: &confighttp.HTTPServerSettings{
				Endpoint: defaultHTTPEndpoint,
			},
		},
	}
}

// CreateTracesReceiver creates a  trace receiver based on provided config.
func createTracesReceiver(
	ctx context.Context,
	params component.ReceiverCreateParams,
	cfg config.Receiver,
	nextConsumer consumer.Traces,
) (component.TracesReceiver, error) {
	r, err := createReceiver(cfg, params.Logger)
	if err != nil {
		return nil, err
	}
	if err = r.registerTraceConsumer(ctx, nextConsumer); err != nil {
		return nil, err
	}
	return r, nil
}

// CreateMetricsReceiver creates a metrics receiver based on provided config.
func createMetricsReceiver(
	ctx context.Context,
	params component.ReceiverCreateParams,
	cfg config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	r, err := createReceiver(cfg, params.Logger)
	if err != nil {
		return nil, err
	}
	if err = r.registerMetricsConsumer(ctx, consumer); err != nil {
		return nil, err
	}
	return r, nil
}

// CreateLogReceiver creates a log receiver based on provided config.
func createLogReceiver(
	ctx context.Context,
	params component.ReceiverCreateParams,
	cfg config.Receiver,
	consumer consumer.Logs,
) (component.LogsReceiver, error) {
	r, err := createReceiver(cfg, params.Logger)
	if err != nil {
		return nil, err
	}
	if err = r.registerLogsConsumer(ctx, consumer); err != nil {
		return nil, err
	}
	return r, nil
}

func createReceiver(cfg config.Receiver, logger *zap.Logger) (*otlpReceiver, error) {
	rCfg := cfg.(*Config)

	// There must be one receiver for both metrics and traces. We maintain a map of
	// receivers per config.

	// Check to see if there is already a receiver for this config.
	receiver, ok := receivers[rCfg]
	if !ok {
		var err error
		// We don't have a receiver, so create one.
		receiver, err = newOtlpReceiver(rCfg, logger)
		if err != nil {
			return nil, err
		}
		// Remember the receiver in the map
		receivers[rCfg] = receiver
	}
	return receiver, nil
}

// This is the map of already created OTLP receivers for particular configurations.
// We maintain this map because the Factory is asked trace and metric receivers separately
// when it gets CreateTracesReceiver() and CreateMetricsReceiver() but they must not
// create separate objects, they must use one otlpReceiver object per configuration.
// When the receiver is shutdown it should be removed from this map so the same configuration
// can be recreated successfully.
var receivers = map[*Config]*otlpReceiver{}
