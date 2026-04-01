# Sheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Length** | **float64** | Length in mm | 
**Width** | **float64** | Width in mm | 
**Quantity** | **int32** | 0 &#x3D; unlimited | 
**Grain** | [**GrainDirection**](GrainDirection.md) |  | 
**IsOffcut** | Pointer to **bool** | Offcut sheets are prioritized by the solver | [optional] [default to false]
**TrimMargins** | Pointer to [**Margins**](Margins.md) |  | [optional] 
**ArticleNumber** | Pointer to **string** | Optional article/SKU reference for this sheet type | [optional] 

## Methods

### NewSheet

`func NewSheet(id string, length float64, width float64, quantity int32, grain GrainDirection, ) *Sheet`

NewSheet instantiates a new Sheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSheetWithDefaults

`func NewSheetWithDefaults() *Sheet`

NewSheetWithDefaults instantiates a new Sheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sheet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sheet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sheet) SetId(v string)`

SetId sets Id field to given value.


### GetLength

`func (o *Sheet) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Sheet) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Sheet) SetLength(v float64)`

SetLength sets Length field to given value.


### GetWidth

`func (o *Sheet) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Sheet) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Sheet) SetWidth(v float64)`

SetWidth sets Width field to given value.


### GetQuantity

`func (o *Sheet) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Sheet) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Sheet) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetGrain

`func (o *Sheet) GetGrain() GrainDirection`

GetGrain returns the Grain field if non-nil, zero value otherwise.

### GetGrainOk

`func (o *Sheet) GetGrainOk() (*GrainDirection, bool)`

GetGrainOk returns a tuple with the Grain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrain

`func (o *Sheet) SetGrain(v GrainDirection)`

SetGrain sets Grain field to given value.


### GetIsOffcut

`func (o *Sheet) GetIsOffcut() bool`

GetIsOffcut returns the IsOffcut field if non-nil, zero value otherwise.

### GetIsOffcutOk

`func (o *Sheet) GetIsOffcutOk() (*bool, bool)`

GetIsOffcutOk returns a tuple with the IsOffcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffcut

`func (o *Sheet) SetIsOffcut(v bool)`

SetIsOffcut sets IsOffcut field to given value.

### HasIsOffcut

`func (o *Sheet) HasIsOffcut() bool`

HasIsOffcut returns a boolean if a field has been set.

### GetTrimMargins

`func (o *Sheet) GetTrimMargins() Margins`

GetTrimMargins returns the TrimMargins field if non-nil, zero value otherwise.

### GetTrimMarginsOk

`func (o *Sheet) GetTrimMarginsOk() (*Margins, bool)`

GetTrimMarginsOk returns a tuple with the TrimMargins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimMargins

`func (o *Sheet) SetTrimMargins(v Margins)`

SetTrimMargins sets TrimMargins field to given value.

### HasTrimMargins

`func (o *Sheet) HasTrimMargins() bool`

HasTrimMargins returns a boolean if a field has been set.

### GetArticleNumber

`func (o *Sheet) GetArticleNumber() string`

GetArticleNumber returns the ArticleNumber field if non-nil, zero value otherwise.

### GetArticleNumberOk

`func (o *Sheet) GetArticleNumberOk() (*string, bool)`

GetArticleNumberOk returns a tuple with the ArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleNumber

`func (o *Sheet) SetArticleNumber(v string)`

SetArticleNumber sets ArticleNumber field to given value.

### HasArticleNumber

`func (o *Sheet) HasArticleNumber() bool`

HasArticleNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


