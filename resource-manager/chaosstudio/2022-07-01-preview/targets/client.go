package targets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetsClient struct {
	Client *resourcemanager.Client
}

func NewTargetsClientWithBaseURI(api environments.Api) (*TargetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "targets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TargetsClient: %+v", err)
	}

	return &TargetsClient{
		Client: client,
	}, nil
}
