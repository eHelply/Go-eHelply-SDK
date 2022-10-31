/*
eHelply SDK - 1.1.120

eHelply SDK for SuperStack Services

API version: 1.1.120
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Page Page state
type Page struct {
	Items []interface{} `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// NewPage instantiates a new Page object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPage(items []interface{}, pagination Pagination) *Page {
	this := Page{}
	this.Items = items
	this.Pagination = pagination
	return &this
}

// NewPageWithDefaults instantiates a new Page object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageWithDefaults() *Page {
	this := Page{}
	return &this
}

// GetItems returns the Items field value
func (o *Page) GetItems() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Page) GetItemsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Page) SetItems(v []interface{}) {
	o.Items = v
}

// GetPagination returns the Pagination field value
func (o *Page) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *Page) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *Page) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o Page) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullablePage struct {
	value *Page
	isSet bool
}

func (v NullablePage) Get() *Page {
	return v.value
}

func (v *NullablePage) Set(val *Page) {
	v.value = val
	v.isSet = true
}

func (v NullablePage) IsSet() bool {
	return v.isSet
}

func (v *NullablePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePage(val *Page) *NullablePage {
	return &NullablePage{value: val, isSet: true}
}

func (v NullablePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


