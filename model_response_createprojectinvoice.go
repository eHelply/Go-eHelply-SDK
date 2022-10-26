/*
eHelply SDK - 1.1.116

eHelply SDK for SuperStack Services

API version: 1.1.116
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ResponseCreateprojectinvoice struct for ResponseCreateprojectinvoice
type ResponseCreateprojectinvoice struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseCreateprojectinvoice instantiates a new ResponseCreateprojectinvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateprojectinvoice() *ResponseCreateprojectinvoice {
	this := ResponseCreateprojectinvoice{}
	return &this
}

// NewResponseCreateprojectinvoiceWithDefaults instantiates a new ResponseCreateprojectinvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateprojectinvoiceWithDefaults() *ResponseCreateprojectinvoice {
	this := ResponseCreateprojectinvoice{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseCreateprojectinvoice) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateprojectinvoice) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseCreateprojectinvoice) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseCreateprojectinvoice) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseCreateprojectinvoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseCreateprojectinvoice struct {
	value *ResponseCreateprojectinvoice
	isSet bool
}

func (v NullableResponseCreateprojectinvoice) Get() *ResponseCreateprojectinvoice {
	return v.value
}

func (v *NullableResponseCreateprojectinvoice) Set(val *ResponseCreateprojectinvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateprojectinvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateprojectinvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateprojectinvoice(val *ResponseCreateprojectinvoice) *NullableResponseCreateprojectinvoice {
	return &NullableResponseCreateprojectinvoice{value: val, isSet: true}
}

func (v NullableResponseCreateprojectinvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateprojectinvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


