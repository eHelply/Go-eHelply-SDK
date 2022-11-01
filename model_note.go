/*
eHelply SDK - 1.1.117

eHelply SDK for SuperStack Services

API version: 1.1.117
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Note struct for Note
type Note struct {
	ParticipantUuid string `json:"participant_uuid"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// NewNote instantiates a new Note object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNote(participantUuid string, content string, createdAt string) *Note {
	this := Note{}
	this.ParticipantUuid = participantUuid
	this.Content = content
	this.CreatedAt = createdAt
	return &this
}

// NewNoteWithDefaults instantiates a new Note object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteWithDefaults() *Note {
	this := Note{}
	return &this
}

// GetParticipantUuid returns the ParticipantUuid field value
func (o *Note) GetParticipantUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParticipantUuid
}

// GetParticipantUuidOk returns a tuple with the ParticipantUuid field value
// and a boolean to check if the value has been set.
func (o *Note) GetParticipantUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParticipantUuid, true
}

// SetParticipantUuid sets field value
func (o *Note) SetParticipantUuid(v string) {
	o.ParticipantUuid = v
}

// GetContent returns the Content field value
func (o *Note) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Note) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Note) SetContent(v string) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Note) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Note) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Note) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o Note) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["participant_uuid"] = o.ParticipantUuid
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableNote struct {
	value *Note
	isSet bool
}

func (v NullableNote) Get() *Note {
	return v.value
}

func (v *NullableNote) Set(val *Note) {
	v.value = val
	v.isSet = true
}

func (v NullableNote) IsSet() bool {
	return v.isSet
}

func (v *NullableNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNote(val *Note) *NullableNote {
	return &NullableNote{value: val, isSet: true}
}

func (v NullableNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


