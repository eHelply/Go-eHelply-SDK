/*
eHelply SDK - 1.1.90

eHelply SDK for SuperStack Services

API version: 1.1.90
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// BasicMetaCreate Basic meta
type BasicMetaCreate struct {
	Name *string `json:"name,omitempty"`
	Slug *bool `json:"slug,omitempty"`
}

// NewBasicMetaCreate instantiates a new BasicMetaCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicMetaCreate() *BasicMetaCreate {
	this := BasicMetaCreate{}
	var slug bool = true
	this.Slug = &slug
	return &this
}

// NewBasicMetaCreateWithDefaults instantiates a new BasicMetaCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicMetaCreateWithDefaults() *BasicMetaCreate {
	this := BasicMetaCreate{}
	var slug bool = true
	this.Slug = &slug
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BasicMetaCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicMetaCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BasicMetaCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BasicMetaCreate) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *BasicMetaCreate) GetSlug() bool {
	if o == nil || o.Slug == nil {
		var ret bool
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicMetaCreate) GetSlugOk() (*bool, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *BasicMetaCreate) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given bool and assigns it to the Slug field.
func (o *BasicMetaCreate) SetSlug(v bool) {
	o.Slug = &v
}

func (o BasicMetaCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	return json.Marshal(toSerialize)
}

type NullableBasicMetaCreate struct {
	value *BasicMetaCreate
	isSet bool
}

func (v NullableBasicMetaCreate) Get() *BasicMetaCreate {
	return v.value
}

func (v *NullableBasicMetaCreate) Set(val *BasicMetaCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicMetaCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicMetaCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicMetaCreate(val *BasicMetaCreate) *NullableBasicMetaCreate {
	return &NullableBasicMetaCreate{value: val, isSet: true}
}

func (v NullableBasicMetaCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicMetaCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


