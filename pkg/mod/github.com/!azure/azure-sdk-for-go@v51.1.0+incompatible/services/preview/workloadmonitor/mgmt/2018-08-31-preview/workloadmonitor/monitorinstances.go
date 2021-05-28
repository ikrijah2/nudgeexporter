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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// MonitorInstancesClient is the aPIs for workload monitoring
type MonitorInstancesClient struct {
	BaseClient
}

// NewMonitorInstancesClient creates an instance of the MonitorInstancesClient client.
func NewMonitorInstancesClient(subscriptionID string) MonitorInstancesClient {
	return NewMonitorInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMonitorInstancesClientWithBaseURI creates an instance of the MonitorInstancesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewMonitorInstancesClientWithBaseURI(baseURI string, subscriptionID string) MonitorInstancesClient {
	return MonitorInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
// monitorInstanceID - monitorInstance Id.
// selectParameter - properties to be returned in the response.
// expand - include properties inline in the response.
func (client MonitorInstancesClient) Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorInstanceID uuid.UUID, selectParameter string, expand string) (result MonitorInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorInstancesClient.Get")
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
		return result, validation.NewError("workloadmonitor.MonitorInstancesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, monitorInstanceID, selectParameter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client MonitorInstancesClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, monitorInstanceID uuid.UUID, selectParameter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorInstanceId": autorest.Encode("path", monitorInstanceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceNamespace": autorest.Encode("path", resourceNamespace),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitorInstances/{monitorInstanceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MonitorInstancesClient) GetResponder(resp *http.Response) (result MonitorInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResource sends the list by resource request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceNamespace - the Namespace of the resource.
// resourceType - the type of the resource.
// resourceName - name of the resource.
// selectParameter - properties to be returned in the response.
// filter - filter to be applied on the operation.
// apply - apply aggregation.
// orderby - sort the result on one or more properties.
// expand - include properties inline in the response.
// top - limit the result to the specified number of rows.
// skiptoken - the page-continuation token to use with a paged version of this API.
func (client MonitorInstancesClient) ListByResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result MonitorInstancesCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorInstancesClient.ListByResource")
		defer func() {
			sc := -1
			if result.mic.Response.Response != nil {
				sc = result.mic.Response.Response.StatusCode
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
		return result, validation.NewError("workloadmonitor.MonitorInstancesClient", "ListByResource", err.Error())
	}

	result.fn = client.listByResourceNextResults
	req, err := client.ListByResourcePreparer(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, selectParameter, filter, apply, orderby, expand, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "ListByResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.mic.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "ListByResource", resp, "Failure sending request")
		return
	}

	result.mic, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "ListByResource", resp, "Failure responding to request")
		return
	}
	if result.mic.hasNextLink() && result.mic.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourcePreparer prepares the ListByResource request.
func (client MonitorInstancesClient) ListByResourcePreparer(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceNamespace": autorest.Encode("path", resourceNamespace),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(top) > 0 {
		queryParameters["$top"] = autorest.Encode("query", top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitorInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceSender sends the ListByResource request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorInstancesClient) ListByResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceResponder handles the response to the ListByResource request. The method always
// closes the http.Response Body.
func (client MonitorInstancesClient) ListByResourceResponder(resp *http.Response) (result MonitorInstancesCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceNextResults retrieves the next set of results, if any.
func (client MonitorInstancesClient) listByResourceNextResults(ctx context.Context, lastResults MonitorInstancesCollection) (result MonitorInstancesCollection, err error) {
	req, err := lastResults.monitorInstancesCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "listByResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "listByResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.MonitorInstancesClient", "listByResourceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client MonitorInstancesClient) ListByResourceComplete(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result MonitorInstancesCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorInstancesClient.ListByResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResource(ctx, resourceGroupName, resourceNamespace, resourceType, resourceName, selectParameter, filter, apply, orderby, expand, top, skiptoken)
	return
}