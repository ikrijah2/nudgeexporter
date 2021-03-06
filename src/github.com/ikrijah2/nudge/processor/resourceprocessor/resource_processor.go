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

package resourceprocessor

import (
	"context"

	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/processor/processorhelper"
)

type resourceProcessor struct {
	attrProc *processorhelper.AttrProc
}

// ProcessTraces implements the TProcessor interface
func (rp *resourceProcessor) ProcessTraces(_ context.Context, td pdata.Traces) (pdata.Traces, error) {
	rss := td.ResourceSpans()
	for i := 0; i < rss.Len(); i++ {
		rp.attrProc.Process(rss.At(i).Resource().Attributes())
	}
	return td, nil
}

// ProcessMetrics implements the MProcessor interface
func (rp *resourceProcessor) ProcessMetrics(_ context.Context, md pdata.Metrics) (pdata.Metrics, error) {
	rms := md.ResourceMetrics()
	for i := 0; i < rms.Len(); i++ {
		rp.attrProc.Process(rms.At(i).Resource().Attributes())
	}
	return md, nil
}

// ProcessLogs implements the LProcessor interface
func (rp *resourceProcessor) ProcessLogs(_ context.Context, ld pdata.Logs) (pdata.Logs, error) {
	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rp.attrProc.Process(rls.At(i).Resource().Attributes())
	}
	return ld, nil
}
