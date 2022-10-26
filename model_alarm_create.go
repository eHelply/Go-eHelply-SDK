/*
eHelply SDK - 1.1.117

eHelply SDK for SuperStack Services

API version: 1.1.117
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AlarmCreate struct for AlarmCreate
type AlarmCreate struct {
	Process *string `json:"process,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Name *string `json:"name,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewAlarmCreate instantiates a new AlarmCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmCreate() *AlarmCreate {
	this := AlarmCreate{}
	return &this
}

// NewAlarmCreateWithDefaults instantiates a new AlarmCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmCreateWithDefaults() *AlarmCreate {
	this := AlarmCreate{}
	return &this
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *AlarmCreate) GetProcess() string {
	if o == nil || o.Process == nil {
		var ret string
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmCreate) GetProcessOk() (*string, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *AlarmCreate) HasProcess() bool {
	if o != nil && o.Process != nil {
		return true
	}

	return false
}

// SetProcess gets a reference to the given string and assigns it to the Process field.
func (o *AlarmCreate) SetProcess(v string) {
	o.Process = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AlarmCreate) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmCreate) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AlarmCreate) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AlarmCreate) SetSeverity(v string) {
	o.Severity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlarmCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlarmCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlarmCreate) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AlarmCreate) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmCreate) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AlarmCreate) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AlarmCreate) SetSummary(v string) {
	o.Summary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlarmCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlarmCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlarmCreate) SetDescription(v string) {
	o.Description = &v
}

func (o AlarmCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmCreate struct {
	value *AlarmCreate
	isSet bool
}

func (v NullableAlarmCreate) Get() *AlarmCreate {
	return v.value
}

func (v *NullableAlarmCreate) Set(val *AlarmCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmCreate(val *AlarmCreate) *NullableAlarmCreate {
	return &NullableAlarmCreate{value: val, isSet: true}
}

func (v NullableAlarmCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


