/*
eHelply SDK - 1.1.103

eHelply SDK for SuperStack Services

API version: 1.1.103
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserTokenReturn Tokens, naming scheme based off of cognito return fields
type UserTokenReturn struct {
	AccessToken string `json:"AccessToken"`
	ExpiresIn int32 `json:"ExpiresIn"`
	TokenType string `json:"TokenType"`
	IdToken string `json:"IdToken"`
}

// NewUserTokenReturn instantiates a new UserTokenReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTokenReturn(accessToken string, expiresIn int32, tokenType string, idToken string) *UserTokenReturn {
	this := UserTokenReturn{}
	this.AccessToken = accessToken
	this.ExpiresIn = expiresIn
	this.TokenType = tokenType
	this.IdToken = idToken
	return &this
}

// NewUserTokenReturnWithDefaults instantiates a new UserTokenReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTokenReturnWithDefaults() *UserTokenReturn {
	this := UserTokenReturn{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *UserTokenReturn) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *UserTokenReturn) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *UserTokenReturn) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *UserTokenReturn) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *UserTokenReturn) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *UserTokenReturn) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetTokenType returns the TokenType field value
func (o *UserTokenReturn) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *UserTokenReturn) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *UserTokenReturn) SetTokenType(v string) {
	o.TokenType = v
}

// GetIdToken returns the IdToken field value
func (o *UserTokenReturn) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *UserTokenReturn) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *UserTokenReturn) SetIdToken(v string) {
	o.IdToken = v
}

func (o UserTokenReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AccessToken"] = o.AccessToken
	}
	if true {
		toSerialize["ExpiresIn"] = o.ExpiresIn
	}
	if true {
		toSerialize["TokenType"] = o.TokenType
	}
	if true {
		toSerialize["IdToken"] = o.IdToken
	}
	return json.Marshal(toSerialize)
}

type NullableUserTokenReturn struct {
	value *UserTokenReturn
	isSet bool
}

func (v NullableUserTokenReturn) Get() *UserTokenReturn {
	return v.value
}

func (v *NullableUserTokenReturn) Set(val *UserTokenReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTokenReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTokenReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTokenReturn(val *UserTokenReturn) *NullableUserTokenReturn {
	return &NullableUserTokenReturn{value: val, isSet: true}
}

func (v NullableUserTokenReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTokenReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


