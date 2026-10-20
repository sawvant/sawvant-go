# \DemoAPI

All URIs are relative to *https://api.sawvant.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDemoOptimization**](DemoAPI.md#CreateDemoOptimization) | **Post** /v1/demo/optimize | Run a small optimization without an API key



## CreateDemoOptimization

> DemoOptimizeResponse CreateDemoOptimization(ctx).OptimizeRequest(optimizeRequest).Execute()

Run a small optimization without an API key



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
	resp, r, err := apiClient.DemoAPI.CreateDemoOptimization(context.Background()).OptimizeRequest(optimizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DemoAPI.CreateDemoOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDemoOptimization`: DemoOptimizeResponse
	fmt.Fprintf(os.Stdout, "Response from `DemoAPI.CreateDemoOptimization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDemoOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optimizeRequest** | [**OptimizeRequest**](OptimizeRequest.md) |  | 

### Return type

[**DemoOptimizeResponse**](DemoOptimizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

