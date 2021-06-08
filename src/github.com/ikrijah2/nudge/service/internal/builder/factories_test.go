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

package builder

import (
	"context"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/exporter/exporterhelper"
	"github.com/ikrijah2/nudge/extension/extensionhelper"
	"github.com/ikrijah2/nudge/internal/testcomponents"
	"github.com/ikrijah2/nudge/processor/processorhelper"
	"github.com/ikrijah2/nudge/receiver/receiverhelper"
)

func createTestFactories() component.Factories {
	exampleReceiverFactory := testcomponents.ExampleReceiverFactory
	exampleProcessorFactory := testcomponents.ExampleProcessorFactory
	exampleExporterFactory := testcomponents.ExampleExporterFactory
	badExtensionFactory := newBadExtensionFactory()
	badReceiverFactory := newBadReceiverFactory()
	badProcessorFactory := newBadProcessorFactory()
	badExporterFactory := newBadExporterFactory()

	factories := component.Factories{
		Extensions: map[config.Type]component.ExtensionFactory{
			badExtensionFactory.Type(): badExtensionFactory,
		},
		Receivers: map[config.Type]component.ReceiverFactory{
			exampleReceiverFactory.Type(): exampleReceiverFactory,
			badReceiverFactory.Type():     badReceiverFactory,
		},
		Processors: map[config.Type]component.ProcessorFactory{
			exampleProcessorFactory.Type(): exampleProcessorFactory,
			badProcessorFactory.Type():     badProcessorFactory,
		},
		Exporters: map[config.Type]component.ExporterFactory{
			exampleExporterFactory.Type(): exampleExporterFactory,
			badExporterFactory.Type():     badExporterFactory,
		},
	}

	return factories
}

func newBadReceiverFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory("bf", func() config.Receiver {
		return &config.ReceiverSettings{TypeVal: "bf"}
	})
}

func newBadProcessorFactory() component.ProcessorFactory {
	return processorhelper.NewFactory("bf", func() config.Processor {
		return &config.ProcessorSettings{
			TypeVal: "bf",
		}
	})
}

func newBadExporterFactory() component.ExporterFactory {
	return exporterhelper.NewFactory("bf", func() config.Exporter {
		return &config.ExporterSettings{
			TypeVal: "bf",
		}
	})
}

func newBadExtensionFactory() component.ExtensionFactory {
	return extensionhelper.NewFactory(
		"bf",
		func() config.Extension {
			return &config.ExtensionSettings{
				TypeVal: "bf",
			}
		},
		func(ctx context.Context, params component.ExtensionCreateParams, extension config.Extension) (component.Extension, error) {
			return nil, nil
		},
	)
}
