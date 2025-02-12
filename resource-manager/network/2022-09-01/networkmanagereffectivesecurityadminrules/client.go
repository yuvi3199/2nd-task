package networkmanagereffectivesecurityadminrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkManagerEffectiveSecurityAdminRulesClient struct {
	Client *resourcemanager.Client
}

func NewNetworkManagerEffectiveSecurityAdminRulesClientWithBaseURI(api environments.Api) (*NetworkManagerEffectiveSecurityAdminRulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "networkmanagereffectivesecurityadminrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkManagerEffectiveSecurityAdminRulesClient: %+v", err)
	}

	return &NetworkManagerEffectiveSecurityAdminRulesClient{
		Client: client,
	}, nil
}
