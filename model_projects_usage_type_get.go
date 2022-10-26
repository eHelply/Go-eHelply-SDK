/*
eHelply SDK - 1.1.112

eHelply SDK for SuperStack Services

API version: 1.1.112
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsUsageTypeGet Used for get endpoint
type ProjectsUsageTypeGet struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Summary string `json:"summary"`
	Category string `json:"category"`
	Service string `json:"service"`
	UnitPrices []ProjectsUsageTypeUnitPrice `json:"unit_prices"`
}

// NewProjectsUsageTypeGet instantiates a new ProjectsUsageTypeGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsUsageTypeGet(key string, name string, summary string, category string, service string, unitPrices []ProjectsUsageTypeUnitPrice) *ProjectsUsageTypeGet {
	this := ProjectsUsageTypeGet{}
	this.Key = key
	this.Name = name
	this.Summary = summary
	this.Category = category
	this.Service = service
	this.UnitPrices = unitPrices
	return &this
}

// NewProjectsUsageTypeGetWithDefaults instantiates a new ProjectsUsageTypeGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsUsageTypeGetWithDefaults() *ProjectsUsageTypeGet {
	this := ProjectsUsageTypeGet{}
	return &this
}

// GetKey returns the Key field value
func (o *ProjectsUsageTypeGet) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeGet) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ProjectsUsageTypeGet) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ProjectsUsageTypeGet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectsUsageTypeGet) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value
func (o *ProjectsUsageTypeGet) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeGet) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ProjectsUsageTypeGet) SetSummary(v string) {
	o.Summary = v
}

// GetCategory returns the Category field value
func (o *ProjectsUsageTypeGet) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeGet) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ProjectsUsageTypeGet) SetCategory(v string) {
	o.Category = v
}

// GetService returns the Service field value
func (o *ProjectsUsageTypeGet) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeGet) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ProjectsUsageTypeGet) SetService(v string) {
	o.Service = v
}

// GetUnitPrices returns the UnitPrices field value
func (o *ProjectsUsageTypeGet) GetUnitPrices() []ProjectsUsageTypeUnitPrice {
	if o == nil {
		var ret []ProjectsUsageTypeUnitPrice
		return ret
	}

	return o.UnitPrices
}

// GetUnitPricesOk returns a tuple with the UnitPrices field value
// and a boolean to check if the value has been set.
func (o *ProjectsUsageTypeGet) GetUnitPricesOk() ([]ProjectsUsageTypeUnitPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitPrices, true
}

// SetUnitPrices sets field value
func (o *ProjectsUsageTypeGet) SetUnitPrices(v []ProjectsUsageTypeUnitPrice) {
	o.UnitPrices = v
}

func (o ProjectsUsageTypeGet) MarshalJSON() ([]byte, error) {
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

type NullableProjectsUsageTypeGet struct {
	value *ProjectsUsageTypeGet
	isSet bool
}

func (v NullableProjectsUsageTypeGet) Get() *ProjectsUsageTypeGet {
	return v.value
}

func (v *NullableProjectsUsageTypeGet) Set(val *ProjectsUsageTypeGet) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsUsageTypeGet) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsUsageTypeGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsUsageTypeGet(val *ProjectsUsageTypeGet) *NullableProjectsUsageTypeGet {
	return &NullableProjectsUsageTypeGet{value: val, isSet: true}
}

func (v NullableProjectsUsageTypeGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsUsageTypeGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


