// Package frontdoor implements the Azure ARM Frontdoor service API version .
//
// FrontDoor Client
package frontdoor

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
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Frontdoor
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Frontdoor.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckFrontDoorNameAvailability check the availability of a Front Door resource name.
// Parameters:
// checkFrontDoorNameAvailabilityInput - input to check.
func (client BaseClient) CheckFrontDoorNameAvailability(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckFrontDoorNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: checkFrontDoorNameAvailabilityInput,
			Constraints: []validation.Constraint{{Target: "checkFrontDoorNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.BaseClient", "CheckFrontDoorNameAvailability", err.Error())
	}

	req, err := client.CheckFrontDoorNameAvailabilityPreparer(ctx, checkFrontDoorNameAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BaseClient", "CheckFrontDoorNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckFrontDoorNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.BaseClient", "CheckFrontDoorNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckFrontDoorNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BaseClient", "CheckFrontDoorNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckFrontDoorNameAvailabilityPreparer prepares the CheckFrontDoorNameAvailability request.
func (client BaseClient) CheckFrontDoorNameAvailabilityPreparer(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput) (*http.Request, error) {
	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Network/checkFrontDoorNameAvailability"),
		autorest.WithJSON(checkFrontDoorNameAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckFrontDoorNameAvailabilitySender sends the CheckFrontDoorNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckFrontDoorNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CheckFrontDoorNameAvailabilityResponder handles the response to the CheckFrontDoorNameAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckFrontDoorNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CheckFrontDoorNameAvailabilityWithSubscription check the availability of a Front Door subdomain.
// Parameters:
// checkFrontDoorNameAvailabilityInput - input to check.
func (client BaseClient) CheckFrontDoorNameAvailabilityWithSubscription(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckFrontDoorNameAvailabilityWithSubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: checkFrontDoorNameAvailabilityInput,
			Constraints: []validation.Constraint{{Target: "checkFrontDoorNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("frontdoor.BaseClient", "CheckFrontDoorNameAvailabilityWithSubscription", err.Error())
	}

	req, err := client.CheckFrontDoorNameAvailabilityWithSubscriptionPreparer(ctx, checkFrontDoorNameAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BaseClient", "CheckFrontDoorNameAvailabilityWithSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckFrontDoorNameAvailabilityWithSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "frontdoor.BaseClient", "CheckFrontDoorNameAvailabilityWithSubscription", resp, "Failure sending request")
		return
	}

	result, err = client.CheckFrontDoorNameAvailabilityWithSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "frontdoor.BaseClient", "CheckFrontDoorNameAvailabilityWithSubscription", resp, "Failure responding to request")
		return
	}

	return
}

// CheckFrontDoorNameAvailabilityWithSubscriptionPreparer prepares the CheckFrontDoorNameAvailabilityWithSubscription request.
func (client BaseClient) CheckFrontDoorNameAvailabilityWithSubscriptionPreparer(ctx context.Context, checkFrontDoorNameAvailabilityInput CheckNameAvailabilityInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/checkFrontDoorNameAvailability", pathParameters),
		autorest.WithJSON(checkFrontDoorNameAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckFrontDoorNameAvailabilityWithSubscriptionSender sends the CheckFrontDoorNameAvailabilityWithSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckFrontDoorNameAvailabilityWithSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckFrontDoorNameAvailabilityWithSubscriptionResponder handles the response to the CheckFrontDoorNameAvailabilityWithSubscription request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckFrontDoorNameAvailabilityWithSubscriptionResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
