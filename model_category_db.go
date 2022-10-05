/*
eHelply SDK - 1.1.110

eHelply SDK for SuperStack Services

API version: 1.1.110
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CategoryDb **:param** uuid                                **type:** string **:param** name                                **type:** string **:param** project_uuid                        **type:** string or None  **:param** meta_uuid                           **type:** string or None
type CategoryDb struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
}

// NewCategoryDb instantiates a new CategoryDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryDb(uuid string, name string) *CategoryDb {
	this := CategoryDb{}
	this.Uuid = uuid
	this.Name = name
	return &this
}

// NewCategoryDbWithDefaults instantiates a new CategoryDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryDbWithDefaults() *CategoryDb {
	this := CategoryDb{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *CategoryDb) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CategoryDb) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CategoryDb) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *CategoryDb) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryDb) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryDb) SetName(v string) {
	o.Name = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *CategoryDb) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryDb) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *CategoryDb) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *CategoryDb) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *CategoryDb) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryDb) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *CategoryDb) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *CategoryDb) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

func (o CategoryDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.MetaUuid != nil {
		toSerialize["meta_uuid"] = o.MetaUuid
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryDb struct {
	value *CategoryDb
	isSet bool
}

func (v NullableCategoryDb) Get() *CategoryDb {
	return v.value
}

func (v *NullableCategoryDb) Set(val *CategoryDb) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryDb) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryDb(val *CategoryDb) *NullableCategoryDb {
	return &NullableCategoryDb{value: val, isSet: true}
}

func (v NullableCategoryDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


