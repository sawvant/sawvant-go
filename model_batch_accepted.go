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

// checks if the BatchAccepted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchAccepted{}

// BatchAccepted struct for BatchAccepted
type BatchAccepted struct {
	BatchId string `json:"batch_id"`
	JobIds []string `json:"job_ids"`
	// Number of jobs in the request
	Total int32 `json:"total"`
	// Number of jobs successfully enqueued
	Created int32 `json:"created"`
	// Present only when one or more jobs failed to enqueue
	Results []BatchJobResult `json:"results,omitempty"`
}

type _BatchAccepted BatchAccepted

// NewBatchAccepted instantiates a new BatchAccepted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchAccepted(batchId string, jobIds []string, total int32, created int32) *BatchAccepted {
	this := BatchAccepted{}
	this.BatchId = batchId
	this.JobIds = jobIds
	this.Total = total
	this.Created = created
	return &this
}

// NewBatchAcceptedWithDefaults instantiates a new BatchAccepted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchAcceptedWithDefaults() *BatchAccepted {
	this := BatchAccepted{}
	return &this
}

// GetBatchId returns the BatchId field value
func (o *BatchAccepted) GetBatchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value
// and a boolean to check if the value has been set.
func (o *BatchAccepted) GetBatchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchId, true
}

// SetBatchId sets field value
func (o *BatchAccepted) SetBatchId(v string) {
	o.BatchId = v
}

// GetJobIds returns the JobIds field value
func (o *BatchAccepted) GetJobIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.JobIds
}

// GetJobIdsOk returns a tuple with the JobIds field value
// and a boolean to check if the value has been set.
func (o *BatchAccepted) GetJobIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobIds, true
}

// SetJobIds sets field value
func (o *BatchAccepted) SetJobIds(v []string) {
	o.JobIds = v
}

// GetTotal returns the Total field value
func (o *BatchAccepted) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *BatchAccepted) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *BatchAccepted) SetTotal(v int32) {
	o.Total = v
}

// GetCreated returns the Created field value
func (o *BatchAccepted) GetCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *BatchAccepted) GetCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *BatchAccepted) SetCreated(v int32) {
	o.Created = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BatchAccepted) GetResults() []BatchJobResult {
	if o == nil || IsNil(o.Results) {
		var ret []BatchJobResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchAccepted) GetResultsOk() ([]BatchJobResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BatchAccepted) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BatchJobResult and assigns it to the Results field.
func (o *BatchAccepted) SetResults(v []BatchJobResult) {
	o.Results = v
}

func (o BatchAccepted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchAccepted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["batch_id"] = o.BatchId
	toSerialize["job_ids"] = o.JobIds
	toSerialize["total"] = o.Total
	toSerialize["created"] = o.Created
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

func (o *BatchAccepted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"batch_id",
		"job_ids",
		"total",
		"created",
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

	varBatchAccepted := _BatchAccepted{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchAccepted)

	if err != nil {
		return err
	}

	*o = BatchAccepted(varBatchAccepted)

	return err
}

type NullableBatchAccepted struct {
	value *BatchAccepted
	isSet bool
}

func (v NullableBatchAccepted) Get() *BatchAccepted {
	return v.value
}

func (v *NullableBatchAccepted) Set(val *BatchAccepted) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchAccepted) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchAccepted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchAccepted(val *BatchAccepted) *NullableBatchAccepted {
	return &NullableBatchAccepted{value: val, isSet: true}
}

func (v NullableBatchAccepted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchAccepted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


