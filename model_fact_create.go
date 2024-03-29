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

// FactCreate struct for FactCreate
type FactCreate struct {
	Name string `json:"name"`
	Data map[string]interface{} `json:"data"`
	Public bool `json:"public"`
}

// NewFactCreate instantiates a new FactCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFactCreate(name string, data map[string]interface{}, public bool) *FactCreate {
	this := FactCreate{}
	this.Name = name
	this.Data = data
	this.Public = public
	return &this
}

// NewFactCreateWithDefaults instantiates a new FactCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactCreateWithDefaults() *FactCreate {
	this := FactCreate{}
	return &this
}

// GetName returns the Name field value
func (o *FactCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FactCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FactCreate) SetName(v string) {
	o.Name = v
}

// GetData returns the Data field value
func (o *FactCreate) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FactCreate) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *FactCreate) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetPublic returns the Public field value
func (o *FactCreate) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *FactCreate) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *FactCreate) SetPublic(v bool) {
	o.Public = v
}

func (o FactCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["public"] = o.Public
	}
	return json.Marshal(toSerialize)
}

type NullableFactCreate struct {
	value *FactCreate
	isSet bool
}

func (v NullableFactCreate) Get() *FactCreate {
	return v.value
}

func (v *NullableFactCreate) Set(val *FactCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFactCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFactCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFactCreate(val *FactCreate) *NullableFactCreate {
	return &NullableFactCreate{value: val, isSet: true}
}

func (v NullableFactCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFactCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


