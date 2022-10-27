# NoteDynamoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Content** | **string** |  | 
**Time** | **string** |  | 
**Meta** | [**NoteMeta**](NoteMeta.md) |  | 

## Methods

### NewNoteDynamoResponse

`func NewNoteDynamoResponse(uuid string, content string, time string, meta NoteMeta, ) *NoteDynamoResponse`

NewNoteDynamoResponse instantiates a new NoteDynamoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteDynamoResponseWithDefaults

`func NewNoteDynamoResponseWithDefaults() *NoteDynamoResponse`

NewNoteDynamoResponseWithDefaults instantiates a new NoteDynamoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NoteDynamoResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NoteDynamoResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NoteDynamoResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetContent

`func (o *NoteDynamoResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteDynamoResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteDynamoResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetTime

`func (o *NoteDynamoResponse) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NoteDynamoResponse) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NoteDynamoResponse) SetTime(v string)`

SetTime sets Time field to given value.


### GetMeta

`func (o *NoteDynamoResponse) GetMeta() NoteMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NoteDynamoResponse) GetMetaOk() (*NoteMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NoteDynamoResponse) SetMeta(v NoteMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


