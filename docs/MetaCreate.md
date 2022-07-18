# MetaCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**BasicMetaCreate**](BasicMetaCreate.md) |  | [optional] 
**Detailed** | Pointer to [**DetailedMetaCreate**](DetailedMetaCreate.md) |  | [optional] 
**Custom** | Pointer to [**MetaCustom**](MetaCustom.md) |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**Children** | Pointer to [**[]MetaChildren**](MetaChildren.md) |  | [optional] 
**ParentUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaCreate

`func NewMetaCreate() *MetaCreate`

NewMetaCreate instantiates a new MetaCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaCreateWithDefaults

`func NewMetaCreateWithDefaults() *MetaCreate`

NewMetaCreateWithDefaults instantiates a new MetaCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *MetaCreate) GetBasic() BasicMetaCreate`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *MetaCreate) GetBasicOk() (*BasicMetaCreate, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *MetaCreate) SetBasic(v BasicMetaCreate)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *MetaCreate) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetDetailed

`func (o *MetaCreate) GetDetailed() DetailedMetaCreate`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *MetaCreate) GetDetailedOk() (*DetailedMetaCreate, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *MetaCreate) SetDetailed(v DetailedMetaCreate)`

SetDetailed sets Detailed field to given value.

### HasDetailed

`func (o *MetaCreate) HasDetailed() bool`

HasDetailed returns a boolean if a field has been set.

### GetCustom

`func (o *MetaCreate) GetCustom() MetaCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *MetaCreate) GetCustomOk() (*MetaCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *MetaCreate) SetCustom(v MetaCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *MetaCreate) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetFields

`func (o *MetaCreate) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaCreate) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaCreate) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetChildren

`func (o *MetaCreate) GetChildren() []MetaChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MetaCreate) GetChildrenOk() (*[]MetaChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MetaCreate) SetChildren(v []MetaChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MetaCreate) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetParentUuid

`func (o *MetaCreate) GetParentUuid() string`

GetParentUuid returns the ParentUuid field if non-nil, zero value otherwise.

### GetParentUuidOk

`func (o *MetaCreate) GetParentUuidOk() (*string, bool)`

GetParentUuidOk returns a tuple with the ParentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuid

`func (o *MetaCreate) SetParentUuid(v string)`

SetParentUuid sets ParentUuid field to given value.

### HasParentUuid

`func (o *MetaCreate) HasParentUuid() bool`

HasParentUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


