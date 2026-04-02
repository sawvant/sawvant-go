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

// checks if the Machine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Machine{}

// Machine struct for Machine
type Machine struct {
	// Kerf width in mm
	BladeThickness float64 `json:"blade_thickness"`
	// Cut pattern complexity (1-3)
	MaxLevels int32 `json:"max_levels"`
	// Maximum stack height in mm for batch cutting
	MaxStackHeight *float64 `json:"max_stack_height,omitempty"`
	CutDirection CutDirection `json:"cut_direction"`
}

type _Machine Machine

// NewMachine instantiates a new Machine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachine(bladeThickness float64, maxLevels int32, cutDirection CutDirection) *Machine {
	this := Machine{}
	this.BladeThickness = bladeThickness
	this.MaxLevels = maxLevels
	this.CutDirection = cutDirection
	return &this
}

// NewMachineWithDefaults instantiates a new Machine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineWithDefaults() *Machine {
	this := Machine{}
	return &this
}

// GetBladeThickness returns the BladeThickness field value
func (o *Machine) GetBladeThickness() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BladeThickness
}

// GetBladeThicknessOk returns a tuple with the BladeThickness field value
// and a boolean to check if the value has been set.
func (o *Machine) GetBladeThicknessOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BladeThickness, true
}

// SetBladeThickness sets field value
func (o *Machine) SetBladeThickness(v float64) {
	o.BladeThickness = v
}

// GetMaxLevels returns the MaxLevels field value
func (o *Machine) GetMaxLevels() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxLevels
}

// GetMaxLevelsOk returns a tuple with the MaxLevels field value
// and a boolean to check if the value has been set.
func (o *Machine) GetMaxLevelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLevels, true
}

// SetMaxLevels sets field value
func (o *Machine) SetMaxLevels(v int32) {
	o.MaxLevels = v
}

// GetMaxStackHeight returns the MaxStackHeight field value if set, zero value otherwise.
func (o *Machine) GetMaxStackHeight() float64 {
	if o == nil || IsNil(o.MaxStackHeight) {
		var ret float64
		return ret
	}
	return *o.MaxStackHeight
}

// GetMaxStackHeightOk returns a tuple with the MaxStackHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Machine) GetMaxStackHeightOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxStackHeight) {
		return nil, false
	}
	return o.MaxStackHeight, true
}

// HasMaxStackHeight returns a boolean if a field has been set.
func (o *Machine) HasMaxStackHeight() bool {
	if o != nil && !IsNil(o.MaxStackHeight) {
		return true
	}

	return false
}

// SetMaxStackHeight gets a reference to the given float64 and assigns it to the MaxStackHeight field.
func (o *Machine) SetMaxStackHeight(v float64) {
	o.MaxStackHeight = &v
}

// GetCutDirection returns the CutDirection field value
func (o *Machine) GetCutDirection() CutDirection {
	if o == nil {
		var ret CutDirection
		return ret
	}

	return o.CutDirection
}

// GetCutDirectionOk returns a tuple with the CutDirection field value
// and a boolean to check if the value has been set.
func (o *Machine) GetCutDirectionOk() (*CutDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CutDirection, true
}

// SetCutDirection sets field value
func (o *Machine) SetCutDirection(v CutDirection) {
	o.CutDirection = v
}

func (o Machine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Machine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blade_thickness"] = o.BladeThickness
	toSerialize["max_levels"] = o.MaxLevels
	if !IsNil(o.MaxStackHeight) {
		toSerialize["max_stack_height"] = o.MaxStackHeight
	}
	toSerialize["cut_direction"] = o.CutDirection
	return toSerialize, nil
}

func (o *Machine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"blade_thickness",
		"max_levels",
		"cut_direction",
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

	varMachine := _Machine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMachine)

	if err != nil {
		return err
	}

	*o = Machine(varMachine)

	return err
}

type NullableMachine struct {
	value *Machine
	isSet bool
}

func (v NullableMachine) Get() *Machine {
	return v.value
}

func (v *NullableMachine) Set(val *Machine) {
	v.value = val
	v.isSet = true
}

func (v NullableMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachine(val *Machine) *NullableMachine {
	return &NullableMachine{value: val, isSet: true}
}

func (v NullableMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


