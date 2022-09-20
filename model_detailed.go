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

// Detailed struct for Detailed
type Detailed struct {
	SummaryUuid *string `json:"summary_uuid,omitempty"`
	DescriptionUuid *string `json:"description_uuid,omitempty"`
}

// NewDetailed instantiates a new Detailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailed() *Detailed {
	this := Detailed{}
	return &this
}

// NewDetailedWithDefaults instantiates a new Detailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedWithDefaults() *Detailed {
	this := Detailed{}
	return &this
}

// GetSummaryUuid returns the SummaryUuid field value if set, zero value otherwise.
func (o *Detailed) GetSummaryUuid() string {
	if o == nil || o.SummaryUuid == nil {
		var ret string
		return ret
	}
	return *o.SummaryUuid
}

// GetSummaryUuidOk returns a tuple with the SummaryUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detailed) GetSummaryUuidOk() (*string, bool) {
	if o == nil || o.SummaryUuid == nil {
		return nil, false
	}
	return o.SummaryUuid, true
}

// HasSummaryUuid returns a boolean if a field has been set.
func (o *Detailed) HasSummaryUuid() bool {
	if o != nil && o.SummaryUuid != nil {
		return true
	}

	return false
}

// SetSummaryUuid gets a reference to the given string and assigns it to the SummaryUuid field.
func (o *Detailed) SetSummaryUuid(v string) {
	o.SummaryUuid = &v
}

// GetDescriptionUuid returns the DescriptionUuid field value if set, zero value otherwise.
func (o *Detailed) GetDescriptionUuid() string {
	if o == nil || o.DescriptionUuid == nil {
		var ret string
		return ret
	}
	return *o.DescriptionUuid
}

// GetDescriptionUuidOk returns a tuple with the DescriptionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detailed) GetDescriptionUuidOk() (*string, bool) {
	if o == nil || o.DescriptionUuid == nil {
		return nil, false
	}
	return o.DescriptionUuid, true
}

// HasDescriptionUuid returns a boolean if a field has been set.
func (o *Detailed) HasDescriptionUuid() bool {
	if o != nil && o.DescriptionUuid != nil {
		return true
	}

	return false
}

// SetDescriptionUuid gets a reference to the given string and assigns it to the DescriptionUuid field.
func (o *Detailed) SetDescriptionUuid(v string) {
	o.DescriptionUuid = &v
}

func (o Detailed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SummaryUuid != nil {
		toSerialize["summary_uuid"] = o.SummaryUuid
	}
	if o.DescriptionUuid != nil {
		toSerialize["description_uuid"] = o.DescriptionUuid
	}
	return json.Marshal(toSerialize)
}

type NullableDetailed struct {
	value *Detailed
	isSet bool
}

func (v NullableDetailed) Get() *Detailed {
	return v.value
}

func (v *NullableDetailed) Set(val *Detailed) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailed) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailed(val *Detailed) *NullableDetailed {
	return &NullableDetailed{value: val, isSet: true}
}

func (v NullableDetailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


