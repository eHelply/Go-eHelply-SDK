/*
eHelply SDK - 1.1.111

eHelply SDK for SuperStack Services

API version: 1.1.111
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CatalogReturn struct for CatalogReturn
type CatalogReturn struct {
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	Name *string `json:"name,omitempty"`
	Featured map[string]interface{} `json:"featured,omitempty"`
	SubCatalogs map[string]interface{} `json:"sub_catalogs,omitempty"`
	Uuid string `json:"uuid"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
	ProjectUuid string `json:"project_uuid"`
	ProductUuids []string `json:"product_uuids,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewCatalogReturn instantiates a new CatalogReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogReturn(uuid string, projectUuid string) *CatalogReturn {
	this := CatalogReturn{}
	this.Uuid = uuid
	this.ProjectUuid = projectUuid
	return &this
}

// NewCatalogReturnWithDefaults instantiates a new CatalogReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogReturnWithDefaults() *CatalogReturn {
	this := CatalogReturn{}
	return &this
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *CatalogReturn) GetMetaData() map[string]interface{} {
	if o == nil || o.MetaData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *CatalogReturn) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *CatalogReturn) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogReturn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogReturn) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogReturn) SetName(v string) {
	o.Name = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *CatalogReturn) GetFeatured() map[string]interface{} {
	if o == nil || o.Featured == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetFeaturedOk() (map[string]interface{}, bool) {
	if o == nil || o.Featured == nil {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *CatalogReturn) HasFeatured() bool {
	if o != nil && o.Featured != nil {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given map[string]interface{} and assigns it to the Featured field.
func (o *CatalogReturn) SetFeatured(v map[string]interface{}) {
	o.Featured = v
}

// GetSubCatalogs returns the SubCatalogs field value if set, zero value otherwise.
func (o *CatalogReturn) GetSubCatalogs() map[string]interface{} {
	if o == nil || o.SubCatalogs == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SubCatalogs
}

// GetSubCatalogsOk returns a tuple with the SubCatalogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetSubCatalogsOk() (map[string]interface{}, bool) {
	if o == nil || o.SubCatalogs == nil {
		return nil, false
	}
	return o.SubCatalogs, true
}

// HasSubCatalogs returns a boolean if a field has been set.
func (o *CatalogReturn) HasSubCatalogs() bool {
	if o != nil && o.SubCatalogs != nil {
		return true
	}

	return false
}

// SetSubCatalogs gets a reference to the given map[string]interface{} and assigns it to the SubCatalogs field.
func (o *CatalogReturn) SetSubCatalogs(v map[string]interface{}) {
	o.SubCatalogs = v
}

// GetUuid returns the Uuid field value
func (o *CatalogReturn) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CatalogReturn) SetUuid(v string) {
	o.Uuid = v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *CatalogReturn) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *CatalogReturn) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *CatalogReturn) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetProjectUuid returns the ProjectUuid field value
func (o *CatalogReturn) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *CatalogReturn) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetProductUuids returns the ProductUuids field value if set, zero value otherwise.
func (o *CatalogReturn) GetProductUuids() []string {
	if o == nil || o.ProductUuids == nil {
		var ret []string
		return ret
	}
	return o.ProductUuids
}

// GetProductUuidsOk returns a tuple with the ProductUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetProductUuidsOk() ([]string, bool) {
	if o == nil || o.ProductUuids == nil {
		return nil, false
	}
	return o.ProductUuids, true
}

// HasProductUuids returns a boolean if a field has been set.
func (o *CatalogReturn) HasProductUuids() bool {
	if o != nil && o.ProductUuids != nil {
		return true
	}

	return false
}

// SetProductUuids gets a reference to the given []string and assigns it to the ProductUuids field.
func (o *CatalogReturn) SetProductUuids(v []string) {
	o.ProductUuids = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CatalogReturn) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CatalogReturn) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CatalogReturn) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CatalogReturn) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CatalogReturn) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CatalogReturn) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *CatalogReturn) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogReturn) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CatalogReturn) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *CatalogReturn) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o CatalogReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Featured != nil {
		toSerialize["featured"] = o.Featured
	}
	if o.SubCatalogs != nil {
		toSerialize["sub_catalogs"] = o.SubCatalogs
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if o.MetaUuid != nil {
		toSerialize["meta_uuid"] = o.MetaUuid
	}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.ProductUuids != nil {
		toSerialize["product_uuids"] = o.ProductUuids
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

type NullableCatalogReturn struct {
	value *CatalogReturn
	isSet bool
}

func (v NullableCatalogReturn) Get() *CatalogReturn {
	return v.value
}

func (v *NullableCatalogReturn) Set(val *CatalogReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogReturn(val *CatalogReturn) *NullableCatalogReturn {
	return &NullableCatalogReturn{value: val, isSet: true}
}

func (v NullableCatalogReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


