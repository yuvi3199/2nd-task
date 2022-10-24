package communitygalleryimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImagePurchasePlan struct {
	Name      *string `json:"name,omitempty"`
	Product   *string `json:"product,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
}