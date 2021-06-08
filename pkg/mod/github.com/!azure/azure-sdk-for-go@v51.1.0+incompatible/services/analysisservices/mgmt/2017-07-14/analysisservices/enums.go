package analysisservices

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

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Paused ...
	Paused ProvisioningState = "Paused"
	// Pausing ...
	Pausing ProvisioningState = "Pausing"
	// Preparing ...
	Preparing ProvisioningState = "Preparing"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Resuming ...
	Resuming ProvisioningState = "Resuming"
	// Scaling ...
	Scaling ProvisioningState = "Scaling"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Suspended ...
	Suspended ProvisioningState = "Suspended"
	// Suspending ...
	Suspending ProvisioningState = "Suspending"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Deleting, Failed, Paused, Pausing, Preparing, Provisioning, Resuming, Scaling, Succeeded, Suspended, Suspending, Updating}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic ...
	Basic SkuTier = "Basic"
	// Development ...
	Development SkuTier = "Development"
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Basic, Development, Standard}
}

// State enumerates the values for state.
type State string

const (
	// StateDeleting ...
	StateDeleting State = "Deleting"
	// StateFailed ...
	StateFailed State = "Failed"
	// StatePaused ...
	StatePaused State = "Paused"
	// StatePausing ...
	StatePausing State = "Pausing"
	// StatePreparing ...
	StatePreparing State = "Preparing"
	// StateProvisioning ...
	StateProvisioning State = "Provisioning"
	// StateResuming ...
	StateResuming State = "Resuming"
	// StateScaling ...
	StateScaling State = "Scaling"
	// StateSucceeded ...
	StateSucceeded State = "Succeeded"
	// StateSuspended ...
	StateSuspended State = "Suspended"
	// StateSuspending ...
	StateSuspending State = "Suspending"
	// StateUpdating ...
	StateUpdating State = "Updating"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateDeleting, StateFailed, StatePaused, StatePausing, StatePreparing, StateProvisioning, StateResuming, StateScaling, StateSucceeded, StateSuspended, StateSuspending, StateUpdating}
}

// Status enumerates the values for status.
type Status string

const (
	// Live ...
	Live Status = "Live"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Live}
}
