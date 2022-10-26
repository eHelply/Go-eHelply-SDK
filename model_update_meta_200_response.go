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

// UpdateMeta200Response struct for UpdateMeta200Response
type UpdateMeta200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewUpdateMeta200Response instantiates a new UpdateMeta200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMeta200Response() *UpdateMeta200Response {
	this := UpdateMeta200Response{}
	return &this
}

// NewUpdateMeta200ResponseWithDefaults instantiates a new UpdateMeta200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMeta200ResponseWithDefaults() *UpdateMeta200Response {
	this := UpdateMeta200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateMeta200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeta200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateMeta200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateMeta200Response) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateMeta200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateMeta200Response struct {
	value *UpdateMeta200Response
	isSet bool
}

func (v NullableUpdateMeta200Response) Get() *UpdateMeta200Response {
	return v.value
}

func (v *NullableUpdateMeta200Response) Set(val *UpdateMeta200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMeta200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMeta200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMeta200Response(val *UpdateMeta200Response) *NullableUpdateMeta200Response {
	return &NullableUpdateMeta200Response{value: val, isSet: true}
}

func (v NullableUpdateMeta200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMeta200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


