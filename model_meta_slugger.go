/*
eHelply SDK - 1.1.92

eHelply SDK for SuperStack Services

API version: 1.1.92
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// MetaSlugger Meta slugger
type MetaSlugger struct {
	Name string `json:"name"`
}

// NewMetaSlugger instantiates a new MetaSlugger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaSlugger(name string) *MetaSlugger {
	this := MetaSlugger{}
	this.Name = name
	return &this
}

// NewMetaSluggerWithDefaults instantiates a new MetaSlugger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaSluggerWithDefaults() *MetaSlugger {
	this := MetaSlugger{}
	return &this
}

// GetName returns the Name field value
func (o *MetaSlugger) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetaSlugger) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetaSlugger) SetName(v string) {
	o.Name = v
}

func (o MetaSlugger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableMetaSlugger struct {
	value *MetaSlugger
	isSet bool
}

func (v NullableMetaSlugger) Get() *MetaSlugger {
	return v.value
}

func (v *NullableMetaSlugger) Set(val *MetaSlugger) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaSlugger) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaSlugger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaSlugger(val *MetaSlugger) *NullableMetaSlugger {
	return &NullableMetaSlugger{value: val, isSet: true}
}

func (v NullableMetaSlugger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaSlugger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


