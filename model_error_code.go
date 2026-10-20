/*
 * Sawvant Cutting Optimization API — Go SDK
 *
 * File generated from our OpenAPI spec; DO NOT EDIT.
 */

package sawvant

import (
	"encoding/json"
	"fmt"
)

// ErrorCode Stable, machine-readable error code. New codes may be added.
type ErrorCode string

// List of ErrorCode
const (
	INVALID_BODY ErrorCode = "INVALID_BODY"
	BODY_TOO_LARGE ErrorCode = "BODY_TOO_LARGE"
	INVALID_JSON ErrorCode = "INVALID_JSON"
	INVALID_STRATEGY ErrorCode = "INVALID_STRATEGY"
	PARTS_EMPTY ErrorCode = "PARTS_EMPTY"
	SHEETS_EMPTY ErrorCode = "SHEETS_EMPTY"
	TOO_MANY_PARTS ErrorCode = "TOO_MANY_PARTS"
	TOO_MANY_SHEETS ErrorCode = "TOO_MANY_SHEETS"
	VALIDATION_FAILED ErrorCode = "VALIDATION_FAILED"
	INFEASIBLE ErrorCode = "INFEASIBLE"
	INTERNAL_ERROR ErrorCode = "INTERNAL_ERROR"
	SERVICE_UNAVAILABLE ErrorCode = "SERVICE_UNAVAILABLE"
)

// All allowed values of ErrorCode enum
var AllowedErrorCodeEnumValues = []ErrorCode{
	"INVALID_BODY",
	"BODY_TOO_LARGE",
	"INVALID_JSON",
	"INVALID_STRATEGY",
	"PARTS_EMPTY",
	"SHEETS_EMPTY",
	"TOO_MANY_PARTS",
	"TOO_MANY_SHEETS",
	"VALIDATION_FAILED",
	"INFEASIBLE",
	"INTERNAL_ERROR",
	"SERVICE_UNAVAILABLE",
}

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// NewErrorCodeFromValue returns a pointer to a valid ErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodeFromValue(v string) (*ErrorCode, error) {
	ev := ErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCode: valid values are %v", v, AllowedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCode) IsValid() bool {
	for _, existing := range AllowedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

