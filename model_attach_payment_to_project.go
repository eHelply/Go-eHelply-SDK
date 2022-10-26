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

// AttachPaymentToProject struct for AttachPaymentToProject
type AttachPaymentToProject struct {
	PaymentType string `json:"payment_type"`
	Number string `json:"number"`
	ExpMonth int32 `json:"exp_month"`
	ExpYear int32 `json:"exp_year"`
	Cvc string `json:"cvc"`
}

// NewAttachPaymentToProject instantiates a new AttachPaymentToProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachPaymentToProject(paymentType string, number string, expMonth int32, expYear int32, cvc string) *AttachPaymentToProject {
	this := AttachPaymentToProject{}
	this.PaymentType = paymentType
	this.Number = number
	this.ExpMonth = expMonth
	this.ExpYear = expYear
	this.Cvc = cvc
	return &this
}

// NewAttachPaymentToProjectWithDefaults instantiates a new AttachPaymentToProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachPaymentToProjectWithDefaults() *AttachPaymentToProject {
	this := AttachPaymentToProject{}
	return &this
}

// GetPaymentType returns the PaymentType field value
func (o *AttachPaymentToProject) GetPaymentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value
// and a boolean to check if the value has been set.
func (o *AttachPaymentToProject) GetPaymentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentType, true
}

// SetPaymentType sets field value
func (o *AttachPaymentToProject) SetPaymentType(v string) {
	o.PaymentType = v
}

// GetNumber returns the Number field value
func (o *AttachPaymentToProject) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *AttachPaymentToProject) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *AttachPaymentToProject) SetNumber(v string) {
	o.Number = v
}

// GetExpMonth returns the ExpMonth field value
func (o *AttachPaymentToProject) GetExpMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value
// and a boolean to check if the value has been set.
func (o *AttachPaymentToProject) GetExpMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpMonth, true
}

// SetExpMonth sets field value
func (o *AttachPaymentToProject) SetExpMonth(v int32) {
	o.ExpMonth = v
}

// GetExpYear returns the ExpYear field value
func (o *AttachPaymentToProject) GetExpYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value
// and a boolean to check if the value has been set.
func (o *AttachPaymentToProject) GetExpYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpYear, true
}

// SetExpYear sets field value
func (o *AttachPaymentToProject) SetExpYear(v int32) {
	o.ExpYear = v
}

// GetCvc returns the Cvc field value
func (o *AttachPaymentToProject) GetCvc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value
// and a boolean to check if the value has been set.
func (o *AttachPaymentToProject) GetCvcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvc, true
}

// SetCvc sets field value
func (o *AttachPaymentToProject) SetCvc(v string) {
	o.Cvc = v
}

func (o AttachPaymentToProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_type"] = o.PaymentType
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["exp_month"] = o.ExpMonth
	}
	if true {
		toSerialize["exp_year"] = o.ExpYear
	}
	if true {
		toSerialize["cvc"] = o.Cvc
	}
	return json.Marshal(toSerialize)
}

type NullableAttachPaymentToProject struct {
	value *AttachPaymentToProject
	isSet bool
}

func (v NullableAttachPaymentToProject) Get() *AttachPaymentToProject {
	return v.value
}

func (v *NullableAttachPaymentToProject) Set(val *AttachPaymentToProject) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachPaymentToProject) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachPaymentToProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachPaymentToProject(val *AttachPaymentToProject) *NullableAttachPaymentToProject {
	return &NullableAttachPaymentToProject{value: val, isSet: true}
}

func (v NullableAttachPaymentToProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachPaymentToProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


