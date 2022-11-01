/*
eHelply SDK - 1.1.117

eHelply SDK for SuperStack Services

API version: 1.1.117
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// DeleteMeta200Response struct for DeleteMeta200Response
type DeleteMeta200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewDeleteMeta200Response instantiates a new DeleteMeta200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMeta200Response() *DeleteMeta200Response {
	this := DeleteMeta200Response{}
	return &this
}

// NewDeleteMeta200ResponseWithDefaults instantiates a new DeleteMeta200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMeta200ResponseWithDefaults() *DeleteMeta200Response {
	this := DeleteMeta200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteMeta200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMeta200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteMeta200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteMeta200Response) SetMessage(v string) {
	o.Message = &v
}

func (o DeleteMeta200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteMeta200Response struct {
	value *DeleteMeta200Response
	isSet bool
}

func (v NullableDeleteMeta200Response) Get() *DeleteMeta200Response {
	return v.value
}

func (v *NullableDeleteMeta200Response) Set(val *DeleteMeta200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMeta200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMeta200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMeta200Response(val *DeleteMeta200Response) *NullableDeleteMeta200Response {
	return &NullableDeleteMeta200Response{value: val, isSet: true}
}

func (v NullableDeleteMeta200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMeta200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


