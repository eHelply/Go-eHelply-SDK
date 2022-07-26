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

// GetSecret struct for GetSecret
type GetSecret struct {
	Name string `json:"name"`
}

// NewGetSecret instantiates a new GetSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSecret(name string) *GetSecret {
	this := GetSecret{}
	this.Name = name
	return &this
}

// NewGetSecretWithDefaults instantiates a new GetSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSecretWithDefaults() *GetSecret {
	this := GetSecret{}
	return &this
}

// GetName returns the Name field value
func (o *GetSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetSecret) SetName(v string) {
	o.Name = v
}

func (o GetSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGetSecret struct {
	value *GetSecret
	isSet bool
}

func (v NullableGetSecret) Get() *GetSecret {
	return v.value
}

func (v *NullableGetSecret) Set(val *GetSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSecret(val *GetSecret) *NullableGetSecret {
	return &NullableGetSecret{value: val, isSet: true}
}

func (v NullableGetSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


