/*
eHelply SDK - 1.1.101

eHelply SDK for SuperStack Services

API version: 1.1.101
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// StaffResponse **:param** uuid                                **type:** string **:param** project_uuid                        **type:** string or None  **:param** entity                              **type:** string or None  **:param** place                               **type:** string or None  **:param** company                             **type:** string or None  **:param** schedule                            **type:** string or None  **:param** catalog                             **type:** string or None  **:param** reviews                             **type:** string or None  **:param** created_at                          **type:** string or None  **:param** updated_at                          **type:** string or None  **:param** deleted_at                          **type:** string or None
type StaffResponse struct {
	Uuid string `json:"uuid"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
	Entity map[string]interface{} `json:"entity,omitempty"`
	Place map[string]interface{} `json:"place,omitempty"`
	PlaceRoles []string `json:"place_roles,omitempty"`
	Company map[string]interface{} `json:"company,omitempty"`
	CompanyRoles []string `json:"company_roles,omitempty"`
	Schedule map[string]interface{} `json:"schedule,omitempty"`
	Catalog map[string]interface{} `json:"catalog,omitempty"`
	Reviews map[string]interface{} `json:"reviews,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewStaffResponse instantiates a new StaffResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaffResponse(uuid string) *StaffResponse {
	this := StaffResponse{}
	this.Uuid = uuid
	return &this
}

// NewStaffResponseWithDefaults instantiates a new StaffResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaffResponseWithDefaults() *StaffResponse {
	this := StaffResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *StaffResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *StaffResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *StaffResponse) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *StaffResponse) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *StaffResponse) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *StaffResponse) GetEntity() map[string]interface{} {
	if o == nil || o.Entity == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetEntityOk() (map[string]interface{}, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *StaffResponse) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *StaffResponse) SetEntity(v map[string]interface{}) {
	o.Entity = v
}

// GetPlace returns the Place field value if set, zero value otherwise.
func (o *StaffResponse) GetPlace() map[string]interface{} {
	if o == nil || o.Place == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Place
}

// GetPlaceOk returns a tuple with the Place field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetPlaceOk() (map[string]interface{}, bool) {
	if o == nil || o.Place == nil {
		return nil, false
	}
	return o.Place, true
}

// HasPlace returns a boolean if a field has been set.
func (o *StaffResponse) HasPlace() bool {
	if o != nil && o.Place != nil {
		return true
	}

	return false
}

// SetPlace gets a reference to the given map[string]interface{} and assigns it to the Place field.
func (o *StaffResponse) SetPlace(v map[string]interface{}) {
	o.Place = v
}

// GetPlaceRoles returns the PlaceRoles field value if set, zero value otherwise.
func (o *StaffResponse) GetPlaceRoles() []string {
	if o == nil || o.PlaceRoles == nil {
		var ret []string
		return ret
	}
	return o.PlaceRoles
}

// GetPlaceRolesOk returns a tuple with the PlaceRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetPlaceRolesOk() ([]string, bool) {
	if o == nil || o.PlaceRoles == nil {
		return nil, false
	}
	return o.PlaceRoles, true
}

// HasPlaceRoles returns a boolean if a field has been set.
func (o *StaffResponse) HasPlaceRoles() bool {
	if o != nil && o.PlaceRoles != nil {
		return true
	}

	return false
}

// SetPlaceRoles gets a reference to the given []string and assigns it to the PlaceRoles field.
func (o *StaffResponse) SetPlaceRoles(v []string) {
	o.PlaceRoles = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *StaffResponse) GetCompany() map[string]interface{} {
	if o == nil || o.Company == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetCompanyOk() (map[string]interface{}, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *StaffResponse) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given map[string]interface{} and assigns it to the Company field.
func (o *StaffResponse) SetCompany(v map[string]interface{}) {
	o.Company = v
}

// GetCompanyRoles returns the CompanyRoles field value if set, zero value otherwise.
func (o *StaffResponse) GetCompanyRoles() []string {
	if o == nil || o.CompanyRoles == nil {
		var ret []string
		return ret
	}
	return o.CompanyRoles
}

// GetCompanyRolesOk returns a tuple with the CompanyRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetCompanyRolesOk() ([]string, bool) {
	if o == nil || o.CompanyRoles == nil {
		return nil, false
	}
	return o.CompanyRoles, true
}

// HasCompanyRoles returns a boolean if a field has been set.
func (o *StaffResponse) HasCompanyRoles() bool {
	if o != nil && o.CompanyRoles != nil {
		return true
	}

	return false
}

// SetCompanyRoles gets a reference to the given []string and assigns it to the CompanyRoles field.
func (o *StaffResponse) SetCompanyRoles(v []string) {
	o.CompanyRoles = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *StaffResponse) GetSchedule() map[string]interface{} {
	if o == nil || o.Schedule == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetScheduleOk() (map[string]interface{}, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *StaffResponse) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given map[string]interface{} and assigns it to the Schedule field.
func (o *StaffResponse) SetSchedule(v map[string]interface{}) {
	o.Schedule = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *StaffResponse) GetCatalog() map[string]interface{} {
	if o == nil || o.Catalog == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetCatalogOk() (map[string]interface{}, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *StaffResponse) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given map[string]interface{} and assigns it to the Catalog field.
func (o *StaffResponse) SetCatalog(v map[string]interface{}) {
	o.Catalog = v
}

// GetReviews returns the Reviews field value if set, zero value otherwise.
func (o *StaffResponse) GetReviews() map[string]interface{} {
	if o == nil || o.Reviews == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Reviews
}

// GetReviewsOk returns a tuple with the Reviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetReviewsOk() (map[string]interface{}, bool) {
	if o == nil || o.Reviews == nil {
		return nil, false
	}
	return o.Reviews, true
}

// HasReviews returns a boolean if a field has been set.
func (o *StaffResponse) HasReviews() bool {
	if o != nil && o.Reviews != nil {
		return true
	}

	return false
}

// SetReviews gets a reference to the given map[string]interface{} and assigns it to the Reviews field.
func (o *StaffResponse) SetReviews(v map[string]interface{}) {
	o.Reviews = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StaffResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StaffResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *StaffResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *StaffResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *StaffResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *StaffResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *StaffResponse) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaffResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *StaffResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *StaffResponse) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o StaffResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.Place != nil {
		toSerialize["place"] = o.Place
	}
	if o.PlaceRoles != nil {
		toSerialize["place_roles"] = o.PlaceRoles
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.CompanyRoles != nil {
		toSerialize["company_roles"] = o.CompanyRoles
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Catalog != nil {
		toSerialize["catalog"] = o.Catalog
	}
	if o.Reviews != nil {
		toSerialize["reviews"] = o.Reviews
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

type NullableStaffResponse struct {
	value *StaffResponse
	isSet bool
}

func (v NullableStaffResponse) Get() *StaffResponse {
	return v.value
}

func (v *NullableStaffResponse) Set(val *StaffResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStaffResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStaffResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaffResponse(val *StaffResponse) *NullableStaffResponse {
	return &NullableStaffResponse{value: val, isSet: true}
}

func (v NullableStaffResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaffResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


