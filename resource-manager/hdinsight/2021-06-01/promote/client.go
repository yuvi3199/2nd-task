package promote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PromoteClient struct {
	Client *resourcemanager.Client
}

func NewPromoteClientWithBaseURI(api environments.Api) (*PromoteClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "promote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PromoteClient: %+v", err)
	}

	return &PromoteClient{
		Client: client,
	}, nil
}
