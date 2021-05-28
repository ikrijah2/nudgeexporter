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

// TagsManagementClient is the client for the Tags methods of the Resources
// service.
type TagsManagementClient struct {
	ManagementClient
}

// NewTagsManagementClient creates an instance of the TagsManagementClient
// client.
func NewTagsManagementClient(subscriptionID string) TagsManagementClient {
	return NewTagsManagementClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTagsManagementClientWithBaseURI creates an instance of the
// TagsManagementClient client.
func NewTagsManagementClientWithBaseURI(baseURI string, subscriptionID string) TagsManagementClient {
	return TagsManagementClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a subscription resource tag.
//
// tagName is the name of the tag.
func (client TagsManagementClient) CreateOrUpdate(tagName string) (result TagDetails, ae error) {
	req, err := client.CreateOrUpdatePreparer(tagName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/TagsManagementClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client TagsManagementClient) CreateOrUpdatePreparer(tagName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
		"tagName":        url.QueryEscape(tagName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client TagsManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusCreated)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client TagsManagementClient) CreateOrUpdateResponder(resp *http.Response) (result TagDetails, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateValue create a subscription resource tag value.
//
// tagName is the name of the tag. tagValue is the value of the tag.
func (client TagsManagementClient) CreateOrUpdateValue(tagName string, tagValue string) (result TagValue, ae error) {
	req, err := client.CreateOrUpdateValuePreparer(tagName, tagValue)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "CreateOrUpdateValue", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateValueSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "CreateOrUpdateValue", "Failure sending request")
	}

	result, err = client.CreateOrUpdateValueResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/TagsManagementClient", "CreateOrUpdateValue", "Failure responding to request")
	}

	return
}

// CreateOrUpdateValuePreparer prepares the CreateOrUpdateValue request.
func (client TagsManagementClient) CreateOrUpdateValuePreparer(tagName string, tagValue string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
		"tagName":        url.QueryEscape(tagName),
		"tagValue":       url.QueryEscape(tagValue),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateValueSender sends the CreateOrUpdateValue request. The method will close the
// http.Response Body if it receives an error.
func (client TagsManagementClient) CreateOrUpdateValueSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusCreated)
}

// CreateOrUpdateValueResponder handles the response to the CreateOrUpdateValue request. The method always
// closes the http.Response Body.
func (client TagsManagementClient) CreateOrUpdateValueResponder(resp *http.Response) (result TagValue, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a subscription resource tag.
//
// tagName is the name of the tag.
func (client TagsManagementClient) Delete(tagName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(tagName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/TagsManagementClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TagsManagementClient) DeletePreparer(tagName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
		"tagName":        url.QueryEscape(tagName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TagsManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TagsManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteValue delete a subscription resource tag value.
//
// tagName is the name of the tag. tagValue is the value of the tag.
func (client TagsManagementClient) DeleteValue(tagName string, tagValue string) (result autorest.Response, ae error) {
	req, err := client.DeleteValuePreparer(tagName, tagValue)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "DeleteValue", "Failure preparing request")
	}

	resp, err := client.DeleteValueSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "DeleteValue", "Failure sending request")
	}

	result, err = client.DeleteValueResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/TagsManagementClient", "DeleteValue", "Failure responding to request")
	}

	return
}

// DeleteValuePreparer prepares the DeleteValue request.
func (client TagsManagementClient) DeleteValuePreparer(tagName string, tagValue string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
		"tagName":        url.QueryEscape(tagName),
		"tagValue":       url.QueryEscape(tagValue),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteValueSender sends the DeleteValue request. The method will close the
// http.Response Body if it receives an error.
func (client TagsManagementClient) DeleteValueSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK, http.StatusNoContent)
}

// DeleteValueResponder handles the response to the DeleteValue request. The method always
// closes the http.Response Body.
func (client TagsManagementClient) DeleteValueResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List get a list of subscription resource tags.
func (client TagsManagementClient) List() (result TagsListResult, ae error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/TagsManagementClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client TagsManagementClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TagsManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TagsManagementClient) ListResponder(resp *http.Response) (result TagsListResult, err error) {
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
func (client TagsManagementClient) ListNextResults(lastResults TagsListResult) (result TagsListResult, ae error) {
	req, err := lastResults.TagsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/TagsManagementClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "resources/TagsManagementClient", "List", "Failure responding to next results request request")
	}

	return
}