/*
eHelply SDK - 1.1.106

eHelply SDK for SuperStack Services

API version: 1.1.106
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ContactMethod struct for ContactMethod
type ContactMethod struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewContactMethod instantiates a new ContactMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactMethod() *ContactMethod {
	this := ContactMethod{}
	var name string = ""
	this.Name = &name
	var value string = ""
	this.Value = &value
	return &this
}

// NewContactMethodWithDefaults instantiates a new ContactMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactMethodWithDefaults() *ContactMethod {
	this := ContactMethod{}
	var name string = ""
	this.Name = &name
	var value string = ""
	this.Value = &value
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContactMethod) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactMethod) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContactMethod) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContactMethod) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ContactMethod) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactMethod) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContactMethod) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ContactMethod) SetValue(v string) {
	o.Value = &v
}

func (o ContactMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableContactMethod struct {
	value *ContactMethod
	isSet bool
}

func (v NullableContactMethod) Get() *ContactMethod {
	return v.value
}

func (v *NullableContactMethod) Set(val *ContactMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableContactMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableContactMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactMethod(val *ContactMethod) *NullableContactMethod {
	return &NullableContactMethod{value: val, isSet: true}
}

func (v NullableContactMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


