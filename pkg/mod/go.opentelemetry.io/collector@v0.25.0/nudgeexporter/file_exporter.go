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

package fileexporter

import (
	"context"
	"io"
	"sync"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/pdata"
	"github.com/ikrijah2/nudge/internal"
)

// Marshaler configuration used for marhsaling Protobuf to JSON. Use default config.
var marshaler = &jsonpb.Marshaler{}

// fileExporter is the implementation of file exporter that writes telemetry data to a file
// in Protobuf-JSON format.
type fileExporter struct {
	file  io.WriteCloser
	mutex sync.Mutex
}

func (e *fileExporter) ConsumeTraces(_ context.Context, td pdata.Traces) error {
	return exportMessageAsLine(e, internal.TracesToOtlp(td.InternalRep()))
}

func (e *fileExporter) ConsumeMetrics(_ context.Context, md pdata.Metrics) error {
	return exportMessageAsLine(e, internal.MetricsToOtlp(md.InternalRep()))
}

func (e *fileExporter) ConsumeLogs(_ context.Context, ld pdata.Logs) error {
	request := internal.LogsToOtlp(ld.InternalRep())
	return exportMessageAsLine(e, request)
}

func exportMessageAsLine(e *fileExporter, message proto.Message) error {
	// Ensure only one write operation happens at a time.
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if err := marshaler.Marshal(e.file, message); err != nil {
		return err
	}
	if _, err := io.WriteString(e.file, "\n"); err != nil {
		return err
	}
	return nil
}

func (e *fileExporter) Start(context.Context, component.Host) error {
	return nil
}

// Shutdown stops the exporter and is invoked during shutdown.
func (e *fileExporter) Shutdown(context.Context) error {
	return e.file.Close()
}
