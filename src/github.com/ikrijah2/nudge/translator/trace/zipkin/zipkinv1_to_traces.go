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

package zipkin

import (
	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/translator/internaldata"
)

// V1JSONBatchToInternalTraces transforms a JSON blob with a list of Zipkin v1 spans into pdata.Traces.
func V1JSONBatchToInternalTraces(blob []byte, parseStringTags bool) (pdata.Traces, error) {
	traces := pdata.NewTraces()

	ocTraces, err := v1JSONBatchToOCProto(blob, parseStringTags)
	if err != nil {
		return traces, err
	}

	for _, td := range ocTraces {
		tmp := internaldata.OCToTraces(td.Node, td.Resource, td.Spans)
		tmp.ResourceSpans().MoveAndAppendTo(traces.ResourceSpans())
	}
	return traces, nil
}
