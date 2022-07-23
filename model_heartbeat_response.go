/*
eHelply SDK - 1.1.91

eHelply SDK for SuperStack Services

API version: 1.1.91
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// HeartbeatResponse struct for HeartbeatResponse
type HeartbeatResponse struct {
	Uuid *string `json:"uuid,omitempty"`
	ServiceUuid string `json:"service_uuid"`
	Stage string `json:"stage"`
	Process string `json:"process"`
	Health string `json:"health"`
	Platform map[string]interface{} `json:"platform"`
	CreatedAt *string `json:"created_at,omitempty"`
}

// NewHeartbeatResponse instantiates a new HeartbeatResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeartbeatResponse(serviceUuid string, stage string, process string, health string, platform map[string]interface{}) *HeartbeatResponse {
	this := HeartbeatResponse{}
	this.ServiceUuid = serviceUuid
	this.Stage = stage
	this.Process = process
	this.Health = health
	this.Platform = platform
	return &this
}

// NewHeartbeatResponseWithDefaults instantiates a new HeartbeatResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeartbeatResponseWithDefaults() *HeartbeatResponse {
	this := HeartbeatResponse{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HeartbeatResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartbeatResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HeartbeatResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HeartbeatResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetServiceUuid returns the ServiceUuid field value
func (o *HeartbeatResponse) GetServiceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value
// and a boolean to check if the value has been set.
func (o *HeartbeatResponse) GetServiceUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUuid, true
}

// SetServiceUuid sets field value
func (o *HeartbeatResponse) SetServiceUuid(v string) {
	o.ServiceUuid = v
}

// GetStage returns the Stage field value
func (o *HeartbeatResponse) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *HeartbeatResponse) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *HeartbeatResponse) SetStage(v string) {
	o.Stage = v
}

// GetProcess returns the Process field value
func (o *HeartbeatResponse) GetProcess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Process
}

// GetProcessOk returns a tuple with the Process field value
// and a boolean to check if the value has been set.
func (o *HeartbeatResponse) GetProcessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Process, true
}

// SetProcess sets field value
func (o *HeartbeatResponse) SetProcess(v string) {
	o.Process = v
}

// GetHealth returns the Health field value
func (o *HeartbeatResponse) GetHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *HeartbeatResponse) GetHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *HeartbeatResponse) SetHealth(v string) {
	o.Health = v
}

// GetPlatform returns the Platform field value
func (o *HeartbeatResponse) GetPlatform() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *HeartbeatResponse) GetPlatformOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform, true
}

// SetPlatform sets field value
func (o *HeartbeatResponse) SetPlatform(v map[string]interface{}) {
	o.Platform = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HeartbeatResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeartbeatResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HeartbeatResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *HeartbeatResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o HeartbeatResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["service_uuid"] = o.ServiceUuid
	}
	if true {
		toSerialize["stage"] = o.Stage
	}
	if true {
		toSerialize["process"] = o.Process
	}
	if true {
		toSerialize["health"] = o.Health
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableHeartbeatResponse struct {
	value *HeartbeatResponse
	isSet bool
}

func (v NullableHeartbeatResponse) Get() *HeartbeatResponse {
	return v.value
}

func (v *NullableHeartbeatResponse) Set(val *HeartbeatResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHeartbeatResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHeartbeatResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeartbeatResponse(val *HeartbeatResponse) *NullableHeartbeatResponse {
	return &NullableHeartbeatResponse{value: val, isSet: true}
}

func (v NullableHeartbeatResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeartbeatResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


