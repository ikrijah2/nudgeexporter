package compute

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

// AvailabilitySetsManagementClient is the the Compute Management Client.
type AvailabilitySetsManagementClient struct {
	ManagementClient
}

// NewAvailabilitySetsManagementClient creates an instance of the
// AvailabilitySetsManagementClient client.
func NewAvailabilitySetsManagementClient(subscriptionID string) AvailabilitySetsManagementClient {
	return NewAvailabilitySetsManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAvailabilitySetsManagementClientWithBaseURI creates an instance of the
// AvailabilitySetsManagementClient client.
func NewAvailabilitySetsManagementClientWithBaseURI(baseURI string, subscriptionID string) AvailabilitySetsManagementClient {
	return AvailabilitySetsManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update the availability set.
//
// resourceGroupName is the name of the resource group. name is parameters
// supplied to the Create Availability Set operation. parameters is
// parameters supplied to the Create Availability Set operation.
func (client AvailabilitySetsManagementClient) CreateOrUpdate(resourceGroupName string, name string, parameters AvailabilitySet) (result AvailabilitySet, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, name, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AvailabilitySetsManagementClient) CreateOrUpdatePreparer(resourceGroupName string, name string, parameters AvailabilitySet) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              url.QueryEscape(name),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{name}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsManagementClient) CreateOrUpdateResponder(resp *http.Response) (result AvailabilitySet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the availability set.
//
// resourceGroupName is the name of the resource group. availabilitySetName is
// the name of the availability set.
func (client AvailabilitySetsManagementClient) Delete(resourceGroupName string, availabilitySetName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, availabilitySetName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AvailabilitySetsManagementClient) DeletePreparer(resourceGroupName string, availabilitySetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"availabilitySetName": url.QueryEscape(availabilitySetName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the operation to get the availability set.
//
// resourceGroupName is the name of the resource group. availabilitySetName is
// the name of the availability set.
func (client AvailabilitySetsManagementClient) Get(resourceGroupName string, availabilitySetName string) (result AvailabilitySet, ae error) {
	req, err := client.GetPreparer(resourceGroupName, availabilitySetName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AvailabilitySetsManagementClient) GetPreparer(resourceGroupName string, availabilitySetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"availabilitySetName": url.QueryEscape(availabilitySetName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsManagementClient) GetResponder(resp *http.Response) (result AvailabilitySet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the operation to list the availability sets.
//
// resourceGroupName is the name of the resource group.
func (client AvailabilitySetsManagementClient) List(resourceGroupName string) (result AvailabilitySetListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AvailabilitySetsManagementClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsManagementClient) ListResponder(resp *http.Response) (result AvailabilitySetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAvailableSizes lists virtual-machine-sizes available to be used for an
// availability set.
//
// resourceGroupName is the name of the resource group. availabilitySetName is
// the name of the availability set.
func (client AvailabilitySetsManagementClient) ListAvailableSizes(resourceGroupName string, availabilitySetName string) (result VirtualMachineSizeListResult, ae error) {
	req, err := client.ListAvailableSizesPreparer(resourceGroupName, availabilitySetName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "ListAvailableSizes", "Failure preparing request")
	}

	resp, err := client.ListAvailableSizesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "ListAvailableSizes", "Failure sending request")
	}

	result, err = client.ListAvailableSizesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/AvailabilitySetsManagementClient", "ListAvailableSizes", "Failure responding to request")
	}

	return
}

// ListAvailableSizesPreparer prepares the ListAvailableSizes request.
func (client AvailabilitySetsManagementClient) ListAvailableSizesPreparer(resourceGroupName string, availabilitySetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"availabilitySetName": url.QueryEscape(availabilitySetName),
		"resourceGroupName":   url.QueryEscape(resourceGroupName),
		"subscriptionId":      url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAvailableSizesSender sends the ListAvailableSizes request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilitySetsManagementClient) ListAvailableSizesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListAvailableSizesResponder handles the response to the ListAvailableSizes request. The method always
// closes the http.Response Body.
func (client AvailabilitySetsManagementClient) ListAvailableSizesResponder(resp *http.Response) (result VirtualMachineSizeListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
