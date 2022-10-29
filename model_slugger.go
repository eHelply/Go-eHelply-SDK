/*
eHelply SDK - 1.1.115

eHelply SDK for SuperStack Services

API version: 1.1.115
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Slugger Meta slugger
type Slugger struct {
	Name string `json:"name"`
}

// NewSlugger instantiates a new Slugger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlugger(name string) *Slugger {
	this := Slugger{}
	this.Name = name
	return &this
}

// NewSluggerWithDefaults instantiates a new Slugger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSluggerWithDefaults() *Slugger {
	this := Slugger{}
	return &this
}

// GetName returns the Name field value
func (o *Slugger) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Slugger) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Slugger) SetName(v string) {
	o.Name = v
}

func (o Slugger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSlugger struct {
	value *Slugger
	isSet bool
}

func (v NullableSlugger) Get() *Slugger {
	return v.value
}

func (v *NullableSlugger) Set(val *Slugger) {
	v.value = val
	v.isSet = true
}

func (v NullableSlugger) IsSet() bool {
	return v.isSet
}

func (v *NullableSlugger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlugger(val *Slugger) *NullableSlugger {
	return &NullableSlugger{value: val, isSet: true}
}

func (v NullableSlugger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlugger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


