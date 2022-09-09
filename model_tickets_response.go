/*
eHelply SDK - 1.1.113

eHelply SDK for SuperStack Services

API version: 1.1.113
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// TicketsResponse struct for TicketsResponse
type TicketsResponse struct {
	Subject string `json:"subject"`
	Priority string `json:"priority"`
	TicketId string `json:"ticket_id"`
}

// NewTicketsResponse instantiates a new TicketsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketsResponse(subject string, priority string, ticketId string) *TicketsResponse {
	this := TicketsResponse{}
	this.Subject = subject
	this.Priority = priority
	this.TicketId = ticketId
	return &this
}

// NewTicketsResponseWithDefaults instantiates a new TicketsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketsResponseWithDefaults() *TicketsResponse {
	this := TicketsResponse{}
	return &this
}

// GetSubject returns the Subject field value
func (o *TicketsResponse) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *TicketsResponse) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *TicketsResponse) SetSubject(v string) {
	o.Subject = v
}

// GetPriority returns the Priority field value
func (o *TicketsResponse) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *TicketsResponse) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *TicketsResponse) SetPriority(v string) {
	o.Priority = v
}

// GetTicketId returns the TicketId field value
func (o *TicketsResponse) GetTicketId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketId
}

// GetTicketIdOk returns a tuple with the TicketId field value
// and a boolean to check if the value has been set.
func (o *TicketsResponse) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketId, true
}

// SetTicketId sets field value
func (o *TicketsResponse) SetTicketId(v string) {
	o.TicketId = v
}

func (o TicketsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if true {
		toSerialize["ticket_id"] = o.TicketId
	}
	return json.Marshal(toSerialize)
}

type NullableTicketsResponse struct {
	value *TicketsResponse
	isSet bool
}

func (v NullableTicketsResponse) Get() *TicketsResponse {
	return v.value
}

func (v *NullableTicketsResponse) Set(val *TicketsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketsResponse(val *TicketsResponse) *NullableTicketsResponse {
	return &NullableTicketsResponse{value: val, isSet: true}
}

func (v NullableTicketsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


