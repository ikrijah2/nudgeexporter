package web

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

// TopLevelDomainsClient is the webSite Management Client
type TopLevelDomainsClient struct {
	BaseClient
}

// NewTopLevelDomainsClient creates an instance of the TopLevelDomainsClient client.
func NewTopLevelDomainsClient(subscriptionID string) TopLevelDomainsClient {
	return NewTopLevelDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTopLevelDomainsClientWithBaseURI creates an instance of the TopLevelDomainsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTopLevelDomainsClientWithBaseURI(baseURI string, subscriptionID string) TopLevelDomainsClient {
	return TopLevelDomainsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get description for Get details of a top-level domain.
// Parameters:
// name - name of the top-level domain.
func (client TopLevelDomainsClient) Get(ctx context.Context, name string) (result TopLevelDomain, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TopLevelDomainsClient) GetPreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) GetResponder(resp *http.Response) (result TopLevelDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List description for Get all top-level domains supported for registration.
func (client TopLevelDomainsClient) List(ctx context.Context) (result TopLevelDomainCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.List")
		defer func() {
			sc := -1
			if result.tldc.Response.Response != nil {
				sc = result.tldc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tldc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "List", resp, "Failure sending request")
		return
	}

	result.tldc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.tldc.hasNextLink() && result.tldc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client TopLevelDomainsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) ListResponder(resp *http.Response) (result TopLevelDomainCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TopLevelDomainsClient) listNextResults(ctx context.Context, lastResults TopLevelDomainCollection) (result TopLevelDomainCollection, err error) {
	req, err := lastResults.topLevelDomainCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TopLevelDomainsClient) ListComplete(ctx context.Context) (result TopLevelDomainCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.List")
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

// ListAgreements description for Gets all legal agreements that user needs to accept before purchasing a domain.
// Parameters:
// name - name of the top-level domain.
// agreementOption - domain agreement options.
func (client TopLevelDomainsClient) ListAgreements(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption) (result TldLegalAgreementCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.ListAgreements")
		defer func() {
			sc := -1
			if result.tlac.Response.Response != nil {
				sc = result.tlac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAgreementsNextResults
	req, err := client.ListAgreementsPreparer(ctx, name, agreementOption)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "ListAgreements", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAgreementsSender(req)
	if err != nil {
		result.tlac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "ListAgreements", resp, "Failure sending request")
		return
	}

	result.tlac, err = client.ListAgreementsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "ListAgreements", resp, "Failure responding to request")
		return
	}
	if result.tlac.hasNextLink() && result.tlac.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAgreementsPreparer prepares the ListAgreements request.
func (client TopLevelDomainsClient) ListAgreementsPreparer(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}/listAgreements", pathParameters),
		autorest.WithJSON(agreementOption),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAgreementsSender sends the ListAgreements request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) ListAgreementsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAgreementsResponder handles the response to the ListAgreements request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) ListAgreementsResponder(resp *http.Response) (result TldLegalAgreementCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAgreementsNextResults retrieves the next set of results, if any.
func (client TopLevelDomainsClient) listAgreementsNextResults(ctx context.Context, lastResults TldLegalAgreementCollection) (result TldLegalAgreementCollection, err error) {
	req, err := lastResults.tldLegalAgreementCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listAgreementsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAgreementsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listAgreementsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAgreementsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.TopLevelDomainsClient", "listAgreementsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAgreementsComplete enumerates all values, automatically crossing page boundaries as required.
func (client TopLevelDomainsClient) ListAgreementsComplete(ctx context.Context, name string, agreementOption TopLevelDomainAgreementOption) (result TldLegalAgreementCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TopLevelDomainsClient.ListAgreements")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAgreements(ctx, name, agreementOption)
	return
}
