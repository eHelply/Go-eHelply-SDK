# StatsVitalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ServiceUuid** | **string** |  | 
**HeartbeatUuid** | **string** |  | 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 
**Vitals** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewStatsVitalsResponse

`func NewStatsVitalsResponse(serviceUuid string, heartbeatUuid string, ) *StatsVitalsResponse`

NewStatsVitalsResponse instantiates a new StatsVitalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsVitalsResponseWithDefaults

`func NewStatsVitalsResponseWithDefaults() *StatsVitalsResponse`

NewStatsVitalsResponseWithDefaults instantiates a new StatsVitalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *StatsVitalsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StatsVitalsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StatsVitalsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StatsVitalsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetServiceUuid

`func (o *StatsVitalsResponse) GetServiceUuid() string`

GetServiceUuid returns the ServiceUuid field if non-nil, zero value otherwise.

### GetServiceUuidOk

`func (o *StatsVitalsResponse) GetServiceUuidOk() (*string, bool)`

GetServiceUuidOk returns a tuple with the ServiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUuid

`func (o *StatsVitalsResponse) SetServiceUuid(v string)`

SetServiceUuid sets ServiceUuid field to given value.


### GetHeartbeatUuid

`func (o *StatsVitalsResponse) GetHeartbeatUuid() string`

GetHeartbeatUuid returns the HeartbeatUuid field if non-nil, zero value otherwise.

### GetHeartbeatUuidOk

`func (o *StatsVitalsResponse) GetHeartbeatUuidOk() (*string, bool)`

GetHeartbeatUuidOk returns a tuple with the HeartbeatUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatUuid

`func (o *StatsVitalsResponse) SetHeartbeatUuid(v string)`

SetHeartbeatUuid sets HeartbeatUuid field to given value.


### GetStats

`func (o *StatsVitalsResponse) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StatsVitalsResponse) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StatsVitalsResponse) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *StatsVitalsResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetVitals

`func (o *StatsVitalsResponse) GetVitals() map[string]interface{}`

GetVitals returns the Vitals field if non-nil, zero value otherwise.

### GetVitalsOk

`func (o *StatsVitalsResponse) GetVitalsOk() (*map[string]interface{}, bool)`

GetVitalsOk returns a tuple with the Vitals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVitals

`func (o *StatsVitalsResponse) SetVitals(v map[string]interface{})`

SetVitals sets Vitals field to given value.

### HasVitals

`func (o *StatsVitalsResponse) HasVitals() bool`

HasVitals returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StatsVitalsResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StatsVitalsResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StatsVitalsResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StatsVitalsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


