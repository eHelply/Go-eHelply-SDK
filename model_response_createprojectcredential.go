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

// ResponseCreateprojectcredential struct for ResponseCreateprojectcredential
type ResponseCreateprojectcredential struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseCreateprojectcredential instantiates a new ResponseCreateprojectcredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateprojectcredential() *ResponseCreateprojectcredential {
	this := ResponseCreateprojectcredential{}
	return &this
}

// NewResponseCreateprojectcredentialWithDefaults instantiates a new ResponseCreateprojectcredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateprojectcredentialWithDefaults() *ResponseCreateprojectcredential {
	this := ResponseCreateprojectcredential{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseCreateprojectcredential) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateprojectcredential) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseCreateprojectcredential) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseCreateprojectcredential) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseCreateprojectcredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseCreateprojectcredential struct {
	value *ResponseCreateprojectcredential
	isSet bool
}

func (v NullableResponseCreateprojectcredential) Get() *ResponseCreateprojectcredential {
	return v.value
}

func (v *NullableResponseCreateprojectcredential) Set(val *ResponseCreateprojectcredential) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateprojectcredential) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateprojectcredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateprojectcredential(val *ResponseCreateprojectcredential) *NullableResponseCreateprojectcredential {
	return &NullableResponseCreateprojectcredential{value: val, isSet: true}
}

func (v NullableResponseCreateprojectcredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateprojectcredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


