# Machine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BladeThickness** | **float64** | Kerf width in mm | 
**MaxLevels** | **int32** | Cut pattern complexity (1-3) | 
**MaxStackHeight** | Pointer to **float64** | Maximum stack height in mm for batch cutting | [optional] 
**CutDirection** | [**CutDirection**](CutDirection.md) |  | 

## Methods

### NewMachine

`func NewMachine(bladeThickness float64, maxLevels int32, cutDirection CutDirection, ) *Machine`

NewMachine instantiates a new Machine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineWithDefaults

`func NewMachineWithDefaults() *Machine`

NewMachineWithDefaults instantiates a new Machine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBladeThickness

`func (o *Machine) GetBladeThickness() float64`

GetBladeThickness returns the BladeThickness field if non-nil, zero value otherwise.

### GetBladeThicknessOk

`func (o *Machine) GetBladeThicknessOk() (*float64, bool)`

GetBladeThicknessOk returns a tuple with the BladeThickness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBladeThickness

`func (o *Machine) SetBladeThickness(v float64)`

SetBladeThickness sets BladeThickness field to given value.


### GetMaxLevels

`func (o *Machine) GetMaxLevels() int32`

GetMaxLevels returns the MaxLevels field if non-nil, zero value otherwise.

### GetMaxLevelsOk

`func (o *Machine) GetMaxLevelsOk() (*int32, bool)`

GetMaxLevelsOk returns a tuple with the MaxLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLevels

`func (o *Machine) SetMaxLevels(v int32)`

SetMaxLevels sets MaxLevels field to given value.


### GetMaxStackHeight

`func (o *Machine) GetMaxStackHeight() float64`

GetMaxStackHeight returns the MaxStackHeight field if non-nil, zero value otherwise.

### GetMaxStackHeightOk

`func (o *Machine) GetMaxStackHeightOk() (*float64, bool)`

GetMaxStackHeightOk returns a tuple with the MaxStackHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStackHeight

`func (o *Machine) SetMaxStackHeight(v float64)`

SetMaxStackHeight sets MaxStackHeight field to given value.

### HasMaxStackHeight

`func (o *Machine) HasMaxStackHeight() bool`

HasMaxStackHeight returns a boolean if a field has been set.

### GetCutDirection

`func (o *Machine) GetCutDirection() CutDirection`

GetCutDirection returns the CutDirection field if non-nil, zero value otherwise.

### GetCutDirectionOk

`func (o *Machine) GetCutDirectionOk() (*CutDirection, bool)`

GetCutDirectionOk returns a tuple with the CutDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutDirection

`func (o *Machine) SetCutDirection(v CutDirection)`

SetCutDirection sets CutDirection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


