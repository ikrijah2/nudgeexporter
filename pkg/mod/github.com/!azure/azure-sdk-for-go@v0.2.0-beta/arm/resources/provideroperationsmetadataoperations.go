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

// ProviderOperationsMetadataOperationsManagementClient is the client for the
// ProviderOperationsMetadataOperations methods of the Resources service.
type ProviderOperationsMetadataOperationsManagementClient struct {
	ManagementClient
}

// NewProviderOperationsMetadataOperationsManagementClient creates an instance
// of the ProviderOperationsMetadataOperationsManagementClient client.
func NewProviderOperationsMetadataOperationsManagementClient(subscriptionID string) ProviderOperationsMetadataOperationsManagementClient {
	return NewProviderOperationsMetadataOperationsManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderOperationsMetadataOperationsManagementClientWithBaseURI creates
// an instance of the ProviderOperationsMetadataOperationsManagementClient
// client.
func NewProviderOperationsMetadataOperationsManagementClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsMetadataOperationsManagementClient {
	return ProviderOperationsMetadataOperationsManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets provider operations metadata
//
// resourceProviderNamespace is namespace of the resource provider.
func (client ProviderOperationsMetadataOperationsManagementClient) Get(resourceProviderNamespace string, apiVersion string, expand string) (result ProviderOperationsMetadata, ae error) {
	req, err := client.GetPreparer(resourceProviderNamespace, apiVersion, expand)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationsMetadataOperationsManagementClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationsMetadataOperationsManagementClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/ProviderOperationsMetadataOperationsManagementClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProviderOperationsMetadataOperationsManagementClient) GetPreparer(resourceProviderNamespace string, apiVersion string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = expand
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/providerOperations/{resourceProviderNamespace}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationsMetadataOperationsManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProviderOperationsMetadataOperationsManagementClient) GetResponder(resp *http.Response) (result ProviderOperationsMetadata, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets provider operations metadata list
//
func (client ProviderOperationsMetadataOperationsManagementClient) List(apiVersion string, expand string) (result ProviderOperationsMetadataListResult, ae error) {
	req, err := client.ListPreparer(apiVersion, expand)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationsMetadataOperationsManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ProviderOperationsMetadataOperationsManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/ProviderOperationsMetadataOperationsManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProviderOperationsMetadataOperationsManagementClient) ListPreparer(apiVersion string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = expand
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/providerOperations"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationsMetadataOperationsManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProviderOperationsMetadataOperationsManagementClient) ListResponder(resp *http.Response) (result ProviderOperationsMetadataListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
