# CatalogReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Featured** | Pointer to **map[string]interface{}** |  | [optional] 
**SubCatalogs** | Pointer to **map[string]interface{}** |  | [optional] 
**Uuid** | **string** |  | 
**MetaUuid** | Pointer to **string** |  | [optional] 
**ProjectUuid** | **string** |  | 
**ProductUuids** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCatalogReturn

`func NewCatalogReturn(uuid string, projectUuid string, ) *CatalogReturn`

NewCatalogReturn instantiates a new CatalogReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogReturnWithDefaults

`func NewCatalogReturnWithDefaults() *CatalogReturn`

NewCatalogReturnWithDefaults instantiates a new CatalogReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *CatalogReturn) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *CatalogReturn) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *CatalogReturn) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *CatalogReturn) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetName

`func (o *CatalogReturn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogReturn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogReturn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogReturn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFeatured

`func (o *CatalogReturn) GetFeatured() map[string]interface{}`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CatalogReturn) GetFeaturedOk() (*map[string]interface{}, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CatalogReturn) SetFeatured(v map[string]interface{})`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CatalogReturn) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetSubCatalogs

`func (o *CatalogReturn) GetSubCatalogs() map[string]interface{}`

GetSubCatalogs returns the SubCatalogs field if non-nil, zero value otherwise.

### GetSubCatalogsOk

`func (o *CatalogReturn) GetSubCatalogsOk() (*map[string]interface{}, bool)`

GetSubCatalogsOk returns a tuple with the SubCatalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCatalogs

`func (o *CatalogReturn) SetSubCatalogs(v map[string]interface{})`

SetSubCatalogs sets SubCatalogs field to given value.

### HasSubCatalogs

`func (o *CatalogReturn) HasSubCatalogs() bool`

HasSubCatalogs returns a boolean if a field has been set.

### GetUuid

`func (o *CatalogReturn) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CatalogReturn) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CatalogReturn) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetMetaUuid

`func (o *CatalogReturn) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *CatalogReturn) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *CatalogReturn) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *CatalogReturn) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetProjectUuid

`func (o *CatalogReturn) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *CatalogReturn) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *CatalogReturn) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.


### GetProductUuids

`func (o *CatalogReturn) GetProductUuids() []string`

GetProductUuids returns the ProductUuids field if non-nil, zero value otherwise.

### GetProductUuidsOk

`func (o *CatalogReturn) GetProductUuidsOk() (*[]string, bool)`

GetProductUuidsOk returns a tuple with the ProductUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUuids

`func (o *CatalogReturn) SetProductUuids(v []string)`

SetProductUuids sets ProductUuids field to given value.

### HasProductUuids

`func (o *CatalogReturn) HasProductUuids() bool`

HasProductUuids returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CatalogReturn) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CatalogReturn) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CatalogReturn) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CatalogReturn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CatalogReturn) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CatalogReturn) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CatalogReturn) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CatalogReturn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CatalogReturn) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CatalogReturn) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CatalogReturn) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CatalogReturn) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


