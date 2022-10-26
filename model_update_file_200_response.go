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

// UpdateFile200Response struct for UpdateFile200Response
type UpdateFile200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewUpdateFile200Response instantiates a new UpdateFile200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFile200Response() *UpdateFile200Response {
	this := UpdateFile200Response{}
	return &this
}

// NewUpdateFile200ResponseWithDefaults instantiates a new UpdateFile200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFile200ResponseWithDefaults() *UpdateFile200Response {
	this := UpdateFile200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateFile200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateFile200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateFile200Response) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateFile200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFile200Response struct {
	value *UpdateFile200Response
	isSet bool
}

func (v NullableUpdateFile200Response) Get() *UpdateFile200Response {
	return v.value
}

func (v *NullableUpdateFile200Response) Set(val *UpdateFile200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFile200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFile200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFile200Response(val *UpdateFile200Response) *NullableUpdateFile200Response {
	return &NullableUpdateFile200Response{value: val, isSet: true}
}

func (v NullableUpdateFile200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFile200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


