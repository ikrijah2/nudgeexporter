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

package batchprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ikrijah2/nudge/consumer/pdata"
	"github.com/ikrijah2/nudge/internal/testdata"
)

func TestSplitLogs_noop(t *testing.T) {
	td := testdata.GenerateLogDataManyLogsSameResource(20)
	splitSize := 40
	split := splitLogs(splitSize, td)
	assert.Equal(t, td, split)

	td.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().Resize(5)
	assert.EqualValues(t, td, split)
}

func TestSplitLogs(t *testing.T) {
	td := testdata.GenerateLogDataManyLogsSameResource(20)
	logs := td.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs()
	for i := 0; i < logs.Len(); i++ {
		logs.At(i).SetName(getTestLogName(0, i))
	}
	cp := pdata.NewLogs()
	cp.ResourceLogs().Resize(1)
	cp.ResourceLogs().At(0).InstrumentationLibraryLogs().Resize(1)
	cp.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().Resize(5)
	cpLogs := cp.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs()
	td.ResourceLogs().At(0).Resource().CopyTo(
		cp.ResourceLogs().At(0).Resource())
	td.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).InstrumentationLibrary().CopyTo(
		cp.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).InstrumentationLibrary())
	logs.At(19).CopyTo(cpLogs.At(0))
	logs.At(18).CopyTo(cpLogs.At(1))
	logs.At(17).CopyTo(cpLogs.At(2))
	logs.At(16).CopyTo(cpLogs.At(3))
	logs.At(15).CopyTo(cpLogs.At(4))

	splitSize := 5
	split := splitLogs(splitSize, td)
	assert.Equal(t, splitSize, split.LogRecordCount())
	assert.Equal(t, cp, split)
	assert.Equal(t, 15, td.LogRecordCount())
	assert.Equal(t, "test-log-int-0-19", split.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().At(0).Name())
	assert.Equal(t, "test-log-int-0-15", split.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().At(4).Name())
}

func TestSplitLogsMultipleResourceLogs(t *testing.T) {
	td := testdata.GenerateLogDataManyLogsSameResource(20)
	logs := td.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs()
	for i := 0; i < logs.Len(); i++ {
		logs.At(i).SetName(getTestLogName(0, i))
	}
	td.ResourceLogs().Resize(2)
	// add second index to resource logs
	testdata.GenerateLogDataManyLogsSameResource(20).
		ResourceLogs().At(0).CopyTo(td.ResourceLogs().At(1))
	logs = td.ResourceLogs().At(1).InstrumentationLibraryLogs().At(0).Logs()
	for i := 0; i < logs.Len(); i++ {
		logs.At(i).SetName(getTestLogName(1, i))
	}

	splitSize := 5
	split := splitLogs(splitSize, td)
	assert.Equal(t, splitSize, split.LogRecordCount())
	assert.Equal(t, 35, td.LogRecordCount())
	assert.Equal(t, "test-log-int-1-19", split.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().At(0).Name())
	assert.Equal(t, "test-log-int-1-15", split.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().At(4).Name())
}

func TestSplitLogsMultipleResourceLogs_split_size_greater_than_log_size(t *testing.T) {
	td := testdata.GenerateLogDataManyLogsSameResource(20)
	logs := td.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs()
	for i := 0; i < logs.Len(); i++ {
		logs.At(i).SetName(getTestLogName(0, i))
	}
	td.ResourceLogs().Resize(2)
	// add second index to resource logs
	testdata.GenerateLogDataManyLogsSameResource(20).
		ResourceLogs().At(0).CopyTo(td.ResourceLogs().At(1))
	logs = td.ResourceLogs().At(1).InstrumentationLibraryLogs().At(0).Logs()
	for i := 0; i < logs.Len(); i++ {
		logs.At(i).SetName(getTestLogName(1, i))
	}

	splitSize := 25
	split := splitLogs(splitSize, td)
	assert.Equal(t, splitSize, split.LogRecordCount())
	assert.Equal(t, 40-splitSize, td.LogRecordCount())
	assert.Equal(t, 1, td.ResourceLogs().Len())
	assert.Equal(t, "test-log-int-1-19", split.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().At(0).Name())
	assert.Equal(t, "test-log-int-1-0", split.ResourceLogs().At(0).InstrumentationLibraryLogs().At(0).Logs().At(19).Name())
	assert.Equal(t, "test-log-int-0-19", split.ResourceLogs().At(1).InstrumentationLibraryLogs().At(0).Logs().At(0).Name())
	assert.Equal(t, "test-log-int-0-15", split.ResourceLogs().At(1).InstrumentationLibraryLogs().At(0).Logs().At(4).Name())
}
