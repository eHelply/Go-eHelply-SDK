/*
eHelply SDK - 1.1.99

eHelply SDK for SuperStack Services

API version: 1.1.99
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsUsageTypeUnitPrice struct for ProjectsUsageTypeUnitPrice
type ProjectsUsageTypeUnitPrice struct {
	MinQuantity int32 `json:"min_quantity"`
	MaxQuantity int32 `json:"max_quantity"`
	UnitPrice int32 `json:"unit_price"`
}

// NewProjectsUsageTypeUnitPrice instantiates a new ProjectsUsageTypeUnitPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsUsageTypeUnitPrice(minQuantity int32, maxQuantity int32, unitPrice int32) *ProjectsUsageTypeUnitPrice {
	this := ProjectsUsageTypeUnitPrice{}
	this.MinQuantity = minQuantity
	this.MaxQuantity = maxQuantity
	this.UnitPrice = unitPrice
	return &this
}

// NewProjectsUsageTypeUnitPriceWithDefaults instantiates a new ProjectsUsageTypeUnitPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsUsageTypeUnitPriceWithDefaults() *ProjectsUsageTypeUnitPrice {
	this := ProjectsUsageTypeUnitPrice{}
	return &this
}

// GetMinQuantity returns the MinQuantity field value
func (o *ProjectsUsageTypeUnitPrice) GetMinQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinQuantity
}

// GetMinQuantityOk returns a tuple with the MinQuantity field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeUnitPrice) GetMinQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinQuantity, true
}

// SetMinQuantity sets field value
func (o *ProjectsUsageTypeUnitPrice) SetMinQuantity(v int32) {
	o.MinQuantity = v
}

// GetMaxQuantity returns the MaxQuantity field value
func (o *ProjectsUsageTypeUnitPrice) GetMaxQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxQuantity
}

// GetMaxQuantityOk returns a tuple with the MaxQuantity field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeUnitPrice) GetMaxQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxQuantity, true
}

// SetMaxQuantity sets field value
func (o *ProjectsUsageTypeUnitPrice) SetMaxQuantity(v int32) {
	o.MaxQuantity = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *ProjectsUsageTypeUnitPrice) GetUnitPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeUnitPrice) GetUnitPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *ProjectsUsageTypeUnitPrice) SetUnitPrice(v int32) {
	o.UnitPrice = v
}

func (o ProjectsUsageTypeUnitPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["min_quantity"] = o.MinQuantity
	}
	if true {
		toSerialize["max_quantity"] = o.MaxQuantity
	}
	if true {
		toSerialize["unit_price"] = o.UnitPrice
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsUsageTypeUnitPrice struct {
	value *ProjectsUsageTypeUnitPrice
	isSet bool
}

func (v NullableProjectsUsageTypeUnitPrice) Get() *ProjectsUsageTypeUnitPrice {
	return v.value
}

func (v *NullableProjectsUsageTypeUnitPrice) Set(val *ProjectsUsageTypeUnitPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsUsageTypeUnitPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsUsageTypeUnitPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsUsageTypeUnitPrice(val *ProjectsUsageTypeUnitPrice) *NullableProjectsUsageTypeUnitPrice {
	return &NullableProjectsUsageTypeUnitPrice{value: val, isSet: true}
}

func (v NullableProjectsUsageTypeUnitPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsUsageTypeUnitPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


