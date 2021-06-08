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

// Package consumer contains interfaces that receive and process consumerdata.
package consumer

import (
	"context"

	"github.com/ikrijah2/nudge/consumer/pdata"
)

// Metrics is the new metrics consumer interface that receives pdata.Metrics, processes it
// as needed, and sends it to the next processing node if any or to the destination.
type Metrics interface {
	// ConsumeMetrics receives pdata.Metrics for consumption.
	ConsumeMetrics(ctx context.Context, md pdata.Metrics) error
}

// Traces is an interface that receives pdata.Traces, processes it
// as needed, and sends it to the next processing node if any or to the destination.
type Traces interface {
	// ConsumeTraces receives pdata.Traces for consumption.
	ConsumeTraces(ctx context.Context, td pdata.Traces) error
}

// Logs is an interface that receives pdata.Logs, processes it
// as needed, and sends it to the next processing node if any or to the destination.
type Logs interface {
	// ConsumeLogs receives pdata.Logs for consumption.
	ConsumeLogs(ctx context.Context, ld pdata.Logs) error
}
