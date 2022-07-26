/*
eHelply SDK - 1.1.102

eHelply SDK for SuperStack Services

API version: 1.1.102
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// SecurityCreateToken struct for SecurityCreateToken
type SecurityCreateToken struct {
	Length *int32 `json:"length,omitempty"`
}

// NewSecurityCreateToken instantiates a new SecurityCreateToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityCreateToken() *SecurityCreateToken {
	this := SecurityCreateToken{}
	var length int32 = 64
	this.Length = &length
	return &this
}

// NewSecurityCreateTokenWithDefaults instantiates a new SecurityCreateToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityCreateTokenWithDefaults() *SecurityCreateToken {
	this := SecurityCreateToken{}
	var length int32 = 64
	this.Length = &length
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *SecurityCreateToken) GetLength() int32 {
	if o == nil || o.Length == nil {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityCreateToken) GetLengthOk() (*int32, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *SecurityCreateToken) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *SecurityCreateToken) SetLength(v int32) {
	o.Length = &v
}

func (o SecurityCreateToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityCreateToken struct {
	value *SecurityCreateToken
	isSet bool
}

func (v NullableSecurityCreateToken) Get() *SecurityCreateToken {
	return v.value
}

func (v *NullableSecurityCreateToken) Set(val *SecurityCreateToken) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityCreateToken) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityCreateToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityCreateToken(val *SecurityCreateToken) *NullableSecurityCreateToken {
	return &NullableSecurityCreateToken{value: val, isSet: true}
}

func (v NullableSecurityCreateToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityCreateToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


