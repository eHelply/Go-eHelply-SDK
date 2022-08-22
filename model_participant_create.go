/*
eHelply SDK - 1.1.103

eHelply SDK for SuperStack Services

API version: 1.1.103
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ParticipantCreate Contains field for when we create a participant only
type ParticipantCreate struct {
	// Dictionary containing any data you would like
	Meta map[string]interface{} `json:"meta,omitempty"`
	UserUuid *string `json:"user_uuid,omitempty"`
}

// NewParticipantCreate instantiates a new ParticipantCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipantCreate() *ParticipantCreate {
	this := ParticipantCreate{}
	return &this
}

// NewParticipantCreateWithDefaults instantiates a new ParticipantCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantCreateWithDefaults() *ParticipantCreate {
	this := ParticipantCreate{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ParticipantCreate) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantCreate) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ParticipantCreate) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ParticipantCreate) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *ParticipantCreate) GetUserUuid() string {
	if o == nil || o.UserUuid == nil {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantCreate) GetUserUuidOk() (*string, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *ParticipantCreate) HasUserUuid() bool {
	if o != nil && o.UserUuid != nil {
		return true
	}

	return false
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *ParticipantCreate) SetUserUuid(v string) {
	o.UserUuid = &v
}

func (o ParticipantCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.UserUuid != nil {
		toSerialize["user_uuid"] = o.UserUuid
	}
	return json.Marshal(toSerialize)
}

type NullableParticipantCreate struct {
	value *ParticipantCreate
	isSet bool
}

func (v NullableParticipantCreate) Get() *ParticipantCreate {
	return v.value
}

func (v *NullableParticipantCreate) Set(val *ParticipantCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantCreate(val *ParticipantCreate) *NullableParticipantCreate {
	return &NullableParticipantCreate{value: val, isSet: true}
}

func (v NullableParticipantCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


