/*
eHelply SDK - 1.1.113

eHelply SDK for SuperStack Services

API version: 1.1.113
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ResponseUpdateprojectcredential struct for ResponseUpdateprojectcredential
type ResponseUpdateprojectcredential struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseUpdateprojectcredential instantiates a new ResponseUpdateprojectcredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseUpdateprojectcredential() *ResponseUpdateprojectcredential {
	this := ResponseUpdateprojectcredential{}
	return &this
}

// NewResponseUpdateprojectcredentialWithDefaults instantiates a new ResponseUpdateprojectcredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseUpdateprojectcredentialWithDefaults() *ResponseUpdateprojectcredential {
	this := ResponseUpdateprojectcredential{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseUpdateprojectcredential) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseUpdateprojectcredential) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseUpdateprojectcredential) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseUpdateprojectcredential) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseUpdateprojectcredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseUpdateprojectcredential struct {
	value *ResponseUpdateprojectcredential
	isSet bool
}

func (v NullableResponseUpdateprojectcredential) Get() *ResponseUpdateprojectcredential {
	return v.value
}

func (v *NullableResponseUpdateprojectcredential) Set(val *ResponseUpdateprojectcredential) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseUpdateprojectcredential) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseUpdateprojectcredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseUpdateprojectcredential(val *ResponseUpdateprojectcredential) *NullableResponseUpdateprojectcredential {
	return &NullableResponseUpdateprojectcredential{value: val, isSet: true}
}

func (v NullableResponseUpdateprojectcredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseUpdateprojectcredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


