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

// MetaCustom struct for MetaCustom
type MetaCustom struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	List []CustomList `json:"list,omitempty"`
}

// NewMetaCustom instantiates a new MetaCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaCustom() *MetaCustom {
	this := MetaCustom{}
	return &this
}

// NewMetaCustomWithDefaults instantiates a new MetaCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaCustomWithDefaults() *MetaCustom {
	this := MetaCustom{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaCustom) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCustom) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaCustom) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaCustom) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetaCustom) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCustom) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetaCustom) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetaCustom) SetDescription(v string) {
	o.Description = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *MetaCustom) GetList() []CustomList {
	if o == nil || o.List == nil {
		var ret []CustomList
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCustom) GetListOk() ([]CustomList, bool) {
	if o == nil || o.List == nil {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *MetaCustom) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given []CustomList and assigns it to the List field.
func (o *MetaCustom) SetList(v []CustomList) {
	o.List = v
}

func (o MetaCustom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	return json.Marshal(toSerialize)
}

type NullableMetaCustom struct {
	value *MetaCustom
	isSet bool
}

func (v NullableMetaCustom) Get() *MetaCustom {
	return v.value
}

func (v *NullableMetaCustom) Set(val *MetaCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaCustom(val *MetaCustom) *NullableMetaCustom {
	return &NullableMetaCustom{value: val, isSet: true}
}

func (v NullableMetaCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


