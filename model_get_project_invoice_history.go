/*
eHelply SDK - 1.1.107

eHelply SDK for SuperStack Services

API version: 1.1.107
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// GetProjectInvoiceHistory struct for GetProjectInvoiceHistory
type GetProjectInvoiceHistory struct {
	ProjectUuid string `json:"project_uuid"`
	InvoiceHistory []History `json:"invoice_history,omitempty"`
}

// NewGetProjectInvoiceHistory instantiates a new GetProjectInvoiceHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectInvoiceHistory(projectUuid string) *GetProjectInvoiceHistory {
	this := GetProjectInvoiceHistory{}
	this.ProjectUuid = projectUuid
	return &this
}

// NewGetProjectInvoiceHistoryWithDefaults instantiates a new GetProjectInvoiceHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectInvoiceHistoryWithDefaults() *GetProjectInvoiceHistory {
	this := GetProjectInvoiceHistory{}
	return &this
}

// GetProjectUuid returns the ProjectUuid field value
func (o *GetProjectInvoiceHistory) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceHistory) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *GetProjectInvoiceHistory) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetInvoiceHistory returns the InvoiceHistory field value if set, zero value otherwise.
func (o *GetProjectInvoiceHistory) GetInvoiceHistory() []History {
	if o == nil || o.InvoiceHistory == nil {
		var ret []History
		return ret
	}
	return o.InvoiceHistory
}

// GetInvoiceHistoryOk returns a tuple with the InvoiceHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceHistory) GetInvoiceHistoryOk() ([]History, bool) {
	if o == nil || o.InvoiceHistory == nil {
		return nil, false
	}
	return o.InvoiceHistory, true
}

// HasInvoiceHistory returns a boolean if a field has been set.
func (o *GetProjectInvoiceHistory) HasInvoiceHistory() bool {
	if o != nil && o.InvoiceHistory != nil {
		return true
	}

	return false
}

// SetInvoiceHistory gets a reference to the given []History and assigns it to the InvoiceHistory field.
func (o *GetProjectInvoiceHistory) SetInvoiceHistory(v []History) {
	o.InvoiceHistory = v
}

func (o GetProjectInvoiceHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.InvoiceHistory != nil {
		toSerialize["invoice_history"] = o.InvoiceHistory
	}
	return json.Marshal(toSerialize)
}

type NullableGetProjectInvoiceHistory struct {
	value *GetProjectInvoiceHistory
	isSet bool
}

func (v NullableGetProjectInvoiceHistory) Get() *GetProjectInvoiceHistory {
	return v.value
}

func (v *NullableGetProjectInvoiceHistory) Set(val *GetProjectInvoiceHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectInvoiceHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectInvoiceHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectInvoiceHistory(val *GetProjectInvoiceHistory) *NullableGetProjectInvoiceHistory {
	return &NullableGetProjectInvoiceHistory{value: val, isSet: true}
}

func (v NullableGetProjectInvoiceHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectInvoiceHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


