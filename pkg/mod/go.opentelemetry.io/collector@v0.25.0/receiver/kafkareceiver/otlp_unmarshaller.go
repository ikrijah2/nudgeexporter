// Copyright 2020 The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafkareceiver

import (
	"go.opentelemetry.io/collector/consumer/pdata"
)

type otlpTracesPbUnmarshaller struct {
}

var _ TracesUnmarshaller = (*otlpTracesPbUnmarshaller)(nil)

func (p *otlpTracesPbUnmarshaller) Unmarshal(bytes []byte) (pdata.Traces, error) {
	return pdata.TracesFromOtlpProtoBytes(bytes)
}

func (*otlpTracesPbUnmarshaller) Encoding() string {
	return defaultEncoding
}

type otlpLogsPbUnmarshaller struct {
}

var _ LogsUnmarshaller = (*otlpLogsPbUnmarshaller)(nil)

func (p *otlpLogsPbUnmarshaller) Unmarshal(bytes []byte) (pdata.Logs, error) {
	return pdata.LogsFromOtlpProtoBytes(bytes)
}

func (*otlpLogsPbUnmarshaller) Encoding() string {
	return defaultEncoding
}
