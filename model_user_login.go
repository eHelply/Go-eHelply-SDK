/*
eHelply SDK - 1.1.119

eHelply SDK for SuperStack Services

API version: 1.1.119
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserLogin Password and username required to login
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUserLogin instantiates a new UserLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLogin(username string, password string) *UserLogin {
	this := UserLogin{}
	this.Username = username
	this.Password = password
	return &this
}

// NewUserLoginWithDefaults instantiates a new UserLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginWithDefaults() *UserLogin {
	this := UserLogin{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserLogin) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserLogin) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *UserLogin) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserLogin) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserLogin) SetPassword(v string) {
	o.Password = v
}

func (o UserLogin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserLogin struct {
	value *UserLogin
	isSet bool
}

func (v NullableUserLogin) Get() *UserLogin {
	return v.value
}

func (v *NullableUserLogin) Set(val *UserLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLogin(val *UserLogin) *NullableUserLogin {
	return &NullableUserLogin{value: val, isSet: true}
}

func (v NullableUserLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


