# PlaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Summary** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] [default to true]
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Addresses** | Pointer to [**[]AddressBase**](AddressBase.md) |  | [optional] 
**Contact** | Pointer to [**ContactBase**](ContactBase.md) |  | [optional] 
**Uuid** | **string** |  | 
**ProjectUuid** | Pointer to **string** |  | [optional] 
**MetaUuid** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagBase**](TagBase.md) |  | [optional] 
**Categories** | Pointer to [**[]CategoryBase**](CategoryBase.md) |  | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewPlaceResponse

`func NewPlaceResponse(name string, uuid string, ) *PlaceResponse`

NewPlaceResponse instantiates a new PlaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceResponseWithDefaults

`func NewPlaceResponseWithDefaults() *PlaceResponse`

NewPlaceResponseWithDefaults instantiates a new PlaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlaceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSummary

`func (o *PlaceResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PlaceResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PlaceResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PlaceResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPublic

`func (o *PlaceResponse) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PlaceResponse) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PlaceResponse) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *PlaceResponse) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetMeta

`func (o *PlaceResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PlaceResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PlaceResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PlaceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAddresses

`func (o *PlaceResponse) GetAddresses() []AddressBase`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PlaceResponse) GetAddressesOk() (*[]AddressBase, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PlaceResponse) SetAddresses(v []AddressBase)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PlaceResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetContact

`func (o *PlaceResponse) GetContact() ContactBase`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PlaceResponse) GetContactOk() (*ContactBase, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PlaceResponse) SetContact(v ContactBase)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PlaceResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetUuid

`func (o *PlaceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PlaceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PlaceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetProjectUuid

`func (o *PlaceResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *PlaceResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *PlaceResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *PlaceResponse) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetMetaUuid

`func (o *PlaceResponse) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *PlaceResponse) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *PlaceResponse) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *PlaceResponse) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetTags

`func (o *PlaceResponse) GetTags() []TagBase`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PlaceResponse) GetTagsOk() (*[]TagBase, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PlaceResponse) SetTags(v []TagBase)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PlaceResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCategories

`func (o *PlaceResponse) GetCategories() []CategoryBase`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PlaceResponse) GetCategoriesOk() (*[]CategoryBase, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PlaceResponse) SetCategories(v []CategoryBase)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PlaceResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCompany

`func (o *PlaceResponse) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PlaceResponse) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PlaceResponse) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PlaceResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PlaceResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlaceResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlaceResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PlaceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PlaceResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlaceResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlaceResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PlaceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PlaceResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PlaceResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PlaceResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PlaceResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


