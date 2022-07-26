/*
eHelply SDK - 1.1.95

eHelply SDK for SuperStack Services

API version: 1.1.95
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
	"os"
)

// SecurityEncryptionKeyResponse struct for SecurityEncryptionKeyResponse
type SecurityEncryptionKeyResponse struct {
	Uuid string `json:"uuid"`
	Key *os.File `json:"key"`
	Category string `json:"category"`
	DeletedAt string `json:"deleted_at"`
	CreatedAt string `json:"created_at"`
	RetrievedAt string `json:"retrieved_at"`
}

// NewSecurityEncryptionKeyResponse instantiates a new SecurityEncryptionKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEncryptionKeyResponse(uuid string, key *os.File, category string, deletedAt string, createdAt string, retrievedAt string) *SecurityEncryptionKeyResponse {
	this := SecurityEncryptionKeyResponse{}
	this.Uuid = uuid
	this.Key = key
	this.Category = category
	this.DeletedAt = deletedAt
	this.CreatedAt = createdAt
	this.RetrievedAt = retrievedAt
	return &this
}

// NewSecurityEncryptionKeyResponseWithDefaults instantiates a new SecurityEncryptionKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEncryptionKeyResponseWithDefaults() *SecurityEncryptionKeyResponse {
	this := SecurityEncryptionKeyResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *SecurityEncryptionKeyResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SecurityEncryptionKeyResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetKey returns the Key field value
func (o *SecurityEncryptionKeyResponse) GetKey() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyResponse) GetKeyOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SecurityEncryptionKeyResponse) SetKey(v *os.File) {
	o.Key = v
}

// GetCategory returns the Category field value
func (o *SecurityEncryptionKeyResponse) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyResponse) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *SecurityEncryptionKeyResponse) SetCategory(v string) {
	o.Category = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *SecurityEncryptionKeyResponse) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *SecurityEncryptionKeyResponse) SetDeletedAt(v string) {
	o.DeletedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SecurityEncryptionKeyResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SecurityEncryptionKeyResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetRetrievedAt returns the RetrievedAt field value
func (o *SecurityEncryptionKeyResponse) GetRetrievedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetrievedAt
}

// GetRetrievedAtOk returns a tuple with the RetrievedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityEncryptionKeyResponse) GetRetrievedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetrievedAt, true
}

// SetRetrievedAt sets field value
func (o *SecurityEncryptionKeyResponse) SetRetrievedAt(v string) {
	o.RetrievedAt = v
}

func (o SecurityEncryptionKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["retrieved_at"] = o.RetrievedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityEncryptionKeyResponse struct {
	value *SecurityEncryptionKeyResponse
	isSet bool
}

func (v NullableSecurityEncryptionKeyResponse) Get() *SecurityEncryptionKeyResponse {
	return v.value
}

func (v *NullableSecurityEncryptionKeyResponse) Set(val *SecurityEncryptionKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEncryptionKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEncryptionKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEncryptionKeyResponse(val *SecurityEncryptionKeyResponse) *NullableSecurityEncryptionKeyResponse {
	return &NullableSecurityEncryptionKeyResponse{value: val, isSet: true}
}

func (v NullableSecurityEncryptionKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEncryptionKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


