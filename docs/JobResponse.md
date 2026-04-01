# JobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** |  | 
**Status** | **string** |  | 
**Progress** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Result** | Pointer to [**OptimizeResult**](OptimizeResult.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | E.g. unplaced part IDs | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewJobResponse

`func NewJobResponse(jobId string, status string, progress int32, createdAt time.Time, ) *JobResponse`

NewJobResponse instantiates a new JobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResponseWithDefaults

`func NewJobResponseWithDefaults() *JobResponse`

NewJobResponseWithDefaults instantiates a new JobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *JobResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *JobResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *JobResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *JobResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetCreatedAt

`func (o *JobResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *JobResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *JobResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *JobResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *JobResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *JobResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *JobResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *JobResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *JobResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetResult

`func (o *JobResponse) GetResult() OptimizeResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JobResponse) GetResultOk() (*OptimizeResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JobResponse) SetResult(v OptimizeResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *JobResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetWarnings

`func (o *JobResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *JobResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *JobResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *JobResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetError

`func (o *JobResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JobResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JobResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *JobResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


