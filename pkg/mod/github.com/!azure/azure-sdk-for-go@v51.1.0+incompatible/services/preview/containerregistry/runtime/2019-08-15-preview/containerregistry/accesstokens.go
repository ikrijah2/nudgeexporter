package containerregistry

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

// AccessTokensClient is the metadata API definition for the Azure Container Registry runtime
type AccessTokensClient struct {
	BaseClient
}

// NewAccessTokensClient creates an instance of the AccessTokensClient client.
func NewAccessTokensClient(loginURI string) AccessTokensClient {
	return AccessTokensClient{New(loginURI)}
}

// Get exchange ACR Refresh token for an ACR Access Token
// Parameters:
// service - indicates the name of your Azure container registry.
// scope - which is expected to be a valid scope, and can be specified more than once for multiple scope
// requests. You obtained this from the Www-Authenticate response header from the challenge.
// refreshToken - must be a valid ACR refresh token
func (client AccessTokensClient) Get(ctx context.Context, service string, scope string, refreshToken string) (result AccessToken, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessTokensClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, service, scope, refreshToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AccessTokensClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.AccessTokensClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AccessTokensClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AccessTokensClient) GetPreparer(ctx context.Context, service string, scope string, refreshToken string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	formDataParameters := map[string]interface{}{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
		"scope":         scope,
		"service":       service,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPath("/oauth2/token"),
		autorest.WithFormData(autorest.MapToValues(formDataParameters)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AccessTokensClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AccessTokensClient) GetResponder(resp *http.Response) (result AccessToken, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetFromLogin exchange Username, Password and Scope an ACR Access Token
// Parameters:
// service - indicates the name of your Azure container registry.
// scope - expected to be a valid scope, and can be specified more than once for multiple scope requests. You
// can obtain this from the Www-Authenticate response header from the challenge.
func (client AccessTokensClient) GetFromLogin(ctx context.Context, service string, scope string) (result AccessToken, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccessTokensClient.GetFromLogin")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetFromLoginPreparer(ctx, service, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AccessTokensClient", "GetFromLogin", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetFromLoginSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.AccessTokensClient", "GetFromLogin", resp, "Failure sending request")
		return
	}

	result, err = client.GetFromLoginResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AccessTokensClient", "GetFromLogin", resp, "Failure responding to request")
		return
	}

	return
}

// GetFromLoginPreparer prepares the GetFromLogin request.
func (client AccessTokensClient) GetFromLoginPreparer(ctx context.Context, service string, scope string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	queryParameters := map[string]interface{}{
		"scope":   autorest.Encode("query", scope),
		"service": autorest.Encode("query", service),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPath("/oauth2/token"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetFromLoginSender sends the GetFromLogin request. The method will close the
// http.Response Body if it receives an error.
func (client AccessTokensClient) GetFromLoginSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetFromLoginResponder handles the response to the GetFromLogin request. The method always
// closes the http.Response Body.
func (client AccessTokensClient) GetFromLoginResponder(resp *http.Response) (result AccessToken, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
