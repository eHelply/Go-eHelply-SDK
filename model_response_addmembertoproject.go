/*
eHelply SDK - 1.1.89

eHelply SDK for SuperStack Services

API version: 1.1.89
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ResponseAddmembertoproject struct for ResponseAddmembertoproject
type ResponseAddmembertoproject struct {
	Message *string `json:"message,omitempty"`
}

// NewResponseAddmembertoproject instantiates a new ResponseAddmembertoproject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAddmembertoproject() *ResponseAddmembertoproject {
	this := ResponseAddmembertoproject{}
	return &this
}

// NewResponseAddmembertoprojectWithDefaults instantiates a new ResponseAddmembertoproject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAddmembertoprojectWithDefaults() *ResponseAddmembertoproject {
	this := ResponseAddmembertoproject{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseAddmembertoproject) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAddmembertoproject) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseAddmembertoproject) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseAddmembertoproject) SetMessage(v string) {
	o.Message = &v
}

func (o ResponseAddmembertoproject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableResponseAddmembertoproject struct {
	value *ResponseAddmembertoproject
	isSet bool
}

func (v NullableResponseAddmembertoproject) Get() *ResponseAddmembertoproject {
	return v.value
}

func (v *NullableResponseAddmembertoproject) Set(val *ResponseAddmembertoproject) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAddmembertoproject) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAddmembertoproject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAddmembertoproject(val *ResponseAddmembertoproject) *NullableResponseAddmembertoproject {
	return &NullableResponseAddmembertoproject{value: val, isSet: true}
}

func (v NullableResponseAddmembertoproject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAddmembertoproject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


