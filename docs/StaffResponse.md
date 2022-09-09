# StaffResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityUuid** | **string** |  | 
**ProjectUuid** | Pointer to **string** |  | [optional] 
**ScheduleUuid** | Pointer to **string** |  | [optional] 
**CatalogUuid** | Pointer to **string** |  | [optional] 
**ReviewGroupUuid** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**Entity** | Pointer to **string** |  | [optional] 
**Place** | Pointer to **map[string]interface{}** |  | [optional] 
**PlaceRoles** | Pointer to **[]string** |  | [optional] 
**Company** | Pointer to **map[string]interface{}** |  | [optional] 
**CompanyRoles** | Pointer to **[]string** |  | [optional] 
**Schedule** | Pointer to **map[string]interface{}** |  | [optional] 
**Catalog** | Pointer to **map[string]interface{}** |  | [optional] 
**Reviews** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewStaffResponse

`func NewStaffResponse(entityUuid string, uuid string, ) *StaffResponse`

NewStaffResponse instantiates a new StaffResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaffResponseWithDefaults

`func NewStaffResponseWithDefaults() *StaffResponse`

NewStaffResponseWithDefaults instantiates a new StaffResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityUuid

`func (o *StaffResponse) GetEntityUuid() string`

GetEntityUuid returns the EntityUuid field if non-nil, zero value otherwise.

### GetEntityUuidOk

`func (o *StaffResponse) GetEntityUuidOk() (*string, bool)`

GetEntityUuidOk returns a tuple with the EntityUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUuid

`func (o *StaffResponse) SetEntityUuid(v string)`

SetEntityUuid sets EntityUuid field to given value.


### GetProjectUuid

`func (o *StaffResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *StaffResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *StaffResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *StaffResponse) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetScheduleUuid

`func (o *StaffResponse) GetScheduleUuid() string`

GetScheduleUuid returns the ScheduleUuid field if non-nil, zero value otherwise.

### GetScheduleUuidOk

`func (o *StaffResponse) GetScheduleUuidOk() (*string, bool)`

GetScheduleUuidOk returns a tuple with the ScheduleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUuid

`func (o *StaffResponse) SetScheduleUuid(v string)`

SetScheduleUuid sets ScheduleUuid field to given value.

### HasScheduleUuid

`func (o *StaffResponse) HasScheduleUuid() bool`

HasScheduleUuid returns a boolean if a field has been set.

### GetCatalogUuid

`func (o *StaffResponse) GetCatalogUuid() string`

GetCatalogUuid returns the CatalogUuid field if non-nil, zero value otherwise.

### GetCatalogUuidOk

`func (o *StaffResponse) GetCatalogUuidOk() (*string, bool)`

GetCatalogUuidOk returns a tuple with the CatalogUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogUuid

`func (o *StaffResponse) SetCatalogUuid(v string)`

SetCatalogUuid sets CatalogUuid field to given value.

### HasCatalogUuid

`func (o *StaffResponse) HasCatalogUuid() bool`

HasCatalogUuid returns a boolean if a field has been set.

### GetReviewGroupUuid

`func (o *StaffResponse) GetReviewGroupUuid() string`

GetReviewGroupUuid returns the ReviewGroupUuid field if non-nil, zero value otherwise.

### GetReviewGroupUuidOk

`func (o *StaffResponse) GetReviewGroupUuidOk() (*string, bool)`

GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewGroupUuid

`func (o *StaffResponse) SetReviewGroupUuid(v string)`

SetReviewGroupUuid sets ReviewGroupUuid field to given value.

### HasReviewGroupUuid

`func (o *StaffResponse) HasReviewGroupUuid() bool`

HasReviewGroupUuid returns a boolean if a field has been set.

### GetUuid

`func (o *StaffResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StaffResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StaffResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetEntity

`func (o *StaffResponse) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *StaffResponse) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *StaffResponse) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *StaffResponse) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetPlace

`func (o *StaffResponse) GetPlace() map[string]interface{}`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *StaffResponse) GetPlaceOk() (*map[string]interface{}, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *StaffResponse) SetPlace(v map[string]interface{})`

SetPlace sets Place field to given value.

### HasPlace

`func (o *StaffResponse) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetPlaceRoles

`func (o *StaffResponse) GetPlaceRoles() []string`

GetPlaceRoles returns the PlaceRoles field if non-nil, zero value otherwise.

### GetPlaceRolesOk

`func (o *StaffResponse) GetPlaceRolesOk() (*[]string, bool)`

GetPlaceRolesOk returns a tuple with the PlaceRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceRoles

`func (o *StaffResponse) SetPlaceRoles(v []string)`

SetPlaceRoles sets PlaceRoles field to given value.

### HasPlaceRoles

`func (o *StaffResponse) HasPlaceRoles() bool`

HasPlaceRoles returns a boolean if a field has been set.

### GetCompany

`func (o *StaffResponse) GetCompany() map[string]interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *StaffResponse) GetCompanyOk() (*map[string]interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *StaffResponse) SetCompany(v map[string]interface{})`

SetCompany sets Company field to given value.

### HasCompany

`func (o *StaffResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyRoles

`func (o *StaffResponse) GetCompanyRoles() []string`

GetCompanyRoles returns the CompanyRoles field if non-nil, zero value otherwise.

### GetCompanyRolesOk

`func (o *StaffResponse) GetCompanyRolesOk() (*[]string, bool)`

GetCompanyRolesOk returns a tuple with the CompanyRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyRoles

`func (o *StaffResponse) SetCompanyRoles(v []string)`

SetCompanyRoles sets CompanyRoles field to given value.

### HasCompanyRoles

`func (o *StaffResponse) HasCompanyRoles() bool`

HasCompanyRoles returns a boolean if a field has been set.

### GetSchedule

`func (o *StaffResponse) GetSchedule() map[string]interface{}`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *StaffResponse) GetScheduleOk() (*map[string]interface{}, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *StaffResponse) SetSchedule(v map[string]interface{})`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *StaffResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCatalog

`func (o *StaffResponse) GetCatalog() map[string]interface{}`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *StaffResponse) GetCatalogOk() (*map[string]interface{}, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *StaffResponse) SetCatalog(v map[string]interface{})`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *StaffResponse) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetReviews

`func (o *StaffResponse) GetReviews() map[string]interface{}`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *StaffResponse) GetReviewsOk() (*map[string]interface{}, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *StaffResponse) SetReviews(v map[string]interface{})`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *StaffResponse) HasReviews() bool`

HasReviews returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StaffResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StaffResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StaffResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StaffResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StaffResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StaffResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StaffResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StaffResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *StaffResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *StaffResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *StaffResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *StaffResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


