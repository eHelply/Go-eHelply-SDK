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
	"time"
)

// User User object, contains all required parameters
type User struct {
	CognitoId *string `json:"cognito_id,omitempty"`
	CognitoData map[string]interface{} `json:"cognito_data,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Email *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Country *string `json:"country,omitempty"`
	Picture *string `json:"picture,omitempty"`
	GpsLocation map[string]interface{} `json:"gps_location,omitempty"`
	VerifiedLegalTerms *bool `json:"verified_legal_terms,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	DateDeleted *time.Time `json:"date_deleted,omitempty"`
	Uuid string `json:"uuid"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	FirstLogin *bool `json:"first_login,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(uuid string) *User {
	this := User{}
	var verifiedLegalTerms bool = false
	this.VerifiedLegalTerms = &verifiedLegalTerms
	this.Uuid = uuid
	var firstLogin bool = true
	this.FirstLogin = &firstLogin
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var verifiedLegalTerms bool = false
	this.VerifiedLegalTerms = &verifiedLegalTerms
	var firstLogin bool = true
	this.FirstLogin = &firstLogin
	return &this
}

// GetCognitoId returns the CognitoId field value if set, zero value otherwise.
func (o *User) GetCognitoId() string {
	if o == nil || o.CognitoId == nil {
		var ret string
		return ret
	}
	return *o.CognitoId
}

// GetCognitoIdOk returns a tuple with the CognitoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCognitoIdOk() (*string, bool) {
	if o == nil || o.CognitoId == nil {
		return nil, false
	}
	return o.CognitoId, true
}

// HasCognitoId returns a boolean if a field has been set.
func (o *User) HasCognitoId() bool {
	if o != nil && o.CognitoId != nil {
		return true
	}

	return false
}

// SetCognitoId gets a reference to the given string and assigns it to the CognitoId field.
func (o *User) SetCognitoId(v string) {
	o.CognitoId = &v
}

// GetCognitoData returns the CognitoData field value if set, zero value otherwise.
func (o *User) GetCognitoData() map[string]interface{} {
	if o == nil || o.CognitoData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CognitoData
}

// GetCognitoDataOk returns a tuple with the CognitoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCognitoDataOk() (map[string]interface{}, bool) {
	if o == nil || o.CognitoData == nil {
		return nil, false
	}
	return o.CognitoData, true
}

// HasCognitoData returns a boolean if a field has been set.
func (o *User) HasCognitoData() bool {
	if o != nil && o.CognitoData != nil {
		return true
	}

	return false
}

// SetCognitoData gets a reference to the given map[string]interface{} and assigns it to the CognitoData field.
func (o *User) SetCognitoData(v map[string]interface{}) {
	o.CognitoData = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *User) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *User) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *User) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *User) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *User) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *User) SetCountry(v string) {
	o.Country = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *User) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *User) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *User) SetPicture(v string) {
	o.Picture = &v
}

// GetGpsLocation returns the GpsLocation field value if set, zero value otherwise.
func (o *User) GetGpsLocation() map[string]interface{} {
	if o == nil || o.GpsLocation == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.GpsLocation
}

// GetGpsLocationOk returns a tuple with the GpsLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGpsLocationOk() (map[string]interface{}, bool) {
	if o == nil || o.GpsLocation == nil {
		return nil, false
	}
	return o.GpsLocation, true
}

// HasGpsLocation returns a boolean if a field has been set.
func (o *User) HasGpsLocation() bool {
	if o != nil && o.GpsLocation != nil {
		return true
	}

	return false
}

// SetGpsLocation gets a reference to the given map[string]interface{} and assigns it to the GpsLocation field.
func (o *User) SetGpsLocation(v map[string]interface{}) {
	o.GpsLocation = v
}

// GetVerifiedLegalTerms returns the VerifiedLegalTerms field value if set, zero value otherwise.
func (o *User) GetVerifiedLegalTerms() bool {
	if o == nil || o.VerifiedLegalTerms == nil {
		var ret bool
		return ret
	}
	return *o.VerifiedLegalTerms
}

// GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerifiedLegalTermsOk() (*bool, bool) {
	if o == nil || o.VerifiedLegalTerms == nil {
		return nil, false
	}
	return o.VerifiedLegalTerms, true
}

// HasVerifiedLegalTerms returns a boolean if a field has been set.
func (o *User) HasVerifiedLegalTerms() bool {
	if o != nil && o.VerifiedLegalTerms != nil {
		return true
	}

	return false
}

// SetVerifiedLegalTerms gets a reference to the given bool and assigns it to the VerifiedLegalTerms field.
func (o *User) SetVerifiedLegalTerms(v bool) {
	o.VerifiedLegalTerms = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *User) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *User) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *User) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *User) GetDateUpdated() time.Time {
	if o == nil || o.DateUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDateUpdatedOk() (*time.Time, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *User) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given time.Time and assigns it to the DateUpdated field.
func (o *User) SetDateUpdated(v time.Time) {
	o.DateUpdated = &v
}

// GetDateDeleted returns the DateDeleted field value if set, zero value otherwise.
func (o *User) GetDateDeleted() time.Time {
	if o == nil || o.DateDeleted == nil {
		var ret time.Time
		return ret
	}
	return *o.DateDeleted
}

// GetDateDeletedOk returns a tuple with the DateDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDateDeletedOk() (*time.Time, bool) {
	if o == nil || o.DateDeleted == nil {
		return nil, false
	}
	return o.DateDeleted, true
}

// HasDateDeleted returns a boolean if a field has been set.
func (o *User) HasDateDeleted() bool {
	if o != nil && o.DateDeleted != nil {
		return true
	}

	return false
}

// SetDateDeleted gets a reference to the given time.Time and assigns it to the DateDeleted field.
func (o *User) SetDateDeleted(v time.Time) {
	o.DateDeleted = &v
}

// GetUuid returns the Uuid field value
func (o *User) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *User) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *User) SetUuid(v string) {
	o.Uuid = v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *User) GetLastLogin() time.Time {
	if o == nil || o.LastLogin == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetFirstLogin returns the FirstLogin field value if set, zero value otherwise.
func (o *User) GetFirstLogin() bool {
	if o == nil || o.FirstLogin == nil {
		var ret bool
		return ret
	}
	return *o.FirstLogin
}

// GetFirstLoginOk returns a tuple with the FirstLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstLoginOk() (*bool, bool) {
	if o == nil || o.FirstLogin == nil {
		return nil, false
	}
	return o.FirstLogin, true
}

// HasFirstLogin returns a boolean if a field has been set.
func (o *User) HasFirstLogin() bool {
	if o != nil && o.FirstLogin != nil {
		return true
	}

	return false
}

// SetFirstLogin gets a reference to the given bool and assigns it to the FirstLogin field.
func (o *User) SetFirstLogin(v bool) {
	o.FirstLogin = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CognitoId != nil {
		toSerialize["cognito_id"] = o.CognitoId
	}
	if o.CognitoData != nil {
		toSerialize["cognito_data"] = o.CognitoData
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.GpsLocation != nil {
		toSerialize["gps_location"] = o.GpsLocation
	}
	if o.VerifiedLegalTerms != nil {
		toSerialize["verified_legal_terms"] = o.VerifiedLegalTerms
	}
	if o.DateCreated != nil {
		toSerialize["date_created"] = o.DateCreated
	}
	if o.DateUpdated != nil {
		toSerialize["date_updated"] = o.DateUpdated
	}
	if o.DateDeleted != nil {
		toSerialize["date_deleted"] = o.DateDeleted
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
	}
	if o.FirstLogin != nil {
		toSerialize["first_login"] = o.FirstLogin
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


