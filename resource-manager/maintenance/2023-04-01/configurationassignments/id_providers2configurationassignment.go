package configurationassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2ConfigurationAssignmentId{}

// Providers2ConfigurationAssignmentId is a struct representing the Resource ID for a Providers 2 Configuration Assignment
type Providers2ConfigurationAssignmentId struct {
	SubscriptionId              string
	ResourceGroupName           string
	ProviderName                string
	ResourceType                string
	ResourceName                string
	ConfigurationAssignmentName string
}

// NewProviders2ConfigurationAssignmentID returns a new Providers2ConfigurationAssignmentId struct
func NewProviders2ConfigurationAssignmentID(subscriptionId string, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string) Providers2ConfigurationAssignmentId {
	return Providers2ConfigurationAssignmentId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		ProviderName:                providerName,
		ResourceType:                resourceType,
		ResourceName:                resourceName,
		ConfigurationAssignmentName: configurationAssignmentName,
	}
}

// ParseProviders2ConfigurationAssignmentID parses 'input' into a Providers2ConfigurationAssignmentId
func ParseProviders2ConfigurationAssignmentID(input string) (*Providers2ConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2ConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceType", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.ConfigurationAssignmentName, ok = parsed.Parsed["configurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ParseProviders2ConfigurationAssignmentIDInsensitively parses 'input' case-insensitively into a Providers2ConfigurationAssignmentId
// note: this method should only be used for API response data and not user input
func ParseProviders2ConfigurationAssignmentIDInsensitively(input string) (*Providers2ConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2ConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2ConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceType", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.ConfigurationAssignmentName, ok = parsed.Parsed["configurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ValidateProviders2ConfigurationAssignmentID checks that 'input' can be parsed as a Providers 2 Configuration Assignment ID
func ValidateProviders2ConfigurationAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2ConfigurationAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Configuration Assignment ID
func (id Providers2ConfigurationAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.Maintenance/configurationAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ResourceType, id.ResourceName, id.ConfigurationAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Configuration Assignment ID
func (id Providers2ConfigurationAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("resourceType", "resourceTypeValue"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMaintenance", "Microsoft.Maintenance", "Microsoft.Maintenance"),
		resourceids.StaticSegment("staticConfigurationAssignments", "configurationAssignments", "configurationAssignments"),
		resourceids.UserSpecifiedSegment("configurationAssignmentName", "configurationAssignmentValue"),
	}
}

// String returns a human-readable description of this Providers 2 Configuration Assignment ID
func (id Providers2ConfigurationAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Resource Type: %q", id.ResourceType),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Configuration Assignment Name: %q", id.ConfigurationAssignmentName),
	}
	return fmt.Sprintf("Providers 2 Configuration Assignment (%s)", strings.Join(components, "\n"))
}
