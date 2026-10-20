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

// checks if the BatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchRequest{}

// BatchRequest struct for BatchRequest
type BatchRequest struct {
	// Individual optimization requests. At most 50 with strategy `fast` or 10 with strategy `thorough`. 
	Jobs []OptimizeRequest `json:"jobs"`
}

type _BatchRequest BatchRequest

// NewBatchRequest instantiates a new BatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRequest(jobs []OptimizeRequest) *BatchRequest {
	this := BatchRequest{}
	this.Jobs = jobs
	return &this
}

// NewBatchRequestWithDefaults instantiates a new BatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRequestWithDefaults() *BatchRequest {
	this := BatchRequest{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *BatchRequest) GetJobs() []OptimizeRequest {
	if o == nil {
		var ret []OptimizeRequest
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *BatchRequest) GetJobsOk() ([]OptimizeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *BatchRequest) SetJobs(v []OptimizeRequest) {
	o.Jobs = v
}

func (o BatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobs"] = o.Jobs
	return toSerialize, nil
}

func (o *BatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobs",
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

	varBatchRequest := _BatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchRequest)

	if err != nil {
		return err
	}

	*o = BatchRequest(varBatchRequest)

	return err
}

type NullableBatchRequest struct {
	value *BatchRequest
	isSet bool
}

func (v NullableBatchRequest) Get() *BatchRequest {
	return v.value
}

func (v *NullableBatchRequest) Set(val *BatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRequest(val *BatchRequest) *NullableBatchRequest {
	return &NullableBatchRequest{value: val, isSet: true}
}

func (v NullableBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


