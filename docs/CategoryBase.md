# CategoryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MetaUuid** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**ProjectUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewCategoryBase

`func NewCategoryBase(name string, ) *CategoryBase`

NewCategoryBase instantiates a new CategoryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryBaseWithDefaults

`func NewCategoryBaseWithDefaults() *CategoryBase`

NewCategoryBaseWithDefaults instantiates a new CategoryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategoryBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryBase) SetName(v string)`

SetName sets Name field to given value.


### GetMetaUuid

`func (o *CategoryBase) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *CategoryBase) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *CategoryBase) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *CategoryBase) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetMeta

`func (o *CategoryBase) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CategoryBase) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CategoryBase) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CategoryBase) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetProjectUuid

`func (o *CategoryBase) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *CategoryBase) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *CategoryBase) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *CategoryBase) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


