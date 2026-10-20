/*
 * Sawvant Cutting Optimization API — Go SDK
 *
 * File generated from our OpenAPI spec; DO NOT EDIT.
 */

package sawvant

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type DemoAPI interface {

	/*
	CreateDemoOptimization Run a small optimization without an API key

	Public endpoint behind the playground on sawvant.com. No API key or
account needed. Runs the fast strategy synchronously and returns the
result together with an SVG cutting plan per layout, so there is no
job to poll.

Limits: at most 30 part types, a total part quantity of 100, 3 sheet
types and a 64 KB request body. `strategy` must be omitted or `fast`;
`cost_tariffs` is ignored. Each client IP may make 10 requests per
minute and 100 per day. The input is not stored.

Browsers may only call this endpoint from sawvant.com.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateDemoOptimizationRequest
	*/
	CreateDemoOptimization(ctx context.Context) ApiCreateDemoOptimizationRequest

	// CreateDemoOptimizationExecute executes the request
	//  @return DemoOptimizeResponse
	CreateDemoOptimizationExecute(r ApiCreateDemoOptimizationRequest) (*DemoOptimizeResponse, *http.Response, error)
}

// DemoAPIService DemoAPI service
type DemoAPIService service

type ApiCreateDemoOptimizationRequest struct {
	ctx context.Context
	ApiService DemoAPI
	optimizeRequest *OptimizeRequest
}

func (r ApiCreateDemoOptimizationRequest) OptimizeRequest(optimizeRequest OptimizeRequest) ApiCreateDemoOptimizationRequest {
	r.optimizeRequest = &optimizeRequest
	return r
}

func (r ApiCreateDemoOptimizationRequest) Execute() (*DemoOptimizeResponse, *http.Response, error) {
	return r.ApiService.CreateDemoOptimizationExecute(r)
}

/*
CreateDemoOptimization Run a small optimization without an API key

Public endpoint behind the playground on sawvant.com. No API key or
account needed. Runs the fast strategy synchronously and returns the
result together with an SVG cutting plan per layout, so there is no
job to poll.

Limits: at most 30 part types, a total part quantity of 100, 3 sheet
types and a 64 KB request body. `strategy` must be omitted or `fast`;
`cost_tariffs` is ignored. Each client IP may make 10 requests per
minute and 100 per day. The input is not stored.

Browsers may only call this endpoint from sawvant.com.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDemoOptimizationRequest
*/
func (a *DemoAPIService) CreateDemoOptimization(ctx context.Context) ApiCreateDemoOptimizationRequest {
	return ApiCreateDemoOptimizationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DemoOptimizeResponse
func (a *DemoAPIService) CreateDemoOptimizationExecute(r ApiCreateDemoOptimizationRequest) (*DemoOptimizeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DemoOptimizeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DemoAPIService.CreateDemoOptimization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/demo/optimize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.optimizeRequest == nil {
		return localVarReturnValue, nil, reportError("optimizeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.optimizeRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v DemoError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v DemoError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v DemoError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v DemoError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v DemoError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v DemoError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
