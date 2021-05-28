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

package skus

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-09-01/skus"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ResourceSkuCapacityScaleType = original.ResourceSkuCapacityScaleType

const (
	Automatic ResourceSkuCapacityScaleType = original.Automatic
	Manual    ResourceSkuCapacityScaleType = original.Manual
	None      ResourceSkuCapacityScaleType = original.None
)

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
	Zone     ResourceSkuRestrictionsType = original.Zone
)

type BaseClient = original.BaseClient
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuCapacity = original.ResourceSkuCapacity
type ResourceSkuCosts = original.ResourceSkuCosts
type ResourceSkuLocationInfo = original.ResourceSkuLocationInfo
type ResourceSkuRestrictionInfo = original.ResourceSkuRestrictionInfo
type ResourceSkuRestrictions = original.ResourceSkuRestrictions
type ResourceSkusClient = original.ResourceSkusClient
type ResourceSkusResult = original.ResourceSkusResult
type ResourceSkusResultIterator = original.ResourceSkusResultIterator
type ResourceSkusResultPage = original.ResourceSkusResultPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewResourceSkusClient(subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClient(subscriptionID)
}
func NewResourceSkusClientWithBaseURI(baseURI string, subscriptionID string) ResourceSkusClient {
	return original.NewResourceSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkusResultIterator(page ResourceSkusResultPage) ResourceSkusResultIterator {
	return original.NewResourceSkusResultIterator(page)
}
func NewResourceSkusResultPage(cur ResourceSkusResult, getNextPage func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)) ResourceSkusResultPage {
	return original.NewResourceSkusResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleResourceSkuCapacityScaleTypeValues() []ResourceSkuCapacityScaleType {
	return original.PossibleResourceSkuCapacityScaleTypeValues()
}
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return original.PossibleResourceSkuRestrictionsReasonCodeValues()
}
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return original.PossibleResourceSkuRestrictionsTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}