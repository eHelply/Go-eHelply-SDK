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

// FieldDynamo Field Dynamo
type FieldDynamo struct {
	Type map[string]interface{} `json:"type,omitempty"`
	Placeholder *string `json:"placeholder,omitempty"`
	Validations map[string]interface{} `json:"validations,omitempty"`
	Hint *string `json:"hint,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Label *string `json:"label,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
	Uuid string `json:"uuid"`
}

// NewFieldDynamo instantiates a new FieldDynamo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldDynamo(uuid string) *FieldDynamo {
	this := FieldDynamo{}
	this.Uuid = uuid
	return &this
}

// NewFieldDynamoWithDefaults instantiates a new FieldDynamo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldDynamoWithDefaults() *FieldDynamo {
	this := FieldDynamo{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FieldDynamo) GetType() map[string]interface{} {
	if o == nil || o.Type == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FieldDynamo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *FieldDynamo) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *FieldDynamo) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *FieldDynamo) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *FieldDynamo) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *FieldDynamo) GetValidations() map[string]interface{} {
	if o == nil || o.Validations == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetValidationsOk() (map[string]interface{}, bool) {
	if o == nil || o.Validations == nil {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *FieldDynamo) HasValidations() bool {
	if o != nil && o.Validations != nil {
		return true
	}

	return false
}

// SetValidations gets a reference to the given map[string]interface{} and assigns it to the Validations field.
func (o *FieldDynamo) SetValidations(v map[string]interface{}) {
	o.Validations = v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *FieldDynamo) GetHint() string {
	if o == nil || o.Hint == nil {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetHintOk() (*string, bool) {
	if o == nil || o.Hint == nil {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *FieldDynamo) HasHint() bool {
	if o != nil && o.Hint != nil {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *FieldDynamo) SetHint(v string) {
	o.Hint = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *FieldDynamo) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *FieldDynamo) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *FieldDynamo) SetIcon(v string) {
	o.Icon = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FieldDynamo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FieldDynamo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FieldDynamo) SetLabel(v string) {
	o.Label = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FieldDynamo) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FieldDynamo) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *FieldDynamo) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetUuid returns the Uuid field value
func (o *FieldDynamo) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *FieldDynamo) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *FieldDynamo) SetUuid(v string) {
	o.Uuid = v
}

func (o FieldDynamo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableFieldDynamo struct {
	value *FieldDynamo
	isSet bool
}

func (v NullableFieldDynamo) Get() *FieldDynamo {
	return v.value
}

func (v *NullableFieldDynamo) Set(val *FieldDynamo) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldDynamo) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldDynamo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldDynamo(val *FieldDynamo) *NullableFieldDynamo {
	return &NullableFieldDynamo{value: val, isSet: true}
}

func (v NullableFieldDynamo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldDynamo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


