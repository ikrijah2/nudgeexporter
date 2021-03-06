package network

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

// UsagesResourceProviderClient is the the Windows Azure Network management
// API provides a RESTful set of web services that interact with Windows
// Azure Networks service to manage your network resrources. The API has
// entities that capture the relationship between an end user and the Windows
// Azure Networks service.
type UsagesResourceProviderClient struct {
	ResourceProviderClient
}

// NewUsagesResourceProviderClient creates an instance of the
// UsagesResourceProviderClient client.
func NewUsagesResourceProviderClient(subscriptionID string) UsagesResourceProviderClient {
	return NewUsagesResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsagesResourceProviderClientWithBaseURI creates an instance of the
// UsagesResourceProviderClient client.
func NewUsagesResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) UsagesResourceProviderClient {
	return UsagesResourceProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists compute usages for a subscription.
//
// location is the location upon which resource usage is queried.
func (client UsagesResourceProviderClient) List(location string) (result UsagesListResult, ae error) {
	req, err := client.ListPreparer(location)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/UsagesResourceProviderClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/UsagesResourceProviderClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/UsagesResourceProviderClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client UsagesResourceProviderClient) ListPreparer(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsagesResourceProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsagesResourceProviderClient) ListResponder(resp *http.Response) (result UsagesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
