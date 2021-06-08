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

package opencensusreceiver

import (
	"context"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/configgrpc"
	"github.com/ikrijah2/nudge/config/confignet"
	"github.com/ikrijah2/nudge/consumer"
	"github.com/ikrijah2/nudge/receiver/receiverhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "opencensus"
)

// NewFactory creates a new OpenCensus receiver factory.
func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		receiverhelper.WithTraces(createTracesReceiver),
		receiverhelper.WithMetrics(createMetricsReceiver))
}

func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings: config.ReceiverSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		GRPCServerSettings: configgrpc.GRPCServerSettings{
			NetAddr: confignet.NetAddr{
				Endpoint:  "0.0.0.0:55678",
				Transport: "tcp",
			},
			// We almost write 0 bytes, so no need to tune WriteBufferSize.
			ReadBufferSize: 512 * 1024,
		},
	}
}

func createTracesReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	cfg config.Receiver,
	nextConsumer consumer.Traces,
) (component.TracesReceiver, error) {
	r, err := createReceiver(cfg)
	if err != nil {
		return nil, err
	}

	r.traceConsumer = nextConsumer

	return r, nil
}

func createMetricsReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	cfg config.Receiver,
	nextConsumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	r, err := createReceiver(cfg)
	if err != nil {
		return nil, err
	}

	r.metricsConsumer = nextConsumer

	return r, nil
}

func createReceiver(cfg config.Receiver) (*ocReceiver, error) {
	rCfg := cfg.(*Config)

	// There must be one receiver for both metrics and traces. We maintain a map of
	// receivers per config.

	// Check to see if there is already a receiver for this config.
	receiver, ok := receivers[rCfg]
	if !ok {
		// Build the configuration options.
		opts, err := rCfg.buildOptions()
		if err != nil {
			return nil, err
		}

		// We don't have a receiver, so create one.
		receiver, err = newOpenCensusReceiver(
			rCfg.Name(), rCfg.NetAddr.Transport, rCfg.NetAddr.Endpoint, nil, nil, opts...)
		if err != nil {
			return nil, err
		}
		// Remember the receiver in the map
		receivers[rCfg] = receiver
	}
	return receiver, nil
}

// This is the map of already created OpenCensus receivers for particular configurations.
// We maintain this map because the Factory is asked trace and metric receivers separately
// when it gets CreateTracesReceiver() and CreateMetricsReceiver() but they must not
// create separate objects, they must use one ocReceiver object per configuration.
var receivers = map[*Config]*ocReceiver{}
