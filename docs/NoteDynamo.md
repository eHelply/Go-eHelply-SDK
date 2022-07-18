# NoteDynamo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Content** | Pointer to ***os.File** |  | [optional] 
**Time** | **string** |  | 
**Meta** | [**NoteMeta**](NoteMeta.md) |  | 

## Methods

### NewNoteDynamo

`func NewNoteDynamo(uuid string, time string, meta NoteMeta, ) *NoteDynamo`

NewNoteDynamo instantiates a new NoteDynamo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteDynamoWithDefaults

`func NewNoteDynamoWithDefaults() *NoteDynamo`

NewNoteDynamoWithDefaults instantiates a new NoteDynamo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NoteDynamo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NoteDynamo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NoteDynamo) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetContent

`func (o *NoteDynamo) GetContent() *os.File`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteDynamo) GetContentOk() (**os.File, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteDynamo) SetContent(v *os.File)`

SetContent sets Content field to given value.

### HasContent

`func (o *NoteDynamo) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTime

`func (o *NoteDynamo) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NoteDynamo) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NoteDynamo) SetTime(v string)`

SetTime sets Time field to given value.


### GetMeta

`func (o *NoteDynamo) GetMeta() NoteMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NoteDynamo) GetMetaOk() (*NoteMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NoteDynamo) SetMeta(v NoteMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


