package resources

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

// ProvidersManagementClient is the client for the Providers methods of the
// Resources service.
type ProvidersManagementClient struct {
	ManagementClient
}

// NewProvidersManagementClient creates an instance of the
// ProvidersManagementClient client.
func NewProvidersManagementClient(subscriptionID string) ProvidersManagementClient {
	return NewProvidersManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProvidersManagementClientWithBaseURI creates an instance of the
// ProvidersManagementClient client.
func NewProvidersManagementClientWithBaseURI(baseURI string, subscriptionID string) ProvidersManagementClient {
	return ProvidersManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a resource provider.
//
// resourceProviderNamespace is namespace of the resource provider.
func (client ProvidersManagementClient) Get(resourceProviderNamespace string) (result Provider, ae error) {
	req, err := client.GetPreparer(resourceProviderNamespace)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProvidersManagementClient) GetPreparer(resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProvidersManagementClient) GetResponder(resp *http.Response) (result Provider, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of resource providers.
//
// top is query parameters. If null is passed returns all deployments.
func (client ProvidersManagementClient) List(top *int) (result ProviderListResult, ae error) {
	req, err := client.ListPreparer(top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProvidersManagementClient) ListPreparer(top *int) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProvidersManagementClient) ListResponder(resp *http.Response) (result ProviderListResult, err error) {
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
func (client ProvidersManagementClient) ListNextResults(lastResults ProviderListResult) (result ProviderListResult, ae error) {
	req, err := lastResults.ProviderListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "List", "Failure responding to next results request request")
	}

	return
}

// Register registers provider to be used with a subscription.
//
// resourceProviderNamespace is namespace of the resource provider.
func (client ProvidersManagementClient) Register(resourceProviderNamespace string) (result Provider, ae error) {
	req, err := client.RegisterPreparer(resourceProviderNamespace)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Register", "Failure preparing request")
	}

	resp, err := client.RegisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Register", "Failure sending request")
	}

	result, err = client.RegisterResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Register", "Failure responding to request")
	}

	return
}

// RegisterPreparer prepares the Register request.
func (client ProvidersManagementClient) RegisterPreparer(resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/register"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// RegisterSender sends the Register request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersManagementClient) RegisterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// RegisterResponder handles the response to the Register request. The method always
// closes the http.Response Body.
func (client ProvidersManagementClient) RegisterResponder(resp *http.Response) (result Provider, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Unregister unregisters provider from a subscription.
//
// resourceProviderNamespace is namespace of the resource provider.
func (client ProvidersManagementClient) Unregister(resourceProviderNamespace string) (result Provider, ae error) {
	req, err := client.UnregisterPreparer(resourceProviderNamespace)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Unregister", "Failure preparing request")
	}

	resp, err := client.UnregisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Unregister", "Failure sending request")
	}

	result, err = client.UnregisterResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/ProvidersManagementClient", "Unregister", "Failure responding to request")
	}

	return
}

// UnregisterPreparer prepares the Unregister request.
func (client ProvidersManagementClient) UnregisterPreparer(resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/unregister"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// UnregisterSender sends the Unregister request. The method will close the
// http.Response Body if it receives an error.
func (client ProvidersManagementClient) UnregisterSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// UnregisterResponder handles the response to the Unregister request. The method always
// closes the http.Response Body.
func (client ProvidersManagementClient) UnregisterResponder(resp *http.Response) (result Provider, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
