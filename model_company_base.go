/*
eHelply SDK - 1.1.89

eHelply SDK for SuperStack Services

API version: 1.1.89
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CompanyBase **:param** name                                **type:** string **:param** summary                             **type:** string or None  **:param** public                              **type:** bool or None  **:param** meta                                **type:** dict or None  **:param** contact                             **type:** ContactBase or None
type CompanyBase struct {
	Name *string `json:"name,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Public *bool `json:"public,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Contact *ContactBase `json:"contact,omitempty"`
}

// NewCompanyBase instantiates a new CompanyBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyBase() *CompanyBase {
	this := CompanyBase{}
	var public bool = true
	this.Public = &public
	return &this
}

// NewCompanyBaseWithDefaults instantiates a new CompanyBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyBaseWithDefaults() *CompanyBase {
	this := CompanyBase{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanyBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanyBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanyBase) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *CompanyBase) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyBase) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *CompanyBase) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *CompanyBase) SetSummary(v string) {
	o.Summary = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *CompanyBase) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyBase) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *CompanyBase) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *CompanyBase) SetPublic(v bool) {
	o.Public = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CompanyBase) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyBase) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CompanyBase) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *CompanyBase) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *CompanyBase) GetContact() ContactBase {
	if o == nil || o.Contact == nil {
		var ret ContactBase
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyBase) GetContactOk() (*ContactBase, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *CompanyBase) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactBase and assigns it to the Contact field.
func (o *CompanyBase) SetContact(v ContactBase) {
	o.Contact = &v
}

func (o CompanyBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyBase struct {
	value *CompanyBase
	isSet bool
}

func (v NullableCompanyBase) Get() *CompanyBase {
	return v.value
}

func (v *NullableCompanyBase) Set(val *CompanyBase) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyBase) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyBase(val *CompanyBase) *NullableCompanyBase {
	return &NullableCompanyBase{value: val, isSet: true}
}

func (v NullableCompanyBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


