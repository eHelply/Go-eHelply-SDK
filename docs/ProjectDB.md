# ProjectDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **string** |  | 
**CurrentSpend** | **int32** |  | 
**MaxSpend** | **int32** |  | 
**ProjectStatus** | **string** |  | 
**ArchivedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectDB

`func NewProjectDB(uuid string, name string, createdAt string, currentSpend int32, maxSpend int32, projectStatus string, ) *ProjectDB`

NewProjectDB instantiates a new ProjectDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDBWithDefaults

`func NewProjectDBWithDefaults() *ProjectDB`

NewProjectDBWithDefaults instantiates a new ProjectDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ProjectDB) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectDB) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectDB) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *ProjectDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDB) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ProjectDB) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectDB) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectDB) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrentSpend

`func (o *ProjectDB) GetCurrentSpend() int32`

GetCurrentSpend returns the CurrentSpend field if non-nil, zero value otherwise.

### GetCurrentSpendOk

`func (o *ProjectDB) GetCurrentSpendOk() (*int32, bool)`

GetCurrentSpendOk returns a tuple with the CurrentSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSpend

`func (o *ProjectDB) SetCurrentSpend(v int32)`

SetCurrentSpend sets CurrentSpend field to given value.


### GetMaxSpend

`func (o *ProjectDB) GetMaxSpend() int32`

GetMaxSpend returns the MaxSpend field if non-nil, zero value otherwise.

### GetMaxSpendOk

`func (o *ProjectDB) GetMaxSpendOk() (*int32, bool)`

GetMaxSpendOk returns a tuple with the MaxSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpend

`func (o *ProjectDB) SetMaxSpend(v int32)`

SetMaxSpend sets MaxSpend field to given value.


### GetProjectStatus

`func (o *ProjectDB) GetProjectStatus() string`

GetProjectStatus returns the ProjectStatus field if non-nil, zero value otherwise.

### GetProjectStatusOk

`func (o *ProjectDB) GetProjectStatusOk() (*string, bool)`

GetProjectStatusOk returns a tuple with the ProjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatus

`func (o *ProjectDB) SetProjectStatus(v string)`

SetProjectStatus sets ProjectStatus field to given value.


### GetArchivedAt

`func (o *ProjectDB) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *ProjectDB) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *ProjectDB) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *ProjectDB) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


