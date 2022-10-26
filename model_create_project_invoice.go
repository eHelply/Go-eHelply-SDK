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

// CreateProjectInvoice struct for CreateProjectInvoice
type CreateProjectInvoice struct {
	Year int32 `json:"year"`
	Month int32 `json:"month"`
}

// NewCreateProjectInvoice instantiates a new CreateProjectInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectInvoice(year int32, month int32) *CreateProjectInvoice {
	this := CreateProjectInvoice{}
	this.Year = year
	this.Month = month
	return &this
}

// NewCreateProjectInvoiceWithDefaults instantiates a new CreateProjectInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectInvoiceWithDefaults() *CreateProjectInvoice {
	this := CreateProjectInvoice{}
	return &this
}

// GetYear returns the Year field value
func (o *CreateProjectInvoice) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *CreateProjectInvoice) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *CreateProjectInvoice) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *CreateProjectInvoice) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *CreateProjectInvoice) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *CreateProjectInvoice) SetMonth(v int32) {
	o.Month = v
}

func (o CreateProjectInvoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProjectInvoice struct {
	value *CreateProjectInvoice
	isSet bool
}

func (v NullableCreateProjectInvoice) Get() *CreateProjectInvoice {
	return v.value
}

func (v *NullableCreateProjectInvoice) Set(val *CreateProjectInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectInvoice(val *CreateProjectInvoice) *NullableCreateProjectInvoice {
	return &NullableCreateProjectInvoice{value: val, isSet: true}
}

func (v NullableCreateProjectInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


