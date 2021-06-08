package batch

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

// AccountKeyType enumerates the values for account key type.
type AccountKeyType string

const (
	// Primary ...
	Primary AccountKeyType = "Primary"
	// Secondary ...
	Secondary AccountKeyType = "Secondary"
)

// PossibleAccountKeyTypeValues returns an array of possible values for the AccountKeyType const type.
func PossibleAccountKeyTypeValues() []AccountKeyType {
	return []AccountKeyType{Primary, Secondary}
}

// PackageState enumerates the values for package state.
type PackageState string

const (
	// Active ...
	Active PackageState = "active"
	// Pending ...
	Pending PackageState = "pending"
	// Unmapped ...
	Unmapped PackageState = "unmapped"
)

// PossiblePackageStateValues returns an array of possible values for the PackageState const type.
func PossiblePackageStateValues() []PackageState {
	return []PackageState{Active, Pending, Unmapped}
}

// PoolAllocationMode enumerates the values for pool allocation mode.
type PoolAllocationMode string

const (
	// BatchService ...
	BatchService PoolAllocationMode = "BatchService"
	// UserSubscription ...
	UserSubscription PoolAllocationMode = "UserSubscription"
)

// PossiblePoolAllocationModeValues returns an array of possible values for the PoolAllocationMode const type.
func PossiblePoolAllocationModeValues() []PoolAllocationMode {
	return []PoolAllocationMode{BatchService, UserSubscription}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Cancelled ...
	Cancelled ProvisioningState = "Cancelled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Invalid ...
	Invalid ProvisioningState = "Invalid"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Cancelled, Creating, Deleting, Failed, Invalid, Succeeded}
}
