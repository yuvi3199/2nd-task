package iotsecuritysolutions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotSecuritySolutionDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// IotSecuritySolutionDelete ...
func (c IotSecuritySolutionsClient) IotSecuritySolutionDelete(ctx context.Context, id IotSecuritySolutionId) (result IotSecuritySolutionDeleteOperationResponse, err error) {
	req, err := c.preparerForIotSecuritySolutionDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotSecuritySolutionDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotSecuritySolutionDelete prepares the IotSecuritySolutionDelete request.
func (c IotSecuritySolutionsClient) preparerForIotSecuritySolutionDelete(ctx context.Context, id IotSecuritySolutionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIotSecuritySolutionDelete handles the response to the IotSecuritySolutionDelete request. The method always
// closes the http.Response Body.
func (c IotSecuritySolutionsClient) responderForIotSecuritySolutionDelete(resp *http.Response) (result IotSecuritySolutionDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
