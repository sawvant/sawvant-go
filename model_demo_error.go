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

// checks if the DemoError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DemoError{}

// DemoError struct for DemoError
type DemoError struct {
	Error DemoErrorError `json:"error"`
}

type _DemoError DemoError

// NewDemoError instantiates a new DemoError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDemoError(error_ DemoErrorError) *DemoError {
	this := DemoError{}
	this.Error = error_
	return &this
}

// NewDemoErrorWithDefaults instantiates a new DemoError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDemoErrorWithDefaults() *DemoError {
	this := DemoError{}
	return &this
}

// GetError returns the Error field value
func (o *DemoError) GetError() DemoErrorError {
	if o == nil {
		var ret DemoErrorError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *DemoError) GetErrorOk() (*DemoErrorError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *DemoError) SetError(v DemoErrorError) {
	o.Error = v
}

func (o DemoError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DemoError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *DemoError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varDemoError := _DemoError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDemoError)

	if err != nil {
		return err
	}

	*o = DemoError(varDemoError)

	return err
}

type NullableDemoError struct {
	value *DemoError
	isSet bool
}

func (v NullableDemoError) Get() *DemoError {
	return v.value
}

func (v *NullableDemoError) Set(val *DemoError) {
	v.value = val
	v.isSet = true
}

func (v NullableDemoError) IsSet() bool {
	return v.isSet
}

func (v *NullableDemoError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDemoError(val *DemoError) *NullableDemoError {
	return &NullableDemoError{value: val, isSet: true}
}

func (v NullableDemoError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDemoError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


