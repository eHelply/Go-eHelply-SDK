/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceSuperStackMetaFeature struct for ServiceSuperStackMetaFeature
type ServiceSuperStackMetaFeature struct {
	Name string `json:"name"`
	Summary string `json:"summary"`
}

// NewServiceSuperStackMetaFeature instantiates a new ServiceSuperStackMetaFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSuperStackMetaFeature(name string, summary string) *ServiceSuperStackMetaFeature {
	this := ServiceSuperStackMetaFeature{}
	this.Name = name
	this.Summary = summary
	return &this
}

// NewServiceSuperStackMetaFeatureWithDefaults instantiates a new ServiceSuperStackMetaFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSuperStackMetaFeatureWithDefaults() *ServiceSuperStackMetaFeature {
	this := ServiceSuperStackMetaFeature{}
	return &this
}

// GetName returns the Name field value
func (o *ServiceSuperStackMetaFeature) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaFeature) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceSuperStackMetaFeature) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value
func (o *ServiceSuperStackMetaFeature) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaFeature) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ServiceSuperStackMetaFeature) SetSummary(v string) {
	o.Summary = v
}

func (o ServiceSuperStackMetaFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSuperStackMetaFeature struct {
	value *ServiceSuperStackMetaFeature
	isSet bool
}

func (v NullableServiceSuperStackMetaFeature) Get() *ServiceSuperStackMetaFeature {
	return v.value
}

func (v *NullableServiceSuperStackMetaFeature) Set(val *ServiceSuperStackMetaFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSuperStackMetaFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSuperStackMetaFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSuperStackMetaFeature(val *ServiceSuperStackMetaFeature) *NullableServiceSuperStackMetaFeature {
	return &NullableServiceSuperStackMetaFeature{value: val, isSet: true}
}

func (v NullableServiceSuperStackMetaFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSuperStackMetaFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


