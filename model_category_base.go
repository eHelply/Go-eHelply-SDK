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

// CategoryBase **:param** name                                **type:** string **:param** meta                                **type:** dict or None  **:param** project_uuid                        **type:** string or None
type CategoryBase struct {
	Name string `json:"name"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

// NewCategoryBase instantiates a new CategoryBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryBase(name string) *CategoryBase {
	this := CategoryBase{}
	this.Name = name
	return &this
}

// NewCategoryBaseWithDefaults instantiates a new CategoryBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryBaseWithDefaults() *CategoryBase {
	this := CategoryBase{}
	return &this
}

// GetName returns the Name field value
func (o *CategoryBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryBase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryBase) SetName(v string) {
	o.Name = v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *CategoryBase) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryBase) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *CategoryBase) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *CategoryBase) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CategoryBase) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryBase) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CategoryBase) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *CategoryBase) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *CategoryBase) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryBase) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *CategoryBase) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *CategoryBase) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

func (o CategoryBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.MetaUuid != nil {
		toSerialize["meta_uuid"] = o.MetaUuid
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryBase struct {
	value *CategoryBase
	isSet bool
}

func (v NullableCategoryBase) Get() *CategoryBase {
	return v.value
}

func (v *NullableCategoryBase) Set(val *CategoryBase) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryBase) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryBase(val *CategoryBase) *NullableCategoryBase {
	return &NullableCategoryBase{value: val, isSet: true}
}

func (v NullableCategoryBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


