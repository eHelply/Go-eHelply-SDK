/*
eHelply SDK - 1.1.98

eHelply SDK for SuperStack Services

API version: 1.1.98
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsProjectUsageDB struct for ProjectsProjectUsageDB
type ProjectsProjectUsageDB struct {
	Uuid string `json:"uuid"`
	ProjectUuid string `json:"project_uuid"`
	UsageKey string `json:"usage_key"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	Quantity int32 `json:"quantity"`
	EstimatedCost int32 `json:"estimated_cost"`
	UpdatedAt string `json:"updated_at"`
}

// NewProjectsProjectUsageDB instantiates a new ProjectsProjectUsageDB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsProjectUsageDB(uuid string, projectUuid string, usageKey string, year int32, month int32, quantity int32, estimatedCost int32, updatedAt string) *ProjectsProjectUsageDB {
	this := ProjectsProjectUsageDB{}
	this.Uuid = uuid
	this.ProjectUuid = projectUuid
	this.UsageKey = usageKey
	this.Year = year
	this.Month = month
	this.Quantity = quantity
	this.EstimatedCost = estimatedCost
	this.UpdatedAt = updatedAt
	return &this
}

// NewProjectsProjectUsageDBWithDefaults instantiates a new ProjectsProjectUsageDB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsProjectUsageDBWithDefaults() *ProjectsProjectUsageDB {
	this := ProjectsProjectUsageDB{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ProjectsProjectUsageDB) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ProjectsProjectUsageDB) SetUuid(v string) {
	o.Uuid = v
}

// GetProjectUuid returns the ProjectUuid field value
func (o *ProjectsProjectUsageDB) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *ProjectsProjectUsageDB) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetUsageKey returns the UsageKey field value
func (o *ProjectsProjectUsageDB) GetUsageKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsageKey
}

// GetUsageKeyOk returns a tuple with the UsageKey field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetUsageKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageKey, true
}

// SetUsageKey sets field value
func (o *ProjectsProjectUsageDB) SetUsageKey(v string) {
	o.UsageKey = v
}

// GetYear returns the Year field value
func (o *ProjectsProjectUsageDB) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *ProjectsProjectUsageDB) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *ProjectsProjectUsageDB) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *ProjectsProjectUsageDB) SetMonth(v int32) {
	o.Month = v
}

// GetQuantity returns the Quantity field value
func (o *ProjectsProjectUsageDB) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ProjectsProjectUsageDB) SetQuantity(v int32) {
	o.Quantity = v
}

// GetEstimatedCost returns the EstimatedCost field value
func (o *ProjectsProjectUsageDB) GetEstimatedCost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedCost
}

// GetEstimatedCostOk returns a tuple with the EstimatedCost field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetEstimatedCostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedCost, true
}

// SetEstimatedCost sets field value
func (o *ProjectsProjectUsageDB) SetEstimatedCost(v int32) {
	o.EstimatedCost = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProjectsProjectUsageDB) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectUsageDB) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProjectsProjectUsageDB) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ProjectsProjectUsageDB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if true {
		toSerialize["usage_key"] = o.UsageKey
	}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["estimated_cost"] = o.EstimatedCost
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsProjectUsageDB struct {
	value *ProjectsProjectUsageDB
	isSet bool
}

func (v NullableProjectsProjectUsageDB) Get() *ProjectsProjectUsageDB {
	return v.value
}

func (v *NullableProjectsProjectUsageDB) Set(val *ProjectsProjectUsageDB) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsProjectUsageDB) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsProjectUsageDB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsProjectUsageDB(val *ProjectsProjectUsageDB) *NullableProjectsProjectUsageDB {
	return &NullableProjectsProjectUsageDB{value: val, isSet: true}
}

func (v NullableProjectsProjectUsageDB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsProjectUsageDB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


