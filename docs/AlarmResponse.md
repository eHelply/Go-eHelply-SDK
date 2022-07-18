# AlarmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ServiceUuid** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**Process** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**TicketUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**LatestAt** | Pointer to **string** |  | [optional] 
**AcknowledgedAt** | Pointer to **string** |  | [optional] 
**IgnoredAt** | Pointer to **string** |  | [optional] 
**ClearedAt** | Pointer to **string** |  | [optional] 
**TerminatedAt** | Pointer to **string** |  | [optional] 
**AssigneeUuid** | Pointer to **string** |  | [optional] 
**AcknowledgerUuid** | Pointer to **string** |  | [optional] 
**IgnorerUuid** | Pointer to **string** |  | [optional] 
**TerminaterUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewAlarmResponse

`func NewAlarmResponse() *AlarmResponse`

NewAlarmResponse instantiates a new AlarmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmResponseWithDefaults

`func NewAlarmResponseWithDefaults() *AlarmResponse`

NewAlarmResponseWithDefaults instantiates a new AlarmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AlarmResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlarmResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlarmResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AlarmResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetServiceUuid

`func (o *AlarmResponse) GetServiceUuid() string`

GetServiceUuid returns the ServiceUuid field if non-nil, zero value otherwise.

### GetServiceUuidOk

`func (o *AlarmResponse) GetServiceUuidOk() (*string, bool)`

GetServiceUuidOk returns a tuple with the ServiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUuid

`func (o *AlarmResponse) SetServiceUuid(v string)`

SetServiceUuid sets ServiceUuid field to given value.

### HasServiceUuid

`func (o *AlarmResponse) HasServiceUuid() bool`

HasServiceUuid returns a boolean if a field has been set.

### GetCount

`func (o *AlarmResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AlarmResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AlarmResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AlarmResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStage

`func (o *AlarmResponse) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *AlarmResponse) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *AlarmResponse) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *AlarmResponse) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetProcess

`func (o *AlarmResponse) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *AlarmResponse) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *AlarmResponse) SetProcess(v string)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *AlarmResponse) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetSeverity

`func (o *AlarmResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlarmResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlarmResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlarmResponse) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTicketUuid

`func (o *AlarmResponse) GetTicketUuid() string`

GetTicketUuid returns the TicketUuid field if non-nil, zero value otherwise.

### GetTicketUuidOk

`func (o *AlarmResponse) GetTicketUuidOk() (*string, bool)`

GetTicketUuidOk returns a tuple with the TicketUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUuid

`func (o *AlarmResponse) SetTicketUuid(v string)`

SetTicketUuid sets TicketUuid field to given value.

### HasTicketUuid

`func (o *AlarmResponse) HasTicketUuid() bool`

HasTicketUuid returns a boolean if a field has been set.

### GetName

`func (o *AlarmResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlarmResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlarmResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlarmResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *AlarmResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AlarmResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AlarmResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AlarmResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDescription

`func (o *AlarmResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlarmResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlarmResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlarmResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNote

`func (o *AlarmResponse) GetNote() []map[string]interface{}`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AlarmResponse) GetNoteOk() (*[]map[string]interface{}, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AlarmResponse) SetNote(v []map[string]interface{})`

SetNote sets Note field to given value.

### HasNote

`func (o *AlarmResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlarmResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlarmResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlarmResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlarmResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLatestAt

`func (o *AlarmResponse) GetLatestAt() string`

GetLatestAt returns the LatestAt field if non-nil, zero value otherwise.

### GetLatestAtOk

`func (o *AlarmResponse) GetLatestAtOk() (*string, bool)`

GetLatestAtOk returns a tuple with the LatestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAt

`func (o *AlarmResponse) SetLatestAt(v string)`

SetLatestAt sets LatestAt field to given value.

### HasLatestAt

`func (o *AlarmResponse) HasLatestAt() bool`

HasLatestAt returns a boolean if a field has been set.

### GetAcknowledgedAt

`func (o *AlarmResponse) GetAcknowledgedAt() string`

GetAcknowledgedAt returns the AcknowledgedAt field if non-nil, zero value otherwise.

### GetAcknowledgedAtOk

`func (o *AlarmResponse) GetAcknowledgedAtOk() (*string, bool)`

GetAcknowledgedAtOk returns a tuple with the AcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedAt

`func (o *AlarmResponse) SetAcknowledgedAt(v string)`

SetAcknowledgedAt sets AcknowledgedAt field to given value.

### HasAcknowledgedAt

`func (o *AlarmResponse) HasAcknowledgedAt() bool`

HasAcknowledgedAt returns a boolean if a field has been set.

### GetIgnoredAt

`func (o *AlarmResponse) GetIgnoredAt() string`

GetIgnoredAt returns the IgnoredAt field if non-nil, zero value otherwise.

### GetIgnoredAtOk

`func (o *AlarmResponse) GetIgnoredAtOk() (*string, bool)`

GetIgnoredAtOk returns a tuple with the IgnoredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredAt

`func (o *AlarmResponse) SetIgnoredAt(v string)`

SetIgnoredAt sets IgnoredAt field to given value.

### HasIgnoredAt

`func (o *AlarmResponse) HasIgnoredAt() bool`

HasIgnoredAt returns a boolean if a field has been set.

### GetClearedAt

`func (o *AlarmResponse) GetClearedAt() string`

GetClearedAt returns the ClearedAt field if non-nil, zero value otherwise.

### GetClearedAtOk

`func (o *AlarmResponse) GetClearedAtOk() (*string, bool)`

GetClearedAtOk returns a tuple with the ClearedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearedAt

`func (o *AlarmResponse) SetClearedAt(v string)`

SetClearedAt sets ClearedAt field to given value.

### HasClearedAt

`func (o *AlarmResponse) HasClearedAt() bool`

HasClearedAt returns a boolean if a field has been set.

### GetTerminatedAt

`func (o *AlarmResponse) GetTerminatedAt() string`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *AlarmResponse) GetTerminatedAtOk() (*string, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *AlarmResponse) SetTerminatedAt(v string)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *AlarmResponse) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### GetAssigneeUuid

`func (o *AlarmResponse) GetAssigneeUuid() string`

GetAssigneeUuid returns the AssigneeUuid field if non-nil, zero value otherwise.

### GetAssigneeUuidOk

`func (o *AlarmResponse) GetAssigneeUuidOk() (*string, bool)`

GetAssigneeUuidOk returns a tuple with the AssigneeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeUuid

`func (o *AlarmResponse) SetAssigneeUuid(v string)`

SetAssigneeUuid sets AssigneeUuid field to given value.

### HasAssigneeUuid

`func (o *AlarmResponse) HasAssigneeUuid() bool`

HasAssigneeUuid returns a boolean if a field has been set.

### GetAcknowledgerUuid

`func (o *AlarmResponse) GetAcknowledgerUuid() string`

GetAcknowledgerUuid returns the AcknowledgerUuid field if non-nil, zero value otherwise.

### GetAcknowledgerUuidOk

`func (o *AlarmResponse) GetAcknowledgerUuidOk() (*string, bool)`

GetAcknowledgerUuidOk returns a tuple with the AcknowledgerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgerUuid

`func (o *AlarmResponse) SetAcknowledgerUuid(v string)`

SetAcknowledgerUuid sets AcknowledgerUuid field to given value.

### HasAcknowledgerUuid

`func (o *AlarmResponse) HasAcknowledgerUuid() bool`

HasAcknowledgerUuid returns a boolean if a field has been set.

### GetIgnorerUuid

`func (o *AlarmResponse) GetIgnorerUuid() string`

GetIgnorerUuid returns the IgnorerUuid field if non-nil, zero value otherwise.

### GetIgnorerUuidOk

`func (o *AlarmResponse) GetIgnorerUuidOk() (*string, bool)`

GetIgnorerUuidOk returns a tuple with the IgnorerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorerUuid

`func (o *AlarmResponse) SetIgnorerUuid(v string)`

SetIgnorerUuid sets IgnorerUuid field to given value.

### HasIgnorerUuid

`func (o *AlarmResponse) HasIgnorerUuid() bool`

HasIgnorerUuid returns a boolean if a field has been set.

### GetTerminaterUuid

`func (o *AlarmResponse) GetTerminaterUuid() string`

GetTerminaterUuid returns the TerminaterUuid field if non-nil, zero value otherwise.

### GetTerminaterUuidOk

`func (o *AlarmResponse) GetTerminaterUuidOk() (*string, bool)`

GetTerminaterUuidOk returns a tuple with the TerminaterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminaterUuid

`func (o *AlarmResponse) SetTerminaterUuid(v string)`

SetTerminaterUuid sets TerminaterUuid field to given value.

### HasTerminaterUuid

`func (o *AlarmResponse) HasTerminaterUuid() bool`

HasTerminaterUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


