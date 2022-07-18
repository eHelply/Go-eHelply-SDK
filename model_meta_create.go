/*
eHelply SDK - 1.1.87

eHelply SDK for SuperStack Services

API version: 1.1.87
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// MetaCreate Meta
type MetaCreate struct {
	Basic *BasicMetaCreate `json:"basic,omitempty"`
	Detailed *DetailedMetaCreate `json:"detailed,omitempty"`
	Custom *MetaCustom `json:"custom,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	Children []MetaChildren `json:"children,omitempty"`
	ParentUuid *string `json:"parent_uuid,omitempty"`
}

// NewMetaCreate instantiates a new MetaCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaCreate() *MetaCreate {
	this := MetaCreate{}
	return &this
}

// NewMetaCreateWithDefaults instantiates a new MetaCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaCreateWithDefaults() *MetaCreate {
	this := MetaCreate{}
	return &this
}

// GetBasic returns the Basic field value if set, zero value otherwise.
func (o *MetaCreate) GetBasic() BasicMetaCreate {
	if o == nil || o.Basic == nil {
		var ret BasicMetaCreate
		return ret
	}
	return *o.Basic
}

// GetBasicOk returns a tuple with the Basic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCreate) GetBasicOk() (*BasicMetaCreate, bool) {
	if o == nil || o.Basic == nil {
		return nil, false
	}
	return o.Basic, true
}

// HasBasic returns a boolean if a field has been set.
func (o *MetaCreate) HasBasic() bool {
	if o != nil && o.Basic != nil {
		return true
	}

	return false
}

// SetBasic gets a reference to the given BasicMetaCreate and assigns it to the Basic field.
func (o *MetaCreate) SetBasic(v BasicMetaCreate) {
	o.Basic = &v
}

// GetDetailed returns the Detailed field value if set, zero value otherwise.
func (o *MetaCreate) GetDetailed() DetailedMetaCreate {
	if o == nil || o.Detailed == nil {
		var ret DetailedMetaCreate
		return ret
	}
	return *o.Detailed
}

// GetDetailedOk returns a tuple with the Detailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCreate) GetDetailedOk() (*DetailedMetaCreate, bool) {
	if o == nil || o.Detailed == nil {
		return nil, false
	}
	return o.Detailed, true
}

// HasDetailed returns a boolean if a field has been set.
func (o *MetaCreate) HasDetailed() bool {
	if o != nil && o.Detailed != nil {
		return true
	}

	return false
}

// SetDetailed gets a reference to the given DetailedMetaCreate and assigns it to the Detailed field.
func (o *MetaCreate) SetDetailed(v DetailedMetaCreate) {
	o.Detailed = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *MetaCreate) GetCustom() MetaCustom {
	if o == nil || o.Custom == nil {
		var ret MetaCustom
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCreate) GetCustomOk() (*MetaCustom, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *MetaCreate) HasCustom() bool {
	if o != nil && o.Custom != nil {
		return true
	}

	return false
}

// SetCustom gets a reference to the given MetaCustom and assigns it to the Custom field.
func (o *MetaCreate) SetCustom(v MetaCustom) {
	o.Custom = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MetaCreate) GetFields() []Field {
	if o == nil || o.Fields == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCreate) GetFieldsOk() ([]Field, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetaCreate) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *MetaCreate) SetFields(v []Field) {
	o.Fields = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *MetaCreate) GetChildren() []MetaChildren {
	if o == nil || o.Children == nil {
		var ret []MetaChildren
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCreate) GetChildrenOk() ([]MetaChildren, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *MetaCreate) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []MetaChildren and assigns it to the Children field.
func (o *MetaCreate) SetChildren(v []MetaChildren) {
	o.Children = v
}

// GetParentUuid returns the ParentUuid field value if set, zero value otherwise.
func (o *MetaCreate) GetParentUuid() string {
	if o == nil || o.ParentUuid == nil {
		var ret string
		return ret
	}
	return *o.ParentUuid
}

// GetParentUuidOk returns a tuple with the ParentUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaCreate) GetParentUuidOk() (*string, bool) {
	if o == nil || o.ParentUuid == nil {
		return nil, false
	}
	return o.ParentUuid, true
}

// HasParentUuid returns a boolean if a field has been set.
func (o *MetaCreate) HasParentUuid() bool {
	if o != nil && o.ParentUuid != nil {
		return true
	}

	return false
}

// SetParentUuid gets a reference to the given string and assigns it to the ParentUuid field.
func (o *MetaCreate) SetParentUuid(v string) {
	o.ParentUuid = &v
}

func (o MetaCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Basic != nil {
		toSerialize["basic"] = o.Basic
	}
	if o.Detailed != nil {
		toSerialize["detailed"] = o.Detailed
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.ParentUuid != nil {
		toSerialize["parent_uuid"] = o.ParentUuid
	}
	return json.Marshal(toSerialize)
}

type NullableMetaCreate struct {
	value *MetaCreate
	isSet bool
}

func (v NullableMetaCreate) Get() *MetaCreate {
	return v.value
}

func (v *NullableMetaCreate) Set(val *MetaCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaCreate(val *MetaCreate) *NullableMetaCreate {
	return &NullableMetaCreate{value: val, isSet: true}
}

func (v NullableMetaCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


