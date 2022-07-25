/*
eHelply SDK - 1.1.92

eHelply SDK for SuperStack Services

API version: 1.1.92
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// StripeCustomerSecretResponse struct for StripeCustomerSecretResponse
type StripeCustomerSecretResponse struct {
	Secret string `json:"secret"`
}

// NewStripeCustomerSecretResponse instantiates a new StripeCustomerSecretResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeCustomerSecretResponse(secret string) *StripeCustomerSecretResponse {
	this := StripeCustomerSecretResponse{}
	this.Secret = secret
	return &this
}

// NewStripeCustomerSecretResponseWithDefaults instantiates a new StripeCustomerSecretResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeCustomerSecretResponseWithDefaults() *StripeCustomerSecretResponse {
	this := StripeCustomerSecretResponse{}
	return &this
}

// GetSecret returns the Secret field value
func (o *StripeCustomerSecretResponse) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *StripeCustomerSecretResponse) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *StripeCustomerSecretResponse) SetSecret(v string) {
	o.Secret = v
}

func (o StripeCustomerSecretResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableStripeCustomerSecretResponse struct {
	value *StripeCustomerSecretResponse
	isSet bool
}

func (v NullableStripeCustomerSecretResponse) Get() *StripeCustomerSecretResponse {
	return v.value
}

func (v *NullableStripeCustomerSecretResponse) Set(val *StripeCustomerSecretResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeCustomerSecretResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeCustomerSecretResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeCustomerSecretResponse(val *StripeCustomerSecretResponse) *NullableStripeCustomerSecretResponse {
	return &NullableStripeCustomerSecretResponse{value: val, isSet: true}
}

func (v NullableStripeCustomerSecretResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeCustomerSecretResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


