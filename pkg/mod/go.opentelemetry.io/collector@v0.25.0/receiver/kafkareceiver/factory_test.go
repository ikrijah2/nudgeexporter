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

package kafkareceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configcheck"
	"go.opentelemetry.io/collector/consumer/pdata"
)

func TestCreateDefaultConfig(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, configcheck.ValidateConfig(cfg))
	assert.Equal(t, []string{defaultBroker}, cfg.Brokers)
	assert.Equal(t, defaultTopic, cfg.Topic)
	assert.Equal(t, defaultGroupID, cfg.GroupID)
	assert.Equal(t, defaultClientID, cfg.ClientID)
}

func TestCreateTracesReceiver(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	cfg.Brokers = []string{"invalid:9092"}
	cfg.ProtocolVersion = "2.0.0"
	f := kafkaReceiverFactory{tracesUnmarshalers: defaultTracesUnmarshallers()}
	r, err := f.createTracesReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
	// no available broker
	require.Error(t, err)
	assert.Nil(t, r)
}

func TestCreateTracesReceiver_error(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	cfg.ProtocolVersion = "2.0.0"
	// disable contacting broker at startup
	cfg.Metadata.Full = false
	f := kafkaReceiverFactory{tracesUnmarshalers: defaultTracesUnmarshallers()}
	r, err := f.createTracesReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
	require.NoError(t, err)
	assert.NotNil(t, r)
}

func TestWithTracesUnmarshallers(t *testing.T) {
	unmarshaller := &customTracesUnmarshaller{}
	f := NewFactory(WithAddTracesUnmarshallers(map[string]TracesUnmarshaller{unmarshaller.Encoding(): unmarshaller}))
	cfg := createDefaultConfig().(*Config)
	// disable contacting broker
	cfg.Metadata.Full = false
	cfg.ProtocolVersion = "2.0.0"

	t.Run("custom_encoding", func(t *testing.T) {
		cfg.Encoding = unmarshaller.Encoding()
		receiver, err := f.CreateTracesReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
		require.NoError(t, err)
		require.NotNil(t, receiver)
	})
	t.Run("default_encoding", func(t *testing.T) {
		cfg.Encoding = new(otlpTracesPbUnmarshaller).Encoding()
		receiver, err := f.CreateTracesReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
		require.NoError(t, err)
		assert.NotNil(t, receiver)
	})
}

func TestCreateLogsReceiver(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	cfg.Brokers = []string{"invalid:9092"}
	cfg.ProtocolVersion = "2.0.0"
	f := kafkaReceiverFactory{logsUnmarshaller: defaultLogsUnmarshallers()}
	r, err := f.createLogsReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
	// no available broker
	require.Error(t, err)
	assert.Nil(t, r)
}

func TestCreateLogsReceiver_error(t *testing.T) {
	cfg := createDefaultConfig().(*Config)
	cfg.ProtocolVersion = "2.0.0"
	// disable contacting broker at startup
	cfg.Metadata.Full = false
	f := kafkaReceiverFactory{logsUnmarshaller: defaultLogsUnmarshallers()}
	r, err := f.createLogsReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
	require.NoError(t, err)
	assert.NotNil(t, r)
}

func TestWithLogsUnmarshallers(t *testing.T) {
	unmarshaller := &customLogsUnmarshaller{}
	f := NewFactory(WithAddLogsUnmarshallers(map[string]LogsUnmarshaller{unmarshaller.Encoding(): unmarshaller}))
	cfg := createDefaultConfig().(*Config)
	// disable contacting broker
	cfg.Metadata.Full = false
	cfg.ProtocolVersion = "2.0.0"

	t.Run("custom_encoding", func(t *testing.T) {
		cfg.Encoding = unmarshaller.Encoding()
		exporter, err := f.CreateLogsReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
		require.NoError(t, err)
		require.NotNil(t, exporter)
	})
	t.Run("default_encoding", func(t *testing.T) {
		cfg.Encoding = new(otlpLogsPbUnmarshaller).Encoding()
		exporter, err := f.CreateLogsReceiver(context.Background(), component.ReceiverCreateParams{}, cfg, nil)
		require.NoError(t, err)
		assert.NotNil(t, exporter)
	})
}

type customTracesUnmarshaller struct {
}

type customLogsUnmarshaller struct {
}

var _ TracesUnmarshaller = (*customTracesUnmarshaller)(nil)

func (c customTracesUnmarshaller) Unmarshal([]byte) (pdata.Traces, error) {
	panic("implement me")
}

func (c customTracesUnmarshaller) Encoding() string {
	return "custom"
}

func (c customLogsUnmarshaller) Unmarshal([]byte) (pdata.Logs, error) {
	panic("implement me")
}

func (c customLogsUnmarshaller) Encoding() string {
	return "custom"
}
