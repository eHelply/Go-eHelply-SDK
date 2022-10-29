/*
eHelply SDK - 1.1.115

eHelply SDK for SuperStack Services

API version: 1.1.115
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CreateField200Response struct for CreateField200Response
type CreateField200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewCreateField200Response instantiates a new CreateField200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateField200Response() *CreateField200Response {
	this := CreateField200Response{}
	return &this
}

// NewCreateField200ResponseWithDefaults instantiates a new CreateField200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateField200ResponseWithDefaults() *CreateField200Response {
	this := CreateField200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateField200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateField200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateField200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateField200Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateField200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableCreateField200Response struct {
	value *CreateField200Response
	isSet bool
}

func (v NullableCreateField200Response) Get() *CreateField200Response {
	return v.value
}

func (v *NullableCreateField200Response) Set(val *CreateField200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateField200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateField200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateField200Response(val *CreateField200Response) *NullableCreateField200Response {
	return &NullableCreateField200Response{value: val, isSet: true}
}

func (v NullableCreateField200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateField200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


