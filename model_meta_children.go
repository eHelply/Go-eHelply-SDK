/*
eHelply SDK - 1.1.121

eHelply SDK for SuperStack Services

API version: 1.1.121
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// MetaChildren struct for MetaChildren
type MetaChildren struct {
	ChildName *string `json:"child_name,omitempty"`
	ChildDescription *string `json:"child_description,omitempty"`
	ChildUuid *string `json:"child_uuid,omitempty"`
}

// NewMetaChildren instantiates a new MetaChildren object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaChildren() *MetaChildren {
	this := MetaChildren{}
	return &this
}

// NewMetaChildrenWithDefaults instantiates a new MetaChildren object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaChildrenWithDefaults() *MetaChildren {
	this := MetaChildren{}
	return &this
}

// GetChildName returns the ChildName field value if set, zero value otherwise.
func (o *MetaChildren) GetChildName() string {
	if o == nil || o.ChildName == nil {
		var ret string
		return ret
	}
	return *o.ChildName
}

// GetChildNameOk returns a tuple with the ChildName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaChildren) GetChildNameOk() (*string, bool) {
	if o == nil || o.ChildName == nil {
		return nil, false
	}
	return o.ChildName, true
}

// HasChildName returns a boolean if a field has been set.
func (o *MetaChildren) HasChildName() bool {
	if o != nil && o.ChildName != nil {
		return true
	}

	return false
}

// SetChildName gets a reference to the given string and assigns it to the ChildName field.
func (o *MetaChildren) SetChildName(v string) {
	o.ChildName = &v
}

// GetChildDescription returns the ChildDescription field value if set, zero value otherwise.
func (o *MetaChildren) GetChildDescription() string {
	if o == nil || o.ChildDescription == nil {
		var ret string
		return ret
	}
	return *o.ChildDescription
}

// GetChildDescriptionOk returns a tuple with the ChildDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaChildren) GetChildDescriptionOk() (*string, bool) {
	if o == nil || o.ChildDescription == nil {
		return nil, false
	}
	return o.ChildDescription, true
}

// HasChildDescription returns a boolean if a field has been set.
func (o *MetaChildren) HasChildDescription() bool {
	if o != nil && o.ChildDescription != nil {
		return true
	}

	return false
}

// SetChildDescription gets a reference to the given string and assigns it to the ChildDescription field.
func (o *MetaChildren) SetChildDescription(v string) {
	o.ChildDescription = &v
}

// GetChildUuid returns the ChildUuid field value if set, zero value otherwise.
func (o *MetaChildren) GetChildUuid() string {
	if o == nil || o.ChildUuid == nil {
		var ret string
		return ret
	}
	return *o.ChildUuid
}

// GetChildUuidOk returns a tuple with the ChildUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaChildren) GetChildUuidOk() (*string, bool) {
	if o == nil || o.ChildUuid == nil {
		return nil, false
	}
	return o.ChildUuid, true
}

// HasChildUuid returns a boolean if a field has been set.
func (o *MetaChildren) HasChildUuid() bool {
	if o != nil && o.ChildUuid != nil {
		return true
	}

	return false
}

// SetChildUuid gets a reference to the given string and assigns it to the ChildUuid field.
func (o *MetaChildren) SetChildUuid(v string) {
	o.ChildUuid = &v
}

func (o MetaChildren) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChildName != nil {
		toSerialize["child_name"] = o.ChildName
	}
	if o.ChildDescription != nil {
		toSerialize["child_description"] = o.ChildDescription
	}
	if o.ChildUuid != nil {
		toSerialize["child_uuid"] = o.ChildUuid
	}
	return json.Marshal(toSerialize)
}

type NullableMetaChildren struct {
	value *MetaChildren
	isSet bool
}

func (v NullableMetaChildren) Get() *MetaChildren {
	return v.value
}

func (v *NullableMetaChildren) Set(val *MetaChildren) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaChildren) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaChildren) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaChildren(val *MetaChildren) *NullableMetaChildren {
	return &NullableMetaChildren{value: val, isSet: true}
}

func (v NullableMetaChildren) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaChildren) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


