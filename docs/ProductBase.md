# ProductBase

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

## Methods

### NewProductBase

`func NewProductBase(price int32, quantityForPublic int32, ) *ProductBase`

NewProductBase instantiates a new ProductBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductBaseWithDefaults

`func NewProductBaseWithDefaults() *ProductBase`

NewProductBaseWithDefaults instantiates a new ProductBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *ProductBase) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ProductBase) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ProductBase) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ProductBase) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetCollectionUuid

`func (o *ProductBase) GetCollectionUuid() string`

GetCollectionUuid returns the CollectionUuid field if non-nil, zero value otherwise.

### GetCollectionUuidOk

`func (o *ProductBase) GetCollectionUuidOk() (*string, bool)`

GetCollectionUuidOk returns a tuple with the CollectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionUuid

`func (o *ProductBase) SetCollectionUuid(v string)`

SetCollectionUuid sets CollectionUuid field to given value.

### HasCollectionUuid

`func (o *ProductBase) HasCollectionUuid() bool`

HasCollectionUuid returns a boolean if a field has been set.

### GetReviewGroupUuid

`func (o *ProductBase) GetReviewGroupUuid() string`

GetReviewGroupUuid returns the ReviewGroupUuid field if non-nil, zero value otherwise.

### GetReviewGroupUuidOk

`func (o *ProductBase) GetReviewGroupUuidOk() (*string, bool)`

GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewGroupUuid

`func (o *ProductBase) SetReviewGroupUuid(v string)`

SetReviewGroupUuid sets ReviewGroupUuid field to given value.

### HasReviewGroupUuid

`func (o *ProductBase) HasReviewGroupUuid() bool`

HasReviewGroupUuid returns a boolean if a field has been set.

### GetAddons

`func (o *ProductBase) GetAddons() []string`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *ProductBase) GetAddonsOk() (*[]string, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *ProductBase) SetAddons(v []string)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *ProductBase) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetName

`func (o *ProductBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ProductBase) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductBase) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductBase) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetQuantityForPublic

`func (o *ProductBase) GetQuantityForPublic() int32`

GetQuantityForPublic returns the QuantityForPublic field if non-nil, zero value otherwise.

### GetQuantityForPublicOk

`func (o *ProductBase) GetQuantityForPublicOk() (*int32, bool)`

GetQuantityForPublicOk returns a tuple with the QuantityForPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityForPublic

`func (o *ProductBase) SetQuantityForPublic(v int32)`

SetQuantityForPublic sets QuantityForPublic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


