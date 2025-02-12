package dsccompilationjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobStreamListResult struct {
	NextLink *string      `json:"nextLink,omitempty"`
	Value    *[]JobStream `json:"value,omitempty"`
}
