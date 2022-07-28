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

// ParticipantUpdate Fields for updating participants including the participant uuid
type ParticipantUpdate struct {
	// Dictionary containing any data you would like
	Meta map[string]interface{} `json:"meta,omitempty"`
	UserUuid *string `json:"user_uuid,omitempty"`
	Uuid string `json:"uuid"`
}

// NewParticipantUpdate instantiates a new ParticipantUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipantUpdate(uuid string) *ParticipantUpdate {
	this := ParticipantUpdate{}
	this.Uuid = uuid
	return &this
}

// NewParticipantUpdateWithDefaults instantiates a new ParticipantUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantUpdateWithDefaults() *ParticipantUpdate {
	this := ParticipantUpdate{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ParticipantUpdate) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUpdate) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ParticipantUpdate) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *ParticipantUpdate) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *ParticipantUpdate) GetUserUuid() string {
	if o == nil || o.UserUuid == nil {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUpdate) GetUserUuidOk() (*string, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *ParticipantUpdate) HasUserUuid() bool {
	if o != nil && o.UserUuid != nil {
		return true
	}

	return false
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *ParticipantUpdate) SetUserUuid(v string) {
	o.UserUuid = &v
}

// GetUuid returns the Uuid field value
func (o *ParticipantUpdate) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ParticipantUpdate) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ParticipantUpdate) SetUuid(v string) {
	o.Uuid = v
}

func (o ParticipantUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.UserUuid != nil {
		toSerialize["user_uuid"] = o.UserUuid
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableParticipantUpdate struct {
	value *ParticipantUpdate
	isSet bool
}

func (v NullableParticipantUpdate) Get() *ParticipantUpdate {
	return v.value
}

func (v *NullableParticipantUpdate) Set(val *ParticipantUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantUpdate(val *ParticipantUpdate) *NullableParticipantUpdate {
	return &NullableParticipantUpdate{value: val, isSet: true}
}

func (v NullableParticipantUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


