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

// SecurityEncryptionKeyGet struct for SecurityEncryptionKeyGet
type SecurityEncryptionKeyGet struct {
	Key string `json:"key"`
	Category string `json:"category"`
	CreatedAt string `json:"created_at"`
	RetrievedAt string `json:"retrieved_at"`
}

// NewSecurityEncryptionKeyGet instantiates a new SecurityEncryptionKeyGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEncryptionKeyGet(key string, category string, createdAt string, retrievedAt string) *SecurityEncryptionKeyGet {
	this := SecurityEncryptionKeyGet{}
	this.Key = key
	this.Category = category
	this.CreatedAt = createdAt
	this.RetrievedAt = retrievedAt
	return &this
}

// NewSecurityEncryptionKeyGetWithDefaults instantiates a new SecurityEncryptionKeyGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEncryptionKeyGetWithDefaults() *SecurityEncryptionKeyGet {
	this := SecurityEncryptionKeyGet{}
	return &this
}

// GetKey returns the Key field value
func (o *SecurityEncryptionKeyGet) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyGet) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SecurityEncryptionKeyGet) SetKey(v string) {
	o.Key = v
}

// GetCategory returns the Category field value
func (o *SecurityEncryptionKeyGet) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyGet) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *SecurityEncryptionKeyGet) SetCategory(v string) {
	o.Category = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SecurityEncryptionKeyGet) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyGet) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SecurityEncryptionKeyGet) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetRetrievedAt returns the RetrievedAt field value
func (o *SecurityEncryptionKeyGet) GetRetrievedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetrievedAt
}

// GetRetrievedAtOk returns a tuple with the RetrievedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyGet) GetRetrievedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetrievedAt, true
}

// SetRetrievedAt sets field value
func (o *SecurityEncryptionKeyGet) SetRetrievedAt(v string) {
	o.RetrievedAt = v
}

func (o SecurityEncryptionKeyGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["retrieved_at"] = o.RetrievedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityEncryptionKeyGet struct {
	value *SecurityEncryptionKeyGet
	isSet bool
}

func (v NullableSecurityEncryptionKeyGet) Get() *SecurityEncryptionKeyGet {
	return v.value
}

func (v *NullableSecurityEncryptionKeyGet) Set(val *SecurityEncryptionKeyGet) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEncryptionKeyGet) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEncryptionKeyGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEncryptionKeyGet(val *SecurityEncryptionKeyGet) *NullableSecurityEncryptionKeyGet {
	return &NullableSecurityEncryptionKeyGet{value: val, isSet: true}
}

func (v NullableSecurityEncryptionKeyGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEncryptionKeyGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


