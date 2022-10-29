/*
eHelply SDK - 1.1.119

eHelply SDK for SuperStack Services

API version: 1.1.119
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceSuperStackMetaGettingStartedEndpointTeaser struct for ServiceSuperStackMetaGettingStartedEndpointTeaser
type ServiceSuperStackMetaGettingStartedEndpointTeaser struct {
	Path string `json:"path"`
	Method string `json:"method"`
	Description string `json:"description"`
}

// NewServiceSuperStackMetaGettingStartedEndpointTeaser instantiates a new ServiceSuperStackMetaGettingStartedEndpointTeaser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSuperStackMetaGettingStartedEndpointTeaser(path string, method string, description string) *ServiceSuperStackMetaGettingStartedEndpointTeaser {
	this := ServiceSuperStackMetaGettingStartedEndpointTeaser{}
	this.Path = path
	this.Method = method
	this.Description = description
	return &this
}

// NewServiceSuperStackMetaGettingStartedEndpointTeaserWithDefaults instantiates a new ServiceSuperStackMetaGettingStartedEndpointTeaser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSuperStackMetaGettingStartedEndpointTeaserWithDefaults() *ServiceSuperStackMetaGettingStartedEndpointTeaser {
	this := ServiceSuperStackMetaGettingStartedEndpointTeaser{}
	return &this
}

// GetPath returns the Path field value
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) SetPath(v string) {
	o.Path = v
}

// GetMethod returns the Method field value
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) SetMethod(v string) {
	o.Method = v
}

// GetDescription returns the Description field value
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceSuperStackMetaGettingStartedEndpointTeaser) SetDescription(v string) {
	o.Description = v
}

func (o ServiceSuperStackMetaGettingStartedEndpointTeaser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSuperStackMetaGettingStartedEndpointTeaser struct {
	value *ServiceSuperStackMetaGettingStartedEndpointTeaser
	isSet bool
}

func (v NullableServiceSuperStackMetaGettingStartedEndpointTeaser) Get() *ServiceSuperStackMetaGettingStartedEndpointTeaser {
	return v.value
}

func (v *NullableServiceSuperStackMetaGettingStartedEndpointTeaser) Set(val *ServiceSuperStackMetaGettingStartedEndpointTeaser) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSuperStackMetaGettingStartedEndpointTeaser) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSuperStackMetaGettingStartedEndpointTeaser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSuperStackMetaGettingStartedEndpointTeaser(val *ServiceSuperStackMetaGettingStartedEndpointTeaser) *NullableServiceSuperStackMetaGettingStartedEndpointTeaser {
	return &NullableServiceSuperStackMetaGettingStartedEndpointTeaser{value: val, isSet: true}
}

func (v NullableServiceSuperStackMetaGettingStartedEndpointTeaser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSuperStackMetaGettingStartedEndpointTeaser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

