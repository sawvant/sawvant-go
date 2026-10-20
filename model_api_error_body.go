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

// checks if the ApiErrorBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorBody{}

// ApiErrorBody struct for ApiErrorBody
type ApiErrorBody struct {
	Code ErrorCode `json:"code"`
	// Human-readable explanation. Do not parse it; use `code` and `details`.
	Message string `json:"message"`
	Details []FieldError `json:"details,omitempty"`
}

type _ApiErrorBody ApiErrorBody

// NewApiErrorBody instantiates a new ApiErrorBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorBody(code ErrorCode, message string) *ApiErrorBody {
	this := ApiErrorBody{}
	this.Code = code
	this.Message = message
	return &this
}

// NewApiErrorBodyWithDefaults instantiates a new ApiErrorBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorBodyWithDefaults() *ApiErrorBody {
	this := ApiErrorBody{}
	return &this
}

// GetCode returns the Code field value
func (o *ApiErrorBody) GetCode() ErrorCode {
	if o == nil {
		var ret ErrorCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiErrorBody) GetCodeOk() (*ErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiErrorBody) SetCode(v ErrorCode) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ApiErrorBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiErrorBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiErrorBody) SetMessage(v string) {
	o.Message = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ApiErrorBody) GetDetails() []FieldError {
	if o == nil || IsNil(o.Details) {
		var ret []FieldError
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorBody) GetDetailsOk() ([]FieldError, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ApiErrorBody) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []FieldError and assigns it to the Details field.
func (o *ApiErrorBody) SetDetails(v []FieldError) {
	o.Details = v
}

func (o ApiErrorBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

func (o *ApiErrorBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varApiErrorBody := _ApiErrorBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiErrorBody)

	if err != nil {
		return err
	}

	*o = ApiErrorBody(varApiErrorBody)

	return err
}

type NullableApiErrorBody struct {
	value *ApiErrorBody
	isSet bool
}

func (v NullableApiErrorBody) Get() *ApiErrorBody {
	return v.value
}

func (v *NullableApiErrorBody) Set(val *ApiErrorBody) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorBody) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorBody(val *ApiErrorBody) *NullableApiErrorBody {
	return &NullableApiErrorBody{value: val, isSet: true}
}

func (v NullableApiErrorBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


