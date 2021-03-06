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
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/exporter/kafkaexporter"
)

// Config defines configuration for Kafka receiver.
type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`
	// The list of kafka brokers (default localhost:9092)
	Brokers []string `mapstructure:"brokers"`
	// Kafka protocol version
	ProtocolVersion string `mapstructure:"protocol_version"`
	// The name of the kafka topic to consume from (default "otlp_spans")
	Topic string `mapstructure:"topic"`
	// Encoding of the messages (default "otlp_proto")
	Encoding string `mapstructure:"encoding"`
	// The consumer group that receiver will be consuming messages from (default "otel-collector")
	GroupID string `mapstructure:"group_id"`
	// The consumer client ID that receiver will use (default "otel-collector")
	ClientID string `mapstructure:"client_id"`

	// Metadata is the namespace for metadata management properties used by the
	// Client, and shared by the Producer/Consumer.
	Metadata kafkaexporter.Metadata `mapstructure:"metadata"`

	Authentication kafkaexporter.Authentication `mapstructure:"auth"`
}

var _ config.Receiver = (*Config)(nil)

// Validate checks the receiver configuration is valid
func (cfg *Config) Validate() error {
	return nil
}
