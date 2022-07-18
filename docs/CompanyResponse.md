# CompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] [default to true]
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Contact** | Pointer to [**ContactBase**](ContactBase.md) |  | [optional] 
**Uuid** | **string** |  | 
**ProjectUuid** | Pointer to **string** |  | [optional] 
**MetaUuid** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagBase**](TagBase.md) |  | [optional] 
**Categories** | Pointer to [**[]CategoryBase**](CategoryBase.md) |  | [optional] 
**Places** | Pointer to [**[]PlaceBase**](PlaceBase.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCompanyResponse

`func NewCompanyResponse(uuid string, ) *CompanyResponse`

NewCompanyResponse instantiates a new CompanyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyResponseWithDefaults

`func NewCompanyResponseWithDefaults() *CompanyResponse`

NewCompanyResponseWithDefaults instantiates a new CompanyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompanyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *CompanyResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CompanyResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CompanyResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CompanyResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPublic

`func (o *CompanyResponse) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CompanyResponse) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CompanyResponse) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CompanyResponse) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetMeta

`func (o *CompanyResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CompanyResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CompanyResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CompanyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetContact

`func (o *CompanyResponse) GetContact() ContactBase`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CompanyResponse) GetContactOk() (*ContactBase, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CompanyResponse) SetContact(v ContactBase)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CompanyResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetUuid

`func (o *CompanyResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CompanyResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CompanyResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetProjectUuid

`func (o *CompanyResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *CompanyResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *CompanyResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *CompanyResponse) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetMetaUuid

`func (o *CompanyResponse) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *CompanyResponse) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *CompanyResponse) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *CompanyResponse) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetTags

`func (o *CompanyResponse) GetTags() []TagBase`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CompanyResponse) GetTagsOk() (*[]TagBase, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CompanyResponse) SetTags(v []TagBase)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CompanyResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCategories

`func (o *CompanyResponse) GetCategories() []CategoryBase`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CompanyResponse) GetCategoriesOk() (*[]CategoryBase, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CompanyResponse) SetCategories(v []CategoryBase)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CompanyResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPlaces

`func (o *CompanyResponse) GetPlaces() []PlaceBase`

GetPlaces returns the Places field if non-nil, zero value otherwise.

### GetPlacesOk

`func (o *CompanyResponse) GetPlacesOk() (*[]PlaceBase, bool)`

GetPlacesOk returns a tuple with the Places field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaces

`func (o *CompanyResponse) SetPlaces(v []PlaceBase)`

SetPlaces sets Places field to given value.

### HasPlaces

`func (o *CompanyResponse) HasPlaces() bool`

HasPlaces returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CompanyResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompanyResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompanyResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CompanyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CompanyResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CompanyResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CompanyResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CompanyResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CompanyResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CompanyResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


