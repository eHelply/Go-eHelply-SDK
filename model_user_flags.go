/*
eHelply SDK - 1.1.121

eHelply SDK for SuperStack Services

API version: 1.1.121
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UserFlags Flags that may be attached to a user
type UserFlags struct {
	RequiresTour *bool `json:"requires_tour,omitempty"`
	MissingData *bool `json:"missing_data,omitempty"`
	LegalUpdates *bool `json:"legal_updates,omitempty"`
	Newsletters *bool `json:"newsletters,omitempty"`
}

// NewUserFlags instantiates a new UserFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFlags() *UserFlags {
	this := UserFlags{}
	var requiresTour bool = false
	this.RequiresTour = &requiresTour
	var missingData bool = false
	this.MissingData = &missingData
	var legalUpdates bool = false
	this.LegalUpdates = &legalUpdates
	var newsletters bool = false
	this.Newsletters = &newsletters
	return &this
}

// NewUserFlagsWithDefaults instantiates a new UserFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFlagsWithDefaults() *UserFlags {
	this := UserFlags{}
	var requiresTour bool = false
	this.RequiresTour = &requiresTour
	var missingData bool = false
	this.MissingData = &missingData
	var legalUpdates bool = false
	this.LegalUpdates = &legalUpdates
	var newsletters bool = false
	this.Newsletters = &newsletters
	return &this
}

// GetRequiresTour returns the RequiresTour field value if set, zero value otherwise.
func (o *UserFlags) GetRequiresTour() bool {
	if o == nil || o.RequiresTour == nil {
		var ret bool
		return ret
	}
	return *o.RequiresTour
}

// GetRequiresTourOk returns a tuple with the RequiresTour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFlags) GetRequiresTourOk() (*bool, bool) {
	if o == nil || o.RequiresTour == nil {
		return nil, false
	}
	return o.RequiresTour, true
}

// HasRequiresTour returns a boolean if a field has been set.
func (o *UserFlags) HasRequiresTour() bool {
	if o != nil && o.RequiresTour != nil {
		return true
	}

	return false
}

// SetRequiresTour gets a reference to the given bool and assigns it to the RequiresTour field.
func (o *UserFlags) SetRequiresTour(v bool) {
	o.RequiresTour = &v
}

// GetMissingData returns the MissingData field value if set, zero value otherwise.
func (o *UserFlags) GetMissingData() bool {
	if o == nil || o.MissingData == nil {
		var ret bool
		return ret
	}
	return *o.MissingData
}

// GetMissingDataOk returns a tuple with the MissingData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFlags) GetMissingDataOk() (*bool, bool) {
	if o == nil || o.MissingData == nil {
		return nil, false
	}
	return o.MissingData, true
}

// HasMissingData returns a boolean if a field has been set.
func (o *UserFlags) HasMissingData() bool {
	if o != nil && o.MissingData != nil {
		return true
	}

	return false
}

// SetMissingData gets a reference to the given bool and assigns it to the MissingData field.
func (o *UserFlags) SetMissingData(v bool) {
	o.MissingData = &v
}

// GetLegalUpdates returns the LegalUpdates field value if set, zero value otherwise.
func (o *UserFlags) GetLegalUpdates() bool {
	if o == nil || o.LegalUpdates == nil {
		var ret bool
		return ret
	}
	return *o.LegalUpdates
}

// GetLegalUpdatesOk returns a tuple with the LegalUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFlags) GetLegalUpdatesOk() (*bool, bool) {
	if o == nil || o.LegalUpdates == nil {
		return nil, false
	}
	return o.LegalUpdates, true
}

// HasLegalUpdates returns a boolean if a field has been set.
func (o *UserFlags) HasLegalUpdates() bool {
	if o != nil && o.LegalUpdates != nil {
		return true
	}

	return false
}

// SetLegalUpdates gets a reference to the given bool and assigns it to the LegalUpdates field.
func (o *UserFlags) SetLegalUpdates(v bool) {
	o.LegalUpdates = &v
}

// GetNewsletters returns the Newsletters field value if set, zero value otherwise.
func (o *UserFlags) GetNewsletters() bool {
	if o == nil || o.Newsletters == nil {
		var ret bool
		return ret
	}
	return *o.Newsletters
}

// GetNewslettersOk returns a tuple with the Newsletters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFlags) GetNewslettersOk() (*bool, bool) {
	if o == nil || o.Newsletters == nil {
		return nil, false
	}
	return o.Newsletters, true
}

// HasNewsletters returns a boolean if a field has been set.
func (o *UserFlags) HasNewsletters() bool {
	if o != nil && o.Newsletters != nil {
		return true
	}

	return false
}

// SetNewsletters gets a reference to the given bool and assigns it to the Newsletters field.
func (o *UserFlags) SetNewsletters(v bool) {
	o.Newsletters = &v
}

func (o UserFlags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequiresTour != nil {
		toSerialize["requires_tour"] = o.RequiresTour
	}
	if o.MissingData != nil {
		toSerialize["missing_data"] = o.MissingData
	}
	if o.LegalUpdates != nil {
		toSerialize["legal_updates"] = o.LegalUpdates
	}
	if o.Newsletters != nil {
		toSerialize["newsletters"] = o.Newsletters
	}
	return json.Marshal(toSerialize)
}

type NullableUserFlags struct {
	value *UserFlags
	isSet bool
}

func (v NullableUserFlags) Get() *UserFlags {
	return v.value
}

func (v *NullableUserFlags) Set(val *UserFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFlags(val *UserFlags) *NullableUserFlags {
	return &NullableUserFlags{value: val, isSet: true}
}

func (v NullableUserFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


