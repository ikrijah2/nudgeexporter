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

// RouteTablesResourceProviderClient is the the Windows Azure Network
// management API provides a RESTful set of web services that interact with
// Windows Azure Networks service to manage your network resrources. The API
// has entities that capture the relationship between an end user and the
// Windows Azure Networks service.
type RouteTablesResourceProviderClient struct {
	ResourceProviderClient
}

// NewRouteTablesResourceProviderClient creates an instance of the
// RouteTablesResourceProviderClient client.
func NewRouteTablesResourceProviderClient(subscriptionID string) RouteTablesResourceProviderClient {
	return NewRouteTablesResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRouteTablesResourceProviderClientWithBaseURI creates an instance of the
// RouteTablesResourceProviderClient client.
func NewRouteTablesResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) RouteTablesResourceProviderClient {
	return RouteTablesResourceProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put RouteTable operation creates/updates a route tablein
// the specified resource group.
//
// resourceGroupName is the name of the resource group. routeTableName is the
// name of the route table. parameters is parameters supplied to the
// create/update Route Table operation
func (client RouteTablesResourceProviderClient) CreateOrUpdate(resourceGroupName string, routeTableName string, parameters RouteTable) (result RouteTable, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, routeTableName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RouteTablesResourceProviderClient) CreateOrUpdatePreparer(resourceGroupName string, routeTableName string, parameters RouteTable) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"routeTableName":    url.QueryEscape(routeTableName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTablesResourceProviderClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RouteTablesResourceProviderClient) CreateOrUpdateResponder(resp *http.Response) (result RouteTable, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete RouteTable operation deletes the specifed Route Table
//
// resourceGroupName is the name of the resource group. routeTableName is the
// name of the route table.
func (client RouteTablesResourceProviderClient) Delete(resourceGroupName string, routeTableName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, routeTableName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RouteTablesResourceProviderClient) DeletePreparer(resourceGroupName string, routeTableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"routeTableName":    url.QueryEscape(routeTableName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTablesResourceProviderClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RouteTablesResourceProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get RouteTables operation retrieves information about the specified
// route table.
//
// resourceGroupName is the name of the resource group. routeTableName is the
// name of the route table.
func (client RouteTablesResourceProviderClient) Get(resourceGroupName string, routeTableName string) (result RouteTable, ae error) {
	req, err := client.GetPreparer(resourceGroupName, routeTableName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RouteTablesResourceProviderClient) GetPreparer(resourceGroupName string, routeTableName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"routeTableName":    url.QueryEscape(routeTableName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTablesResourceProviderClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RouteTablesResourceProviderClient) GetResponder(resp *http.Response) (result RouteTable, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the list RouteTables returns all route tables in a resource group
//
// resourceGroupName is the name of the resource group.
func (client RouteTablesResourceProviderClient) List(resourceGroupName string) (result RouteTableListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RouteTablesResourceProviderClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTablesResourceProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RouteTablesResourceProviderClient) ListResponder(resp *http.Response) (result RouteTableListResult, err error) {
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
func (client RouteTablesResourceProviderClient) ListNextResults(lastResults RouteTableListResult) (result RouteTableListResult, ae error) {
	req, err := lastResults.RouteTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "List", "Failure responding to next results request request")
	}

	return
}

// ListAll the list RouteTables returns all route tables in a subscription
func (client RouteTablesResourceProviderClient) ListAll() (result RouteTableListResult, ae error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "ListAll", "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "ListAll", "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "ListAll", "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client RouteTablesResourceProviderClient) ListAllPreparer() (*http.Request, error) {
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
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTablesResourceProviderClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client RouteTablesResourceProviderClient) ListAllResponder(resp *http.Response) (result RouteTableListResult, err error) {
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
func (client RouteTablesResourceProviderClient) ListAllNextResults(lastResults RouteTableListResult) (result RouteTableListResult, ae error) {
	req, err := lastResults.RouteTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "ListAll", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "ListAll", "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/RouteTablesResourceProviderClient", "ListAll", "Failure responding to next results request request")
	}

	return
}
