# MetaCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**List** | Pointer to [**[]CustomList**](CustomList.md) |  | [optional] 

## Methods

### NewMetaCustom

`func NewMetaCustom() *MetaCustom`

NewMetaCustom instantiates a new MetaCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaCustomWithDefaults

`func NewMetaCustomWithDefaults() *MetaCustom`

NewMetaCustomWithDefaults instantiates a new MetaCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetaCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaCustom) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaCustom) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MetaCustom) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetaCustom) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetaCustom) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetaCustom) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetList

`func (o *MetaCustom) GetList() []CustomList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MetaCustom) GetListOk() (*[]CustomList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MetaCustom) SetList(v []CustomList)`

SetList sets List field to given value.

### HasList

`func (o *MetaCustom) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


