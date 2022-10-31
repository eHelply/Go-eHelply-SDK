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

// ServiceSuperStackMetaUseCase struct for ServiceSuperStackMetaUseCase
type ServiceSuperStackMetaUseCase struct {
	Name string `json:"name"`
	Summary string `json:"summary"`
}

// NewServiceSuperStackMetaUseCase instantiates a new ServiceSuperStackMetaUseCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSuperStackMetaUseCase(name string, summary string) *ServiceSuperStackMetaUseCase {
	this := ServiceSuperStackMetaUseCase{}
	this.Name = name
	this.Summary = summary
	return &this
}

// NewServiceSuperStackMetaUseCaseWithDefaults instantiates a new ServiceSuperStackMetaUseCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSuperStackMetaUseCaseWithDefaults() *ServiceSuperStackMetaUseCase {
	this := ServiceSuperStackMetaUseCase{}
	return &this
}

// GetName returns the Name field value
func (o *ServiceSuperStackMetaUseCase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaUseCase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceSuperStackMetaUseCase) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value
func (o *ServiceSuperStackMetaUseCase) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaUseCase) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ServiceSuperStackMetaUseCase) SetSummary(v string) {
	o.Summary = v
}

func (o ServiceSuperStackMetaUseCase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSuperStackMetaUseCase struct {
	value *ServiceSuperStackMetaUseCase
	isSet bool
}

func (v NullableServiceSuperStackMetaUseCase) Get() *ServiceSuperStackMetaUseCase {
	return v.value
}

func (v *NullableServiceSuperStackMetaUseCase) Set(val *ServiceSuperStackMetaUseCase) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSuperStackMetaUseCase) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSuperStackMetaUseCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSuperStackMetaUseCase(val *ServiceSuperStackMetaUseCase) *NullableServiceSuperStackMetaUseCase {
	return &NullableServiceSuperStackMetaUseCase{value: val, isSet: true}
}

func (v NullableServiceSuperStackMetaUseCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSuperStackMetaUseCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


