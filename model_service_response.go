/*
eHelply SDK - 1.1.88

eHelply SDK for SuperStack Services

API version: 1.1.88
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceResponse struct for ServiceResponse
type ServiceResponse struct {
	Name string `json:"name"`
	Key string `json:"key"`
	Version string `json:"version"`
	Summary string `json:"summary"`
	Authors []string `json:"authors"`
	AuthorEmails []string `json:"author_emails"`
	Uuid *string `json:"uuid,omitempty"`
	Heartbeats []interface{} `json:"heartbeats,omitempty"`
	Alarms []AlarmResponse `json:"alarms,omitempty"`
	HiddenStages []string `json:"hidden_stages,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewServiceResponse instantiates a new ServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceResponse(name string, key string, version string, summary string, authors []string, authorEmails []string) *ServiceResponse {
	this := ServiceResponse{}
	this.Name = name
	this.Key = key
	this.Version = version
	this.Summary = summary
	this.Authors = authors
	this.AuthorEmails = authorEmails
	return &this
}

// NewServiceResponseWithDefaults instantiates a new ServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceResponseWithDefaults() *ServiceResponse {
	this := ServiceResponse{}
	return &this
}

// GetName returns the Name field value
func (o *ServiceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceResponse) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *ServiceResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ServiceResponse) SetKey(v string) {
	o.Key = v
}

// GetVersion returns the Version field value
func (o *ServiceResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ServiceResponse) SetVersion(v string) {
	o.Version = v
}

// GetSummary returns the Summary field value
func (o *ServiceResponse) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ServiceResponse) SetSummary(v string) {
	o.Summary = v
}

// GetAuthors returns the Authors field value
func (o *ServiceResponse) GetAuthors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetAuthorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authors, true
}

// SetAuthors sets field value
func (o *ServiceResponse) SetAuthors(v []string) {
	o.Authors = v
}

// GetAuthorEmails returns the AuthorEmails field value
func (o *ServiceResponse) GetAuthorEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorEmails
}

// GetAuthorEmailsOk returns a tuple with the AuthorEmails field value
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetAuthorEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorEmails, true
}

// SetAuthorEmails sets field value
func (o *ServiceResponse) SetAuthorEmails(v []string) {
	o.AuthorEmails = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ServiceResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ServiceResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ServiceResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetHeartbeats returns the Heartbeats field value if set, zero value otherwise.
func (o *ServiceResponse) GetHeartbeats() []interface{} {
	if o == nil || o.Heartbeats == nil {
		var ret []interface{}
		return ret
	}
	return o.Heartbeats
}

// GetHeartbeatsOk returns a tuple with the Heartbeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetHeartbeatsOk() ([]interface{}, bool) {
	if o == nil || o.Heartbeats == nil {
		return nil, false
	}
	return o.Heartbeats, true
}

// HasHeartbeats returns a boolean if a field has been set.
func (o *ServiceResponse) HasHeartbeats() bool {
	if o != nil && o.Heartbeats != nil {
		return true
	}

	return false
}

// SetHeartbeats gets a reference to the given []interface{} and assigns it to the Heartbeats field.
func (o *ServiceResponse) SetHeartbeats(v []interface{}) {
	o.Heartbeats = v
}

// GetAlarms returns the Alarms field value if set, zero value otherwise.
func (o *ServiceResponse) GetAlarms() []AlarmResponse {
	if o == nil || o.Alarms == nil {
		var ret []AlarmResponse
		return ret
	}
	return o.Alarms
}

// GetAlarmsOk returns a tuple with the Alarms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetAlarmsOk() ([]AlarmResponse, bool) {
	if o == nil || o.Alarms == nil {
		return nil, false
	}
	return o.Alarms, true
}

// HasAlarms returns a boolean if a field has been set.
func (o *ServiceResponse) HasAlarms() bool {
	if o != nil && o.Alarms != nil {
		return true
	}

	return false
}

// SetAlarms gets a reference to the given []AlarmResponse and assigns it to the Alarms field.
func (o *ServiceResponse) SetAlarms(v []AlarmResponse) {
	o.Alarms = v
}

// GetHiddenStages returns the HiddenStages field value if set, zero value otherwise.
func (o *ServiceResponse) GetHiddenStages() []string {
	if o == nil || o.HiddenStages == nil {
		var ret []string
		return ret
	}
	return o.HiddenStages
}

// GetHiddenStagesOk returns a tuple with the HiddenStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetHiddenStagesOk() ([]string, bool) {
	if o == nil || o.HiddenStages == nil {
		return nil, false
	}
	return o.HiddenStages, true
}

// HasHiddenStages returns a boolean if a field has been set.
func (o *ServiceResponse) HasHiddenStages() bool {
	if o != nil && o.HiddenStages != nil {
		return true
	}

	return false
}

// SetHiddenStages gets a reference to the given []string and assigns it to the HiddenStages field.
func (o *ServiceResponse) SetHiddenStages(v []string) {
	o.HiddenStages = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ServiceResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceResponse) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ServiceResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["authors"] = o.Authors
	}
	if true {
		toSerialize["author_emails"] = o.AuthorEmails
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Heartbeats != nil {
		toSerialize["heartbeats"] = o.Heartbeats
	}
	if o.Alarms != nil {
		toSerialize["alarms"] = o.Alarms
	}
	if o.HiddenStages != nil {
		toSerialize["hidden_stages"] = o.HiddenStages
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableServiceResponse struct {
	value *ServiceResponse
	isSet bool
}

func (v NullableServiceResponse) Get() *ServiceResponse {
	return v.value
}

func (v *NullableServiceResponse) Set(val *ServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceResponse(val *ServiceResponse) *NullableServiceResponse {
	return &NullableServiceResponse{value: val, isSet: true}
}

func (v NullableServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


