
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/scalingplan` Documentation

The `scalingplan` SDK allows for interaction with the Azure Resource Manager Service `desktopvirtualization` (API Version `2022-09-09`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/scalingplan"
```


### Client Initialization

```go
client := scalingplan.NewScalingPlanClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScalingPlanClient.Create`

```go
ctx := context.TODO()
id := scalingplan.NewScalingPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanValue")

payload := scalingplan.ScalingPlan{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanClient.Delete`

```go
ctx := context.TODO()
id := scalingplan.NewScalingPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanClient.Get`

```go
ctx := context.TODO()
id := scalingplan.NewScalingPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanClient.ListByHostPool`

```go
ctx := context.TODO()
id := scalingplan.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue")

// alternatively `client.ListByHostPool(ctx, id, scalingplan.DefaultListByHostPoolOperationOptions())` can be used to do batched pagination
items, err := client.ListByHostPoolComplete(ctx, id, scalingplan.DefaultListByHostPoolOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScalingPlanClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := scalingplan.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, scalingplan.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, scalingplan.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScalingPlanClient.ListBySubscription`

```go
ctx := context.TODO()
id := scalingplan.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, scalingplan.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, scalingplan.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScalingPlanClient.Update`

```go
ctx := context.TODO()
id := scalingplan.NewScalingPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanValue")

payload := scalingplan.ScalingPlanPatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
