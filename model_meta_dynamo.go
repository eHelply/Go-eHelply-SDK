/*
eHelply SDK - 1.1.95

eHelply SDK for SuperStack Services

API version: 1.1.95
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// MetaDynamo A meta from DynamoDB
type MetaDynamo struct {
	Basic *Basic `json:"basic,omitempty"`
	Detailed *Detailed `json:"detailed,omitempty"`
	Custom *MetaCustom `json:"custom,omitempty"`
	Dates *DatesMeta `json:"dates,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	Children []MetaChildren `json:"children,omitempty"`
	ParentUuid *string `json:"parent_uuid,omitempty"`
	Uuid string `json:"uuid"`
}

// NewMetaDynamo instantiates a new MetaDynamo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaDynamo(uuid string) *MetaDynamo {
	this := MetaDynamo{}
	this.Uuid = uuid
	return &this
}

// NewMetaDynamoWithDefaults instantiates a new MetaDynamo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaDynamoWithDefaults() *MetaDynamo {
	this := MetaDynamo{}
	return &this
}

// GetBasic returns the Basic field value if set, zero value otherwise.
func (o *MetaDynamo) GetBasic() Basic {
	if o == nil || o.Basic == nil {
		var ret Basic
		return ret
	}
	return *o.Basic
}

// GetBasicOk returns a tuple with the Basic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetBasicOk() (*Basic, bool) {
	if o == nil || o.Basic == nil {
		return nil, false
	}
	return o.Basic, true
}

// HasBasic returns a boolean if a field has been set.
func (o *MetaDynamo) HasBasic() bool {
	if o != nil && o.Basic != nil {
		return true
	}

	return false
}

// SetBasic gets a reference to the given Basic and assigns it to the Basic field.
func (o *MetaDynamo) SetBasic(v Basic) {
	o.Basic = &v
}

// GetDetailed returns the Detailed field value if set, zero value otherwise.
func (o *MetaDynamo) GetDetailed() Detailed {
	if o == nil || o.Detailed == nil {
		var ret Detailed
		return ret
	}
	return *o.Detailed
}

// GetDetailedOk returns a tuple with the Detailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetDetailedOk() (*Detailed, bool) {
	if o == nil || o.Detailed == nil {
		return nil, false
	}
	return o.Detailed, true
}

// HasDetailed returns a boolean if a field has been set.
func (o *MetaDynamo) HasDetailed() bool {
	if o != nil && o.Detailed != nil {
		return true
	}

	return false
}

// SetDetailed gets a reference to the given Detailed and assigns it to the Detailed field.
func (o *MetaDynamo) SetDetailed(v Detailed) {
	o.Detailed = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *MetaDynamo) GetCustom() MetaCustom {
	if o == nil || o.Custom == nil {
		var ret MetaCustom
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetCustomOk() (*MetaCustom, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *MetaDynamo) HasCustom() bool {
	if o != nil && o.Custom != nil {
		return true
	}

	return false
}

// SetCustom gets a reference to the given MetaCustom and assigns it to the Custom field.
func (o *MetaDynamo) SetCustom(v MetaCustom) {
	o.Custom = &v
}

// GetDates returns the Dates field value if set, zero value otherwise.
func (o *MetaDynamo) GetDates() DatesMeta {
	if o == nil || o.Dates == nil {
		var ret DatesMeta
		return ret
	}
	return *o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetDatesOk() (*DatesMeta, bool) {
	if o == nil || o.Dates == nil {
		return nil, false
	}
	return o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *MetaDynamo) HasDates() bool {
	if o != nil && o.Dates != nil {
		return true
	}

	return false
}

// SetDates gets a reference to the given DatesMeta and assigns it to the Dates field.
func (o *MetaDynamo) SetDates(v DatesMeta) {
	o.Dates = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MetaDynamo) GetFields() []Field {
	if o == nil || o.Fields == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetFieldsOk() ([]Field, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetaDynamo) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *MetaDynamo) SetFields(v []Field) {
	o.Fields = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *MetaDynamo) GetChildren() []MetaChildren {
	if o == nil || o.Children == nil {
		var ret []MetaChildren
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetChildrenOk() ([]MetaChildren, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *MetaDynamo) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []MetaChildren and assigns it to the Children field.
func (o *MetaDynamo) SetChildren(v []MetaChildren) {
	o.Children = v
}

// GetParentUuid returns the ParentUuid field value if set, zero value otherwise.
func (o *MetaDynamo) GetParentUuid() string {
	if o == nil || o.ParentUuid == nil {
		var ret string
		return ret
	}
	return *o.ParentUuid
}

// GetParentUuidOk returns a tuple with the ParentUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetParentUuidOk() (*string, bool) {
	if o == nil || o.ParentUuid == nil {
		return nil, false
	}
	return o.ParentUuid, true
}

// HasParentUuid returns a boolean if a field has been set.
func (o *MetaDynamo) HasParentUuid() bool {
	if o != nil && o.ParentUuid != nil {
		return true
	}

	return false
}

// SetParentUuid gets a reference to the given string and assigns it to the ParentUuid field.
func (o *MetaDynamo) SetParentUuid(v string) {
	o.ParentUuid = &v
}

// GetUuid returns the Uuid field value
func (o *MetaDynamo) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *MetaDynamo) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *MetaDynamo) SetUuid(v string) {
	o.Uuid = v
}

func (o MetaDynamo) MarshalJSON() ([]byte, error) {
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
	if o.Dates != nil {
		toSerialize["dates"] = o.Dates
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
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableMetaDynamo struct {
	value *MetaDynamo
	isSet bool
}

func (v NullableMetaDynamo) Get() *MetaDynamo {
	return v.value
}

func (v *NullableMetaDynamo) Set(val *MetaDynamo) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaDynamo) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaDynamo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaDynamo(val *MetaDynamo) *NullableMetaDynamo {
	return &NullableMetaDynamo{value: val, isSet: true}
}

func (v NullableMetaDynamo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaDynamo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


