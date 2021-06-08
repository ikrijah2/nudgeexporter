package hybridcompute

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InstanceViewTypes enumerates the values for instance view types.
type InstanceViewTypes string

const (
	// InstanceView ...
	InstanceView InstanceViewTypes = "instanceView"
)

// PossibleInstanceViewTypesValues returns an array of possible values for the InstanceViewTypes const type.
func PossibleInstanceViewTypesValues() []InstanceViewTypes {
	return []InstanceViewTypes{InstanceView}
}

// StatusLevelTypes enumerates the values for status level types.
type StatusLevelTypes string

const (
	// Error ...
	Error StatusLevelTypes = "Error"
	// Info ...
	Info StatusLevelTypes = "Info"
	// Warning ...
	Warning StatusLevelTypes = "Warning"
)

// PossibleStatusLevelTypesValues returns an array of possible values for the StatusLevelTypes const type.
func PossibleStatusLevelTypesValues() []StatusLevelTypes {
	return []StatusLevelTypes{Error, Info, Warning}
}

// StatusTypes enumerates the values for status types.
type StatusTypes string

const (
	// StatusTypesConnected ...
	StatusTypesConnected StatusTypes = "Connected"
	// StatusTypesDisconnected ...
	StatusTypesDisconnected StatusTypes = "Disconnected"
	// StatusTypesError ...
	StatusTypesError StatusTypes = "Error"
)

// PossibleStatusTypesValues returns an array of possible values for the StatusTypes const type.
func PossibleStatusTypesValues() []StatusTypes {
	return []StatusTypes{StatusTypesConnected, StatusTypesDisconnected, StatusTypesError}
}
