# BatchAccepted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **string** |  | 
**JobIds** | **[]string** |  | 
**Total** | **int32** | Number of jobs in the request | 
**Created** | **int32** | Number of jobs successfully enqueued | 
**Results** | Pointer to [**[]BatchJobResult**](BatchJobResult.md) | Present only when one or more jobs failed to enqueue | [optional] 

## Methods

### NewBatchAccepted

`func NewBatchAccepted(batchId string, jobIds []string, total int32, created int32, ) *BatchAccepted`

NewBatchAccepted instantiates a new BatchAccepted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchAcceptedWithDefaults

`func NewBatchAcceptedWithDefaults() *BatchAccepted`

NewBatchAcceptedWithDefaults instantiates a new BatchAccepted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *BatchAccepted) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *BatchAccepted) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *BatchAccepted) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### GetJobIds

`func (o *BatchAccepted) GetJobIds() []string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *BatchAccepted) GetJobIdsOk() (*[]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *BatchAccepted) SetJobIds(v []string)`

SetJobIds sets JobIds field to given value.


### GetTotal

`func (o *BatchAccepted) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BatchAccepted) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BatchAccepted) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetCreated

`func (o *BatchAccepted) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BatchAccepted) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BatchAccepted) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetResults

`func (o *BatchAccepted) GetResults() []BatchJobResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchAccepted) GetResultsOk() (*[]BatchJobResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchAccepted) SetResults(v []BatchJobResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BatchAccepted) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


