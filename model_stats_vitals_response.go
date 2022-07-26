/*
eHelply SDK - 1.1.102

eHelply SDK for SuperStack Services

API version: 1.1.102
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// StatsVitalsResponse struct for StatsVitalsResponse
type StatsVitalsResponse struct {
	Uuid *string `json:"uuid,omitempty"`
	ServiceUuid string `json:"service_uuid"`
	HeartbeatUuid string `json:"heartbeat_uuid"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Vitals map[string]interface{} `json:"vitals,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
}

// NewStatsVitalsResponse instantiates a new StatsVitalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsVitalsResponse(serviceUuid string, heartbeatUuid string) *StatsVitalsResponse {
	this := StatsVitalsResponse{}
	this.ServiceUuid = serviceUuid
	this.HeartbeatUuid = heartbeatUuid
	return &this
}

// NewStatsVitalsResponseWithDefaults instantiates a new StatsVitalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsVitalsResponseWithDefaults() *StatsVitalsResponse {
	this := StatsVitalsResponse{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StatsVitalsResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsVitalsResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StatsVitalsResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StatsVitalsResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetServiceUuid returns the ServiceUuid field value
func (o *StatsVitalsResponse) GetServiceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value
// and a boolean to check if the value has been set.
func (o *StatsVitalsResponse) GetServiceUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUuid, true
}

// SetServiceUuid sets field value
func (o *StatsVitalsResponse) SetServiceUuid(v string) {
	o.ServiceUuid = v
}

// GetHeartbeatUuid returns the HeartbeatUuid field value
func (o *StatsVitalsResponse) GetHeartbeatUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeartbeatUuid
}

// GetHeartbeatUuidOk returns a tuple with the HeartbeatUuid field value
// and a boolean to check if the value has been set.
func (o *StatsVitalsResponse) GetHeartbeatUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeartbeatUuid, true
}

// SetHeartbeatUuid sets field value
func (o *StatsVitalsResponse) SetHeartbeatUuid(v string) {
	o.HeartbeatUuid = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *StatsVitalsResponse) GetStats() map[string]interface{} {
	if o == nil || o.Stats == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsVitalsResponse) GetStatsOk() (map[string]interface{}, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *StatsVitalsResponse) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given map[string]interface{} and assigns it to the Stats field.
func (o *StatsVitalsResponse) SetStats(v map[string]interface{}) {
	o.Stats = v
}

// GetVitals returns the Vitals field value if set, zero value otherwise.
func (o *StatsVitalsResponse) GetVitals() map[string]interface{} {
	if o == nil || o.Vitals == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Vitals
}

// GetVitalsOk returns a tuple with the Vitals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsVitalsResponse) GetVitalsOk() (map[string]interface{}, bool) {
	if o == nil || o.Vitals == nil {
		return nil, false
	}
	return o.Vitals, true
}

// HasVitals returns a boolean if a field has been set.
func (o *StatsVitalsResponse) HasVitals() bool {
	if o != nil && o.Vitals != nil {
		return true
	}

	return false
}

// SetVitals gets a reference to the given map[string]interface{} and assigns it to the Vitals field.
func (o *StatsVitalsResponse) SetVitals(v map[string]interface{}) {
	o.Vitals = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StatsVitalsResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsVitalsResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StatsVitalsResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *StatsVitalsResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o StatsVitalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["service_uuid"] = o.ServiceUuid
	}
	if true {
		toSerialize["heartbeat_uuid"] = o.HeartbeatUuid
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	if o.Vitals != nil {
		toSerialize["vitals"] = o.Vitals
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableStatsVitalsResponse struct {
	value *StatsVitalsResponse
	isSet bool
}

func (v NullableStatsVitalsResponse) Get() *StatsVitalsResponse {
	return v.value
}

func (v *NullableStatsVitalsResponse) Set(val *StatsVitalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsVitalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsVitalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsVitalsResponse(val *StatsVitalsResponse) *NullableStatsVitalsResponse {
	return &NullableStatsVitalsResponse{value: val, isSet: true}
}

func (v NullableStatsVitalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsVitalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


