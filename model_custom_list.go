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

// CustomList struct for CustomList
type CustomList struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewCustomList instantiates a new CustomList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomList() *CustomList {
	this := CustomList{}
	return &this
}

// NewCustomListWithDefaults instantiates a new CustomList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomListWithDefaults() *CustomList {
	this := CustomList{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomList) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomList) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomList) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomList) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomList) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomList) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomList) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomList) SetDescription(v string) {
	o.Description = &v
}

func (o CustomList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCustomList struct {
	value *CustomList
	isSet bool
}

func (v NullableCustomList) Get() *CustomList {
	return v.value
}

func (v *NullableCustomList) Set(val *CustomList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomList(val *CustomList) *NullableCustomList {
	return &NullableCustomList{value: val, isSet: true}
}

func (v NullableCustomList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


