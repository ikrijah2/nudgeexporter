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
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// WorkflowsManagementClient is the client for the Workflows methods of the
// Logic service.
type WorkflowsManagementClient struct {
	ManagementClient
}

// NewWorkflowsManagementClient creates an instance of the
// WorkflowsManagementClient client.
func NewWorkflowsManagementClient(subscriptionID string) WorkflowsManagementClient {
	return NewWorkflowsManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowsManagementClientWithBaseURI creates an instance of the
// WorkflowsManagementClient client.
func NewWorkflowsManagementClientWithBaseURI(baseURI string, subscriptionID string) WorkflowsManagementClient {
	return WorkflowsManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. workflow is the workflow.
func (client WorkflowsManagementClient) CreateOrUpdate(resourceGroupName string, workflowName string, workflow Workflow) (result Workflow, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, workflowName, workflow)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkflowsManagementClient) CreateOrUpdatePreparer(resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithJSON(workflow),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusCreated)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) CreateOrUpdateResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsManagementClient) Delete(resourceGroupName string, workflowName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkflowsManagementClient) DeletePreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Disable disables a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsManagementClient) Disable(resourceGroupName string, workflowName string) (result autorest.Response, ae error) {
	req, err := client.DisablePreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Disable", "Failure preparing request")
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Disable", "Failure sending request")
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Disable", "Failure responding to request")
	}

	return
}

// DisablePreparer prepares the Disable request.
func (client WorkflowsManagementClient) DisablePreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/disable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) DisableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) DisableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Enable enables a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsManagementClient) Enable(resourceGroupName string, workflowName string) (result autorest.Response, ae error) {
	req, err := client.EnablePreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Enable", "Failure preparing request")
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Enable", "Failure sending request")
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Enable", "Failure responding to request")
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client WorkflowsManagementClient) EnablePreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/enable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) EnableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) EnableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsManagementClient) Get(resourceGroupName string, workflowName string) (result Workflow, ae error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowsManagementClient) GetPreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) GetResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets a list of workflows by resource group.
//
// resourceGroupName is the resource group name. top is the number of items to
// be included in the result. filter is the filter to apply on the operation.
func (client WorkflowsManagementClient) ListByResourceGroup(resourceGroupName string, top *int, filter string) (result WorkflowListResult, ae error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName, top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListByResourceGroup", "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListByResourceGroup", "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListByResourceGroup", "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WorkflowsManagementClient) ListByResourceGroupPreparer(resourceGroupName string, top *int, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) ListByResourceGroupResponder(resp *http.Response) (result WorkflowListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client WorkflowsManagementClient) ListByResourceGroupNextResults(lastResults WorkflowListResult) (result WorkflowListResult, ae error) {
	req, err := lastResults.WorkflowListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListByResourceGroup", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListByResourceGroup", "Failure sending next results request request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListByResourceGroup", "Failure responding to next results request request")
	}

	return
}

// ListBySubscription gets a list of workflows by subscription.
//
// top is the number of items to be included in the result. filter is the
// filter to apply on the operation.
func (client WorkflowsManagementClient) ListBySubscription(top *int, filter string) (result WorkflowListResult, ae error) {
	req, err := client.ListBySubscriptionPreparer(top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListBySubscription", "Failure preparing request")
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListBySubscription", "Failure sending request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListBySubscription", "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client WorkflowsManagementClient) ListBySubscriptionPreparer(top *int, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/workflows"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) ListBySubscriptionResponder(resp *http.Response) (result WorkflowListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionNextResults retrieves the next set of results, if any.
func (client WorkflowsManagementClient) ListBySubscriptionNextResults(lastResults WorkflowListResult) (result WorkflowListResult, ae error) {
	req, err := lastResults.WorkflowListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListBySubscription", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListBySubscription", "Failure sending next results request request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "ListBySubscription", "Failure responding to next results request request")
	}

	return
}

// Run runs a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. parameters is the parameters.
func (client WorkflowsManagementClient) Run(resourceGroupName string, workflowName string, parameters RunWorkflowParameters) (result WorkflowRun, ae error) {
	req, err := client.RunPreparer(resourceGroupName, workflowName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Run", "Failure preparing request")
	}

	resp, err := client.RunSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Run", "Failure sending request")
	}

	result, err = client.RunResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Run", "Failure responding to request")
	}

	return
}

// RunPreparer prepares the Run request.
func (client WorkflowsManagementClient) RunPreparer(resourceGroupName string, workflowName string, parameters RunWorkflowParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/run"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) RunSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusAccepted)
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) RunResponder(resp *http.Response) (result WorkflowRun, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. workflow is the workflow.
func (client WorkflowsManagementClient) Update(resourceGroupName string, workflowName string, workflow Workflow) (result Workflow, ae error) {
	req, err := client.UpdatePreparer(resourceGroupName, workflowName, workflow)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Update", "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Update", "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Update", "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WorkflowsManagementClient) UpdatePreparer(resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithJSON(workflow),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) UpdateResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate validates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. workflow is the workflow.
func (client WorkflowsManagementClient) Validate(resourceGroupName string, workflowName string, workflow Workflow) (result autorest.Response, ae error) {
	req, err := client.ValidatePreparer(resourceGroupName, workflowName, workflow)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Validate", "Failure preparing request")
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Validate", "Failure sending request")
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowsManagementClient", "Validate", "Failure responding to request")
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client WorkflowsManagementClient) ValidatePreparer(resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/validate"),
		autorest.WithJSON(workflow),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsManagementClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client WorkflowsManagementClient) ValidateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}