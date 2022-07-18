# NoteMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalAuthor** | Pointer to **string** |  | [optional] 
**Author** | **string** |  | 
**PreviousVersion** | Pointer to **string** |  | [optional] 
**NextVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewNoteMeta

`func NewNoteMeta(author string, ) *NoteMeta`

NewNoteMeta instantiates a new NoteMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteMetaWithDefaults

`func NewNoteMetaWithDefaults() *NoteMeta`

NewNoteMetaWithDefaults instantiates a new NoteMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalAuthor

`func (o *NoteMeta) GetOriginalAuthor() string`

GetOriginalAuthor returns the OriginalAuthor field if non-nil, zero value otherwise.

### GetOriginalAuthorOk

`func (o *NoteMeta) GetOriginalAuthorOk() (*string, bool)`

GetOriginalAuthorOk returns a tuple with the OriginalAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAuthor

`func (o *NoteMeta) SetOriginalAuthor(v string)`

SetOriginalAuthor sets OriginalAuthor field to given value.

### HasOriginalAuthor

`func (o *NoteMeta) HasOriginalAuthor() bool`

HasOriginalAuthor returns a boolean if a field has been set.

### GetAuthor

`func (o *NoteMeta) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NoteMeta) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NoteMeta) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetPreviousVersion

`func (o *NoteMeta) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *NoteMeta) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *NoteMeta) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *NoteMeta) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### GetNextVersion

`func (o *NoteMeta) GetNextVersion() string`

GetNextVersion returns the NextVersion field if non-nil, zero value otherwise.

### GetNextVersionOk

`func (o *NoteMeta) GetNextVersionOk() (*string, bool)`

GetNextVersionOk returns a tuple with the NextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextVersion

`func (o *NoteMeta) SetNextVersion(v string)`

SetNextVersion sets NextVersion field to given value.

### HasNextVersion

`func (o *NoteMeta) HasNextVersion() bool`

HasNextVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


