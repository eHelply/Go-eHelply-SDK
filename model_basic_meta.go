/*
eHelply SDK - 1.1.107

eHelply SDK for SuperStack Services

API version: 1.1.107
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// BasicMeta Basic meta
type BasicMeta struct {
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewBasicMeta instantiates a new BasicMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicMeta() *BasicMeta {
	this := BasicMeta{}
	return &this
}

// NewBasicMetaWithDefaults instantiates a new BasicMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicMetaWithDefaults() *BasicMeta {
	this := BasicMeta{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BasicMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BasicMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BasicMeta) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *BasicMeta) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicMeta) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *BasicMeta) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *BasicMeta) SetSlug(v string) {
	o.Slug = &v
}

func (o BasicMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	return json.Marshal(toSerialize)
}

type NullableBasicMeta struct {
	value *BasicMeta
	isSet bool
}

func (v NullableBasicMeta) Get() *BasicMeta {
	return v.value
}

func (v *NullableBasicMeta) Set(val *BasicMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicMeta(val *BasicMeta) *NullableBasicMeta {
	return &NullableBasicMeta{value: val, isSet: true}
}

func (v NullableBasicMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


