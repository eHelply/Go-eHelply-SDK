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

// ResponseRevokeprojectcredit struct for ResponseRevokeprojectcredit
type ResponseRevokeprojectcredit struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseRevokeprojectcredit instantiates a new ResponseRevokeprojectcredit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRevokeprojectcredit() *ResponseRevokeprojectcredit {
	this := ResponseRevokeprojectcredit{}
	return &this
}

// NewResponseRevokeprojectcreditWithDefaults instantiates a new ResponseRevokeprojectcredit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRevokeprojectcreditWithDefaults() *ResponseRevokeprojectcredit {
	this := ResponseRevokeprojectcredit{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseRevokeprojectcredit) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseRevokeprojectcredit) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseRevokeprojectcredit) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseRevokeprojectcredit) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseRevokeprojectcredit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseRevokeprojectcredit struct {
	value *ResponseRevokeprojectcredit
	isSet bool
}

func (v NullableResponseRevokeprojectcredit) Get() *ResponseRevokeprojectcredit {
	return v.value
}

func (v *NullableResponseRevokeprojectcredit) Set(val *ResponseRevokeprojectcredit) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRevokeprojectcredit) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRevokeprojectcredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRevokeprojectcredit(val *ResponseRevokeprojectcredit) *NullableResponseRevokeprojectcredit {
	return &NullableResponseRevokeprojectcredit{value: val, isSet: true}
}

func (v NullableResponseRevokeprojectcredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRevokeprojectcredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


