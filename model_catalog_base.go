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

// CatalogBase struct for CatalogBase
type CatalogBase struct {
	MetaData map[string]interface{} `json:"meta_data,omitempty"`
	Name *string `json:"name,omitempty"`
	Featured map[string]interface{} `json:"featured,omitempty"`
	SubCatalogs map[string]interface{} `json:"sub_catalogs,omitempty"`
}

// NewCatalogBase instantiates a new CatalogBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogBase() *CatalogBase {
	this := CatalogBase{}
	return &this
}

// NewCatalogBaseWithDefaults instantiates a new CatalogBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogBaseWithDefaults() *CatalogBase {
	this := CatalogBase{}
	return &this
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *CatalogBase) GetMetaData() map[string]interface{} {
	if o == nil || o.MetaData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogBase) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *CatalogBase) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *CatalogBase) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogBase) SetName(v string) {
	o.Name = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *CatalogBase) GetFeatured() map[string]interface{} {
	if o == nil || o.Featured == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogBase) GetFeaturedOk() (map[string]interface{}, bool) {
	if o == nil || o.Featured == nil {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *CatalogBase) HasFeatured() bool {
	if o != nil && o.Featured != nil {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given map[string]interface{} and assigns it to the Featured field.
func (o *CatalogBase) SetFeatured(v map[string]interface{}) {
	o.Featured = v
}

// GetSubCatalogs returns the SubCatalogs field value if set, zero value otherwise.
func (o *CatalogBase) GetSubCatalogs() map[string]interface{} {
	if o == nil || o.SubCatalogs == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SubCatalogs
}

// GetSubCatalogsOk returns a tuple with the SubCatalogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogBase) GetSubCatalogsOk() (map[string]interface{}, bool) {
	if o == nil || o.SubCatalogs == nil {
		return nil, false
	}
	return o.SubCatalogs, true
}

// HasSubCatalogs returns a boolean if a field has been set.
func (o *CatalogBase) HasSubCatalogs() bool {
	if o != nil && o.SubCatalogs != nil {
		return true
	}

	return false
}

// SetSubCatalogs gets a reference to the given map[string]interface{} and assigns it to the SubCatalogs field.
func (o *CatalogBase) SetSubCatalogs(v map[string]interface{}) {
	o.SubCatalogs = v
}

func (o CatalogBase) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableCatalogBase struct {
	value *CatalogBase
	isSet bool
}

func (v NullableCatalogBase) Get() *CatalogBase {
	return v.value
}

func (v *NullableCatalogBase) Set(val *CatalogBase) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogBase) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogBase(val *CatalogBase) *NullableCatalogBase {
	return &NullableCatalogBase{value: val, isSet: true}
}

func (v NullableCatalogBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


