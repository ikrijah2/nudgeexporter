package resourcehealth

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

// EmergingIssuesClient is the the Resource Health Client.
type EmergingIssuesClient struct {
	BaseClient
}

// NewEmergingIssuesClient creates an instance of the EmergingIssuesClient client.
func NewEmergingIssuesClient(subscriptionID string) EmergingIssuesClient {
	return NewEmergingIssuesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEmergingIssuesClientWithBaseURI creates an instance of the EmergingIssuesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEmergingIssuesClientWithBaseURI(baseURI string, subscriptionID string) EmergingIssuesClient {
	return EmergingIssuesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets Azure services' emerging issues.
func (client EmergingIssuesClient) Get(ctx context.Context) (result EmergingIssuesGetResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmergingIssuesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client EmergingIssuesClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"issueName": autorest.Encode("path", "default"),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ResourceHealth/emergingIssues/{issueName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EmergingIssuesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EmergingIssuesClient) GetResponder(resp *http.Response) (result EmergingIssuesGetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists Azure services' emerging issues.
func (client EmergingIssuesClient) List(ctx context.Context) (result EmergingIssueListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmergingIssuesClient.List")
		defer func() {
			sc := -1
			if result.eilr.Response.Response != nil {
				sc = result.eilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.eilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "List", resp, "Failure sending request")
		return
	}

	result.eilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.eilr.hasNextLink() && result.eilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EmergingIssuesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.ResourceHealth/emergingIssues"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EmergingIssuesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EmergingIssuesClient) ListResponder(resp *http.Response) (result EmergingIssueListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EmergingIssuesClient) listNextResults(ctx context.Context, lastResults EmergingIssueListResult) (result EmergingIssueListResult, err error) {
	req, err := lastResults.emergingIssueListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.EmergingIssuesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EmergingIssuesClient) ListComplete(ctx context.Context) (result EmergingIssueListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EmergingIssuesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
