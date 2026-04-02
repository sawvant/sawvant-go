/*
 * Sawvant Cutting Optimization API — Go SDK
 *
 * File generated from our OpenAPI spec; DO NOT EDIT.
 */

package sawvant

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the JobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobResponse{}

// JobResponse struct for JobResponse
type JobResponse struct {
	JobId string `json:"job_id"`
	Status string `json:"status"`
	Progress int32 `json:"progress"`
	CreatedAt time.Time `json:"created_at"`
	StartedAt *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Result *OptimizeResult `json:"result,omitempty"`
	// E.g. unplaced part IDs
	Warnings []string `json:"warnings,omitempty"`
	Error *string `json:"error,omitempty"`
}

type _JobResponse JobResponse

// NewJobResponse instantiates a new JobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobResponse(jobId string, status string, progress int32, createdAt time.Time) *JobResponse {
	this := JobResponse{}
	this.JobId = jobId
	this.Status = status
	this.Progress = progress
	this.CreatedAt = createdAt
	return &this
}

// NewJobResponseWithDefaults instantiates a new JobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobResponseWithDefaults() *JobResponse {
	this := JobResponse{}
	return &this
}

// GetJobId returns the JobId field value
func (o *JobResponse) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *JobResponse) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *JobResponse) SetJobId(v string) {
	o.JobId = v
}

// GetStatus returns the Status field value
func (o *JobResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *JobResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JobResponse) SetStatus(v string) {
	o.Status = v
}

// GetProgress returns the Progress field value
func (o *JobResponse) GetProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *JobResponse) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *JobResponse) SetProgress(v int32) {
	o.Progress = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *JobResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *JobResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *JobResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *JobResponse) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponse) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *JobResponse) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *JobResponse) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *JobResponse) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *JobResponse) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *JobResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *JobResponse) GetResult() OptimizeResult {
	if o == nil || IsNil(o.Result) {
		var ret OptimizeResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponse) GetResultOk() (*OptimizeResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *JobResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OptimizeResult and assigns it to the Result field.
func (o *JobResponse) SetResult(v OptimizeResult) {
	o.Result = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *JobResponse) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *JobResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *JobResponse) SetWarnings(v []string) {
	o.Warnings = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *JobResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *JobResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *JobResponse) SetError(v string) {
	o.Error = &v
}

func (o JobResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["job_id"] = o.JobId
	toSerialize["status"] = o.Status
	toSerialize["progress"] = o.Progress
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.CompletedAt) {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *JobResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"job_id",
		"status",
		"progress",
		"created_at",
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

	varJobResponse := _JobResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJobResponse)

	if err != nil {
		return err
	}

	*o = JobResponse(varJobResponse)

	return err
}

type NullableJobResponse struct {
	value *JobResponse
	isSet bool
}

func (v NullableJobResponse) Get() *JobResponse {
	return v.value
}

func (v *NullableJobResponse) Set(val *JobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResponse(val *JobResponse) *NullableJobResponse {
	return &NullableJobResponse{value: val, isSet: true}
}

func (v NullableJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


