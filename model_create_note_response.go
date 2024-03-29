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

// CreateNoteResponse struct for CreateNoteResponse
type CreateNoteResponse struct {
	Uuid string `json:"uuid"`
}

// NewCreateNoteResponse instantiates a new CreateNoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNoteResponse(uuid string) *CreateNoteResponse {
	this := CreateNoteResponse{}
	this.Uuid = uuid
	return &this
}

// NewCreateNoteResponseWithDefaults instantiates a new CreateNoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNoteResponseWithDefaults() *CreateNoteResponse {
	this := CreateNoteResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *CreateNoteResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CreateNoteResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CreateNoteResponse) SetUuid(v string) {
	o.Uuid = v
}

func (o CreateNoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNoteResponse struct {
	value *CreateNoteResponse
	isSet bool
}

func (v NullableCreateNoteResponse) Get() *CreateNoteResponse {
	return v.value
}

func (v *NullableCreateNoteResponse) Set(val *CreateNoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNoteResponse(val *CreateNoteResponse) *NullableCreateNoteResponse {
	return &NullableCreateNoteResponse{value: val, isSet: true}
}

func (v NullableCreateNoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


