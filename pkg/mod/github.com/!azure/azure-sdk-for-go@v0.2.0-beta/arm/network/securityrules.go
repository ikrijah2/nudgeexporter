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

// SecurityRulesResourceProviderClient is the the Windows Azure Network
// management API provides a RESTful set of web services that interact with
// Windows Azure Networks service to manage your network resrources. The API
// has entities that capture the relationship between an end user and the
// Windows Azure Networks service.
type SecurityRulesResourceProviderClient struct {
	ResourceProviderClient
}

// NewSecurityRulesResourceProviderClient creates an instance of the
// SecurityRulesResourceProviderClient client.
func NewSecurityRulesResourceProviderClient(subscriptionID string) SecurityRulesResourceProviderClient {
	return NewSecurityRulesResourceProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecurityRulesResourceProviderClientWithBaseURI creates an instance of
// the SecurityRulesResourceProviderClient client.
func NewSecurityRulesResourceProviderClientWithBaseURI(baseURI string, subscriptionID string) SecurityRulesResourceProviderClient {
	return SecurityRulesResourceProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put network security rule operation creates/updates a
// security rule in the specified network security group
//
// resourceGroupName is the name of the resource group.
// networkSecurityGroupName is the name of the network security group.
// securityRuleName is the name of the security rule. securityRuleParameters
// is parameters supplied to the create/update network security rule
// operation
func (client SecurityRulesResourceProviderClient) CreateOrUpdate(resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule) (result SecurityRule, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityRulesResourceProviderClient) CreateOrUpdatePreparer(resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": url.QueryEscape(networkSecurityGroupName),
		"resourceGroupName":        url.QueryEscape(resourceGroupName),
		"securityRuleName":         url.QueryEscape(securityRuleName),
		"subscriptionId":           url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"),
		autorest.WithJSON(securityRuleParameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesResourceProviderClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityRulesResourceProviderClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the delete network security rule operation deletes the specified
// network security rule.
//
// resourceGroupName is the name of the resource group.
// networkSecurityGroupName is the name of the network security group.
// securityRuleName is the name of the security rule.
func (client SecurityRulesResourceProviderClient) Delete(resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, networkSecurityGroupName, securityRuleName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecurityRulesResourceProviderClient) DeletePreparer(resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": url.QueryEscape(networkSecurityGroupName),
		"resourceGroupName":        url.QueryEscape(resourceGroupName),
		"securityRuleName":         url.QueryEscape(securityRuleName),
		"subscriptionId":           url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesResourceProviderClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SecurityRulesResourceProviderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get NetworkSecurityRule operation retreives information about the
// specified network security rule.
//
// resourceGroupName is the name of the resource group.
// networkSecurityGroupName is the name of the network security group.
// securityRuleName is the name of the security rule.
func (client SecurityRulesResourceProviderClient) Get(resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result SecurityRule, ae error) {
	req, err := client.GetPreparer(resourceGroupName, networkSecurityGroupName, securityRuleName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecurityRulesResourceProviderClient) GetPreparer(resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": url.QueryEscape(networkSecurityGroupName),
		"resourceGroupName":        url.QueryEscape(resourceGroupName),
		"securityRuleName":         url.QueryEscape(securityRuleName),
		"subscriptionId":           url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesResourceProviderClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecurityRulesResourceProviderClient) GetResponder(resp *http.Response) (result SecurityRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List network security rule opertion retrieves all the security
// rules in a network security group.
//
// resourceGroupName is the name of the resource group.
// networkSecurityGroupName is the name of the network security group.
func (client SecurityRulesResourceProviderClient) List(resourceGroupName string, networkSecurityGroupName string) (result SecurityRuleListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName, networkSecurityGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SecurityRulesResourceProviderClient) ListPreparer(resourceGroupName string, networkSecurityGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": url.QueryEscape(networkSecurityGroupName),
		"resourceGroupName":        url.QueryEscape(resourceGroupName),
		"subscriptionId":           url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesResourceProviderClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecurityRulesResourceProviderClient) ListResponder(resp *http.Response) (result SecurityRuleListResult, err error) {
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
func (client SecurityRulesResourceProviderClient) ListNextResults(lastResults SecurityRuleListResult) (result SecurityRuleListResult, ae error) {
	req, err := lastResults.SecurityRuleListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/SecurityRulesResourceProviderClient", "List", "Failure responding to next results request request")
	}

	return
}
