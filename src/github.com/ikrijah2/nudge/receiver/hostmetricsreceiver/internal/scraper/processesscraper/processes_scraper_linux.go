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

// +build linux openbsd

package processesscraper

import (
	"github.com/shirou/gopsutil/load"

	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/receiver/hostmetricsreceiver/internal/metadata"
)

const unixSystemSpecificMetricsLen = 1

func appendUnixSystemSpecificProcessesMetrics(metrics pdata.MetricSlice, startIndex int, now pdata.Timestamp, misc *load.MiscStat) error {
	initializeProcessesCreatedMetric(metrics.At(startIndex), now, misc)
	return nil
}

func initializeProcessesCreatedMetric(metric pdata.Metric, now pdata.Timestamp, misc *load.MiscStat) {
	metadata.Metrics.SystemProcessesCreated.Init(metric)

	ddps := metric.IntSum().DataPoints()
	ddps.Resize(1)
	ddps.At(0).SetTimestamp(now)
	ddps.At(0).SetValue(int64(misc.ProcsCreated))
}
