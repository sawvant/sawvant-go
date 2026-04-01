# OptimizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parts** | [**[]Part**](Part.md) |  | 
**Sheets** | [**[]Sheet**](Sheet.md) |  | 
**Machine** | [**Machine**](Machine.md) |  | 
**Strategy** | Pointer to **string** | Solve strategy. \&quot;fast\&quot; runs all greedy solvers concurrently. \&quot;thorough\&quot; adds Gilmore-Gomory column generation for optimal patterns. Each strategy has its own rate limit quota.  | [optional] [default to "fast"]
**CostTariffs** | Pointer to [**CostTariffs**](CostTariffs.md) |  | [optional] 

## Methods

### NewOptimizeRequest

`func NewOptimizeRequest(parts []Part, sheets []Sheet, machine Machine, ) *OptimizeRequest`

NewOptimizeRequest instantiates a new OptimizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizeRequestWithDefaults

`func NewOptimizeRequestWithDefaults() *OptimizeRequest`

NewOptimizeRequestWithDefaults instantiates a new OptimizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParts

`func (o *OptimizeRequest) GetParts() []Part`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *OptimizeRequest) GetPartsOk() (*[]Part, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *OptimizeRequest) SetParts(v []Part)`

SetParts sets Parts field to given value.


### GetSheets

`func (o *OptimizeRequest) GetSheets() []Sheet`

GetSheets returns the Sheets field if non-nil, zero value otherwise.

### GetSheetsOk

`func (o *OptimizeRequest) GetSheetsOk() (*[]Sheet, bool)`

GetSheetsOk returns a tuple with the Sheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheets

`func (o *OptimizeRequest) SetSheets(v []Sheet)`

SetSheets sets Sheets field to given value.


### GetMachine

`func (o *OptimizeRequest) GetMachine() Machine`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *OptimizeRequest) GetMachineOk() (*Machine, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *OptimizeRequest) SetMachine(v Machine)`

SetMachine sets Machine field to given value.


### GetStrategy

`func (o *OptimizeRequest) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *OptimizeRequest) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *OptimizeRequest) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *OptimizeRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetCostTariffs

`func (o *OptimizeRequest) GetCostTariffs() CostTariffs`

GetCostTariffs returns the CostTariffs field if non-nil, zero value otherwise.

### GetCostTariffsOk

`func (o *OptimizeRequest) GetCostTariffsOk() (*CostTariffs, bool)`

GetCostTariffsOk returns a tuple with the CostTariffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostTariffs

`func (o *OptimizeRequest) SetCostTariffs(v CostTariffs)`

SetCostTariffs sets CostTariffs field to given value.

### HasCostTariffs

`func (o *OptimizeRequest) HasCostTariffs() bool`

HasCostTariffs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


