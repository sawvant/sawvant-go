# OptimizeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layouts** | [**[]Layout**](Layout.md) |  | 
**Summary** | [**Summary**](Summary.md) |  | 
**Metrics** | Pointer to [**Metrics**](Metrics.md) |  | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 

## Methods

### NewOptimizeResult

`func NewOptimizeResult(layouts []Layout, summary Summary, ) *OptimizeResult`

NewOptimizeResult instantiates a new OptimizeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizeResultWithDefaults

`func NewOptimizeResultWithDefaults() *OptimizeResult`

NewOptimizeResultWithDefaults instantiates a new OptimizeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayouts

`func (o *OptimizeResult) GetLayouts() []Layout`

GetLayouts returns the Layouts field if non-nil, zero value otherwise.

### GetLayoutsOk

`func (o *OptimizeResult) GetLayoutsOk() (*[]Layout, bool)`

GetLayoutsOk returns a tuple with the Layouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayouts

`func (o *OptimizeResult) SetLayouts(v []Layout)`

SetLayouts sets Layouts field to given value.


### GetSummary

`func (o *OptimizeResult) GetSummary() Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OptimizeResult) GetSummaryOk() (*Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OptimizeResult) SetSummary(v Summary)`

SetSummary sets Summary field to given value.


### GetMetrics

`func (o *OptimizeResult) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *OptimizeResult) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *OptimizeResult) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *OptimizeResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetCost

`func (o *OptimizeResult) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OptimizeResult) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OptimizeResult) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *OptimizeResult) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


