package labservices

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

// EnvironmentsClient is the the Managed Labs Client.
type EnvironmentsClient struct {
	BaseClient
}

// NewEnvironmentsClient creates an instance of the EnvironmentsClient client.
func NewEnvironmentsClient(subscriptionID string) EnvironmentsClient {
	return NewEnvironmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEnvironmentsClientWithBaseURI creates an instance of the EnvironmentsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEnvironmentsClientWithBaseURI(baseURI string, subscriptionID string) EnvironmentsClient {
	return EnvironmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Claim claims the environment and assigns it to the user
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
func (client EnvironmentsClient) Claim(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.Claim")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ClaimPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Claim", nil, "Failure preparing request")
		return
	}

	resp, err := client.ClaimSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Claim", resp, "Failure sending request")
		return
	}

	result, err = client.ClaimResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Claim", resp, "Failure responding to request")
		return
	}

	return
}

// ClaimPreparer prepares the Claim request.
func (client EnvironmentsClient) ClaimPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}/claim", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ClaimSender sends the Claim request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) ClaimSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ClaimResponder handles the response to the Claim request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) ClaimResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create or replace an existing Environment.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
// environment - represents an environment instance
func (client EnvironmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, environment Environment) (result Environment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName, environment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EnvironmentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, environment Environment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", pathParameters),
		autorest.WithJSON(environment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) CreateOrUpdateResponder(resp *http.Response) (result Environment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete environment. This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
func (client EnvironmentsClient) Delete(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result EnvironmentsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EnvironmentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) DeleteSender(req *http.Request) (future EnvironmentsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client EnvironmentsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "labservices.EnvironmentsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("labservices.EnvironmentsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get environment
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
// expand - specify the $expand query. Example: 'properties($expand=networkInterface)'
func (client EnvironmentsClient) Get(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, expand string) (result Environment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client EnvironmentsClient) GetPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) GetResponder(resp *http.Response) (result Environment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list environments in a given environment setting.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// expand - specify the $expand query. Example: 'properties($expand=networkInterface)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client EnvironmentsClient) List(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationEnvironmentPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.List")
		defer func() {
			sc := -1
			if result.rwce.Response.Response != nil {
				sc = result.rwce.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwce.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.rwce, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rwce.hasNextLink() && result.rwce.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EnvironmentsClient) ListPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationEnvironment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EnvironmentsClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationEnvironment) (result ResponseWithContinuationEnvironment, err error) {
	req, err := lastResults.responseWithContinuationEnvironmentPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EnvironmentsClient) ListComplete(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationEnvironmentIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, expand, filter, top, orderby)
	return
}

// ResetPassword resets the user password on an environment This operation can take a while to complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
// resetPasswordPayload - represents the payload for resetting passwords.
func (client EnvironmentsClient) ResetPassword(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, resetPasswordPayload ResetPasswordPayload) (result EnvironmentsResetPasswordFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.ResetPassword")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resetPasswordPayload,
			Constraints: []validation.Constraint{{Target: "resetPasswordPayload.EnvironmentID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.EnvironmentsClient", "ResetPassword", err.Error())
	}

	req, err := client.ResetPasswordPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName, resetPasswordPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "ResetPassword", nil, "Failure preparing request")
		return
	}

	result, err = client.ResetPasswordSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "ResetPassword", nil, "Failure sending request")
		return
	}

	return
}

// ResetPasswordPreparer prepares the ResetPassword request.
func (client EnvironmentsClient) ResetPasswordPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, resetPasswordPayload ResetPasswordPayload) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}/resetPassword", pathParameters),
		autorest.WithJSON(resetPasswordPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetPasswordSender sends the ResetPassword request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) ResetPasswordSender(req *http.Request) (future EnvironmentsResetPasswordFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client EnvironmentsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "labservices.EnvironmentsResetPasswordFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("labservices.EnvironmentsResetPasswordFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// ResetPasswordResponder handles the response to the ResetPassword request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) ResetPasswordResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start starts an environment by starting all resources inside the environment. This operation can take a while to
// complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
func (client EnvironmentsClient) Start(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result EnvironmentsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Start", nil, "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client EnvironmentsClient) StartPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) StartSender(req *http.Request) (future EnvironmentsStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client EnvironmentsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "labservices.EnvironmentsStartFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("labservices.EnvironmentsStartFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops an environment by stopping all resources inside the environment This operation can take a while to
// complete
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
func (client EnvironmentsClient) Stop(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (result EnvironmentsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Stop", nil, "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client EnvironmentsClient) StopPreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) StopSender(req *http.Request) (future EnvironmentsStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client EnvironmentsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "labservices.EnvironmentsStopFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("labservices.EnvironmentsStopFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update modify properties of environments.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// labName - the name of the lab.
// environmentSettingName - the name of the environment Setting.
// environmentName - the name of the environment.
// environment - represents an environment instance
func (client EnvironmentsClient) Update(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, environment EnvironmentFragment) (result Environment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnvironmentsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labAccountName, labName, environmentSettingName, environmentName, environment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.EnvironmentsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EnvironmentsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, labName string, environmentSettingName string, environmentName string, environment EnvironmentFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":        autorest.Encode("path", environmentName),
		"environmentSettingName": autorest.Encode("path", environmentSettingName),
		"labAccountName":         autorest.Encode("path", labAccountName),
		"labName":                autorest.Encode("path", labName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", pathParameters),
		autorest.WithJSON(environment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EnvironmentsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EnvironmentsClient) UpdateResponder(resp *http.Response) (result Environment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
