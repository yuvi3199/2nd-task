package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScalingTrigger interface {
}

func unmarshalScalingTriggerImplementation(input []byte) (ScalingTrigger, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScalingTrigger into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AveragePartitionLoadTrigger") {
		var out AveragePartitionLoadScalingTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AveragePartitionLoadScalingTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AverageServiceLoadTrigger") {
		var out AverageServiceLoadScalingTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AverageServiceLoadScalingTrigger: %+v", err)
		}
		return out, nil
	}

	type RawScalingTriggerImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawScalingTriggerImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}