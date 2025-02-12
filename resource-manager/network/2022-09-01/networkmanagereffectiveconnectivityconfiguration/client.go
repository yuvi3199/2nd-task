package networkmanagereffectiveconnectivityconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkManagerEffectiveConnectivityConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewNetworkManagerEffectiveConnectivityConfigurationClientWithBaseURI(api environments.Api) (*NetworkManagerEffectiveConnectivityConfigurationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "networkmanagereffectiveconnectivityconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkManagerEffectiveConnectivityConfigurationClient: %+v", err)
	}

	return &NetworkManagerEffectiveConnectivityConfigurationClient{
		Client: client,
	}, nil
}
