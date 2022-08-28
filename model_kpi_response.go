/*
eHelply SDK - 1.1.106

eHelply SDK for SuperStack Services

API version: 1.1.106
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// KpiResponse struct for KpiResponse
type KpiResponse struct {
	Uuid *string `json:"uuid,omitempty"`
	ServiceUuid string `json:"service_uuid"`
	Kpis []interface{} `json:"kpis"`
	CreatedAt *string `json:"created_at,omitempty"`
	FetchedAt *string `json:"fetched_at,omitempty"`
}

// NewKpiResponse instantiates a new KpiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKpiResponse(serviceUuid string, kpis []interface{}) *KpiResponse {
	this := KpiResponse{}
	this.ServiceUuid = serviceUuid
	this.Kpis = kpis
	return &this
}

// NewKpiResponseWithDefaults instantiates a new KpiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKpiResponseWithDefaults() *KpiResponse {
	this := KpiResponse{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *KpiResponse) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KpiResponse) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *KpiResponse) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *KpiResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetServiceUuid returns the ServiceUuid field value
func (o *KpiResponse) GetServiceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value
// and a boolean to check if the value has been set.
func (o *KpiResponse) GetServiceUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceUuid, true
}

// SetServiceUuid sets field value
func (o *KpiResponse) SetServiceUuid(v string) {
	o.ServiceUuid = v
}

// GetKpis returns the Kpis field value
func (o *KpiResponse) GetKpis() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Kpis
}

// GetKpisOk returns a tuple with the Kpis field value
// and a boolean to check if the value has been set.
func (o *KpiResponse) GetKpisOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kpis, true
}

// SetKpis sets field value
func (o *KpiResponse) SetKpis(v []interface{}) {
	o.Kpis = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KpiResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KpiResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KpiResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *KpiResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetFetchedAt returns the FetchedAt field value if set, zero value otherwise.
func (o *KpiResponse) GetFetchedAt() string {
	if o == nil || o.FetchedAt == nil {
		var ret string
		return ret
	}
	return *o.FetchedAt
}

// GetFetchedAtOk returns a tuple with the FetchedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KpiResponse) GetFetchedAtOk() (*string, bool) {
	if o == nil || o.FetchedAt == nil {
		return nil, false
	}
	return o.FetchedAt, true
}

// HasFetchedAt returns a boolean if a field has been set.
func (o *KpiResponse) HasFetchedAt() bool {
	if o != nil && o.FetchedAt != nil {
		return true
	}

	return false
}

// SetFetchedAt gets a reference to the given string and assigns it to the FetchedAt field.
func (o *KpiResponse) SetFetchedAt(v string) {
	o.FetchedAt = &v
}

func (o KpiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["service_uuid"] = o.ServiceUuid
	}
	if true {
		toSerialize["kpis"] = o.Kpis
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.FetchedAt != nil {
		toSerialize["fetched_at"] = o.FetchedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKpiResponse struct {
	value *KpiResponse
	isSet bool
}

func (v NullableKpiResponse) Get() *KpiResponse {
	return v.value
}

func (v *NullableKpiResponse) Set(val *KpiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKpiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKpiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKpiResponse(val *KpiResponse) *NullableKpiResponse {
	return &NullableKpiResponse{value: val, isSet: true}
}

func (v NullableKpiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKpiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


