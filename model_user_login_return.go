/*
eHelply SDK - 1.1.109

eHelply SDK for SuperStack Services

API version: 1.1.109
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserLoginReturn Tokens (refresh, access, identity, ect)
type UserLoginReturn struct {
	AccessToken string `json:"AccessToken"`
	ExpiresIn int32 `json:"ExpiresIn"`
	TokenType string `json:"TokenType"`
	IdToken string `json:"IdToken"`
	RefreshToken string `json:"RefreshToken"`
}

// NewUserLoginReturn instantiates a new UserLoginReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLoginReturn(accessToken string, expiresIn int32, tokenType string, idToken string, refreshToken string) *UserLoginReturn {
	this := UserLoginReturn{}
	this.AccessToken = accessToken
	this.ExpiresIn = expiresIn
	this.TokenType = tokenType
	this.IdToken = idToken
	this.RefreshToken = refreshToken
	return &this
}

// NewUserLoginReturnWithDefaults instantiates a new UserLoginReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginReturnWithDefaults() *UserLoginReturn {
	this := UserLoginReturn{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *UserLoginReturn) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *UserLoginReturn) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *UserLoginReturn) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *UserLoginReturn) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *UserLoginReturn) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *UserLoginReturn) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetTokenType returns the TokenType field value
func (o *UserLoginReturn) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *UserLoginReturn) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *UserLoginReturn) SetTokenType(v string) {
	o.TokenType = v
}

// GetIdToken returns the IdToken field value
func (o *UserLoginReturn) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *UserLoginReturn) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *UserLoginReturn) SetIdToken(v string) {
	o.IdToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *UserLoginReturn) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *UserLoginReturn) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *UserLoginReturn) SetRefreshToken(v string) {
	o.RefreshToken = v
}

func (o UserLoginReturn) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["RefreshToken"] = o.RefreshToken
	}
	return json.Marshal(toSerialize)
}

type NullableUserLoginReturn struct {
	value *UserLoginReturn
	isSet bool
}

func (v NullableUserLoginReturn) Get() *UserLoginReturn {
	return v.value
}

func (v *NullableUserLoginReturn) Set(val *UserLoginReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLoginReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLoginReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLoginReturn(val *UserLoginReturn) *NullableUserLoginReturn {
	return &NullableUserLoginReturn{value: val, isSet: true}
}

func (v NullableUserLoginReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLoginReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


