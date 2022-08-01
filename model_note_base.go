/*
eHelply SDK - 1.1.100

eHelply SDK for SuperStack Services

API version: 1.1.100
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// NoteBase Note This is used as the payload to endpoints
type NoteBase struct {
	Content string `json:"content"`
	Time string `json:"time"`
	Meta NoteMeta `json:"meta"`
}

// NewNoteBase instantiates a new NoteBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteBase(content string, time string, meta NoteMeta) *NoteBase {
	this := NoteBase{}
	this.Content = content
	this.Time = time
	this.Meta = meta
	return &this
}

// NewNoteBaseWithDefaults instantiates a new NoteBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteBaseWithDefaults() *NoteBase {
	this := NoteBase{}
	return &this
}

// GetContent returns the Content field value
func (o *NoteBase) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *NoteBase) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *NoteBase) SetContent(v string) {
	o.Content = v
}

// GetTime returns the Time field value
func (o *NoteBase) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *NoteBase) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *NoteBase) SetTime(v string) {
	o.Time = v
}

// GetMeta returns the Meta field value
func (o *NoteBase) GetMeta() NoteMeta {
	if o == nil {
		var ret NoteMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *NoteBase) GetMetaOk() (*NoteMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *NoteBase) SetMeta(v NoteMeta) {
	o.Meta = v
}

func (o NoteBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableNoteBase struct {
	value *NoteBase
	isSet bool
}

func (v NullableNoteBase) Get() *NoteBase {
	return v.value
}

func (v *NullableNoteBase) Set(val *NoteBase) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteBase) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteBase(val *NoteBase) *NullableNoteBase {
	return &NullableNoteBase{value: val, isSet: true}
}

func (v NullableNoteBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


