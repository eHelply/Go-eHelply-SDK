/*
eHelply SDK - 1.1.107

eHelply SDK for SuperStack Services

API version: 1.1.107
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserPasswordReset Information used for resetting the password
type UserPasswordReset struct {
	Email string `json:"email"`
}

// NewUserPasswordReset instantiates a new UserPasswordReset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPasswordReset(email string) *UserPasswordReset {
	this := UserPasswordReset{}
	this.Email = email
	return &this
}

// NewUserPasswordResetWithDefaults instantiates a new UserPasswordReset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPasswordResetWithDefaults() *UserPasswordReset {
	this := UserPasswordReset{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserPasswordReset) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserPasswordReset) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserPasswordReset) SetEmail(v string) {
	o.Email = v
}

func (o UserPasswordReset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableUserPasswordReset struct {
	value *UserPasswordReset
	isSet bool
}

func (v NullableUserPasswordReset) Get() *UserPasswordReset {
	return v.value
}

func (v *NullableUserPasswordReset) Set(val *UserPasswordReset) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPasswordReset) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPasswordReset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPasswordReset(val *UserPasswordReset) *NullableUserPasswordReset {
	return &NullableUserPasswordReset{value: val, isSet: true}
}

func (v NullableUserPasswordReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPasswordReset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


