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

// checks if the OptimizeAccepted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizeAccepted{}

// OptimizeAccepted struct for OptimizeAccepted
type OptimizeAccepted struct {
	JobId string `json:"job_id"`
	Status string `json:"status"`
	PollUrl string `json:"poll_url"`
	StreamUrl string `json:"stream_url"`
}

type _OptimizeAccepted OptimizeAccepted

// NewOptimizeAccepted instantiates a new OptimizeAccepted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizeAccepted(jobId string, status string, pollUrl string, streamUrl string) *OptimizeAccepted {
	this := OptimizeAccepted{}
	this.JobId = jobId
	this.Status = status
	this.PollUrl = pollUrl
	this.StreamUrl = streamUrl
	return &this
}

// NewOptimizeAcceptedWithDefaults instantiates a new OptimizeAccepted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizeAcceptedWithDefaults() *OptimizeAccepted {
	this := OptimizeAccepted{}
	return &this
}

// GetJobId returns the JobId field value
func (o *OptimizeAccepted) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *OptimizeAccepted) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *OptimizeAccepted) SetJobId(v string) {
	o.JobId = v
}

// GetStatus returns the Status field value
func (o *OptimizeAccepted) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OptimizeAccepted) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OptimizeAccepted) SetStatus(v string) {
	o.Status = v
}

// GetPollUrl returns the PollUrl field value
func (o *OptimizeAccepted) GetPollUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PollUrl
}

// GetPollUrlOk returns a tuple with the PollUrl field value
// and a boolean to check if the value has been set.
func (o *OptimizeAccepted) GetPollUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PollUrl, true
}

// SetPollUrl sets field value
func (o *OptimizeAccepted) SetPollUrl(v string) {
	o.PollUrl = v
}

// GetStreamUrl returns the StreamUrl field value
func (o *OptimizeAccepted) GetStreamUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamUrl
}

// GetStreamUrlOk returns a tuple with the StreamUrl field value
// and a boolean to check if the value has been set.
func (o *OptimizeAccepted) GetStreamUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamUrl, true
}

// SetStreamUrl sets field value
func (o *OptimizeAccepted) SetStreamUrl(v string) {
	o.StreamUrl = v
}

func (o OptimizeAccepted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptimizeAccepted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job_id"] = o.JobId
	toSerialize["status"] = o.Status
	toSerialize["poll_url"] = o.PollUrl
	toSerialize["stream_url"] = o.StreamUrl
	return toSerialize, nil
}

func (o *OptimizeAccepted) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"job_id",
		"status",
		"poll_url",
		"stream_url",
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

	varOptimizeAccepted := _OptimizeAccepted{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOptimizeAccepted)

	if err != nil {
		return err
	}

	*o = OptimizeAccepted(varOptimizeAccepted)

	return err
}

type NullableOptimizeAccepted struct {
	value *OptimizeAccepted
	isSet bool
}

func (v NullableOptimizeAccepted) Get() *OptimizeAccepted {
	return v.value
}

func (v *NullableOptimizeAccepted) Set(val *OptimizeAccepted) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizeAccepted) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizeAccepted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizeAccepted(val *OptimizeAccepted) *NullableOptimizeAccepted {
	return &NullableOptimizeAccepted{value: val, isSet: true}
}

func (v NullableOptimizeAccepted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptimizeAccepted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


