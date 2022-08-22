/*
eHelply SDK - 1.1.103

eHelply SDK for SuperStack Services

API version: 1.1.103
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProductBase struct for ProductBase
type ProductBase struct {
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	CollectionUuid *string `json:"collection_uuid,omitempty"`
	ReviewGroupUuid *string `json:"review_group_uuid,omitempty"`
	Addons []string `json:"addons,omitempty"`
	Name *string `json:"name,omitempty"`
	Price int32 `json:"price"`
	QuantityForPublic int32 `json:"quantity_for_public"`
}

// NewProductBase instantiates a new ProductBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductBase(price int32, quantityForPublic int32) *ProductBase {
	this := ProductBase{}
	this.Price = price
	this.QuantityForPublic = quantityForPublic
	return &this
}

// NewProductBaseWithDefaults instantiates a new ProductBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductBaseWithDefaults() *ProductBase {
	this := ProductBase{}
	return &this
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *ProductBase) GetMetaData() map[string]interface{} {
	if o == nil || o.MetaData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBase) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *ProductBase) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *ProductBase) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetCollectionUuid returns the CollectionUuid field value if set, zero value otherwise.
func (o *ProductBase) GetCollectionUuid() string {
	if o == nil || o.CollectionUuid == nil {
		var ret string
		return ret
	}
	return *o.CollectionUuid
}

// GetCollectionUuidOk returns a tuple with the CollectionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBase) GetCollectionUuidOk() (*string, bool) {
	if o == nil || o.CollectionUuid == nil {
		return nil, false
	}
	return o.CollectionUuid, true
}

// HasCollectionUuid returns a boolean if a field has been set.
func (o *ProductBase) HasCollectionUuid() bool {
	if o != nil && o.CollectionUuid != nil {
		return true
	}

	return false
}

// SetCollectionUuid gets a reference to the given string and assigns it to the CollectionUuid field.
func (o *ProductBase) SetCollectionUuid(v string) {
	o.CollectionUuid = &v
}

// GetReviewGroupUuid returns the ReviewGroupUuid field value if set, zero value otherwise.
func (o *ProductBase) GetReviewGroupUuid() string {
	if o == nil || o.ReviewGroupUuid == nil {
		var ret string
		return ret
	}
	return *o.ReviewGroupUuid
}

// GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBase) GetReviewGroupUuidOk() (*string, bool) {
	if o == nil || o.ReviewGroupUuid == nil {
		return nil, false
	}
	return o.ReviewGroupUuid, true
}

// HasReviewGroupUuid returns a boolean if a field has been set.
func (o *ProductBase) HasReviewGroupUuid() bool {
	if o != nil && o.ReviewGroupUuid != nil {
		return true
	}

	return false
}

// SetReviewGroupUuid gets a reference to the given string and assigns it to the ReviewGroupUuid field.
func (o *ProductBase) SetReviewGroupUuid(v string) {
	o.ReviewGroupUuid = &v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *ProductBase) GetAddons() []string {
	if o == nil || o.Addons == nil {
		var ret []string
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBase) GetAddonsOk() ([]string, bool) {
	if o == nil || o.Addons == nil {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *ProductBase) HasAddons() bool {
	if o != nil && o.Addons != nil {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []string and assigns it to the Addons field.
func (o *ProductBase) SetAddons(v []string) {
	o.Addons = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductBase) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value
func (o *ProductBase) GetPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ProductBase) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *ProductBase) SetPrice(v int32) {
	o.Price = v
}

// GetQuantityForPublic returns the QuantityForPublic field value
func (o *ProductBase) GetQuantityForPublic() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuantityForPublic
}

// GetQuantityForPublicOk returns a tuple with the QuantityForPublic field value
// and a boolean to check if the value has been set.
func (o *ProductBase) GetQuantityForPublicOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuantityForPublic, true
}

// SetQuantityForPublic sets field value
func (o *ProductBase) SetQuantityForPublic(v int32) {
	o.QuantityForPublic = v
}

func (o ProductBase) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableProductBase struct {
	value *ProductBase
	isSet bool
}

func (v NullableProductBase) Get() *ProductBase {
	return v.value
}

func (v *NullableProductBase) Set(val *ProductBase) {
	v.value = val
	v.isSet = true
}

func (v NullableProductBase) IsSet() bool {
	return v.isSet
}

func (v *NullableProductBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductBase(val *ProductBase) *NullableProductBase {
	return &NullableProductBase{value: val, isSet: true}
}

func (v NullableProductBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


