/*
eHelply SDK - 1.1.100

eHelply SDK for SuperStack Services

API version: 1.1.100
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsUsageTypeDB Used for DB row entry
type ProjectsUsageTypeDB struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Summary string `json:"summary"`
	Category string `json:"category"`
	Service string `json:"service"`
	UnitPrices []ProjectsUsageTypeUnitPrice `json:"unit_prices"`
}

// NewProjectsUsageTypeDB instantiates a new ProjectsUsageTypeDB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsUsageTypeDB(key string, name string, summary string, category string, service string, unitPrices []ProjectsUsageTypeUnitPrice) *ProjectsUsageTypeDB {
	this := ProjectsUsageTypeDB{}
	this.Key = key
	this.Name = name
	this.Summary = summary
	this.Category = category
	this.Service = service
	this.UnitPrices = unitPrices
	return &this
}

// NewProjectsUsageTypeDBWithDefaults instantiates a new ProjectsUsageTypeDB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsUsageTypeDBWithDefaults() *ProjectsUsageTypeDB {
	this := ProjectsUsageTypeDB{}
	return &this
}

// GetKey returns the Key field value
func (o *ProjectsUsageTypeDB) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeDB) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ProjectsUsageTypeDB) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ProjectsUsageTypeDB) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeDB) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectsUsageTypeDB) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value
func (o *ProjectsUsageTypeDB) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeDB) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ProjectsUsageTypeDB) SetSummary(v string) {
	o.Summary = v
}

// GetCategory returns the Category field value
func (o *ProjectsUsageTypeDB) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeDB) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ProjectsUsageTypeDB) SetCategory(v string) {
	o.Category = v
}

// GetService returns the Service field value
func (o *ProjectsUsageTypeDB) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeDB) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ProjectsUsageTypeDB) SetService(v string) {
	o.Service = v
}

// GetUnitPrices returns the UnitPrices field value
func (o *ProjectsUsageTypeDB) GetUnitPrices() []ProjectsUsageTypeUnitPrice {
	if o == nil {
		var ret []ProjectsUsageTypeUnitPrice
		return ret
	}

	return o.UnitPrices
}

// GetUnitPricesOk returns a tuple with the UnitPrices field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeDB) GetUnitPricesOk() ([]ProjectsUsageTypeUnitPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitPrices, true
}

// SetUnitPrices sets field value
func (o *ProjectsUsageTypeDB) SetUnitPrices(v []ProjectsUsageTypeUnitPrice) {
	o.UnitPrices = v
}

func (o ProjectsUsageTypeDB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["unit_prices"] = o.UnitPrices
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsUsageTypeDB struct {
	value *ProjectsUsageTypeDB
	isSet bool
}

func (v NullableProjectsUsageTypeDB) Get() *ProjectsUsageTypeDB {
	return v.value
}

func (v *NullableProjectsUsageTypeDB) Set(val *ProjectsUsageTypeDB) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsUsageTypeDB) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsUsageTypeDB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsUsageTypeDB(val *ProjectsUsageTypeDB) *NullableProjectsUsageTypeDB {
	return &NullableProjectsUsageTypeDB{value: val, isSet: true}
}

func (v NullableProjectsUsageTypeDB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsUsageTypeDB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


