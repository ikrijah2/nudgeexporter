package job

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
	"github.com/satori/go.uuid"
	"net/http"
)

// Client is the creates an Azure Data Lake Analytics job client.
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return Client{New()}
}

// Build builds (compiles) the specified job in the specified Data Lake Analytics account for job correctness and
// validation.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// parameters - the parameters to build a job.
func (client Client) Build(ctx context.Context, accountName string, parameters Information) (result Information, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Build")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("job.Client", "Build", err.Error())
	}

	req, err := client.BuildPreparer(ctx, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Build", nil, "Failure preparing request")
		return
	}

	resp, err := client.BuildSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "Build", resp, "Failure sending request")
		return
	}

	result, err = client.BuildResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Build", resp, "Failure responding to request")
		return
	}

	return
}

// BuildPreparer prepares the Build request.
func (client Client) BuildPreparer(ctx context.Context, accountName string, parameters Information) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2016-03-20-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ErrorMessage = nil
	parameters.SubmitTime = nil
	parameters.StartTime = nil
	parameters.EndTime = nil
	parameters.State = ""
	parameters.Result = ""
	parameters.LogFolder = nil
	parameters.StateAuditRecords = nil
	parameters.HierarchyQueueNode = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/BuildJob"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BuildSender sends the Build request. The method will close the
// http.Response Body if it receives an error.
func (client Client) BuildSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// BuildResponder handles the response to the Build request. The method always
// closes the http.Response Body.
func (client Client) BuildResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Cancel cancels the running job specified by the job ID.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// jobIdentity - jobInfo ID to cancel.
func (client Client) Cancel(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Cancel")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "job.Client", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Cancel", resp, "Failure responding to request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client Client) CancelPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2016-03-20-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}/CancelJob", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client Client) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create submits a job to the specified Data Lake Analytics account.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// jobIdentity - the job ID (a GUID) for the job being submitted.
// parameters - the parameters to submit a job.
func (client Client) Create(ctx context.Context, accountName string, jobIdentity uuid.UUID, parameters Information) (result Information, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("job.Client", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, accountName, jobIdentity, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID, parameters Information) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2016-03-20-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ErrorMessage = nil
	parameters.SubmitTime = nil
	parameters.StartTime = nil
	parameters.EndTime = nil
	parameters.State = ""
	parameters.Result = ""
	parameters.LogFolder = nil
	parameters.StateAuditRecords = nil
	parameters.HierarchyQueueNode = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the job information for the specified job ID.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// jobIdentity - jobInfo ID.
func (client Client) Get(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result Information, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2016-03-20-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDebugDataPath gets the job debug data information specified by the job ID.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// jobIdentity - jobInfo ID.
func (client Client) GetDebugDataPath(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result DataPath, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetDebugDataPath")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDebugDataPathPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetDebugDataPath", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDebugDataPathSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "GetDebugDataPath", resp, "Failure sending request")
		return
	}

	result, err = client.GetDebugDataPathResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetDebugDataPath", resp, "Failure responding to request")
		return
	}

	return
}

// GetDebugDataPathPreparer prepares the GetDebugDataPath request.
func (client Client) GetDebugDataPathPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2016-03-20-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}/GetDebugDataPath", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDebugDataPathSender sends the GetDebugDataPath request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetDebugDataPathSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDebugDataPathResponder handles the response to the GetDebugDataPath request. The method always
// closes the http.Response Body.
func (client Client) GetDebugDataPathResponder(resp *http.Response) (result DataPath, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStatistics gets statistics of the specified job.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// jobIdentity - job Information ID.
func (client Client) GetStatistics(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result Statistics, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetStatistics")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStatisticsPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetStatistics", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatisticsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "GetStatistics", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatisticsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetStatistics", resp, "Failure responding to request")
		return
	}

	return
}

// GetStatisticsPreparer prepares the GetStatistics request.
func (client Client) GetStatisticsPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2016-03-20-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/Jobs/{jobIdentity}/GetStatistics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStatisticsSender sends the GetStatistics request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetStatisticsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetStatisticsResponder handles the response to the GetStatistics request. The method always
// closes the http.Response Body.
func (client Client) GetStatisticsResponder(resp *http.Response) (result Statistics, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the jobs, if any, associated with the specified Data Lake Analytics account. The response includes a link
// to the next page of results, if any.
// Parameters:
// accountName - the Azure Data Lake Analytics account to execute job operations on.
// filter - oData filter. Optional.
// top - the number of items to return. Optional.
// skip - the number of items to skip over before returning elements. Optional.
// expand - oData expansion. Expand related resources in line with the retrieved resources, e.g.
// Categories?$expand=Products would expand Product data in line with each Category entry. Optional.
// selectParameter - oData Select statement. Limits the properties on each entry to just those requested, e.g.
// Categories?$select=CategoryName,Description. Optional.
// orderby - orderBy clause. One or more comma-separated expressions with an optional "asc" (the default) or
// "desc" depending on the order you'd like the values sorted, e.g. Categories?$orderby=CategoryName desc.
// Optional.
// count - the Boolean value of true or false to request a count of the matching resources included with the
// resources in the response, e.g. Categories?$count=true. Optional.
// search - a free form search. A free-text search expression to match for whether a particular entry should be
// included in the feed, e.g. Categories?$search=blue OR green. Optional.
// formatParameter - the return format. Return the response in particular format without access to request
// headers for standard content-type negotiation (e.g Orders?$format=json). Optional.
func (client Client) List(ctx context.Context, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result InfoListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, accountName, filter, top, skip, expand, selectParameter, orderby, count, search, formatParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "List", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "List", resp, "Failure responding to request")
		return
	}
	if result.ilr.hasNextLink() && result.ilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2016-03-20-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}
	if len(search) > 0 {
		queryParameters["$search"] = autorest.Encode("query", search)
	}
	if len(formatParameter) > 0 {
		queryParameters["$format"] = autorest.Encode("query", formatParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/Jobs"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result InfoListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client Client) listNextResults(ctx context.Context, lastResults InfoListResult) (result InfoListResult, err error) {
	req, err := lastResults.infoListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "job.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "job.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result InfoListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, accountName, filter, top, skip, expand, selectParameter, orderby, count, search, formatParameter)
	return
}
