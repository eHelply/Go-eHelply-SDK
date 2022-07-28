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

// ResponseGeneratetoken struct for ResponseGeneratetoken
type ResponseGeneratetoken struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseGeneratetoken instantiates a new ResponseGeneratetoken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseGeneratetoken() *ResponseGeneratetoken {
	this := ResponseGeneratetoken{}
	return &this
}

// NewResponseGeneratetokenWithDefaults instantiates a new ResponseGeneratetoken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseGeneratetokenWithDefaults() *ResponseGeneratetoken {
	this := ResponseGeneratetoken{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseGeneratetoken) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseGeneratetoken) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseGeneratetoken) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseGeneratetoken) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseGeneratetoken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseGeneratetoken struct {
	value *ResponseGeneratetoken
	isSet bool
}

func (v NullableResponseGeneratetoken) Get() *ResponseGeneratetoken {
	return v.value
}

func (v *NullableResponseGeneratetoken) Set(val *ResponseGeneratetoken) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseGeneratetoken) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseGeneratetoken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseGeneratetoken(val *ResponseGeneratetoken) *NullableResponseGeneratetoken {
	return &NullableResponseGeneratetoken{value: val, isSet: true}
}

func (v NullableResponseGeneratetoken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseGeneratetoken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


