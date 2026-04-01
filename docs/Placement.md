# Placement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartId** | **string** |  | 
**SheetId** | **string** | Which sheet type this part is placed on | 
**X** | **float64** | mm from top-left of usable area | 
**Y** | **float64** | mm from top-left of usable area | 
**Width** | **float64** | Post-edge-banding dimension | 
**Height** | **float64** | Post-edge-banding dimension | 
**Rotated** | **bool** |  | 
**Grain** | [**GrainDirection**](GrainDirection.md) |  | 

## Methods

### NewPlacement

`func NewPlacement(partId string, sheetId string, x float64, y float64, width float64, height float64, rotated bool, grain GrainDirection, ) *Placement`

NewPlacement instantiates a new Placement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementWithDefaults

`func NewPlacementWithDefaults() *Placement`

NewPlacementWithDefaults instantiates a new Placement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartId

`func (o *Placement) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *Placement) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *Placement) SetPartId(v string)`

SetPartId sets PartId field to given value.


### GetSheetId

`func (o *Placement) GetSheetId() string`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *Placement) GetSheetIdOk() (*string, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *Placement) SetSheetId(v string)`

SetSheetId sets SheetId field to given value.


### GetX

`func (o *Placement) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Placement) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Placement) SetX(v float64)`

SetX sets X field to given value.


### GetY

`func (o *Placement) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Placement) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Placement) SetY(v float64)`

SetY sets Y field to given value.


### GetWidth

`func (o *Placement) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Placement) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Placement) SetWidth(v float64)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *Placement) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Placement) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Placement) SetHeight(v float64)`

SetHeight sets Height field to given value.


### GetRotated

`func (o *Placement) GetRotated() bool`

GetRotated returns the Rotated field if non-nil, zero value otherwise.

### GetRotatedOk

`func (o *Placement) GetRotatedOk() (*bool, bool)`

GetRotatedOk returns a tuple with the Rotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotated

`func (o *Placement) SetRotated(v bool)`

SetRotated sets Rotated field to given value.


### GetGrain

`func (o *Placement) GetGrain() GrainDirection`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *Placement) GetGrainOk() (*GrainDirection, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *Placement) SetGrain(v GrainDirection)`

SetGrain sets Grain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


