/*
eHelply SDK - 1.1.104

eHelply SDK for SuperStack Services

API version: 1.1.104
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// StripeAccountResponse struct for StripeAccountResponse
type StripeAccountResponse struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ProjectUuid string `json:"project_uuid"`
	StripeCustomerId string `json:"stripe_customer_id"`
}

// NewStripeAccountResponse instantiates a new StripeAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeAccountResponse(firstName string, lastName string, projectUuid string, stripeCustomerId string) *StripeAccountResponse {
	this := StripeAccountResponse{}
	this.FirstName = firstName
	this.LastName = lastName
	this.ProjectUuid = projectUuid
	this.StripeCustomerId = stripeCustomerId
	return &this
}

// NewStripeAccountResponseWithDefaults instantiates a new StripeAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeAccountResponseWithDefaults() *StripeAccountResponse {
	this := StripeAccountResponse{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *StripeAccountResponse) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *StripeAccountResponse) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *StripeAccountResponse) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *StripeAccountResponse) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *StripeAccountResponse) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *StripeAccountResponse) SetLastName(v string) {
	o.LastName = v
}

// GetProjectUuid returns the ProjectUuid field value
func (o *StripeAccountResponse) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *StripeAccountResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *StripeAccountResponse) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetStripeCustomerId returns the StripeCustomerId field value
func (o *StripeAccountResponse) GetStripeCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StripeCustomerId
}

// GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field value
// and a boolean to check if the value has been set.
func (o *StripeAccountResponse) GetStripeCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StripeCustomerId, true
}

// SetStripeCustomerId sets field value
func (o *StripeAccountResponse) SetStripeCustomerId(v string) {
	o.StripeCustomerId = v
}

func (o StripeAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if true {
		toSerialize["stripe_customer_id"] = o.StripeCustomerId
	}
	return json.Marshal(toSerialize)
}

type NullableStripeAccountResponse struct {
	value *StripeAccountResponse
	isSet bool
}

func (v NullableStripeAccountResponse) Get() *StripeAccountResponse {
	return v.value
}

func (v *NullableStripeAccountResponse) Set(val *StripeAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeAccountResponse(val *StripeAccountResponse) *NullableStripeAccountResponse {
	return &NullableStripeAccountResponse{value: val, isSet: true}
}

func (v NullableStripeAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


