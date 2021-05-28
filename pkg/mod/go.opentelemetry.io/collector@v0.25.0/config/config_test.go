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

package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errInvalidRecvConfig = errors.New("invalid receiver config")
var errInvalidExpConfig = errors.New("invalid exporter config")
var errInvalidProcConfig = errors.New("invalid processor config")
var errInvalidExtConfig = errors.New("invalid extension config")

type nopRecvConfig struct {
	ReceiverSettings
}

func (nc *nopRecvConfig) Validate() error {
	if nc.TypeVal != "nop" {
		return errInvalidRecvConfig
	}
	return nil
}

type nopExpConfig struct {
	ExporterSettings
}

func (nc *nopExpConfig) Validate() error {
	if nc.TypeVal != "nop" {
		return errInvalidExpConfig
	}
	return nil
}

type nopProcConfig struct {
	ProcessorSettings
}

func (nc *nopProcConfig) Validate() error {
	if nc.TypeVal != "nop" {
		return errInvalidProcConfig
	}
	return nil
}

type nopExtConfig struct {
	ExtensionSettings
}

func (nc *nopExtConfig) Validate() error {
	if nc.TypeVal != "nop" {
		return errInvalidExtConfig
	}
	return nil
}

func TestConfigValidate(t *testing.T) {
	var testCases = []struct {
		name     string // test case name (also file name containing config yaml)
		cfgFn    func() *Config
		expected error
	}{
		{
			name:     "valid",
			cfgFn:    generateConfig,
			expected: nil,
		},
		{
			name: "missing-exporters",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Exporters = nil
				return cfg
			},
			expected: errMissingExporters,
		},
		{
			name: "missing-receivers",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Receivers = nil
				return cfg
			},
			expected: errMissingReceivers,
		},
		{
			name: "invalid-extension-reference",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Service.Extensions = append(cfg.Service.Extensions, "nop/2")
				return cfg
			},
			expected: errors.New(`service references extension "nop/2" which does not exist`),
		},
		{
			name: "invalid-receiver-reference",
			cfgFn: func() *Config {
				cfg := generateConfig()
				pipe := cfg.Service.Pipelines["traces"]
				pipe.Receivers = append(pipe.Receivers, "nop/2")
				return cfg
			},
			expected: errors.New(`pipeline "traces" references receiver "nop/2" which does not exist`),
		},
		{
			name: "invalid-processor-reference",
			cfgFn: func() *Config {
				cfg := generateConfig()
				pipe := cfg.Service.Pipelines["traces"]
				pipe.Processors = append(pipe.Processors, "nop/2")
				return cfg
			},
			expected: errors.New(`pipeline "traces" references processor "nop/2" which does not exist`),
		},
		{
			name: "invalid-exporter-reference",
			cfgFn: func() *Config {
				cfg := generateConfig()
				pipe := cfg.Service.Pipelines["traces"]
				pipe.Exporters = append(pipe.Exporters, "nop/2")
				return cfg
			},
			expected: errors.New(`pipeline "traces" references exporter "nop/2" which does not exist`),
		},
		{
			name: "missing-pipeline-receivers",
			cfgFn: func() *Config {
				cfg := generateConfig()
				pipe := cfg.Service.Pipelines["traces"]
				pipe.Receivers = nil
				return cfg
			},
			expected: errors.New(`pipeline "traces" must have at least one receiver`),
		},
		{
			name: "missing-pipeline-exporters",
			cfgFn: func() *Config {
				cfg := generateConfig()
				pipe := cfg.Service.Pipelines["traces"]
				pipe.Exporters = nil
				return cfg
			},
			expected: errors.New(`pipeline "traces" must have at least one exporter`),
		},
		{
			name: "missing-pipelines",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Service.Pipelines = nil
				return cfg
			},
			expected: errMissingServicePipelines,
		},
		{
			name: "invalid-receiver-config",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Receivers["nop"] = &nopRecvConfig{
					ReceiverSettings: ReceiverSettings{
						TypeVal: "invalid_rec_type",
					},
				}
				return cfg
			},
			expected: fmt.Errorf(`receiver "nop" has invalid configuration: %w`, errInvalidRecvConfig),
		},
		{
			name: "invalid-exporter-config",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Exporters["nop"] = &nopExpConfig{
					ExporterSettings: ExporterSettings{
						TypeVal: "invalid_exp_type",
					},
				}
				return cfg
			},
			expected: fmt.Errorf(`exporter "nop" has invalid configuration: %w`, errInvalidExpConfig),
		},
		{
			name: "invalid-processor-config",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Processors["nop"] = &nopProcConfig{
					ProcessorSettings: ProcessorSettings{
						TypeVal: "invalid_proc_type",
					},
				}
				return cfg
			},
			expected: fmt.Errorf(`processor "nop" has invalid configuration: %w`, errInvalidProcConfig),
		},
		{
			name: "invalid-extension-config",
			cfgFn: func() *Config {
				cfg := generateConfig()
				cfg.Extensions["nop"] = &nopExtConfig{
					ExtensionSettings: ExtensionSettings{
						TypeVal: "invalid_ext_type",
					},
				}
				return cfg
			},
			expected: fmt.Errorf(`extension "nop" has invalid configuration: %w`, errInvalidExtConfig),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			cfg := test.cfgFn()
			assert.Equal(t, test.expected, cfg.Validate())
		})
	}
}

func generateConfig() *Config {
	return &Config{
		Receivers: map[string]Receiver{
			"nop": &nopRecvConfig{
				ReceiverSettings: ReceiverSettings{
					TypeVal: "nop",
				},
			},
		},
		Exporters: map[string]Exporter{
			"nop": &nopExpConfig{
				ExporterSettings: ExporterSettings{
					TypeVal: "nop",
				},
			},
		},
		Processors: map[string]Processor{
			"nop": &nopProcConfig{
				ProcessorSettings: ProcessorSettings{
					TypeVal: "nop",
				},
			},
		},
		Extensions: map[string]Extension{
			"nop": &nopExtConfig{
				ExtensionSettings: ExtensionSettings{
					TypeVal: "nop",
				},
			},
		},
		Service: Service{
			Extensions: []string{"nop"},
			Pipelines: map[string]*Pipeline{
				"traces": {
					Name:       "traces",
					InputType:  TracesDataType,
					Receivers:  []string{"nop"},
					Processors: []string{"nop"},
					Exporters:  []string{"nop"},
				},
			},
		},
	}
}
