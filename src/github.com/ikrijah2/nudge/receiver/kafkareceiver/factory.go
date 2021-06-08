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
	"time"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/consumer"
	"github.com/ikrijah2/nudge/exporter/kafkaexporter"
	"github.com/ikrijah2/nudge/receiver/receiverhelper"
)

const (
	typeStr         = "kafka"
	defaultTopic    = "otlp_spans"
	defaultEncoding = "otlp_proto"
	defaultBroker   = "localhost:9092"
	defaultClientID = "otel-collector"
	defaultGroupID  = defaultClientID

	// default from sarama.NewConfig()
	defaultMetadataRetryMax = 3
	// default from sarama.NewConfig()
	defaultMetadataRetryBackoff = time.Millisecond * 250
	// default from sarama.NewConfig()
	defaultMetadataFull = true
)

// FactoryOption applies changes to kafkaReceiverFactory.
type FactoryOption func(factory *kafkaReceiverFactory)

// WithAddTracesUnmarshallers adds marshallers.
func WithAddTracesUnmarshallers(encodingMarshaller map[string]TracesUnmarshaller) FactoryOption {
	return func(factory *kafkaReceiverFactory) {
		for encoding, unmarshaller := range encodingMarshaller {
			factory.tracesUnmarshalers[encoding] = unmarshaller
		}
	}
}

// WithAddLogsUnmarshallers adds LogsUnmarshallers.
func WithAddLogsUnmarshallers(encodingMarshaller map[string]LogsUnmarshaller) FactoryOption {
	return func(factory *kafkaReceiverFactory) {
		for encoding, unmarshaller := range encodingMarshaller {
			factory.logsUnmarshaller[encoding] = unmarshaller
		}
	}
}

// NewFactory creates Kafka receiver factory.
func NewFactory(options ...FactoryOption) component.ReceiverFactory {
	f := &kafkaReceiverFactory{
		tracesUnmarshalers: defaultTracesUnmarshallers(),
		logsUnmarshaller:   defaultLogsUnmarshallers(),
	}
	for _, o := range options {
		o(f)
	}
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		receiverhelper.WithTraces(f.createTracesReceiver),
		receiverhelper.WithLogs(f.createLogsReceiver),
	)
}

func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings: config.ReceiverSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		Topic:    defaultTopic,
		Encoding: defaultEncoding,
		Brokers:  []string{defaultBroker},
		ClientID: defaultClientID,
		GroupID:  defaultGroupID,
		Metadata: kafkaexporter.Metadata{
			Full: defaultMetadataFull,
			Retry: kafkaexporter.MetadataRetry{
				Max:     defaultMetadataRetryMax,
				Backoff: defaultMetadataRetryBackoff,
			},
		},
	}
}

type kafkaReceiverFactory struct {
	tracesUnmarshalers map[string]TracesUnmarshaller
	logsUnmarshaller   map[string]LogsUnmarshaller
}

func (f *kafkaReceiverFactory) createTracesReceiver(
	_ context.Context,
	params component.ReceiverCreateParams,
	cfg config.Receiver,
	nextConsumer consumer.Traces,
) (component.TracesReceiver, error) {
	c := cfg.(*Config)
	r, err := newTracesReceiver(*c, params, f.tracesUnmarshalers, nextConsumer)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (f *kafkaReceiverFactory) createLogsReceiver(
	_ context.Context,
	params component.ReceiverCreateParams,
	cfg config.Receiver,
	nextConsumer consumer.Logs,
) (component.LogsReceiver, error) {
	c := cfg.(*Config)
	r, err := newLogsReceiver(*c, params, f.logsUnmarshaller, nextConsumer)
	if err != nil {
		return nil, err
	}
	return r, nil
}
