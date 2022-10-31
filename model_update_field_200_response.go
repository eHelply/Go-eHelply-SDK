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

// UpdateField200Response struct for UpdateField200Response
type UpdateField200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewUpdateField200Response instantiates a new UpdateField200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateField200Response() *UpdateField200Response {
	this := UpdateField200Response{}
	return &this
}

// NewUpdateField200ResponseWithDefaults instantiates a new UpdateField200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateField200ResponseWithDefaults() *UpdateField200Response {
	this := UpdateField200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateField200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateField200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateField200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateField200Response) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateField200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateField200Response struct {
	value *UpdateField200Response
	isSet bool
}

func (v NullableUpdateField200Response) Get() *UpdateField200Response {
	return v.value
}

func (v *NullableUpdateField200Response) Set(val *UpdateField200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateField200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateField200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateField200Response(val *UpdateField200Response) *NullableUpdateField200Response {
	return &NullableUpdateField200Response{value: val, isSet: true}
}

func (v NullableUpdateField200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateField200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


