/*
eHelply SDK - 1.1.114

eHelply SDK for SuperStack Services

API version: 1.1.114
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AlarmAcknowledge struct for AlarmAcknowledge
type AlarmAcknowledge struct {
	AcknowledgerUuid string `json:"acknowledger_uuid"`
}

// NewAlarmAcknowledge instantiates a new AlarmAcknowledge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmAcknowledge(acknowledgerUuid string) *AlarmAcknowledge {
	this := AlarmAcknowledge{}
	this.AcknowledgerUuid = acknowledgerUuid
	return &this
}

// NewAlarmAcknowledgeWithDefaults instantiates a new AlarmAcknowledge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmAcknowledgeWithDefaults() *AlarmAcknowledge {
	this := AlarmAcknowledge{}
	return &this
}

// GetAcknowledgerUuid returns the AcknowledgerUuid field value
func (o *AlarmAcknowledge) GetAcknowledgerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcknowledgerUuid
}

// GetAcknowledgerUuidOk returns a tuple with the AcknowledgerUuid field value
// and a boolean to check if the value has been set.
func (o *AlarmAcknowledge) GetAcknowledgerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcknowledgerUuid, true
}

// SetAcknowledgerUuid sets field value
func (o *AlarmAcknowledge) SetAcknowledgerUuid(v string) {
	o.AcknowledgerUuid = v
}

func (o AlarmAcknowledge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["acknowledger_uuid"] = o.AcknowledgerUuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmAcknowledge struct {
	value *AlarmAcknowledge
	isSet bool
}

func (v NullableAlarmAcknowledge) Get() *AlarmAcknowledge {
	return v.value
}

func (v *NullableAlarmAcknowledge) Set(val *AlarmAcknowledge) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmAcknowledge) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmAcknowledge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmAcknowledge(val *AlarmAcknowledge) *NullableAlarmAcknowledge {
	return &NullableAlarmAcknowledge{value: val, isSet: true}
}

func (v NullableAlarmAcknowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmAcknowledge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


