/*
eHelply SDK - 1.1.119

eHelply SDK for SuperStack Services

API version: 1.1.119
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CreateProjectCredential struct for CreateProjectCredential
type CreateProjectCredential struct {
	ServiceName string `json:"service_name"`
	Secrets []Credential `json:"secrets"`
}

// NewCreateProjectCredential instantiates a new CreateProjectCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectCredential(serviceName string, secrets []Credential) *CreateProjectCredential {
	this := CreateProjectCredential{}
	this.ServiceName = serviceName
	this.Secrets = secrets
	return &this
}

// NewCreateProjectCredentialWithDefaults instantiates a new CreateProjectCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectCredentialWithDefaults() *CreateProjectCredential {
	this := CreateProjectCredential{}
	return &this
}

// GetServiceName returns the ServiceName field value
func (o *CreateProjectCredential) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *CreateProjectCredential) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *CreateProjectCredential) SetServiceName(v string) {
	o.ServiceName = v
}

// GetSecrets returns the Secrets field value
func (o *CreateProjectCredential) GetSecrets() []Credential {
	if o == nil {
		var ret []Credential
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *CreateProjectCredential) GetSecretsOk() ([]Credential, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *CreateProjectCredential) SetSecrets(v []Credential) {
	o.Secrets = v
}

func (o CreateProjectCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["service_name"] = o.ServiceName
	}
	if true {
		toSerialize["secrets"] = o.Secrets
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProjectCredential struct {
	value *CreateProjectCredential
	isSet bool
}

func (v NullableCreateProjectCredential) Get() *CreateProjectCredential {
	return v.value
}

func (v *NullableCreateProjectCredential) Set(val *CreateProjectCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectCredential(val *CreateProjectCredential) *NullableCreateProjectCredential {
	return &NullableCreateProjectCredential{value: val, isSet: true}
}

func (v NullableCreateProjectCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


