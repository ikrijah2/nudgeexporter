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

package healthcheckextension

import (
	"context"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/config/confignet"
	"github.com/ikrijah2/nudge/extension/extensionhelper"
)

const (
	// The value of extension "type" in configuration.
	typeStr = "health_check"

	// Use 0.0.0.0 to make the health check endpoint accessible
	// in container orchestration environments like Kubernetes.
	defaultEndpoint = "0.0.0.0:13133"
)

// NewFactory creates a factory for HealthCheck extension.
func NewFactory() component.ExtensionFactory {
	return extensionhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		createExtension)
}

func createDefaultConfig() config.Extension {
	return &Config{
		ExtensionSettings: config.NewExtensionSettings(typeStr),
		TCPAddr: confignet.TCPAddr{
			Endpoint: defaultEndpoint,
		},
	}
}

func createExtension(_ context.Context, params component.ExtensionCreateParams, cfg config.Extension) (component.Extension, error) {
	config := cfg.(*Config)

	return newServer(*config, params.Logger), nil
}
