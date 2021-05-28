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

// ExpressRouteCircuitsResourceProviderClient is the the Windows Azure Network
// management API provides a RESTful set of web services that interact with
// Windows Azure Networks service to manage your network resrources. The API
// has entities that capture the relationship between an end user and the
// Windows Azure Networks service.
type ExpressRouteCircuitsResourceProviderClient struct {
	ResourceProviderClient
}

// NewExpressRouteCircuitsResourceProviderClient creates an instance of the
// ExpressRouteCircuitsResourceProviderClient client.
func NewExpressRouteCircuitsResourceProviderClient(subscriptionID string) ExpressRouteCircuitsResourceProviderClient {
	return NewExpressRouteCircuitsResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCircuitsResourceProviderClientWithBaseURI creates an
// instance of the ExpressRouteCircuitsResourceProviderClient client.
func NewExpressRouteCircuitsResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitsResourceProviderClient {
	return ExpressRouteCircuitsResourceProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put ExpressRouteCircuit operation creates/updates a
// ExpressRouteCircuit
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit. parameters is parameters supplied to the
// create/delete ExpressRouteCircuit operation
func (client ExpressRouteCircuitsResourceProviderClient) CreateOrUpdate(resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (result ExpressRouteCircuit, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, circuitName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitsResourceProviderClient) CreateOrUpdatePreparer(resourceGroupName string, circuitName string, parameters ExpressRouteCircuit) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteCircuit, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the delete ExpressRouteCircuit operation deletes the specified
// ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route Circuit.
func (client ExpressRouteCircuitsResourceProviderClient) Delete(resourceGroupName string, circuitName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteCircuitsResourceProviderClient) DeletePreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get ExpressRouteCircuit operation retreives information about the
// specified ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit.
func (client ExpressRouteCircuitsResourceProviderClient) Get(resourceGroupName string, circuitName string) (result ExpressRouteCircuit, ae error) {
	req, err := client.GetPreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCircuitsResourceProviderClient) GetPreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) GetResponder(resp *http.Response) (result ExpressRouteCircuit, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List ExpressRouteCircuit opertion retrieves all the
// ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ExpressRouteCircuitsResourceProviderClient) List(resourceGroupName string) (result ExpressRouteCircuitListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCircuitsResourceProviderClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) ListResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
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
func (client ExpressRouteCircuitsResourceProviderClient) ListNextResults(lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, ae error) {
	req, err := lastResults.ExpressRouteCircuitListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "List", "Failure responding to next results request request")
	}

	return
}

// ListAll the List ExpressRouteCircuit opertion retrieves all the
// ExpressRouteCircuits in a subscription.
func (client ExpressRouteCircuitsResourceProviderClient) ListAll() (result ExpressRouteCircuitListResult, ae error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListAll", "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListAll", "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListAll", "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client ExpressRouteCircuitsResourceProviderClient) ListAllPreparer() (*http.Request, error) {
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
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) ListAllResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
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
func (client ExpressRouteCircuitsResourceProviderClient) ListAllNextResults(lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, ae error) {
	req, err := lastResults.ExpressRouteCircuitListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListAll", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListAll", "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListAll", "Failure responding to next results request request")
	}

	return
}

// ListArpTable the ListArpTable from ExpressRouteCircuit opertion retrieves
// the currently advertised arp table associated with the
// ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit.
func (client ExpressRouteCircuitsResourceProviderClient) ListArpTable(resourceGroupName string, circuitName string) (result ExpressRouteCircuitsArpTableListResult, ae error) {
	req, err := client.ListArpTablePreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListArpTable", "Failure preparing request")
	}

	resp, err := client.ListArpTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListArpTable", "Failure sending request")
	}

	result, err = client.ListArpTableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListArpTable", "Failure responding to request")
	}

	return
}

// ListArpTablePreparer prepares the ListArpTable request.
func (client ExpressRouteCircuitsResourceProviderClient) ListArpTablePreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}arpTable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListArpTableSender sends the ListArpTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) ListArpTableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListArpTableResponder handles the response to the ListArpTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) ListArpTableResponder(resp *http.Response) (result ExpressRouteCircuitsArpTableListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListArpTableNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsResourceProviderClient) ListArpTableNextResults(lastResults ExpressRouteCircuitsArpTableListResult) (result ExpressRouteCircuitsArpTableListResult, ae error) {
	req, err := lastResults.ExpressRouteCircuitsArpTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListArpTable", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListArpTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListArpTable", "Failure sending next results request request")
	}

	result, err = client.ListArpTableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListArpTable", "Failure responding to next results request request")
	}

	return
}

// ListRoutesTable the ListRoutesTable from ExpressRouteCircuit opertion
// retrieves the currently advertised routes table associated with the
// ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit.
func (client ExpressRouteCircuitsResourceProviderClient) ListRoutesTable(resourceGroupName string, circuitName string) (result ExpressRouteCircuitsRoutesTableListResult, ae error) {
	req, err := client.ListRoutesTablePreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListRoutesTable", "Failure preparing request")
	}

	resp, err := client.ListRoutesTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListRoutesTable", "Failure sending request")
	}

	result, err = client.ListRoutesTableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListRoutesTable", "Failure responding to request")
	}

	return
}

// ListRoutesTablePreparer prepares the ListRoutesTable request.
func (client ExpressRouteCircuitsResourceProviderClient) ListRoutesTablePreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}routesTable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListRoutesTableSender sends the ListRoutesTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) ListRoutesTableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListRoutesTableResponder handles the response to the ListRoutesTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) ListRoutesTableResponder(resp *http.Response) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListRoutesTableNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsResourceProviderClient) ListRoutesTableNextResults(lastResults ExpressRouteCircuitsRoutesTableListResult) (result ExpressRouteCircuitsRoutesTableListResult, ae error) {
	req, err := lastResults.ExpressRouteCircuitsRoutesTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListRoutesTable", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListRoutesTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListRoutesTable", "Failure sending next results request request")
	}

	result, err = client.ListRoutesTableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListRoutesTable", "Failure responding to next results request request")
	}

	return
}

// ListStats the Liststats ExpressRouteCircuit opertion retrieves all the
// stats from a ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the loadBalancer.
func (client ExpressRouteCircuitsResourceProviderClient) ListStats(resourceGroupName string, circuitName string) (result ExpressRouteCircuitsStatsListResult, ae error) {
	req, err := client.ListStatsPreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListStats", "Failure preparing request")
	}

	resp, err := client.ListStatsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListStats", "Failure sending request")
	}

	result, err = client.ListStatsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListStats", "Failure responding to request")
	}

	return
}

// ListStatsPreparer prepares the ListStats request.
func (client ExpressRouteCircuitsResourceProviderClient) ListStatsPreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}stats"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListStatsSender sends the ListStats request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsResourceProviderClient) ListStatsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListStatsResponder handles the response to the ListStats request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsResourceProviderClient) ListStatsResponder(resp *http.Response) (result ExpressRouteCircuitsStatsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListStatsNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsResourceProviderClient) ListStatsNextResults(lastResults ExpressRouteCircuitsStatsListResult) (result ExpressRouteCircuitsStatsListResult, ae error) {
	req, err := lastResults.ExpressRouteCircuitsStatsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListStats", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListStatsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListStats", "Failure sending next results request request")
	}

	result, err = client.ListStatsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsResourceProviderClient", "ListStats", "Failure responding to next results request request")
	}

	return
}
