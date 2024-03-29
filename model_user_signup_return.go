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

// UserSignupReturn Default participant UUID
type UserSignupReturn struct {
	ParticipantUuid string `json:"participant_uuid"`
}

// NewUserSignupReturn instantiates a new UserSignupReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSignupReturn(participantUuid string) *UserSignupReturn {
	this := UserSignupReturn{}
	this.ParticipantUuid = participantUuid
	return &this
}

// NewUserSignupReturnWithDefaults instantiates a new UserSignupReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSignupReturnWithDefaults() *UserSignupReturn {
	this := UserSignupReturn{}
	return &this
}

// GetParticipantUuid returns the ParticipantUuid field value
func (o *UserSignupReturn) GetParticipantUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParticipantUuid
}

// GetParticipantUuidOk returns a tuple with the ParticipantUuid field value
// and a boolean to check if the value has been set.
func (o *UserSignupReturn) GetParticipantUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParticipantUuid, true
}

// SetParticipantUuid sets field value
func (o *UserSignupReturn) SetParticipantUuid(v string) {
	o.ParticipantUuid = v
}

func (o UserSignupReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["participant_uuid"] = o.ParticipantUuid
	}
	return json.Marshal(toSerialize)
}

type NullableUserSignupReturn struct {
	value *UserSignupReturn
	isSet bool
}

func (v NullableUserSignupReturn) Get() *UserSignupReturn {
	return v.value
}

func (v *NullableUserSignupReturn) Set(val *UserSignupReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSignupReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSignupReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSignupReturn(val *UserSignupReturn) *NullableUserSignupReturn {
	return &NullableUserSignupReturn{value: val, isSet: true}
}

func (v NullableUserSignupReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSignupReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


