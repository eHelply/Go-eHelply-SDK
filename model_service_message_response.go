/*
eHelply SDK - 1.1.111

eHelply SDK for SuperStack Services

API version: 1.1.111
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceMessageResponse struct for ServiceMessageResponse
type ServiceMessageResponse struct {
	Message string `json:"message"`
}

// NewServiceMessageResponse instantiates a new ServiceMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceMessageResponse(message string) *ServiceMessageResponse {
	this := ServiceMessageResponse{}
	this.Message = message
	return &this
}

// NewServiceMessageResponseWithDefaults instantiates a new ServiceMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceMessageResponseWithDefaults() *ServiceMessageResponse {
	this := ServiceMessageResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *ServiceMessageResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ServiceMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ServiceMessageResponse) SetMessage(v string) {
	o.Message = v
}

func (o ServiceMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableServiceMessageResponse struct {
	value *ServiceMessageResponse
	isSet bool
}

func (v NullableServiceMessageResponse) Get() *ServiceMessageResponse {
	return v.value
}

func (v *NullableServiceMessageResponse) Set(val *ServiceMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceMessageResponse(val *ServiceMessageResponse) *NullableServiceMessageResponse {
	return &NullableServiceMessageResponse{value: val, isSet: true}
}

func (v NullableServiceMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


