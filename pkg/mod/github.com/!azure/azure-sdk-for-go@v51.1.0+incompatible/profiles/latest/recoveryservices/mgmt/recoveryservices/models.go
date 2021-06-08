// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package recoveryservices

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2016-06-01/recoveryservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AuthType = original.AuthType

const (
	AAD                  AuthType = original.AAD
	AccessControlService AuthType = original.AccessControlService
	ACS                  AuthType = original.ACS
	AzureActiveDirectory AuthType = original.AzureActiveDirectory
	Invalid              AuthType = original.Invalid
)

type AuthTypeBasicResourceCertificateDetails = original.AuthTypeBasicResourceCertificateDetails

const (
	AuthTypeAccessControlService       AuthTypeBasicResourceCertificateDetails = original.AuthTypeAccessControlService
	AuthTypeAzureActiveDirectory       AuthTypeBasicResourceCertificateDetails = original.AuthTypeAzureActiveDirectory
	AuthTypeResourceCertificateDetails AuthTypeBasicResourceCertificateDetails = original.AuthTypeResourceCertificateDetails
)

type PrivateEndpointConnectionStatus = original.PrivateEndpointConnectionStatus

const (
	Approved     PrivateEndpointConnectionStatus = original.Approved
	Disconnected PrivateEndpointConnectionStatus = original.Disconnected
	Pending      PrivateEndpointConnectionStatus = original.Pending
	Rejected     PrivateEndpointConnectionStatus = original.Rejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStatePending   ProvisioningState = original.ProvisioningStatePending
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None           ResourceIdentityType = original.None
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type SkuName = original.SkuName

const (
	RS0      SkuName = original.RS0
	Standard SkuName = original.Standard
)

type TriggerType = original.TriggerType

const (
	ForcedUpgrade TriggerType = original.ForcedUpgrade
	UserTriggered TriggerType = original.UserTriggered
)

type UsagesUnit = original.UsagesUnit

const (
	Bytes          UsagesUnit = original.Bytes
	BytesPerSecond UsagesUnit = original.BytesPerSecond
	Count          UsagesUnit = original.Count
	CountPerSecond UsagesUnit = original.CountPerSecond
	Percent        UsagesUnit = original.Percent
	Seconds        UsagesUnit = original.Seconds
)

type VaultPrivateEndpointState = original.VaultPrivateEndpointState

const (
	VaultPrivateEndpointStateEnabled VaultPrivateEndpointState = original.VaultPrivateEndpointStateEnabled
	VaultPrivateEndpointStateNone    VaultPrivateEndpointState = original.VaultPrivateEndpointStateNone
)

type VaultUpgradeState = original.VaultUpgradeState

const (
	Failed     VaultUpgradeState = original.Failed
	InProgress VaultUpgradeState = original.InProgress
	Unknown    VaultUpgradeState = original.Unknown
	Upgraded   VaultUpgradeState = original.Upgraded
)

type BaseClient = original.BaseClient
type BasicResourceCertificateDetails = original.BasicResourceCertificateDetails
type CertificateRequest = original.CertificateRequest
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type Client = original.Client
type ClientDiscoveryDisplay = original.ClientDiscoveryDisplay
type ClientDiscoveryForLogSpecification = original.ClientDiscoveryForLogSpecification
type ClientDiscoveryForProperties = original.ClientDiscoveryForProperties
type ClientDiscoveryForServiceSpecification = original.ClientDiscoveryForServiceSpecification
type ClientDiscoveryResponse = original.ClientDiscoveryResponse
type ClientDiscoveryResponseIterator = original.ClientDiscoveryResponseIterator
type ClientDiscoveryResponsePage = original.ClientDiscoveryResponsePage
type ClientDiscoveryValueForSingleAPI = original.ClientDiscoveryValueForSingleAPI
type IdentityData = original.IdentityData
type JobsSummary = original.JobsSummary
type MonitoringSummary = original.MonitoringSummary
type NameInfo = original.NameInfo
type OperationsClient = original.OperationsClient
type PatchTrackedResource = original.PatchTrackedResource
type PatchVault = original.PatchVault
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionVaultProperties = original.PrivateEndpointConnectionVaultProperties
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResources = original.PrivateLinkResources
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkResourcesIterator = original.PrivateLinkResourcesIterator
type PrivateLinkResourcesPage = original.PrivateLinkResourcesPage
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type RawCertificateData = original.RawCertificateData
type RegisteredIdentitiesClient = original.RegisteredIdentitiesClient
type ReplicationUsage = original.ReplicationUsage
type ReplicationUsageList = original.ReplicationUsageList
type ReplicationUsagesClient = original.ReplicationUsagesClient
type Resource = original.Resource
type ResourceCertificateAndAadDetails = original.ResourceCertificateAndAadDetails
type ResourceCertificateAndAcsDetails = original.ResourceCertificateAndAcsDetails
type ResourceCertificateDetails = original.ResourceCertificateDetails
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type UpgradeDetails = original.UpgradeDetails
type UsagesClient = original.UsagesClient
type Vault = original.Vault
type VaultCertificateResponse = original.VaultCertificateResponse
type VaultCertificatesClient = original.VaultCertificatesClient
type VaultExtendedInfo = original.VaultExtendedInfo
type VaultExtendedInfoClient = original.VaultExtendedInfoClient
type VaultExtendedInfoResource = original.VaultExtendedInfoResource
type VaultList = original.VaultList
type VaultListIterator = original.VaultListIterator
type VaultListPage = original.VaultListPage
type VaultProperties = original.VaultProperties
type VaultUsage = original.VaultUsage
type VaultUsageList = original.VaultUsageList
type VaultsClient = original.VaultsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientDiscoveryResponseIterator(page ClientDiscoveryResponsePage) ClientDiscoveryResponseIterator {
	return original.NewClientDiscoveryResponseIterator(page)
}
func NewClientDiscoveryResponsePage(cur ClientDiscoveryResponse, getNextPage func(context.Context, ClientDiscoveryResponse) (ClientDiscoveryResponse, error)) ClientDiscoveryResponsePage {
	return original.NewClientDiscoveryResponsePage(cur, getNextPage)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesIterator(page PrivateLinkResourcesPage) PrivateLinkResourcesIterator {
	return original.NewPrivateLinkResourcesIterator(page)
}
func NewPrivateLinkResourcesPage(cur PrivateLinkResources, getNextPage func(context.Context, PrivateLinkResources) (PrivateLinkResources, error)) PrivateLinkResourcesPage {
	return original.NewPrivateLinkResourcesPage(cur, getNextPage)
}
func NewRegisteredIdentitiesClient(subscriptionID string) RegisteredIdentitiesClient {
	return original.NewRegisteredIdentitiesClient(subscriptionID)
}
func NewRegisteredIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) RegisteredIdentitiesClient {
	return original.NewRegisteredIdentitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicationUsagesClient(subscriptionID string) ReplicationUsagesClient {
	return original.NewReplicationUsagesClient(subscriptionID)
}
func NewReplicationUsagesClientWithBaseURI(baseURI string, subscriptionID string) ReplicationUsagesClient {
	return original.NewReplicationUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultCertificatesClient(subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClient(subscriptionID)
}
func NewVaultCertificatesClientWithBaseURI(baseURI string, subscriptionID string) VaultCertificatesClient {
	return original.NewVaultCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultExtendedInfoClient(subscriptionID string) VaultExtendedInfoClient {
	return original.NewVaultExtendedInfoClient(subscriptionID)
}
func NewVaultExtendedInfoClientWithBaseURI(baseURI string, subscriptionID string) VaultExtendedInfoClient {
	return original.NewVaultExtendedInfoClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultListIterator(page VaultListPage) VaultListIterator {
	return original.NewVaultListIterator(page)
}
func NewVaultListPage(cur VaultList, getNextPage func(context.Context, VaultList) (VaultList, error)) VaultListPage {
	return original.NewVaultListPage(cur, getNextPage)
}
func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthTypeBasicResourceCertificateDetailsValues() []AuthTypeBasicResourceCertificateDetails {
	return original.PossibleAuthTypeBasicResourceCertificateDetailsValues()
}
func PossibleAuthTypeValues() []AuthType {
	return original.PossibleAuthTypeValues()
}
func PossiblePrivateEndpointConnectionStatusValues() []PrivateEndpointConnectionStatus {
	return original.PossiblePrivateEndpointConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTriggerTypeValues() []TriggerType {
	return original.PossibleTriggerTypeValues()
}
func PossibleUsagesUnitValues() []UsagesUnit {
	return original.PossibleUsagesUnitValues()
}
func PossibleVaultPrivateEndpointStateValues() []VaultPrivateEndpointState {
	return original.PossibleVaultPrivateEndpointStateValues()
}
func PossibleVaultUpgradeStateValues() []VaultUpgradeState {
	return original.PossibleVaultUpgradeStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
