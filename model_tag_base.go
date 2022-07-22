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

// TagBase **:param** name                                **type:** string **:param** project_uuid                        **type:** string or None
type TagBase struct {
	Name string `json:"name"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

// NewTagBase instantiates a new TagBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagBase(name string) *TagBase {
	this := TagBase{}
	this.Name = name
	return &this
}

// NewTagBaseWithDefaults instantiates a new TagBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagBaseWithDefaults() *TagBase {
	this := TagBase{}
	return &this
}

// GetName returns the Name field value
func (o *TagBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagBase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagBase) SetName(v string) {
	o.Name = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *TagBase) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagBase) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *TagBase) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *TagBase) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

func (o TagBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	return json.Marshal(toSerialize)
}

type NullableTagBase struct {
	value *TagBase
	isSet bool
}

func (v NullableTagBase) Get() *TagBase {
	return v.value
}

func (v *NullableTagBase) Set(val *TagBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTagBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTagBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagBase(val *TagBase) *NullableTagBase {
	return &NullableTagBase{value: val, isSet: true}
}

func (v NullableTagBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


