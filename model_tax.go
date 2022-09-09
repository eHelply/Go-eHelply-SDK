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

// Tax struct for Tax
type Tax struct {
	Name string `json:"name"`
	Rate int32 `json:"rate"`
}

// NewTax instantiates a new Tax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTax(name string, rate int32) *Tax {
	this := Tax{}
	this.Name = name
	this.Rate = rate
	return &this
}

// NewTaxWithDefaults instantiates a new Tax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxWithDefaults() *Tax {
	this := Tax{}
	return &this
}

// GetName returns the Name field value
func (o *Tax) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tax) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tax) SetName(v string) {
	o.Name = v
}

// GetRate returns the Rate field value
func (o *Tax) GetRate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *Tax) GetRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *Tax) SetRate(v int32) {
	o.Rate = v
}

func (o Tax) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["rate"] = o.Rate
	}
	return json.Marshal(toSerialize)
}

type NullableTax struct {
	value *Tax
	isSet bool
}

func (v NullableTax) Get() *Tax {
	return v.value
}

func (v *NullableTax) Set(val *Tax) {
	v.value = val
	v.isSet = true
}

func (v NullableTax) IsSet() bool {
	return v.isSet
}

func (v *NullableTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTax(val *Tax) *NullableTax {
	return &NullableTax{value: val, isSet: true}
}

func (v NullableTax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


