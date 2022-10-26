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

// Dates struct for Dates
type Dates struct {
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewDates instantiates a new Dates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDates() *Dates {
	this := Dates{}
	var createdAt string = "2022-10-26T19:13:31.399861"
	this.CreatedAt = &createdAt
	var updatedAt string = "2022-10-26T19:13:31.399861"
	this.UpdatedAt = &updatedAt
	return &this
}

// NewDatesWithDefaults instantiates a new Dates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatesWithDefaults() *Dates {
	this := Dates{}
	var createdAt string = "2022-10-26T19:13:31.399861"
	this.CreatedAt = &createdAt
	var updatedAt string = "2022-10-26T19:13:31.399861"
	this.UpdatedAt = &updatedAt
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Dates) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dates) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Dates) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Dates) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Dates) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dates) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Dates) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Dates) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Dates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDates struct {
	value *Dates
	isSet bool
}

func (v NullableDates) Get() *Dates {
	return v.value
}

func (v *NullableDates) Set(val *Dates) {
	v.value = val
	v.isSet = true
}

func (v NullableDates) IsSet() bool {
	return v.isSet
}

func (v *NullableDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDates(val *Dates) *NullableDates {
	return &NullableDates{value: val, isSet: true}
}

func (v NullableDates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


