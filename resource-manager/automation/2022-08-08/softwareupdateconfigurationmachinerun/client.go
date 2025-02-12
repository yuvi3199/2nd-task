package softwareupdateconfigurationmachinerun

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationMachineRunClient struct {
	Client *resourcemanager.Client
}

func NewSoftwareUpdateConfigurationMachineRunClientWithBaseURI(api environments.Api) (*SoftwareUpdateConfigurationMachineRunClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "softwareupdateconfigurationmachinerun", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SoftwareUpdateConfigurationMachineRunClient: %+v", err)
	}

	return &SoftwareUpdateConfigurationMachineRunClient{
		Client: client,
	}, nil
}
