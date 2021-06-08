package baremetalinfrastructure

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

// AzureBareMetalInstancesClient is the the BareMetalInfrastructure Management client
type AzureBareMetalInstancesClient struct {
	BaseClient
}

// NewAzureBareMetalInstancesClient creates an instance of the AzureBareMetalInstancesClient client.
func NewAzureBareMetalInstancesClient(subscriptionID string) AzureBareMetalInstancesClient {
	return NewAzureBareMetalInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAzureBareMetalInstancesClientWithBaseURI creates an instance of the AzureBareMetalInstancesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewAzureBareMetalInstancesClientWithBaseURI(baseURI string, subscriptionID string) AzureBareMetalInstancesClient {
	return AzureBareMetalInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes a Azure BareMetal instance with the specified subscription, resource group, and instance name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - name of the Azure BareMetal on Azure instance.
func (client AzureBareMetalInstancesClient) Delete(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result AzureBareMetalInstancesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
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
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, azureBareMetalInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AzureBareMetalInstancesClient) DeletePreparer(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"azureBareMetalInstanceName": autorest.Encode("path", azureBareMetalInstanceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) DeleteSender(req *http.Request) (future AzureBareMetalInstancesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AzureBareMetalInstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("baremetalinfrastructure.AzureBareMetalInstancesDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an Azure BareMetal instance for the specified subscription, resource group, and instance name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - name of the Azure BareMetal on Azure instance.
func (client AzureBareMetalInstancesClient) Get(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result AzureBareMetalInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.Get")
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
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, azureBareMetalInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AzureBareMetalInstancesClient) GetPreparer(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"azureBareMetalInstanceName": autorest.Encode("path", azureBareMetalInstanceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) GetResponder(resp *http.Response) (result AzureBareMetalInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of AzureBareMetal instances in the specified subscription and resource group
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client AzureBareMetalInstancesClient) List(ctx context.Context, resourceGroupName string) (result AzureBareMetalInstancesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.List")
		defer func() {
			sc := -1
			if result.abmilr.Response.Response != nil {
				sc = result.abmilr.Response.Response.StatusCode
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
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.abmilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.abmilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.abmilr.hasNextLink() && result.abmilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client AzureBareMetalInstancesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) ListResponder(resp *http.Response) (result AzureBareMetalInstancesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AzureBareMetalInstancesClient) listNextResults(ctx context.Context, lastResults AzureBareMetalInstancesListResult) (result AzureBareMetalInstancesListResult, err error) {
	req, err := lastResults.azureBareMetalInstancesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AzureBareMetalInstancesClient) ListComplete(ctx context.Context, resourceGroupName string) (result AzureBareMetalInstancesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListBySubscription gets a list of AzureBareMetal instances in the specified subscription. The operations returns
// various properties of each Azure BareMetal instance.
func (client AzureBareMetalInstancesClient) ListBySubscription(ctx context.Context) (result AzureBareMetalInstancesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.abmilr.Response.Response != nil {
				sc = result.abmilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.abmilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.abmilr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.abmilr.hasNextLink() && result.abmilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client AzureBareMetalInstancesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) ListBySubscriptionResponder(resp *http.Response) (result AzureBareMetalInstancesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client AzureBareMetalInstancesClient) listBySubscriptionNextResults(ctx context.Context, lastResults AzureBareMetalInstancesListResult) (result AzureBareMetalInstancesListResult, err error) {
	req, err := lastResults.azureBareMetalInstancesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client AzureBareMetalInstancesClient) ListBySubscriptionComplete(ctx context.Context) (result AzureBareMetalInstancesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Restart the operation to restart an AzureBareMetal instance
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - name of the Azure BareMetal on Azure instance.
func (client AzureBareMetalInstancesClient) Restart(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result AzureBareMetalInstancesRestartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.Restart")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
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
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "Restart", err.Error())
	}

	req, err := client.RestartPreparer(ctx, resourceGroupName, azureBareMetalInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Restart", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Restart", nil, "Failure sending request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client AzureBareMetalInstancesClient) RestartPreparer(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"azureBareMetalInstanceName": autorest.Encode("path", azureBareMetalInstanceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) RestartSender(req *http.Request) (future AzureBareMetalInstancesRestartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AzureBareMetalInstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesRestartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("baremetalinfrastructure.AzureBareMetalInstancesRestartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Shutdown the operation to shutdown an AzureBareMetal instance
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - name of the Azure BareMetal on Azure instance.
func (client AzureBareMetalInstancesClient) Shutdown(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result AzureBareMetalInstancesShutdownFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.Shutdown")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
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
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "Shutdown", err.Error())
	}

	req, err := client.ShutdownPreparer(ctx, resourceGroupName, azureBareMetalInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Shutdown", nil, "Failure preparing request")
		return
	}

	result, err = client.ShutdownSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Shutdown", nil, "Failure sending request")
		return
	}

	return
}

// ShutdownPreparer prepares the Shutdown request.
func (client AzureBareMetalInstancesClient) ShutdownPreparer(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"azureBareMetalInstanceName": autorest.Encode("path", azureBareMetalInstanceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}/shutdown", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ShutdownSender sends the Shutdown request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) ShutdownSender(req *http.Request) (future AzureBareMetalInstancesShutdownFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AzureBareMetalInstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesShutdownFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("baremetalinfrastructure.AzureBareMetalInstancesShutdownFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// ShutdownResponder handles the response to the Shutdown request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) ShutdownResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start the operation to start an AzureBareMetal instance
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - name of the Azure BareMetal on Azure instance.
func (client AzureBareMetalInstancesClient) Start(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (result AzureBareMetalInstancesStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
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
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "Start", err.Error())
	}

	req, err := client.StartPreparer(ctx, resourceGroupName, azureBareMetalInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client AzureBareMetalInstancesClient) StartPreparer(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"azureBareMetalInstanceName": autorest.Encode("path", azureBareMetalInstanceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) StartSender(req *http.Request) (future AzureBareMetalInstancesStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client AzureBareMetalInstancesClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesStartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("baremetalinfrastructure.AzureBareMetalInstancesStartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update patches the Tags field of a Azure BareMetal instance for the specified subscription, resource group, and
// instance name.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - name of the Azure BareMetal on Azure instance.
// tagsParameter - request body that only contains the new Tags field
func (client AzureBareMetalInstancesClient) Update(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, tagsParameter Tags) (result AzureBareMetalInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AzureBareMetalInstancesClient.Update")
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
		return result, validation.NewError("baremetalinfrastructure.AzureBareMetalInstancesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, azureBareMetalInstanceName, tagsParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "baremetalinfrastructure.AzureBareMetalInstancesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AzureBareMetalInstancesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, tagsParameter Tags) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"azureBareMetalInstanceName": autorest.Encode("path", azureBareMetalInstanceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}", pathParameters),
		autorest.WithJSON(tagsParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AzureBareMetalInstancesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AzureBareMetalInstancesClient) UpdateResponder(resp *http.Response) (result AzureBareMetalInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
