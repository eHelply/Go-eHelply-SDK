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

// GetProjectCredential struct for GetProjectCredential
type GetProjectCredential struct {
	ServiceName string `json:"service_name"`
	Secrets []GetSecret `json:"secrets"`
}

// NewGetProjectCredential instantiates a new GetProjectCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectCredential(serviceName string, secrets []GetSecret) *GetProjectCredential {
	this := GetProjectCredential{}
	this.ServiceName = serviceName
	this.Secrets = secrets
	return &this
}

// NewGetProjectCredentialWithDefaults instantiates a new GetProjectCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectCredentialWithDefaults() *GetProjectCredential {
	this := GetProjectCredential{}
	return &this
}

// GetServiceName returns the ServiceName field value
func (o *GetProjectCredential) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *GetProjectCredential) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *GetProjectCredential) SetServiceName(v string) {
	o.ServiceName = v
}

// GetSecrets returns the Secrets field value
func (o *GetProjectCredential) GetSecrets() []GetSecret {
	if o == nil {
		var ret []GetSecret
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *GetProjectCredential) GetSecretsOk() ([]GetSecret, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *GetProjectCredential) SetSecrets(v []GetSecret) {
	o.Secrets = v
}

func (o GetProjectCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["service_name"] = o.ServiceName
	}
	if true {
		toSerialize["secrets"] = o.Secrets
	}
	return json.Marshal(toSerialize)
}

type NullableGetProjectCredential struct {
	value *GetProjectCredential
	isSet bool
}

func (v NullableGetProjectCredential) Get() *GetProjectCredential {
	return v.value
}

func (v *NullableGetProjectCredential) Set(val *GetProjectCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectCredential(val *GetProjectCredential) *NullableGetProjectCredential {
	return &NullableGetProjectCredential{value: val, isSet: true}
}

func (v NullableGetProjectCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


