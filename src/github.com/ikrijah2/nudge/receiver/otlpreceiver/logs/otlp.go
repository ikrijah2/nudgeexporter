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

package logs

import (
	"context"

	"github.com/ikrijah2/nudge/client"
	"github.com/ikrijah2/nudge/consumer"
	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/internal"
	collectorlog "github.com/ikrijah2/nudge/internal/data/protogen/collector/logs/v1"
	"github.com/ikrijah2/nudge/obsreport"
)

const (
	dataFormatProtobuf = "protobuf"
)

// Receiver is the type used to handle spans from OpenTelemetry exporters.
type Receiver struct {
	instanceName string
	nextConsumer consumer.Logs
}

// New creates a new Receiver reference.
func New(instanceName string, nextConsumer consumer.Logs) *Receiver {
	r := &Receiver{
		instanceName: instanceName,
		nextConsumer: nextConsumer,
	}

	return r
}

const (
	receiverTagValue  = "otlp_log"
	receiverTransport = "grpc"
)

// Export implements the service Export logs func.
func (r *Receiver) Export(ctx context.Context, req *collectorlog.ExportLogsServiceRequest) (*collectorlog.ExportLogsServiceResponse, error) {
	// We need to ensure that it propagates the receiver name as a tag
	ctxWithReceiverName := obsreport.ReceiverContext(ctx, r.instanceName, receiverTransport)

	ld := pdata.LogsFromInternalRep(internal.LogsFromOtlp(req))
	err := r.sendToNextConsumer(ctxWithReceiverName, ld)
	if err != nil {
		return nil, err
	}

	return &collectorlog.ExportLogsServiceResponse{}, nil
}

func (r *Receiver) sendToNextConsumer(ctx context.Context, ld pdata.Logs) error {
	numSpans := ld.LogRecordCount()
	if numSpans == 0 {
		return nil
	}

	if c, ok := client.FromGRPC(ctx); ok {
		ctx = client.NewContext(ctx, c)
	}

	ctx = obsreport.StartLogsReceiveOp(ctx, r.instanceName, receiverTransport)
	err := r.nextConsumer.ConsumeLogs(ctx, ld)
	obsreport.EndLogsReceiveOp(ctx, dataFormatProtobuf, numSpans, err)

	return err
}
