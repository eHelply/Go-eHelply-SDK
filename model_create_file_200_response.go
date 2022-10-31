/*
eHelply SDK - 1.1.120

eHelply SDK for SuperStack Services

API version: 1.1.120
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CreateFile200Response struct for CreateFile200Response
type CreateFile200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewCreateFile200Response instantiates a new CreateFile200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFile200Response() *CreateFile200Response {
	this := CreateFile200Response{}
	return &this
}

// NewCreateFile200ResponseWithDefaults instantiates a new CreateFile200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFile200ResponseWithDefaults() *CreateFile200Response {
	this := CreateFile200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateFile200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFile200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateFile200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateFile200Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateFile200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFile200Response struct {
	value *CreateFile200Response
	isSet bool
}

func (v NullableCreateFile200Response) Get() *CreateFile200Response {
	return v.value
}

func (v *NullableCreateFile200Response) Set(val *CreateFile200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFile200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFile200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFile200Response(val *CreateFile200Response) *NullableCreateFile200Response {
	return &NullableCreateFile200Response{value: val, isSet: true}
}

func (v NullableCreateFile200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFile200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


