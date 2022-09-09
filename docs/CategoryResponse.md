# CategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Name** | **string** |  | 
**ProjectUuid** | Pointer to **string** |  | [optional] 
**MetaUuid** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCategoryResponse

`func NewCategoryResponse(uuid string, name string, ) *CategoryResponse`

NewCategoryResponse instantiates a new CategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryResponseWithDefaults

`func NewCategoryResponseWithDefaults() *CategoryResponse`

NewCategoryResponseWithDefaults instantiates a new CategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *CategoryResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CategoryResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CategoryResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *CategoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProjectUuid

`func (o *CategoryResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *CategoryResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *CategoryResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *CategoryResponse) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetMetaUuid

`func (o *CategoryResponse) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *CategoryResponse) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *CategoryResponse) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *CategoryResponse) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetMeta

`func (o *CategoryResponse) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CategoryResponse) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CategoryResponse) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CategoryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


