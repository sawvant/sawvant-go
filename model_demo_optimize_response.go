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

// checks if the DemoOptimizeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DemoOptimizeResponse{}

// DemoOptimizeResponse struct for DemoOptimizeResponse
type DemoOptimizeResponse struct {
	Layouts []DemoLayout `json:"layouts"`
	Summary Summary `json:"summary"`
	Metrics Metrics `json:"metrics"`
}

type _DemoOptimizeResponse DemoOptimizeResponse

// NewDemoOptimizeResponse instantiates a new DemoOptimizeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDemoOptimizeResponse(layouts []DemoLayout, summary Summary, metrics Metrics) *DemoOptimizeResponse {
	this := DemoOptimizeResponse{}
	this.Layouts = layouts
	this.Summary = summary
	this.Metrics = metrics
	return &this
}

// NewDemoOptimizeResponseWithDefaults instantiates a new DemoOptimizeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDemoOptimizeResponseWithDefaults() *DemoOptimizeResponse {
	this := DemoOptimizeResponse{}
	return &this
}

// GetLayouts returns the Layouts field value
func (o *DemoOptimizeResponse) GetLayouts() []DemoLayout {
	if o == nil {
		var ret []DemoLayout
		return ret
	}

	return o.Layouts
}

// GetLayoutsOk returns a tuple with the Layouts field value
// and a boolean to check if the value has been set.
func (o *DemoOptimizeResponse) GetLayoutsOk() ([]DemoLayout, bool) {
	if o == nil {
		return nil, false
	}
	return o.Layouts, true
}

// SetLayouts sets field value
func (o *DemoOptimizeResponse) SetLayouts(v []DemoLayout) {
	o.Layouts = v
}

// GetSummary returns the Summary field value
func (o *DemoOptimizeResponse) GetSummary() Summary {
	if o == nil {
		var ret Summary
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DemoOptimizeResponse) GetSummaryOk() (*Summary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *DemoOptimizeResponse) SetSummary(v Summary) {
	o.Summary = v
}

// GetMetrics returns the Metrics field value
func (o *DemoOptimizeResponse) GetMetrics() Metrics {
	if o == nil {
		var ret Metrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *DemoOptimizeResponse) GetMetricsOk() (*Metrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *DemoOptimizeResponse) SetMetrics(v Metrics) {
	o.Metrics = v
}

func (o DemoOptimizeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DemoOptimizeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["layouts"] = o.Layouts
	toSerialize["summary"] = o.Summary
	toSerialize["metrics"] = o.Metrics
	return toSerialize, nil
}

func (o *DemoOptimizeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"layouts",
		"summary",
		"metrics",
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

	varDemoOptimizeResponse := _DemoOptimizeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDemoOptimizeResponse)

	if err != nil {
		return err
	}

	*o = DemoOptimizeResponse(varDemoOptimizeResponse)

	return err
}

type NullableDemoOptimizeResponse struct {
	value *DemoOptimizeResponse
	isSet bool
}

func (v NullableDemoOptimizeResponse) Get() *DemoOptimizeResponse {
	return v.value
}

func (v *NullableDemoOptimizeResponse) Set(val *DemoOptimizeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDemoOptimizeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDemoOptimizeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDemoOptimizeResponse(val *DemoOptimizeResponse) *NullableDemoOptimizeResponse {
	return &NullableDemoOptimizeResponse{value: val, isSet: true}
}

func (v NullableDemoOptimizeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDemoOptimizeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


