/*
eHelply SDK - 1.1.102

eHelply SDK for SuperStack Services

API version: 1.1.102
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// PlaceBase **:param** name                                **type:** string **:param** summary                             **type:** string or None  **:param** public                              **type:** bool or None  **:param** meta                                **type:** dict or None  **:param** addresses                           **type:** PlaceAddress or None  **:param** contact                             **type:** ContactBase or None
type PlaceBase struct {
	Name string `json:"name"`
	Summary *string `json:"summary,omitempty"`
	Public *bool `json:"public,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Addresses []AddressBase `json:"addresses,omitempty"`
	Contact *ContactBase `json:"contact,omitempty"`
}

// NewPlaceBase instantiates a new PlaceBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceBase(name string) *PlaceBase {
	this := PlaceBase{}
	this.Name = name
	var public bool = true
	this.Public = &public
	return &this
}

// NewPlaceBaseWithDefaults instantiates a new PlaceBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceBaseWithDefaults() *PlaceBase {
	this := PlaceBase{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetName returns the Name field value
func (o *PlaceBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlaceBase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlaceBase) SetName(v string) {
	o.Name = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PlaceBase) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBase) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PlaceBase) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PlaceBase) SetSummary(v string) {
	o.Summary = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *PlaceBase) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBase) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *PlaceBase) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *PlaceBase) SetPublic(v bool) {
	o.Public = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PlaceBase) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBase) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PlaceBase) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *PlaceBase) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *PlaceBase) GetAddresses() []AddressBase {
	if o == nil || o.Addresses == nil {
		var ret []AddressBase
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBase) GetAddressesOk() ([]AddressBase, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *PlaceBase) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressBase and assigns it to the Addresses field.
func (o *PlaceBase) SetAddresses(v []AddressBase) {
	o.Addresses = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *PlaceBase) GetContact() ContactBase {
	if o == nil || o.Contact == nil {
		var ret ContactBase
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceBase) GetContactOk() (*ContactBase, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *PlaceBase) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactBase and assigns it to the Contact field.
func (o *PlaceBase) SetContact(v ContactBase) {
	o.Contact = &v
}

func (o PlaceBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullablePlaceBase struct {
	value *PlaceBase
	isSet bool
}

func (v NullablePlaceBase) Get() *PlaceBase {
	return v.value
}

func (v *NullablePlaceBase) Set(val *PlaceBase) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceBase) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceBase(val *PlaceBase) *NullablePlaceBase {
	return &NullablePlaceBase{value: val, isSet: true}
}

func (v NullablePlaceBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


