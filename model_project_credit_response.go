/*
eHelply SDK - 1.1.109

eHelply SDK for SuperStack Services

API version: 1.1.109
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ProjectCreditResponse struct for ProjectCreditResponse
type ProjectCreditResponse struct {
	Uuid string `json:"uuid"`
	ProjectUuid string `json:"project_uuid"`
	CreditsGranted int32 `json:"credits_granted"`
	CreditsConsumed *int32 `json:"credits_consumed,omitempty"`
	GrantedBy string `json:"granted_by"`
	GrantedReason string `json:"granted_reason"`
	GrantedAt string `json:"granted_at"`
	FullyConsumedAt *string `json:"fully_consumed_at,omitempty"`
	RevokedReason *string `json:"revoked_reason,omitempty"`
	RevokedAt *string `json:"revoked_at,omitempty"`
}

// NewProjectCreditResponse instantiates a new ProjectCreditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreditResponse(uuid string, projectUuid string, creditsGranted int32, grantedBy string, grantedReason string, grantedAt string) *ProjectCreditResponse {
	this := ProjectCreditResponse{}
	this.Uuid = uuid
	this.ProjectUuid = projectUuid
	this.CreditsGranted = creditsGranted
	this.GrantedBy = grantedBy
	this.GrantedReason = grantedReason
	this.GrantedAt = grantedAt
	return &this
}

// NewProjectCreditResponseWithDefaults instantiates a new ProjectCreditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreditResponseWithDefaults() *ProjectCreditResponse {
	this := ProjectCreditResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ProjectCreditResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ProjectCreditResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetProjectUuid returns the ProjectUuid field value
func (o *ProjectCreditResponse) GetProjectUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUuid, true
}

// SetProjectUuid sets field value
func (o *ProjectCreditResponse) SetProjectUuid(v string) {
	o.ProjectUuid = v
}

// GetCreditsGranted returns the CreditsGranted field value
func (o *ProjectCreditResponse) GetCreditsGranted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreditsGranted
}

// GetCreditsGrantedOk returns a tuple with the CreditsGranted field value
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetCreditsGrantedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsGranted, true
}

// SetCreditsGranted sets field value
func (o *ProjectCreditResponse) SetCreditsGranted(v int32) {
	o.CreditsGranted = v
}

// GetCreditsConsumed returns the CreditsConsumed field value if set, zero value otherwise.
func (o *ProjectCreditResponse) GetCreditsConsumed() int32 {
	if o == nil || o.CreditsConsumed == nil {
		var ret int32
		return ret
	}
	return *o.CreditsConsumed
}

// GetCreditsConsumedOk returns a tuple with the CreditsConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetCreditsConsumedOk() (*int32, bool) {
	if o == nil || o.CreditsConsumed == nil {
		return nil, false
	}
	return o.CreditsConsumed, true
}

// HasCreditsConsumed returns a boolean if a field has been set.
func (o *ProjectCreditResponse) HasCreditsConsumed() bool {
	if o != nil && o.CreditsConsumed != nil {
		return true
	}

	return false
}

// SetCreditsConsumed gets a reference to the given int32 and assigns it to the CreditsConsumed field.
func (o *ProjectCreditResponse) SetCreditsConsumed(v int32) {
	o.CreditsConsumed = &v
}

// GetGrantedBy returns the GrantedBy field value
func (o *ProjectCreditResponse) GetGrantedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantedBy
}

// GetGrantedByOk returns a tuple with the GrantedBy field value
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetGrantedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantedBy, true
}

// SetGrantedBy sets field value
func (o *ProjectCreditResponse) SetGrantedBy(v string) {
	o.GrantedBy = v
}

// GetGrantedReason returns the GrantedReason field value
func (o *ProjectCreditResponse) GetGrantedReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantedReason
}

// GetGrantedReasonOk returns a tuple with the GrantedReason field value
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetGrantedReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantedReason, true
}

// SetGrantedReason sets field value
func (o *ProjectCreditResponse) SetGrantedReason(v string) {
	o.GrantedReason = v
}

// GetGrantedAt returns the GrantedAt field value
func (o *ProjectCreditResponse) GetGrantedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantedAt
}

// GetGrantedAtOk returns a tuple with the GrantedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetGrantedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantedAt, true
}

// SetGrantedAt sets field value
func (o *ProjectCreditResponse) SetGrantedAt(v string) {
	o.GrantedAt = v
}

// GetFullyConsumedAt returns the FullyConsumedAt field value if set, zero value otherwise.
func (o *ProjectCreditResponse) GetFullyConsumedAt() string {
	if o == nil || o.FullyConsumedAt == nil {
		var ret string
		return ret
	}
	return *o.FullyConsumedAt
}

// GetFullyConsumedAtOk returns a tuple with the FullyConsumedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetFullyConsumedAtOk() (*string, bool) {
	if o == nil || o.FullyConsumedAt == nil {
		return nil, false
	}
	return o.FullyConsumedAt, true
}

// HasFullyConsumedAt returns a boolean if a field has been set.
func (o *ProjectCreditResponse) HasFullyConsumedAt() bool {
	if o != nil && o.FullyConsumedAt != nil {
		return true
	}

	return false
}

// SetFullyConsumedAt gets a reference to the given string and assigns it to the FullyConsumedAt field.
func (o *ProjectCreditResponse) SetFullyConsumedAt(v string) {
	o.FullyConsumedAt = &v
}

// GetRevokedReason returns the RevokedReason field value if set, zero value otherwise.
func (o *ProjectCreditResponse) GetRevokedReason() string {
	if o == nil || o.RevokedReason == nil {
		var ret string
		return ret
	}
	return *o.RevokedReason
}

// GetRevokedReasonOk returns a tuple with the RevokedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetRevokedReasonOk() (*string, bool) {
	if o == nil || o.RevokedReason == nil {
		return nil, false
	}
	return o.RevokedReason, true
}

// HasRevokedReason returns a boolean if a field has been set.
func (o *ProjectCreditResponse) HasRevokedReason() bool {
	if o != nil && o.RevokedReason != nil {
		return true
	}

	return false
}

// SetRevokedReason gets a reference to the given string and assigns it to the RevokedReason field.
func (o *ProjectCreditResponse) SetRevokedReason(v string) {
	o.RevokedReason = &v
}

// GetRevokedAt returns the RevokedAt field value if set, zero value otherwise.
func (o *ProjectCreditResponse) GetRevokedAt() string {
	if o == nil || o.RevokedAt == nil {
		var ret string
		return ret
	}
	return *o.RevokedAt
}

// GetRevokedAtOk returns a tuple with the RevokedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreditResponse) GetRevokedAtOk() (*string, bool) {
	if o == nil || o.RevokedAt == nil {
		return nil, false
	}
	return o.RevokedAt, true
}

// HasRevokedAt returns a boolean if a field has been set.
func (o *ProjectCreditResponse) HasRevokedAt() bool {
	if o != nil && o.RevokedAt != nil {
		return true
	}

	return false
}

// SetRevokedAt gets a reference to the given string and assigns it to the RevokedAt field.
func (o *ProjectCreditResponse) SetRevokedAt(v string) {
	o.RevokedAt = &v
}

func (o ProjectCreditResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	if true {
		toSerialize["credits_granted"] = o.CreditsGranted
	}
	if o.CreditsConsumed != nil {
		toSerialize["credits_consumed"] = o.CreditsConsumed
	}
	if true {
		toSerialize["granted_by"] = o.GrantedBy
	}
	if true {
		toSerialize["granted_reason"] = o.GrantedReason
	}
	if true {
		toSerialize["granted_at"] = o.GrantedAt
	}
	if o.FullyConsumedAt != nil {
		toSerialize["fully_consumed_at"] = o.FullyConsumedAt
	}
	if o.RevokedReason != nil {
		toSerialize["revoked_reason"] = o.RevokedReason
	}
	if o.RevokedAt != nil {
		toSerialize["revoked_at"] = o.RevokedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProjectCreditResponse struct {
	value *ProjectCreditResponse
	isSet bool
}

func (v NullableProjectCreditResponse) Get() *ProjectCreditResponse {
	return v.value
}

func (v *NullableProjectCreditResponse) Set(val *ProjectCreditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreditResponse(val *ProjectCreditResponse) *NullableProjectCreditResponse {
	return &NullableProjectCreditResponse{value: val, isSet: true}
}

func (v NullableProjectCreditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

