# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantUuid** | **string** |  | 
**Content** | **string** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewNote

`func NewNote(participantUuid string, content string, createdAt string, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipantUuid

`func (o *Note) GetParticipantUuid() string`

GetParticipantUuid returns the ParticipantUuid field if non-nil, zero value otherwise.

### GetParticipantUuidOk

`func (o *Note) GetParticipantUuidOk() (*string, bool)`

GetParticipantUuidOk returns a tuple with the ParticipantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantUuid

`func (o *Note) SetParticipantUuid(v string)`

SetParticipantUuid sets ParticipantUuid field to given value.


### GetContent

`func (o *Note) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Note) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Note) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *Note) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Note) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Note) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


