# DemoLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SheetId** | **string** |  | 
**Quantity** | **int32** | Number of identical sheets using this layout pattern | 
**Placements** | [**[]Placement**](Placement.md) |  | 
**Svg** | **string** | Cutting plan of this layout as a standalone SVG document | 

## Methods

### NewDemoLayout

`func NewDemoLayout(sheetId string, quantity int32, placements []Placement, svg string, ) *DemoLayout`

NewDemoLayout instantiates a new DemoLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDemoLayoutWithDefaults

`func NewDemoLayoutWithDefaults() *DemoLayout`

NewDemoLayoutWithDefaults instantiates a new DemoLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSheetId

`func (o *DemoLayout) GetSheetId() string`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *DemoLayout) GetSheetIdOk() (*string, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *DemoLayout) SetSheetId(v string)`

SetSheetId sets SheetId field to given value.


### GetQuantity

`func (o *DemoLayout) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DemoLayout) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DemoLayout) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetPlacements

`func (o *DemoLayout) GetPlacements() []Placement`

GetPlacements returns the Placements field if non-nil, zero value otherwise.

### GetPlacementsOk

`func (o *DemoLayout) GetPlacementsOk() (*[]Placement, bool)`

GetPlacementsOk returns a tuple with the Placements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacements

`func (o *DemoLayout) SetPlacements(v []Placement)`

SetPlacements sets Placements field to given value.


### GetSvg

`func (o *DemoLayout) GetSvg() string`

GetSvg returns the Svg field if non-nil, zero value otherwise.

### GetSvgOk

`func (o *DemoLayout) GetSvgOk() (*string, bool)`

GetSvgOk returns a tuple with the Svg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvg

`func (o *DemoLayout) SetSvg(v string)`

SetSvg sets Svg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


