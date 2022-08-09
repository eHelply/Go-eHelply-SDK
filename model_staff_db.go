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

// StaffDb **:param** uuid                                **type:** string **:param** project_uuid                        **type:** string or None  **:param** entity_uuid                         **type:** string or None  **:param** schedule_uuid                       **type:** string or None  **:param** catalog_uuid                        **type:** string or None  **:param** review_group_uuid                   **type:** string or None  **:param** created_at                          **type:** string or None  **:param** updated_at                          **type:** string or None  **:param** deleted_at                          **type:** string or None
type StaffDb struct {
	Uuid string `json:"uuid"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
	EntityUuid *string `json:"entity_uuid,omitempty"`
	ScheduleUuid *string `json:"schedule_uuid,omitempty"`
	CatalogUuid *string `json:"catalog_uuid,omitempty"`
	ReviewGroupUuid *string `json:"review_group_uuid,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewStaffDb instantiates a new StaffDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaffDb(uuid string) *StaffDb {
	this := StaffDb{}
	this.Uuid = uuid
	return &this
}

// NewStaffDbWithDefaults instantiates a new StaffDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaffDbWithDefaults() *StaffDb {
	this := StaffDb{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *StaffDb) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *StaffDb) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *StaffDb) SetUuid(v string) {
	o.Uuid = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *StaffDb) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *StaffDb) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *StaffDb) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetEntityUuid returns the EntityUuid field value if set, zero value otherwise.
func (o *StaffDb) GetEntityUuid() string {
	if o == nil || o.EntityUuid == nil {
		var ret string
		return ret
	}
	return *o.EntityUuid
}

// GetEntityUuidOk returns a tuple with the EntityUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetEntityUuidOk() (*string, bool) {
	if o == nil || o.EntityUuid == nil {
		return nil, false
	}
	return o.EntityUuid, true
}

// HasEntityUuid returns a boolean if a field has been set.
func (o *StaffDb) HasEntityUuid() bool {
	if o != nil && o.EntityUuid != nil {
		return true
	}

	return false
}

// SetEntityUuid gets a reference to the given string and assigns it to the EntityUuid field.
func (o *StaffDb) SetEntityUuid(v string) {
	o.EntityUuid = &v
}

// GetScheduleUuid returns the ScheduleUuid field value if set, zero value otherwise.
func (o *StaffDb) GetScheduleUuid() string {
	if o == nil || o.ScheduleUuid == nil {
		var ret string
		return ret
	}
	return *o.ScheduleUuid
}

// GetScheduleUuidOk returns a tuple with the ScheduleUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetScheduleUuidOk() (*string, bool) {
	if o == nil || o.ScheduleUuid == nil {
		return nil, false
	}
	return o.ScheduleUuid, true
}

// HasScheduleUuid returns a boolean if a field has been set.
func (o *StaffDb) HasScheduleUuid() bool {
	if o != nil && o.ScheduleUuid != nil {
		return true
	}

	return false
}

// SetScheduleUuid gets a reference to the given string and assigns it to the ScheduleUuid field.
func (o *StaffDb) SetScheduleUuid(v string) {
	o.ScheduleUuid = &v
}

// GetCatalogUuid returns the CatalogUuid field value if set, zero value otherwise.
func (o *StaffDb) GetCatalogUuid() string {
	if o == nil || o.CatalogUuid == nil {
		var ret string
		return ret
	}
	return *o.CatalogUuid
}

// GetCatalogUuidOk returns a tuple with the CatalogUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetCatalogUuidOk() (*string, bool) {
	if o == nil || o.CatalogUuid == nil {
		return nil, false
	}
	return o.CatalogUuid, true
}

// HasCatalogUuid returns a boolean if a field has been set.
func (o *StaffDb) HasCatalogUuid() bool {
	if o != nil && o.CatalogUuid != nil {
		return true
	}

	return false
}

// SetCatalogUuid gets a reference to the given string and assigns it to the CatalogUuid field.
func (o *StaffDb) SetCatalogUuid(v string) {
	o.CatalogUuid = &v
}

// GetReviewGroupUuid returns the ReviewGroupUuid field value if set, zero value otherwise.
func (o *StaffDb) GetReviewGroupUuid() string {
	if o == nil || o.ReviewGroupUuid == nil {
		var ret string
		return ret
	}
	return *o.ReviewGroupUuid
}

// GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetReviewGroupUuidOk() (*string, bool) {
	if o == nil || o.ReviewGroupUuid == nil {
		return nil, false
	}
	return o.ReviewGroupUuid, true
}

// HasReviewGroupUuid returns a boolean if a field has been set.
func (o *StaffDb) HasReviewGroupUuid() bool {
	if o != nil && o.ReviewGroupUuid != nil {
		return true
	}

	return false
}

// SetReviewGroupUuid gets a reference to the given string and assigns it to the ReviewGroupUuid field.
func (o *StaffDb) SetReviewGroupUuid(v string) {
	o.ReviewGroupUuid = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StaffDb) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StaffDb) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *StaffDb) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *StaffDb) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *StaffDb) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *StaffDb) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *StaffDb) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffDb) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *StaffDb) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *StaffDb) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o StaffDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.EntityUuid != nil {
		toSerialize["entity_uuid"] = o.EntityUuid
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
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableStaffDb struct {
	value *StaffDb
	isSet bool
}

func (v NullableStaffDb) Get() *StaffDb {
	return v.value
}

func (v *NullableStaffDb) Set(val *StaffDb) {
	v.value = val
	v.isSet = true
}

func (v NullableStaffDb) IsSet() bool {
	return v.isSet
}

func (v *NullableStaffDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaffDb(val *StaffDb) *NullableStaffDb {
	return &NullableStaffDb{value: val, isSet: true}
}

func (v NullableStaffDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaffDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


