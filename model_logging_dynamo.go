/*
eHelply SDK - 1.1.111

eHelply SDK for SuperStack Services

API version: 1.1.111
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// LoggingDynamo struct for LoggingDynamo
type LoggingDynamo struct {
	Service string `json:"service"`
	Time string `json:"time"`
	Log map[string]interface{} `json:"log"`
	Severity string `json:"severity"`
	Subject string `json:"subject"`
}

// NewLoggingDynamo instantiates a new LoggingDynamo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingDynamo(service string, time string, log map[string]interface{}, severity string, subject string) *LoggingDynamo {
	this := LoggingDynamo{}
	this.Service = service
	this.Time = time
	this.Log = log
	this.Severity = severity
	this.Subject = subject
	return &this
}

// NewLoggingDynamoWithDefaults instantiates a new LoggingDynamo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingDynamoWithDefaults() *LoggingDynamo {
	this := LoggingDynamo{}
	return &this
}

// GetService returns the Service field value
func (o *LoggingDynamo) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *LoggingDynamo) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *LoggingDynamo) SetService(v string) {
	o.Service = v
}

// GetTime returns the Time field value
func (o *LoggingDynamo) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *LoggingDynamo) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *LoggingDynamo) SetTime(v string) {
	o.Time = v
}

// GetLog returns the Log field value
func (o *LoggingDynamo) GetLog() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *LoggingDynamo) GetLogOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Log, true
}

// SetLog sets field value
func (o *LoggingDynamo) SetLog(v map[string]interface{}) {
	o.Log = v
}

// GetSeverity returns the Severity field value
func (o *LoggingDynamo) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *LoggingDynamo) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *LoggingDynamo) SetSeverity(v string) {
	o.Severity = v
}

// GetSubject returns the Subject field value
func (o *LoggingDynamo) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *LoggingDynamo) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *LoggingDynamo) SetSubject(v string) {
	o.Subject = v
}

func (o LoggingDynamo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["log"] = o.Log
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableLoggingDynamo struct {
	value *LoggingDynamo
	isSet bool
}

func (v NullableLoggingDynamo) Get() *LoggingDynamo {
	return v.value
}

func (v *NullableLoggingDynamo) Set(val *LoggingDynamo) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingDynamo) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingDynamo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingDynamo(val *LoggingDynamo) *NullableLoggingDynamo {
	return &NullableLoggingDynamo{value: val, isSet: true}
}

func (v NullableLoggingDynamo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingDynamo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


