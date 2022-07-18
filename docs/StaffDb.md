# StaffDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**ProjectUuid** | Pointer to **string** |  | [optional] 
**EntityUuid** | Pointer to **string** |  | [optional] 
**ScheduleUuid** | Pointer to **string** |  | [optional] 
**CatalogUuid** | Pointer to **string** |  | [optional] 
**ReviewGroupUuid** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewStaffDb

`func NewStaffDb(uuid string, ) *StaffDb`

NewStaffDb instantiates a new StaffDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaffDbWithDefaults

`func NewStaffDbWithDefaults() *StaffDb`

NewStaffDbWithDefaults instantiates a new StaffDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *StaffDb) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StaffDb) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StaffDb) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetProjectUuid

`func (o *StaffDb) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *StaffDb) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *StaffDb) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *StaffDb) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetEntityUuid

`func (o *StaffDb) GetEntityUuid() string`

GetEntityUuid returns the EntityUuid field if non-nil, zero value otherwise.

### GetEntityUuidOk

`func (o *StaffDb) GetEntityUuidOk() (*string, bool)`

GetEntityUuidOk returns a tuple with the EntityUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUuid

`func (o *StaffDb) SetEntityUuid(v string)`

SetEntityUuid sets EntityUuid field to given value.

### HasEntityUuid

`func (o *StaffDb) HasEntityUuid() bool`

HasEntityUuid returns a boolean if a field has been set.

### GetScheduleUuid

`func (o *StaffDb) GetScheduleUuid() string`

GetScheduleUuid returns the ScheduleUuid field if non-nil, zero value otherwise.

### GetScheduleUuidOk

`func (o *StaffDb) GetScheduleUuidOk() (*string, bool)`

GetScheduleUuidOk returns a tuple with the ScheduleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUuid

`func (o *StaffDb) SetScheduleUuid(v string)`

SetScheduleUuid sets ScheduleUuid field to given value.

### HasScheduleUuid

`func (o *StaffDb) HasScheduleUuid() bool`

HasScheduleUuid returns a boolean if a field has been set.

### GetCatalogUuid

`func (o *StaffDb) GetCatalogUuid() string`

GetCatalogUuid returns the CatalogUuid field if non-nil, zero value otherwise.

### GetCatalogUuidOk

`func (o *StaffDb) GetCatalogUuidOk() (*string, bool)`

GetCatalogUuidOk returns a tuple with the CatalogUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogUuid

`func (o *StaffDb) SetCatalogUuid(v string)`

SetCatalogUuid sets CatalogUuid field to given value.

### HasCatalogUuid

`func (o *StaffDb) HasCatalogUuid() bool`

HasCatalogUuid returns a boolean if a field has been set.

### GetReviewGroupUuid

`func (o *StaffDb) GetReviewGroupUuid() string`

GetReviewGroupUuid returns the ReviewGroupUuid field if non-nil, zero value otherwise.

### GetReviewGroupUuidOk

`func (o *StaffDb) GetReviewGroupUuidOk() (*string, bool)`

GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewGroupUuid

`func (o *StaffDb) SetReviewGroupUuid(v string)`

SetReviewGroupUuid sets ReviewGroupUuid field to given value.

### HasReviewGroupUuid

`func (o *StaffDb) HasReviewGroupUuid() bool`

HasReviewGroupUuid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StaffDb) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StaffDb) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StaffDb) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StaffDb) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StaffDb) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StaffDb) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StaffDb) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StaffDb) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *StaffDb) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *StaffDb) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *StaffDb) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *StaffDb) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


