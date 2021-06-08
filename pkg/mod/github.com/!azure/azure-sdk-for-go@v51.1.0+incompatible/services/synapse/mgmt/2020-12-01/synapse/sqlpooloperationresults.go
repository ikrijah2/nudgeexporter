package synapse

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SQLPoolOperationResultsClient is the azure Synapse Analytics Management Client
type SQLPoolOperationResultsClient struct {
	BaseClient
}

// NewSQLPoolOperationResultsClient creates an instance of the SQLPoolOperationResultsClient client.
func NewSQLPoolOperationResultsClient(subscriptionID string) SQLPoolOperationResultsClient {
	return NewSQLPoolOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolOperationResultsClientWithBaseURI creates an instance of the SQLPoolOperationResultsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewSQLPoolOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolOperationResultsClient {
	return SQLPoolOperationResultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetLocationHeaderResult get the status of a SQL pool operation
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// operationID - operation ID
func (client SQLPoolOperationResultsClient) GetLocationHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, operationID string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolOperationResultsClient.GetLocationHeaderResult")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolOperationResultsClient", "GetLocationHeaderResult", err.Error())
	}

	req, err := client.GetLocationHeaderResultPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationResultsClient", "GetLocationHeaderResult", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLocationHeaderResultSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationResultsClient", "GetLocationHeaderResult", resp, "Failure sending request")
		return
	}

	result, err = client.GetLocationHeaderResultResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationResultsClient", "GetLocationHeaderResult", resp, "Failure responding to request")
		return
	}

	return
}

// GetLocationHeaderResultPreparer prepares the GetLocationHeaderResult request.
func (client SQLPoolOperationResultsClient) GetLocationHeaderResultPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLocationHeaderResultSender sends the GetLocationHeaderResult request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolOperationResultsClient) GetLocationHeaderResultSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetLocationHeaderResultResponder handles the response to the GetLocationHeaderResult request. The method always
// closes the http.Response Body.
func (client SQLPoolOperationResultsClient) GetLocationHeaderResultResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
