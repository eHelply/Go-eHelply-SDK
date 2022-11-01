/*
eHelply SDK - 1.1.121

eHelply SDK for SuperStack Services

API version: 1.1.121
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// TicketResponse struct for TicketResponse
type TicketResponse struct {
	TicketId string `json:"ticket_id"`
	ContactId string `json:"contact_id"`
	Subject string `json:"subject"`
	Priority string `json:"priority"`
}

// NewTicketResponse instantiates a new TicketResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketResponse(ticketId string, contactId string, subject string, priority string) *TicketResponse {
	this := TicketResponse{}
	this.TicketId = ticketId
	this.ContactId = contactId
	this.Subject = subject
	this.Priority = priority
	return &this
}

// NewTicketResponseWithDefaults instantiates a new TicketResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketResponseWithDefaults() *TicketResponse {
	this := TicketResponse{}
	return &this
}

// GetTicketId returns the TicketId field value
func (o *TicketResponse) GetTicketId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketId
}

// GetTicketIdOk returns a tuple with the TicketId field value
// and a boolean to check if the value has been set.
func (o *TicketResponse) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketId, true
}

// SetTicketId sets field value
func (o *TicketResponse) SetTicketId(v string) {
	o.TicketId = v
}

// GetContactId returns the ContactId field value
func (o *TicketResponse) GetContactId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value
// and a boolean to check if the value has been set.
func (o *TicketResponse) GetContactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactId, true
}

// SetContactId sets field value
func (o *TicketResponse) SetContactId(v string) {
	o.ContactId = v
}

// GetSubject returns the Subject field value
func (o *TicketResponse) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *TicketResponse) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *TicketResponse) SetSubject(v string) {
	o.Subject = v
}

// GetPriority returns the Priority field value
func (o *TicketResponse) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *TicketResponse) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *TicketResponse) SetPriority(v string) {
	o.Priority = v
}

func (o TicketResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ticket_id"] = o.TicketId
	}
	if true {
		toSerialize["contact_id"] = o.ContactId
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableTicketResponse struct {
	value *TicketResponse
	isSet bool
}

func (v NullableTicketResponse) Get() *TicketResponse {
	return v.value
}

func (v *NullableTicketResponse) Set(val *TicketResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketResponse(val *TicketResponse) *NullableTicketResponse {
	return &NullableTicketResponse{value: val, isSet: true}
}

func (v NullableTicketResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


