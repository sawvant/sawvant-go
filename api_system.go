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


type SystemAPI interface {

	/*
	GetHealth Health check

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetHealthRequest
	*/
	GetHealth(ctx context.Context) ApiGetHealthRequest

	// GetHealthExecute executes the request
	//  @return HealthResponse
	GetHealthExecute(r ApiGetHealthRequest) (*HealthResponse, *http.Response, error)

	/*
	ListTiers List available subscription tiers

	Returns all available subscription tiers with their rate limits, pricing,
and feature gates. This is the single source of truth for tier configuration.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListTiersRequest
	*/
	ListTiers(ctx context.Context) ApiListTiersRequest

	// ListTiersExecute executes the request
	//  @return []TierDefinition
	ListTiersExecute(r ApiListTiersRequest) ([]TierDefinition, *http.Response, error)
}

// SystemAPIService SystemAPI service
type SystemAPIService service

type ApiGetHealthRequest struct {
	ctx context.Context
	ApiService SystemAPI
}

func (r ApiGetHealthRequest) Execute() (*HealthResponse, *http.Response, error) {
	return r.ApiService.GetHealthExecute(r)
}

/*
GetHealth Health check

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHealthRequest
*/
func (a *SystemAPIService) GetHealth(ctx context.Context) ApiGetHealthRequest {
	return ApiGetHealthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HealthResponse
func (a *SystemAPIService) GetHealthExecute(r ApiGetHealthRequest) (*HealthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HealthResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemAPIService.GetHealth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/health"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiListTiersRequest struct {
	ctx context.Context
	ApiService SystemAPI
}

func (r ApiListTiersRequest) Execute() ([]TierDefinition, *http.Response, error) {
	return r.ApiService.ListTiersExecute(r)
}

/*
ListTiers List available subscription tiers

Returns all available subscription tiers with their rate limits, pricing,
and feature gates. This is the single source of truth for tier configuration.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListTiersRequest
*/
func (a *SystemAPIService) ListTiers(ctx context.Context) ApiListTiersRequest {
	return ApiListTiersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TierDefinition
func (a *SystemAPIService) ListTiersExecute(r ApiListTiersRequest) ([]TierDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TierDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemAPIService.ListTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
