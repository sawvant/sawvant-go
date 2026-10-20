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

// checks if the DemoLayout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DemoLayout{}

// DemoLayout struct for DemoLayout
type DemoLayout struct {
	SheetId string `json:"sheet_id"`
	// Number of identical sheets using this layout pattern
	Quantity int32 `json:"quantity"`
	Placements []Placement `json:"placements"`
	// Cutting plan of this layout as a standalone SVG document
	Svg string `json:"svg"`
}

type _DemoLayout DemoLayout

// NewDemoLayout instantiates a new DemoLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDemoLayout(sheetId string, quantity int32, placements []Placement, svg string) *DemoLayout {
	this := DemoLayout{}
	this.SheetId = sheetId
	this.Quantity = quantity
	this.Placements = placements
	this.Svg = svg
	return &this
}

// NewDemoLayoutWithDefaults instantiates a new DemoLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDemoLayoutWithDefaults() *DemoLayout {
	this := DemoLayout{}
	return &this
}

// GetSheetId returns the SheetId field value
func (o *DemoLayout) GetSheetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SheetId
}

// GetSheetIdOk returns a tuple with the SheetId field value
// and a boolean to check if the value has been set.
func (o *DemoLayout) GetSheetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SheetId, true
}

// SetSheetId sets field value
func (o *DemoLayout) SetSheetId(v string) {
	o.SheetId = v
}

// GetQuantity returns the Quantity field value
func (o *DemoLayout) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *DemoLayout) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *DemoLayout) SetQuantity(v int32) {
	o.Quantity = v
}

// GetPlacements returns the Placements field value
func (o *DemoLayout) GetPlacements() []Placement {
	if o == nil {
		var ret []Placement
		return ret
	}

	return o.Placements
}

// GetPlacementsOk returns a tuple with the Placements field value
// and a boolean to check if the value has been set.
func (o *DemoLayout) GetPlacementsOk() ([]Placement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placements, true
}

// SetPlacements sets field value
func (o *DemoLayout) SetPlacements(v []Placement) {
	o.Placements = v
}

// GetSvg returns the Svg field value
func (o *DemoLayout) GetSvg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Svg
}

// GetSvgOk returns a tuple with the Svg field value
// and a boolean to check if the value has been set.
func (o *DemoLayout) GetSvgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Svg, true
}

// SetSvg sets field value
func (o *DemoLayout) SetSvg(v string) {
	o.Svg = v
}

func (o DemoLayout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DemoLayout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sheet_id"] = o.SheetId
	toSerialize["quantity"] = o.Quantity
	toSerialize["placements"] = o.Placements
	toSerialize["svg"] = o.Svg
	return toSerialize, nil
}

func (o *DemoLayout) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sheet_id",
		"quantity",
		"placements",
		"svg",
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

	varDemoLayout := _DemoLayout{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDemoLayout)

	if err != nil {
		return err
	}

	*o = DemoLayout(varDemoLayout)

	return err
}

type NullableDemoLayout struct {
	value *DemoLayout
	isSet bool
}

func (v NullableDemoLayout) Get() *DemoLayout {
	return v.value
}

func (v *NullableDemoLayout) Set(val *DemoLayout) {
	v.value = val
	v.isSet = true
}

func (v NullableDemoLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableDemoLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDemoLayout(val *DemoLayout) *NullableDemoLayout {
	return &NullableDemoLayout{value: val, isSet: true}
}

func (v NullableDemoLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDemoLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


