/*
eHelply SDK - 1.1.106

eHelply SDK for SuperStack Services

API version: 1.1.106
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AlarmIgnore struct for AlarmIgnore
type AlarmIgnore struct {
	IgnorerUuid string `json:"ignorer_uuid"`
}

// NewAlarmIgnore instantiates a new AlarmIgnore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmIgnore(ignorerUuid string) *AlarmIgnore {
	this := AlarmIgnore{}
	this.IgnorerUuid = ignorerUuid
	return &this
}

// NewAlarmIgnoreWithDefaults instantiates a new AlarmIgnore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmIgnoreWithDefaults() *AlarmIgnore {
	this := AlarmIgnore{}
	return &this
}

// GetIgnorerUuid returns the IgnorerUuid field value
func (o *AlarmIgnore) GetIgnorerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IgnorerUuid
}

// GetIgnorerUuidOk returns a tuple with the IgnorerUuid field value
// and a boolean to check if the value has been set.
func (o *AlarmIgnore) GetIgnorerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnorerUuid, true
}

// SetIgnorerUuid sets field value
func (o *AlarmIgnore) SetIgnorerUuid(v string) {
	o.IgnorerUuid = v
}

func (o AlarmIgnore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ignorer_uuid"] = o.IgnorerUuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmIgnore struct {
	value *AlarmIgnore
	isSet bool
}

func (v NullableAlarmIgnore) Get() *AlarmIgnore {
	return v.value
}

func (v *NullableAlarmIgnore) Set(val *AlarmIgnore) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmIgnore) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmIgnore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmIgnore(val *AlarmIgnore) *NullableAlarmIgnore {
	return &NullableAlarmIgnore{value: val, isSet: true}
}

func (v NullableAlarmIgnore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmIgnore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


