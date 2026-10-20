# BatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]OptimizeRequest**](OptimizeRequest.md) | Individual optimization requests. At most 50 with strategy &#x60;fast&#x60; or 10 with strategy &#x60;thorough&#x60;.  | 

## Methods

### NewBatchRequest

`func NewBatchRequest(jobs []OptimizeRequest, ) *BatchRequest`

NewBatchRequest instantiates a new BatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestWithDefaults

`func NewBatchRequestWithDefaults() *BatchRequest`

NewBatchRequestWithDefaults instantiates a new BatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *BatchRequest) GetJobs() []OptimizeRequest`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *BatchRequest) GetJobsOk() (*[]OptimizeRequest, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *BatchRequest) SetJobs(v []OptimizeRequest)`

SetJobs sets Jobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


