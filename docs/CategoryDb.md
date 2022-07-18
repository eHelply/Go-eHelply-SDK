# CategoryDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Name** | **string** |  | 
**ProjectUuid** | Pointer to **string** |  | [optional] 
**MetaUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewCategoryDb

`func NewCategoryDb(uuid string, name string, ) *CategoryDb`

NewCategoryDb instantiates a new CategoryDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryDbWithDefaults

`func NewCategoryDbWithDefaults() *CategoryDb`

NewCategoryDbWithDefaults instantiates a new CategoryDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *CategoryDb) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CategoryDb) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CategoryDb) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *CategoryDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryDb) SetName(v string)`

SetName sets Name field to given value.


### GetProjectUuid

`func (o *CategoryDb) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *CategoryDb) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *CategoryDb) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *CategoryDb) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetMetaUuid

`func (o *CategoryDb) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *CategoryDb) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *CategoryDb) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *CategoryDb) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


