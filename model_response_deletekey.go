/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ResponseDeletekey struct for ResponseDeletekey
type ResponseDeletekey struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseDeletekey instantiates a new ResponseDeletekey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeletekey() *ResponseDeletekey {
	this := ResponseDeletekey{}
	return &this
}

// NewResponseDeletekeyWithDefaults instantiates a new ResponseDeletekey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeletekeyWithDefaults() *ResponseDeletekey {
	this := ResponseDeletekey{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseDeletekey) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeletekey) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseDeletekey) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseDeletekey) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseDeletekey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseDeletekey struct {
	value *ResponseDeletekey
	isSet bool
}

func (v NullableResponseDeletekey) Get() *ResponseDeletekey {
	return v.value
}

func (v *NullableResponseDeletekey) Set(val *ResponseDeletekey) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeletekey) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeletekey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeletekey(val *ResponseDeletekey) *NullableResponseDeletekey {
	return &NullableResponseDeletekey{value: val, isSet: true}
}

func (v NullableResponseDeletekey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeletekey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


