/*
eHelply SDK - 1.1.97

eHelply SDK for SuperStack Services

API version: 1.1.97
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Selection Selection Model
type Selection struct {
	Name *string `json:"name,omitempty"`
	Value *float32 `json:"value,omitempty"`
	Icon *string `json:"icon,omitempty"`
}

// NewSelection instantiates a new Selection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelection() *Selection {
	this := Selection{}
	return &this
}

// NewSelectionWithDefaults instantiates a new Selection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectionWithDefaults() *Selection {
	this := Selection{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Selection) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selection) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Selection) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Selection) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Selection) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selection) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Selection) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *Selection) SetValue(v float32) {
	o.Value = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Selection) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selection) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Selection) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Selection) SetIcon(v string) {
	o.Icon = &v
}

func (o Selection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullableSelection struct {
	value *Selection
	isSet bool
}

func (v NullableSelection) Get() *Selection {
	return v.value
}

func (v *NullableSelection) Set(val *Selection) {
	v.value = val
	v.isSet = true
}

func (v NullableSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelection(val *Selection) *NullableSelection {
	return &NullableSelection{value: val, isSet: true}
}

func (v NullableSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


