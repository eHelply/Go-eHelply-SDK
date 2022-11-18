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

// UpdateProjectCredentialRequest struct for UpdateProjectCredentialRequest
type UpdateProjectCredentialRequest struct {
	Secrets []Credential `json:"secrets"`
}

// NewUpdateProjectCredentialRequest instantiates a new UpdateProjectCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectCredentialRequest(secrets []Credential) *UpdateProjectCredentialRequest {
	this := UpdateProjectCredentialRequest{}
	this.Secrets = secrets
	return &this
}

// NewUpdateProjectCredentialRequestWithDefaults instantiates a new UpdateProjectCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectCredentialRequestWithDefaults() *UpdateProjectCredentialRequest {
	this := UpdateProjectCredentialRequest{}
	return &this
}

// GetSecrets returns the Secrets field value
func (o *UpdateProjectCredentialRequest) GetSecrets() []Credential {
	if o == nil {
		var ret []Credential
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectCredentialRequest) GetSecretsOk() ([]Credential, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *UpdateProjectCredentialRequest) SetSecrets(v []Credential) {
	o.Secrets = v
}

func (o UpdateProjectCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["secrets"] = o.Secrets
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProjectCredentialRequest struct {
	value *UpdateProjectCredentialRequest
	isSet bool
}

func (v NullableUpdateProjectCredentialRequest) Get() *UpdateProjectCredentialRequest {
	return v.value
}

func (v *NullableUpdateProjectCredentialRequest) Set(val *UpdateProjectCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectCredentialRequest(val *UpdateProjectCredentialRequest) *NullableUpdateProjectCredentialRequest {
	return &NullableUpdateProjectCredentialRequest{value: val, isSet: true}
}

func (v NullableUpdateProjectCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


