/*
eHelply SDK - 1.1.108

eHelply SDK for SuperStack Services

API version: 1.1.108
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AlarmResponse struct for AlarmResponse
type AlarmResponse struct {
	Uuid *string `json:"uuid,omitempty"`
	ServiceUuid *string `json:"service_uuid,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Stage *string `json:"stage,omitempty"`
	Process *string `json:"process,omitempty"`
	Severity *string `json:"severity,omitempty"`
	TicketUuid *string `json:"ticket_uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	Note []map[string]interface{} `json:"note,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	LatestAt *string `json:"latest_at,omitempty"`
	AcknowledgedAt *string `json:"acknowledged_at,omitempty"`
	IgnoredAt *string `json:"ignored_at,omitempty"`
	ClearedAt *string `json:"cleared_at,omitempty"`
	TerminatedAt *string `json:"terminated_at,omitempty"`
	AssigneeUuid *string `json:"assignee_uuid,omitempty"`
	AcknowledgerUuid *string `json:"acknowledger_uuid,omitempty"`
	IgnorerUuid *string `json:"ignorer_uuid,omitempty"`
	TerminaterUuid *string `json:"terminater_uuid,omitempty"`
}

// NewAlarmResponse instantiates a new AlarmResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmResponse() *AlarmResponse {
	this := AlarmResponse{}
	return &this
}

// NewAlarmResponseWithDefaults instantiates a new AlarmResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmResponseWithDefaults() *AlarmResponse {
	this := AlarmResponse{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *AlarmResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *AlarmResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *AlarmResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetServiceUuid returns the ServiceUuid field value if set, zero value otherwise.
func (o *AlarmResponse) GetServiceUuid() string {
	if o == nil || o.ServiceUuid == nil {
		var ret string
		return ret
	}
	return *o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetServiceUuidOk() (*string, bool) {
	if o == nil || o.ServiceUuid == nil {
		return nil, false
	}
	return o.ServiceUuid, true
}

// HasServiceUuid returns a boolean if a field has been set.
func (o *AlarmResponse) HasServiceUuid() bool {
	if o != nil && o.ServiceUuid != nil {
		return true
	}

	return false
}

// SetServiceUuid gets a reference to the given string and assigns it to the ServiceUuid field.
func (o *AlarmResponse) SetServiceUuid(v string) {
	o.ServiceUuid = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AlarmResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AlarmResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AlarmResponse) SetCount(v int32) {
	o.Count = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *AlarmResponse) GetStage() string {
	if o == nil || o.Stage == nil {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetStageOk() (*string, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *AlarmResponse) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *AlarmResponse) SetStage(v string) {
	o.Stage = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *AlarmResponse) GetProcess() string {
	if o == nil || o.Process == nil {
		var ret string
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetProcessOk() (*string, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *AlarmResponse) HasProcess() bool {
	if o != nil && o.Process != nil {
		return true
	}

	return false
}

// SetProcess gets a reference to the given string and assigns it to the Process field.
func (o *AlarmResponse) SetProcess(v string) {
	o.Process = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AlarmResponse) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AlarmResponse) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AlarmResponse) SetSeverity(v string) {
	o.Severity = &v
}

// GetTicketUuid returns the TicketUuid field value if set, zero value otherwise.
func (o *AlarmResponse) GetTicketUuid() string {
	if o == nil || o.TicketUuid == nil {
		var ret string
		return ret
	}
	return *o.TicketUuid
}

// GetTicketUuidOk returns a tuple with the TicketUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetTicketUuidOk() (*string, bool) {
	if o == nil || o.TicketUuid == nil {
		return nil, false
	}
	return o.TicketUuid, true
}

// HasTicketUuid returns a boolean if a field has been set.
func (o *AlarmResponse) HasTicketUuid() bool {
	if o != nil && o.TicketUuid != nil {
		return true
	}

	return false
}

// SetTicketUuid gets a reference to the given string and assigns it to the TicketUuid field.
func (o *AlarmResponse) SetTicketUuid(v string) {
	o.TicketUuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlarmResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlarmResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlarmResponse) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AlarmResponse) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AlarmResponse) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AlarmResponse) SetSummary(v string) {
	o.Summary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlarmResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlarmResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlarmResponse) SetDescription(v string) {
	o.Description = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *AlarmResponse) GetNote() []map[string]interface{} {
	if o == nil || o.Note == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetNoteOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *AlarmResponse) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given []map[string]interface{} and assigns it to the Note field.
func (o *AlarmResponse) SetNote(v []map[string]interface{}) {
	o.Note = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlarmResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlarmResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AlarmResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLatestAt returns the LatestAt field value if set, zero value otherwise.
func (o *AlarmResponse) GetLatestAt() string {
	if o == nil || o.LatestAt == nil {
		var ret string
		return ret
	}
	return *o.LatestAt
}

// GetLatestAtOk returns a tuple with the LatestAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetLatestAtOk() (*string, bool) {
	if o == nil || o.LatestAt == nil {
		return nil, false
	}
	return o.LatestAt, true
}

// HasLatestAt returns a boolean if a field has been set.
func (o *AlarmResponse) HasLatestAt() bool {
	if o != nil && o.LatestAt != nil {
		return true
	}

	return false
}

// SetLatestAt gets a reference to the given string and assigns it to the LatestAt field.
func (o *AlarmResponse) SetLatestAt(v string) {
	o.LatestAt = &v
}

// GetAcknowledgedAt returns the AcknowledgedAt field value if set, zero value otherwise.
func (o *AlarmResponse) GetAcknowledgedAt() string {
	if o == nil || o.AcknowledgedAt == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgedAt
}

// GetAcknowledgedAtOk returns a tuple with the AcknowledgedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetAcknowledgedAtOk() (*string, bool) {
	if o == nil || o.AcknowledgedAt == nil {
		return nil, false
	}
	return o.AcknowledgedAt, true
}

// HasAcknowledgedAt returns a boolean if a field has been set.
func (o *AlarmResponse) HasAcknowledgedAt() bool {
	if o != nil && o.AcknowledgedAt != nil {
		return true
	}

	return false
}

// SetAcknowledgedAt gets a reference to the given string and assigns it to the AcknowledgedAt field.
func (o *AlarmResponse) SetAcknowledgedAt(v string) {
	o.AcknowledgedAt = &v
}

// GetIgnoredAt returns the IgnoredAt field value if set, zero value otherwise.
func (o *AlarmResponse) GetIgnoredAt() string {
	if o == nil || o.IgnoredAt == nil {
		var ret string
		return ret
	}
	return *o.IgnoredAt
}

// GetIgnoredAtOk returns a tuple with the IgnoredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetIgnoredAtOk() (*string, bool) {
	if o == nil || o.IgnoredAt == nil {
		return nil, false
	}
	return o.IgnoredAt, true
}

// HasIgnoredAt returns a boolean if a field has been set.
func (o *AlarmResponse) HasIgnoredAt() bool {
	if o != nil && o.IgnoredAt != nil {
		return true
	}

	return false
}

// SetIgnoredAt gets a reference to the given string and assigns it to the IgnoredAt field.
func (o *AlarmResponse) SetIgnoredAt(v string) {
	o.IgnoredAt = &v
}

// GetClearedAt returns the ClearedAt field value if set, zero value otherwise.
func (o *AlarmResponse) GetClearedAt() string {
	if o == nil || o.ClearedAt == nil {
		var ret string
		return ret
	}
	return *o.ClearedAt
}

// GetClearedAtOk returns a tuple with the ClearedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetClearedAtOk() (*string, bool) {
	if o == nil || o.ClearedAt == nil {
		return nil, false
	}
	return o.ClearedAt, true
}

// HasClearedAt returns a boolean if a field has been set.
func (o *AlarmResponse) HasClearedAt() bool {
	if o != nil && o.ClearedAt != nil {
		return true
	}

	return false
}

// SetClearedAt gets a reference to the given string and assigns it to the ClearedAt field.
func (o *AlarmResponse) SetClearedAt(v string) {
	o.ClearedAt = &v
}

// GetTerminatedAt returns the TerminatedAt field value if set, zero value otherwise.
func (o *AlarmResponse) GetTerminatedAt() string {
	if o == nil || o.TerminatedAt == nil {
		var ret string
		return ret
	}
	return *o.TerminatedAt
}

// GetTerminatedAtOk returns a tuple with the TerminatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetTerminatedAtOk() (*string, bool) {
	if o == nil || o.TerminatedAt == nil {
		return nil, false
	}
	return o.TerminatedAt, true
}

// HasTerminatedAt returns a boolean if a field has been set.
func (o *AlarmResponse) HasTerminatedAt() bool {
	if o != nil && o.TerminatedAt != nil {
		return true
	}

	return false
}

// SetTerminatedAt gets a reference to the given string and assigns it to the TerminatedAt field.
func (o *AlarmResponse) SetTerminatedAt(v string) {
	o.TerminatedAt = &v
}

// GetAssigneeUuid returns the AssigneeUuid field value if set, zero value otherwise.
func (o *AlarmResponse) GetAssigneeUuid() string {
	if o == nil || o.AssigneeUuid == nil {
		var ret string
		return ret
	}
	return *o.AssigneeUuid
}

// GetAssigneeUuidOk returns a tuple with the AssigneeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetAssigneeUuidOk() (*string, bool) {
	if o == nil || o.AssigneeUuid == nil {
		return nil, false
	}
	return o.AssigneeUuid, true
}

// HasAssigneeUuid returns a boolean if a field has been set.
func (o *AlarmResponse) HasAssigneeUuid() bool {
	if o != nil && o.AssigneeUuid != nil {
		return true
	}

	return false
}

// SetAssigneeUuid gets a reference to the given string and assigns it to the AssigneeUuid field.
func (o *AlarmResponse) SetAssigneeUuid(v string) {
	o.AssigneeUuid = &v
}

// GetAcknowledgerUuid returns the AcknowledgerUuid field value if set, zero value otherwise.
func (o *AlarmResponse) GetAcknowledgerUuid() string {
	if o == nil || o.AcknowledgerUuid == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgerUuid
}

// GetAcknowledgerUuidOk returns a tuple with the AcknowledgerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetAcknowledgerUuidOk() (*string, bool) {
	if o == nil || o.AcknowledgerUuid == nil {
		return nil, false
	}
	return o.AcknowledgerUuid, true
}

// HasAcknowledgerUuid returns a boolean if a field has been set.
func (o *AlarmResponse) HasAcknowledgerUuid() bool {
	if o != nil && o.AcknowledgerUuid != nil {
		return true
	}

	return false
}

// SetAcknowledgerUuid gets a reference to the given string and assigns it to the AcknowledgerUuid field.
func (o *AlarmResponse) SetAcknowledgerUuid(v string) {
	o.AcknowledgerUuid = &v
}

// GetIgnorerUuid returns the IgnorerUuid field value if set, zero value otherwise.
func (o *AlarmResponse) GetIgnorerUuid() string {
	if o == nil || o.IgnorerUuid == nil {
		var ret string
		return ret
	}
	return *o.IgnorerUuid
}

// GetIgnorerUuidOk returns a tuple with the IgnorerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetIgnorerUuidOk() (*string, bool) {
	if o == nil || o.IgnorerUuid == nil {
		return nil, false
	}
	return o.IgnorerUuid, true
}

// HasIgnorerUuid returns a boolean if a field has been set.
func (o *AlarmResponse) HasIgnorerUuid() bool {
	if o != nil && o.IgnorerUuid != nil {
		return true
	}

	return false
}

// SetIgnorerUuid gets a reference to the given string and assigns it to the IgnorerUuid field.
func (o *AlarmResponse) SetIgnorerUuid(v string) {
	o.IgnorerUuid = &v
}

// GetTerminaterUuid returns the TerminaterUuid field value if set, zero value otherwise.
func (o *AlarmResponse) GetTerminaterUuid() string {
	if o == nil || o.TerminaterUuid == nil {
		var ret string
		return ret
	}
	return *o.TerminaterUuid
}

// GetTerminaterUuidOk returns a tuple with the TerminaterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmResponse) GetTerminaterUuidOk() (*string, bool) {
	if o == nil || o.TerminaterUuid == nil {
		return nil, false
	}
	return o.TerminaterUuid, true
}

// HasTerminaterUuid returns a boolean if a field has been set.
func (o *AlarmResponse) HasTerminaterUuid() bool {
	if o != nil && o.TerminaterUuid != nil {
		return true
	}

	return false
}

// SetTerminaterUuid gets a reference to the given string and assigns it to the TerminaterUuid field.
func (o *AlarmResponse) SetTerminaterUuid(v string) {
	o.TerminaterUuid = &v
}

func (o AlarmResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.ServiceUuid != nil {
		toSerialize["service_uuid"] = o.ServiceUuid
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.TicketUuid != nil {
		toSerialize["ticket_uuid"] = o.TicketUuid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LatestAt != nil {
		toSerialize["latest_at"] = o.LatestAt
	}
	if o.AcknowledgedAt != nil {
		toSerialize["acknowledged_at"] = o.AcknowledgedAt
	}
	if o.IgnoredAt != nil {
		toSerialize["ignored_at"] = o.IgnoredAt
	}
	if o.ClearedAt != nil {
		toSerialize["cleared_at"] = o.ClearedAt
	}
	if o.TerminatedAt != nil {
		toSerialize["terminated_at"] = o.TerminatedAt
	}
	if o.AssigneeUuid != nil {
		toSerialize["assignee_uuid"] = o.AssigneeUuid
	}
	if o.AcknowledgerUuid != nil {
		toSerialize["acknowledger_uuid"] = o.AcknowledgerUuid
	}
	if o.IgnorerUuid != nil {
		toSerialize["ignorer_uuid"] = o.IgnorerUuid
	}
	if o.TerminaterUuid != nil {
		toSerialize["terminater_uuid"] = o.TerminaterUuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmResponse struct {
	value *AlarmResponse
	isSet bool
}

func (v NullableAlarmResponse) Get() *AlarmResponse {
	return v.value
}

func (v *NullableAlarmResponse) Set(val *AlarmResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmResponse(val *AlarmResponse) *NullableAlarmResponse {
	return &NullableAlarmResponse{value: val, isSet: true}
}

func (v NullableAlarmResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


