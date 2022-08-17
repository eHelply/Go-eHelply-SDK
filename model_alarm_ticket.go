/*
eHelply SDK - 1.1.108

eHelply SDK for SuperStack Services

API version: 1.1.108
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AlarmTicket struct for AlarmTicket
type AlarmTicket struct {
	TicketUuid string `json:"ticket_uuid"`
}

// NewAlarmTicket instantiates a new AlarmTicket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmTicket(ticketUuid string) *AlarmTicket {
	this := AlarmTicket{}
	this.TicketUuid = ticketUuid
	return &this
}

// NewAlarmTicketWithDefaults instantiates a new AlarmTicket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmTicketWithDefaults() *AlarmTicket {
	this := AlarmTicket{}
	return &this
}

// GetTicketUuid returns the TicketUuid field value
func (o *AlarmTicket) GetTicketUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketUuid
}

// GetTicketUuidOk returns a tuple with the TicketUuid field value
// and a boolean to check if the value has been set.
func (o *AlarmTicket) GetTicketUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketUuid, true
}

// SetTicketUuid sets field value
func (o *AlarmTicket) SetTicketUuid(v string) {
	o.TicketUuid = v
}

func (o AlarmTicket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ticket_uuid"] = o.TicketUuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmTicket struct {
	value *AlarmTicket
	isSet bool
}

func (v NullableAlarmTicket) Get() *AlarmTicket {
	return v.value
}

func (v *NullableAlarmTicket) Set(val *AlarmTicket) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmTicket(val *AlarmTicket) *NullableAlarmTicket {
	return &NullableAlarmTicket{value: val, isSet: true}
}

func (v NullableAlarmTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


