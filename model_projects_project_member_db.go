/*
eHelply SDK - 1.1.88

eHelply SDK for SuperStack Services

API version: 1.1.88
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectsProjectMemberDB struct for ProjectsProjectMemberDB
type ProjectsProjectMemberDB struct {
	Uuid string `json:"uuid"`
	ProjectUuid string `json:"project_uuid"`
	EntityUuid string `json:"entity_uuid"`
	Role string `json:"role"`
	CreatedAt string `json:"created_at"`
}

// NewProjectsProjectMemberDB instantiates a new ProjectsProjectMemberDB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsProjectMemberDB(uuid string, projectUuid string, entityUuid string, role string, createdAt string) *ProjectsProjectMemberDB {
	this := ProjectsProjectMemberDB{}
	this.Uuid = uuid
	this.ProjectUuid = projectUuid
	this.EntityUuid = entityUuid
	this.Role = role
	this.CreatedAt = createdAt
	return &this
}

// NewProjectsProjectMemberDBWithDefaults instantiates a new ProjectsProjectMemberDB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsProjectMemberDBWithDefaults() *ProjectsProjectMemberDB {
	this := ProjectsProjectMemberDB{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ProjectsProjectMemberDB) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectMemberDB) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ProjectsProjectMemberDB) SetUuid(v string) {
	o.Uuid = v
}

// GetProjectUuid returns the ProjectUuid field value
func (o *ProjectsProjectMemberDB) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectMemberDB) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *ProjectsProjectMemberDB) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetEntityUuid returns the EntityUuid field value
func (o *ProjectsProjectMemberDB) GetEntityUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityUuid
}

// GetEntityUuidOk returns a tuple with the EntityUuid field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectMemberDB) GetEntityUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityUuid, true
}

// SetEntityUuid sets field value
func (o *ProjectsProjectMemberDB) SetEntityUuid(v string) {
	o.EntityUuid = v
}

// GetRole returns the Role field value
func (o *ProjectsProjectMemberDB) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectMemberDB) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProjectsProjectMemberDB) SetRole(v string) {
	o.Role = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectsProjectMemberDB) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectsProjectMemberDB) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectsProjectMemberDB) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o ProjectsProjectMemberDB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if true {
		toSerialize["entity_uuid"] = o.EntityUuid
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsProjectMemberDB struct {
	value *ProjectsProjectMemberDB
	isSet bool
}

func (v NullableProjectsProjectMemberDB) Get() *ProjectsProjectMemberDB {
	return v.value
}

func (v *NullableProjectsProjectMemberDB) Set(val *ProjectsProjectMemberDB) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsProjectMemberDB) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsProjectMemberDB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsProjectMemberDB(val *ProjectsProjectMemberDB) *NullableProjectsProjectMemberDB {
	return &NullableProjectsProjectMemberDB{value: val, isSet: true}
}

func (v NullableProjectsProjectMemberDB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsProjectMemberDB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

