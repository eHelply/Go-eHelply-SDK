# NoteBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Time** | **string** |  | 
**Meta** | [**NoteMeta**](NoteMeta.md) |  | 

## Methods

### NewNoteBase

`func NewNoteBase(content string, time string, meta NoteMeta, ) *NoteBase`

NewNoteBase instantiates a new NoteBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteBaseWithDefaults

`func NewNoteBaseWithDefaults() *NoteBase`

NewNoteBaseWithDefaults instantiates a new NoteBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *NoteBase) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteBase) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteBase) SetContent(v string)`

SetContent sets Content field to given value.


### GetTime

`func (o *NoteBase) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NoteBase) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NoteBase) SetTime(v string)`

SetTime sets Time field to given value.


### GetMeta

`func (o *NoteBase) GetMeta() NoteMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NoteBase) GetMetaOk() (*NoteMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NoteBase) SetMeta(v NoteMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


