/*
eHelply SDK - 1.1.112

eHelply SDK for SuperStack Services

API version: 1.1.112
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ResponseDeleteusagetype struct for ResponseDeleteusagetype
type ResponseDeleteusagetype struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseDeleteusagetype instantiates a new ResponseDeleteusagetype object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeleteusagetype() *ResponseDeleteusagetype {
	this := ResponseDeleteusagetype{}
	return &this
}

// NewResponseDeleteusagetypeWithDefaults instantiates a new ResponseDeleteusagetype object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeleteusagetypeWithDefaults() *ResponseDeleteusagetype {
	this := ResponseDeleteusagetype{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseDeleteusagetype) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeleteusagetype) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseDeleteusagetype) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseDeleteusagetype) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseDeleteusagetype) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseDeleteusagetype struct {
	value *ResponseDeleteusagetype
	isSet bool
}

func (v NullableResponseDeleteusagetype) Get() *ResponseDeleteusagetype {
	return v.value
}

func (v *NullableResponseDeleteusagetype) Set(val *ResponseDeleteusagetype) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeleteusagetype) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeleteusagetype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeleteusagetype(val *ResponseDeleteusagetype) *NullableResponseDeleteusagetype {
	return &NullableResponseDeleteusagetype{value: val, isSet: true}
}

func (v NullableResponseDeleteusagetype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeleteusagetype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


