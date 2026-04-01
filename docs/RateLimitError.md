# RateLimitError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Limit** | **int32** | Request limit for this tier and strategy (24h sliding window) | 
**Remaining** | **int32** | Remaining requests in the current window | 
**Strategy** | **string** | Which strategy quota was exceeded | 

## Methods

### NewRateLimitError

`func NewRateLimitError(error_ string, limit int32, remaining int32, strategy string, ) *RateLimitError`

NewRateLimitError instantiates a new RateLimitError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitErrorWithDefaults

`func NewRateLimitErrorWithDefaults() *RateLimitError`

NewRateLimitErrorWithDefaults instantiates a new RateLimitError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RateLimitError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RateLimitError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RateLimitError) SetError(v string)`

SetError sets Error field to given value.


### GetLimit

`func (o *RateLimitError) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RateLimitError) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RateLimitError) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetRemaining

`func (o *RateLimitError) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *RateLimitError) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *RateLimitError) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.


### GetStrategy

`func (o *RateLimitError) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *RateLimitError) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *RateLimitError) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


