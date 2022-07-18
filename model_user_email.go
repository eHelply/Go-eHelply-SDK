/*
eHelply SDK - 1.1.87

eHelply SDK for SuperStack Services

API version: 1.1.87
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserEmail User email information
type UserEmail struct {
	Address *string `json:"address,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}

// NewUserEmail instantiates a new UserEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEmail() *UserEmail {
	this := UserEmail{}
	var verified bool = false
	this.Verified = &verified
	return &this
}

// NewUserEmailWithDefaults instantiates a new UserEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEmailWithDefaults() *UserEmail {
	this := UserEmail{}
	var verified bool = false
	this.Verified = &verified
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UserEmail) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEmail) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UserEmail) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UserEmail) SetAddress(v string) {
	o.Address = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *UserEmail) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEmail) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *UserEmail) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *UserEmail) SetVerified(v bool) {
	o.Verified = &v
}

func (o UserEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}
	return json.Marshal(toSerialize)
}

type NullableUserEmail struct {
	value *UserEmail
	isSet bool
}

func (v NullableUserEmail) Get() *UserEmail {
	return v.value
}

func (v *NullableUserEmail) Set(val *UserEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEmail(val *UserEmail) *NullableUserEmail {
	return &NullableUserEmail{value: val, isSet: true}
}

func (v NullableUserEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


