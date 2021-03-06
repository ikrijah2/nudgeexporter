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

package prometheusremotewriteexporter

import (
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/confighttp"
	"github.com/ikrijah2/nudge/exporter/exporterhelper"
)

// Config defines configuration for Remote Write exporter.
type Config struct {
	*config.ExporterSettings       `mapstructure:"-"`
	exporterhelper.TimeoutSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	exporterhelper.QueueSettings   `mapstructure:"sending_queue"`
	exporterhelper.RetrySettings   `mapstructure:"retry_on_failure"`

	// prefix attached to each exported metric name
	// See: https://prometheus.io/docs/practices/naming/#metric-names
	Namespace string `mapstructure:"namespace"`

	// ExternalLabels defines a map of label keys and values that are allowed to start with reserved prefix "__"
	ExternalLabels map[string]string `mapstructure:"external_labels"`

	HTTPClientSettings confighttp.HTTPClientSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
}

var _ config.Exporter = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *Config) Validate() error {
	return nil
}
