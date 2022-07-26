/*
eHelply SDK - 1.1.99

eHelply SDK for SuperStack Services

API version: 1.1.99
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// SecurityKeyCreate Used for create endpoint
type SecurityKeyCreate struct {
	Name string `json:"name"`
	Summary string `json:"summary"`
}

// NewSecurityKeyCreate instantiates a new SecurityKeyCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityKeyCreate(name string, summary string) *SecurityKeyCreate {
	this := SecurityKeyCreate{}
	this.Name = name
	this.Summary = summary
	return &this
}

// NewSecurityKeyCreateWithDefaults instantiates a new SecurityKeyCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityKeyCreateWithDefaults() *SecurityKeyCreate {
	this := SecurityKeyCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SecurityKeyCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityKeyCreate) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value
func (o *SecurityKeyCreate) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *SecurityKeyCreate) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *SecurityKeyCreate) SetSummary(v string) {
	o.Summary = v
}

func (o SecurityKeyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityKeyCreate struct {
	value *SecurityKeyCreate
	isSet bool
}

func (v NullableSecurityKeyCreate) Get() *SecurityKeyCreate {
	return v.value
}

func (v *NullableSecurityKeyCreate) Set(val *SecurityKeyCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityKeyCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityKeyCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityKeyCreate(val *SecurityKeyCreate) *NullableSecurityKeyCreate {
	return &NullableSecurityKeyCreate{value: val, isSet: true}
}

func (v NullableSecurityKeyCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityKeyCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


