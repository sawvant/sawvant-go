# \OptimizationAPI

All URIs are relative to *https://api.sawvant.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOptimization**](OptimizationAPI.md#CreateOptimization) | **Post** /v1/optimize | Submit a cutting optimization job



## CreateOptimization

> OptimizeAccepted CreateOptimization(ctx).OptimizeRequest(optimizeRequest).Execute()

Submit a cutting optimization job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	optimizeRequest := *openapiclient.NewOptimizeRequest([]openapiclient.Part{*openapiclient.NewPart("Id_example", float64(123), float64(123), int32(123), openapiclient.GrainDirection("none"))}, []openapiclient.Sheet{*openapiclient.NewSheet("Id_example", float64(123), float64(123), int32(123), openapiclient.GrainDirection("none"))}, *openapiclient.NewMachine(float64(123), int32(123), openapiclient.CutDirection("default"))) // OptimizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptimizationAPI.CreateOptimization(context.Background()).OptimizeRequest(optimizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.CreateOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOptimization`: OptimizeAccepted
	fmt.Fprintf(os.Stdout, "Response from `OptimizationAPI.CreateOptimization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optimizeRequest** | [**OptimizeRequest**](OptimizeRequest.md) |  | 

### Return type

[**OptimizeAccepted**](OptimizeAccepted.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

