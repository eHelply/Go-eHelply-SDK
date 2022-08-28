/*
eHelply SDK - 1.1.111

eHelply SDK for SuperStack Services

API version: 1.1.111
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// DetailedMetaGet Detailed meta based on Notes
type DetailedMetaGet struct {
	Summary *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	SummaryHistory []string `json:"summary_history,omitempty"`
	DescriptionHistory []string `json:"description_history,omitempty"`
}

// NewDetailedMetaGet instantiates a new DetailedMetaGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedMetaGet() *DetailedMetaGet {
	this := DetailedMetaGet{}
	return &this
}

// NewDetailedMetaGetWithDefaults instantiates a new DetailedMetaGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedMetaGetWithDefaults() *DetailedMetaGet {
	this := DetailedMetaGet{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *DetailedMetaGet) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedMetaGet) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *DetailedMetaGet) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *DetailedMetaGet) SetSummary(v string) {
	o.Summary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DetailedMetaGet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedMetaGet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DetailedMetaGet) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DetailedMetaGet) SetDescription(v string) {
	o.Description = &v
}

// GetSummaryHistory returns the SummaryHistory field value if set, zero value otherwise.
func (o *DetailedMetaGet) GetSummaryHistory() []string {
	if o == nil || o.SummaryHistory == nil {
		var ret []string
		return ret
	}
	return o.SummaryHistory
}

// GetSummaryHistoryOk returns a tuple with the SummaryHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedMetaGet) GetSummaryHistoryOk() ([]string, bool) {
	if o == nil || o.SummaryHistory == nil {
		return nil, false
	}
	return o.SummaryHistory, true
}

// HasSummaryHistory returns a boolean if a field has been set.
func (o *DetailedMetaGet) HasSummaryHistory() bool {
	if o != nil && o.SummaryHistory != nil {
		return true
	}

	return false
}

// SetSummaryHistory gets a reference to the given []string and assigns it to the SummaryHistory field.
func (o *DetailedMetaGet) SetSummaryHistory(v []string) {
	o.SummaryHistory = v
}

// GetDescriptionHistory returns the DescriptionHistory field value if set, zero value otherwise.
func (o *DetailedMetaGet) GetDescriptionHistory() []string {
	if o == nil || o.DescriptionHistory == nil {
		var ret []string
		return ret
	}
	return o.DescriptionHistory
}

// GetDescriptionHistoryOk returns a tuple with the DescriptionHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedMetaGet) GetDescriptionHistoryOk() ([]string, bool) {
	if o == nil || o.DescriptionHistory == nil {
		return nil, false
	}
	return o.DescriptionHistory, true
}

// HasDescriptionHistory returns a boolean if a field has been set.
func (o *DetailedMetaGet) HasDescriptionHistory() bool {
	if o != nil && o.DescriptionHistory != nil {
		return true
	}

	return false
}

// SetDescriptionHistory gets a reference to the given []string and assigns it to the DescriptionHistory field.
func (o *DetailedMetaGet) SetDescriptionHistory(v []string) {
	o.DescriptionHistory = v
}

func (o DetailedMetaGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.SummaryHistory != nil {
		toSerialize["summary_history"] = o.SummaryHistory
	}
	if o.DescriptionHistory != nil {
		toSerialize["description_history"] = o.DescriptionHistory
	}
	return json.Marshal(toSerialize)
}

type NullableDetailedMetaGet struct {
	value *DetailedMetaGet
	isSet bool
}

func (v NullableDetailedMetaGet) Get() *DetailedMetaGet {
	return v.value
}

func (v *NullableDetailedMetaGet) Set(val *DetailedMetaGet) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedMetaGet) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedMetaGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedMetaGet(val *DetailedMetaGet) *NullableDetailedMetaGet {
	return &NullableDetailedMetaGet{value: val, isSet: true}
}

func (v NullableDetailedMetaGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedMetaGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


