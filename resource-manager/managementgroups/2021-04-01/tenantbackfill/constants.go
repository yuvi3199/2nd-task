package tenantbackfill

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Status string

const (
	StatusCancelled                Status = "Cancelled"
	StatusCompleted                Status = "Completed"
	StatusFailed                   Status = "Failed"
	StatusNotStarted               Status = "NotStarted"
	StatusNotStartedButGroupsExist Status = "NotStartedButGroupsExist"
	StatusStarted                  Status = "Started"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusCancelled),
		string(StatusCompleted),
		string(StatusFailed),
		string(StatusNotStarted),
		string(StatusNotStartedButGroupsExist),
		string(StatusStarted),
	}
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"cancelled":                StatusCancelled,
		"completed":                StatusCompleted,
		"failed":                   StatusFailed,
		"notstarted":               StatusNotStarted,
		"notstartedbutgroupsexist": StatusNotStartedButGroupsExist,
		"started":                  StatusStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
