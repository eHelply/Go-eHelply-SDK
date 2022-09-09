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

// ResponseDeleteprojectcredential struct for ResponseDeleteprojectcredential
type ResponseDeleteprojectcredential struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseDeleteprojectcredential instantiates a new ResponseDeleteprojectcredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeleteprojectcredential() *ResponseDeleteprojectcredential {
	this := ResponseDeleteprojectcredential{}
	return &this
}

// NewResponseDeleteprojectcredentialWithDefaults instantiates a new ResponseDeleteprojectcredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeleteprojectcredentialWithDefaults() *ResponseDeleteprojectcredential {
	this := ResponseDeleteprojectcredential{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseDeleteprojectcredential) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeleteprojectcredential) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseDeleteprojectcredential) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseDeleteprojectcredential) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseDeleteprojectcredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseDeleteprojectcredential struct {
	value *ResponseDeleteprojectcredential
	isSet bool
}

func (v NullableResponseDeleteprojectcredential) Get() *ResponseDeleteprojectcredential {
	return v.value
}

func (v *NullableResponseDeleteprojectcredential) Set(val *ResponseDeleteprojectcredential) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeleteprojectcredential) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeleteprojectcredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeleteprojectcredential(val *ResponseDeleteprojectcredential) *NullableResponseDeleteprojectcredential {
	return &NullableResponseDeleteprojectcredential{value: val, isSet: true}
}

func (v NullableResponseDeleteprojectcredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeleteprojectcredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


