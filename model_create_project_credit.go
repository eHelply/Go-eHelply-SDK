/*
eHelply SDK - 1.1.114

eHelply SDK for SuperStack Services

API version: 1.1.114
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CreateProjectCredit struct for CreateProjectCredit
type CreateProjectCredit struct {
	CreditsGranted int32 `json:"credits_granted"`
	GrantedReason string `json:"granted_reason"`
}

// NewCreateProjectCredit instantiates a new CreateProjectCredit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectCredit(creditsGranted int32, grantedReason string) *CreateProjectCredit {
	this := CreateProjectCredit{}
	this.CreditsGranted = creditsGranted
	this.GrantedReason = grantedReason
	return &this
}

// NewCreateProjectCreditWithDefaults instantiates a new CreateProjectCredit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectCreditWithDefaults() *CreateProjectCredit {
	this := CreateProjectCredit{}
	return &this
}

// GetCreditsGranted returns the CreditsGranted field value
func (o *CreateProjectCredit) GetCreditsGranted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreditsGranted
}

// GetCreditsGrantedOk returns a tuple with the CreditsGranted field value
// and a boolean to check if the value has been set.
func (o *CreateProjectCredit) GetCreditsGrantedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsGranted, true
}

// SetCreditsGranted sets field value
func (o *CreateProjectCredit) SetCreditsGranted(v int32) {
	o.CreditsGranted = v
}

// GetGrantedReason returns the GrantedReason field value
func (o *CreateProjectCredit) GetGrantedReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantedReason
}

// GetGrantedReasonOk returns a tuple with the GrantedReason field value
// and a boolean to check if the value has been set.
func (o *CreateProjectCredit) GetGrantedReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantedReason, true
}

// SetGrantedReason sets field value
func (o *CreateProjectCredit) SetGrantedReason(v string) {
	o.GrantedReason = v
}

func (o CreateProjectCredit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credits_granted"] = o.CreditsGranted
	}
	if true {
		toSerialize["granted_reason"] = o.GrantedReason
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProjectCredit struct {
	value *CreateProjectCredit
	isSet bool
}

func (v NullableCreateProjectCredit) Get() *CreateProjectCredit {
	return v.value
}

func (v *NullableCreateProjectCredit) Set(val *CreateProjectCredit) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectCredit(val *CreateProjectCredit) *NullableCreateProjectCredit {
	return &NullableCreateProjectCredit{value: val, isSet: true}
}

func (v NullableCreateProjectCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


