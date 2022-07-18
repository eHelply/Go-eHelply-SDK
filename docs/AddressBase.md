# AddressBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**PostalZipCode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**ProvinceState** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Lat** | Pointer to **string** |  | [optional] 
**Lng** | Pointer to **string** |  | [optional] 

## Methods

### NewAddressBase

`func NewAddressBase() *AddressBase`

NewAddressBase instantiates a new AddressBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBaseWithDefaults

`func NewAddressBaseWithDefaults() *AddressBase`

NewAddressBaseWithDefaults instantiates a new AddressBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *AddressBase) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AddressBase) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AddressBase) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *AddressBase) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetAddressLine1

`func (o *AddressBase) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressBase) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressBase) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *AddressBase) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AddressBase) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressBase) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressBase) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressBase) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetPostalZipCode

`func (o *AddressBase) GetPostalZipCode() string`

GetPostalZipCode returns the PostalZipCode field if non-nil, zero value otherwise.

### GetPostalZipCodeOk

`func (o *AddressBase) GetPostalZipCodeOk() (*string, bool)`

GetPostalZipCodeOk returns a tuple with the PostalZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalZipCode

`func (o *AddressBase) SetPostalZipCode(v string)`

SetPostalZipCode sets PostalZipCode field to given value.

### HasPostalZipCode

`func (o *AddressBase) HasPostalZipCode() bool`

HasPostalZipCode returns a boolean if a field has been set.

### GetCity

`func (o *AddressBase) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressBase) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressBase) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressBase) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetProvinceState

`func (o *AddressBase) GetProvinceState() string`

GetProvinceState returns the ProvinceState field if non-nil, zero value otherwise.

### GetProvinceStateOk

`func (o *AddressBase) GetProvinceStateOk() (*string, bool)`

GetProvinceStateOk returns a tuple with the ProvinceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceState

`func (o *AddressBase) SetProvinceState(v string)`

SetProvinceState sets ProvinceState field to given value.

### HasProvinceState

`func (o *AddressBase) HasProvinceState() bool`

HasProvinceState returns a boolean if a field has been set.

### GetCountry

`func (o *AddressBase) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressBase) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressBase) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AddressBase) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLat

`func (o *AddressBase) GetLat() string`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *AddressBase) GetLatOk() (*string, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *AddressBase) SetLat(v string)`

SetLat sets Lat field to given value.

### HasLat

`func (o *AddressBase) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *AddressBase) GetLng() string`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *AddressBase) GetLngOk() (*string, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *AddressBase) SetLng(v string)`

SetLng sets Lng field to given value.

### HasLng

`func (o *AddressBase) HasLng() bool`

HasLng returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


