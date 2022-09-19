/*
eHelply SDK - 1.1.114

eHelply SDK for SuperStack Services

API version: 1.1.114
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Validations Validations
type Validations struct {
	Value []string `json:"value,omitempty"`
}

// NewValidations instantiates a new Validations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidations() *Validations {
	this := Validations{}
	return &this
}

// NewValidationsWithDefaults instantiates a new Validations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationsWithDefaults() *Validations {
	this := Validations{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Validations) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Validations) GetValueOk() ([]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Validations) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *Validations) SetValue(v []string) {
	o.Value = v
}

func (o Validations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableValidations struct {
	value *Validations
	isSet bool
}

func (v NullableValidations) Get() *Validations {
	return v.value
}

func (v *NullableValidations) Set(val *Validations) {
	v.value = val
	v.isSet = true
}

func (v NullableValidations) IsSet() bool {
	return v.isSet
}

func (v *NullableValidations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidations(val *Validations) *NullableValidations {
	return &NullableValidations{value: val, isSet: true}
}

func (v NullableValidations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


