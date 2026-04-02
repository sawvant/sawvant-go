/*
 * Sawvant Cutting Optimization API — Go SDK
 *
 * File generated from our OpenAPI spec; DO NOT EDIT.
 */

package sawvant

import (
	"encoding/json"
)

// checks if the EdgeCorrection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeCorrection{}

// EdgeCorrection struct for EdgeCorrection
type EdgeCorrection struct {
	// mm added for edge banding
	Top *float64 `json:"top,omitempty"`
	Bottom *float64 `json:"bottom,omitempty"`
	Left *float64 `json:"left,omitempty"`
	Right *float64 `json:"right,omitempty"`
}

// NewEdgeCorrection instantiates a new EdgeCorrection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeCorrection() *EdgeCorrection {
	this := EdgeCorrection{}
	var top float64 = 0
	this.Top = &top
	var bottom float64 = 0
	this.Bottom = &bottom
	var left float64 = 0
	this.Left = &left
	var right float64 = 0
	this.Right = &right
	return &this
}

// NewEdgeCorrectionWithDefaults instantiates a new EdgeCorrection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeCorrectionWithDefaults() *EdgeCorrection {
	this := EdgeCorrection{}
	var top float64 = 0
	this.Top = &top
	var bottom float64 = 0
	this.Bottom = &bottom
	var left float64 = 0
	this.Left = &left
	var right float64 = 0
	this.Right = &right
	return &this
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *EdgeCorrection) GetTop() float64 {
	if o == nil || IsNil(o.Top) {
		var ret float64
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeCorrection) GetTopOk() (*float64, bool) {
	if o == nil || IsNil(o.Top) {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *EdgeCorrection) HasTop() bool {
	if o != nil && !IsNil(o.Top) {
		return true
	}

	return false
}

// SetTop gets a reference to the given float64 and assigns it to the Top field.
func (o *EdgeCorrection) SetTop(v float64) {
	o.Top = &v
}

// GetBottom returns the Bottom field value if set, zero value otherwise.
func (o *EdgeCorrection) GetBottom() float64 {
	if o == nil || IsNil(o.Bottom) {
		var ret float64
		return ret
	}
	return *o.Bottom
}

// GetBottomOk returns a tuple with the Bottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeCorrection) GetBottomOk() (*float64, bool) {
	if o == nil || IsNil(o.Bottom) {
		return nil, false
	}
	return o.Bottom, true
}

// HasBottom returns a boolean if a field has been set.
func (o *EdgeCorrection) HasBottom() bool {
	if o != nil && !IsNil(o.Bottom) {
		return true
	}

	return false
}

// SetBottom gets a reference to the given float64 and assigns it to the Bottom field.
func (o *EdgeCorrection) SetBottom(v float64) {
	o.Bottom = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *EdgeCorrection) GetLeft() float64 {
	if o == nil || IsNil(o.Left) {
		var ret float64
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeCorrection) GetLeftOk() (*float64, bool) {
	if o == nil || IsNil(o.Left) {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *EdgeCorrection) HasLeft() bool {
	if o != nil && !IsNil(o.Left) {
		return true
	}

	return false
}

// SetLeft gets a reference to the given float64 and assigns it to the Left field.
func (o *EdgeCorrection) SetLeft(v float64) {
	o.Left = &v
}

// GetRight returns the Right field value if set, zero value otherwise.
func (o *EdgeCorrection) GetRight() float64 {
	if o == nil || IsNil(o.Right) {
		var ret float64
		return ret
	}
	return *o.Right
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeCorrection) GetRightOk() (*float64, bool) {
	if o == nil || IsNil(o.Right) {
		return nil, false
	}
	return o.Right, true
}

// HasRight returns a boolean if a field has been set.
func (o *EdgeCorrection) HasRight() bool {
	if o != nil && !IsNil(o.Right) {
		return true
	}

	return false
}

// SetRight gets a reference to the given float64 and assigns it to the Right field.
func (o *EdgeCorrection) SetRight(v float64) {
	o.Right = &v
}

func (o EdgeCorrection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeCorrection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Top) {
		toSerialize["top"] = o.Top
	}
	if !IsNil(o.Bottom) {
		toSerialize["bottom"] = o.Bottom
	}
	if !IsNil(o.Left) {
		toSerialize["left"] = o.Left
	}
	if !IsNil(o.Right) {
		toSerialize["right"] = o.Right
	}
	return toSerialize, nil
}

type NullableEdgeCorrection struct {
	value *EdgeCorrection
	isSet bool
}

func (v NullableEdgeCorrection) Get() *EdgeCorrection {
	return v.value
}

func (v *NullableEdgeCorrection) Set(val *EdgeCorrection) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeCorrection) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeCorrection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeCorrection(val *EdgeCorrection) *NullableEdgeCorrection {
	return &NullableEdgeCorrection{value: val, isSet: true}
}

func (v NullableEdgeCorrection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeCorrection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


