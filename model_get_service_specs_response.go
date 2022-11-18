/*
eHelply SDK - 1.1.122

eHelply SDK for SuperStack Services

API version: 1.1.122
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// GetServiceSpecsResponse struct for GetServiceSpecsResponse
type GetServiceSpecsResponse struct {
	Specs []string `json:"specs"`
}

// NewGetServiceSpecsResponse instantiates a new GetServiceSpecsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceSpecsResponse(specs []string) *GetServiceSpecsResponse {
	this := GetServiceSpecsResponse{}
	this.Specs = specs
	return &this
}

// NewGetServiceSpecsResponseWithDefaults instantiates a new GetServiceSpecsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceSpecsResponseWithDefaults() *GetServiceSpecsResponse {
	this := GetServiceSpecsResponse{}
	return &this
}

// GetSpecs returns the Specs field value
func (o *GetServiceSpecsResponse) GetSpecs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value
// and a boolean to check if the value has been set.
func (o *GetServiceSpecsResponse) GetSpecsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Specs, true
}

// SetSpecs sets field value
func (o *GetServiceSpecsResponse) SetSpecs(v []string) {
	o.Specs = v
}

func (o GetServiceSpecsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["specs"] = o.Specs
	}
	return json.Marshal(toSerialize)
}

type NullableGetServiceSpecsResponse struct {
	value *GetServiceSpecsResponse
	isSet bool
}

func (v NullableGetServiceSpecsResponse) Get() *GetServiceSpecsResponse {
	return v.value
}

func (v *NullableGetServiceSpecsResponse) Set(val *GetServiceSpecsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceSpecsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceSpecsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceSpecsResponse(val *GetServiceSpecsResponse) *NullableGetServiceSpecsResponse {
	return &NullableGetServiceSpecsResponse{value: val, isSet: true}
}

func (v NullableGetServiceSpecsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceSpecsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


