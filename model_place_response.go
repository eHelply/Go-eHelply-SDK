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

// PlaceResponse **:param** name                                **type:** string **:param** summary                             **type:** string or None  **:param** public                              **type:** bool or None  **:param** meta                                **type:** dict or None  **:param** addresses                           **type:** PlaceAddress or None  **:param** contact                             **type:** ContactBase or None
type PlaceResponse struct {
	ProjectUuid *string `json:"project_uuid,omitempty"`
	Name string `json:"name"`
	Summary *string `json:"summary,omitempty"`
	Public *bool `json:"public,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Addresses []AddressBase `json:"addresses,omitempty"`
	Contact *ContactBase `json:"contact,omitempty"`
	Picture *string `json:"picture,omitempty"`
	Uuid string `json:"uuid"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
	Tags []TagBase `json:"tags,omitempty"`
	Categories []CategoryBase `json:"categories,omitempty"`
	Company *Company `json:"company,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewPlaceResponse instantiates a new PlaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceResponse(name string, uuid string) *PlaceResponse {
	this := PlaceResponse{}
	this.Name = name
	var public bool = true
	this.Public = &public
	this.Uuid = uuid
	return &this
}

// NewPlaceResponseWithDefaults instantiates a new PlaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceResponseWithDefaults() *PlaceResponse {
	this := PlaceResponse{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *PlaceResponse) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *PlaceResponse) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *PlaceResponse) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetName returns the Name field value
func (o *PlaceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlaceResponse) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PlaceResponse) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PlaceResponse) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PlaceResponse) SetSummary(v string) {
	o.Summary = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *PlaceResponse) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *PlaceResponse) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *PlaceResponse) SetPublic(v bool) {
	o.Public = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PlaceResponse) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PlaceResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *PlaceResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *PlaceResponse) GetAddresses() []AddressBase {
	if o == nil || o.Addresses == nil {
		var ret []AddressBase
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetAddressesOk() ([]AddressBase, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *PlaceResponse) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressBase and assigns it to the Addresses field.
func (o *PlaceResponse) SetAddresses(v []AddressBase) {
	o.Addresses = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *PlaceResponse) GetContact() ContactBase {
	if o == nil || o.Contact == nil {
		var ret ContactBase
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetContactOk() (*ContactBase, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *PlaceResponse) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactBase and assigns it to the Contact field.
func (o *PlaceResponse) SetContact(v ContactBase) {
	o.Contact = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *PlaceResponse) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *PlaceResponse) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *PlaceResponse) SetPicture(v string) {
	o.Picture = &v
}

// GetUuid returns the Uuid field value
func (o *PlaceResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PlaceResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *PlaceResponse) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *PlaceResponse) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *PlaceResponse) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PlaceResponse) GetTags() []TagBase {
	if o == nil || o.Tags == nil {
		var ret []TagBase
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetTagsOk() ([]TagBase, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PlaceResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagBase and assigns it to the Tags field.
func (o *PlaceResponse) SetTags(v []TagBase) {
	o.Tags = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *PlaceResponse) GetCategories() []CategoryBase {
	if o == nil || o.Categories == nil {
		var ret []CategoryBase
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetCategoriesOk() ([]CategoryBase, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *PlaceResponse) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []CategoryBase and assigns it to the Categories field.
func (o *PlaceResponse) SetCategories(v []CategoryBase) {
	o.Categories = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *PlaceResponse) GetCompany() Company {
	if o == nil || o.Company == nil {
		var ret Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetCompanyOk() (*Company, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *PlaceResponse) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given Company and assigns it to the Company field.
func (o *PlaceResponse) SetCompany(v Company) {
	o.Company = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PlaceResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PlaceResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PlaceResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PlaceResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PlaceResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PlaceResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *PlaceResponse) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *PlaceResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *PlaceResponse) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o PlaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if true {
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
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if true {
		toSerialize["uuid"] = o.Uuid
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
	if o.Company != nil {
		toSerialize["company"] = o.Company
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

type NullablePlaceResponse struct {
	value *PlaceResponse
	isSet bool
}

func (v NullablePlaceResponse) Get() *PlaceResponse {
	return v.value
}

func (v *NullablePlaceResponse) Set(val *PlaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceResponse(val *PlaceResponse) *NullablePlaceResponse {
	return &NullablePlaceResponse{value: val, isSet: true}
}

func (v NullablePlaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


