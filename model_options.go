/*
eHelply SDK - 1.1.88

eHelply SDK for SuperStack Services

API version: 1.1.88
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// Options struct for Options
type Options struct {
	Required *bool `json:"required,omitempty"`
	Label *string `json:"label,omitempty"`
	InsetLabel *string `json:"insetLabel,omitempty"`
	Placeholder *string `json:"placeholder,omitempty"`
	Hint *string `json:"hint,omitempty"`
	Icon *string `json:"icon,omitempty"`
	MaxLength *float32 `json:"maxLength,omitempty"`
	Counter *bool `json:"counter,omitempty"`
	Caption *string `json:"caption,omitempty"`
	Color *string `json:"color,omitempty"`
	Size *string `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
	IconPosition *string `json:"iconPosition,omitempty"`
	Selections []OptionGroup `json:"selections,omitempty"`
}

// NewOptions instantiates a new Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptions() *Options {
	this := Options{}
	return &this
}

// NewOptionsWithDefaults instantiates a new Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsWithDefaults() *Options {
	this := Options{}
	return &this
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Options) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *Options) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *Options) SetRequired(v bool) {
	o.Required = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Options) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Options) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Options) SetLabel(v string) {
	o.Label = &v
}

// GetInsetLabel returns the InsetLabel field value if set, zero value otherwise.
func (o *Options) GetInsetLabel() string {
	if o == nil || o.InsetLabel == nil {
		var ret string
		return ret
	}
	return *o.InsetLabel
}

// GetInsetLabelOk returns a tuple with the InsetLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetInsetLabelOk() (*string, bool) {
	if o == nil || o.InsetLabel == nil {
		return nil, false
	}
	return o.InsetLabel, true
}

// HasInsetLabel returns a boolean if a field has been set.
func (o *Options) HasInsetLabel() bool {
	if o != nil && o.InsetLabel != nil {
		return true
	}

	return false
}

// SetInsetLabel gets a reference to the given string and assigns it to the InsetLabel field.
func (o *Options) SetInsetLabel(v string) {
	o.InsetLabel = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *Options) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *Options) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *Options) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *Options) GetHint() string {
	if o == nil || o.Hint == nil {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetHintOk() (*string, bool) {
	if o == nil || o.Hint == nil {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *Options) HasHint() bool {
	if o != nil && o.Hint != nil {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *Options) SetHint(v string) {
	o.Hint = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Options) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Options) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Options) SetIcon(v string) {
	o.Icon = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *Options) GetMaxLength() float32 {
	if o == nil || o.MaxLength == nil {
		var ret float32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetMaxLengthOk() (*float32, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *Options) HasMaxLength() bool {
	if o != nil && o.MaxLength != nil {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given float32 and assigns it to the MaxLength field.
func (o *Options) SetMaxLength(v float32) {
	o.MaxLength = &v
}

// GetCounter returns the Counter field value if set, zero value otherwise.
func (o *Options) GetCounter() bool {
	if o == nil || o.Counter == nil {
		var ret bool
		return ret
	}
	return *o.Counter
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetCounterOk() (*bool, bool) {
	if o == nil || o.Counter == nil {
		return nil, false
	}
	return o.Counter, true
}

// HasCounter returns a boolean if a field has been set.
func (o *Options) HasCounter() bool {
	if o != nil && o.Counter != nil {
		return true
	}

	return false
}

// SetCounter gets a reference to the given bool and assigns it to the Counter field.
func (o *Options) SetCounter(v bool) {
	o.Counter = &v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *Options) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *Options) HasCaption() bool {
	if o != nil && o.Caption != nil {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *Options) SetCaption(v string) {
	o.Caption = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Options) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Options) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Options) SetColor(v string) {
	o.Color = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Options) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Options) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Options) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Options) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Options) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Options) SetType(v string) {
	o.Type = &v
}

// GetIconPosition returns the IconPosition field value if set, zero value otherwise.
func (o *Options) GetIconPosition() string {
	if o == nil || o.IconPosition == nil {
		var ret string
		return ret
	}
	return *o.IconPosition
}

// GetIconPositionOk returns a tuple with the IconPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetIconPositionOk() (*string, bool) {
	if o == nil || o.IconPosition == nil {
		return nil, false
	}
	return o.IconPosition, true
}

// HasIconPosition returns a boolean if a field has been set.
func (o *Options) HasIconPosition() bool {
	if o != nil && o.IconPosition != nil {
		return true
	}

	return false
}

// SetIconPosition gets a reference to the given string and assigns it to the IconPosition field.
func (o *Options) SetIconPosition(v string) {
	o.IconPosition = &v
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *Options) GetSelections() []OptionGroup {
	if o == nil || o.Selections == nil {
		var ret []OptionGroup
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Options) GetSelectionsOk() ([]OptionGroup, bool) {
	if o == nil || o.Selections == nil {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *Options) HasSelections() bool {
	if o != nil && o.Selections != nil {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []OptionGroup and assigns it to the Selections field.
func (o *Options) SetSelections(v []OptionGroup) {
	o.Selections = v
}

func (o Options) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.InsetLabel != nil {
		toSerialize["insetLabel"] = o.InsetLabel
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.Hint != nil {
		toSerialize["hint"] = o.Hint
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.Counter != nil {
		toSerialize["counter"] = o.Counter
	}
	if o.Caption != nil {
		toSerialize["caption"] = o.Caption
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IconPosition != nil {
		toSerialize["iconPosition"] = o.IconPosition
	}
	if o.Selections != nil {
		toSerialize["selections"] = o.Selections
	}
	return json.Marshal(toSerialize)
}

type NullableOptions struct {
	value *Options
	isSet bool
}

func (v NullableOptions) Get() *Options {
	return v.value
}

func (v *NullableOptions) Set(val *Options) {
	v.value = val
	v.isSet = true
}

func (v NullableOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptions(val *Options) *NullableOptions {
	return &NullableOptions{value: val, isSet: true}
}

func (v NullableOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

