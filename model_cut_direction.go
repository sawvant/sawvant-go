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

// CutDirection the model 'CutDirection'
type CutDirection string

// List of CutDirection
const (
	DEFAULT CutDirection = "default"
	RIP CutDirection = "rip"
	CROSS CutDirection = "cross"
)

// All allowed values of CutDirection enum
var AllowedCutDirectionEnumValues = []CutDirection{
	"default",
	"rip",
	"cross",
}

func (v *CutDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CutDirection(value)
	for _, existing := range AllowedCutDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CutDirection", value)
}

// NewCutDirectionFromValue returns a pointer to a valid CutDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCutDirectionFromValue(v string) (*CutDirection, error) {
	ev := CutDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CutDirection: valid values are %v", v, AllowedCutDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CutDirection) IsValid() bool {
	for _, existing := range AllowedCutDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CutDirection value
func (v CutDirection) Ptr() *CutDirection {
	return &v
}

type NullableCutDirection struct {
	value *CutDirection
	isSet bool
}

func (v NullableCutDirection) Get() *CutDirection {
	return v.value
}

func (v *NullableCutDirection) Set(val *CutDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableCutDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableCutDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCutDirection(val *CutDirection) *NullableCutDirection {
	return &NullableCutDirection{value: val, isSet: true}
}

func (v NullableCutDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCutDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

