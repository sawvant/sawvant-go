# Part

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Length** | **float64** | Length in mm | 
**Width** | **float64** | Width in mm | 
**Quantity** | **int32** |  | 
**Grain** | [**GrainDirection**](GrainDirection.md) |  | 
**EdgeBanding** | Pointer to [**EdgeCorrection**](EdgeCorrection.md) |  | [optional] 

## Methods

### NewPart

`func NewPart(id string, length float64, width float64, quantity int32, grain GrainDirection, ) *Part`

NewPart instantiates a new Part object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartWithDefaults

`func NewPartWithDefaults() *Part`

NewPartWithDefaults instantiates a new Part object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Part) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Part) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Part) SetId(v string)`

SetId sets Id field to given value.


### GetLength

`func (o *Part) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Part) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Part) SetLength(v float64)`

SetLength sets Length field to given value.


### GetWidth

`func (o *Part) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Part) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Part) SetWidth(v float64)`

SetWidth sets Width field to given value.


### GetQuantity

`func (o *Part) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Part) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Part) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetGrain

`func (o *Part) GetGrain() GrainDirection`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *Part) GetGrainOk() (*GrainDirection, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *Part) SetGrain(v GrainDirection)`

SetGrain sets Grain field to given value.


### GetEdgeBanding

`func (o *Part) GetEdgeBanding() EdgeCorrection`

GetEdgeBanding returns the EdgeBanding field if non-nil, zero value otherwise.

### GetEdgeBandingOk

`func (o *Part) GetEdgeBandingOk() (*EdgeCorrection, bool)`

GetEdgeBandingOk returns a tuple with the EdgeBanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeBanding

`func (o *Part) SetEdgeBanding(v EdgeCorrection)`

SetEdgeBanding sets EdgeBanding field to given value.

### HasEdgeBanding

`func (o *Part) HasEdgeBanding() bool`

HasEdgeBanding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


