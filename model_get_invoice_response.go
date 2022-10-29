/*
eHelply SDK - 1.1.119

eHelply SDK for SuperStack Services

API version: 1.1.119
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// GetInvoiceResponse struct for GetInvoiceResponse
type GetInvoiceResponse struct {
	InvoiceUuid string `json:"invoice_uuid"`
	LineItems []LineItem `json:"line_items"`
	Subtotal int32 `json:"subtotal"`
	Discounts []Discount `json:"discounts"`
	Taxes []Tax `json:"taxes"`
	Total int32 `json:"total"`
	Notes []Note `json:"notes"`
	Paid *bool `json:"paid,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdateAt string `json:"update_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewGetInvoiceResponse instantiates a new GetInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoiceResponse(invoiceUuid string, lineItems []LineItem, subtotal int32, discounts []Discount, taxes []Tax, total int32, notes []Note, createdAt string, updateAt string) *GetInvoiceResponse {
	this := GetInvoiceResponse{}
	this.InvoiceUuid = invoiceUuid
	this.LineItems = lineItems
	this.Subtotal = subtotal
	this.Discounts = discounts
	this.Taxes = taxes
	this.Total = total
	this.Notes = notes
	var paid bool = false
	this.Paid = &paid
	this.CreatedAt = createdAt
	this.UpdateAt = updateAt
	return &this
}

// NewGetInvoiceResponseWithDefaults instantiates a new GetInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoiceResponseWithDefaults() *GetInvoiceResponse {
	this := GetInvoiceResponse{}
	var paid bool = false
	this.Paid = &paid
	return &this
}

// GetInvoiceUuid returns the InvoiceUuid field value
func (o *GetInvoiceResponse) GetInvoiceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceUuid
}

// GetInvoiceUuidOk returns a tuple with the InvoiceUuid field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetInvoiceUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceUuid, true
}

// SetInvoiceUuid sets field value
func (o *GetInvoiceResponse) SetInvoiceUuid(v string) {
	o.InvoiceUuid = v
}

// GetLineItems returns the LineItems field value
func (o *GetInvoiceResponse) GetLineItems() []LineItem {
	if o == nil {
		var ret []LineItem
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *GetInvoiceResponse) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetSubtotal returns the Subtotal field value
func (o *GetInvoiceResponse) GetSubtotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetSubtotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *GetInvoiceResponse) SetSubtotal(v int32) {
	o.Subtotal = v
}

// GetDiscounts returns the Discounts field value
func (o *GetInvoiceResponse) GetDiscounts() []Discount {
	if o == nil {
		var ret []Discount
		return ret
	}

	return o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetDiscountsOk() ([]Discount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discounts, true
}

// SetDiscounts sets field value
func (o *GetInvoiceResponse) SetDiscounts(v []Discount) {
	o.Discounts = v
}

// GetTaxes returns the Taxes field value
func (o *GetInvoiceResponse) GetTaxes() []Tax {
	if o == nil {
		var ret []Tax
		return ret
	}

	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetTaxesOk() ([]Tax, bool) {
	if o == nil {
		return nil, false
	}
	return o.Taxes, true
}

// SetTaxes sets field value
func (o *GetInvoiceResponse) SetTaxes(v []Tax) {
	o.Taxes = v
}

// GetTotal returns the Total field value
func (o *GetInvoiceResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *GetInvoiceResponse) SetTotal(v int32) {
	o.Total = v
}

// GetNotes returns the Notes field value
func (o *GetInvoiceResponse) GetNotes() []Note {
	if o == nil {
		var ret []Note
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetNotesOk() ([]Note, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *GetInvoiceResponse) SetNotes(v []Note) {
	o.Notes = v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *GetInvoiceResponse) GetPaid() bool {
	if o == nil || o.Paid == nil {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetPaidOk() (*bool, bool) {
	if o == nil || o.Paid == nil {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *GetInvoiceResponse) HasPaid() bool {
	if o != nil && o.Paid != nil {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *GetInvoiceResponse) SetPaid(v bool) {
	o.Paid = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetInvoiceResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetInvoiceResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdateAt returns the UpdateAt field value
func (o *GetInvoiceResponse) GetUpdateAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateAt
}

// GetUpdateAtOk returns a tuple with the UpdateAt field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetUpdateAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateAt, true
}

// SetUpdateAt sets field value
func (o *GetInvoiceResponse) SetUpdateAt(v string) {
	o.UpdateAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *GetInvoiceResponse) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *GetInvoiceResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *GetInvoiceResponse) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o GetInvoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["invoice_uuid"] = o.InvoiceUuid
	}
	if true {
		toSerialize["line_items"] = o.LineItems
	}
	if true {
		toSerialize["subtotal"] = o.Subtotal
	}
	if true {
		toSerialize["discounts"] = o.Discounts
	}
	if true {
		toSerialize["taxes"] = o.Taxes
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["notes"] = o.Notes
	}
	if o.Paid != nil {
		toSerialize["paid"] = o.Paid
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["update_at"] = o.UpdateAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGetInvoiceResponse struct {
	value *GetInvoiceResponse
	isSet bool
}

func (v NullableGetInvoiceResponse) Get() *GetInvoiceResponse {
	return v.value
}

func (v *NullableGetInvoiceResponse) Set(val *GetInvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoiceResponse(val *GetInvoiceResponse) *NullableGetInvoiceResponse {
	return &NullableGetInvoiceResponse{value: val, isSet: true}
}

func (v NullableGetInvoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


