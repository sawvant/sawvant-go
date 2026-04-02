# \JobsAPI

All URIs are relative to *https://api.sawvant.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJob**](JobsAPI.md#GetJob) | **Get** /v1/jobs/{id} | Get job status
[**StreamJob**](JobsAPI.md#StreamJob) | **Get** /v1/jobs/{id}/stream | Stream job progress via SSE



## GetJob

> JobResponse GetJob(ctx, id).Execute()

Get job status



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: JobResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobResponse**](JobResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamJob

> string StreamJob(ctx, id).Execute()

Stream job progress via SSE



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.StreamJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.StreamJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamJob`: string
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.StreamJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

