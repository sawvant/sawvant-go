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

// checks if the FieldError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldError{}

// FieldError struct for FieldError
type FieldError struct {
	// JSON path of the field, e.g. `parts[2].width` or `machine.max_levels`. For `INVALID_JSON` the path has no array index (`parts.width`). 
	Field string `json:"field"`
	Reason string `json:"reason"`
}

type _FieldError FieldError

// NewFieldError instantiates a new FieldError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldError(field string, reason string) *FieldError {
	this := FieldError{}
	this.Field = field
	this.Reason = reason
	return &this
}

// NewFieldErrorWithDefaults instantiates a new FieldError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldErrorWithDefaults() *FieldError {
	this := FieldError{}
	return &this
}

// GetField returns the Field field value
func (o *FieldError) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *FieldError) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *FieldError) SetField(v string) {
	o.Field = v
}

// GetReason returns the Reason field value
func (o *FieldError) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *FieldError) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *FieldError) SetReason(v string) {
	o.Reason = v
}

func (o FieldError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *FieldError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
		"reason",
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

	varFieldError := _FieldError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFieldError)

	if err != nil {
		return err
	}

	*o = FieldError(varFieldError)

	return err
}

type NullableFieldError struct {
	value *FieldError
	isSet bool
}

func (v NullableFieldError) Get() *FieldError {
	return v.value
}

func (v *NullableFieldError) Set(val *FieldError) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldError) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldError(val *FieldError) *NullableFieldError {
	return &NullableFieldError{value: val, isSet: true}
}

func (v NullableFieldError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


