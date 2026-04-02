/*
 * Sawvant Cutting Optimization API — Go SDK
 *
 * File generated from our OpenAPI spec; DO NOT EDIT.
 */

package sawvant

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TierDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TierDefinition{}

// TierDefinition struct for TierDefinition
type TierDefinition struct {
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	PriceEurCents int32 `json:"price_eur_cents"`
	Period *string `json:"period,omitempty"`
	// Maximum fast strategy requests per 24h sliding window
	RateLimitFast int32 `json:"rate_limit_fast"`
	// Maximum thorough strategy requests per 24h sliding window
	RateLimitThorough int32 `json:"rate_limit_thorough"`
	// Feature gates enabled for this tier
	Features map[string]bool `json:"features"`
}

type _TierDefinition TierDefinition

// NewTierDefinition instantiates a new TierDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTierDefinition(name string, displayName string, priceEurCents int32, rateLimitFast int32, rateLimitThorough int32, features map[string]bool) *TierDefinition {
	this := TierDefinition{}
	this.Name = name
	this.DisplayName = displayName
	this.PriceEurCents = priceEurCents
	this.RateLimitFast = rateLimitFast
	this.RateLimitThorough = rateLimitThorough
	this.Features = features
	return &this
}

// NewTierDefinitionWithDefaults instantiates a new TierDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTierDefinitionWithDefaults() *TierDefinition {
	this := TierDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *TierDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TierDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TierDefinition) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *TierDefinition) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *TierDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *TierDefinition) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPriceEurCents returns the PriceEurCents field value
func (o *TierDefinition) GetPriceEurCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriceEurCents
}

// GetPriceEurCentsOk returns a tuple with the PriceEurCents field value
// and a boolean to check if the value has been set.
func (o *TierDefinition) GetPriceEurCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceEurCents, true
}

// SetPriceEurCents sets field value
func (o *TierDefinition) SetPriceEurCents(v int32) {
	o.PriceEurCents = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *TierDefinition) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TierDefinition) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *TierDefinition) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *TierDefinition) SetPeriod(v string) {
	o.Period = &v
}

// GetRateLimitFast returns the RateLimitFast field value
func (o *TierDefinition) GetRateLimitFast() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RateLimitFast
}

// GetRateLimitFastOk returns a tuple with the RateLimitFast field value
// and a boolean to check if the value has been set.
func (o *TierDefinition) GetRateLimitFastOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateLimitFast, true
}

// SetRateLimitFast sets field value
func (o *TierDefinition) SetRateLimitFast(v int32) {
	o.RateLimitFast = v
}

// GetRateLimitThorough returns the RateLimitThorough field value
func (o *TierDefinition) GetRateLimitThorough() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RateLimitThorough
}

// GetRateLimitThoroughOk returns a tuple with the RateLimitThorough field value
// and a boolean to check if the value has been set.
func (o *TierDefinition) GetRateLimitThoroughOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateLimitThorough, true
}

// SetRateLimitThorough sets field value
func (o *TierDefinition) SetRateLimitThorough(v int32) {
	o.RateLimitThorough = v
}

// GetFeatures returns the Features field value
func (o *TierDefinition) GetFeatures() map[string]bool {
	if o == nil {
		var ret map[string]bool
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *TierDefinition) GetFeaturesOk() (map[string]bool, bool) {
	if o == nil {
		return map[string]bool{}, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *TierDefinition) SetFeatures(v map[string]bool) {
	o.Features = v
}

func (o TierDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TierDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName
	toSerialize["price_eur_cents"] = o.PriceEurCents
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	toSerialize["rate_limit_fast"] = o.RateLimitFast
	toSerialize["rate_limit_thorough"] = o.RateLimitThorough
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

func (o *TierDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"display_name",
		"price_eur_cents",
		"rate_limit_fast",
		"rate_limit_thorough",
		"features",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTierDefinition := _TierDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTierDefinition)

	if err != nil {
		return err
	}

	*o = TierDefinition(varTierDefinition)

	return err
}

type NullableTierDefinition struct {
	value *TierDefinition
	isSet bool
}

func (v NullableTierDefinition) Get() *TierDefinition {
	return v.value
}

func (v *NullableTierDefinition) Set(val *TierDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTierDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTierDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTierDefinition(val *TierDefinition) *NullableTierDefinition {
	return &NullableTierDefinition{value: val, isSet: true}
}

func (v NullableTierDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTierDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


