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

// checks if the Cost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cost{}

// Cost struct for Cost
type Cost struct {
	SetupCost float64 `json:"setup_cost"`
	CuttingCost float64 `json:"cutting_cost"`
	StackingCost float64 `json:"stacking_cost"`
	RotationCost float64 `json:"rotation_cost"`
	CycleCost float64 `json:"cycle_cost"`
	TotalCost float64 `json:"total_cost"`
}

type _Cost Cost

// NewCost instantiates a new Cost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCost(setupCost float64, cuttingCost float64, stackingCost float64, rotationCost float64, cycleCost float64, totalCost float64) *Cost {
	this := Cost{}
	this.SetupCost = setupCost
	this.CuttingCost = cuttingCost
	this.StackingCost = stackingCost
	this.RotationCost = rotationCost
	this.CycleCost = cycleCost
	this.TotalCost = totalCost
	return &this
}

// NewCostWithDefaults instantiates a new Cost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostWithDefaults() *Cost {
	this := Cost{}
	return &this
}

// GetSetupCost returns the SetupCost field value
func (o *Cost) GetSetupCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SetupCost
}

// GetSetupCostOk returns a tuple with the SetupCost field value
// and a boolean to check if the value has been set.
func (o *Cost) GetSetupCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetupCost, true
}

// SetSetupCost sets field value
func (o *Cost) SetSetupCost(v float64) {
	o.SetupCost = v
}

// GetCuttingCost returns the CuttingCost field value
func (o *Cost) GetCuttingCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CuttingCost
}

// GetCuttingCostOk returns a tuple with the CuttingCost field value
// and a boolean to check if the value has been set.
func (o *Cost) GetCuttingCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CuttingCost, true
}

// SetCuttingCost sets field value
func (o *Cost) SetCuttingCost(v float64) {
	o.CuttingCost = v
}

// GetStackingCost returns the StackingCost field value
func (o *Cost) GetStackingCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.StackingCost
}

// GetStackingCostOk returns a tuple with the StackingCost field value
// and a boolean to check if the value has been set.
func (o *Cost) GetStackingCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StackingCost, true
}

// SetStackingCost sets field value
func (o *Cost) SetStackingCost(v float64) {
	o.StackingCost = v
}

// GetRotationCost returns the RotationCost field value
func (o *Cost) GetRotationCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.RotationCost
}

// GetRotationCostOk returns a tuple with the RotationCost field value
// and a boolean to check if the value has been set.
func (o *Cost) GetRotationCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RotationCost, true
}

// SetRotationCost sets field value
func (o *Cost) SetRotationCost(v float64) {
	o.RotationCost = v
}

// GetCycleCost returns the CycleCost field value
func (o *Cost) GetCycleCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CycleCost
}

// GetCycleCostOk returns a tuple with the CycleCost field value
// and a boolean to check if the value has been set.
func (o *Cost) GetCycleCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CycleCost, true
}

// SetCycleCost sets field value
func (o *Cost) SetCycleCost(v float64) {
	o.CycleCost = v
}

// GetTotalCost returns the TotalCost field value
func (o *Cost) GetTotalCost() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value
// and a boolean to check if the value has been set.
func (o *Cost) GetTotalCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCost, true
}

// SetTotalCost sets field value
func (o *Cost) SetTotalCost(v float64) {
	o.TotalCost = v
}

func (o Cost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["setup_cost"] = o.SetupCost
	toSerialize["cutting_cost"] = o.CuttingCost
	toSerialize["stacking_cost"] = o.StackingCost
	toSerialize["rotation_cost"] = o.RotationCost
	toSerialize["cycle_cost"] = o.CycleCost
	toSerialize["total_cost"] = o.TotalCost
	return toSerialize, nil
}

func (o *Cost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"setup_cost",
		"cutting_cost",
		"stacking_cost",
		"rotation_cost",
		"cycle_cost",
		"total_cost",
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

	varCost := _Cost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCost)

	if err != nil {
		return err
	}

	*o = Cost(varCost)

	return err
}

type NullableCost struct {
	value *Cost
	isSet bool
}

func (v NullableCost) Get() *Cost {
	return v.value
}

func (v *NullableCost) Set(val *Cost) {
	v.value = val
	v.isSet = true
}

func (v NullableCost) IsSet() bool {
	return v.isSet
}

func (v *NullableCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCost(val *Cost) *NullableCost {
	return &NullableCost{value: val, isSet: true}
}

func (v NullableCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


