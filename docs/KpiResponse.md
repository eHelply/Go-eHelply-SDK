# KpiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ServiceUuid** | **string** |  | 
**Kpis** | **[]interface{}** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] 
**FetchedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewKpiResponse

`func NewKpiResponse(serviceUuid string, kpis []interface{}, ) *KpiResponse`

NewKpiResponse instantiates a new KpiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKpiResponseWithDefaults

`func NewKpiResponseWithDefaults() *KpiResponse`

NewKpiResponseWithDefaults instantiates a new KpiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *KpiResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *KpiResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *KpiResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *KpiResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetServiceUuid

`func (o *KpiResponse) GetServiceUuid() string`

GetServiceUuid returns the ServiceUuid field if non-nil, zero value otherwise.

### GetServiceUuidOk

`func (o *KpiResponse) GetServiceUuidOk() (*string, bool)`

GetServiceUuidOk returns a tuple with the ServiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUuid

`func (o *KpiResponse) SetServiceUuid(v string)`

SetServiceUuid sets ServiceUuid field to given value.


### GetKpis

`func (o *KpiResponse) GetKpis() []interface{}`

GetKpis returns the Kpis field if non-nil, zero value otherwise.

### GetKpisOk

`func (o *KpiResponse) GetKpisOk() (*[]interface{}, bool)`

GetKpisOk returns a tuple with the Kpis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpis

`func (o *KpiResponse) SetKpis(v []interface{})`

SetKpis sets Kpis field to given value.


### GetCreatedAt

`func (o *KpiResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KpiResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KpiResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KpiResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFetchedAt

`func (o *KpiResponse) GetFetchedAt() string`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *KpiResponse) GetFetchedAtOk() (*string, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *KpiResponse) SetFetchedAt(v string)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *KpiResponse) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


