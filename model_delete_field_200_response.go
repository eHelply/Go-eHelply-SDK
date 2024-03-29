/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// DeleteField200Response struct for DeleteField200Response
type DeleteField200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewDeleteField200Response instantiates a new DeleteField200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteField200Response() *DeleteField200Response {
	this := DeleteField200Response{}
	return &this
}

// NewDeleteField200ResponseWithDefaults instantiates a new DeleteField200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteField200ResponseWithDefaults() *DeleteField200Response {
	this := DeleteField200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteField200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteField200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteField200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteField200Response) SetMessage(v string) {
	o.Message = &v
}

func (o DeleteField200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteField200Response struct {
	value *DeleteField200Response
	isSet bool
}

func (v NullableDeleteField200Response) Get() *DeleteField200Response {
	return v.value
}

func (v *NullableDeleteField200Response) Set(val *DeleteField200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteField200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteField200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteField200Response(val *DeleteField200Response) *NullableDeleteField200Response {
	return &NullableDeleteField200Response{value: val, isSet: true}
}

func (v NullableDeleteField200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteField200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


