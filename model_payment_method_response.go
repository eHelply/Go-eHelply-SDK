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

// PaymentMethodResponse struct for PaymentMethodResponse
type PaymentMethodResponse struct {
	PaymentId string `json:"payment_id"`
	Last4Digits string `json:"last_4_digits"`
	CardBrand string `json:"card_brand"`
	ProjectUuid string `json:"project_uuid"`
}

// NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodResponse(paymentId string, last4Digits string, cardBrand string, projectUuid string) *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	this.PaymentId = paymentId
	this.Last4Digits = last4Digits
	this.CardBrand = cardBrand
	this.ProjectUuid = projectUuid
	return &this
}

// NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentMethodResponse) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentMethodResponse) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetLast4Digits returns the Last4Digits field value
func (o *PaymentMethodResponse) GetLast4Digits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Last4Digits
}

// GetLast4DigitsOk returns a tuple with the Last4Digits field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetLast4DigitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last4Digits, true
}

// SetLast4Digits sets field value
func (o *PaymentMethodResponse) SetLast4Digits(v string) {
	o.Last4Digits = v
}

// GetCardBrand returns the CardBrand field value
func (o *PaymentMethodResponse) GetCardBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetCardBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardBrand, true
}

// SetCardBrand sets field value
func (o *PaymentMethodResponse) SetCardBrand(v string) {
	o.CardBrand = v
}

// GetProjectUuid returns the ProjectUuid field value
func (o *PaymentMethodResponse) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *PaymentMethodResponse) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

func (o PaymentMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	if true {
		toSerialize["last_4_digits"] = o.Last4Digits
	}
	if true {
		toSerialize["card_brand"] = o.CardBrand
	}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodResponse struct {
	value *PaymentMethodResponse
	isSet bool
}

func (v NullablePaymentMethodResponse) Get() *PaymentMethodResponse {
	return v.value
}

func (v *NullablePaymentMethodResponse) Set(val *PaymentMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodResponse(val *PaymentMethodResponse) *NullablePaymentMethodResponse {
	return &NullablePaymentMethodResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


