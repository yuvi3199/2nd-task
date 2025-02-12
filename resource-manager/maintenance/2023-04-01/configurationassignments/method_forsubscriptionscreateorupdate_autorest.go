package configurationassignments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForSubscriptionsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationAssignment
}

// ForSubscriptionsCreateOrUpdate ...
func (c ConfigurationAssignmentsClient) ForSubscriptionsCreateOrUpdate(ctx context.Context, id ConfigurationAssignmentId, input ConfigurationAssignment) (result ForSubscriptionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForForSubscriptionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForSubscriptionsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForSubscriptionsCreateOrUpdate prepares the ForSubscriptionsCreateOrUpdate request.
func (c ConfigurationAssignmentsClient) preparerForForSubscriptionsCreateOrUpdate(ctx context.Context, id ConfigurationAssignmentId, input ConfigurationAssignment) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForForSubscriptionsCreateOrUpdate handles the response to the ForSubscriptionsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c ConfigurationAssignmentsClient) responderForForSubscriptionsCreateOrUpdate(resp *http.Response) (result ForSubscriptionsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
