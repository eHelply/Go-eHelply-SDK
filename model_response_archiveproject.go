/*
eHelply SDK - 1.1.106

eHelply SDK for SuperStack Services

API version: 1.1.106
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ResponseArchiveproject struct for ResponseArchiveproject
type ResponseArchiveproject struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseArchiveproject instantiates a new ResponseArchiveproject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseArchiveproject() *ResponseArchiveproject {
	this := ResponseArchiveproject{}
	return &this
}

// NewResponseArchiveprojectWithDefaults instantiates a new ResponseArchiveproject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseArchiveprojectWithDefaults() *ResponseArchiveproject {
	this := ResponseArchiveproject{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseArchiveproject) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseArchiveproject) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseArchiveproject) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseArchiveproject) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseArchiveproject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseArchiveproject struct {
	value *ResponseArchiveproject
	isSet bool
}

func (v NullableResponseArchiveproject) Get() *ResponseArchiveproject {
	return v.value
}

func (v *NullableResponseArchiveproject) Set(val *ResponseArchiveproject) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseArchiveproject) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseArchiveproject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseArchiveproject(val *ResponseArchiveproject) *NullableResponseArchiveproject {
	return &NullableResponseArchiveproject{value: val, isSet: true}
}

func (v NullableResponseArchiveproject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseArchiveproject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


