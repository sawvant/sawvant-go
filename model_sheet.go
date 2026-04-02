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

// checks if the Sheet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sheet{}

// Sheet struct for Sheet
type Sheet struct {
	Id string `json:"id"`
	// Length in mm
	Length float64 `json:"length"`
	// Width in mm
	Width float64 `json:"width"`
	// 0 = unlimited
	Quantity int32 `json:"quantity"`
	Grain GrainDirection `json:"grain"`
	// Offcut sheets are prioritized by the solver
	IsOffcut *bool `json:"is_offcut,omitempty"`
	TrimMargins *Margins `json:"trim_margins,omitempty"`
	// Optional article/SKU reference for this sheet type
	ArticleNumber *string `json:"article_number,omitempty"`
}

type _Sheet Sheet

// NewSheet instantiates a new Sheet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSheet(id string, length float64, width float64, quantity int32, grain GrainDirection) *Sheet {
	this := Sheet{}
	this.Id = id
	this.Length = length
	this.Width = width
	this.Quantity = quantity
	this.Grain = grain
	var isOffcut bool = false
	this.IsOffcut = &isOffcut
	return &this
}

// NewSheetWithDefaults instantiates a new Sheet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSheetWithDefaults() *Sheet {
	this := Sheet{}
	var isOffcut bool = false
	this.IsOffcut = &isOffcut
	return &this
}

// GetId returns the Id field value
func (o *Sheet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Sheet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Sheet) SetId(v string) {
	o.Id = v
}

// GetLength returns the Length field value
func (o *Sheet) GetLength() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *Sheet) GetLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *Sheet) SetLength(v float64) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *Sheet) GetWidth() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Sheet) GetWidthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Sheet) SetWidth(v float64) {
	o.Width = v
}

// GetQuantity returns the Quantity field value
func (o *Sheet) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Sheet) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Sheet) SetQuantity(v int32) {
	o.Quantity = v
}

// GetGrain returns the Grain field value
func (o *Sheet) GetGrain() GrainDirection {
	if o == nil {
		var ret GrainDirection
		return ret
	}

	return o.Grain
}

// GetGrainOk returns a tuple with the Grain field value
// and a boolean to check if the value has been set.
func (o *Sheet) GetGrainOk() (*GrainDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grain, true
}

// SetGrain sets field value
func (o *Sheet) SetGrain(v GrainDirection) {
	o.Grain = v
}

// GetIsOffcut returns the IsOffcut field value if set, zero value otherwise.
func (o *Sheet) GetIsOffcut() bool {
	if o == nil || IsNil(o.IsOffcut) {
		var ret bool
		return ret
	}
	return *o.IsOffcut
}

// GetIsOffcutOk returns a tuple with the IsOffcut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sheet) GetIsOffcutOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOffcut) {
		return nil, false
	}
	return o.IsOffcut, true
}

// HasIsOffcut returns a boolean if a field has been set.
func (o *Sheet) HasIsOffcut() bool {
	if o != nil && !IsNil(o.IsOffcut) {
		return true
	}

	return false
}

// SetIsOffcut gets a reference to the given bool and assigns it to the IsOffcut field.
func (o *Sheet) SetIsOffcut(v bool) {
	o.IsOffcut = &v
}

// GetTrimMargins returns the TrimMargins field value if set, zero value otherwise.
func (o *Sheet) GetTrimMargins() Margins {
	if o == nil || IsNil(o.TrimMargins) {
		var ret Margins
		return ret
	}
	return *o.TrimMargins
}

// GetTrimMarginsOk returns a tuple with the TrimMargins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sheet) GetTrimMarginsOk() (*Margins, bool) {
	if o == nil || IsNil(o.TrimMargins) {
		return nil, false
	}
	return o.TrimMargins, true
}

// HasTrimMargins returns a boolean if a field has been set.
func (o *Sheet) HasTrimMargins() bool {
	if o != nil && !IsNil(o.TrimMargins) {
		return true
	}

	return false
}

// SetTrimMargins gets a reference to the given Margins and assigns it to the TrimMargins field.
func (o *Sheet) SetTrimMargins(v Margins) {
	o.TrimMargins = &v
}

// GetArticleNumber returns the ArticleNumber field value if set, zero value otherwise.
func (o *Sheet) GetArticleNumber() string {
	if o == nil || IsNil(o.ArticleNumber) {
		var ret string
		return ret
	}
	return *o.ArticleNumber
}

// GetArticleNumberOk returns a tuple with the ArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sheet) GetArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ArticleNumber) {
		return nil, false
	}
	return o.ArticleNumber, true
}

// HasArticleNumber returns a boolean if a field has been set.
func (o *Sheet) HasArticleNumber() bool {
	if o != nil && !IsNil(o.ArticleNumber) {
		return true
	}

	return false
}

// SetArticleNumber gets a reference to the given string and assigns it to the ArticleNumber field.
func (o *Sheet) SetArticleNumber(v string) {
	o.ArticleNumber = &v
}

func (o Sheet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sheet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["length"] = o.Length
	toSerialize["width"] = o.Width
	toSerialize["quantity"] = o.Quantity
	toSerialize["grain"] = o.Grain
	if !IsNil(o.IsOffcut) {
		toSerialize["is_offcut"] = o.IsOffcut
	}
	if !IsNil(o.TrimMargins) {
		toSerialize["trim_margins"] = o.TrimMargins
	}
	if !IsNil(o.ArticleNumber) {
		toSerialize["article_number"] = o.ArticleNumber
	}
	return toSerialize, nil
}

func (o *Sheet) UnmarshalJSON(data []byte) (err error) {
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

	varSheet := _Sheet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSheet)

	if err != nil {
		return err
	}

	*o = Sheet(varSheet)

	return err
}

type NullableSheet struct {
	value *Sheet
	isSet bool
}

func (v NullableSheet) Get() *Sheet {
	return v.value
}

func (v *NullableSheet) Set(val *Sheet) {
	v.value = val
	v.isSet = true
}

func (v NullableSheet) IsSet() bool {
	return v.isSet
}

func (v *NullableSheet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSheet(val *Sheet) *NullableSheet {
	return &NullableSheet{value: val, isSet: true}
}

func (v NullableSheet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSheet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


