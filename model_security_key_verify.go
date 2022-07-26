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

// SecurityKeyVerify Used for verify key endpoint
type SecurityKeyVerify struct {
	Access string `json:"access"`
	Secret string `json:"secret"`
}

// NewSecurityKeyVerify instantiates a new SecurityKeyVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityKeyVerify(access string, secret string) *SecurityKeyVerify {
	this := SecurityKeyVerify{}
	this.Access = access
	this.Secret = secret
	return &this
}

// NewSecurityKeyVerifyWithDefaults instantiates a new SecurityKeyVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityKeyVerifyWithDefaults() *SecurityKeyVerify {
	this := SecurityKeyVerify{}
	return &this
}

// GetAccess returns the Access field value
func (o *SecurityKeyVerify) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyVerify) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *SecurityKeyVerify) SetAccess(v string) {
	o.Access = v
}

// GetSecret returns the Secret field value
func (o *SecurityKeyVerify) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyVerify) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *SecurityKeyVerify) SetSecret(v string) {
	o.Secret = v
}

func (o SecurityKeyVerify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access"] = o.Access
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityKeyVerify struct {
	value *SecurityKeyVerify
	isSet bool
}

func (v NullableSecurityKeyVerify) Get() *SecurityKeyVerify {
	return v.value
}

func (v *NullableSecurityKeyVerify) Set(val *SecurityKeyVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityKeyVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityKeyVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityKeyVerify(val *SecurityKeyVerify) *NullableSecurityKeyVerify {
	return &NullableSecurityKeyVerify{value: val, isSet: true}
}

func (v NullableSecurityKeyVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityKeyVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


