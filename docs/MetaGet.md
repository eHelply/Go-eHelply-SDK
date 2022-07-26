# MetaGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Basic** | [**BasicMeta**](BasicMeta.md) |  | 
**Detailed** | [**DetailedMetaGet**](DetailedMetaGet.md) |  | 
**Custom** | Pointer to [**MetaCustom**](MetaCustom.md) |  | [optional] 
**Dates** | Pointer to [**DatesMeta**](DatesMeta.md) |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**Children** | Pointer to [**[]MetaChildren**](MetaChildren.md) |  | [optional] 
**ParentUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaGet

`func NewMetaGet(uuid string, basic BasicMeta, detailed DetailedMetaGet, ) *MetaGet`

NewMetaGet instantiates a new MetaGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaGetWithDefaults

`func NewMetaGetWithDefaults() *MetaGet`

NewMetaGetWithDefaults instantiates a new MetaGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *MetaGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MetaGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MetaGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetBasic

`func (o *MetaGet) GetBasic() BasicMeta`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *MetaGet) GetBasicOk() (*BasicMeta, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *MetaGet) SetBasic(v BasicMeta)`

SetBasic sets Basic field to given value.


### GetDetailed

`func (o *MetaGet) GetDetailed() DetailedMetaGet`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *MetaGet) GetDetailedOk() (*DetailedMetaGet, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *MetaGet) SetDetailed(v DetailedMetaGet)`

SetDetailed sets Detailed field to given value.


### GetCustom

`func (o *MetaGet) GetCustom() MetaCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *MetaGet) GetCustomOk() (*MetaCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *MetaGet) SetCustom(v MetaCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *MetaGet) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDates

`func (o *MetaGet) GetDates() DatesMeta`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *MetaGet) GetDatesOk() (*DatesMeta, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *MetaGet) SetDates(v DatesMeta)`

SetDates sets Dates field to given value.

### HasDates

`func (o *MetaGet) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetFields

`func (o *MetaGet) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaGet) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaGet) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaGet) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetChildren

`func (o *MetaGet) GetChildren() []MetaChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MetaGet) GetChildrenOk() (*[]MetaChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MetaGet) SetChildren(v []MetaChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MetaGet) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetParentUuid

`func (o *MetaGet) GetParentUuid() string`

GetParentUuid returns the ParentUuid field if non-nil, zero value otherwise.

### GetParentUuidOk

`func (o *MetaGet) GetParentUuidOk() (*string, bool)`

GetParentUuidOk returns a tuple with the ParentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuid

`func (o *MetaGet) SetParentUuid(v string)`

SetParentUuid sets ParentUuid field to given value.

### HasParentUuid

`func (o *MetaGet) HasParentUuid() bool`

HasParentUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


