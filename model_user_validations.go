/*
eHelply SDK - 1.1.89

eHelply SDK for SuperStack Services

API version: 1.1.89
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserValidations Fields used to validate a user's field
type UserValidations struct {
	Value string `json:"value"`
	ValidationType string `json:"validation_type"`
}

// NewUserValidations instantiates a new UserValidations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserValidations(value string, validationType string) *UserValidations {
	this := UserValidations{}
	this.Value = value
	this.ValidationType = validationType
	return &this
}

// NewUserValidationsWithDefaults instantiates a new UserValidations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserValidationsWithDefaults() *UserValidations {
	this := UserValidations{}
	return &this
}

// GetValue returns the Value field value
func (o *UserValidations) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UserValidations) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UserValidations) SetValue(v string) {
	o.Value = v
}

// GetValidationType returns the ValidationType field value
func (o *UserValidations) GetValidationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationType
}

// GetValidationTypeOk returns a tuple with the ValidationType field value
// and a boolean to check if the value has been set.
func (o *UserValidations) GetValidationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationType, true
}

// SetValidationType sets field value
func (o *UserValidations) SetValidationType(v string) {
	o.ValidationType = v
}

func (o UserValidations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["validation_type"] = o.ValidationType
	}
	return json.Marshal(toSerialize)
}

type NullableUserValidations struct {
	value *UserValidations
	isSet bool
}

func (v NullableUserValidations) Get() *UserValidations {
	return v.value
}

func (v *NullableUserValidations) Set(val *UserValidations) {
	v.value = val
	v.isSet = true
}

func (v NullableUserValidations) IsSet() bool {
	return v.isSet
}

func (v *NullableUserValidations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserValidations(val *UserValidations) *NullableUserValidations {
	return &NullableUserValidations{value: val, isSet: true}
}

func (v NullableUserValidations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserValidations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


