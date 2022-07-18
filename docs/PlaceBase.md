# PlaceBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] [default to true]
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Addresses** | Pointer to [**[]AddressBase**](AddressBase.md) |  | [optional] 
**Contact** | Pointer to [**ContactBase**](ContactBase.md) |  | [optional] 

## Methods

### NewPlaceBase

`func NewPlaceBase(name string, ) *PlaceBase`

NewPlaceBase instantiates a new PlaceBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceBaseWithDefaults

`func NewPlaceBaseWithDefaults() *PlaceBase`

NewPlaceBaseWithDefaults instantiates a new PlaceBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlaceBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaceBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaceBase) SetName(v string)`

SetName sets Name field to given value.


### GetSummary

`func (o *PlaceBase) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PlaceBase) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PlaceBase) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PlaceBase) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPublic

`func (o *PlaceBase) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PlaceBase) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PlaceBase) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *PlaceBase) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetMeta

`func (o *PlaceBase) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PlaceBase) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PlaceBase) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PlaceBase) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAddresses

`func (o *PlaceBase) GetAddresses() []AddressBase`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PlaceBase) GetAddressesOk() (*[]AddressBase, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PlaceBase) SetAddresses(v []AddressBase)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PlaceBase) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetContact

`func (o *PlaceBase) GetContact() ContactBase`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PlaceBase) GetContactOk() (*ContactBase, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PlaceBase) SetContact(v ContactBase)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PlaceBase) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


