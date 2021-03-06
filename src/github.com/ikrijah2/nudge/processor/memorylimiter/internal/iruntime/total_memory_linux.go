// Copyright  The OpenTelemetry Authors
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

// +build linux

package iruntime

import "github.com/ikrijah2/nudge/processor/memorylimiter/internal/cgroups"

// TotalMemory returns total available memory.
// This implementation is meant for linux and uses cgroups to determine available memory.
func TotalMemory() (int64, error) {
	cgroups, err := cgroups.NewCGroupsForCurrentProcess()
	if err != nil {
		return 0, err
	}
	memoryQuota, defined, err := cgroups.MemoryQuota()
	if err != nil || !defined {
		return 0, err
	}
	return memoryQuota, nil
}
