# BatchJobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Position of this job in the request array | 
**JobId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewBatchJobResult

`func NewBatchJobResult(index int32, ) *BatchJobResult`

NewBatchJobResult instantiates a new BatchJobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobResultWithDefaults

`func NewBatchJobResultWithDefaults() *BatchJobResult`

NewBatchJobResultWithDefaults instantiates a new BatchJobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *BatchJobResult) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *BatchJobResult) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *BatchJobResult) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetJobId

`func (o *BatchJobResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BatchJobResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BatchJobResult) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BatchJobResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetError

`func (o *BatchJobResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchJobResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchJobResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BatchJobResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


