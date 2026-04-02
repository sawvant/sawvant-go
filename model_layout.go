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

// checks if the Layout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Layout{}

// Layout struct for Layout
type Layout struct {
	SheetId string `json:"sheet_id"`
	// Number of identical sheets using this layout pattern
	Quantity int32 `json:"quantity"`
	Placements []Placement `json:"placements"`
}

type _Layout Layout

// NewLayout instantiates a new Layout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayout(sheetId string, quantity int32, placements []Placement) *Layout {
	this := Layout{}
	this.SheetId = sheetId
	this.Quantity = quantity
	this.Placements = placements
	return &this
}

// NewLayoutWithDefaults instantiates a new Layout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutWithDefaults() *Layout {
	this := Layout{}
	return &this
}

// GetSheetId returns the SheetId field value
func (o *Layout) GetSheetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SheetId
}

// GetSheetIdOk returns a tuple with the SheetId field value
// and a boolean to check if the value has been set.
func (o *Layout) GetSheetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SheetId, true
}

// SetSheetId sets field value
func (o *Layout) SetSheetId(v string) {
	o.SheetId = v
}

// GetQuantity returns the Quantity field value
func (o *Layout) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Layout) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Layout) SetQuantity(v int32) {
	o.Quantity = v
}

// GetPlacements returns the Placements field value
func (o *Layout) GetPlacements() []Placement {
	if o == nil {
		var ret []Placement
		return ret
	}

	return o.Placements
}

// GetPlacementsOk returns a tuple with the Placements field value
// and a boolean to check if the value has been set.
func (o *Layout) GetPlacementsOk() ([]Placement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placements, true
}

// SetPlacements sets field value
func (o *Layout) SetPlacements(v []Placement) {
	o.Placements = v
}

func (o Layout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Layout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sheet_id"] = o.SheetId
	toSerialize["quantity"] = o.Quantity
	toSerialize["placements"] = o.Placements
	return toSerialize, nil
}

func (o *Layout) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sheet_id",
		"quantity",
		"placements",
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

	varLayout := _Layout{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLayout)

	if err != nil {
		return err
	}

	*o = Layout(varLayout)

	return err
}

type NullableLayout struct {
	value *Layout
	isSet bool
}

func (v NullableLayout) Get() *Layout {
	return v.value
}

func (v *NullableLayout) Set(val *Layout) {
	v.value = val
	v.isSet = true
}

func (v NullableLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayout(val *Layout) *NullableLayout {
	return &NullableLayout{value: val, isSet: true}
}

func (v NullableLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


