/*
eHelply SDK - 1.1.119

eHelply SDK for SuperStack Services

API version: 1.1.119
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// TagResponse **:param** uuid                                **type:** string **:param** name                                **type:** string **:param** project_uuid                        **type:** string or None
type TagResponse struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

// NewTagResponse instantiates a new TagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagResponse(uuid string, name string) *TagResponse {
	this := TagResponse{}
	this.Uuid = uuid
	this.Name = name
	return &this
}

// NewTagResponseWithDefaults instantiates a new TagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagResponseWithDefaults() *TagResponse {
	this := TagResponse{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *TagResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *TagResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *TagResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *TagResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagResponse) SetName(v string) {
	o.Name = v
}

// GetProjectUuid returns the ProjectUuid field value if set, zero value otherwise.
func (o *TagResponse) GetProjectUuid() string {
	if o == nil || o.ProjectUuid == nil {
		var ret string
		return ret
	}
	return *o.ProjectUuid
}

// GetProjectUuidOk returns a tuple with the ProjectUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponse) GetProjectUuidOk() (*string, bool) {
	if o == nil || o.ProjectUuid == nil {
		return nil, false
	}
	return o.ProjectUuid, true
}

// HasProjectUuid returns a boolean if a field has been set.
func (o *TagResponse) HasProjectUuid() bool {
	if o != nil && o.ProjectUuid != nil {
		return true
	}

	return false
}

// SetProjectUuid gets a reference to the given string and assigns it to the ProjectUuid field.
func (o *TagResponse) SetProjectUuid(v string) {
	o.ProjectUuid = &v
}

func (o TagResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ProjectUuid != nil {
		toSerialize["project_uuid"] = o.ProjectUuid
	}
	return json.Marshal(toSerialize)
}

type NullableTagResponse struct {
	value *TagResponse
	isSet bool
}

func (v NullableTagResponse) Get() *TagResponse {
	return v.value
}

func (v *NullableTagResponse) Set(val *TagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagResponse(val *TagResponse) *NullableTagResponse {
	return &NullableTagResponse{value: val, isSet: true}
}

func (v NullableTagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


