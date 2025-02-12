package expressroutecrossconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCrossConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCrossConnectionsClientWithBaseURI(api environments.Api) (*ExpressRouteCrossConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutecrossconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCrossConnectionsClient: %+v", err)
	}

	return &ExpressRouteCrossConnectionsClient{
		Client: client,
	}, nil
}
