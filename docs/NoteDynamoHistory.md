# NoteDynamoHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Content** | Pointer to ***os.File** |  | [optional] 
**Time** | **string** |  | 
**Meta** | [**NoteMeta**](NoteMeta.md) |  | 
**History** | Pointer to [**[]NoteDynamoResponse**](NoteDynamoResponse.md) |  | [optional] [default to []]

## Methods

### NewNoteDynamoHistory

`func NewNoteDynamoHistory(uuid string, time string, meta NoteMeta, ) *NoteDynamoHistory`

NewNoteDynamoHistory instantiates a new NoteDynamoHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteDynamoHistoryWithDefaults

`func NewNoteDynamoHistoryWithDefaults() *NoteDynamoHistory`

NewNoteDynamoHistoryWithDefaults instantiates a new NoteDynamoHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NoteDynamoHistory) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NoteDynamoHistory) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NoteDynamoHistory) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetContent

`func (o *NoteDynamoHistory) GetContent() *os.File`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteDynamoHistory) GetContentOk() (**os.File, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteDynamoHistory) SetContent(v *os.File)`

SetContent sets Content field to given value.

### HasContent

`func (o *NoteDynamoHistory) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTime

`func (o *NoteDynamoHistory) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NoteDynamoHistory) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NoteDynamoHistory) SetTime(v string)`

SetTime sets Time field to given value.


### GetMeta

`func (o *NoteDynamoHistory) GetMeta() NoteMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NoteDynamoHistory) GetMetaOk() (*NoteMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NoteDynamoHistory) SetMeta(v NoteMeta)`

SetMeta sets Meta field to given value.


### GetHistory

`func (o *NoteDynamoHistory) GetHistory() []NoteDynamoResponse`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *NoteDynamoHistory) GetHistoryOk() (*[]NoteDynamoResponse, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *NoteDynamoHistory) SetHistory(v []NoteDynamoResponse)`

SetHistory sets History field to given value.

### HasHistory

`func (o *NoteDynamoHistory) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


