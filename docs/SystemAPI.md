# \SystemAPI

All URIs are relative to *https://api.sawvant.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealth**](SystemAPI.md#GetHealth) | **Get** /health | Health check
[**ListTiers**](SystemAPI.md#ListTiers) | **Get** /v1/tiers | List available subscription tiers



## GetHealth

> HealthResponse GetHealth(ctx).Execute()

Health check

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.GetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealth`: HealthResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


### Return type

[**HealthResponse**](HealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTiers

> []TierDefinition ListTiers(ctx).Execute()

List available subscription tiers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemAPI.ListTiers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTiers`: []TierDefinition
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTiersRequest struct via the builder pattern


### Return type

[**[]TierDefinition**](TierDefinition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

