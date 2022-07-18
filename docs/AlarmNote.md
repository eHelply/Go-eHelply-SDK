# AlarmNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorUuid** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewAlarmNote

`func NewAlarmNote(authorUuid string, message string, ) *AlarmNote`

NewAlarmNote instantiates a new AlarmNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmNoteWithDefaults

`func NewAlarmNoteWithDefaults() *AlarmNote`

NewAlarmNoteWithDefaults instantiates a new AlarmNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorUuid

`func (o *AlarmNote) GetAuthorUuid() string`

GetAuthorUuid returns the AuthorUuid field if non-nil, zero value otherwise.

### GetAuthorUuidOk

`func (o *AlarmNote) GetAuthorUuidOk() (*string, bool)`

GetAuthorUuidOk returns a tuple with the AuthorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorUuid

`func (o *AlarmNote) SetAuthorUuid(v string)`

SetAuthorUuid sets AuthorUuid field to given value.


### GetMessage

`func (o *AlarmNote) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlarmNote) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlarmNote) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


