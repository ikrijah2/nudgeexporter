package network

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// VirtualNetworksResourceProviderClient is the the Windows Azure Network
// management API provides a RESTful set of web services that interact with
// Windows Azure Networks service to manage your network resrources. The API
// has entities that capture the relationship between an end user and the
// Windows Azure Networks service.
type VirtualNetworksResourceProviderClient struct {
	ResourceProviderClient
}

// NewVirtualNetworksResourceProviderClient creates an instance of the
// VirtualNetworksResourceProviderClient client.
func NewVirtualNetworksResourceProviderClient(subscriptionID string) VirtualNetworksResourceProviderClient {
	return NewVirtualNetworksResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworksResourceProviderClientWithBaseURI creates an instance of
// the VirtualNetworksResourceProviderClient client.
func NewVirtualNetworksResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworksResourceProviderClient {
	return VirtualNetworksResourceProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put VirtualNetwork operation creates/updates a virtual
// network in the specified resource group.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. parameters is parameters supplied to the
// create/update Virtual Network operation
func (client VirtualNetworksResourceProviderClient) CreateOrUpdate(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (result VirtualNetwork, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworksResourceProviderClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksResourceProviderClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworksResourceProviderClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete VirtualNetwork operation deletes the specifed virtual
// network
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network.
func (client VirtualNetworksResourceProviderClient) Delete(resourceGroupName string, virtualNetworkName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, virtualNetworkName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworksResourceProviderClient) DeletePreparer(resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksResourceProviderClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworksResourceProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get VirtualNetwork operation retrieves information about the
// specified virtual network.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network.
func (client VirtualNetworksResourceProviderClient) Get(resourceGroupName string, virtualNetworkName string) (result VirtualNetwork, ae error) {
	req, err := client.GetPreparer(resourceGroupName, virtualNetworkName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworksResourceProviderClient) GetPreparer(resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksResourceProviderClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworksResourceProviderClient) GetResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the list VirtualNetwork returns all Virtual Networks in a resource
// group
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworksResourceProviderClient) List(resourceGroupName string) (result VirtualNetworkListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworksResourceProviderClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksResourceProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworksResourceProviderClient) ListResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client VirtualNetworksResourceProviderClient) ListNextResults(lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, ae error) {
	req, err := lastResults.VirtualNetworkListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "List", "Failure responding to next results request request")
	}

	return
}

// ListAll the list VirtualNetwork returns all Virtual Networks in a
// subscription
func (client VirtualNetworksResourceProviderClient) ListAll() (result VirtualNetworkListResult, ae error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "ListAll", "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "ListAll", "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "ListAll", "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client VirtualNetworksResourceProviderClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualnetworks"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksResourceProviderClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client VirtualNetworksResourceProviderClient) ListAllResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client VirtualNetworksResourceProviderClient) ListAllNextResults(lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, ae error) {
	req, err := lastResults.VirtualNetworkListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "ListAll", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "ListAll", "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/VirtualNetworksResourceProviderClient", "ListAll", "Failure responding to next results request request")
	}

	return
}
