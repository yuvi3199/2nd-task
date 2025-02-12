package sapdiskconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SAPDiskConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewSAPDiskConfigurationsClientWithBaseURI(api environments.Api) (*SAPDiskConfigurationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "sapdiskconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SAPDiskConfigurationsClient: %+v", err)
	}

	return &SAPDiskConfigurationsClient{
		Client: client,
	}, nil
}
