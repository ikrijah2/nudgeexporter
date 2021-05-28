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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/processor/attributesprocessor"
	"go.opentelemetry.io/collector/processor/memorylimiter"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.opentelemetry.io/collector/processor/resourceprocessor"
	"go.opentelemetry.io/collector/processor/spanprocessor"
)

func TestDefaultProcessors(t *testing.T) {
	allFactories, err := Components()
	require.NoError(t, err)

	procFactories := allFactories.Processors

	tests := []struct {
		processor   config.Type
		getConfigFn getProcessorConfigFn
	}{
		{
			processor: "attributes",
			getConfigFn: func() config.Processor {
				cfg := procFactories["attributes"].CreateDefaultConfig().(*attributesprocessor.Config)
				cfg.Actions = []processorhelper.ActionKeyValue{
					{Key: "attribute1", Action: processorhelper.INSERT, Value: 123},
				}
				return cfg
			},
		},
		{
			processor: "batch",
		},
		{
			processor: "filter",
		},
		{
			processor: "memory_limiter",
			getConfigFn: func() config.Processor {
				cfg := procFactories["memory_limiter"].CreateDefaultConfig().(*memorylimiter.Config)
				cfg.CheckInterval = 100 * time.Millisecond
				cfg.MemoryLimitMiB = 1024 * 1024
				return cfg
			},
		},
		{
			processor: "probabilistic_sampler",
		},
		{
			processor: "resource",
			getConfigFn: func() config.Processor {
				cfg := procFactories["resource"].CreateDefaultConfig().(*resourceprocessor.Config)
				cfg.AttributesActions = []processorhelper.ActionKeyValue{
					{Key: "attribute1", Action: processorhelper.INSERT, Value: 123},
				}
				return cfg
			},
		},
		{
			processor: "span",
			getConfigFn: func() config.Processor {
				cfg := procFactories["span"].CreateDefaultConfig().(*spanprocessor.Config)
				cfg.Rename.FromAttributes = []string{"test-key"}
				return cfg
			},
		},
	}

	assert.Equal(t, len(tests), len(procFactories))
	for _, tt := range tests {
		t.Run(string(tt.processor), func(t *testing.T) {
			factory, ok := procFactories[tt.processor]
			require.True(t, ok)
			assert.Equal(t, tt.processor, factory.Type())
			assert.Equal(t, tt.processor, factory.CreateDefaultConfig().Type())

			verifyProcessorLifecycle(t, factory, tt.getConfigFn)
		})
	}
}

// getProcessorConfigFn is used customize the configuration passed to the verification.
// This is used to change ports or provide values required but not provided by the
// default configuration.
type getProcessorConfigFn func() config.Processor

// verifyProcessorLifecycle is used to test if an processor type can handle the typical
// lifecycle of a component. The getConfigFn parameter only need to be specified if
// the test can't be done with the default configuration for the component.
func verifyProcessorLifecycle(t *testing.T, factory component.ProcessorFactory, getConfigFn getProcessorConfigFn) {
	ctx := context.Background()
	host := newAssertNoErrorHost(t)
	processorCreateParams := component.ProcessorCreateParams{
		Logger:               zap.NewNop(),
		ApplicationStartInfo: component.DefaultApplicationStartInfo(),
	}

	if getConfigFn == nil {
		getConfigFn = factory.CreateDefaultConfig
	}

	createFns := []createProcessorFn{
		wrapCreateLogsProc(factory),
		wrapCreateTracesProc(factory),
		wrapCreateMetricsProc(factory),
	}

	for _, createFn := range createFns {
		firstExp, err := createFn(ctx, processorCreateParams, getConfigFn())
		if errors.Is(err, componenterror.ErrDataTypeIsNotSupported) {
			continue
		}
		require.NoError(t, err)
		require.NoError(t, firstExp.Start(ctx, host))
		require.NoError(t, firstExp.Shutdown(ctx))

		secondExp, err := createFn(ctx, processorCreateParams, getConfigFn())
		require.NoError(t, err)
		require.NoError(t, secondExp.Start(ctx, host))
		require.NoError(t, secondExp.Shutdown(ctx))
	}
}

type createProcessorFn func(
	ctx context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
) (component.Processor, error)

func wrapCreateLogsProc(factory component.ProcessorFactory) createProcessorFn {
	return func(ctx context.Context, params component.ProcessorCreateParams, cfg config.Processor) (component.Processor, error) {
		return factory.CreateLogsProcessor(ctx, params, cfg, consumertest.NewNop())
	}
}

func wrapCreateMetricsProc(factory component.ProcessorFactory) createProcessorFn {
	return func(ctx context.Context, params component.ProcessorCreateParams, cfg config.Processor) (component.Processor, error) {
		return factory.CreateMetricsProcessor(ctx, params, cfg, consumertest.NewNop())
	}
}

func wrapCreateTracesProc(factory component.ProcessorFactory) createProcessorFn {
	return func(ctx context.Context, params component.ProcessorCreateParams, cfg config.Processor) (component.Processor, error) {
		return factory.CreateTracesProcessor(ctx, params, cfg, consumertest.NewNop())
	}
}
