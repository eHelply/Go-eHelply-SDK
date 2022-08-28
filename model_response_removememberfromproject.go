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

// ResponseRemovememberfromproject struct for ResponseRemovememberfromproject
type ResponseRemovememberfromproject struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseRemovememberfromproject instantiates a new ResponseRemovememberfromproject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRemovememberfromproject() *ResponseRemovememberfromproject {
	this := ResponseRemovememberfromproject{}
	return &this
}

// NewResponseRemovememberfromprojectWithDefaults instantiates a new ResponseRemovememberfromproject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRemovememberfromprojectWithDefaults() *ResponseRemovememberfromproject {
	this := ResponseRemovememberfromproject{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseRemovememberfromproject) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseRemovememberfromproject) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseRemovememberfromproject) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseRemovememberfromproject) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseRemovememberfromproject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseRemovememberfromproject struct {
	value *ResponseRemovememberfromproject
	isSet bool
}

func (v NullableResponseRemovememberfromproject) Get() *ResponseRemovememberfromproject {
	return v.value
}

func (v *NullableResponseRemovememberfromproject) Set(val *ResponseRemovememberfromproject) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRemovememberfromproject) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRemovememberfromproject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRemovememberfromproject(val *ResponseRemovememberfromproject) *NullableResponseRemovememberfromproject {
	return &NullableResponseRemovememberfromproject{value: val, isSet: true}
}

func (v NullableResponseRemovememberfromproject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRemovememberfromproject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


