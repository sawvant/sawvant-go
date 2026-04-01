# Layout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SheetId** | **string** |  | 
**Quantity** | **int32** | Number of identical sheets using this layout pattern | 
**Placements** | [**[]Placement**](Placement.md) |  | 

## Methods

### NewLayout

`func NewLayout(sheetId string, quantity int32, placements []Placement, ) *Layout`

NewLayout instantiates a new Layout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutWithDefaults

`func NewLayoutWithDefaults() *Layout`

NewLayoutWithDefaults instantiates a new Layout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSheetId

`func (o *Layout) GetSheetId() string`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *Layout) GetSheetIdOk() (*string, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *Layout) SetSheetId(v string)`

SetSheetId sets SheetId field to given value.


### GetQuantity

`func (o *Layout) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Layout) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Layout) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetPlacements

`func (o *Layout) GetPlacements() []Placement`

GetPlacements returns the Placements field if non-nil, zero value otherwise.

### GetPlacementsOk

`func (o *Layout) GetPlacementsOk() (*[]Placement, bool)`

GetPlacementsOk returns a tuple with the Placements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacements

`func (o *Layout) SetPlacements(v []Placement)`

SetPlacements sets Placements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


