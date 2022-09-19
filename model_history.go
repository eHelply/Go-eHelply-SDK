/*
eHelply SDK - 1.1.107

eHelply SDK for SuperStack Services

API version: 1.1.107
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// History struct for History
type History struct {
	Year int32 `json:"year"`
	Month int32 `json:"month"`
}

// NewHistory instantiates a new History object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistory(year int32, month int32) *History {
	this := History{}
	this.Year = year
	this.Month = month
	return &this
}

// NewHistoryWithDefaults instantiates a new History object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryWithDefaults() *History {
	this := History{}
	return &this
}

// GetYear returns the Year field value
func (o *History) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *History) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *History) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *History) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *History) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *History) SetMonth(v int32) {
	o.Month = v
}

func (o History) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	return json.Marshal(toSerialize)
}

type NullableHistory struct {
	value *History
	isSet bool
}

func (v NullableHistory) Get() *History {
	return v.value
}

func (v *NullableHistory) Set(val *History) {
	v.value = val
	v.isSet = true
}

func (v NullableHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistory(val *History) *NullableHistory {
	return &NullableHistory{value: val, isSet: true}
}

func (v NullableHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


