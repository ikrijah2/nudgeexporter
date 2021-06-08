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

// Package collector handles the command-line, configuration, and runs the OC collector.
package service

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"testing"

	"github.com/prometheus/common/expfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/service/defaultcomponents"
	"github.com/ikrijah2/nudge/service/internal/builder"
	"github.com/ikrijah2/nudge/service/parserprovider"
	"github.com/ikrijah2/nudge/testutil"
)

func TestApplication_Start(t *testing.T) {
	factories, err := defaultcomponents.Components()
	require.NoError(t, err)

	loggingHookCalled := false
	hook := func(entry zapcore.Entry) error {
		loggingHookCalled = true
		return nil
	}

	app, err := New(Parameters{Factories: factories, ApplicationStartInfo: component.DefaultApplicationStartInfo(), LoggingOptions: []zap.Option{zap.Hooks(hook)}})
	require.NoError(t, err)
	assert.Equal(t, app.rootCmd, app.Command())

	const testPrefix = "a_test"
	metricsPort := testutil.GetAvailablePort(t)
	healthCheckEndpoint := testutil.GetAvailableLocalAddress(t)
	app.rootCmd.SetArgs([]string{
		"--config=testdata/otelcol-config.yaml",
		"--metrics-addr=localhost:" + strconv.FormatUint(uint64(metricsPort), 10),
		"--metrics-prefix=" + testPrefix,
		"--set=extensions.health_check.endpoint=" + healthCheckEndpoint,
	})

	appDone := make(chan struct{})
	go func() {
		defer close(appDone)
		assert.NoError(t, app.Run())
	}()

	assert.Equal(t, Starting, <-app.GetStateChannel())
	assert.Equal(t, Running, <-app.GetStateChannel())
	require.True(t, isAppAvailable(t, "http://"+healthCheckEndpoint))
	assert.Equal(t, app.logger, app.GetLogger())
	assert.True(t, loggingHookCalled)

	// All labels added to all collector metrics by default are listed below.
	// These labels are hard coded here in order to avoid inadvertent changes:
	// at this point changing labels should be treated as a breaking changing
	// and requires a good justification. The reason is that changes to metric
	// names or labels can break alerting, dashboards, etc that are used to
	// monitor the Collector in production deployments.
	mandatoryLabels := []string{
		"service_instance_id",
	}
	assertMetrics(t, testPrefix, metricsPort, mandatoryLabels)

	// Trigger another configuration load.
	require.NoError(t, app.reloadService(context.Background()))
	require.True(t, isAppAvailable(t, "http://"+healthCheckEndpoint))

	app.signalsChannel <- syscall.SIGTERM
	<-appDone
	assert.Equal(t, Closing, <-app.GetStateChannel())
	assert.Equal(t, Closed, <-app.GetStateChannel())
}

type mockAppTelemetry struct{}

func (tel *mockAppTelemetry) init(chan<- error, uint64, *zap.Logger) error {
	return nil
}

func (tel *mockAppTelemetry) shutdown() error {
	return errors.New("err1")
}

func TestApplication_ReportError(t *testing.T) {
	// use a mock AppTelemetry struct to return an error on shutdown
	preservedAppTelemetry := applicationTelemetry
	applicationTelemetry = &mockAppTelemetry{}
	defer func() { applicationTelemetry = preservedAppTelemetry }()

	factories, err := defaultcomponents.Components()
	require.NoError(t, err)

	app, err := New(Parameters{Factories: factories, ApplicationStartInfo: component.DefaultApplicationStartInfo()})
	require.NoError(t, err)

	app.rootCmd.SetArgs([]string{"--config=testdata/otelcol-config-minimal.yaml"})

	appDone := make(chan struct{})
	go func() {
		defer close(appDone)
		assert.EqualError(t, app.Run(), "failed to shutdown application telemetry: err1")
	}()

	assert.Equal(t, Starting, <-app.GetStateChannel())
	assert.Equal(t, Running, <-app.GetStateChannel())
	app.service.ReportFatalError(errors.New("err2"))
	<-appDone
	assert.Equal(t, Closing, <-app.GetStateChannel())
	assert.Equal(t, Closed, <-app.GetStateChannel())
}

func TestApplication_StartAsGoRoutine(t *testing.T) {
	factories, err := defaultcomponents.Components()
	require.NoError(t, err)

	params := Parameters{
		ApplicationStartInfo: component.DefaultApplicationStartInfo(),
		ParserProvider:       new(minimalParserLoader),
		Factories:            factories,
	}
	app, err := New(params)
	require.NoError(t, err)
	app.Command().SetArgs([]string{
		"--metrics-level=NONE",
	})

	appDone := make(chan struct{})
	go func() {
		defer close(appDone)
		appErr := app.Run()
		if appErr != nil {
			err = appErr
		}
	}()

	assert.Equal(t, Starting, <-app.GetStateChannel())
	assert.Equal(t, Running, <-app.GetStateChannel())

	app.Shutdown()
	app.Shutdown()
	<-appDone
	assert.Equal(t, Closing, <-app.GetStateChannel())
	assert.Equal(t, Closed, <-app.GetStateChannel())
}

// isAppAvailable checks if the healthcheck server at the given endpoint is
// returning `available`.
func isAppAvailable(t *testing.T, healthCheckEndPoint string) bool {
	client := &http.Client{}
	resp, err := client.Get(healthCheckEndPoint)
	require.NoError(t, err)

	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func assertMetrics(t *testing.T, prefix string, metricsPort uint16, mandatoryLabels []string) {
	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("http://localhost:%d/metrics", metricsPort))
	require.NoError(t, err)

	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)

	var parser expfmt.TextParser
	parsed, err := parser.TextToMetricFamilies(reader)
	require.NoError(t, err)

	for metricName, metricFamily := range parsed {
		// require is used here so test fails with a single message.
		require.True(
			t,
			strings.HasPrefix(metricName, prefix),
			"expected prefix %q but string starts with %q",
			prefix,
			metricName[:len(prefix)+1]+"...")

		for _, metric := range metricFamily.Metric {
			var labelNames []string
			for _, labelPair := range metric.Label {
				labelNames = append(labelNames, *labelPair.Name)
			}

			for _, mandatoryLabel := range mandatoryLabels {
				// require is used here so test fails with a single message.
				require.Contains(t, labelNames, mandatoryLabel, "mandatory label %q not present", mandatoryLabel)
			}
		}
	}
}

type minimalParserLoader struct{}

func (*minimalParserLoader) Get() (*config.Parser, error) {
	configStr := `
receivers:
  otlp:
    protocols:
      grpc:
exporters:
  logging:
processors:
  batch:

extensions:

service:
  extensions:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging]
`
	return config.NewParserFromBuffer(strings.NewReader(configStr))
}

type errParserLoader struct {
	err error
}

func (epl *errParserLoader) Get() (*config.Parser, error) {
	return nil, epl.err
}

func TestApplication_reloadService(t *testing.T) {
	factories, err := defaultcomponents.Components()
	require.NoError(t, err)
	ctx := context.Background()
	sentinelError := errors.New("sentinel error")

	tests := []struct {
		name           string
		parserProvider parserprovider.ParserProvider
		service        *service
		skip           bool
	}{
		{
			name:           "first_load_err",
			parserProvider: &errParserLoader{err: sentinelError},
		},
		{
			name:           "retire_service_ok_load_err",
			parserProvider: &errParserLoader{err: sentinelError},
			service: &service{
				logger:          zap.NewNop(),
				builtExporters:  builder.Exporters{},
				builtPipelines:  builder.BuiltPipelines{},
				builtReceivers:  builder.Receivers{},
				builtExtensions: builder.Extensions{},
			},
		},
		{
			name:           "retire_service_ok_load_ok",
			parserProvider: new(minimalParserLoader),
			service: &service{
				logger:          zap.NewNop(),
				builtExporters:  builder.Exporters{},
				builtPipelines:  builder.BuiltPipelines{},
				builtReceivers:  builder.Receivers{},
				builtExtensions: builder.Extensions{},
			},
			skip: runtime.GOOS == "darwin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Log("Skipping test", tt.name)
				return
			}

			app := Application{
				logger:         zap.NewNop(),
				parserProvider: tt.parserProvider,
				factories:      factories,
				service:        tt.service,
			}

			err := app.reloadService(ctx)

			if err != nil {
				assert.ErrorIs(t, err, sentinelError)
				return
			}

			// If successful need to shutdown active service.
			assert.NoError(t, app.service.Shutdown(ctx))
		})
	}
}
