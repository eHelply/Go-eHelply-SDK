/*
eHelply SDK - 1.1.88

eHelply SDK for SuperStack Services

API version: 1.1.88
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsProjectUpdate Used for update endpoint
type ProjectsProjectUpdate struct {
	MaxSpend *int32 `json:"max_spend,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewProjectsProjectUpdate instantiates a new ProjectsProjectUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsProjectUpdate() *ProjectsProjectUpdate {
	this := ProjectsProjectUpdate{}
	return &this
}

// NewProjectsProjectUpdateWithDefaults instantiates a new ProjectsProjectUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsProjectUpdateWithDefaults() *ProjectsProjectUpdate {
	this := ProjectsProjectUpdate{}
	return &this
}

// GetMaxSpend returns the MaxSpend field value if set, zero value otherwise.
func (o *ProjectsProjectUpdate) GetMaxSpend() int32 {
	if o == nil || o.MaxSpend == nil {
		var ret int32
		return ret
	}
	return *o.MaxSpend
}

// GetMaxSpendOk returns a tuple with the MaxSpend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUpdate) GetMaxSpendOk() (*int32, bool) {
	if o == nil || o.MaxSpend == nil {
		return nil, false
	}
	return o.MaxSpend, true
}

// HasMaxSpend returns a boolean if a field has been set.
func (o *ProjectsProjectUpdate) HasMaxSpend() bool {
	if o != nil && o.MaxSpend != nil {
		return true
	}

	return false
}

// SetMaxSpend gets a reference to the given int32 and assigns it to the MaxSpend field.
func (o *ProjectsProjectUpdate) SetMaxSpend(v int32) {
	o.MaxSpend = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProjectsProjectUpdate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUpdate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectsProjectUpdate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProjectsProjectUpdate) SetStatus(v string) {
	o.Status = &v
}

func (o ProjectsProjectUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxSpend != nil {
		toSerialize["max_spend"] = o.MaxSpend
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsProjectUpdate struct {
	value *ProjectsProjectUpdate
	isSet bool
}

func (v NullableProjectsProjectUpdate) Get() *ProjectsProjectUpdate {
	return v.value
}

func (v *NullableProjectsProjectUpdate) Set(val *ProjectsProjectUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsProjectUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsProjectUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsProjectUpdate(val *ProjectsProjectUpdate) *NullableProjectsProjectUpdate {
	return &NullableProjectsProjectUpdate{value: val, isSet: true}
}

func (v NullableProjectsProjectUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsProjectUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


