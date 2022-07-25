/*
eHelply SDK - 1.1.93

eHelply SDK for SuperStack Services

API version: 1.1.93
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserPasswordResetConfirmation Information for resetting the new password with a confirmation code
type UserPasswordResetConfirmation struct {
	Email string `json:"email"`
	ConfirmationCode string `json:"confirmation_code"`
	Password string `json:"password"`
}

// NewUserPasswordResetConfirmation instantiates a new UserPasswordResetConfirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPasswordResetConfirmation(email string, confirmationCode string, password string) *UserPasswordResetConfirmation {
	this := UserPasswordResetConfirmation{}
	this.Email = email
	this.ConfirmationCode = confirmationCode
	this.Password = password
	return &this
}

// NewUserPasswordResetConfirmationWithDefaults instantiates a new UserPasswordResetConfirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPasswordResetConfirmationWithDefaults() *UserPasswordResetConfirmation {
	this := UserPasswordResetConfirmation{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserPasswordResetConfirmation) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserPasswordResetConfirmation) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserPasswordResetConfirmation) SetEmail(v string) {
	o.Email = v
}

// GetConfirmationCode returns the ConfirmationCode field value
func (o *UserPasswordResetConfirmation) GetConfirmationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmationCode
}

// GetConfirmationCodeOk returns a tuple with the ConfirmationCode field value
// and a boolean to check if the value has been set.
func (o *UserPasswordResetConfirmation) GetConfirmationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmationCode, true
}

// SetConfirmationCode sets field value
func (o *UserPasswordResetConfirmation) SetConfirmationCode(v string) {
	o.ConfirmationCode = v
}

// GetPassword returns the Password field value
func (o *UserPasswordResetConfirmation) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserPasswordResetConfirmation) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserPasswordResetConfirmation) SetPassword(v string) {
	o.Password = v
}

func (o UserPasswordResetConfirmation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["confirmation_code"] = o.ConfirmationCode
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserPasswordResetConfirmation struct {
	value *UserPasswordResetConfirmation
	isSet bool
}

func (v NullableUserPasswordResetConfirmation) Get() *UserPasswordResetConfirmation {
	return v.value
}

func (v *NullableUserPasswordResetConfirmation) Set(val *UserPasswordResetConfirmation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPasswordResetConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPasswordResetConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPasswordResetConfirmation(val *UserPasswordResetConfirmation) *NullableUserPasswordResetConfirmation {
	return &NullableUserPasswordResetConfirmation{value: val, isSet: true}
}

func (v NullableUserPasswordResetConfirmation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPasswordResetConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


