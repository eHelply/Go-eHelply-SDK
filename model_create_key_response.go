/*
eHelply SDK - 1.1.90

eHelply SDK for SuperStack Services

API version: 1.1.90
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CreateKeyResponse struct for CreateKeyResponse
type CreateKeyResponse struct {
	Uuid string `json:"uuid"`
	Access string `json:"access"`
	Secret string `json:"secret"`
}

// NewCreateKeyResponse instantiates a new CreateKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyResponse(uuid string, access string, secret string) *CreateKeyResponse {
	this := CreateKeyResponse{}
	this.Uuid = uuid
	this.Access = access
	this.Secret = secret
	return &this
}

// NewCreateKeyResponseWithDefaults instantiates a new CreateKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyResponseWithDefaults() *CreateKeyResponse {
	this := CreateKeyResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *CreateKeyResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CreateKeyResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CreateKeyResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetAccess returns the Access field value
func (o *CreateKeyResponse) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *CreateKeyResponse) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *CreateKeyResponse) SetAccess(v string) {
	o.Access = v
}

// GetSecret returns the Secret field value
func (o *CreateKeyResponse) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *CreateKeyResponse) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *CreateKeyResponse) SetSecret(v string) {
	o.Secret = v
}

func (o CreateKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["access"] = o.Access
	}
	if true {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableCreateKeyResponse struct {
	value *CreateKeyResponse
	isSet bool
}

func (v NullableCreateKeyResponse) Get() *CreateKeyResponse {
	return v.value
}

func (v *NullableCreateKeyResponse) Set(val *CreateKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeyResponse(val *CreateKeyResponse) *NullableCreateKeyResponse {
	return &NullableCreateKeyResponse{value: val, isSet: true}
}

func (v NullableCreateKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


