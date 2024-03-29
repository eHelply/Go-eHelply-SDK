/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsProjectGet Used for get endpoint
type ProjectsProjectGet struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Status *string `json:"status,omitempty"`
	ArchivedAt *string `json:"archived_at,omitempty"`
}

// NewProjectsProjectGet instantiates a new ProjectsProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsProjectGet(uuid string, name string) *ProjectsProjectGet {
	this := ProjectsProjectGet{}
	this.Uuid = uuid
	this.Name = name
	return &this
}

// NewProjectsProjectGetWithDefaults instantiates a new ProjectsProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsProjectGetWithDefaults() *ProjectsProjectGet {
	this := ProjectsProjectGet{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ProjectsProjectGet) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectGet) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ProjectsProjectGet) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *ProjectsProjectGet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectsProjectGet) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProjectsProjectGet) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsProjectGet) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectsProjectGet) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProjectsProjectGet) SetStatus(v string) {
	o.Status = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *ProjectsProjectGet) GetArchivedAt() string {
	if o == nil || o.ArchivedAt == nil {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsProjectGet) GetArchivedAtOk() (*string, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *ProjectsProjectGet) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *ProjectsProjectGet) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

func (o ProjectsProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ArchivedAt != nil {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsProjectGet struct {
	value *ProjectsProjectGet
	isSet bool
}

func (v NullableProjectsProjectGet) Get() *ProjectsProjectGet {
	return v.value
}

func (v *NullableProjectsProjectGet) Set(val *ProjectsProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsProjectGet(val *ProjectsProjectGet) *NullableProjectsProjectGet {
	return &NullableProjectsProjectGet{value: val, isSet: true}
}

func (v NullableProjectsProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


