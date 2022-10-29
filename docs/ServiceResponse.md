# ServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Key** | **string** |  | 
**Version** | **string** |  | 
**Summary** | **string** |  | 
**Authors** | **[]string** |  | 
**AuthorEmails** | **[]string** |  | 
**Uuid** | Pointer to **string** |  | [optional] 
**Heartbeats** | Pointer to **[]interface{}** |  | [optional] 
**Alarms** | Pointer to [**[]AlarmResponse**](AlarmResponse.md) |  | [optional] 
**HiddenStages** | Pointer to **[]string** |  | [optional] [default to []]
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**SuperstackMeta** | Pointer to [**ServiceSuperStackMeta**](ServiceSuperStackMeta.md) |  | [optional] 

## Methods

### NewServiceResponse

`func NewServiceResponse(name string, key string, version string, summary string, authors []string, authorEmails []string, ) *ServiceResponse`

NewServiceResponse instantiates a new ServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseWithDefaults

`func NewServiceResponseWithDefaults() *ServiceResponse`

NewServiceResponseWithDefaults instantiates a new ServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *ServiceResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ServiceResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ServiceResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *ServiceResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSummary

`func (o *ServiceResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ServiceResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ServiceResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetAuthors

`func (o *ServiceResponse) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ServiceResponse) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ServiceResponse) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.


### GetAuthorEmails

`func (o *ServiceResponse) GetAuthorEmails() []string`

GetAuthorEmails returns the AuthorEmails field if non-nil, zero value otherwise.

### GetAuthorEmailsOk

`func (o *ServiceResponse) GetAuthorEmailsOk() (*[]string, bool)`

GetAuthorEmailsOk returns a tuple with the AuthorEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmails

`func (o *ServiceResponse) SetAuthorEmails(v []string)`

SetAuthorEmails sets AuthorEmails field to given value.


### GetUuid

`func (o *ServiceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHeartbeats

`func (o *ServiceResponse) GetHeartbeats() []interface{}`

GetHeartbeats returns the Heartbeats field if non-nil, zero value otherwise.

### GetHeartbeatsOk

`func (o *ServiceResponse) GetHeartbeatsOk() (*[]interface{}, bool)`

GetHeartbeatsOk returns a tuple with the Heartbeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeats

`func (o *ServiceResponse) SetHeartbeats(v []interface{})`

SetHeartbeats sets Heartbeats field to given value.

### HasHeartbeats

`func (o *ServiceResponse) HasHeartbeats() bool`

HasHeartbeats returns a boolean if a field has been set.

### GetAlarms

`func (o *ServiceResponse) GetAlarms() []AlarmResponse`

GetAlarms returns the Alarms field if non-nil, zero value otherwise.

### GetAlarmsOk

`func (o *ServiceResponse) GetAlarmsOk() (*[]AlarmResponse, bool)`

GetAlarmsOk returns a tuple with the Alarms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarms

`func (o *ServiceResponse) SetAlarms(v []AlarmResponse)`

SetAlarms sets Alarms field to given value.

### HasAlarms

`func (o *ServiceResponse) HasAlarms() bool`

HasAlarms returns a boolean if a field has been set.

### GetHiddenStages

`func (o *ServiceResponse) GetHiddenStages() []string`

GetHiddenStages returns the HiddenStages field if non-nil, zero value otherwise.

### GetHiddenStagesOk

`func (o *ServiceResponse) GetHiddenStagesOk() (*[]string, bool)`

GetHiddenStagesOk returns a tuple with the HiddenStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenStages

`func (o *ServiceResponse) SetHiddenStages(v []string)`

SetHiddenStages sets HiddenStages field to given value.

### HasHiddenStages

`func (o *ServiceResponse) HasHiddenStages() bool`

HasHiddenStages returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSuperstackMeta

`func (o *ServiceResponse) GetSuperstackMeta() ServiceSuperStackMeta`

GetSuperstackMeta returns the SuperstackMeta field if non-nil, zero value otherwise.

### GetSuperstackMetaOk

`func (o *ServiceResponse) GetSuperstackMetaOk() (*ServiceSuperStackMeta, bool)`

GetSuperstackMetaOk returns a tuple with the SuperstackMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperstackMeta

`func (o *ServiceResponse) SetSuperstackMeta(v ServiceSuperStackMeta)`

SetSuperstackMeta sets SuperstackMeta field to given value.

### HasSuperstackMeta

`func (o *ServiceResponse) HasSuperstackMeta() bool`

HasSuperstackMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


