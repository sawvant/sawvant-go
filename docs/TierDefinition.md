# TierDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**PriceEurCents** | **int32** |  | 
**Period** | Pointer to **string** |  | [optional] 
**RateLimitFast** | **int32** | Maximum fast strategy requests per 24h sliding window | 
**RateLimitThorough** | **int32** | Maximum thorough strategy requests per 24h sliding window | 
**Features** | **map[string]bool** | Feature gates enabled for this tier | 

## Methods

### NewTierDefinition

`func NewTierDefinition(name string, displayName string, priceEurCents int32, rateLimitFast int32, rateLimitThorough int32, features map[string]bool, ) *TierDefinition`

NewTierDefinition instantiates a new TierDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierDefinitionWithDefaults

`func NewTierDefinitionWithDefaults() *TierDefinition`

NewTierDefinitionWithDefaults instantiates a new TierDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TierDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TierDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *TierDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TierDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TierDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPriceEurCents

`func (o *TierDefinition) GetPriceEurCents() int32`

GetPriceEurCents returns the PriceEurCents field if non-nil, zero value otherwise.

### GetPriceEurCentsOk

`func (o *TierDefinition) GetPriceEurCentsOk() (*int32, bool)`

GetPriceEurCentsOk returns a tuple with the PriceEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEurCents

`func (o *TierDefinition) SetPriceEurCents(v int32)`

SetPriceEurCents sets PriceEurCents field to given value.


### GetPeriod

`func (o *TierDefinition) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TierDefinition) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TierDefinition) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TierDefinition) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRateLimitFast

`func (o *TierDefinition) GetRateLimitFast() int32`

GetRateLimitFast returns the RateLimitFast field if non-nil, zero value otherwise.

### GetRateLimitFastOk

`func (o *TierDefinition) GetRateLimitFastOk() (*int32, bool)`

GetRateLimitFastOk returns a tuple with the RateLimitFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitFast

`func (o *TierDefinition) SetRateLimitFast(v int32)`

SetRateLimitFast sets RateLimitFast field to given value.


### GetRateLimitThorough

`func (o *TierDefinition) GetRateLimitThorough() int32`

GetRateLimitThorough returns the RateLimitThorough field if non-nil, zero value otherwise.

### GetRateLimitThoroughOk

`func (o *TierDefinition) GetRateLimitThoroughOk() (*int32, bool)`

GetRateLimitThoroughOk returns a tuple with the RateLimitThorough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitThorough

`func (o *TierDefinition) SetRateLimitThorough(v int32)`

SetRateLimitThorough sets RateLimitThorough field to given value.


### GetFeatures

`func (o *TierDefinition) GetFeatures() map[string]bool`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TierDefinition) GetFeaturesOk() (*map[string]bool, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TierDefinition) SetFeatures(v map[string]bool)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


