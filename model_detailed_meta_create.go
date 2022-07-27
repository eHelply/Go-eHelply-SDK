/*
eHelply SDK - 1.1.103

eHelply SDK for SuperStack Services

API version: 1.1.103
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// DetailedMetaCreate Detailed meta based on Notes
type DetailedMetaCreate struct {
	Summary *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewDetailedMetaCreate instantiates a new DetailedMetaCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedMetaCreate() *DetailedMetaCreate {
	this := DetailedMetaCreate{}
	return &this
}

// NewDetailedMetaCreateWithDefaults instantiates a new DetailedMetaCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedMetaCreateWithDefaults() *DetailedMetaCreate {
	this := DetailedMetaCreate{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *DetailedMetaCreate) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedMetaCreate) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *DetailedMetaCreate) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *DetailedMetaCreate) SetSummary(v string) {
	o.Summary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DetailedMetaCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedMetaCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DetailedMetaCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DetailedMetaCreate) SetDescription(v string) {
	o.Description = &v
}

func (o DetailedMetaCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableDetailedMetaCreate struct {
	value *DetailedMetaCreate
	isSet bool
}

func (v NullableDetailedMetaCreate) Get() *DetailedMetaCreate {
	return v.value
}

func (v *NullableDetailedMetaCreate) Set(val *DetailedMetaCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedMetaCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedMetaCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedMetaCreate(val *DetailedMetaCreate) *NullableDetailedMetaCreate {
	return &NullableDetailedMetaCreate{value: val, isSet: true}
}

func (v NullableDetailedMetaCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedMetaCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


