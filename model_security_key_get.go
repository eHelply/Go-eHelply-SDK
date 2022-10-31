/*
eHelply SDK - 1.1.116

eHelply SDK for SuperStack Services

API version: 1.1.116
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// SecurityKeyGet Used for get endpoint
type SecurityKeyGet struct {
	Uuid string `json:"uuid"`
	Access string `json:"access"`
	Name string `json:"name"`
	Summary string `json:"summary"`
	CreatedAt string `json:"created_at"`
	LastUsedAt string `json:"last_used_at"`
}

// NewSecurityKeyGet instantiates a new SecurityKeyGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityKeyGet(uuid string, access string, name string, summary string, createdAt string, lastUsedAt string) *SecurityKeyGet {
	this := SecurityKeyGet{}
	this.Uuid = uuid
	this.Access = access
	this.Name = name
	this.Summary = summary
	this.CreatedAt = createdAt
	this.LastUsedAt = lastUsedAt
	return &this
}

// NewSecurityKeyGetWithDefaults instantiates a new SecurityKeyGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityKeyGetWithDefaults() *SecurityKeyGet {
	this := SecurityKeyGet{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *SecurityKeyGet) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyGet) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SecurityKeyGet) SetUuid(v string) {
	o.Uuid = v
}

// GetAccess returns the Access field value
func (o *SecurityKeyGet) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyGet) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *SecurityKeyGet) SetAccess(v string) {
	o.Access = v
}

// GetName returns the Name field value
func (o *SecurityKeyGet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityKeyGet) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value
func (o *SecurityKeyGet) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyGet) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *SecurityKeyGet) SetSummary(v string) {
	o.Summary = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SecurityKeyGet) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyGet) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SecurityKeyGet) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetLastUsedAt returns the LastUsedAt field value
func (o *SecurityKeyGet) GetLastUsedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyGet) GetLastUsedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUsedAt, true
}

// SetLastUsedAt sets field value
func (o *SecurityKeyGet) SetLastUsedAt(v string) {
	o.LastUsedAt = v
}

func (o SecurityKeyGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["access"] = o.Access
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityKeyGet struct {
	value *SecurityKeyGet
	isSet bool
}

func (v NullableSecurityKeyGet) Get() *SecurityKeyGet {
	return v.value
}

func (v *NullableSecurityKeyGet) Set(val *SecurityKeyGet) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityKeyGet) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityKeyGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityKeyGet(val *SecurityKeyGet) *NullableSecurityKeyGet {
	return &NullableSecurityKeyGet{value: val, isSet: true}
}

func (v NullableSecurityKeyGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityKeyGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


