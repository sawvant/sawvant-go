/*
 * Sawvant Cutting Optimization API — Go SDK
 *
 * File generated from our OpenAPI spec; DO NOT EDIT.
 */

package sawvant

import (
	"encoding/json"
)

// checks if the CostTariffs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CostTariffs{}

// CostTariffs struct for CostTariffs
type CostTariffs struct {
	// Fixed cost per job
	SetupCost *float64 `json:"setup_cost,omitempty"`
	// Cost per meter of cut length
	CostPerMeter *float64 `json:"cost_per_meter,omitempty"`
	CostPerRotation *float64 `json:"cost_per_rotation,omitempty"`
	CostPerStack *float64 `json:"cost_per_stack,omitempty"`
	CostPerCycle *float64 `json:"cost_per_cycle,omitempty"`
}

// NewCostTariffs instantiates a new CostTariffs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostTariffs() *CostTariffs {
	this := CostTariffs{}
	return &this
}

// NewCostTariffsWithDefaults instantiates a new CostTariffs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostTariffsWithDefaults() *CostTariffs {
	this := CostTariffs{}
	return &this
}

// GetSetupCost returns the SetupCost field value if set, zero value otherwise.
func (o *CostTariffs) GetSetupCost() float64 {
	if o == nil || IsNil(o.SetupCost) {
		var ret float64
		return ret
	}
	return *o.SetupCost
}

// GetSetupCostOk returns a tuple with the SetupCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTariffs) GetSetupCostOk() (*float64, bool) {
	if o == nil || IsNil(o.SetupCost) {
		return nil, false
	}
	return o.SetupCost, true
}

// HasSetupCost returns a boolean if a field has been set.
func (o *CostTariffs) HasSetupCost() bool {
	if o != nil && !IsNil(o.SetupCost) {
		return true
	}

	return false
}

// SetSetupCost gets a reference to the given float64 and assigns it to the SetupCost field.
func (o *CostTariffs) SetSetupCost(v float64) {
	o.SetupCost = &v
}

// GetCostPerMeter returns the CostPerMeter field value if set, zero value otherwise.
func (o *CostTariffs) GetCostPerMeter() float64 {
	if o == nil || IsNil(o.CostPerMeter) {
		var ret float64
		return ret
	}
	return *o.CostPerMeter
}

// GetCostPerMeterOk returns a tuple with the CostPerMeter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTariffs) GetCostPerMeterOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerMeter) {
		return nil, false
	}
	return o.CostPerMeter, true
}

// HasCostPerMeter returns a boolean if a field has been set.
func (o *CostTariffs) HasCostPerMeter() bool {
	if o != nil && !IsNil(o.CostPerMeter) {
		return true
	}

	return false
}

// SetCostPerMeter gets a reference to the given float64 and assigns it to the CostPerMeter field.
func (o *CostTariffs) SetCostPerMeter(v float64) {
	o.CostPerMeter = &v
}

// GetCostPerRotation returns the CostPerRotation field value if set, zero value otherwise.
func (o *CostTariffs) GetCostPerRotation() float64 {
	if o == nil || IsNil(o.CostPerRotation) {
		var ret float64
		return ret
	}
	return *o.CostPerRotation
}

// GetCostPerRotationOk returns a tuple with the CostPerRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTariffs) GetCostPerRotationOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerRotation) {
		return nil, false
	}
	return o.CostPerRotation, true
}

// HasCostPerRotation returns a boolean if a field has been set.
func (o *CostTariffs) HasCostPerRotation() bool {
	if o != nil && !IsNil(o.CostPerRotation) {
		return true
	}

	return false
}

// SetCostPerRotation gets a reference to the given float64 and assigns it to the CostPerRotation field.
func (o *CostTariffs) SetCostPerRotation(v float64) {
	o.CostPerRotation = &v
}

// GetCostPerStack returns the CostPerStack field value if set, zero value otherwise.
func (o *CostTariffs) GetCostPerStack() float64 {
	if o == nil || IsNil(o.CostPerStack) {
		var ret float64
		return ret
	}
	return *o.CostPerStack
}

// GetCostPerStackOk returns a tuple with the CostPerStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTariffs) GetCostPerStackOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerStack) {
		return nil, false
	}
	return o.CostPerStack, true
}

// HasCostPerStack returns a boolean if a field has been set.
func (o *CostTariffs) HasCostPerStack() bool {
	if o != nil && !IsNil(o.CostPerStack) {
		return true
	}

	return false
}

// SetCostPerStack gets a reference to the given float64 and assigns it to the CostPerStack field.
func (o *CostTariffs) SetCostPerStack(v float64) {
	o.CostPerStack = &v
}

// GetCostPerCycle returns the CostPerCycle field value if set, zero value otherwise.
func (o *CostTariffs) GetCostPerCycle() float64 {
	if o == nil || IsNil(o.CostPerCycle) {
		var ret float64
		return ret
	}
	return *o.CostPerCycle
}

// GetCostPerCycleOk returns a tuple with the CostPerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostTariffs) GetCostPerCycleOk() (*float64, bool) {
	if o == nil || IsNil(o.CostPerCycle) {
		return nil, false
	}
	return o.CostPerCycle, true
}

// HasCostPerCycle returns a boolean if a field has been set.
func (o *CostTariffs) HasCostPerCycle() bool {
	if o != nil && !IsNil(o.CostPerCycle) {
		return true
	}

	return false
}

// SetCostPerCycle gets a reference to the given float64 and assigns it to the CostPerCycle field.
func (o *CostTariffs) SetCostPerCycle(v float64) {
	o.CostPerCycle = &v
}

func (o CostTariffs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CostTariffs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SetupCost) {
		toSerialize["setup_cost"] = o.SetupCost
	}
	if !IsNil(o.CostPerMeter) {
		toSerialize["cost_per_meter"] = o.CostPerMeter
	}
	if !IsNil(o.CostPerRotation) {
		toSerialize["cost_per_rotation"] = o.CostPerRotation
	}
	if !IsNil(o.CostPerStack) {
		toSerialize["cost_per_stack"] = o.CostPerStack
	}
	if !IsNil(o.CostPerCycle) {
		toSerialize["cost_per_cycle"] = o.CostPerCycle
	}
	return toSerialize, nil
}

type NullableCostTariffs struct {
	value *CostTariffs
	isSet bool
}

func (v NullableCostTariffs) Get() *CostTariffs {
	return v.value
}

func (v *NullableCostTariffs) Set(val *CostTariffs) {
	v.value = val
	v.isSet = true
}

func (v NullableCostTariffs) IsSet() bool {
	return v.isSet
}

func (v *NullableCostTariffs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostTariffs(val *CostTariffs) *NullableCostTariffs {
	return &NullableCostTariffs{value: val, isSet: true}
}

func (v NullableCostTariffs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostTariffs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


