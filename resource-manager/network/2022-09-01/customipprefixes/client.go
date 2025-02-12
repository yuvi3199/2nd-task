package customipprefixes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomIPPrefixesClient struct {
	Client *resourcemanager.Client
}

func NewCustomIPPrefixesClientWithBaseURI(api environments.Api) (*CustomIPPrefixesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "customipprefixes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomIPPrefixesClient: %+v", err)
	}

	return &CustomIPPrefixesClient{
		Client: client,
	}, nil
}
