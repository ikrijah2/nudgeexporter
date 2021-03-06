package scheduler

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

// JobsManagementClient is the client for the Jobs methods of the Scheduler
// service.
type JobsManagementClient struct {
	ManagementClient
}

// NewJobsManagementClient creates an instance of the JobsManagementClient
// client.
func NewJobsManagementClient(subscriptionID string) JobsManagementClient {
	return NewJobsManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobsManagementClientWithBaseURI creates an instance of the
// JobsManagementClient client.
func NewJobsManagementClientWithBaseURI(baseURI string, subscriptionID string) JobsManagementClient {
	return JobsManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate provisions a new job or updates an existing job.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobName is the job name. job is the job definition.
func (client JobsManagementClient) CreateOrUpdate(resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (result JobDefinition, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, jobCollectionName, jobName, job)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobsManagementClient) CreateOrUpdatePreparer(resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": url.QueryEscape(jobCollectionName),
		"jobName":           url.QueryEscape(jobName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"),
		autorest.WithJSON(job),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobsManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusCreated)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client JobsManagementClient) CreateOrUpdateResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a job.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobName is the job name.
func (client JobsManagementClient) Delete(resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, jobCollectionName, jobName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobsManagementClient) DeletePreparer(resourceGroupName string, jobCollectionName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": url.QueryEscape(jobCollectionName),
		"jobName":           url.QueryEscape(jobName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobsManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobsManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a job.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobName is the job name.
func (client JobsManagementClient) Get(resourceGroupName string, jobCollectionName string, jobName string) (result JobDefinition, ae error) {
	req, err := client.GetPreparer(resourceGroupName, jobCollectionName, jobName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobsManagementClient) GetPreparer(resourceGroupName string, jobCollectionName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": url.QueryEscape(jobCollectionName),
		"jobName":           url.QueryEscape(jobName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsManagementClient) GetResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all jobs under the specified job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. top is the number of jobs to request, in the of range
// [1..100]. skip is the (0-based) index of the job history list from which
// to begin requesting entries.
func (client JobsManagementClient) List(resourceGroupName string, jobCollectionName string, top *int, skip *int) (result JobListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, jobCollectionName, top, skip)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client JobsManagementClient) ListPreparer(resourceGroupName string, jobCollectionName string, top *int, skip *int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": url.QueryEscape(jobCollectionName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if skip != nil {
		queryParameters["$skip"] = skip
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client JobsManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client JobsManagementClient) ListResponder(resp *http.Response) (result JobListResult, err error) {
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
func (client JobsManagementClient) ListNextResults(lastResults JobListResult) (result JobListResult, ae error) {
	req, err := lastResults.JobListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "List", "Failure responding to next results request request")
	}

	return
}

// ListJobHistory lists job history.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobName is the job name. top is the number of job history
// to request, in the of range [1..100]. skip is the (0-based) index of the
// job history list from which to begin requesting entries.
func (client JobsManagementClient) ListJobHistory(resourceGroupName string, jobCollectionName string, jobName string, top *int, skip *int) (result JobHistoryListResult, ae error) {
	req, err := client.ListJobHistoryPreparer(resourceGroupName, jobCollectionName, jobName, top, skip)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "ListJobHistory", "Failure preparing request")
	}

	resp, err := client.ListJobHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "ListJobHistory", "Failure sending request")
	}

	result, err = client.ListJobHistoryResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "ListJobHistory", "Failure responding to request")
	}

	return
}

// ListJobHistoryPreparer prepares the ListJobHistory request.
func (client JobsManagementClient) ListJobHistoryPreparer(resourceGroupName string, jobCollectionName string, jobName string, top *int, skip *int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": url.QueryEscape(jobCollectionName),
		"jobName":           url.QueryEscape(jobName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if skip != nil {
		queryParameters["$skip"] = skip
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/history"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListJobHistorySender sends the ListJobHistory request. The method will close the
// http.Response Body if it receives an error.
func (client JobsManagementClient) ListJobHistorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListJobHistoryResponder handles the response to the ListJobHistory request. The method always
// closes the http.Response Body.
func (client JobsManagementClient) ListJobHistoryResponder(resp *http.Response) (result JobHistoryListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListJobHistoryNextResults retrieves the next set of results, if any.
func (client JobsManagementClient) ListJobHistoryNextResults(lastResults JobHistoryListResult) (result JobHistoryListResult, ae error) {
	req, err := lastResults.JobHistoryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "ListJobHistory", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListJobHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "ListJobHistory", "Failure sending next results request request")
	}

	result, err = client.ListJobHistoryResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "ListJobHistory", "Failure responding to next results request request")
	}

	return
}

// Patch patches an existing job.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobName is the job name. job is the job definition.
func (client JobsManagementClient) Patch(resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (result JobDefinition, ae error) {
	req, err := client.PatchPreparer(resourceGroupName, jobCollectionName, jobName, job)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Patch", "Failure preparing request")
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Patch", "Failure sending request")
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Patch", "Failure responding to request")
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client JobsManagementClient) PatchPreparer(resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": url.QueryEscape(jobCollectionName),
		"jobName":           url.QueryEscape(jobName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"),
		autorest.WithJSON(job),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client JobsManagementClient) PatchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client JobsManagementClient) PatchResponder(resp *http.Response) (result JobDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Run runs a job.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobName is the job name.
func (client JobsManagementClient) Run(resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, ae error) {
	req, err := client.RunPreparer(resourceGroupName, jobCollectionName, jobName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Run", "Failure preparing request")
	}

	resp, err := client.RunSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Run", "Failure sending request")
	}

	result, err = client.RunResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "scheduler/JobsManagementClient", "Run", "Failure responding to request")
	}

	return
}

// RunPreparer prepares the Run request.
func (client JobsManagementClient) RunPreparer(resourceGroupName string, jobCollectionName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": url.QueryEscape(jobCollectionName),
		"jobName":           url.QueryEscape(jobName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/run"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client JobsManagementClient) RunSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client JobsManagementClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
