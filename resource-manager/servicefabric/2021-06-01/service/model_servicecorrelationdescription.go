package service

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceCorrelationDescription struct {
	Scheme      ServiceCorrelationScheme `json:"scheme"`
	ServiceName string                   `json:"serviceName"`
}
