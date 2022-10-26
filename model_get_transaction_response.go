/*
eHelply SDK - 1.1.117

eHelply SDK for SuperStack Services

API version: 1.1.117
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// GetTransactionResponse struct for GetTransactionResponse
type GetTransactionResponse struct {
	Invoice *GetInvoiceResponse `json:"invoice,omitempty"`
	TransactionUuid string `json:"transaction_uuid"`
	StripeId string `json:"stripe_id"`
	Credit int32 `json:"credit"`
	Debit int32 `json:"debit"`
	CreatedAt string `json:"created_at"`
}

// NewGetTransactionResponse instantiates a new GetTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionResponse(transactionUuid string, stripeId string, credit int32, debit int32, createdAt string) *GetTransactionResponse {
	this := GetTransactionResponse{}
	this.TransactionUuid = transactionUuid
	this.StripeId = stripeId
	this.Credit = credit
	this.Debit = debit
	this.CreatedAt = createdAt
	return &this
}

// NewGetTransactionResponseWithDefaults instantiates a new GetTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionResponseWithDefaults() *GetTransactionResponse {
	this := GetTransactionResponse{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *GetTransactionResponse) GetInvoice() GetInvoiceResponse {
	if o == nil || o.Invoice == nil {
		var ret GetInvoiceResponse
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransactionResponse) GetInvoiceOk() (*GetInvoiceResponse, bool) {
	if o == nil || o.Invoice == nil {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *GetTransactionResponse) HasInvoice() bool {
	if o != nil && o.Invoice != nil {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given GetInvoiceResponse and assigns it to the Invoice field.
func (o *GetTransactionResponse) SetInvoice(v GetInvoiceResponse) {
	o.Invoice = &v
}

// GetTransactionUuid returns the TransactionUuid field value
func (o *GetTransactionResponse) GetTransactionUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionUuid
}

// GetTransactionUuidOk returns a tuple with the TransactionUuid field value
// and a boolean to check if the value has been set.
func (o *GetTransactionResponse) GetTransactionUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionUuid, true
}

// SetTransactionUuid sets field value
func (o *GetTransactionResponse) SetTransactionUuid(v string) {
	o.TransactionUuid = v
}

// GetStripeId returns the StripeId field value
func (o *GetTransactionResponse) GetStripeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StripeId
}

// GetStripeIdOk returns a tuple with the StripeId field value
// and a boolean to check if the value has been set.
func (o *GetTransactionResponse) GetStripeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StripeId, true
}

// SetStripeId sets field value
func (o *GetTransactionResponse) SetStripeId(v string) {
	o.StripeId = v
}

// GetCredit returns the Credit field value
func (o *GetTransactionResponse) GetCredit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Credit
}

// GetCreditOk returns a tuple with the Credit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionResponse) GetCreditOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credit, true
}

// SetCredit sets field value
func (o *GetTransactionResponse) SetCredit(v int32) {
	o.Credit = v
}

// GetDebit returns the Debit field value
func (o *GetTransactionResponse) GetDebit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Debit
}

// GetDebitOk returns a tuple with the Debit field value
// and a boolean to check if the value has been set.
func (o *GetTransactionResponse) GetDebitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Debit, true
}

// SetDebit sets field value
func (o *GetTransactionResponse) SetDebit(v int32) {
	o.Debit = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetTransactionResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetTransactionResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetTransactionResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o GetTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Invoice != nil {
		toSerialize["invoice"] = o.Invoice
	}
	if true {
		toSerialize["transaction_uuid"] = o.TransactionUuid
	}
	if true {
		toSerialize["stripe_id"] = o.StripeId
	}
	if true {
		toSerialize["credit"] = o.Credit
	}
	if true {
		toSerialize["debit"] = o.Debit
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransactionResponse struct {
	value *GetTransactionResponse
	isSet bool
}

func (v NullableGetTransactionResponse) Get() *GetTransactionResponse {
	return v.value
}

func (v *NullableGetTransactionResponse) Set(val *GetTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionResponse(val *GetTransactionResponse) *NullableGetTransactionResponse {
	return &NullableGetTransactionResponse{value: val, isSet: true}
}

func (v NullableGetTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


