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

package hardwaresecuritymodules

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/hardwaresecuritymodules/mgmt/2018-10-31-preview/hardwaresecuritymodules"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type JSONWebKeyType = original.JSONWebKeyType

const (
	Allocating    JSONWebKeyType = original.Allocating
	CheckingQuota JSONWebKeyType = original.CheckingQuota
	Connecting    JSONWebKeyType = original.Connecting
	Deleting      JSONWebKeyType = original.Deleting
	Failed        JSONWebKeyType = original.Failed
	Provisioning  JSONWebKeyType = original.Provisioning
	Succeeded     JSONWebKeyType = original.Succeeded
)

type Name = original.Name

const (
	SafeNetLunaNetworkHSMA790 Name = original.SafeNetLunaNetworkHSMA790
)

type APIEntityReference = original.APIEntityReference
type BaseClient = original.BaseClient
type DedicatedHsm = original.DedicatedHsm
type DedicatedHsmClient = original.DedicatedHsmClient
type DedicatedHsmCreateOrUpdateFuture = original.DedicatedHsmCreateOrUpdateFuture
type DedicatedHsmDeleteFuture = original.DedicatedHsmDeleteFuture
type DedicatedHsmError = original.DedicatedHsmError
type DedicatedHsmListResult = original.DedicatedHsmListResult
type DedicatedHsmListResultIterator = original.DedicatedHsmListResultIterator
type DedicatedHsmListResultPage = original.DedicatedHsmListResultPage
type DedicatedHsmOperation = original.DedicatedHsmOperation
type DedicatedHsmOperationDisplay = original.DedicatedHsmOperationDisplay
type DedicatedHsmOperationListResult = original.DedicatedHsmOperationListResult
type DedicatedHsmPatchParameters = original.DedicatedHsmPatchParameters
type DedicatedHsmProperties = original.DedicatedHsmProperties
type DedicatedHsmUpdateFuture = original.DedicatedHsmUpdateFuture
type Error = original.Error
type NetworkInterface = original.NetworkInterface
type NetworkProfile = original.NetworkProfile
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type ResourceListResult = original.ResourceListResult
type Sku = original.Sku

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDedicatedHsmClient(subscriptionID string) DedicatedHsmClient {
	return original.NewDedicatedHsmClient(subscriptionID)
}
func NewDedicatedHsmClientWithBaseURI(baseURI string, subscriptionID string) DedicatedHsmClient {
	return original.NewDedicatedHsmClientWithBaseURI(baseURI, subscriptionID)
}
func NewDedicatedHsmListResultIterator(page DedicatedHsmListResultPage) DedicatedHsmListResultIterator {
	return original.NewDedicatedHsmListResultIterator(page)
}
func NewDedicatedHsmListResultPage(cur DedicatedHsmListResult, getNextPage func(context.Context, DedicatedHsmListResult) (DedicatedHsmListResult, error)) DedicatedHsmListResultPage {
	return original.NewDedicatedHsmListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return original.PossibleJSONWebKeyTypeValues()
}
func PossibleNameValues() []Name {
	return original.PossibleNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
