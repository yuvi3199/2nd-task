package events

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EventSummary
}

type ListByBillingProfileCompleteResult struct {
	Items []EventSummary
}

type ListByBillingProfileOperationOptions struct {
	EndDate   *string
	StartDate *string
}

func DefaultListByBillingProfileOperationOptions() ListByBillingProfileOperationOptions {
	return ListByBillingProfileOperationOptions{}
}

func (o ListByBillingProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByBillingProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDate != nil {
		out.Append("endDate", fmt.Sprintf("%v", *o.EndDate))
	}
	if o.StartDate != nil {
		out.Append("startDate", fmt.Sprintf("%v", *o.StartDate))
	}
	return &out
}

// ListByBillingProfile ...
func (c EventsClient) ListByBillingProfile(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (result ListByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Consumption/events", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]EventSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingProfileComplete retrieves all the results into a single object
func (c EventsClient) ListByBillingProfileComplete(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (ListByBillingProfileCompleteResult, error) {
	return c.ListByBillingProfileCompleteMatchingPredicate(ctx, id, options, EventSummaryOperationPredicate{})
}

// ListByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventsClient) ListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions, predicate EventSummaryOperationPredicate) (result ListByBillingProfileCompleteResult, err error) {
	items := make([]EventSummary, 0)

	resp, err := c.ListByBillingProfile(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByBillingProfileCompleteResult{
		Items: items,
	}
	return
}
