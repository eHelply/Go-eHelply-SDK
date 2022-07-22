/*
eHelply SDK - 1.1.97

eHelply SDK for SuperStack Services

API version: 1.1.97
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AppointmentResponse struct for AppointmentResponse
type AppointmentResponse struct {
	ProjectUuid *string `json:"project_uuid,omitempty"`
	PlaceUuid *string `json:"place_uuid,omitempty"`
	ReviewGroupUuid *string `json:"review_group_uuid,omitempty"`
	ExpectedFinishAt *string `json:"expected_finish_at,omitempty"`
	ExpectedStartAt *string `json:"expected_start_at,omitempty"`
	ActualStartAt *string `json:"actual_start_at,omitempty"`
	ActualFinishAt *string `json:"actual_finish_at,omitempty"`
	Products map[string]interface{} `json:"products,omitempty"`
	MetaUuid *string `json:"meta_uuid,omitempty"`
	Status *string `json:"status,omitempty"`
	CancellationAt *string `json:"cancellation_at,omitempty"`
	CancellationReason *string `json:"cancellation_reason,omitempty"`
	CancellationEntityUuid *string `json:"cancellation_entity_uuid,omitempty"`
	Uuid string `json:"uuid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

// NewAppointmentResponse instantiates a new AppointmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentResponse(uuid string, createdAt string, updatedAt string) *AppointmentResponse {
	this := AppointmentResponse{}
	this.Uuid = uuid
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAppointmentResponseWithDefaults instantiates a new AppointmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentResponseWithDefaults() *AppointmentResponse {
	this := AppointmentResponse{}
	return &this
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *AppointmentResponse) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *AppointmentResponse) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *AppointmentResponse) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetPlaceUuid returns the PlaceUuid field value if set, zero value otherwise.
func (o *AppointmentResponse) GetPlaceUuid() string {
	if o == nil || o.PlaceUuid == nil {
		var ret string
		return ret
	}
	return *o.PlaceUuid
}

// GetPlaceUuidOk returns a tuple with the PlaceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetPlaceUuidOk() (*string, bool) {
	if o == nil || o.PlaceUuid == nil {
		return nil, false
	}
	return o.PlaceUuid, true
}

// HasPlaceUuid returns a boolean if a field has been set.
func (o *AppointmentResponse) HasPlaceUuid() bool {
	if o != nil && o.PlaceUuid != nil {
		return true
	}

	return false
}

// SetPlaceUuid gets a reference to the given string and assigns it to the PlaceUuid field.
func (o *AppointmentResponse) SetPlaceUuid(v string) {
	o.PlaceUuid = &v
}

// GetReviewGroupUuid returns the ReviewGroupUuid field value if set, zero value otherwise.
func (o *AppointmentResponse) GetReviewGroupUuid() string {
	if o == nil || o.ReviewGroupUuid == nil {
		var ret string
		return ret
	}
	return *o.ReviewGroupUuid
}

// GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetReviewGroupUuidOk() (*string, bool) {
	if o == nil || o.ReviewGroupUuid == nil {
		return nil, false
	}
	return o.ReviewGroupUuid, true
}

// HasReviewGroupUuid returns a boolean if a field has been set.
func (o *AppointmentResponse) HasReviewGroupUuid() bool {
	if o != nil && o.ReviewGroupUuid != nil {
		return true
	}

	return false
}

// SetReviewGroupUuid gets a reference to the given string and assigns it to the ReviewGroupUuid field.
func (o *AppointmentResponse) SetReviewGroupUuid(v string) {
	o.ReviewGroupUuid = &v
}

// GetExpectedFinishAt returns the ExpectedFinishAt field value if set, zero value otherwise.
func (o *AppointmentResponse) GetExpectedFinishAt() string {
	if o == nil || o.ExpectedFinishAt == nil {
		var ret string
		return ret
	}
	return *o.ExpectedFinishAt
}

// GetExpectedFinishAtOk returns a tuple with the ExpectedFinishAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetExpectedFinishAtOk() (*string, bool) {
	if o == nil || o.ExpectedFinishAt == nil {
		return nil, false
	}
	return o.ExpectedFinishAt, true
}

// HasExpectedFinishAt returns a boolean if a field has been set.
func (o *AppointmentResponse) HasExpectedFinishAt() bool {
	if o != nil && o.ExpectedFinishAt != nil {
		return true
	}

	return false
}

// SetExpectedFinishAt gets a reference to the given string and assigns it to the ExpectedFinishAt field.
func (o *AppointmentResponse) SetExpectedFinishAt(v string) {
	o.ExpectedFinishAt = &v
}

// GetExpectedStartAt returns the ExpectedStartAt field value if set, zero value otherwise.
func (o *AppointmentResponse) GetExpectedStartAt() string {
	if o == nil || o.ExpectedStartAt == nil {
		var ret string
		return ret
	}
	return *o.ExpectedStartAt
}

// GetExpectedStartAtOk returns a tuple with the ExpectedStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetExpectedStartAtOk() (*string, bool) {
	if o == nil || o.ExpectedStartAt == nil {
		return nil, false
	}
	return o.ExpectedStartAt, true
}

// HasExpectedStartAt returns a boolean if a field has been set.
func (o *AppointmentResponse) HasExpectedStartAt() bool {
	if o != nil && o.ExpectedStartAt != nil {
		return true
	}

	return false
}

// SetExpectedStartAt gets a reference to the given string and assigns it to the ExpectedStartAt field.
func (o *AppointmentResponse) SetExpectedStartAt(v string) {
	o.ExpectedStartAt = &v
}

// GetActualStartAt returns the ActualStartAt field value if set, zero value otherwise.
func (o *AppointmentResponse) GetActualStartAt() string {
	if o == nil || o.ActualStartAt == nil {
		var ret string
		return ret
	}
	return *o.ActualStartAt
}

// GetActualStartAtOk returns a tuple with the ActualStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetActualStartAtOk() (*string, bool) {
	if o == nil || o.ActualStartAt == nil {
		return nil, false
	}
	return o.ActualStartAt, true
}

// HasActualStartAt returns a boolean if a field has been set.
func (o *AppointmentResponse) HasActualStartAt() bool {
	if o != nil && o.ActualStartAt != nil {
		return true
	}

	return false
}

// SetActualStartAt gets a reference to the given string and assigns it to the ActualStartAt field.
func (o *AppointmentResponse) SetActualStartAt(v string) {
	o.ActualStartAt = &v
}

// GetActualFinishAt returns the ActualFinishAt field value if set, zero value otherwise.
func (o *AppointmentResponse) GetActualFinishAt() string {
	if o == nil || o.ActualFinishAt == nil {
		var ret string
		return ret
	}
	return *o.ActualFinishAt
}

// GetActualFinishAtOk returns a tuple with the ActualFinishAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetActualFinishAtOk() (*string, bool) {
	if o == nil || o.ActualFinishAt == nil {
		return nil, false
	}
	return o.ActualFinishAt, true
}

// HasActualFinishAt returns a boolean if a field has been set.
func (o *AppointmentResponse) HasActualFinishAt() bool {
	if o != nil && o.ActualFinishAt != nil {
		return true
	}

	return false
}

// SetActualFinishAt gets a reference to the given string and assigns it to the ActualFinishAt field.
func (o *AppointmentResponse) SetActualFinishAt(v string) {
	o.ActualFinishAt = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *AppointmentResponse) GetProducts() map[string]interface{} {
	if o == nil || o.Products == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetProductsOk() (map[string]interface{}, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *AppointmentResponse) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given map[string]interface{} and assigns it to the Products field.
func (o *AppointmentResponse) SetProducts(v map[string]interface{}) {
	o.Products = v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *AppointmentResponse) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *AppointmentResponse) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *AppointmentResponse) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppointmentResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppointmentResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppointmentResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCancellationAt returns the CancellationAt field value if set, zero value otherwise.
func (o *AppointmentResponse) GetCancellationAt() string {
	if o == nil || o.CancellationAt == nil {
		var ret string
		return ret
	}
	return *o.CancellationAt
}

// GetCancellationAtOk returns a tuple with the CancellationAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetCancellationAtOk() (*string, bool) {
	if o == nil || o.CancellationAt == nil {
		return nil, false
	}
	return o.CancellationAt, true
}

// HasCancellationAt returns a boolean if a field has been set.
func (o *AppointmentResponse) HasCancellationAt() bool {
	if o != nil && o.CancellationAt != nil {
		return true
	}

	return false
}

// SetCancellationAt gets a reference to the given string and assigns it to the CancellationAt field.
func (o *AppointmentResponse) SetCancellationAt(v string) {
	o.CancellationAt = &v
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *AppointmentResponse) GetCancellationReason() string {
	if o == nil || o.CancellationReason == nil {
		var ret string
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetCancellationReasonOk() (*string, bool) {
	if o == nil || o.CancellationReason == nil {
		return nil, false
	}
	return o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *AppointmentResponse) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given string and assigns it to the CancellationReason field.
func (o *AppointmentResponse) SetCancellationReason(v string) {
	o.CancellationReason = &v
}

// GetCancellationEntityUuid returns the CancellationEntityUuid field value if set, zero value otherwise.
func (o *AppointmentResponse) GetCancellationEntityUuid() string {
	if o == nil || o.CancellationEntityUuid == nil {
		var ret string
		return ret
	}
	return *o.CancellationEntityUuid
}

// GetCancellationEntityUuidOk returns a tuple with the CancellationEntityUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetCancellationEntityUuidOk() (*string, bool) {
	if o == nil || o.CancellationEntityUuid == nil {
		return nil, false
	}
	return o.CancellationEntityUuid, true
}

// HasCancellationEntityUuid returns a boolean if a field has been set.
func (o *AppointmentResponse) HasCancellationEntityUuid() bool {
	if o != nil && o.CancellationEntityUuid != nil {
		return true
	}

	return false
}

// SetCancellationEntityUuid gets a reference to the given string and assigns it to the CancellationEntityUuid field.
func (o *AppointmentResponse) SetCancellationEntityUuid(v string) {
	o.CancellationEntityUuid = &v
}

// GetUuid returns the Uuid field value
func (o *AppointmentResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *AppointmentResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AppointmentResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AppointmentResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AppointmentResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AppointmentResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AppointmentResponse) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AppointmentResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *AppointmentResponse) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

func (o AppointmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if o.PlaceUuid != nil {
		toSerialize["place_uuid"] = o.PlaceUuid
	}
	if o.ReviewGroupUuid != nil {
		toSerialize["review_group_uuid"] = o.ReviewGroupUuid
	}
	if o.ExpectedFinishAt != nil {
		toSerialize["expected_finish_at"] = o.ExpectedFinishAt
	}
	if o.ExpectedStartAt != nil {
		toSerialize["expected_start_at"] = o.ExpectedStartAt
	}
	if o.ActualStartAt != nil {
		toSerialize["actual_start_at"] = o.ActualStartAt
	}
	if o.ActualFinishAt != nil {
		toSerialize["actual_finish_at"] = o.ActualFinishAt
	}
	if o.Products != nil {
		toSerialize["products"] = o.Products
	}
	if o.MetaUuid != nil {
		toSerialize["meta_uuid"] = o.MetaUuid
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CancellationAt != nil {
		toSerialize["cancellation_at"] = o.CancellationAt
	}
	if o.CancellationReason != nil {
		toSerialize["cancellation_reason"] = o.CancellationReason
	}
	if o.CancellationEntityUuid != nil {
		toSerialize["cancellation_entity_uuid"] = o.CancellationEntityUuid
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAppointmentResponse struct {
	value *AppointmentResponse
	isSet bool
}

func (v NullableAppointmentResponse) Get() *AppointmentResponse {
	return v.value
}

func (v *NullableAppointmentResponse) Set(val *AppointmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentResponse(val *AppointmentResponse) *NullableAppointmentResponse {
	return &NullableAppointmentResponse{value: val, isSet: true}
}

func (v NullableAppointmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


