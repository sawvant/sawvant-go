# ApiErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**ErrorCode**](ErrorCode.md) |  | 
**Message** | **string** | Human-readable explanation. Do not parse it; use &#x60;code&#x60; and &#x60;details&#x60;. | 
**Details** | Pointer to [**[]FieldError**](FieldError.md) |  | [optional] 

## Methods

### NewApiErrorBody

`func NewApiErrorBody(code ErrorCode, message string, ) *ApiErrorBody`

NewApiErrorBody instantiates a new ApiErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorBodyWithDefaults

`func NewApiErrorBodyWithDefaults() *ApiErrorBody`

NewApiErrorBodyWithDefaults instantiates a new ApiErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiErrorBody) GetCode() ErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiErrorBody) GetCodeOk() (*ErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiErrorBody) SetCode(v ErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ApiErrorBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiErrorBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiErrorBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *ApiErrorBody) GetDetails() []FieldError`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApiErrorBody) GetDetailsOk() (*[]FieldError, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ApiErrorBody) SetDetails(v []FieldError)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ApiErrorBody) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


