package post

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicesStopOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServicesStop ...
func (c POSTClient) ServicesStop(ctx context.Context, id ServiceId) (result ServicesStopOperationResponse, err error) {
	req, err := c.preparerForServicesStop(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "ServicesStop", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServicesStop(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "ServicesStop", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServicesStopThenPoll performs ServicesStop then polls until it's completed
func (c POSTClient) ServicesStopThenPoll(ctx context.Context, id ServiceId) error {
	result, err := c.ServicesStop(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServicesStop: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServicesStop: %+v", err)
	}

	return nil
}

// preparerForServicesStop prepares the ServicesStop request.
func (c POSTClient) preparerForServicesStop(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/stop", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServicesStop sends the ServicesStop request. The method will close the
// http.Response Body if it receives an error.
func (c POSTClient) senderForServicesStop(ctx context.Context, req *http.Request) (future ServicesStopOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
