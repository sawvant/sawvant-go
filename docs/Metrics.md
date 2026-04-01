# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCutLength** | **float64** | Total cut length in mm | 
**TotalCuts** | **int32** |  | 
**TotalRotations** | **int32** |  | 
**TotalStacks** | **int32** |  | 
**CutCycles** | **int32** |  | 

## Methods

### NewMetrics

`func NewMetrics(totalCutLength float64, totalCuts int32, totalRotations int32, totalStacks int32, cutCycles int32, ) *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCutLength

`func (o *Metrics) GetTotalCutLength() float64`

GetTotalCutLength returns the TotalCutLength field if non-nil, zero value otherwise.

### GetTotalCutLengthOk

`func (o *Metrics) GetTotalCutLengthOk() (*float64, bool)`

GetTotalCutLengthOk returns a tuple with the TotalCutLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCutLength

`func (o *Metrics) SetTotalCutLength(v float64)`

SetTotalCutLength sets TotalCutLength field to given value.


### GetTotalCuts

`func (o *Metrics) GetTotalCuts() int32`

GetTotalCuts returns the TotalCuts field if non-nil, zero value otherwise.

### GetTotalCutsOk

`func (o *Metrics) GetTotalCutsOk() (*int32, bool)`

GetTotalCutsOk returns a tuple with the TotalCuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCuts

`func (o *Metrics) SetTotalCuts(v int32)`

SetTotalCuts sets TotalCuts field to given value.


### GetTotalRotations

`func (o *Metrics) GetTotalRotations() int32`

GetTotalRotations returns the TotalRotations field if non-nil, zero value otherwise.

### GetTotalRotationsOk

`func (o *Metrics) GetTotalRotationsOk() (*int32, bool)`

GetTotalRotationsOk returns a tuple with the TotalRotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRotations

`func (o *Metrics) SetTotalRotations(v int32)`

SetTotalRotations sets TotalRotations field to given value.


### GetTotalStacks

`func (o *Metrics) GetTotalStacks() int32`

GetTotalStacks returns the TotalStacks field if non-nil, zero value otherwise.

### GetTotalStacksOk

`func (o *Metrics) GetTotalStacksOk() (*int32, bool)`

GetTotalStacksOk returns a tuple with the TotalStacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStacks

`func (o *Metrics) SetTotalStacks(v int32)`

SetTotalStacks sets TotalStacks field to given value.


### GetCutCycles

`func (o *Metrics) GetCutCycles() int32`

GetCutCycles returns the CutCycles field if non-nil, zero value otherwise.

### GetCutCyclesOk

`func (o *Metrics) GetCutCyclesOk() (*int32, bool)`

GetCutCyclesOk returns a tuple with the CutCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutCycles

`func (o *Metrics) SetCutCycles(v int32)`

SetCutCycles sets CutCycles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


