/*
eHelply SDK - 1.1.87

eHelply SDK for SuperStack Services

API version: 1.1.87
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceCreate struct for ServiceCreate
type ServiceCreate struct {
	Name string `json:"name"`
	Key string `json:"key"`
	Version string `json:"version"`
	Summary string `json:"summary"`
	Authors []string `json:"authors"`
	AuthorEmails []string `json:"author_emails"`
}

// NewServiceCreate instantiates a new ServiceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCreate(name string, key string, version string, summary string, authors []string, authorEmails []string) *ServiceCreate {
	this := ServiceCreate{}
	this.Name = name
	this.Key = key
	this.Version = version
	this.Summary = summary
	this.Authors = authors
	this.AuthorEmails = authorEmails
	return &this
}

// NewServiceCreateWithDefaults instantiates a new ServiceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCreateWithDefaults() *ServiceCreate {
	this := ServiceCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ServiceCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceCreate) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *ServiceCreate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ServiceCreate) SetKey(v string) {
	o.Key = v
}

// GetVersion returns the Version field value
func (o *ServiceCreate) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ServiceCreate) SetVersion(v string) {
	o.Version = v
}

// GetSummary returns the Summary field value
func (o *ServiceCreate) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ServiceCreate) SetSummary(v string) {
	o.Summary = v
}

// GetAuthors returns the Authors field value
func (o *ServiceCreate) GetAuthors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetAuthorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authors, true
}

// SetAuthors sets field value
func (o *ServiceCreate) SetAuthors(v []string) {
	o.Authors = v
}

// GetAuthorEmails returns the AuthorEmails field value
func (o *ServiceCreate) GetAuthorEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorEmails
}

// GetAuthorEmailsOk returns a tuple with the AuthorEmails field value
// and a boolean to check if the value has been set.
func (o *ServiceCreate) GetAuthorEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorEmails, true
}

// SetAuthorEmails sets field value
func (o *ServiceCreate) SetAuthorEmails(v []string) {
	o.AuthorEmails = v
}

func (o ServiceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["authors"] = o.Authors
	}
	if true {
		toSerialize["author_emails"] = o.AuthorEmails
	}
	return json.Marshal(toSerialize)
}

type NullableServiceCreate struct {
	value *ServiceCreate
	isSet bool
}

func (v NullableServiceCreate) Get() *ServiceCreate {
	return v.value
}

func (v *NullableServiceCreate) Set(val *ServiceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceCreate(val *ServiceCreate) *NullableServiceCreate {
	return &NullableServiceCreate{value: val, isSet: true}
}

func (v NullableServiceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


