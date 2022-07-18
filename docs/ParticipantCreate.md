# ParticipantCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** | Dictionary containing any data you would like | [optional] 
**UserUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewParticipantCreate

`func NewParticipantCreate() *ParticipantCreate`

NewParticipantCreate instantiates a new ParticipantCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantCreateWithDefaults

`func NewParticipantCreateWithDefaults() *ParticipantCreate`

NewParticipantCreateWithDefaults instantiates a new ParticipantCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ParticipantCreate) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ParticipantCreate) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ParticipantCreate) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ParticipantCreate) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUserUuid

`func (o *ParticipantCreate) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *ParticipantCreate) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *ParticipantCreate) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *ParticipantCreate) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


