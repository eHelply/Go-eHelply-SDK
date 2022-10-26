/*
eHelply SDK - 1.1.112

eHelply SDK for SuperStack Services

API version: 1.1.112
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// AppointmentBase struct for AppointmentBase
type AppointmentBase struct {
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
}

// NewAppointmentBase instantiates a new AppointmentBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentBase() *AppointmentBase {
	this := AppointmentBase{}
	return &this
}

// NewAppointmentBaseWithDefaults instantiates a new AppointmentBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentBaseWithDefaults() *AppointmentBase {
	this := AppointmentBase{}
	return &this
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *AppointmentBase) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *AppointmentBase) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *AppointmentBase) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

// GetPlaceUuid returns the PlaceUuid field value if set, zero value otherwise.
func (o *AppointmentBase) GetPlaceUuid() string {
	if o == nil || o.PlaceUuid == nil {
		var ret string
		return ret
	}
	return *o.PlaceUuid
}

// GetPlaceUuidOk returns a tuple with the PlaceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetPlaceUuidOk() (*string, bool) {
	if o == nil || o.PlaceUuid == nil {
		return nil, false
	}
	return o.PlaceUuid, true
}

// HasPlaceUuid returns a boolean if a field has been set.
func (o *AppointmentBase) HasPlaceUuid() bool {
	if o != nil && o.PlaceUuid != nil {
		return true
	}

	return false
}

// SetPlaceUuid gets a reference to the given string and assigns it to the PlaceUuid field.
func (o *AppointmentBase) SetPlaceUuid(v string) {
	o.PlaceUuid = &v
}

// GetReviewGroupUuid returns the ReviewGroupUuid field value if set, zero value otherwise.
func (o *AppointmentBase) GetReviewGroupUuid() string {
	if o == nil || o.ReviewGroupUuid == nil {
		var ret string
		return ret
	}
	return *o.ReviewGroupUuid
}

// GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetReviewGroupUuidOk() (*string, bool) {
	if o == nil || o.ReviewGroupUuid == nil {
		return nil, false
	}
	return o.ReviewGroupUuid, true
}

// HasReviewGroupUuid returns a boolean if a field has been set.
func (o *AppointmentBase) HasReviewGroupUuid() bool {
	if o != nil && o.ReviewGroupUuid != nil {
		return true
	}

	return false
}

// SetReviewGroupUuid gets a reference to the given string and assigns it to the ReviewGroupUuid field.
func (o *AppointmentBase) SetReviewGroupUuid(v string) {
	o.ReviewGroupUuid = &v
}

// GetExpectedFinishAt returns the ExpectedFinishAt field value if set, zero value otherwise.
func (o *AppointmentBase) GetExpectedFinishAt() string {
	if o == nil || o.ExpectedFinishAt == nil {
		var ret string
		return ret
	}
	return *o.ExpectedFinishAt
}

// GetExpectedFinishAtOk returns a tuple with the ExpectedFinishAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetExpectedFinishAtOk() (*string, bool) {
	if o == nil || o.ExpectedFinishAt == nil {
		return nil, false
	}
	return o.ExpectedFinishAt, true
}

// HasExpectedFinishAt returns a boolean if a field has been set.
func (o *AppointmentBase) HasExpectedFinishAt() bool {
	if o != nil && o.ExpectedFinishAt != nil {
		return true
	}

	return false
}

// SetExpectedFinishAt gets a reference to the given string and assigns it to the ExpectedFinishAt field.
func (o *AppointmentBase) SetExpectedFinishAt(v string) {
	o.ExpectedFinishAt = &v
}

// GetExpectedStartAt returns the ExpectedStartAt field value if set, zero value otherwise.
func (o *AppointmentBase) GetExpectedStartAt() string {
	if o == nil || o.ExpectedStartAt == nil {
		var ret string
		return ret
	}
	return *o.ExpectedStartAt
}

// GetExpectedStartAtOk returns a tuple with the ExpectedStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetExpectedStartAtOk() (*string, bool) {
	if o == nil || o.ExpectedStartAt == nil {
		return nil, false
	}
	return o.ExpectedStartAt, true
}

// HasExpectedStartAt returns a boolean if a field has been set.
func (o *AppointmentBase) HasExpectedStartAt() bool {
	if o != nil && o.ExpectedStartAt != nil {
		return true
	}

	return false
}

// SetExpectedStartAt gets a reference to the given string and assigns it to the ExpectedStartAt field.
func (o *AppointmentBase) SetExpectedStartAt(v string) {
	o.ExpectedStartAt = &v
}

// GetActualStartAt returns the ActualStartAt field value if set, zero value otherwise.
func (o *AppointmentBase) GetActualStartAt() string {
	if o == nil || o.ActualStartAt == nil {
		var ret string
		return ret
	}
	return *o.ActualStartAt
}

// GetActualStartAtOk returns a tuple with the ActualStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetActualStartAtOk() (*string, bool) {
	if o == nil || o.ActualStartAt == nil {
		return nil, false
	}
	return o.ActualStartAt, true
}

// HasActualStartAt returns a boolean if a field has been set.
func (o *AppointmentBase) HasActualStartAt() bool {
	if o != nil && o.ActualStartAt != nil {
		return true
	}

	return false
}

// SetActualStartAt gets a reference to the given string and assigns it to the ActualStartAt field.
func (o *AppointmentBase) SetActualStartAt(v string) {
	o.ActualStartAt = &v
}

// GetActualFinishAt returns the ActualFinishAt field value if set, zero value otherwise.
func (o *AppointmentBase) GetActualFinishAt() string {
	if o == nil || o.ActualFinishAt == nil {
		var ret string
		return ret
	}
	return *o.ActualFinishAt
}

// GetActualFinishAtOk returns a tuple with the ActualFinishAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetActualFinishAtOk() (*string, bool) {
	if o == nil || o.ActualFinishAt == nil {
		return nil, false
	}
	return o.ActualFinishAt, true
}

// HasActualFinishAt returns a boolean if a field has been set.
func (o *AppointmentBase) HasActualFinishAt() bool {
	if o != nil && o.ActualFinishAt != nil {
		return true
	}

	return false
}

// SetActualFinishAt gets a reference to the given string and assigns it to the ActualFinishAt field.
func (o *AppointmentBase) SetActualFinishAt(v string) {
	o.ActualFinishAt = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *AppointmentBase) GetProducts() map[string]interface{} {
	if o == nil || o.Products == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetProductsOk() (map[string]interface{}, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *AppointmentBase) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given map[string]interface{} and assigns it to the Products field.
func (o *AppointmentBase) SetProducts(v map[string]interface{}) {
	o.Products = v
}

// GetMetaUuid returns the MetaUuid field value if set, zero value otherwise.
func (o *AppointmentBase) GetMetaUuid() string {
	if o == nil || o.MetaUuid == nil {
		var ret string
		return ret
	}
	return *o.MetaUuid
}

// GetMetaUuidOk returns a tuple with the MetaUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetMetaUuidOk() (*string, bool) {
	if o == nil || o.MetaUuid == nil {
		return nil, false
	}
	return o.MetaUuid, true
}

// HasMetaUuid returns a boolean if a field has been set.
func (o *AppointmentBase) HasMetaUuid() bool {
	if o != nil && o.MetaUuid != nil {
		return true
	}

	return false
}

// SetMetaUuid gets a reference to the given string and assigns it to the MetaUuid field.
func (o *AppointmentBase) SetMetaUuid(v string) {
	o.MetaUuid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppointmentBase) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppointmentBase) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppointmentBase) SetStatus(v string) {
	o.Status = &v
}

// GetCancellationAt returns the CancellationAt field value if set, zero value otherwise.
func (o *AppointmentBase) GetCancellationAt() string {
	if o == nil || o.CancellationAt == nil {
		var ret string
		return ret
	}
	return *o.CancellationAt
}

// GetCancellationAtOk returns a tuple with the CancellationAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetCancellationAtOk() (*string, bool) {
	if o == nil || o.CancellationAt == nil {
		return nil, false
	}
	return o.CancellationAt, true
}

// HasCancellationAt returns a boolean if a field has been set.
func (o *AppointmentBase) HasCancellationAt() bool {
	if o != nil && o.CancellationAt != nil {
		return true
	}

	return false
}

// SetCancellationAt gets a reference to the given string and assigns it to the CancellationAt field.
func (o *AppointmentBase) SetCancellationAt(v string) {
	o.CancellationAt = &v
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise.
func (o *AppointmentBase) GetCancellationReason() string {
	if o == nil || o.CancellationReason == nil {
		var ret string
		return ret
	}
	return *o.CancellationReason
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetCancellationReasonOk() (*string, bool) {
	if o == nil || o.CancellationReason == nil {
		return nil, false
	}
	return o.CancellationReason, true
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *AppointmentBase) HasCancellationReason() bool {
	if o != nil && o.CancellationReason != nil {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given string and assigns it to the CancellationReason field.
func (o *AppointmentBase) SetCancellationReason(v string) {
	o.CancellationReason = &v
}

// GetCancellationEntityUuid returns the CancellationEntityUuid field value if set, zero value otherwise.
func (o *AppointmentBase) GetCancellationEntityUuid() string {
	if o == nil || o.CancellationEntityUuid == nil {
		var ret string
		return ret
	}
	return *o.CancellationEntityUuid
}

// GetCancellationEntityUuidOk returns a tuple with the CancellationEntityUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentBase) GetCancellationEntityUuidOk() (*string, bool) {
	if o == nil || o.CancellationEntityUuid == nil {
		return nil, false
	}
	return o.CancellationEntityUuid, true
}

// HasCancellationEntityUuid returns a boolean if a field has been set.
func (o *AppointmentBase) HasCancellationEntityUuid() bool {
	if o != nil && o.CancellationEntityUuid != nil {
		return true
	}

	return false
}

// SetCancellationEntityUuid gets a reference to the given string and assigns it to the CancellationEntityUuid field.
func (o *AppointmentBase) SetCancellationEntityUuid(v string) {
	o.CancellationEntityUuid = &v
}

func (o AppointmentBase) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableAppointmentBase struct {
	value *AppointmentBase
	isSet bool
}

func (v NullableAppointmentBase) Get() *AppointmentBase {
	return v.value
}

func (v *NullableAppointmentBase) Set(val *AppointmentBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentBase(val *AppointmentBase) *NullableAppointmentBase {
	return &NullableAppointmentBase{value: val, isSet: true}
}

func (v NullableAppointmentBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


