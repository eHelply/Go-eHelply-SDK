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

// CreateSlug200Response struct for CreateSlug200Response
type CreateSlug200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewCreateSlug200Response instantiates a new CreateSlug200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSlug200Response() *CreateSlug200Response {
	this := CreateSlug200Response{}
	return &this
}

// NewCreateSlug200ResponseWithDefaults instantiates a new CreateSlug200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSlug200ResponseWithDefaults() *CreateSlug200Response {
	this := CreateSlug200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateSlug200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSlug200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateSlug200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateSlug200Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateSlug200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSlug200Response struct {
	value *CreateSlug200Response
	isSet bool
}

func (v NullableCreateSlug200Response) Get() *CreateSlug200Response {
	return v.value
}

func (v *NullableCreateSlug200Response) Set(val *CreateSlug200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSlug200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSlug200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSlug200Response(val *CreateSlug200Response) *NullableCreateSlug200Response {
	return &NullableCreateSlug200Response{value: val, isSet: true}
}

func (v NullableCreateSlug200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSlug200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


