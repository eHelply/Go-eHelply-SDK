/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// LineItem struct for LineItem
type LineItem struct {
	Name string `json:"name"`
	Quantity int32 `json:"quantity"`
	UnitPrice int32 `json:"unit_price"`
	Total int32 `json:"total"`
}

// NewLineItem instantiates a new LineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItem(name string, quantity int32, unitPrice int32, total int32) *LineItem {
	this := LineItem{}
	this.Name = name
	this.Quantity = quantity
	this.UnitPrice = unitPrice
	this.Total = total
	return &this
}

// NewLineItemWithDefaults instantiates a new LineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemWithDefaults() *LineItem {
	this := LineItem{}
	return &this
}

// GetName returns the Name field value
func (o *LineItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LineItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LineItem) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *LineItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *LineItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *LineItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *LineItem) GetUnitPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *LineItem) GetUnitPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *LineItem) SetUnitPrice(v int32) {
	o.UnitPrice = v
}

// GetTotal returns the Total field value
func (o *LineItem) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *LineItem) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *LineItem) SetTotal(v int32) {
	o.Total = v
}

func (o LineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["unit_price"] = o.UnitPrice
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableLineItem struct {
	value *LineItem
	isSet bool
}

func (v NullableLineItem) Get() *LineItem {
	return v.value
}

func (v *NullableLineItem) Set(val *LineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItem(val *LineItem) *NullableLineItem {
	return &NullableLineItem{value: val, isSet: true}
}

func (v NullableLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


