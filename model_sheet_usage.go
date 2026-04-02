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

// checks if the SheetUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SheetUsage{}

// SheetUsage struct for SheetUsage
type SheetUsage struct {
	SheetId string `json:"sheet_id"`
	// Article/SKU reference (from request, if provided)
	ArticleNumber *string `json:"article_number,omitempty"`
	// Number of this sheet type consumed
	Quantity int32 `json:"quantity"`
	// Yield percentage for this sheet type
	YieldPercent float64 `json:"yield_percent"`
}

type _SheetUsage SheetUsage

// NewSheetUsage instantiates a new SheetUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSheetUsage(sheetId string, quantity int32, yieldPercent float64) *SheetUsage {
	this := SheetUsage{}
	this.SheetId = sheetId
	this.Quantity = quantity
	this.YieldPercent = yieldPercent
	return &this
}

// NewSheetUsageWithDefaults instantiates a new SheetUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSheetUsageWithDefaults() *SheetUsage {
	this := SheetUsage{}
	return &this
}

// GetSheetId returns the SheetId field value
func (o *SheetUsage) GetSheetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SheetId
}

// GetSheetIdOk returns a tuple with the SheetId field value
// and a boolean to check if the value has been set.
func (o *SheetUsage) GetSheetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SheetId, true
}

// SetSheetId sets field value
func (o *SheetUsage) SetSheetId(v string) {
	o.SheetId = v
}

// GetArticleNumber returns the ArticleNumber field value if set, zero value otherwise.
func (o *SheetUsage) GetArticleNumber() string {
	if o == nil || IsNil(o.ArticleNumber) {
		var ret string
		return ret
	}
	return *o.ArticleNumber
}

// GetArticleNumberOk returns a tuple with the ArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SheetUsage) GetArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ArticleNumber) {
		return nil, false
	}
	return o.ArticleNumber, true
}

// HasArticleNumber returns a boolean if a field has been set.
func (o *SheetUsage) HasArticleNumber() bool {
	if o != nil && !IsNil(o.ArticleNumber) {
		return true
	}

	return false
}

// SetArticleNumber gets a reference to the given string and assigns it to the ArticleNumber field.
func (o *SheetUsage) SetArticleNumber(v string) {
	o.ArticleNumber = &v
}

// GetQuantity returns the Quantity field value
func (o *SheetUsage) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *SheetUsage) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *SheetUsage) SetQuantity(v int32) {
	o.Quantity = v
}

// GetYieldPercent returns the YieldPercent field value
func (o *SheetUsage) GetYieldPercent() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.YieldPercent
}

// GetYieldPercentOk returns a tuple with the YieldPercent field value
// and a boolean to check if the value has been set.
func (o *SheetUsage) GetYieldPercentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YieldPercent, true
}

// SetYieldPercent sets field value
func (o *SheetUsage) SetYieldPercent(v float64) {
	o.YieldPercent = v
}

func (o SheetUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SheetUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sheet_id"] = o.SheetId
	if !IsNil(o.ArticleNumber) {
		toSerialize["article_number"] = o.ArticleNumber
	}
	toSerialize["quantity"] = o.Quantity
	toSerialize["yield_percent"] = o.YieldPercent
	return toSerialize, nil
}

func (o *SheetUsage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sheet_id",
		"quantity",
		"yield_percent",
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

	varSheetUsage := _SheetUsage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSheetUsage)

	if err != nil {
		return err
	}

	*o = SheetUsage(varSheetUsage)

	return err
}

type NullableSheetUsage struct {
	value *SheetUsage
	isSet bool
}

func (v NullableSheetUsage) Get() *SheetUsage {
	return v.value
}

func (v *NullableSheetUsage) Set(val *SheetUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableSheetUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableSheetUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSheetUsage(val *SheetUsage) *NullableSheetUsage {
	return &NullableSheetUsage{value: val, isSet: true}
}

func (v NullableSheetUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSheetUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


