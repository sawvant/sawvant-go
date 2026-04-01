# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSheets** | **int32** |  | 
**YieldPercent** | **float64** |  | 
**WastePercent** | **float64** |  | 
**WasteArea** | **float64** | mm² | 
**UnplacedParts** | Pointer to **[]string** |  | [optional] 
**SheetsUsed** | [**[]SheetUsage**](SheetUsage.md) |  | 

## Methods

### NewSummary

`func NewSummary(totalSheets int32, yieldPercent float64, wastePercent float64, wasteArea float64, sheetsUsed []SheetUsage, ) *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSheets

`func (o *Summary) GetTotalSheets() int32`

GetTotalSheets returns the TotalSheets field if non-nil, zero value otherwise.

### GetTotalSheetsOk

`func (o *Summary) GetTotalSheetsOk() (*int32, bool)`

GetTotalSheetsOk returns a tuple with the TotalSheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSheets

`func (o *Summary) SetTotalSheets(v int32)`

SetTotalSheets sets TotalSheets field to given value.


### GetYieldPercent

`func (o *Summary) GetYieldPercent() float64`

GetYieldPercent returns the YieldPercent field if non-nil, zero value otherwise.

### GetYieldPercentOk

`func (o *Summary) GetYieldPercentOk() (*float64, bool)`

GetYieldPercentOk returns a tuple with the YieldPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldPercent

`func (o *Summary) SetYieldPercent(v float64)`

SetYieldPercent sets YieldPercent field to given value.


### GetWastePercent

`func (o *Summary) GetWastePercent() float64`

GetWastePercent returns the WastePercent field if non-nil, zero value otherwise.

### GetWastePercentOk

`func (o *Summary) GetWastePercentOk() (*float64, bool)`

GetWastePercentOk returns a tuple with the WastePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWastePercent

`func (o *Summary) SetWastePercent(v float64)`

SetWastePercent sets WastePercent field to given value.


### GetWasteArea

`func (o *Summary) GetWasteArea() float64`

GetWasteArea returns the WasteArea field if non-nil, zero value otherwise.

### GetWasteAreaOk

`func (o *Summary) GetWasteAreaOk() (*float64, bool)`

GetWasteAreaOk returns a tuple with the WasteArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasteArea

`func (o *Summary) SetWasteArea(v float64)`

SetWasteArea sets WasteArea field to given value.


### GetUnplacedParts

`func (o *Summary) GetUnplacedParts() []string`

GetUnplacedParts returns the UnplacedParts field if non-nil, zero value otherwise.

### GetUnplacedPartsOk

`func (o *Summary) GetUnplacedPartsOk() (*[]string, bool)`

GetUnplacedPartsOk returns a tuple with the UnplacedParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplacedParts

`func (o *Summary) SetUnplacedParts(v []string)`

SetUnplacedParts sets UnplacedParts field to given value.

### HasUnplacedParts

`func (o *Summary) HasUnplacedParts() bool`

HasUnplacedParts returns a boolean if a field has been set.

### GetSheetsUsed

`func (o *Summary) GetSheetsUsed() []SheetUsage`

GetSheetsUsed returns the SheetsUsed field if non-nil, zero value otherwise.

### GetSheetsUsedOk

`func (o *Summary) GetSheetsUsedOk() (*[]SheetUsage, bool)`

GetSheetsUsedOk returns a tuple with the SheetsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetsUsed

`func (o *Summary) SetSheetsUsed(v []SheetUsage)`

SetSheetsUsed sets SheetsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


