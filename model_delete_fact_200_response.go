/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// DeleteFact200Response struct for DeleteFact200Response
type DeleteFact200Response struct {
	Message *string `json:"message,omitempty"`
}

// NewDeleteFact200Response instantiates a new DeleteFact200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFact200Response() *DeleteFact200Response {
	this := DeleteFact200Response{}
	return &this
}

// NewDeleteFact200ResponseWithDefaults instantiates a new DeleteFact200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFact200ResponseWithDefaults() *DeleteFact200Response {
	this := DeleteFact200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DeleteFact200Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFact200Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DeleteFact200Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DeleteFact200Response) SetMessage(v string) {
	o.Message = &v
}

func (o DeleteFact200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteFact200Response struct {
	value *DeleteFact200Response
	isSet bool
}

func (v NullableDeleteFact200Response) Get() *DeleteFact200Response {
	return v.value
}

func (v *NullableDeleteFact200Response) Set(val *DeleteFact200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFact200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFact200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFact200Response(val *DeleteFact200Response) *NullableDeleteFact200Response {
	return &NullableDeleteFact200Response{value: val, isSet: true}
}

func (v NullableDeleteFact200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFact200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


