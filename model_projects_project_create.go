/*
eHelply SDK - 1.1.116

eHelply SDK for SuperStack Services

API version: 1.1.116
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsProjectCreate Used for create endpoint
type ProjectsProjectCreate struct {
	Name string `json:"name"`
}

// NewProjectsProjectCreate instantiates a new ProjectsProjectCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsProjectCreate(name string) *ProjectsProjectCreate {
	this := ProjectsProjectCreate{}
	this.Name = name
	return &this
}

// NewProjectsProjectCreateWithDefaults instantiates a new ProjectsProjectCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsProjectCreateWithDefaults() *ProjectsProjectCreate {
	this := ProjectsProjectCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectsProjectCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectsProjectCreate) SetName(v string) {
	o.Name = v
}

func (o ProjectsProjectCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsProjectCreate struct {
	value *ProjectsProjectCreate
	isSet bool
}

func (v NullableProjectsProjectCreate) Get() *ProjectsProjectCreate {
	return v.value
}

func (v *NullableProjectsProjectCreate) Set(val *ProjectsProjectCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsProjectCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsProjectCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsProjectCreate(val *ProjectsProjectCreate) *NullableProjectsProjectCreate {
	return &NullableProjectsProjectCreate{value: val, isSet: true}
}

func (v NullableProjectsProjectCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsProjectCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


