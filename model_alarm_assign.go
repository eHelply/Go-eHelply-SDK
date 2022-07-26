/*
eHelply SDK - 1.1.94

eHelply SDK for SuperStack Services

API version: 1.1.94
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AlarmAssign struct for AlarmAssign
type AlarmAssign struct {
	AssigneeUuid string `json:"assignee_uuid"`
}

// NewAlarmAssign instantiates a new AlarmAssign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmAssign(assigneeUuid string) *AlarmAssign {
	this := AlarmAssign{}
	this.AssigneeUuid = assigneeUuid
	return &this
}

// NewAlarmAssignWithDefaults instantiates a new AlarmAssign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmAssignWithDefaults() *AlarmAssign {
	this := AlarmAssign{}
	return &this
}

// GetAssigneeUuid returns the AssigneeUuid field value
func (o *AlarmAssign) GetAssigneeUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssigneeUuid
}

// GetAssigneeUuidOk returns a tuple with the AssigneeUuid field value
// and a boolean to check if the value has been set.
func (o *AlarmAssign) GetAssigneeUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssigneeUuid, true
}

// SetAssigneeUuid sets field value
func (o *AlarmAssign) SetAssigneeUuid(v string) {
	o.AssigneeUuid = v
}

func (o AlarmAssign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assignee_uuid"] = o.AssigneeUuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmAssign struct {
	value *AlarmAssign
	isSet bool
}

func (v NullableAlarmAssign) Get() *AlarmAssign {
	return v.value
}

func (v *NullableAlarmAssign) Set(val *AlarmAssign) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmAssign) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmAssign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmAssign(val *AlarmAssign) *NullableAlarmAssign {
	return &NullableAlarmAssign{value: val, isSet: true}
}

func (v NullableAlarmAssign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmAssign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


