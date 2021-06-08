package dns

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

// ZonesManagementClient is the client for managing DNS zones and record.
type ZonesManagementClient struct {
	ManagementClient
}

// NewZonesManagementClient creates an instance of the ZonesManagementClient
// client.
func NewZonesManagementClient(subscriptionID string) ZonesManagementClient {
	return NewZonesManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewZonesManagementClientWithBaseURI creates an instance of the
// ZonesManagementClient client.
func NewZonesManagementClientWithBaseURI(baseURI string, subscriptionID string) ZonesManagementClient {
	return ZonesManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a DNS zone within a resource group.
//
// resourceGroupName is the name of the resource group. zoneName is the name
// of the zone without a terminating dot. parameters is parameters supplied
// to the CreateOrUpdate operation. ifMatch is the etag of Zone. ifNoneMatch
// is defines the If-None-Match condition. Set to '*' to force
// Create-If-Not-Exist. Other values will be ignored.
func (client ZonesManagementClient) CreateOrUpdate(resourceGroupName string, zoneName string, parameters ZoneCreateOrUpdateParameters, ifMatch string, ifNoneMatch string) (result Zone, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, zoneName, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ZonesManagementClient) CreateOrUpdatePreparer(resourceGroupName string, zoneName string, parameters ZoneCreateOrUpdateParameters, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"zoneName":          url.QueryEscape(zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ZonesManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ZonesManagementClient) CreateOrUpdateResponder(resp *http.Response) (result Zone, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete removes a DNS zone from a resource group.
//
// resourceGroupName is the name of the resource group. zoneName is the name
// of the zone without a terminating dot. ifMatch is defines the If-Match
// condition. The delete operation will be performed only if the ETag of the
// zone on the server matches this value.
func (client ZonesManagementClient) Delete(resourceGroupName string, zoneName string, ifMatch string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, zoneName, ifMatch)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ZonesManagementClient) DeletePreparer(resourceGroupName string, zoneName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"zoneName":          url.QueryEscape(zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ZonesManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ZonesManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name
// of the zone without a terminating dot.
func (client ZonesManagementClient) Get(resourceGroupName string, zoneName string) (result Zone, ae error) {
	req, err := client.GetPreparer(resourceGroupName, zoneName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ZonesManagementClient) GetPreparer(resourceGroupName string, zoneName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"zoneName":          url.QueryEscape(zoneName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ZonesManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ZonesManagementClient) GetResponder(resp *http.Response) (result Zone, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListZonesInResourceGroup lists the DNS zones within a resource group.
//
// resourceGroupName is the name of the resource group. top is query
// parameters. If null is passed returns the default number of zones. filter
// is the filter to apply on the operation.
func (client ZonesManagementClient) ListZonesInResourceGroup(resourceGroupName string, top string, filter string) (result ZoneListResult, ae error) {
	req, err := client.ListZonesInResourceGroupPreparer(resourceGroupName, top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "ListZonesInResourceGroup", "Failure preparing request")
	}

	resp, err := client.ListZonesInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "ListZonesInResourceGroup", "Failure sending request")
	}

	result, err = client.ListZonesInResourceGroupResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "ListZonesInResourceGroup", "Failure responding to request")
	}

	return
}

// ListZonesInResourceGroupPreparer prepares the ListZonesInResourceGroup request.
func (client ZonesManagementClient) ListZonesInResourceGroupPreparer(resourceGroupName string, top string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(top) > 0 {
		queryParameters["$top"] = top
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListZonesInResourceGroupSender sends the ListZonesInResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ZonesManagementClient) ListZonesInResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListZonesInResourceGroupResponder handles the response to the ListZonesInResourceGroup request. The method always
// closes the http.Response Body.
func (client ZonesManagementClient) ListZonesInResourceGroupResponder(resp *http.Response) (result ZoneListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListZonesInSubscription lists the DNS zones within a resource group.
//
// top is query parameters. If null is passed returns the default number of
// zones. filter is the filter to apply on the operation.
func (client ZonesManagementClient) ListZonesInSubscription(top string, filter string) (result ZoneListResult, ae error) {
	req, err := client.ListZonesInSubscriptionPreparer(top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "ListZonesInSubscription", "Failure preparing request")
	}

	resp, err := client.ListZonesInSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "ListZonesInSubscription", "Failure sending request")
	}

	result, err = client.ListZonesInSubscriptionResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "dns/ZonesManagementClient", "ListZonesInSubscription", "Failure responding to request")
	}

	return
}

// ListZonesInSubscriptionPreparer prepares the ListZonesInSubscription request.
func (client ZonesManagementClient) ListZonesInSubscriptionPreparer(top string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(top) > 0 {
		queryParameters["$top"] = top
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListZonesInSubscriptionSender sends the ListZonesInSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ZonesManagementClient) ListZonesInSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListZonesInSubscriptionResponder handles the response to the ListZonesInSubscription request. The method always
// closes the http.Response Body.
func (client ZonesManagementClient) ListZonesInSubscriptionResponder(resp *http.Response) (result ZoneListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
