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

// CategoryResponse **:param** uuid                                **type:** string **:param** name                                **type:** string **:param** project_uuid                        **type:** string or None  **:param** meta_uuid                           **type:** string or None
type CategoryResponse struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// NewCategoryResponse instantiates a new CategoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryResponse(uuid string, name string) *CategoryResponse {
	this := CategoryResponse{}
	this.Uuid = uuid
	this.Name = name
	return &this
}

// NewCategoryResponseWithDefaults instantiates a new CategoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryResponseWithDefaults() *CategoryResponse {
	this := CategoryResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *CategoryResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *CategoryResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *CategoryResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *CategoryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryResponse) SetName(v string) {
	o.Name = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *CategoryResponse) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *CategoryResponse) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *CategoryResponse) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *CategoryResponse) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryResponse) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *CategoryResponse) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *CategoryResponse) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CategoryResponse) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryResponse) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CategoryResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *CategoryResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

func (o CategoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.MetaUuid != nil {
		toSerialize["meta_uuid"] = o.MetaUuid
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryResponse struct {
	value *CategoryResponse
	isSet bool
}

func (v NullableCategoryResponse) Get() *CategoryResponse {
	return v.value
}

func (v *NullableCategoryResponse) Set(val *CategoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryResponse(val *CategoryResponse) *NullableCategoryResponse {
	return &NullableCategoryResponse{value: val, isSet: true}
}

func (v NullableCategoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


