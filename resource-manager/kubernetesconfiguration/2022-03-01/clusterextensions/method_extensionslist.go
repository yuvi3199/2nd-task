package clusterextensions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Extension
}

type ExtensionsListCompleteResult struct {
	Items []Extension
}

// ExtensionsList ...
func (c ClusterExtensionsClient) ExtensionsList(ctx context.Context, id ProviderId) (result ExtensionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.KubernetesConfiguration/extensions", id.ID()),
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
		Values *[]Extension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ExtensionsListComplete retrieves all the results into a single object
func (c ClusterExtensionsClient) ExtensionsListComplete(ctx context.Context, id ProviderId) (ExtensionsListCompleteResult, error) {
	return c.ExtensionsListCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ExtensionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ClusterExtensionsClient) ExtensionsListCompleteMatchingPredicate(ctx context.Context, id ProviderId, predicate ExtensionOperationPredicate) (result ExtensionsListCompleteResult, err error) {
	items := make([]Extension, 0)

	resp, err := c.ExtensionsList(ctx, id)
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

	result = ExtensionsListCompleteResult{
		Items: items,
	}
	return
}
