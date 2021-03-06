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

// WorkflowTriggersManagementClient is the client for the WorkflowTriggers
// methods of the Logic service.
type WorkflowTriggersManagementClient struct {
	ManagementClient
}

// NewWorkflowTriggersManagementClient creates an instance of the
// WorkflowTriggersManagementClient client.
func NewWorkflowTriggersManagementClient(subscriptionID string) WorkflowTriggersManagementClient {
	return NewWorkflowTriggersManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowTriggersManagementClientWithBaseURI creates an instance of the
// WorkflowTriggersManagementClient client.
func NewWorkflowTriggersManagementClientWithBaseURI(baseURI string, subscriptionID string) WorkflowTriggersManagementClient {
	return WorkflowTriggersManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a workflow trigger.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. triggerName is the workflow trigger name.
func (client WorkflowTriggersManagementClient) Get(resourceGroupName string, workflowName string, triggerName string) (result WorkflowTrigger, ae error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName, triggerName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowTriggersManagementClient) GetPreparer(resourceGroupName string, workflowName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"triggerName":       url.QueryEscape(triggerName),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersManagementClient) GetResponder(resp *http.Response) (result WorkflowTrigger, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow triggers.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. top is the number of items to be included in the result. filter is
// the filter to apply on the operation.
func (client WorkflowTriggersManagementClient) List(resourceGroupName string, workflowName string, top *int, filter string) (result WorkflowTriggerListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, workflowName, top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowTriggersManagementClient) ListPreparer(resourceGroupName string, workflowName string, top *int, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersManagementClient) ListResponder(resp *http.Response) (result WorkflowTriggerListResult, err error) {
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
func (client WorkflowTriggersManagementClient) ListNextResults(lastResults WorkflowTriggerListResult) (result WorkflowTriggerListResult, ae error) {
	req, err := lastResults.WorkflowTriggerListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "List", "Failure responding to next results request request")
	}

	return
}

// Run runs a workflow trigger.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. triggerName is the workflow trigger name.
func (client WorkflowTriggersManagementClient) Run(resourceGroupName string, workflowName string, triggerName string) (result autorest.Response, ae error) {
	req, err := client.RunPreparer(resourceGroupName, workflowName, triggerName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "Run", "Failure preparing request")
	}

	resp, err := client.RunSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "Run", "Failure sending request")
	}

	result, err = client.RunResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "logic/WorkflowTriggersManagementClient", "Run", "Failure responding to request")
	}

	return
}

// RunPreparer prepares the Run request.
func (client WorkflowTriggersManagementClient) RunPreparer(resourceGroupName string, workflowName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"triggerName":       url.QueryEscape(triggerName),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/run"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggersManagementClient) RunSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client WorkflowTriggersManagementClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
