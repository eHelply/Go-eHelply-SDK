# MetaDynamo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**Basic**](Basic.md) |  | [optional] 
**Detailed** | Pointer to [**Detailed**](Detailed.md) |  | [optional] 
**Custom** | Pointer to [**MetaCustom**](MetaCustom.md) |  | [optional] 
**Dates** | Pointer to [**DatesMeta**](DatesMeta.md) |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**Children** | Pointer to [**[]MetaChildren**](MetaChildren.md) |  | [optional] [default to []]
**ParentUuid** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewMetaDynamo

`func NewMetaDynamo(uuid string, ) *MetaDynamo`

NewMetaDynamo instantiates a new MetaDynamo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDynamoWithDefaults

`func NewMetaDynamoWithDefaults() *MetaDynamo`

NewMetaDynamoWithDefaults instantiates a new MetaDynamo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *MetaDynamo) GetBasic() Basic`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *MetaDynamo) GetBasicOk() (*Basic, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *MetaDynamo) SetBasic(v Basic)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *MetaDynamo) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetDetailed

`func (o *MetaDynamo) GetDetailed() Detailed`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *MetaDynamo) GetDetailedOk() (*Detailed, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *MetaDynamo) SetDetailed(v Detailed)`

SetDetailed sets Detailed field to given value.

### HasDetailed

`func (o *MetaDynamo) HasDetailed() bool`

HasDetailed returns a boolean if a field has been set.

### GetCustom

`func (o *MetaDynamo) GetCustom() MetaCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *MetaDynamo) GetCustomOk() (*MetaCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *MetaDynamo) SetCustom(v MetaCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *MetaDynamo) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDates

`func (o *MetaDynamo) GetDates() DatesMeta`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *MetaDynamo) GetDatesOk() (*DatesMeta, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *MetaDynamo) SetDates(v DatesMeta)`

SetDates sets Dates field to given value.

### HasDates

`func (o *MetaDynamo) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetFields

`func (o *MetaDynamo) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaDynamo) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaDynamo) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaDynamo) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetChildren

`func (o *MetaDynamo) GetChildren() []MetaChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MetaDynamo) GetChildrenOk() (*[]MetaChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MetaDynamo) SetChildren(v []MetaChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *MetaDynamo) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetParentUuid

`func (o *MetaDynamo) GetParentUuid() string`

GetParentUuid returns the ParentUuid field if non-nil, zero value otherwise.

### GetParentUuidOk

`func (o *MetaDynamo) GetParentUuidOk() (*string, bool)`

GetParentUuidOk returns a tuple with the ParentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuid

`func (o *MetaDynamo) SetParentUuid(v string)`

SetParentUuid sets ParentUuid field to given value.

### HasParentUuid

`func (o *MetaDynamo) HasParentUuid() bool`

HasParentUuid returns a boolean if a field has been set.

### GetUuid

`func (o *MetaDynamo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MetaDynamo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MetaDynamo) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


