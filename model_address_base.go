/*
eHelply SDK - 1.1.91

eHelply SDK for SuperStack Services

API version: 1.1.91
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AddressBase **:param** address_type                        **type:** string or None  **:param** address_line_1                      **type:** string or None  **:param** address_line_2                      **type:** string or None  **:param** postal_zip_code                     **type:** string or None  **:param** city                                **type:** string or None  **:param** province_state                      **type:** string or None  **:param** country                             **type:** string or None  **:param** lat                                 **type:** string or None  **:param** lng                                 **type:** string or None
type AddressBase struct {
	AddressType *string `json:"address_type,omitempty"`
	AddressLine1 *string `json:"address_line_1,omitempty"`
	AddressLine2 *string `json:"address_line_2,omitempty"`
	PostalZipCode *string `json:"postal_zip_code,omitempty"`
	City *string `json:"city,omitempty"`
	ProvinceState *string `json:"province_state,omitempty"`
	Country *string `json:"country,omitempty"`
	Lat *string `json:"lat,omitempty"`
	Lng *string `json:"lng,omitempty"`
}

// NewAddressBase instantiates a new AddressBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressBase() *AddressBase {
	this := AddressBase{}
	return &this
}

// NewAddressBaseWithDefaults instantiates a new AddressBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressBaseWithDefaults() *AddressBase {
	this := AddressBase{}
	return &this
}

// GetAddressType returns the AddressType field value if set, zero value otherwise.
func (o *AddressBase) GetAddressType() string {
	if o == nil || o.AddressType == nil {
		var ret string
		return ret
	}
	return *o.AddressType
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetAddressTypeOk() (*string, bool) {
	if o == nil || o.AddressType == nil {
		return nil, false
	}
	return o.AddressType, true
}

// HasAddressType returns a boolean if a field has been set.
func (o *AddressBase) HasAddressType() bool {
	if o != nil && o.AddressType != nil {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given string and assigns it to the AddressType field.
func (o *AddressBase) SetAddressType(v string) {
	o.AddressType = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *AddressBase) GetAddressLine1() string {
	if o == nil || o.AddressLine1 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetAddressLine1Ok() (*string, bool) {
	if o == nil || o.AddressLine1 == nil {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *AddressBase) HasAddressLine1() bool {
	if o != nil && o.AddressLine1 != nil {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *AddressBase) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *AddressBase) GetAddressLine2() string {
	if o == nil || o.AddressLine2 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetAddressLine2Ok() (*string, bool) {
	if o == nil || o.AddressLine2 == nil {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AddressBase) HasAddressLine2() bool {
	if o != nil && o.AddressLine2 != nil {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *AddressBase) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetPostalZipCode returns the PostalZipCode field value if set, zero value otherwise.
func (o *AddressBase) GetPostalZipCode() string {
	if o == nil || o.PostalZipCode == nil {
		var ret string
		return ret
	}
	return *o.PostalZipCode
}

// GetPostalZipCodeOk returns a tuple with the PostalZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetPostalZipCodeOk() (*string, bool) {
	if o == nil || o.PostalZipCode == nil {
		return nil, false
	}
	return o.PostalZipCode, true
}

// HasPostalZipCode returns a boolean if a field has been set.
func (o *AddressBase) HasPostalZipCode() bool {
	if o != nil && o.PostalZipCode != nil {
		return true
	}

	return false
}

// SetPostalZipCode gets a reference to the given string and assigns it to the PostalZipCode field.
func (o *AddressBase) SetPostalZipCode(v string) {
	o.PostalZipCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressBase) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressBase) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressBase) SetCity(v string) {
	o.City = &v
}

// GetProvinceState returns the ProvinceState field value if set, zero value otherwise.
func (o *AddressBase) GetProvinceState() string {
	if o == nil || o.ProvinceState == nil {
		var ret string
		return ret
	}
	return *o.ProvinceState
}

// GetProvinceStateOk returns a tuple with the ProvinceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetProvinceStateOk() (*string, bool) {
	if o == nil || o.ProvinceState == nil {
		return nil, false
	}
	return o.ProvinceState, true
}

// HasProvinceState returns a boolean if a field has been set.
func (o *AddressBase) HasProvinceState() bool {
	if o != nil && o.ProvinceState != nil {
		return true
	}

	return false
}

// SetProvinceState gets a reference to the given string and assigns it to the ProvinceState field.
func (o *AddressBase) SetProvinceState(v string) {
	o.ProvinceState = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressBase) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressBase) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AddressBase) SetCountry(v string) {
	o.Country = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *AddressBase) GetLat() string {
	if o == nil || o.Lat == nil {
		var ret string
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetLatOk() (*string, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *AddressBase) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given string and assigns it to the Lat field.
func (o *AddressBase) SetLat(v string) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *AddressBase) GetLng() string {
	if o == nil || o.Lng == nil {
		var ret string
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBase) GetLngOk() (*string, bool) {
	if o == nil || o.Lng == nil {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *AddressBase) HasLng() bool {
	if o != nil && o.Lng != nil {
		return true
	}

	return false
}

// SetLng gets a reference to the given string and assigns it to the Lng field.
func (o *AddressBase) SetLng(v string) {
	o.Lng = &v
}

func (o AddressBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressType != nil {
		toSerialize["address_type"] = o.AddressType
	}
	if o.AddressLine1 != nil {
		toSerialize["address_line_1"] = o.AddressLine1
	}
	if o.AddressLine2 != nil {
		toSerialize["address_line_2"] = o.AddressLine2
	}
	if o.PostalZipCode != nil {
		toSerialize["postal_zip_code"] = o.PostalZipCode
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.ProvinceState != nil {
		toSerialize["province_state"] = o.ProvinceState
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	return json.Marshal(toSerialize)
}

type NullableAddressBase struct {
	value *AddressBase
	isSet bool
}

func (v NullableAddressBase) Get() *AddressBase {
	return v.value
}

func (v *NullableAddressBase) Set(val *AddressBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBase(val *AddressBase) *NullableAddressBase {
	return &NullableAddressBase{value: val, isSet: true}
}

func (v NullableAddressBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


