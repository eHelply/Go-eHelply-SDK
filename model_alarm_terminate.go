/*
eHelply SDK - 1.1.113

eHelply SDK for SuperStack Services

API version: 1.1.113
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AlarmTerminate struct for AlarmTerminate
type AlarmTerminate struct {
	TerminaterUuid string `json:"terminater_uuid"`
}

// NewAlarmTerminate instantiates a new AlarmTerminate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmTerminate(terminaterUuid string) *AlarmTerminate {
	this := AlarmTerminate{}
	this.TerminaterUuid = terminaterUuid
	return &this
}

// NewAlarmTerminateWithDefaults instantiates a new AlarmTerminate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmTerminateWithDefaults() *AlarmTerminate {
	this := AlarmTerminate{}
	return &this
}

// GetTerminaterUuid returns the TerminaterUuid field value
func (o *AlarmTerminate) GetTerminaterUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TerminaterUuid
}

// GetTerminaterUuidOk returns a tuple with the TerminaterUuid field value
// and a boolean to check if the value has been set.
func (o *AlarmTerminate) GetTerminaterUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminaterUuid, true
}

// SetTerminaterUuid sets field value
func (o *AlarmTerminate) SetTerminaterUuid(v string) {
	o.TerminaterUuid = v
}

func (o AlarmTerminate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["terminater_uuid"] = o.TerminaterUuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmTerminate struct {
	value *AlarmTerminate
	isSet bool
}

func (v NullableAlarmTerminate) Get() *AlarmTerminate {
	return v.value
}

func (v *NullableAlarmTerminate) Set(val *AlarmTerminate) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmTerminate) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmTerminate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmTerminate(val *AlarmTerminate) *NullableAlarmTerminate {
	return &NullableAlarmTerminate{value: val, isSet: true}
}

func (v NullableAlarmTerminate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmTerminate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


