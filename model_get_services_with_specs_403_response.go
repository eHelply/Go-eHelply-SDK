/*
eHelply SDK - 1.1.101

eHelply SDK for SuperStack Services

API version: 1.1.101
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// GetServicesWithSpecs403Response struct for GetServicesWithSpecs403Response
type GetServicesWithSpecs403Response struct {
	Message *string `json:"message,omitempty"`
}

// NewGetServicesWithSpecs403Response instantiates a new GetServicesWithSpecs403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServicesWithSpecs403Response() *GetServicesWithSpecs403Response {
	this := GetServicesWithSpecs403Response{}
	return &this
}

// NewGetServicesWithSpecs403ResponseWithDefaults instantiates a new GetServicesWithSpecs403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServicesWithSpecs403ResponseWithDefaults() *GetServicesWithSpecs403Response {
	this := GetServicesWithSpecs403Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetServicesWithSpecs403Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServicesWithSpecs403Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetServicesWithSpecs403Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetServicesWithSpecs403Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetServicesWithSpecs403Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableGetServicesWithSpecs403Response struct {
	value *GetServicesWithSpecs403Response
	isSet bool
}

func (v NullableGetServicesWithSpecs403Response) Get() *GetServicesWithSpecs403Response {
	return v.value
}

func (v *NullableGetServicesWithSpecs403Response) Set(val *GetServicesWithSpecs403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServicesWithSpecs403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServicesWithSpecs403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServicesWithSpecs403Response(val *GetServicesWithSpecs403Response) *NullableGetServicesWithSpecs403Response {
	return &NullableGetServicesWithSpecs403Response{value: val, isSet: true}
}

func (v NullableGetServicesWithSpecs403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServicesWithSpecs403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


