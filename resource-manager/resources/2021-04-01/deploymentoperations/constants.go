package deploymentoperations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningOperation string

const (
	ProvisioningOperationAction                     ProvisioningOperation = "Action"
	ProvisioningOperationAzureAsyncOperationWaiting ProvisioningOperation = "AzureAsyncOperationWaiting"
	ProvisioningOperationCreate                     ProvisioningOperation = "Create"
	ProvisioningOperationDelete                     ProvisioningOperation = "Delete"
	ProvisioningOperationDeploymentCleanup          ProvisioningOperation = "DeploymentCleanup"
	ProvisioningOperationEvaluateDeploymentOutput   ProvisioningOperation = "EvaluateDeploymentOutput"
	ProvisioningOperationNotSpecified               ProvisioningOperation = "NotSpecified"
	ProvisioningOperationRead                       ProvisioningOperation = "Read"
	ProvisioningOperationResourceCacheWaiting       ProvisioningOperation = "ResourceCacheWaiting"
	ProvisioningOperationWaiting                    ProvisioningOperation = "Waiting"
)

func PossibleValuesForProvisioningOperation() []string {
	return []string{
		string(ProvisioningOperationAction),
		string(ProvisioningOperationAzureAsyncOperationWaiting),
		string(ProvisioningOperationCreate),
		string(ProvisioningOperationDelete),
		string(ProvisioningOperationDeploymentCleanup),
		string(ProvisioningOperationEvaluateDeploymentOutput),
		string(ProvisioningOperationNotSpecified),
		string(ProvisioningOperationRead),
		string(ProvisioningOperationResourceCacheWaiting),
		string(ProvisioningOperationWaiting),
	}
}

func parseProvisioningOperation(input string) (*ProvisioningOperation, error) {
	vals := map[string]ProvisioningOperation{
		"action":                     ProvisioningOperationAction,
		"azureasyncoperationwaiting": ProvisioningOperationAzureAsyncOperationWaiting,
		"create":                     ProvisioningOperationCreate,
		"delete":                     ProvisioningOperationDelete,
		"deploymentcleanup":          ProvisioningOperationDeploymentCleanup,
		"evaluatedeploymentoutput":   ProvisioningOperationEvaluateDeploymentOutput,
		"notspecified":               ProvisioningOperationNotSpecified,
		"read":                       ProvisioningOperationRead,
		"resourcecachewaiting":       ProvisioningOperationResourceCacheWaiting,
		"waiting":                    ProvisioningOperationWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningOperation(input)
	return &out, nil
}
