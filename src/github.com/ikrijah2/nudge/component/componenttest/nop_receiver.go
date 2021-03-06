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

package componenttest

import (
	"context"

	"github.com/ikrijah2/nudge/component"
	"github.com/ikrijah2/nudge/component/componenthelper"
	"github.com/ikrijah2/nudge/config"
	"github.com/ikrijah2/nudge/consumer"
)

// nopReceiverFactory is factory for nopReceiver.
type nopReceiverFactory struct{}

var nopReceiverFactoryInstance = &nopReceiverFactory{}

// NewNopReceiverFactory returns a component.ReceiverFactory that constructs nop exporters.
func NewNopReceiverFactory() component.ReceiverFactory {
	return nopReceiverFactoryInstance
}

// Type gets the type of the Receiver config created by this factory.
func (f *nopReceiverFactory) Type() config.Type {
	return "nop"
}

// CreateDefaultConfig creates the default configuration for the Receiver.
func (f *nopReceiverFactory) CreateDefaultConfig() config.Receiver {
	return &config.ReceiverSettings{
		TypeVal: f.Type(),
	}
}

// CreateTracesReceiver implements component.ReceiverFactory interface.
func (f *nopReceiverFactory) CreateTracesReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	_ config.Receiver,
	_ consumer.Traces,
) (component.TracesReceiver, error) {
	return nopReceiverInstance, nil
}

// CreateMetricsReceiver implements component.ReceiverFactory interface.
func (f *nopReceiverFactory) CreateMetricsReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	_ config.Receiver,
	_ consumer.Metrics,
) (component.MetricsReceiver, error) {
	return nopReceiverInstance, nil
}

// CreateLogsReceiver implements component.ReceiverFactory interface.
func (f *nopReceiverFactory) CreateLogsReceiver(
	_ context.Context,
	_ component.ReceiverCreateParams,
	_ config.Receiver,
	_ consumer.Logs,
) (component.LogsReceiver, error) {
	return nopReceiverInstance, nil
}

var nopReceiverInstance = &nopReceiver{
	Component: componenthelper.New(),
}

// nopReceiver stores consumed traces and metrics for testing purposes.
type nopReceiver struct {
	component.Component
}
