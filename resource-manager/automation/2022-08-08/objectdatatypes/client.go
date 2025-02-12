package objectdatatypes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectDataTypesClient struct {
	Client *resourcemanager.Client
}

func NewObjectDataTypesClientWithBaseURI(api environments.Api) (*ObjectDataTypesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "objectdatatypes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ObjectDataTypesClient: %+v", err)
	}

	return &ObjectDataTypesClient{
		Client: client,
	}, nil
}
