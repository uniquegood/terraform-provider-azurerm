package vnetpeering

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VirtualNetworkPeering
}

type ListByWorkspaceCompleteResult struct {
	Items []VirtualNetworkPeering
}

// ListByWorkspace ...
func (c VNetPeeringClient) ListByWorkspace(ctx context.Context, id WorkspaceId) (result ListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/virtualNetworkPeerings", id.ID()),
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
		Values *[]VirtualNetworkPeering `json:"values"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByWorkspaceComplete retrieves all the results into a single object
func (c VNetPeeringClient) ListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (ListByWorkspaceCompleteResult, error) {
	return c.ListByWorkspaceCompleteMatchingPredicate(ctx, id, VirtualNetworkPeeringOperationPredicate{})
}

// ListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VNetPeeringClient) ListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate VirtualNetworkPeeringOperationPredicate) (result ListByWorkspaceCompleteResult, err error) {
	items := make([]VirtualNetworkPeering, 0)

	resp, err := c.ListByWorkspace(ctx, id)
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

	result = ListByWorkspaceCompleteResult{
		Items: items,
	}
	return
}
