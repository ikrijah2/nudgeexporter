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

package defaultcomponents

import (
	"context"
	"errors"
	"testing"

	promconfig "github.com/prometheus/prometheus/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/component/componenterror"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/consumer/consumertest"
	"github.com/ikrijah2/nudge/receiver/prometheusreceiver"
)

func TestDefaultReceivers(t *testing.T) {
	allFactories, err := Components()
	assert.NoError(t, err)

	rcvrFactories := allFactories.Receivers

	tests := []struct {
		receiver     config.Type
		skipLifecyle bool
		getConfigFn  getReceiverConfigFn
	}{
		{
			receiver: "hostmetrics",
		},
		{
			receiver: "jaeger",
		},
		{
			receiver:     "kafka",
			skipLifecyle: true, // TODO: It needs access to internals to successful start.
		},
		{
			receiver:     "opencensus",
			skipLifecyle: true, // TODO: Usage of CMux doesn't allow proper shutdown.
		},
		{
			receiver: "otlp",
		},
		{
			receiver: "prometheus",
			getConfigFn: func() config.Receiver {
				cfg := rcvrFactories["prometheus"].CreateDefaultConfig().(*prometheusreceiver.Config)
				cfg.PrometheusConfig = &promconfig.Config{
					ScrapeConfigs: []*promconfig.ScrapeConfig{
						{JobName: "test"},
					},
				}
				return cfg
			},
		},
		{
			receiver: "zipkin",
		},
	}

	assert.Equal(t, len(tests), len(rcvrFactories))
	for _, tt := range tests {
		t.Run(string(tt.receiver), func(t *testing.T) {
			factory, ok := rcvrFactories[tt.receiver]
			require.True(t, ok)
			assert.Equal(t, tt.receiver, factory.Type())
			assert.Equal(t, tt.receiver, factory.CreateDefaultConfig().Type())

			if tt.skipLifecyle {
				t.Log("Skipping lifecycle test", tt.receiver)
				return
			}

			verifyReceiverLifecycle(t, factory, tt.getConfigFn)
		})
	}
}

// getReceiverConfigFn is used customize the configuration passed to the verification.
// This is used to change ports or provide values required but not provided by the
// default configuration.
type getReceiverConfigFn func() config.Receiver

// verifyReceiverLifecycle is used to test if a receiver type can handle the typical
// lifecycle of a component. The getConfigFn parameter only need to be specified if
// the test can't be done with the default configuration for the component.
func verifyReceiverLifecycle(t *testing.T, factory component.ReceiverFactory, getConfigFn getReceiverConfigFn) {
	ctx := context.Background()
	host := newAssertNoErrorHost(t)
	receiverCreateParams := component.ReceiverCreateParams{
		Logger:               zap.NewNop(),
		ApplicationStartInfo: component.DefaultApplicationStartInfo(),
	}

	if getConfigFn == nil {
		getConfigFn = factory.CreateDefaultConfig
	}

	createFns := []createReceiverFn{
		wrapCreateLogsRcvr(factory),
		wrapCreateTracesRcvr(factory),
		wrapCreateMetricsRcvr(factory),
	}

	for _, createFn := range createFns {
		firstRcvr, err := createFn(ctx, receiverCreateParams, getConfigFn())
		if errors.Is(err, componenterror.ErrDataTypeIsNotSupported) {
			continue
		}
		require.NoError(t, err)
		require.NoError(t, firstRcvr.Start(ctx, host))
		require.NoError(t, firstRcvr.Shutdown(ctx))

		secondRcvr, err := createFn(ctx, receiverCreateParams, getConfigFn())
		require.NoError(t, err)
		require.NoError(t, secondRcvr.Start(ctx, host))
		require.NoError(t, secondRcvr.Shutdown(ctx))
	}
}

// assertNoErrorHost implements a component.Host that asserts that there were no errors.
type createReceiverFn func(
	ctx context.Context,
	params component.ReceiverCreateParams,
	cfg config.Receiver,
) (component.Receiver, error)

func wrapCreateLogsRcvr(factory component.ReceiverFactory) createReceiverFn {
	return func(ctx context.Context, params component.ReceiverCreateParams, cfg config.Receiver) (component.Receiver, error) {
		return factory.CreateLogsReceiver(ctx, params, cfg, consumertest.NewNop())
	}
}

func wrapCreateMetricsRcvr(factory component.ReceiverFactory) createReceiverFn {
	return func(ctx context.Context, params component.ReceiverCreateParams, cfg config.Receiver) (component.Receiver, error) {
		return factory.CreateMetricsReceiver(ctx, params, cfg, consumertest.NewNop())
	}
}

func wrapCreateTracesRcvr(factory component.ReceiverFactory) createReceiverFn {
	return func(ctx context.Context, params component.ReceiverCreateParams, cfg config.Receiver) (component.Receiver, error) {
		return factory.CreateTracesReceiver(ctx, params, cfg, consumertest.NewNop())
	}
}
