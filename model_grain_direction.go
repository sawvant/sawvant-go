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

// GrainDirection Which axis the grain runs along. On a part it is relative to the part's own length and width; on a sheet it is relative to the sheet's. A part may only be placed so that its grain ends up parallel to the sheet's. A sheet that omits the field is treated as having its grain along its length, so `length` and `width` on a part constrain rotation even when the sheet says nothing. Use `none` on the part for material without a grain, or `free_same` to let the solver choose one orientation and apply it to every instance of that part.
type GrainDirection string

// List of GrainDirection
const (
	NONE GrainDirection = "none"
	LENGTH GrainDirection = "length"
	WIDTH GrainDirection = "width"
	FREE_SAME GrainDirection = "free_same"
)

// All allowed values of GrainDirection enum
var AllowedGrainDirectionEnumValues = []GrainDirection{
	"none",
	"length",
	"width",
	"free_same",
}

func (v *GrainDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrainDirection(value)
	for _, existing := range AllowedGrainDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrainDirection", value)
}

// NewGrainDirectionFromValue returns a pointer to a valid GrainDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrainDirectionFromValue(v string) (*GrainDirection, error) {
	ev := GrainDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrainDirection: valid values are %v", v, AllowedGrainDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrainDirection) IsValid() bool {
	for _, existing := range AllowedGrainDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GrainDirection value
func (v GrainDirection) Ptr() *GrainDirection {
	return &v
}

type NullableGrainDirection struct {
	value *GrainDirection
	isSet bool
}

func (v NullableGrainDirection) Get() *GrainDirection {
	return v.value
}

func (v *NullableGrainDirection) Set(val *GrainDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableGrainDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableGrainDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrainDirection(val *GrainDirection) *NullableGrainDirection {
	return &NullableGrainDirection{value: val, isSet: true}
}

func (v NullableGrainDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrainDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

