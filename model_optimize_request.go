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

// checks if the OptimizeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizeRequest{}

// OptimizeRequest struct for OptimizeRequest
type OptimizeRequest struct {
	Parts []Part `json:"parts"`
	Sheets []Sheet `json:"sheets"`
	Machine Machine `json:"machine"`
	// Solve strategy. \"fast\" runs all greedy solvers concurrently. \"thorough\" adds Gilmore-Gomory column generation for optimal patterns. Each strategy has its own rate limit quota. 
	Strategy *string `json:"strategy,omitempty"`
	CostTariffs *CostTariffs `json:"cost_tariffs,omitempty"`
}

type _OptimizeRequest OptimizeRequest

// NewOptimizeRequest instantiates a new OptimizeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizeRequest(parts []Part, sheets []Sheet, machine Machine) *OptimizeRequest {
	this := OptimizeRequest{}
	this.Parts = parts
	this.Sheets = sheets
	this.Machine = machine
	var strategy string = "fast"
	this.Strategy = &strategy
	return &this
}

// NewOptimizeRequestWithDefaults instantiates a new OptimizeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizeRequestWithDefaults() *OptimizeRequest {
	this := OptimizeRequest{}
	var strategy string = "fast"
	this.Strategy = &strategy
	return &this
}

// GetParts returns the Parts field value
func (o *OptimizeRequest) GetParts() []Part {
	if o == nil {
		var ret []Part
		return ret
	}

	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value
// and a boolean to check if the value has been set.
func (o *OptimizeRequest) GetPartsOk() ([]Part, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parts, true
}

// SetParts sets field value
func (o *OptimizeRequest) SetParts(v []Part) {
	o.Parts = v
}

// GetSheets returns the Sheets field value
func (o *OptimizeRequest) GetSheets() []Sheet {
	if o == nil {
		var ret []Sheet
		return ret
	}

	return o.Sheets
}

// GetSheetsOk returns a tuple with the Sheets field value
// and a boolean to check if the value has been set.
func (o *OptimizeRequest) GetSheetsOk() ([]Sheet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sheets, true
}

// SetSheets sets field value
func (o *OptimizeRequest) SetSheets(v []Sheet) {
	o.Sheets = v
}

// GetMachine returns the Machine field value
func (o *OptimizeRequest) GetMachine() Machine {
	if o == nil {
		var ret Machine
		return ret
	}

	return o.Machine
}

// GetMachineOk returns a tuple with the Machine field value
// and a boolean to check if the value has been set.
func (o *OptimizeRequest) GetMachineOk() (*Machine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Machine, true
}

// SetMachine sets field value
func (o *OptimizeRequest) SetMachine(v Machine) {
	o.Machine = v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *OptimizeRequest) GetStrategy() string {
	if o == nil || IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizeRequest) GetStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *OptimizeRequest) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *OptimizeRequest) SetStrategy(v string) {
	o.Strategy = &v
}

// GetCostTariffs returns the CostTariffs field value if set, zero value otherwise.
func (o *OptimizeRequest) GetCostTariffs() CostTariffs {
	if o == nil || IsNil(o.CostTariffs) {
		var ret CostTariffs
		return ret
	}
	return *o.CostTariffs
}

// GetCostTariffsOk returns a tuple with the CostTariffs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizeRequest) GetCostTariffsOk() (*CostTariffs, bool) {
	if o == nil || IsNil(o.CostTariffs) {
		return nil, false
	}
	return o.CostTariffs, true
}

// HasCostTariffs returns a boolean if a field has been set.
func (o *OptimizeRequest) HasCostTariffs() bool {
	if o != nil && !IsNil(o.CostTariffs) {
		return true
	}

	return false
}

// SetCostTariffs gets a reference to the given CostTariffs and assigns it to the CostTariffs field.
func (o *OptimizeRequest) SetCostTariffs(v CostTariffs) {
	o.CostTariffs = &v
}

func (o OptimizeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptimizeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parts"] = o.Parts
	toSerialize["sheets"] = o.Sheets
	toSerialize["machine"] = o.Machine
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.CostTariffs) {
		toSerialize["cost_tariffs"] = o.CostTariffs
	}
	return toSerialize, nil
}

func (o *OptimizeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parts",
		"sheets",
		"machine",
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

	varOptimizeRequest := _OptimizeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOptimizeRequest)

	if err != nil {
		return err
	}

	*o = OptimizeRequest(varOptimizeRequest)

	return err
}

type NullableOptimizeRequest struct {
	value *OptimizeRequest
	isSet bool
}

func (v NullableOptimizeRequest) Get() *OptimizeRequest {
	return v.value
}

func (v *NullableOptimizeRequest) Set(val *OptimizeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizeRequest(val *OptimizeRequest) *NullableOptimizeRequest {
	return &NullableOptimizeRequest{value: val, isSet: true}
}

func (v NullableOptimizeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptimizeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


