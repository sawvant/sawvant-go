# \OptimizationAPI

All URIs are relative to *https://api.sawvant.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchOptimization**](OptimizationAPI.md#CreateBatchOptimization) | **Post** /v1/optimize/batch | Submit multiple cutting optimization jobs
[**CreateOptimization**](OptimizationAPI.md#CreateOptimization) | **Post** /v1/optimize | Submit a cutting optimization job



## CreateBatchOptimization

> BatchAccepted CreateBatchOptimization(ctx).BatchRequest(batchRequest).Execute()

Submit multiple cutting optimization jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sawvant/sawvant-go"
)

func main() {
	batchRequest := *openapiclient.NewBatchRequest([]openapiclient.OptimizeRequest{*openapiclient.NewOptimizeRequest([]openapiclient.Part{*openapiclient.NewPart("Id_example", float64(123), float64(123), int32(123), openapiclient.GrainDirection("none"))}, []openapiclient.Sheet{*openapiclient.NewSheet("Id_example", float64(123), float64(123), int32(123), openapiclient.GrainDirection("none"))}, *openapiclient.NewMachine(float64(123), int32(123), openapiclient.CutDirection("default")))}) // BatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptimizationAPI.CreateBatchOptimization(context.Background()).BatchRequest(batchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptimizationAPI.CreateBatchOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchOptimization`: BatchAccepted
	fmt.Fprintf(os.Stdout, "Response from `OptimizationAPI.CreateBatchOptimization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRequest** | [**BatchRequest**](BatchRequest.md) |  | 

### Return type

[**BatchAccepted**](BatchAccepted.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/sawvant/sawvant-go"
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

