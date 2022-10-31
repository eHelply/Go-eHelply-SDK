/*
eHelply SDK - 1.1.120

eHelply SDK for SuperStack Services

API version: 1.1.120
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// NoteDynamoHistoryResponse A note from Dynamo DB including n amount of version history of that note
type NoteDynamoHistoryResponse struct {
	Uuid string `json:"uuid"`
	Content string `json:"content"`
	Time string `json:"time"`
	Meta NoteMeta `json:"meta"`
	History []NoteDynamoResponse `json:"history,omitempty"`
}

// NewNoteDynamoHistoryResponse instantiates a new NoteDynamoHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteDynamoHistoryResponse(uuid string, content string, time string, meta NoteMeta) *NoteDynamoHistoryResponse {
	this := NoteDynamoHistoryResponse{}
	this.Uuid = uuid
	this.Content = content
	this.Time = time
	this.Meta = meta
	return &this
}

// NewNoteDynamoHistoryResponseWithDefaults instantiates a new NoteDynamoHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteDynamoHistoryResponseWithDefaults() *NoteDynamoHistoryResponse {
	this := NoteDynamoHistoryResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *NoteDynamoHistoryResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *NoteDynamoHistoryResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *NoteDynamoHistoryResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetContent returns the Content field value
func (o *NoteDynamoHistoryResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *NoteDynamoHistoryResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *NoteDynamoHistoryResponse) SetContent(v string) {
	o.Content = v
}

// GetTime returns the Time field value
func (o *NoteDynamoHistoryResponse) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *NoteDynamoHistoryResponse) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *NoteDynamoHistoryResponse) SetTime(v string) {
	o.Time = v
}

// GetMeta returns the Meta field value
func (o *NoteDynamoHistoryResponse) GetMeta() NoteMeta {
	if o == nil {
		var ret NoteMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *NoteDynamoHistoryResponse) GetMetaOk() (*NoteMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *NoteDynamoHistoryResponse) SetMeta(v NoteMeta) {
	o.Meta = v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *NoteDynamoHistoryResponse) GetHistory() []NoteDynamoResponse {
	if o == nil || o.History == nil {
		var ret []NoteDynamoResponse
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteDynamoHistoryResponse) GetHistoryOk() ([]NoteDynamoResponse, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *NoteDynamoHistoryResponse) HasHistory() bool {
	if o != nil && o.History != nil {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []NoteDynamoResponse and assigns it to the History field.
func (o *NoteDynamoHistoryResponse) SetHistory(v []NoteDynamoResponse) {
	o.History = v
}

func (o NoteDynamoHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	return json.Marshal(toSerialize)
}

type NullableNoteDynamoHistoryResponse struct {
	value *NoteDynamoHistoryResponse
	isSet bool
}

func (v NullableNoteDynamoHistoryResponse) Get() *NoteDynamoHistoryResponse {
	return v.value
}

func (v *NullableNoteDynamoHistoryResponse) Set(val *NoteDynamoHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteDynamoHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteDynamoHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteDynamoHistoryResponse(val *NoteDynamoHistoryResponse) *NullableNoteDynamoHistoryResponse {
	return &NullableNoteDynamoHistoryResponse{value: val, isSet: true}
}

func (v NullableNoteDynamoHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteDynamoHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


