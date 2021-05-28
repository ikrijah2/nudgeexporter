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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AvailablePrivateEndpointTypesClient is the network Client
type AvailablePrivateEndpointTypesClient struct {
	BaseClient
}

// NewAvailablePrivateEndpointTypesClient creates an instance of the AvailablePrivateEndpointTypesClient client.
func NewAvailablePrivateEndpointTypesClient(subscriptionID string) AvailablePrivateEndpointTypesClient {
	return NewAvailablePrivateEndpointTypesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAvailablePrivateEndpointTypesClientWithBaseURI creates an instance of the AvailablePrivateEndpointTypesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewAvailablePrivateEndpointTypesClientWithBaseURI(baseURI string, subscriptionID string) AvailablePrivateEndpointTypesClient {
	return AvailablePrivateEndpointTypesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List returns all of the resource types that can be linked to a Private Endpoint in this subscription in this region.
// Parameters:
// location - the location of the domain name.
func (client AvailablePrivateEndpointTypesClient) List(ctx context.Context, location string) (result AvailablePrivateEndpointTypesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailablePrivateEndpointTypesClient.List")
		defer func() {
			sc := -1
			if result.apetr.Response.Response != nil {
				sc = result.apetr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.apetr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "List", resp, "Failure sending request")
		return
	}

	result.apetr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.apetr.hasNextLink() && result.apetr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AvailablePrivateEndpointTypesClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AvailablePrivateEndpointTypesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AvailablePrivateEndpointTypesClient) ListResponder(resp *http.Response) (result AvailablePrivateEndpointTypesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AvailablePrivateEndpointTypesClient) listNextResults(ctx context.Context, lastResults AvailablePrivateEndpointTypesResult) (result AvailablePrivateEndpointTypesResult, err error) {
	req, err := lastResults.availablePrivateEndpointTypesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailablePrivateEndpointTypesClient) ListComplete(ctx context.Context, location string) (result AvailablePrivateEndpointTypesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailablePrivateEndpointTypesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}

// ListByResourceGroup returns all of the resource types that can be linked to a Private Endpoint in this subscription
// in this region.
// Parameters:
// location - the location of the domain name.
// resourceGroupName - the name of the resource group.
func (client AvailablePrivateEndpointTypesClient) ListByResourceGroup(ctx context.Context, location string, resourceGroupName string) (result AvailablePrivateEndpointTypesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailablePrivateEndpointTypesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.apetr.Response.Response != nil {
				sc = result.apetr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, location, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.apetr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.apetr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.apetr.hasNextLink() && result.apetr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AvailablePrivateEndpointTypesClient) ListByResourceGroupPreparer(ctx context.Context, location string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":          autorest.Encode("path", location),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AvailablePrivateEndpointTypesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AvailablePrivateEndpointTypesClient) ListByResourceGroupResponder(resp *http.Response) (result AvailablePrivateEndpointTypesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AvailablePrivateEndpointTypesClient) listByResourceGroupNextResults(ctx context.Context, lastResults AvailablePrivateEndpointTypesResult) (result AvailablePrivateEndpointTypesResult, err error) {
	req, err := lastResults.availablePrivateEndpointTypesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailablePrivateEndpointTypesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailablePrivateEndpointTypesClient) ListByResourceGroupComplete(ctx context.Context, location string, resourceGroupName string) (result AvailablePrivateEndpointTypesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailablePrivateEndpointTypesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, location, resourceGroupName)
	return
}