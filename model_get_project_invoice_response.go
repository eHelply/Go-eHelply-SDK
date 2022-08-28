/*
eHelply SDK - 1.1.111

eHelply SDK for SuperStack Services

API version: 1.1.111
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// GetProjectInvoiceResponse struct for GetProjectInvoiceResponse
type GetProjectInvoiceResponse struct {
	ProjectUuid string `json:"project_uuid"`
	InvoiceUuid string `json:"invoice_uuid"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	CreatedBy string `json:"created_by"`
	Invoice *GetInvoiceResponse `json:"invoice,omitempty"`
}

// NewGetProjectInvoiceResponse instantiates a new GetProjectInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectInvoiceResponse(projectUuid string, invoiceUuid string, year int32, month int32, createdBy string) *GetProjectInvoiceResponse {
	this := GetProjectInvoiceResponse{}
	this.ProjectUuid = projectUuid
	this.InvoiceUuid = invoiceUuid
	this.Year = year
	this.Month = month
	this.CreatedBy = createdBy
	return &this
}

// NewGetProjectInvoiceResponseWithDefaults instantiates a new GetProjectInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectInvoiceResponseWithDefaults() *GetProjectInvoiceResponse {
	this := GetProjectInvoiceResponse{}
	return &this
}

// GetProjectUuid returns the ProjectUuid field value
func (o *GetProjectInvoiceResponse) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *GetProjectInvoiceResponse) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetInvoiceUuid returns the InvoiceUuid field value
func (o *GetProjectInvoiceResponse) GetInvoiceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceUuid
}

// GetInvoiceUuidOk returns a tuple with the InvoiceUuid field value
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceResponse) GetInvoiceUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceUuid, true
}

// SetInvoiceUuid sets field value
func (o *GetProjectInvoiceResponse) SetInvoiceUuid(v string) {
	o.InvoiceUuid = v
}

// GetYear returns the Year field value
func (o *GetProjectInvoiceResponse) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceResponse) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *GetProjectInvoiceResponse) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value
func (o *GetProjectInvoiceResponse) GetMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceResponse) GetMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *GetProjectInvoiceResponse) SetMonth(v int32) {
	o.Month = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *GetProjectInvoiceResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *GetProjectInvoiceResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *GetProjectInvoiceResponse) GetInvoice() GetInvoiceResponse {
	if o == nil || o.Invoice == nil {
		var ret GetInvoiceResponse
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectInvoiceResponse) GetInvoiceOk() (*GetInvoiceResponse, bool) {
	if o == nil || o.Invoice == nil {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *GetProjectInvoiceResponse) HasInvoice() bool {
	if o != nil && o.Invoice != nil {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given GetInvoiceResponse and assigns it to the Invoice field.
func (o *GetProjectInvoiceResponse) SetInvoice(v GetInvoiceResponse) {
	o.Invoice = &v
}

func (o GetProjectInvoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if true {
		toSerialize["invoice_uuid"] = o.InvoiceUuid
	}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["month"] = o.Month
	}
	if true {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Invoice != nil {
		toSerialize["invoice"] = o.Invoice
	}
	return json.Marshal(toSerialize)
}

type NullableGetProjectInvoiceResponse struct {
	value *GetProjectInvoiceResponse
	isSet bool
}

func (v NullableGetProjectInvoiceResponse) Get() *GetProjectInvoiceResponse {
	return v.value
}

func (v *NullableGetProjectInvoiceResponse) Set(val *GetProjectInvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectInvoiceResponse(val *GetProjectInvoiceResponse) *NullableGetProjectInvoiceResponse {
	return &NullableGetProjectInvoiceResponse{value: val, isSet: true}
}

func (v NullableGetProjectInvoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


