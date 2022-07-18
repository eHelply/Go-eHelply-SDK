# ParticipantUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** | Dictionary containing any data you would like | [optional] 
**UserUuid** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewParticipantUpdate

`func NewParticipantUpdate(uuid string, ) *ParticipantUpdate`

NewParticipantUpdate instantiates a new ParticipantUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantUpdateWithDefaults

`func NewParticipantUpdateWithDefaults() *ParticipantUpdate`

NewParticipantUpdateWithDefaults instantiates a new ParticipantUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ParticipantUpdate) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ParticipantUpdate) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ParticipantUpdate) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ParticipantUpdate) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUserUuid

`func (o *ParticipantUpdate) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *ParticipantUpdate) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *ParticipantUpdate) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *ParticipantUpdate) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

### GetUuid

`func (o *ParticipantUpdate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ParticipantUpdate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ParticipantUpdate) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


