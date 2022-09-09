# TagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Name** | **string** |  | 
**ProjectUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewTagResponse

`func NewTagResponse(uuid string, name string, ) *TagResponse`

NewTagResponse instantiates a new TagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResponseWithDefaults

`func NewTagResponseWithDefaults() *TagResponse`

NewTagResponseWithDefaults instantiates a new TagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *TagResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TagResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TagResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *TagResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProjectUuid

`func (o *TagResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *TagResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *TagResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *TagResponse) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


