/*
eHelply SDK - 1.1.101

eHelply SDK for SuperStack Services

API version: 1.1.101
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Field Field
type Field struct {
	Uuid *string `json:"uuid,omitempty"`
	Type *int32 `json:"type,omitempty"`
	Placeholder *string `json:"placeholder,omitempty"`
	Validations *Validations `json:"validations,omitempty"`
	Hint *string `json:"hint,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Label *string `json:"label,omitempty"`
	Options *Options `json:"options,omitempty"`
}

// NewField instantiates a new Field object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewField() *Field {
	this := Field{}
	return &this
}

// NewFieldWithDefaults instantiates a new Field object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldWithDefaults() *Field {
	this := Field{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Field) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Field) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Field) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Field) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Field) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Field) SetType(v int32) {
	o.Type = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *Field) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *Field) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *Field) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *Field) GetValidations() Validations {
	if o == nil || o.Validations == nil {
		var ret Validations
		return ret
	}
	return *o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetValidationsOk() (*Validations, bool) {
	if o == nil || o.Validations == nil {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *Field) HasValidations() bool {
	if o != nil && o.Validations != nil {
		return true
	}

	return false
}

// SetValidations gets a reference to the given Validations and assigns it to the Validations field.
func (o *Field) SetValidations(v Validations) {
	o.Validations = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *Field) GetHint() string {
	if o == nil || o.Hint == nil {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetHintOk() (*string, bool) {
	if o == nil || o.Hint == nil {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *Field) HasHint() bool {
	if o != nil && o.Hint != nil {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *Field) SetHint(v string) {
	o.Hint = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Field) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Field) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Field) SetIcon(v string) {
	o.Icon = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Field) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Field) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Field) SetLabel(v string) {
	o.Label = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Field) GetOptions() Options {
	if o == nil || o.Options == nil {
		var ret Options
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Field) GetOptionsOk() (*Options, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Field) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Options and assigns it to the Options field.
func (o *Field) SetOptions(v Options) {
	o.Options = &v
}

func (o Field) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.Validations != nil {
		toSerialize["validations"] = o.Validations
	}
	if o.Hint != nil {
		toSerialize["hint"] = o.Hint
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableField struct {
	value *Field
	isSet bool
}

func (v NullableField) Get() *Field {
	return v.value
}

func (v *NullableField) Set(val *Field) {
	v.value = val
	v.isSet = true
}

func (v NullableField) IsSet() bool {
	return v.isSet
}

func (v *NullableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableField(val *Field) *NullableField {
	return &NullableField{value: val, isSet: true}
}

func (v NullableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


