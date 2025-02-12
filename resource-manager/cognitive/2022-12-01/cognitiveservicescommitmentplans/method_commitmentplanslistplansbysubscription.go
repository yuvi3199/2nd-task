package cognitiveservicescommitmentplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentPlansListPlansBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CommitmentPlan
}

type CommitmentPlansListPlansBySubscriptionCompleteResult struct {
	Items []CommitmentPlan
}

// CommitmentPlansListPlansBySubscription ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansBySubscription(ctx context.Context, id commonids.SubscriptionId) (result CommitmentPlansListPlansBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.CognitiveServices/commitmentPlans", id.ID()),
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
		Values *[]CommitmentPlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CommitmentPlansListPlansBySubscriptionComplete retrieves all the results into a single object
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (CommitmentPlansListPlansBySubscriptionCompleteResult, error) {
	return c.CommitmentPlansListPlansBySubscriptionCompleteMatchingPredicate(ctx, id, CommitmentPlanOperationPredicate{})
}

// CommitmentPlansListPlansBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate CommitmentPlanOperationPredicate) (result CommitmentPlansListPlansBySubscriptionCompleteResult, err error) {
	items := make([]CommitmentPlan, 0)

	resp, err := c.CommitmentPlansListPlansBySubscription(ctx, id)
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

	result = CommitmentPlansListPlansBySubscriptionCompleteResult{
		Items: items,
	}
	return
}
