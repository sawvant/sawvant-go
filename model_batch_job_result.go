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

// checks if the BatchJobResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchJobResult{}

// BatchJobResult struct for BatchJobResult
type BatchJobResult struct {
	// Position of this job in the request array
	Index int32 `json:"index"`
	JobId *string `json:"job_id,omitempty"`
	Error *string `json:"error,omitempty"`
}

type _BatchJobResult BatchJobResult

// NewBatchJobResult instantiates a new BatchJobResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchJobResult(index int32) *BatchJobResult {
	this := BatchJobResult{}
	this.Index = index
	return &this
}

// NewBatchJobResultWithDefaults instantiates a new BatchJobResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchJobResultWithDefaults() *BatchJobResult {
	this := BatchJobResult{}
	return &this
}

// GetIndex returns the Index field value
func (o *BatchJobResult) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *BatchJobResult) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *BatchJobResult) SetIndex(v int32) {
	o.Index = v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *BatchJobResult) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchJobResult) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *BatchJobResult) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *BatchJobResult) SetJobId(v string) {
	o.JobId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BatchJobResult) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchJobResult) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BatchJobResult) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BatchJobResult) SetError(v string) {
	o.Error = &v
}

func (o BatchJobResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchJobResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *BatchJobResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
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

	varBatchJobResult := _BatchJobResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchJobResult)

	if err != nil {
		return err
	}

	*o = BatchJobResult(varBatchJobResult)

	return err
}

type NullableBatchJobResult struct {
	value *BatchJobResult
	isSet bool
}

func (v NullableBatchJobResult) Get() *BatchJobResult {
	return v.value
}

func (v *NullableBatchJobResult) Set(val *BatchJobResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchJobResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchJobResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchJobResult(val *BatchJobResult) *NullableBatchJobResult {
	return &NullableBatchJobResult{value: val, isSet: true}
}

func (v NullableBatchJobResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchJobResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


