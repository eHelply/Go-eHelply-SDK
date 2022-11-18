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

// GetServiceServiceWithSpecsResponse struct for GetServiceServiceWithSpecsResponse
type GetServiceServiceWithSpecsResponse struct {
	Services []string `json:"services"`
}

// NewGetServiceServiceWithSpecsResponse instantiates a new GetServiceServiceWithSpecsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceServiceWithSpecsResponse(services []string) *GetServiceServiceWithSpecsResponse {
	this := GetServiceServiceWithSpecsResponse{}
	this.Services = services
	return &this
}

// NewGetServiceServiceWithSpecsResponseWithDefaults instantiates a new GetServiceServiceWithSpecsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceServiceWithSpecsResponseWithDefaults() *GetServiceServiceWithSpecsResponse {
	this := GetServiceServiceWithSpecsResponse{}
	return &this
}

// GetServices returns the Services field value
func (o *GetServiceServiceWithSpecsResponse) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *GetServiceServiceWithSpecsResponse) GetServicesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Services, true
}

// SetServices sets field value
func (o *GetServiceServiceWithSpecsResponse) SetServices(v []string) {
	o.Services = v
}

func (o GetServiceServiceWithSpecsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableGetServiceServiceWithSpecsResponse struct {
	value *GetServiceServiceWithSpecsResponse
	isSet bool
}

func (v NullableGetServiceServiceWithSpecsResponse) Get() *GetServiceServiceWithSpecsResponse {
	return v.value
}

func (v *NullableGetServiceServiceWithSpecsResponse) Set(val *GetServiceServiceWithSpecsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceServiceWithSpecsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceServiceWithSpecsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceServiceWithSpecsResponse(val *GetServiceServiceWithSpecsResponse) *NullableGetServiceServiceWithSpecsResponse {
	return &NullableGetServiceServiceWithSpecsResponse{value: val, isSet: true}
}

func (v NullableGetServiceServiceWithSpecsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceServiceWithSpecsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


