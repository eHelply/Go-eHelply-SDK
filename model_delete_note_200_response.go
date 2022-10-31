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

// DeleteNote200Response struct for DeleteNote200Response
type DeleteNote200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewDeleteNote200Response instantiates a new DeleteNote200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNote200Response() *DeleteNote200Response {
	this := DeleteNote200Response{}
	return &this
}

// NewDeleteNote200ResponseWithDefaults instantiates a new DeleteNote200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNote200ResponseWithDefaults() *DeleteNote200Response {
	this := DeleteNote200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteNote200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteNote200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteNote200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteNote200Response) SetMessage(v string) {
	o.Message = &v
}

func (o DeleteNote200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteNote200Response struct {
	value *DeleteNote200Response
	isSet bool
}

func (v NullableDeleteNote200Response) Get() *DeleteNote200Response {
	return v.value
}

func (v *NullableDeleteNote200Response) Set(val *DeleteNote200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNote200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNote200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNote200Response(val *DeleteNote200Response) *NullableDeleteNote200Response {
	return &NullableDeleteNote200Response{value: val, isSet: true}
}

func (v NullableDeleteNote200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteNote200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


