/*
eHelply SDK - 1.1.119

eHelply SDK for SuperStack Services

API version: 1.1.119
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UpdateNote200Response struct for UpdateNote200Response
type UpdateNote200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewUpdateNote200Response instantiates a new UpdateNote200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNote200Response() *UpdateNote200Response {
	this := UpdateNote200Response{}
	return &this
}

// NewUpdateNote200ResponseWithDefaults instantiates a new UpdateNote200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNote200ResponseWithDefaults() *UpdateNote200Response {
	this := UpdateNote200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateNote200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNote200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateNote200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateNote200Response) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateNote200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNote200Response struct {
	value *UpdateNote200Response
	isSet bool
}

func (v NullableUpdateNote200Response) Get() *UpdateNote200Response {
	return v.value
}

func (v *NullableUpdateNote200Response) Set(val *UpdateNote200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNote200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNote200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNote200Response(val *UpdateNote200Response) *NullableUpdateNote200Response {
	return &NullableUpdateNote200Response{value: val, isSet: true}
}

func (v NullableUpdateNote200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNote200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


