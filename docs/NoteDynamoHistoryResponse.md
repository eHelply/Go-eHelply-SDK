# NoteDynamoHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Content** | **string** |  | 
**Time** | **string** |  | 
**Meta** | [**NoteMeta**](NoteMeta.md) |  | 
**History** | Pointer to [**[]NoteDynamoResponse**](NoteDynamoResponse.md) |  | [optional] [default to []]

## Methods

### NewNoteDynamoHistoryResponse

`func NewNoteDynamoHistoryResponse(uuid string, content string, time string, meta NoteMeta, ) *NoteDynamoHistoryResponse`

NewNoteDynamoHistoryResponse instantiates a new NoteDynamoHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteDynamoHistoryResponseWithDefaults

`func NewNoteDynamoHistoryResponseWithDefaults() *NoteDynamoHistoryResponse`

NewNoteDynamoHistoryResponseWithDefaults instantiates a new NoteDynamoHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NoteDynamoHistoryResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NoteDynamoHistoryResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NoteDynamoHistoryResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetContent

`func (o *NoteDynamoHistoryResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteDynamoHistoryResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteDynamoHistoryResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetTime

`func (o *NoteDynamoHistoryResponse) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NoteDynamoHistoryResponse) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NoteDynamoHistoryResponse) SetTime(v string)`

SetTime sets Time field to given value.


### GetMeta

`func (o *NoteDynamoHistoryResponse) GetMeta() NoteMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NoteDynamoHistoryResponse) GetMetaOk() (*NoteMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NoteDynamoHistoryResponse) SetMeta(v NoteMeta)`

SetMeta sets Meta field to given value.


### GetHistory

`func (o *NoteDynamoHistoryResponse) GetHistory() []NoteDynamoResponse`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *NoteDynamoHistoryResponse) GetHistoryOk() (*[]NoteDynamoResponse, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *NoteDynamoHistoryResponse) SetHistory(v []NoteDynamoResponse)`

SetHistory sets History field to given value.

### HasHistory

`func (o *NoteDynamoHistoryResponse) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


