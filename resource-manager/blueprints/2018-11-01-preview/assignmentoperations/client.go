package assignmentoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentOperationsClient struct {
	Client *resourcemanager.Client
}

func NewAssignmentOperationsClientWithBaseURI(api environments.Api) (*AssignmentOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "assignmentoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssignmentOperationsClient: %+v", err)
	}

	return &AssignmentOperationsClient{
		Client: client,
	}, nil
}
