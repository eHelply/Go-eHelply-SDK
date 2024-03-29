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

// CreateTicket struct for CreateTicket
type CreateTicket struct {
	Priority string `json:"priority"`
	Subject string `json:"subject"`
}

// NewCreateTicket instantiates a new CreateTicket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTicket(priority string, subject string) *CreateTicket {
	this := CreateTicket{}
	this.Priority = priority
	this.Subject = subject
	return &this
}

// NewCreateTicketWithDefaults instantiates a new CreateTicket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTicketWithDefaults() *CreateTicket {
	this := CreateTicket{}
	return &this
}

// GetPriority returns the Priority field value
func (o *CreateTicket) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CreateTicket) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *CreateTicket) SetPriority(v string) {
	o.Priority = v
}

// GetSubject returns the Subject field value
func (o *CreateTicket) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CreateTicket) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CreateTicket) SetSubject(v string) {
	o.Subject = v
}

func (o CreateTicket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTicket struct {
	value *CreateTicket
	isSet bool
}

func (v NullableCreateTicket) Get() *CreateTicket {
	return v.value
}

func (v *NullableCreateTicket) Set(val *CreateTicket) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTicket(val *CreateTicket) *NullableCreateTicket {
	return &NullableCreateTicket{value: val, isSet: true}
}

func (v NullableCreateTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


