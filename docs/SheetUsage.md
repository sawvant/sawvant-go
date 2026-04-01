# SheetUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SheetId** | **string** |  | 
**ArticleNumber** | Pointer to **string** | Article/SKU reference (from request, if provided) | [optional] 
**Quantity** | **int32** | Number of this sheet type consumed | 
**YieldPercent** | **float64** | Yield percentage for this sheet type | 

## Methods

### NewSheetUsage

`func NewSheetUsage(sheetId string, quantity int32, yieldPercent float64, ) *SheetUsage`

NewSheetUsage instantiates a new SheetUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSheetUsageWithDefaults

`func NewSheetUsageWithDefaults() *SheetUsage`

NewSheetUsageWithDefaults instantiates a new SheetUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSheetId

`func (o *SheetUsage) GetSheetId() string`

GetSheetId returns the SheetId field if non-nil, zero value otherwise.

### GetSheetIdOk

`func (o *SheetUsage) GetSheetIdOk() (*string, bool)`

GetSheetIdOk returns a tuple with the SheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetId

`func (o *SheetUsage) SetSheetId(v string)`

SetSheetId sets SheetId field to given value.


### GetArticleNumber

`func (o *SheetUsage) GetArticleNumber() string`

GetArticleNumber returns the ArticleNumber field if non-nil, zero value otherwise.

### GetArticleNumberOk

`func (o *SheetUsage) GetArticleNumberOk() (*string, bool)`

GetArticleNumberOk returns a tuple with the ArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleNumber

`func (o *SheetUsage) SetArticleNumber(v string)`

SetArticleNumber sets ArticleNumber field to given value.

### HasArticleNumber

`func (o *SheetUsage) HasArticleNumber() bool`

HasArticleNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *SheetUsage) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SheetUsage) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SheetUsage) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetYieldPercent

`func (o *SheetUsage) GetYieldPercent() float64`

GetYieldPercent returns the YieldPercent field if non-nil, zero value otherwise.

### GetYieldPercentOk

`func (o *SheetUsage) GetYieldPercentOk() (*float64, bool)`

GetYieldPercentOk returns a tuple with the YieldPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldPercent

`func (o *SheetUsage) SetYieldPercent(v float64)`

SetYieldPercent sets YieldPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


