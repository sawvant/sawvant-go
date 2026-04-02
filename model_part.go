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

// checks if the Part type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Part{}

// Part struct for Part
type Part struct {
	Id string `json:"id"`
	// Length in mm
	Length float64 `json:"length"`
	// Width in mm
	Width float64 `json:"width"`
	Quantity int32 `json:"quantity"`
	Grain GrainDirection `json:"grain"`
	EdgeBanding *EdgeCorrection `json:"edge_banding,omitempty"`
}

type _Part Part

// NewPart instantiates a new Part object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPart(id string, length float64, width float64, quantity int32, grain GrainDirection) *Part {
	this := Part{}
	this.Id = id
	this.Length = length
	this.Width = width
	this.Quantity = quantity
	this.Grain = grain
	return &this
}

// NewPartWithDefaults instantiates a new Part object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartWithDefaults() *Part {
	this := Part{}
	return &this
}

// GetId returns the Id field value
func (o *Part) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Part) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Part) SetId(v string) {
	o.Id = v
}

// GetLength returns the Length field value
func (o *Part) GetLength() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *Part) GetLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *Part) SetLength(v float64) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *Part) GetWidth() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Part) GetWidthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Part) SetWidth(v float64) {
	o.Width = v
}

// GetQuantity returns the Quantity field value
func (o *Part) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Part) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Part) SetQuantity(v int32) {
	o.Quantity = v
}

// GetGrain returns the Grain field value
func (o *Part) GetGrain() GrainDirection {
	if o == nil {
		var ret GrainDirection
		return ret
	}

	return o.Grain
}

// GetGrainOk returns a tuple with the Grain field value
// and a boolean to check if the value has been set.
func (o *Part) GetGrainOk() (*GrainDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grain, true
}

// SetGrain sets field value
func (o *Part) SetGrain(v GrainDirection) {
	o.Grain = v
}

// GetEdgeBanding returns the EdgeBanding field value if set, zero value otherwise.
func (o *Part) GetEdgeBanding() EdgeCorrection {
	if o == nil || IsNil(o.EdgeBanding) {
		var ret EdgeCorrection
		return ret
	}
	return *o.EdgeBanding
}

// GetEdgeBandingOk returns a tuple with the EdgeBanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetEdgeBandingOk() (*EdgeCorrection, bool) {
	if o == nil || IsNil(o.EdgeBanding) {
		return nil, false
	}
	return o.EdgeBanding, true
}

// HasEdgeBanding returns a boolean if a field has been set.
func (o *Part) HasEdgeBanding() bool {
	if o != nil && !IsNil(o.EdgeBanding) {
		return true
	}

	return false
}

// SetEdgeBanding gets a reference to the given EdgeCorrection and assigns it to the EdgeBanding field.
func (o *Part) SetEdgeBanding(v EdgeCorrection) {
	o.EdgeBanding = &v
}

func (o Part) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Part) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["length"] = o.Length
	toSerialize["width"] = o.Width
	toSerialize["quantity"] = o.Quantity
	toSerialize["grain"] = o.Grain
	if !IsNil(o.EdgeBanding) {
		toSerialize["edge_banding"] = o.EdgeBanding
	}
	return toSerialize, nil
}

func (o *Part) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"length",
		"width",
		"quantity",
		"grain",
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

	varPart := _Part{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPart)

	if err != nil {
		return err
	}

	*o = Part(varPart)

	return err
}

type NullablePart struct {
	value *Part
	isSet bool
}

func (v NullablePart) Get() *Part {
	return v.value
}

func (v *NullablePart) Set(val *Part) {
	v.value = val
	v.isSet = true
}

func (v NullablePart) IsSet() bool {
	return v.isSet
}

func (v *NullablePart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePart(val *Part) *NullablePart {
	return &NullablePart{value: val, isSet: true}
}

func (v NullablePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


