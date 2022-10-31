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

// GetAppointment403Response struct for GetAppointment403Response
type GetAppointment403Response struct {
	Message *string `json:"message,omitempty"`
}

// NewGetAppointment403Response instantiates a new GetAppointment403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppointment403Response() *GetAppointment403Response {
	this := GetAppointment403Response{}
	return &this
}

// NewGetAppointment403ResponseWithDefaults instantiates a new GetAppointment403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppointment403ResponseWithDefaults() *GetAppointment403Response {
	this := GetAppointment403Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetAppointment403Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppointment403Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetAppointment403Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetAppointment403Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetAppointment403Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableGetAppointment403Response struct {
	value *GetAppointment403Response
	isSet bool
}

func (v NullableGetAppointment403Response) Get() *GetAppointment403Response {
	return v.value
}

func (v *NullableGetAppointment403Response) Set(val *GetAppointment403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppointment403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppointment403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppointment403Response(val *GetAppointment403Response) *NullableGetAppointment403Response {
	return &NullableGetAppointment403Response{value: val, isSet: true}
}

func (v NullableGetAppointment403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAppointment403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


