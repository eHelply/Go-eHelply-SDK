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

// UserSignup User information used for user signup
type UserSignup struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Country string `json:"country"`
	Lat *string `json:"lat,omitempty"`
	Lng *string `json:"lng,omitempty"`
	VerifiedLegalTerms *bool `json:"verified_legal_terms,omitempty"`
	Picture *string `json:"picture,omitempty"`
	Newsletters *bool `json:"newsletters,omitempty"`
}

// NewUserSignup instantiates a new UserSignup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSignup(username string, password string, email string, firstName string, lastName string, phoneNumber string, country string) *UserSignup {
	this := UserSignup{}
	this.Username = username
	this.Password = password
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.PhoneNumber = phoneNumber
	this.Country = country
	var verifiedLegalTerms bool = false
	this.VerifiedLegalTerms = &verifiedLegalTerms
	return &this
}

// NewUserSignupWithDefaults instantiates a new UserSignup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSignupWithDefaults() *UserSignup {
	this := UserSignup{}
	var verifiedLegalTerms bool = false
	this.VerifiedLegalTerms = &verifiedLegalTerms
	return &this
}

// GetUsername returns the Username field value
func (o *UserSignup) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserSignup) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserSignup) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *UserSignup) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserSignup) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserSignup) SetPassword(v string) {
	o.Password = v
}

// GetEmail returns the Email field value
func (o *UserSignup) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserSignup) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserSignup) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *UserSignup) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UserSignup) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UserSignup) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UserSignup) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UserSignup) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UserSignup) SetLastName(v string) {
	o.LastName = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *UserSignup) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *UserSignup) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *UserSignup) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetCountry returns the Country field value
func (o *UserSignup) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *UserSignup) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *UserSignup) SetCountry(v string) {
	o.Country = v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *UserSignup) GetLat() string {
	if o == nil || o.Lat == nil {
		var ret string
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignup) GetLatOk() (*string, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *UserSignup) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given string and assigns it to the Lat field.
func (o *UserSignup) SetLat(v string) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *UserSignup) GetLng() string {
	if o == nil || o.Lng == nil {
		var ret string
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignup) GetLngOk() (*string, bool) {
	if o == nil || o.Lng == nil {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *UserSignup) HasLng() bool {
	if o != nil && o.Lng != nil {
		return true
	}

	return false
}

// SetLng gets a reference to the given string and assigns it to the Lng field.
func (o *UserSignup) SetLng(v string) {
	o.Lng = &v
}

// GetVerifiedLegalTerms returns the VerifiedLegalTerms field value if set, zero value otherwise.
func (o *UserSignup) GetVerifiedLegalTerms() bool {
	if o == nil || o.VerifiedLegalTerms == nil {
		var ret bool
		return ret
	}
	return *o.VerifiedLegalTerms
}

// GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignup) GetVerifiedLegalTermsOk() (*bool, bool) {
	if o == nil || o.VerifiedLegalTerms == nil {
		return nil, false
	}
	return o.VerifiedLegalTerms, true
}

// HasVerifiedLegalTerms returns a boolean if a field has been set.
func (o *UserSignup) HasVerifiedLegalTerms() bool {
	if o != nil && o.VerifiedLegalTerms != nil {
		return true
	}

	return false
}

// SetVerifiedLegalTerms gets a reference to the given bool and assigns it to the VerifiedLegalTerms field.
func (o *UserSignup) SetVerifiedLegalTerms(v bool) {
	o.VerifiedLegalTerms = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *UserSignup) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignup) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *UserSignup) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *UserSignup) SetPicture(v string) {
	o.Picture = &v
}

// GetNewsletters returns the Newsletters field value if set, zero value otherwise.
func (o *UserSignup) GetNewsletters() bool {
	if o == nil || o.Newsletters == nil {
		var ret bool
		return ret
	}
	return *o.Newsletters
}

// GetNewslettersOk returns a tuple with the Newsletters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignup) GetNewslettersOk() (*bool, bool) {
	if o == nil || o.Newsletters == nil {
		return nil, false
	}
	return o.Newsletters, true
}

// HasNewsletters returns a boolean if a field has been set.
func (o *UserSignup) HasNewsletters() bool {
	if o != nil && o.Newsletters != nil {
		return true
	}

	return false
}

// SetNewsletters gets a reference to the given bool and assigns it to the Newsletters field.
func (o *UserSignup) SetNewsletters(v bool) {
	o.Newsletters = &v
}

func (o UserSignup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	if o.VerifiedLegalTerms != nil {
		toSerialize["verified_legal_terms"] = o.VerifiedLegalTerms
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.Newsletters != nil {
		toSerialize["newsletters"] = o.Newsletters
	}
	return json.Marshal(toSerialize)
}

type NullableUserSignup struct {
	value *UserSignup
	isSet bool
}

func (v NullableUserSignup) Get() *UserSignup {
	return v.value
}

func (v *NullableUserSignup) Set(val *UserSignup) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSignup) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSignup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSignup(val *UserSignup) *NullableUserSignup {
	return &NullableUserSignup{value: val, isSet: true}
}

func (v NullableUserSignup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSignup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


