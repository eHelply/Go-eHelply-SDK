/*
eHelply SDK - 1.1.97

eHelply SDK for SuperStack Services

API version: 1.1.97
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ContactBase **:param** phones                              **type:** List[ContactKeys] or None  **:param** email                               **type:** string or None  **:param** website                             **type:** string or None  **:param** socials                             **type:** List[ContactKeys] or None
type ContactBase struct {
	Phones []ContactMethod `json:"phones,omitempty"`
	Email *string `json:"email,omitempty"`
	Website *string `json:"website,omitempty"`
	Socials []ContactMethod `json:"socials,omitempty"`
}

// NewContactBase instantiates a new ContactBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactBase() *ContactBase {
	this := ContactBase{}
	return &this
}

// NewContactBaseWithDefaults instantiates a new ContactBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactBaseWithDefaults() *ContactBase {
	this := ContactBase{}
	return &this
}

// GetPhones returns the Phones field value if set, zero value otherwise.
func (o *ContactBase) GetPhones() []ContactMethod {
	if o == nil || o.Phones == nil {
		var ret []ContactMethod
		return ret
	}
	return o.Phones
}

// GetPhonesOk returns a tuple with the Phones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactBase) GetPhonesOk() ([]ContactMethod, bool) {
	if o == nil || o.Phones == nil {
		return nil, false
	}
	return o.Phones, true
}

// HasPhones returns a boolean if a field has been set.
func (o *ContactBase) HasPhones() bool {
	if o != nil && o.Phones != nil {
		return true
	}

	return false
}

// SetPhones gets a reference to the given []ContactMethod and assigns it to the Phones field.
func (o *ContactBase) SetPhones(v []ContactMethod) {
	o.Phones = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ContactBase) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactBase) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactBase) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ContactBase) SetEmail(v string) {
	o.Email = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ContactBase) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactBase) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ContactBase) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *ContactBase) SetWebsite(v string) {
	o.Website = &v
}

// GetSocials returns the Socials field value if set, zero value otherwise.
func (o *ContactBase) GetSocials() []ContactMethod {
	if o == nil || o.Socials == nil {
		var ret []ContactMethod
		return ret
	}
	return o.Socials
}

// GetSocialsOk returns a tuple with the Socials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactBase) GetSocialsOk() ([]ContactMethod, bool) {
	if o == nil || o.Socials == nil {
		return nil, false
	}
	return o.Socials, true
}

// HasSocials returns a boolean if a field has been set.
func (o *ContactBase) HasSocials() bool {
	if o != nil && o.Socials != nil {
		return true
	}

	return false
}

// SetSocials gets a reference to the given []ContactMethod and assigns it to the Socials field.
func (o *ContactBase) SetSocials(v []ContactMethod) {
	o.Socials = v
}

func (o ContactBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Phones != nil {
		toSerialize["phones"] = o.Phones
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Socials != nil {
		toSerialize["socials"] = o.Socials
	}
	return json.Marshal(toSerialize)
}

type NullableContactBase struct {
	value *ContactBase
	isSet bool
}

func (v NullableContactBase) Get() *ContactBase {
	return v.value
}

func (v *NullableContactBase) Set(val *ContactBase) {
	v.value = val
	v.isSet = true
}

func (v NullableContactBase) IsSet() bool {
	return v.isSet
}

func (v *NullableContactBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactBase(val *ContactBase) *NullableContactBase {
	return &NullableContactBase{value: val, isSet: true}
}

func (v NullableContactBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


