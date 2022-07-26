/*
eHelply SDK - 1.1.94

eHelply SDK for SuperStack Services

API version: 1.1.94
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// GetServiceSpecResponse struct for GetServiceSpecResponse
type GetServiceSpecResponse struct {
	Contents string `json:"contents"`
}

// NewGetServiceSpecResponse instantiates a new GetServiceSpecResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceSpecResponse(contents string) *GetServiceSpecResponse {
	this := GetServiceSpecResponse{}
	this.Contents = contents
	return &this
}

// NewGetServiceSpecResponseWithDefaults instantiates a new GetServiceSpecResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceSpecResponseWithDefaults() *GetServiceSpecResponse {
	this := GetServiceSpecResponse{}
	return &this
}

// GetContents returns the Contents field value
func (o *GetServiceSpecResponse) GetContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *GetServiceSpecResponse) GetContentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *GetServiceSpecResponse) SetContents(v string) {
	o.Contents = v
}

func (o GetServiceSpecResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableGetServiceSpecResponse struct {
	value *GetServiceSpecResponse
	isSet bool
}

func (v NullableGetServiceSpecResponse) Get() *GetServiceSpecResponse {
	return v.value
}

func (v *NullableGetServiceSpecResponse) Set(val *GetServiceSpecResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceSpecResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceSpecResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceSpecResponse(val *GetServiceSpecResponse) *NullableGetServiceSpecResponse {
	return &NullableGetServiceSpecResponse{value: val, isSet: true}
}

func (v NullableGetServiceSpecResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceSpecResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


