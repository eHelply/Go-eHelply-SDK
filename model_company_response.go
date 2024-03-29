/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CompanyResponse **:param** name                                **type:** string **:param** summary                             **type:** string or None  **:param** public                              **type:** bool or None  **:param** meta                                **type:** dict or None  **:param** contact                             **type:** ContactBase or None
type CompanyResponse struct {
	Name *string `json:"name,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Public *bool `json:"public,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Contact *ContactBase `json:"contact,omitempty"`
	Uuid string `json:"uuid"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
	Tags []TagBase `json:"tags,omitempty"`
	Categories []CategoryBase `json:"categories,omitempty"`
	Places []PlaceBase `json:"places,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewCompanyResponse instantiates a new CompanyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyResponse(uuid string) *CompanyResponse {
	this := CompanyResponse{}
	var public bool = true
	this.Public = &public
	this.Uuid = uuid
	return &this
}

// NewCompanyResponseWithDefaults instantiates a new CompanyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyResponseWithDefaults() *CompanyResponse {
	this := CompanyResponse{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanyResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanyResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanyResponse) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *CompanyResponse) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *CompanyResponse) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *CompanyResponse) SetSummary(v string) {
	o.Summary = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *CompanyResponse) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *CompanyResponse) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *CompanyResponse) SetPublic(v bool) {
	o.Public = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CompanyResponse) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CompanyResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *CompanyResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *CompanyResponse) GetContact() ContactBase {
	if o == nil || o.Contact == nil {
		var ret ContactBase
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetContactOk() (*ContactBase, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *CompanyResponse) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactBase and assigns it to the Contact field.
func (o *CompanyResponse) SetContact(v ContactBase) {
	o.Contact = &v
}

// GetUuid returns the Uuid field value
func (o *CompanyResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CompanyResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *CompanyResponse) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *CompanyResponse) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *CompanyResponse) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *CompanyResponse) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *CompanyResponse) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *CompanyResponse) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CompanyResponse) GetTags() []TagBase {
	if o == nil || o.Tags == nil {
		var ret []TagBase
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetTagsOk() ([]TagBase, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CompanyResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagBase and assigns it to the Tags field.
func (o *CompanyResponse) SetTags(v []TagBase) {
	o.Tags = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CompanyResponse) GetCategories() []CategoryBase {
	if o == nil || o.Categories == nil {
		var ret []CategoryBase
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetCategoriesOk() ([]CategoryBase, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CompanyResponse) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []CategoryBase and assigns it to the Categories field.
func (o *CompanyResponse) SetCategories(v []CategoryBase) {
	o.Categories = v
}

// GetPlaces returns the Places field value if set, zero value otherwise.
func (o *CompanyResponse) GetPlaces() []PlaceBase {
	if o == nil || o.Places == nil {
		var ret []PlaceBase
		return ret
	}
	return o.Places
}

// GetPlacesOk returns a tuple with the Places field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetPlacesOk() ([]PlaceBase, bool) {
	if o == nil || o.Places == nil {
		return nil, false
	}
	return o.Places, true
}

// HasPlaces returns a boolean if a field has been set.
func (o *CompanyResponse) HasPlaces() bool {
	if o != nil && o.Places != nil {
		return true
	}

	return false
}

// SetPlaces gets a reference to the given []PlaceBase and assigns it to the Places field.
func (o *CompanyResponse) SetPlaces(v []PlaceBase) {
	o.Places = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CompanyResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CompanyResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CompanyResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CompanyResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CompanyResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CompanyResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *CompanyResponse) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CompanyResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *CompanyResponse) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o CompanyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.MetaUuid != nil {
		toSerialize["meta_uuid"] = o.MetaUuid
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Places != nil {
		toSerialize["places"] = o.Places
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

type NullableCompanyResponse struct {
	value *CompanyResponse
	isSet bool
}

func (v NullableCompanyResponse) Get() *CompanyResponse {
	return v.value
}

func (v *NullableCompanyResponse) Set(val *CompanyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyResponse(val *CompanyResponse) *NullableCompanyResponse {
	return &NullableCompanyResponse{value: val, isSet: true}
}

func (v NullableCompanyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


