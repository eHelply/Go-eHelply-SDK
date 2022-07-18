# HeartbeatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ServiceUuid** | **string** |  | 
**Stage** | **string** |  | 
**Process** | **string** |  | 
**Health** | **string** |  | 
**Platform** | **map[string]interface{}** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewHeartbeatResponse

`func NewHeartbeatResponse(serviceUuid string, stage string, process string, health string, platform map[string]interface{}, ) *HeartbeatResponse`

NewHeartbeatResponse instantiates a new HeartbeatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeartbeatResponseWithDefaults

`func NewHeartbeatResponseWithDefaults() *HeartbeatResponse`

NewHeartbeatResponseWithDefaults instantiates a new HeartbeatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *HeartbeatResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HeartbeatResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HeartbeatResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HeartbeatResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetServiceUuid

`func (o *HeartbeatResponse) GetServiceUuid() string`

GetServiceUuid returns the ServiceUuid field if non-nil, zero value otherwise.

### GetServiceUuidOk

`func (o *HeartbeatResponse) GetServiceUuidOk() (*string, bool)`

GetServiceUuidOk returns a tuple with the ServiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUuid

`func (o *HeartbeatResponse) SetServiceUuid(v string)`

SetServiceUuid sets ServiceUuid field to given value.


### GetStage

`func (o *HeartbeatResponse) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *HeartbeatResponse) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *HeartbeatResponse) SetStage(v string)`

SetStage sets Stage field to given value.


### GetProcess

`func (o *HeartbeatResponse) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *HeartbeatResponse) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *HeartbeatResponse) SetProcess(v string)`

SetProcess sets Process field to given value.


### GetHealth

`func (o *HeartbeatResponse) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *HeartbeatResponse) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *HeartbeatResponse) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetPlatform

`func (o *HeartbeatResponse) GetPlatform() map[string]interface{}`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *HeartbeatResponse) GetPlatformOk() (*map[string]interface{}, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *HeartbeatResponse) SetPlatform(v map[string]interface{})`

SetPlatform sets Platform field to given value.


### GetCreatedAt

`func (o *HeartbeatResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HeartbeatResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HeartbeatResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HeartbeatResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


