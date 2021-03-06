package workloadmonitor

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// HealthMonitorsClient is the workload Monitor API
type HealthMonitorsClient struct {
	BaseClient
}

// NewHealthMonitorsClient creates an instance of the HealthMonitorsClient client.
func NewHealthMonitorsClient() HealthMonitorsClient {
	return NewHealthMonitorsClientWithBaseURI(DefaultBaseURI)
}

// NewHealthMonitorsClientWithBaseURI creates an instance of the HealthMonitorsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewHealthMonitorsClientWithBaseURI(baseURI string) HealthMonitorsClient {
	return HealthMonitorsClient{NewWithBaseURI(baseURI)}
}

// Get sends the get request.
// Parameters:
// subscriptionID - the subscription Id of the virtual machine.
// resourceGroupName - the resource group of the virtual machine.
// providerName - the provider name (ex: Microsoft.Compute for virtual machines).
// resourceCollectionName - the resource collection name (ex: virtualMachines for virtual machines).
// resourceName - the name of the virtual machine.
// monitorID - the monitor Id of the virtual machine.
// expand - optionally expand the monitor???s evidence and/or configuration. Example:
// $expand=evidence,configuration.
func (client HealthMonitorsClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, expand string) (result HealthMonitor, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HealthMonitorsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client HealthMonitorsClient) GetPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorId":              autorest.Encode("path", monitorID),
		"providerName":           autorest.Encode("path", providerName),
		"resourceCollectionName": autorest.Encode("path", resourceCollectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceName":           autorest.Encode("path", resourceName),
		"subscriptionId":         autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-01-13-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HealthMonitorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HealthMonitorsClient) GetResponder(resp *http.Response) (result HealthMonitor, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStateChange sends the get state change request.
// Parameters:
// subscriptionID - the subscription Id of the virtual machine.
// resourceGroupName - the resource group of the virtual machine.
// providerName - the provider name (ex: Microsoft.Compute for virtual machines).
// resourceCollectionName - the resource collection name (ex: virtualMachines for virtual machines).
// resourceName - the name of the virtual machine.
// monitorID - the monitor Id of the virtual machine.
// timestampUnix - the timestamp of the state change (unix format).
// expand - optionally expand the monitor???s evidence and/or configuration. Example:
// $expand=evidence,configuration.
func (client HealthMonitorsClient) GetStateChange(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, timestampUnix string, expand string) (result HealthMonitorStateChange, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HealthMonitorsClient.GetStateChange")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStateChangePreparer(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, timestampUnix, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "GetStateChange", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStateChangeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "GetStateChange", resp, "Failure sending request")
		return
	}

	result, err = client.GetStateChangeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "GetStateChange", resp, "Failure responding to request")
		return
	}

	return
}

// GetStateChangePreparer prepares the GetStateChange request.
func (client HealthMonitorsClient) GetStateChangePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, timestampUnix string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorId":              autorest.Encode("path", monitorID),
		"providerName":           autorest.Encode("path", providerName),
		"resourceCollectionName": autorest.Encode("path", resourceCollectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceName":           autorest.Encode("path", resourceName),
		"subscriptionId":         autorest.Encode("path", subscriptionID),
		"timestampUnix":          autorest.Encode("path", timestampUnix),
	}

	const APIVersion = "2020-01-13-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history/{timestampUnix}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStateChangeSender sends the GetStateChange request. The method will close the
// http.Response Body if it receives an error.
func (client HealthMonitorsClient) GetStateChangeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetStateChangeResponder handles the response to the GetStateChange request. The method always
// closes the http.Response Body.
func (client HealthMonitorsClient) GetStateChangeResponder(resp *http.Response) (result HealthMonitorStateChange, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// subscriptionID - the subscription Id of the virtual machine.
// resourceGroupName - the resource group of the virtual machine.
// providerName - the provider name (ex: Microsoft.Compute for virtual machines).
// resourceCollectionName - the resource collection name (ex: virtualMachines for virtual machines).
// resourceName - the name of the virtual machine.
// filter - optionally filter by monitor name. Example: $filter=monitorName eq
// 'logical-disks|C:|disk-free-space-mb.'
// expand - optionally expand the monitor???s evidence and/or configuration. Example:
// $expand=evidence,configuration.
func (client HealthMonitorsClient) List(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, filter string, expand string) (result HealthMonitorListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HealthMonitorsClient.List")
		defer func() {
			sc := -1
			if result.hml.Response.Response != nil {
				sc = result.hml.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.hml.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "List", resp, "Failure sending request")
		return
	}

	result.hml, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.hml.hasNextLink() && result.hml.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client HealthMonitorsClient) ListPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":           autorest.Encode("path", providerName),
		"resourceCollectionName": autorest.Encode("path", resourceCollectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceName":           autorest.Encode("path", resourceName),
		"subscriptionId":         autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-01-13-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client HealthMonitorsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client HealthMonitorsClient) ListResponder(resp *http.Response) (result HealthMonitorList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client HealthMonitorsClient) listNextResults(ctx context.Context, lastResults HealthMonitorList) (result HealthMonitorList, err error) {
	req, err := lastResults.healthMonitorListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client HealthMonitorsClient) ListComplete(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, filter string, expand string) (result HealthMonitorListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HealthMonitorsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, filter, expand)
	return
}

// ListStateChanges sends the list state changes request.
// Parameters:
// subscriptionID - the subscription Id of the virtual machine.
// resourceGroupName - the resource group of the virtual machine.
// providerName - the provider name (ex: Microsoft.Compute for virtual machines).
// resourceCollectionName - the resource collection name (ex: virtualMachines for virtual machines).
// resourceName - the name of the virtual machine.
// monitorID - the monitor Id of the virtual machine.
// filter - optionally filter by heartbeat condition. Example: $filter=isHeartbeat eq false.
// expand - optionally expand the monitor???s evidence and/or configuration. Example:
// $expand=evidence,configuration.
// startTimestampUtc - the start of the time window.
// endTimestampUtc - the end of the time window.
func (client HealthMonitorsClient) ListStateChanges(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, filter string, expand string, startTimestampUtc *date.Time, endTimestampUtc *date.Time) (result HealthMonitorStateChangeListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HealthMonitorsClient.ListStateChanges")
		defer func() {
			sc := -1
			if result.hmscl.Response.Response != nil {
				sc = result.hmscl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listStateChangesNextResults
	req, err := client.ListStateChangesPreparer(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, filter, expand, startTimestampUtc, endTimestampUtc)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "ListStateChanges", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListStateChangesSender(req)
	if err != nil {
		result.hmscl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "ListStateChanges", resp, "Failure sending request")
		return
	}

	result.hmscl, err = client.ListStateChangesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "ListStateChanges", resp, "Failure responding to request")
		return
	}
	if result.hmscl.hasNextLink() && result.hmscl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListStateChangesPreparer prepares the ListStateChanges request.
func (client HealthMonitorsClient) ListStateChangesPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, filter string, expand string, startTimestampUtc *date.Time, endTimestampUtc *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorId":              autorest.Encode("path", monitorID),
		"providerName":           autorest.Encode("path", providerName),
		"resourceCollectionName": autorest.Encode("path", resourceCollectionName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"resourceName":           autorest.Encode("path", resourceName),
		"subscriptionId":         autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-01-13-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if startTimestampUtc != nil {
		queryParameters["startTimestampUtc"] = autorest.Encode("query", *startTimestampUtc)
	}
	if endTimestampUtc != nil {
		queryParameters["endTimestampUtc"] = autorest.Encode("query", *endTimestampUtc)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListStateChangesSender sends the ListStateChanges request. The method will close the
// http.Response Body if it receives an error.
func (client HealthMonitorsClient) ListStateChangesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListStateChangesResponder handles the response to the ListStateChanges request. The method always
// closes the http.Response Body.
func (client HealthMonitorsClient) ListStateChangesResponder(resp *http.Response) (result HealthMonitorStateChangeList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listStateChangesNextResults retrieves the next set of results, if any.
func (client HealthMonitorsClient) listStateChangesNextResults(ctx context.Context, lastResults HealthMonitorStateChangeList) (result HealthMonitorStateChangeList, err error) {
	req, err := lastResults.healthMonitorStateChangeListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "listStateChangesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListStateChangesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "listStateChangesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListStateChangesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.HealthMonitorsClient", "listStateChangesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListStateChangesComplete enumerates all values, automatically crossing page boundaries as required.
func (client HealthMonitorsClient) ListStateChangesComplete(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, filter string, expand string, startTimestampUtc *date.Time, endTimestampUtc *date.Time) (result HealthMonitorStateChangeListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HealthMonitorsClient.ListStateChanges")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListStateChanges(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, filter, expand, startTimestampUtc, endTimestampUtc)
	return
}
