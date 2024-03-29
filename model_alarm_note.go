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

// AlarmNote struct for AlarmNote
type AlarmNote struct {
	AuthorUuid string `json:"author_uuid"`
	Message string `json:"message"`
}

// NewAlarmNote instantiates a new AlarmNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmNote(authorUuid string, message string) *AlarmNote {
	this := AlarmNote{}
	this.AuthorUuid = authorUuid
	this.Message = message
	return &this
}

// NewAlarmNoteWithDefaults instantiates a new AlarmNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmNoteWithDefaults() *AlarmNote {
	this := AlarmNote{}
	return &this
}

// GetAuthorUuid returns the AuthorUuid field value
func (o *AlarmNote) GetAuthorUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorUuid
}

// GetAuthorUuidOk returns a tuple with the AuthorUuid field value
// and a boolean to check if the value has been set.
func (o *AlarmNote) GetAuthorUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorUuid, true
}

// SetAuthorUuid sets field value
func (o *AlarmNote) SetAuthorUuid(v string) {
	o.AuthorUuid = v
}

// GetMessage returns the Message field value
func (o *AlarmNote) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlarmNote) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlarmNote) SetMessage(v string) {
	o.Message = v
}

func (o AlarmNote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["author_uuid"] = o.AuthorUuid
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmNote struct {
	value *AlarmNote
	isSet bool
}

func (v NullableAlarmNote) Get() *AlarmNote {
	return v.value
}

func (v *NullableAlarmNote) Set(val *AlarmNote) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmNote) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmNote(val *AlarmNote) *NullableAlarmNote {
	return &NullableAlarmNote{value: val, isSet: true}
}

func (v NullableAlarmNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


