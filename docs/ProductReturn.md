# ProductReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | Pointer to **map[string]interface{}** |  | [optional] 
**CollectionUuid** | Pointer to **string** |  | [optional] 
**ReviewGroupUuid** | Pointer to **string** |  | [optional] 
**Addons** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | **int32** |  | 
**QuantityForPublic** | **int32** |  | 
**Uuid** | **string** |  | 
**MetaUuid** | Pointer to **string** |  | [optional] 
**ProjectUuid** | **string** |  | 
**CatalogUuid** | Pointer to **string** |  | [optional] 
**AddonList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewProductReturn

`func NewProductReturn(price int32, quantityForPublic int32, uuid string, projectUuid string, ) *ProductReturn`

NewProductReturn instantiates a new ProductReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductReturnWithDefaults

`func NewProductReturnWithDefaults() *ProductReturn`

NewProductReturnWithDefaults instantiates a new ProductReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *ProductReturn) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ProductReturn) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ProductReturn) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ProductReturn) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetCollectionUuid

`func (o *ProductReturn) GetCollectionUuid() string`

GetCollectionUuid returns the CollectionUuid field if non-nil, zero value otherwise.

### GetCollectionUuidOk

`func (o *ProductReturn) GetCollectionUuidOk() (*string, bool)`

GetCollectionUuidOk returns a tuple with the CollectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionUuid

`func (o *ProductReturn) SetCollectionUuid(v string)`

SetCollectionUuid sets CollectionUuid field to given value.

### HasCollectionUuid

`func (o *ProductReturn) HasCollectionUuid() bool`

HasCollectionUuid returns a boolean if a field has been set.

### GetReviewGroupUuid

`func (o *ProductReturn) GetReviewGroupUuid() string`

GetReviewGroupUuid returns the ReviewGroupUuid field if non-nil, zero value otherwise.

### GetReviewGroupUuidOk

`func (o *ProductReturn) GetReviewGroupUuidOk() (*string, bool)`

GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewGroupUuid

`func (o *ProductReturn) SetReviewGroupUuid(v string)`

SetReviewGroupUuid sets ReviewGroupUuid field to given value.

### HasReviewGroupUuid

`func (o *ProductReturn) HasReviewGroupUuid() bool`

HasReviewGroupUuid returns a boolean if a field has been set.

### GetAddons

`func (o *ProductReturn) GetAddons() []string`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *ProductReturn) GetAddonsOk() (*[]string, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *ProductReturn) SetAddons(v []string)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *ProductReturn) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetName

`func (o *ProductReturn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductReturn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductReturn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductReturn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ProductReturn) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductReturn) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductReturn) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetQuantityForPublic

`func (o *ProductReturn) GetQuantityForPublic() int32`

GetQuantityForPublic returns the QuantityForPublic field if non-nil, zero value otherwise.

### GetQuantityForPublicOk

`func (o *ProductReturn) GetQuantityForPublicOk() (*int32, bool)`

GetQuantityForPublicOk returns a tuple with the QuantityForPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityForPublic

`func (o *ProductReturn) SetQuantityForPublic(v int32)`

SetQuantityForPublic sets QuantityForPublic field to given value.


### GetUuid

`func (o *ProductReturn) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProductReturn) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProductReturn) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetMetaUuid

`func (o *ProductReturn) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *ProductReturn) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *ProductReturn) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *ProductReturn) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetProjectUuid

`func (o *ProductReturn) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *ProductReturn) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *ProductReturn) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.


### GetCatalogUuid

`func (o *ProductReturn) GetCatalogUuid() string`

GetCatalogUuid returns the CatalogUuid field if non-nil, zero value otherwise.

### GetCatalogUuidOk

`func (o *ProductReturn) GetCatalogUuidOk() (*string, bool)`

GetCatalogUuidOk returns a tuple with the CatalogUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogUuid

`func (o *ProductReturn) SetCatalogUuid(v string)`

SetCatalogUuid sets CatalogUuid field to given value.

### HasCatalogUuid

`func (o *ProductReturn) HasCatalogUuid() bool`

HasCatalogUuid returns a boolean if a field has been set.

### GetAddonList

`func (o *ProductReturn) GetAddonList() []map[string]interface{}`

GetAddonList returns the AddonList field if non-nil, zero value otherwise.

### GetAddonListOk

`func (o *ProductReturn) GetAddonListOk() (*[]map[string]interface{}, bool)`

GetAddonListOk returns a tuple with the AddonList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonList

`func (o *ProductReturn) SetAddonList(v []map[string]interface{})`

SetAddonList sets AddonList field to given value.

### HasAddonList

`func (o *ProductReturn) HasAddonList() bool`

HasAddonList returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductReturn) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductReturn) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductReturn) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductReturn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProductReturn) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductReturn) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductReturn) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProductReturn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ProductReturn) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ProductReturn) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ProductReturn) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ProductReturn) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


