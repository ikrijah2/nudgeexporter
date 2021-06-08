package storagecache

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

// CacheIdentityType enumerates the values for cache identity type.
type CacheIdentityType string

const (
	// None ...
	None CacheIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned CacheIdentityType = "SystemAssigned"
)

// PossibleCacheIdentityTypeValues returns an array of possible values for the CacheIdentityType const type.
func PossibleCacheIdentityTypeValues() []CacheIdentityType {
	return []CacheIdentityType{None, SystemAssigned}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// DomainJoinedType enumerates the values for domain joined type.
type DomainJoinedType string

const (
	// Error ...
	Error DomainJoinedType = "Error"
	// No ...
	No DomainJoinedType = "No"
	// Yes ...
	Yes DomainJoinedType = "Yes"
)

// PossibleDomainJoinedTypeValues returns an array of possible values for the DomainJoinedType const type.
func PossibleDomainJoinedTypeValues() []DomainJoinedType {
	return []DomainJoinedType{Error, No, Yes}
}

// FirmwareStatusType enumerates the values for firmware status type.
type FirmwareStatusType string

const (
	// Available ...
	Available FirmwareStatusType = "available"
	// Unavailable ...
	Unavailable FirmwareStatusType = "unavailable"
)

// PossibleFirmwareStatusTypeValues returns an array of possible values for the FirmwareStatusType const type.
func PossibleFirmwareStatusTypeValues() []FirmwareStatusType {
	return []FirmwareStatusType{Available, Unavailable}
}

// HealthStateType enumerates the values for health state type.
type HealthStateType string

const (
	// Degraded ...
	Degraded HealthStateType = "Degraded"
	// Down ...
	Down HealthStateType = "Down"
	// Flushing ...
	Flushing HealthStateType = "Flushing"
	// Healthy ...
	Healthy HealthStateType = "Healthy"
	// Stopped ...
	Stopped HealthStateType = "Stopped"
	// Stopping ...
	Stopping HealthStateType = "Stopping"
	// Transitioning ...
	Transitioning HealthStateType = "Transitioning"
	// Unknown ...
	Unknown HealthStateType = "Unknown"
	// Upgrading ...
	Upgrading HealthStateType = "Upgrading"
)

// PossibleHealthStateTypeValues returns an array of possible values for the HealthStateType const type.
func PossibleHealthStateTypeValues() []HealthStateType {
	return []HealthStateType{Degraded, Down, Flushing, Healthy, Stopped, Stopping, Transitioning, Unknown, Upgrading}
}

// MetricAggregationType enumerates the values for metric aggregation type.
type MetricAggregationType string

const (
	// MetricAggregationTypeAverage ...
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	// MetricAggregationTypeCount ...
	MetricAggregationTypeCount MetricAggregationType = "Count"
	// MetricAggregationTypeMaximum ...
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
	// MetricAggregationTypeMinimum ...
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	// MetricAggregationTypeNone ...
	MetricAggregationTypeNone MetricAggregationType = "None"
	// MetricAggregationTypeNotSpecified ...
	MetricAggregationTypeNotSpecified MetricAggregationType = "NotSpecified"
	// MetricAggregationTypeTotal ...
	MetricAggregationTypeTotal MetricAggregationType = "Total"
)

// PossibleMetricAggregationTypeValues returns an array of possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{MetricAggregationTypeAverage, MetricAggregationTypeCount, MetricAggregationTypeMaximum, MetricAggregationTypeMinimum, MetricAggregationTypeNone, MetricAggregationTypeNotSpecified, MetricAggregationTypeTotal}
}

// NfsAccessRuleAccess enumerates the values for nfs access rule access.
type NfsAccessRuleAccess string

const (
	// NfsAccessRuleAccessNo ...
	NfsAccessRuleAccessNo NfsAccessRuleAccess = "no"
	// NfsAccessRuleAccessRo ...
	NfsAccessRuleAccessRo NfsAccessRuleAccess = "ro"
	// NfsAccessRuleAccessRw ...
	NfsAccessRuleAccessRw NfsAccessRuleAccess = "rw"
)

// PossibleNfsAccessRuleAccessValues returns an array of possible values for the NfsAccessRuleAccess const type.
func PossibleNfsAccessRuleAccessValues() []NfsAccessRuleAccess {
	return []NfsAccessRuleAccess{NfsAccessRuleAccessNo, NfsAccessRuleAccessRo, NfsAccessRuleAccessRw}
}

// NfsAccessRuleScope enumerates the values for nfs access rule scope.
type NfsAccessRuleScope string

const (
	// Default ...
	Default NfsAccessRuleScope = "default"
	// Host ...
	Host NfsAccessRuleScope = "host"
	// Network ...
	Network NfsAccessRuleScope = "network"
)

// PossibleNfsAccessRuleScopeValues returns an array of possible values for the NfsAccessRuleScope const type.
func PossibleNfsAccessRuleScopeValues() []NfsAccessRuleScope {
	return []NfsAccessRuleScope{Default, Host, Network}
}

// ProvisioningStateType enumerates the values for provisioning state type.
type ProvisioningStateType string

const (
	// Cancelled ...
	Cancelled ProvisioningStateType = "Cancelled"
	// Creating ...
	Creating ProvisioningStateType = "Creating"
	// Deleting ...
	Deleting ProvisioningStateType = "Deleting"
	// Failed ...
	Failed ProvisioningStateType = "Failed"
	// Succeeded ...
	Succeeded ProvisioningStateType = "Succeeded"
	// Updating ...
	Updating ProvisioningStateType = "Updating"
)

// PossibleProvisioningStateTypeValues returns an array of possible values for the ProvisioningStateType const type.
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return []ProvisioningStateType{Cancelled, Creating, Deleting, Failed, Succeeded, Updating}
}

// ReasonCode enumerates the values for reason code.
type ReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID ReasonCode = "QuotaId"
)

// PossibleReasonCodeValues returns an array of possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{NotAvailableForSubscription, QuotaID}
}

// TargetType enumerates the values for target type.
type TargetType string

const (
	// TargetTypeClfs ...
	TargetTypeClfs TargetType = "clfs"
	// TargetTypeNfs3 ...
	TargetTypeNfs3 TargetType = "nfs3"
	// TargetTypeStorageTargetProperties ...
	TargetTypeStorageTargetProperties TargetType = "StorageTargetProperties"
	// TargetTypeUnknown ...
	TargetTypeUnknown TargetType = "unknown"
)

// PossibleTargetTypeValues returns an array of possible values for the TargetType const type.
func PossibleTargetTypeValues() []TargetType {
	return []TargetType{TargetTypeClfs, TargetTypeNfs3, TargetTypeStorageTargetProperties, TargetTypeUnknown}
}

// UsernameDownloadedType enumerates the values for username downloaded type.
type UsernameDownloadedType string

const (
	// UsernameDownloadedTypeError ...
	UsernameDownloadedTypeError UsernameDownloadedType = "Error"
	// UsernameDownloadedTypeNo ...
	UsernameDownloadedTypeNo UsernameDownloadedType = "No"
	// UsernameDownloadedTypeYes ...
	UsernameDownloadedTypeYes UsernameDownloadedType = "Yes"
)

// PossibleUsernameDownloadedTypeValues returns an array of possible values for the UsernameDownloadedType const type.
func PossibleUsernameDownloadedTypeValues() []UsernameDownloadedType {
	return []UsernameDownloadedType{UsernameDownloadedTypeError, UsernameDownloadedTypeNo, UsernameDownloadedTypeYes}
}

// UsernameSource enumerates the values for username source.
type UsernameSource string

const (
	// UsernameSourceAD ...
	UsernameSourceAD UsernameSource = "AD"
	// UsernameSourceFile ...
	UsernameSourceFile UsernameSource = "File"
	// UsernameSourceLDAP ...
	UsernameSourceLDAP UsernameSource = "LDAP"
	// UsernameSourceNone ...
	UsernameSourceNone UsernameSource = "None"
)

// PossibleUsernameSourceValues returns an array of possible values for the UsernameSource const type.
func PossibleUsernameSourceValues() []UsernameSource {
	return []UsernameSource{UsernameSourceAD, UsernameSourceFile, UsernameSourceLDAP, UsernameSourceNone}
}
