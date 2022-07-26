/*
eHelply SDK - 1.1.94

eHelply SDK for SuperStack Services

API version: 1.1.94
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserConfirmation Information to confirm user
type UserConfirmation struct {
	Email string `json:"email"`
	VerificationCode string `json:"verification_code"`
}

// NewUserConfirmation instantiates a new UserConfirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserConfirmation(email string, verificationCode string) *UserConfirmation {
	this := UserConfirmation{}
	this.Email = email
	this.VerificationCode = verificationCode
	return &this
}

// NewUserConfirmationWithDefaults instantiates a new UserConfirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConfirmationWithDefaults() *UserConfirmation {
	this := UserConfirmation{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserConfirmation) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserConfirmation) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserConfirmation) SetEmail(v string) {
	o.Email = v
}

// GetVerificationCode returns the VerificationCode field value
func (o *UserConfirmation) GetVerificationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value
// and a boolean to check if the value has been set.
func (o *UserConfirmation) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationCode, true
}

// SetVerificationCode sets field value
func (o *UserConfirmation) SetVerificationCode(v string) {
	o.VerificationCode = v
}

func (o UserConfirmation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["verification_code"] = o.VerificationCode
	}
	return json.Marshal(toSerialize)
}

type NullableUserConfirmation struct {
	value *UserConfirmation
	isSet bool
}

func (v NullableUserConfirmation) Get() *UserConfirmation {
	return v.value
}

func (v *NullableUserConfirmation) Set(val *UserConfirmation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConfirmation(val *UserConfirmation) *NullableUserConfirmation {
	return &NullableUserConfirmation{value: val, isSet: true}
}

func (v NullableUserConfirmation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


