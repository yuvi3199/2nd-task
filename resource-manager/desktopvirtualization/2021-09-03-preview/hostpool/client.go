package hostpool

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostPoolClient struct {
	Client *resourcemanager.Client
}

func NewHostPoolClientWithBaseURI(api environments.Api) (*HostPoolClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "hostpool", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HostPoolClient: %+v", err)
	}

	return &HostPoolClient{
		Client: client,
	}, nil
}
