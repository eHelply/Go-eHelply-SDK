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

// NoteMeta Metadata associated to a note
type NoteMeta struct {
	OriginalAuthor *string `json:"original_author,omitempty"`
	Author string `json:"author"`
	PreviousVersion *string `json:"previous_version,omitempty"`
	NextVersion *string `json:"next_version,omitempty"`
}

// NewNoteMeta instantiates a new NoteMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteMeta(author string) *NoteMeta {
	this := NoteMeta{}
	this.Author = author
	return &this
}

// NewNoteMetaWithDefaults instantiates a new NoteMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteMetaWithDefaults() *NoteMeta {
	this := NoteMeta{}
	return &this
}

// GetOriginalAuthor returns the OriginalAuthor field value if set, zero value otherwise.
func (o *NoteMeta) GetOriginalAuthor() string {
	if o == nil || o.OriginalAuthor == nil {
		var ret string
		return ret
	}
	return *o.OriginalAuthor
}

// GetOriginalAuthorOk returns a tuple with the OriginalAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteMeta) GetOriginalAuthorOk() (*string, bool) {
	if o == nil || o.OriginalAuthor == nil {
		return nil, false
	}
	return o.OriginalAuthor, true
}

// HasOriginalAuthor returns a boolean if a field has been set.
func (o *NoteMeta) HasOriginalAuthor() bool {
	if o != nil && o.OriginalAuthor != nil {
		return true
	}

	return false
}

// SetOriginalAuthor gets a reference to the given string and assigns it to the OriginalAuthor field.
func (o *NoteMeta) SetOriginalAuthor(v string) {
	o.OriginalAuthor = &v
}

// GetAuthor returns the Author field value
func (o *NoteMeta) GetAuthor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *NoteMeta) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *NoteMeta) SetAuthor(v string) {
	o.Author = v
}

// GetPreviousVersion returns the PreviousVersion field value if set, zero value otherwise.
func (o *NoteMeta) GetPreviousVersion() string {
	if o == nil || o.PreviousVersion == nil {
		var ret string
		return ret
	}
	return *o.PreviousVersion
}

// GetPreviousVersionOk returns a tuple with the PreviousVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteMeta) GetPreviousVersionOk() (*string, bool) {
	if o == nil || o.PreviousVersion == nil {
		return nil, false
	}
	return o.PreviousVersion, true
}

// HasPreviousVersion returns a boolean if a field has been set.
func (o *NoteMeta) HasPreviousVersion() bool {
	if o != nil && o.PreviousVersion != nil {
		return true
	}

	return false
}

// SetPreviousVersion gets a reference to the given string and assigns it to the PreviousVersion field.
func (o *NoteMeta) SetPreviousVersion(v string) {
	o.PreviousVersion = &v
}

// GetNextVersion returns the NextVersion field value if set, zero value otherwise.
func (o *NoteMeta) GetNextVersion() string {
	if o == nil || o.NextVersion == nil {
		var ret string
		return ret
	}
	return *o.NextVersion
}

// GetNextVersionOk returns a tuple with the NextVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteMeta) GetNextVersionOk() (*string, bool) {
	if o == nil || o.NextVersion == nil {
		return nil, false
	}
	return o.NextVersion, true
}

// HasNextVersion returns a boolean if a field has been set.
func (o *NoteMeta) HasNextVersion() bool {
	if o != nil && o.NextVersion != nil {
		return true
	}

	return false
}

// SetNextVersion gets a reference to the given string and assigns it to the NextVersion field.
func (o *NoteMeta) SetNextVersion(v string) {
	o.NextVersion = &v
}

func (o NoteMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OriginalAuthor != nil {
		toSerialize["original_author"] = o.OriginalAuthor
	}
	if true {
		toSerialize["author"] = o.Author
	}
	if o.PreviousVersion != nil {
		toSerialize["previous_version"] = o.PreviousVersion
	}
	if o.NextVersion != nil {
		toSerialize["next_version"] = o.NextVersion
	}
	return json.Marshal(toSerialize)
}

type NullableNoteMeta struct {
	value *NoteMeta
	isSet bool
}

func (v NullableNoteMeta) Get() *NoteMeta {
	return v.value
}

func (v *NullableNoteMeta) Set(val *NoteMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteMeta(val *NoteMeta) *NullableNoteMeta {
	return &NullableNoteMeta{value: val, isSet: true}
}

func (v NullableNoteMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

