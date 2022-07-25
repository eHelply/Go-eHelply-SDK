/*
eHelply SDK - 1.1.92

eHelply SDK for SuperStack Services

API version: 1.1.92
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectDB Used for DB row entry
type ProjectDB struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	CurrentSpend int32 `json:"current_spend"`
	MaxSpend int32 `json:"max_spend"`
	ProjectStatus string `json:"project_status"`
	ArchivedAt *string `json:"archived_at,omitempty"`
}

// NewProjectDB instantiates a new ProjectDB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDB(uuid string, name string, createdAt string, currentSpend int32, maxSpend int32, projectStatus string) *ProjectDB {
	this := ProjectDB{}
	this.Uuid = uuid
	this.Name = name
	this.CreatedAt = createdAt
	this.CurrentSpend = currentSpend
	this.MaxSpend = maxSpend
	this.ProjectStatus = projectStatus
	return &this
}

// NewProjectDBWithDefaults instantiates a new ProjectDB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDBWithDefaults() *ProjectDB {
	this := ProjectDB{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ProjectDB) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ProjectDB) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ProjectDB) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *ProjectDB) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectDB) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectDB) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectDB) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectDB) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectDB) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCurrentSpend returns the CurrentSpend field value
func (o *ProjectDB) GetCurrentSpend() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentSpend
}

// GetCurrentSpendOk returns a tuple with the CurrentSpend field value
// and a boolean to check if the value has been set.
func (o *ProjectDB) GetCurrentSpendOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentSpend, true
}

// SetCurrentSpend sets field value
func (o *ProjectDB) SetCurrentSpend(v int32) {
	o.CurrentSpend = v
}

// GetMaxSpend returns the MaxSpend field value
func (o *ProjectDB) GetMaxSpend() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxSpend
}

// GetMaxSpendOk returns a tuple with the MaxSpend field value
// and a boolean to check if the value has been set.
func (o *ProjectDB) GetMaxSpendOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSpend, true
}

// SetMaxSpend sets field value
func (o *ProjectDB) SetMaxSpend(v int32) {
	o.MaxSpend = v
}

// GetProjectStatus returns the ProjectStatus field value
func (o *ProjectDB) GetProjectStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectStatus
}

// GetProjectStatusOk returns a tuple with the ProjectStatus field value
// and a boolean to check if the value has been set.
func (o *ProjectDB) GetProjectStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectStatus, true
}

// SetProjectStatus sets field value
func (o *ProjectDB) SetProjectStatus(v string) {
	o.ProjectStatus = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *ProjectDB) GetArchivedAt() string {
	if o == nil || o.ArchivedAt == nil {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDB) GetArchivedAtOk() (*string, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *ProjectDB) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *ProjectDB) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

func (o ProjectDB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["current_spend"] = o.CurrentSpend
	}
	if true {
		toSerialize["max_spend"] = o.MaxSpend
	}
	if true {
		toSerialize["project_status"] = o.ProjectStatus
	}
	if o.ArchivedAt != nil {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProjectDB struct {
	value *ProjectDB
	isSet bool
}

func (v NullableProjectDB) Get() *ProjectDB {
	return v.value
}

func (v *NullableProjectDB) Set(val *ProjectDB) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDB) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDB(val *ProjectDB) *NullableProjectDB {
	return &NullableProjectDB{value: val, isSet: true}
}

func (v NullableProjectDB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


