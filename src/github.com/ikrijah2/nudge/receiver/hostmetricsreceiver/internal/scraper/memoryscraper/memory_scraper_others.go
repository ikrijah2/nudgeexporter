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

// +build !linux,!windows

package memoryscraper

import (
	"github.com/shirou/gopsutil/mem"

	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/metadata"
)

const memStatesLen = 3

func appendMemoryUsageStateDataPoints(idps pdata.IntDataPointSlice, now pdata.Timestamp, memInfo *mem.VirtualMemoryStat) {
	initializeMemoryUsageDataPoint(idps.At(0), now, metadata.LabelMemState.Used, int64(memInfo.Used))
	initializeMemoryUsageDataPoint(idps.At(1), now, metadata.LabelMemState.Free, int64(memInfo.Free))
	initializeMemoryUsageDataPoint(idps.At(2), now, metadata.LabelMemState.Inactive, int64(memInfo.Inactive))
}
