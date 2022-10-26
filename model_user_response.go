/*
eHelply SDK - 1.1.116

eHelply SDK for SuperStack Services

API version: 1.1.116
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
	"time"
)

// UserResponse When retrieving a user
type UserResponse struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Country *string `json:"country,omitempty"`
	Picture *string `json:"picture,omitempty"`
	GpsLocation map[string]interface{} `json:"gps_location,omitempty"`
	VerifiedLegalTerms *bool `json:"verified_legal_terms,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	DateDeleted *time.Time `json:"date_deleted,omitempty"`
	Email *UserEmail `json:"email,omitempty"`
	Missing []string `json:"missing,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Participants []map[string]interface{} `json:"participants,omitempty"`
	Flags *UserFlags `json:"flags,omitempty"`
	LastLogin *string `json:"last_login,omitempty"`
}

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse() *UserResponse {
	this := UserResponse{}
	var verifiedLegalTerms bool = false
	this.VerifiedLegalTerms = &verifiedLegalTerms
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	var verifiedLegalTerms bool = false
	this.VerifiedLegalTerms = &verifiedLegalTerms
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserResponse) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserResponse) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserResponse) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserResponse) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserResponse) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserResponse) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserResponse) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UserResponse) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UserResponse) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UserResponse) SetCountry(v string) {
	o.Country = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *UserResponse) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *UserResponse) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *UserResponse) SetPicture(v string) {
	o.Picture = &v
}

// GetGpsLocation returns the GpsLocation field value if set, zero value otherwise.
func (o *UserResponse) GetGpsLocation() map[string]interface{} {
	if o == nil || o.GpsLocation == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.GpsLocation
}

// GetGpsLocationOk returns a tuple with the GpsLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetGpsLocationOk() (map[string]interface{}, bool) {
	if o == nil || o.GpsLocation == nil {
		return nil, false
	}
	return o.GpsLocation, true
}

// HasGpsLocation returns a boolean if a field has been set.
func (o *UserResponse) HasGpsLocation() bool {
	if o != nil && o.GpsLocation != nil {
		return true
	}

	return false
}

// SetGpsLocation gets a reference to the given map[string]interface{} and assigns it to the GpsLocation field.
func (o *UserResponse) SetGpsLocation(v map[string]interface{}) {
	o.GpsLocation = v
}

// GetVerifiedLegalTerms returns the VerifiedLegalTerms field value if set, zero value otherwise.
func (o *UserResponse) GetVerifiedLegalTerms() bool {
	if o == nil || o.VerifiedLegalTerms == nil {
		var ret bool
		return ret
	}
	return *o.VerifiedLegalTerms
}

// GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetVerifiedLegalTermsOk() (*bool, bool) {
	if o == nil || o.VerifiedLegalTerms == nil {
		return nil, false
	}
	return o.VerifiedLegalTerms, true
}

// HasVerifiedLegalTerms returns a boolean if a field has been set.
func (o *UserResponse) HasVerifiedLegalTerms() bool {
	if o != nil && o.VerifiedLegalTerms != nil {
		return true
	}

	return false
}

// SetVerifiedLegalTerms gets a reference to the given bool and assigns it to the VerifiedLegalTerms field.
func (o *UserResponse) SetVerifiedLegalTerms(v bool) {
	o.VerifiedLegalTerms = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *UserResponse) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *UserResponse) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *UserResponse) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *UserResponse) GetDateUpdated() time.Time {
	if o == nil || o.DateUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDateUpdatedOk() (*time.Time, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *UserResponse) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given time.Time and assigns it to the DateUpdated field.
func (o *UserResponse) SetDateUpdated(v time.Time) {
	o.DateUpdated = &v
}

// GetDateDeleted returns the DateDeleted field value if set, zero value otherwise.
func (o *UserResponse) GetDateDeleted() time.Time {
	if o == nil || o.DateDeleted == nil {
		var ret time.Time
		return ret
	}
	return *o.DateDeleted
}

// GetDateDeletedOk returns a tuple with the DateDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDateDeletedOk() (*time.Time, bool) {
	if o == nil || o.DateDeleted == nil {
		return nil, false
	}
	return o.DateDeleted, true
}

// HasDateDeleted returns a boolean if a field has been set.
func (o *UserResponse) HasDateDeleted() bool {
	if o != nil && o.DateDeleted != nil {
		return true
	}

	return false
}

// SetDateDeleted gets a reference to the given time.Time and assigns it to the DateDeleted field.
func (o *UserResponse) SetDateDeleted(v time.Time) {
	o.DateDeleted = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserResponse) GetEmail() UserEmail {
	if o == nil || o.Email == nil {
		var ret UserEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailOk() (*UserEmail, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserResponse) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given UserEmail and assigns it to the Email field.
func (o *UserResponse) SetEmail(v UserEmail) {
	o.Email = &v
}

// GetMissing returns the Missing field value if set, zero value otherwise.
func (o *UserResponse) GetMissing() []string {
	if o == nil || o.Missing == nil {
		var ret []string
		return ret
	}
	return o.Missing
}

// GetMissingOk returns a tuple with the Missing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetMissingOk() ([]string, bool) {
	if o == nil || o.Missing == nil {
		return nil, false
	}
	return o.Missing, true
}

// HasMissing returns a boolean if a field has been set.
func (o *UserResponse) HasMissing() bool {
	if o != nil && o.Missing != nil {
		return true
	}

	return false
}

// SetMissing gets a reference to the given []string and assigns it to the Missing field.
func (o *UserResponse) SetMissing(v []string) {
	o.Missing = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UserResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UserResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UserResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *UserResponse) GetParticipants() []map[string]interface{} {
	if o == nil || o.Participants == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetParticipantsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Participants == nil {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *UserResponse) HasParticipants() bool {
	if o != nil && o.Participants != nil {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []map[string]interface{} and assigns it to the Participants field.
func (o *UserResponse) SetParticipants(v []map[string]interface{}) {
	o.Participants = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *UserResponse) GetFlags() UserFlags {
	if o == nil || o.Flags == nil {
		var ret UserFlags
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFlagsOk() (*UserFlags, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *UserResponse) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given UserFlags and assigns it to the Flags field.
func (o *UserResponse) SetFlags(v UserFlags) {
	o.Flags = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *UserResponse) GetLastLogin() string {
	if o == nil || o.LastLogin == nil {
		var ret string
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLastLoginOk() (*string, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *UserResponse) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given string and assigns it to the LastLogin field.
func (o *UserResponse) SetLastLogin(v string) {
	o.LastLogin = &v
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
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
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Missing != nil {
		toSerialize["missing"] = o.Missing
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Participants != nil {
		toSerialize["participants"] = o.Participants
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


