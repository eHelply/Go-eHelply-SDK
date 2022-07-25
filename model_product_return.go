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

// ProductReturn struct for ProductReturn
type ProductReturn struct {
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	CollectionUuid *string `json:"collection_uuid,omitempty"`
	ReviewGroupUuid *string `json:"review_group_uuid,omitempty"`
	Addons []string `json:"addons,omitempty"`
	Name *string `json:"name,omitempty"`
	Price int32 `json:"price"`
	QuantityForPublic int32 `json:"quantity_for_public"`
	Uuid string `json:"uuid"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
	ProjectUuid string `json:"project_uuid"`
	CatalogUuid *string `json:"catalog_uuid,omitempty"`
	AddonList []map[string]interface{} `json:"addon_list,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewProductReturn instantiates a new ProductReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductReturn(price int32, quantityForPublic int32, uuid string, projectUuid string) *ProductReturn {
	this := ProductReturn{}
	this.Price = price
	this.QuantityForPublic = quantityForPublic
	this.Uuid = uuid
	this.ProjectUuid = projectUuid
	return &this
}

// NewProductReturnWithDefaults instantiates a new ProductReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductReturnWithDefaults() *ProductReturn {
	this := ProductReturn{}
	return &this
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *ProductReturn) GetMetaData() map[string]interface{} {
	if o == nil || o.MetaData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *ProductReturn) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *ProductReturn) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetCollectionUuid returns the CollectionUuid field value if set, zero value otherwise.
func (o *ProductReturn) GetCollectionUuid() string {
	if o == nil || o.CollectionUuid == nil {
		var ret string
		return ret
	}
	return *o.CollectionUuid
}

// GetCollectionUuidOk returns a tuple with the CollectionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetCollectionUuidOk() (*string, bool) {
	if o == nil || o.CollectionUuid == nil {
		return nil, false
	}
	return o.CollectionUuid, true
}

// HasCollectionUuid returns a boolean if a field has been set.
func (o *ProductReturn) HasCollectionUuid() bool {
	if o != nil && o.CollectionUuid != nil {
		return true
	}

	return false
}

// SetCollectionUuid gets a reference to the given string and assigns it to the CollectionUuid field.
func (o *ProductReturn) SetCollectionUuid(v string) {
	o.CollectionUuid = &v
}

// GetReviewGroupUuid returns the ReviewGroupUuid field value if set, zero value otherwise.
func (o *ProductReturn) GetReviewGroupUuid() string {
	if o == nil || o.ReviewGroupUuid == nil {
		var ret string
		return ret
	}
	return *o.ReviewGroupUuid
}

// GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetReviewGroupUuidOk() (*string, bool) {
	if o == nil || o.ReviewGroupUuid == nil {
		return nil, false
	}
	return o.ReviewGroupUuid, true
}

// HasReviewGroupUuid returns a boolean if a field has been set.
func (o *ProductReturn) HasReviewGroupUuid() bool {
	if o != nil && o.ReviewGroupUuid != nil {
		return true
	}

	return false
}

// SetReviewGroupUuid gets a reference to the given string and assigns it to the ReviewGroupUuid field.
func (o *ProductReturn) SetReviewGroupUuid(v string) {
	o.ReviewGroupUuid = &v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *ProductReturn) GetAddons() []string {
	if o == nil || o.Addons == nil {
		var ret []string
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetAddonsOk() ([]string, bool) {
	if o == nil || o.Addons == nil {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *ProductReturn) HasAddons() bool {
	if o != nil && o.Addons != nil {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []string and assigns it to the Addons field.
func (o *ProductReturn) SetAddons(v []string) {
	o.Addons = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductReturn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductReturn) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductReturn) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value
func (o *ProductReturn) GetPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *ProductReturn) SetPrice(v int32) {
	o.Price = v
}

// GetQuantityForPublic returns the QuantityForPublic field value
func (o *ProductReturn) GetQuantityForPublic() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuantityForPublic
}

// GetQuantityForPublicOk returns a tuple with the QuantityForPublic field value
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetQuantityForPublicOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuantityForPublic, true
}

// SetQuantityForPublic sets field value
func (o *ProductReturn) SetQuantityForPublic(v int32) {
	o.QuantityForPublic = v
}

// GetUuid returns the Uuid field value
func (o *ProductReturn) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ProductReturn) SetUuid(v string) {
	o.Uuid = v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *ProductReturn) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *ProductReturn) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *ProductReturn) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetProjectUuid returns the ProjectUuid field value
func (o *ProductReturn) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *ProductReturn) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetCatalogUuid returns the CatalogUuid field value if set, zero value otherwise.
func (o *ProductReturn) GetCatalogUuid() string {
	if o == nil || o.CatalogUuid == nil {
		var ret string
		return ret
	}
	return *o.CatalogUuid
}

// GetCatalogUuidOk returns a tuple with the CatalogUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetCatalogUuidOk() (*string, bool) {
	if o == nil || o.CatalogUuid == nil {
		return nil, false
	}
	return o.CatalogUuid, true
}

// HasCatalogUuid returns a boolean if a field has been set.
func (o *ProductReturn) HasCatalogUuid() bool {
	if o != nil && o.CatalogUuid != nil {
		return true
	}

	return false
}

// SetCatalogUuid gets a reference to the given string and assigns it to the CatalogUuid field.
func (o *ProductReturn) SetCatalogUuid(v string) {
	o.CatalogUuid = &v
}

// GetAddonList returns the AddonList field value if set, zero value otherwise.
func (o *ProductReturn) GetAddonList() []map[string]interface{} {
	if o == nil || o.AddonList == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.AddonList
}

// GetAddonListOk returns a tuple with the AddonList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetAddonListOk() ([]map[string]interface{}, bool) {
	if o == nil || o.AddonList == nil {
		return nil, false
	}
	return o.AddonList, true
}

// HasAddonList returns a boolean if a field has been set.
func (o *ProductReturn) HasAddonList() bool {
	if o != nil && o.AddonList != nil {
		return true
	}

	return false
}

// SetAddonList gets a reference to the given []map[string]interface{} and assigns it to the AddonList field.
func (o *ProductReturn) SetAddonList(v []map[string]interface{}) {
	o.AddonList = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductReturn) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductReturn) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ProductReturn) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProductReturn) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductReturn) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ProductReturn) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ProductReturn) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReturn) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ProductReturn) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *ProductReturn) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o ProductReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if o.CollectionUuid != nil {
		toSerialize["collection_uuid"] = o.CollectionUuid
	}
	if o.ReviewGroupUuid != nil {
		toSerialize["review_group_uuid"] = o.ReviewGroupUuid
	}
	if o.Addons != nil {
		toSerialize["addons"] = o.Addons
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["quantity_for_public"] = o.QuantityForPublic
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
	if o.CatalogUuid != nil {
		toSerialize["catalog_uuid"] = o.CatalogUuid
	}
	if o.AddonList != nil {
		toSerialize["addon_list"] = o.AddonList
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

type NullableProductReturn struct {
	value *ProductReturn
	isSet bool
}

func (v NullableProductReturn) Get() *ProductReturn {
	return v.value
}

func (v *NullableProductReturn) Set(val *ProductReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableProductReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableProductReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductReturn(val *ProductReturn) *NullableProductReturn {
	return &NullableProductReturn{value: val, isSet: true}
}

func (v NullableProductReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


