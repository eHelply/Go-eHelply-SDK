/*
eHelply SDK - 1.1.120

eHelply SDK for SuperStack Services

API version: 1.1.120
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
	"time"
)

// ParticipantUserReturn Contains all fields required when doing a Participant GET but also has user fields (name, location, ect). This is what is returned from all participant endpoints.
type ParticipantUserReturn struct {
	Uuid *string `json:"uuid,omitempty"`
	UserUuid *string `json:"user_uuid,omitempty"`
	ParticipantMeta map[string]interface{} `json:"participant_meta,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Email *Email `json:"email,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Country *string `json:"country,omitempty"`
	GpsLocation map[string]interface{} `json:"gps_location,omitempty"`
	Picture *string `json:"picture,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	VerifiedLegalTerms *bool `json:"verified_legal_terms,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}

// NewParticipantUserReturn instantiates a new ParticipantUserReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipantUserReturn() *ParticipantUserReturn {
	this := ParticipantUserReturn{}
	return &this
}

// NewParticipantUserReturnWithDefaults instantiates a new ParticipantUserReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantUserReturnWithDefaults() *ParticipantUserReturn {
	this := ParticipantUserReturn{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ParticipantUserReturn) SetUuid(v string) {
	o.Uuid = &v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetUserUuid() string {
	if o == nil || o.UserUuid == nil {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetUserUuidOk() (*string, bool) {
	if o == nil || o.UserUuid == nil {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasUserUuid() bool {
	if o != nil && o.UserUuid != nil {
		return true
	}

	return false
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *ParticipantUserReturn) SetUserUuid(v string) {
	o.UserUuid = &v
}

// GetParticipantMeta returns the ParticipantMeta field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetParticipantMeta() map[string]interface{} {
	if o == nil || o.ParticipantMeta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ParticipantMeta
}

// GetParticipantMetaOk returns a tuple with the ParticipantMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetParticipantMetaOk() (map[string]interface{}, bool) {
	if o == nil || o.ParticipantMeta == nil {
		return nil, false
	}
	return o.ParticipantMeta, true
}

// HasParticipantMeta returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasParticipantMeta() bool {
	if o != nil && o.ParticipantMeta != nil {
		return true
	}

	return false
}

// SetParticipantMeta gets a reference to the given map[string]interface{} and assigns it to the ParticipantMeta field.
func (o *ParticipantUserReturn) SetParticipantMeta(v map[string]interface{}) {
	o.ParticipantMeta = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ParticipantUserReturn) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ParticipantUserReturn) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetEmail() Email {
	if o == nil || o.Email == nil {
		var ret Email
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetEmailOk() (*Email, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given Email and assigns it to the Email field.
func (o *ParticipantUserReturn) SetEmail(v Email) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ParticipantUserReturn) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ParticipantUserReturn) SetCountry(v string) {
	o.Country = &v
}

// GetGpsLocation returns the GpsLocation field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetGpsLocation() map[string]interface{} {
	if o == nil || o.GpsLocation == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.GpsLocation
}

// GetGpsLocationOk returns a tuple with the GpsLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetGpsLocationOk() (map[string]interface{}, bool) {
	if o == nil || o.GpsLocation == nil {
		return nil, false
	}
	return o.GpsLocation, true
}

// HasGpsLocation returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasGpsLocation() bool {
	if o != nil && o.GpsLocation != nil {
		return true
	}

	return false
}

// SetGpsLocation gets a reference to the given map[string]interface{} and assigns it to the GpsLocation field.
func (o *ParticipantUserReturn) SetGpsLocation(v map[string]interface{}) {
	o.GpsLocation = v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *ParticipantUserReturn) SetPicture(v string) {
	o.Picture = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetLastLogin() time.Time {
	if o == nil || o.LastLogin == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || o.LastLogin == nil {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasLastLogin() bool {
	if o != nil && o.LastLogin != nil {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *ParticipantUserReturn) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetVerifiedLegalTerms returns the VerifiedLegalTerms field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetVerifiedLegalTerms() bool {
	if o == nil || o.VerifiedLegalTerms == nil {
		var ret bool
		return ret
	}
	return *o.VerifiedLegalTerms
}

// GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetVerifiedLegalTermsOk() (*bool, bool) {
	if o == nil || o.VerifiedLegalTerms == nil {
		return nil, false
	}
	return o.VerifiedLegalTerms, true
}

// HasVerifiedLegalTerms returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasVerifiedLegalTerms() bool {
	if o != nil && o.VerifiedLegalTerms != nil {
		return true
	}

	return false
}

// SetVerifiedLegalTerms gets a reference to the given bool and assigns it to the VerifiedLegalTerms field.
func (o *ParticipantUserReturn) SetVerifiedLegalTerms(v bool) {
	o.VerifiedLegalTerms = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ParticipantUserReturn) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *ParticipantUserReturn) GetDateUpdated() time.Time {
	if o == nil || o.DateUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParticipantUserReturn) GetDateUpdatedOk() (*time.Time, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ParticipantUserReturn) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given time.Time and assigns it to the DateUpdated field.
func (o *ParticipantUserReturn) SetDateUpdated(v time.Time) {
	o.DateUpdated = &v
}

func (o ParticipantUserReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.UserUuid != nil {
		toSerialize["user_uuid"] = o.UserUuid
	}
	if o.ParticipantMeta != nil {
		toSerialize["participant_meta"] = o.ParticipantMeta
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
	if o.GpsLocation != nil {
		toSerialize["gps_location"] = o.GpsLocation
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.LastLogin != nil {
		toSerialize["last_login"] = o.LastLogin
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
	return json.Marshal(toSerialize)
}

type NullableParticipantUserReturn struct {
	value *ParticipantUserReturn
	isSet bool
}

func (v NullableParticipantUserReturn) Get() *ParticipantUserReturn {
	return v.value
}

func (v *NullableParticipantUserReturn) Set(val *ParticipantUserReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipantUserReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipantUserReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipantUserReturn(val *ParticipantUserReturn) *NullableParticipantUserReturn {
	return &NullableParticipantUserReturn{value: val, isSet: true}
}

func (v NullableParticipantUserReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipantUserReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


