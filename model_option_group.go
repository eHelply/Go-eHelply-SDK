/*
eHelply SDK - 1.1.109

eHelply SDK for SuperStack Services

API version: 1.1.109
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// OptionGroup Option Group
type OptionGroup struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Selections []Selection `json:"selections,omitempty"`
}

// NewOptionGroup instantiates a new OptionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionGroup() *OptionGroup {
	this := OptionGroup{}
	return &this
}

// NewOptionGroupWithDefaults instantiates a new OptionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionGroupWithDefaults() *OptionGroup {
	this := OptionGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OptionGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OptionGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OptionGroup) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OptionGroup) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroup) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OptionGroup) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OptionGroup) SetType(v string) {
	o.Type = &v
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *OptionGroup) GetSelections() []Selection {
	if o == nil || o.Selections == nil {
		var ret []Selection
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroup) GetSelectionsOk() ([]Selection, bool) {
	if o == nil || o.Selections == nil {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *OptionGroup) HasSelections() bool {
	if o != nil && o.Selections != nil {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []Selection and assigns it to the Selections field.
func (o *OptionGroup) SetSelections(v []Selection) {
	o.Selections = v
}

func (o OptionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Selections != nil {
		toSerialize["selections"] = o.Selections
	}
	return json.Marshal(toSerialize)
}

type NullableOptionGroup struct {
	value *OptionGroup
	isSet bool
}

func (v NullableOptionGroup) Get() *OptionGroup {
	return v.value
}

func (v *NullableOptionGroup) Set(val *OptionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionGroup(val *OptionGroup) *NullableOptionGroup {
	return &NullableOptionGroup{value: val, isSet: true}
}

func (v NullableOptionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


