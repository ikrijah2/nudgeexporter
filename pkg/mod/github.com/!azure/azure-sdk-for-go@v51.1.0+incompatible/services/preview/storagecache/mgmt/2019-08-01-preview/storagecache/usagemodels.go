package storagecache

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

// UsageModelsClient is the a Storage Cache provides scalable caching service for NAS clients, serving data from either
// NFSv3 or Blob at-rest storage (referred to as "Storage Targets"). These operations allow you to manage caches.
type UsageModelsClient struct {
	BaseClient
}

// NewUsageModelsClient creates an instance of the UsageModelsClient client.
func NewUsageModelsClient(subscriptionID string) UsageModelsClient {
	return NewUsageModelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageModelsClientWithBaseURI creates an instance of the UsageModelsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsageModelsClientWithBaseURI(baseURI string, subscriptionID string) UsageModelsClient {
	return UsageModelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List get the list of cache Usage Models available to this subscription.
func (client UsageModelsClient) List(ctx context.Context) (result UsageModelsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageModelsClient.List")
		defer func() {
			sc := -1
			if result.umr.Response.Response != nil {
				sc = result.umr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.UsageModelsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.umr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.UsageModelsClient", "List", resp, "Failure sending request")
		return
	}

	result.umr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.UsageModelsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.umr.hasNextLink() && result.umr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client UsageModelsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/usageModels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsageModelsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsageModelsClient) ListResponder(resp *http.Response) (result UsageModelsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client UsageModelsClient) listNextResults(ctx context.Context, lastResults UsageModelsResult) (result UsageModelsResult, err error) {
	req, err := lastResults.usageModelsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storagecache.UsageModelsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storagecache.UsageModelsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.UsageModelsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsageModelsClient) ListComplete(ctx context.Context) (result UsageModelsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsageModelsClient.List")
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
