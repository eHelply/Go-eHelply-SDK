/*
eHelply SDK - 1.1.120

eHelply SDK for SuperStack Services

API version: 1.1.120
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceSuperStackMetaGettingStarted struct for ServiceSuperStackMetaGettingStarted
type ServiceSuperStackMetaGettingStarted struct {
	Blurb string `json:"blurb"`
	EndpointTeasers []ServiceSuperStackMetaGettingStartedEndpointTeaser `json:"endpoint_teasers"`
}

// NewServiceSuperStackMetaGettingStarted instantiates a new ServiceSuperStackMetaGettingStarted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSuperStackMetaGettingStarted(blurb string, endpointTeasers []ServiceSuperStackMetaGettingStartedEndpointTeaser) *ServiceSuperStackMetaGettingStarted {
	this := ServiceSuperStackMetaGettingStarted{}
	this.Blurb = blurb
	this.EndpointTeasers = endpointTeasers
	return &this
}

// NewServiceSuperStackMetaGettingStartedWithDefaults instantiates a new ServiceSuperStackMetaGettingStarted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSuperStackMetaGettingStartedWithDefaults() *ServiceSuperStackMetaGettingStarted {
	this := ServiceSuperStackMetaGettingStarted{}
	return &this
}

// GetBlurb returns the Blurb field value
func (o *ServiceSuperStackMetaGettingStarted) GetBlurb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blurb
}

// GetBlurbOk returns a tuple with the Blurb field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaGettingStarted) GetBlurbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blurb, true
}

// SetBlurb sets field value
func (o *ServiceSuperStackMetaGettingStarted) SetBlurb(v string) {
	o.Blurb = v
}

// GetEndpointTeasers returns the EndpointTeasers field value
func (o *ServiceSuperStackMetaGettingStarted) GetEndpointTeasers() []ServiceSuperStackMetaGettingStartedEndpointTeaser {
	if o == nil {
		var ret []ServiceSuperStackMetaGettingStartedEndpointTeaser
		return ret
	}

	return o.EndpointTeasers
}

// GetEndpointTeasersOk returns a tuple with the EndpointTeasers field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaGettingStarted) GetEndpointTeasersOk() ([]ServiceSuperStackMetaGettingStartedEndpointTeaser, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndpointTeasers, true
}

// SetEndpointTeasers sets field value
func (o *ServiceSuperStackMetaGettingStarted) SetEndpointTeasers(v []ServiceSuperStackMetaGettingStartedEndpointTeaser) {
	o.EndpointTeasers = v
}

func (o ServiceSuperStackMetaGettingStarted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blurb"] = o.Blurb
	}
	if true {
		toSerialize["endpoint_teasers"] = o.EndpointTeasers
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSuperStackMetaGettingStarted struct {
	value *ServiceSuperStackMetaGettingStarted
	isSet bool
}

func (v NullableServiceSuperStackMetaGettingStarted) Get() *ServiceSuperStackMetaGettingStarted {
	return v.value
}

func (v *NullableServiceSuperStackMetaGettingStarted) Set(val *ServiceSuperStackMetaGettingStarted) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSuperStackMetaGettingStarted) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSuperStackMetaGettingStarted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSuperStackMetaGettingStarted(val *ServiceSuperStackMetaGettingStarted) *NullableServiceSuperStackMetaGettingStarted {
	return &NullableServiceSuperStackMetaGettingStarted{value: val, isSet: true}
}

func (v NullableServiceSuperStackMetaGettingStarted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSuperStackMetaGettingStarted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


