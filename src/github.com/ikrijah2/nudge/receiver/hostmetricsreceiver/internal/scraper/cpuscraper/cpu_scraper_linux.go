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

// +build linux

package cpuscraper

import (
	"github.com/shirou/gopsutil/cpu"

	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/metadata"
)

const cpuStatesLen = 8

func appendCPUTimeStateDataPoints(ddps pdata.DoubleDataPointSlice, startIdx int, startTime, now pdata.Timestamp, cpuTime cpu.TimesStat) {
	initializeCPUTimeDataPoint(ddps.At(startIdx+0), startTime, now, cpuTime.CPU, metadata.LabelCPUState.User, cpuTime.User)
	initializeCPUTimeDataPoint(ddps.At(startIdx+1), startTime, now, cpuTime.CPU, metadata.LabelCPUState.System, cpuTime.System)
	initializeCPUTimeDataPoint(ddps.At(startIdx+2), startTime, now, cpuTime.CPU, metadata.LabelCPUState.Idle, cpuTime.Idle)
	initializeCPUTimeDataPoint(ddps.At(startIdx+3), startTime, now, cpuTime.CPU, metadata.LabelCPUState.Interrupt, cpuTime.Irq)
	initializeCPUTimeDataPoint(ddps.At(startIdx+4), startTime, now, cpuTime.CPU, metadata.LabelCPUState.Nice, cpuTime.Nice)
	initializeCPUTimeDataPoint(ddps.At(startIdx+5), startTime, now, cpuTime.CPU, metadata.LabelCPUState.Softirq, cpuTime.Softirq)
	initializeCPUTimeDataPoint(ddps.At(startIdx+6), startTime, now, cpuTime.CPU, metadata.LabelCPUState.Steal, cpuTime.Steal)
	initializeCPUTimeDataPoint(ddps.At(startIdx+7), startTime, now, cpuTime.CPU, metadata.LabelCPUState.Wait, cpuTime.Iowait)
}
