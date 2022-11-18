/*
eHelply SDK - 1.1.122

eHelply SDK for SuperStack Services

API version: 1.1.122
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// StaffBase struct for StaffBase
type StaffBase struct {
	EntityUuid string `json:"entity_uuid"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
	ScheduleUuid *string `json:"schedule_uuid,omitempty"`
	CatalogUuid *string `json:"catalog_uuid,omitempty"`
	ReviewGroupUuid *string `json:"review_group_uuid,omitempty"`
}

// NewStaffBase instantiates a new StaffBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaffBase(entityUuid string) *StaffBase {
	this := StaffBase{}
	this.EntityUuid = entityUuid
	return &this
}

// NewStaffBaseWithDefaults instantiates a new StaffBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaffBaseWithDefaults() *StaffBase {
	this := StaffBase{}
	return &this
}

// GetEntityUuid returns the EntityUuid field value
func (o *StaffBase) GetEntityUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityUuid
}

// GetEntityUuidOk returns a tuple with the EntityUuid field value
// and a boolean to check if the value has been set.
func (o *StaffBase) GetEntityUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityUuid, true
}

// SetEntityUuid sets field value
func (o *StaffBase) SetEntityUuid(v string) {
	o.EntityUuid = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *StaffBase) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffBase) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *StaffBase) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *StaffBase) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetScheduleUuid returns the ScheduleUuid field value if set, zero value otherwise.
func (o *StaffBase) GetScheduleUuid() string {
	if o == nil || o.ScheduleUuid == nil {
		var ret string
		return ret
	}
	return *o.ScheduleUuid
}

// GetScheduleUuidOk returns a tuple with the ScheduleUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffBase) GetScheduleUuidOk() (*string, bool) {
	if o == nil || o.ScheduleUuid == nil {
		return nil, false
	}
	return o.ScheduleUuid, true
}

// HasScheduleUuid returns a boolean if a field has been set.
func (o *StaffBase) HasScheduleUuid() bool {
	if o != nil && o.ScheduleUuid != nil {
		return true
	}

	return false
}

// SetScheduleUuid gets a reference to the given string and assigns it to the ScheduleUuid field.
func (o *StaffBase) SetScheduleUuid(v string) {
	o.ScheduleUuid = &v
}

// GetCatalogUuid returns the CatalogUuid field value if set, zero value otherwise.
func (o *StaffBase) GetCatalogUuid() string {
	if o == nil || o.CatalogUuid == nil {
		var ret string
		return ret
	}
	return *o.CatalogUuid
}

// GetCatalogUuidOk returns a tuple with the CatalogUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffBase) GetCatalogUuidOk() (*string, bool) {
	if o == nil || o.CatalogUuid == nil {
		return nil, false
	}
	return o.CatalogUuid, true
}

// HasCatalogUuid returns a boolean if a field has been set.
func (o *StaffBase) HasCatalogUuid() bool {
	if o != nil && o.CatalogUuid != nil {
		return true
	}

	return false
}

// SetCatalogUuid gets a reference to the given string and assigns it to the CatalogUuid field.
func (o *StaffBase) SetCatalogUuid(v string) {
	o.CatalogUuid = &v
}

// GetReviewGroupUuid returns the ReviewGroupUuid field value if set, zero value otherwise.
func (o *StaffBase) GetReviewGroupUuid() string {
	if o == nil || o.ReviewGroupUuid == nil {
		var ret string
		return ret
	}
	return *o.ReviewGroupUuid
}

// GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffBase) GetReviewGroupUuidOk() (*string, bool) {
	if o == nil || o.ReviewGroupUuid == nil {
		return nil, false
	}
	return o.ReviewGroupUuid, true
}

// HasReviewGroupUuid returns a boolean if a field has been set.
func (o *StaffBase) HasReviewGroupUuid() bool {
	if o != nil && o.ReviewGroupUuid != nil {
		return true
	}

	return false
}

// SetReviewGroupUuid gets a reference to the given string and assigns it to the ReviewGroupUuid field.
func (o *StaffBase) SetReviewGroupUuid(v string) {
	o.ReviewGroupUuid = &v
}

func (o StaffBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_uuid"] = o.EntityUuid
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.ScheduleUuid != nil {
		toSerialize["schedule_uuid"] = o.ScheduleUuid
	}
	if o.CatalogUuid != nil {
		toSerialize["catalog_uuid"] = o.CatalogUuid
	}
	if o.ReviewGroupUuid != nil {
		toSerialize["review_group_uuid"] = o.ReviewGroupUuid
	}
	return json.Marshal(toSerialize)
}

type NullableStaffBase struct {
	value *StaffBase
	isSet bool
}

func (v NullableStaffBase) Get() *StaffBase {
	return v.value
}

func (v *NullableStaffBase) Set(val *StaffBase) {
	v.value = val
	v.isSet = true
}

func (v NullableStaffBase) IsSet() bool {
	return v.isSet
}

func (v *NullableStaffBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaffBase(val *StaffBase) *NullableStaffBase {
	return &NullableStaffBase{value: val, isSet: true}
}

func (v NullableStaffBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaffBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


