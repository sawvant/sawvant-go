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

// checks if the ApiError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiError{}

// ApiError Structured error. Used by `POST /v1/optimize`; the other endpoints still return `Error` and move to this shape later. 
type ApiError struct {
	Error ApiErrorBody `json:"error"`
}

type _ApiError ApiError

// NewApiError instantiates a new ApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError(error_ ApiErrorBody) *ApiError {
	this := ApiError{}
	this.Error = error_
	return &this
}

// NewApiErrorWithDefaults instantiates a new ApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorWithDefaults() *ApiError {
	this := ApiError{}
	return &this
}

// GetError returns the Error field value
func (o *ApiError) GetError() ApiErrorBody {
	if o == nil {
		var ret ApiErrorBody
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ApiError) GetErrorOk() (*ApiErrorBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ApiError) SetError(v ApiErrorBody) {
	o.Error = v
}

func (o ApiError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *ApiError) UnmarshalJSON(data []byte) (err error) {
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

	varApiError := _ApiError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiError)

	if err != nil {
		return err
	}

	*o = ApiError(varApiError)

	return err
}

type NullableApiError struct {
	value *ApiError
	isSet bool
}

func (v NullableApiError) Get() *ApiError {
	return v.value
}

func (v *NullableApiError) Set(val *ApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError(val *ApiError) *NullableApiError {
	return &NullableApiError{value: val, isSet: true}
}

func (v NullableApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


