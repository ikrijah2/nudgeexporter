package logic

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

// IntegrationServiceEnvironmentManagedAPIOperationsClient is the REST API for Azure Logic Apps.
type IntegrationServiceEnvironmentManagedAPIOperationsClient struct {
	BaseClient
}

// NewIntegrationServiceEnvironmentManagedAPIOperationsClient creates an instance of the
// IntegrationServiceEnvironmentManagedAPIOperationsClient client.
func NewIntegrationServiceEnvironmentManagedAPIOperationsClient(subscriptionID string) IntegrationServiceEnvironmentManagedAPIOperationsClient {
	return NewIntegrationServiceEnvironmentManagedAPIOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationServiceEnvironmentManagedAPIOperationsClientWithBaseURI creates an instance of the
// IntegrationServiceEnvironmentManagedAPIOperationsClient client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIntegrationServiceEnvironmentManagedAPIOperationsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationServiceEnvironmentManagedAPIOperationsClient {
	return IntegrationServiceEnvironmentManagedAPIOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets the managed Api operations.
// Parameters:
// resourceGroup - the resource group.
// integrationServiceEnvironmentName - the integration service environment name.
// APIName - the api name.
func (client IntegrationServiceEnvironmentManagedAPIOperationsClient) List(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result APIOperationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentManagedAPIOperationsClient.List")
		defer func() {
			sc := -1
			if result.aolr.Response.Response != nil {
				sc = result.aolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroup, integrationServiceEnvironmentName, APIName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedAPIOperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedAPIOperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.aolr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedAPIOperationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.aolr.hasNextLink() && result.aolr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationServiceEnvironmentManagedAPIOperationsClient) ListPreparer(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiName":                           autorest.Encode("path", APIName),
		"integrationServiceEnvironmentName": autorest.Encode("path", integrationServiceEnvironmentName),
		"resourceGroup":                     autorest.Encode("path", resourceGroup),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}/apiOperations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationServiceEnvironmentManagedAPIOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationServiceEnvironmentManagedAPIOperationsClient) ListResponder(resp *http.Response) (result APIOperationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IntegrationServiceEnvironmentManagedAPIOperationsClient) listNextResults(ctx context.Context, lastResults APIOperationListResult) (result APIOperationListResult, err error) {
	req, err := lastResults.aPIOperationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedAPIOperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedAPIOperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationServiceEnvironmentManagedAPIOperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IntegrationServiceEnvironmentManagedAPIOperationsClient) ListComplete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, APIName string) (result APIOperationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationServiceEnvironmentManagedAPIOperationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroup, integrationServiceEnvironmentName, APIName)
	return
}