package findrestorabletimeranges

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FindRestorableTimeRangesClient struct {
	Client *resourcemanager.Client
}

func NewFindRestorableTimeRangesClientWithBaseURI(api environments.Api) (*FindRestorableTimeRangesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "findrestorabletimeranges", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FindRestorableTimeRangesClient: %+v", err)
	}

	return &FindRestorableTimeRangesClient{
		Client: client,
	}, nil
}
