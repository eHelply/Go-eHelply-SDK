/*
eHelply SDK - 1.1.109

eHelply SDK for SuperStack Services

API version: 1.1.109
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ContactResponse struct for ContactResponse
type ContactResponse struct {
	ContactId string `json:"contact_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// NewContactResponse instantiates a new ContactResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactResponse(contactId string, firstName string, lastName string, email string, phone string) *ContactResponse {
	this := ContactResponse{}
	this.ContactId = contactId
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.Phone = phone
	return &this
}

// NewContactResponseWithDefaults instantiates a new ContactResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactResponseWithDefaults() *ContactResponse {
	this := ContactResponse{}
	return &this
}

// GetContactId returns the ContactId field value
func (o *ContactResponse) GetContactId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetContactIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactId, true
}

// SetContactId sets field value
func (o *ContactResponse) SetContactId(v string) {
	o.ContactId = v
}

// GetFirstName returns the FirstName field value
func (o *ContactResponse) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ContactResponse) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *ContactResponse) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ContactResponse) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *ContactResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ContactResponse) SetEmail(v string) {
	o.Email = v
}

// GetPhone returns the Phone field value
func (o *ContactResponse) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *ContactResponse) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *ContactResponse) SetPhone(v string) {
	o.Phone = v
}

func (o ContactResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contact_id"] = o.ContactId
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableContactResponse struct {
	value *ContactResponse
	isSet bool
}

func (v NullableContactResponse) Get() *ContactResponse {
	return v.value
}

func (v *NullableContactResponse) Set(val *ContactResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContactResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContactResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactResponse(val *ContactResponse) *NullableContactResponse {
	return &NullableContactResponse{value: val, isSet: true}
}

func (v NullableContactResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


