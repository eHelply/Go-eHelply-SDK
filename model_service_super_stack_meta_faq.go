/*
eHelply SDK - 1.1.122

eHelply SDK for SuperStack Services

API version: 1.1.122
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceSuperStackMetaFaq struct for ServiceSuperStackMetaFaq
type ServiceSuperStackMetaFaq struct {
	Question string `json:"question"`
	Answer string `json:"answer"`
}

// NewServiceSuperStackMetaFaq instantiates a new ServiceSuperStackMetaFaq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSuperStackMetaFaq(question string, answer string) *ServiceSuperStackMetaFaq {
	this := ServiceSuperStackMetaFaq{}
	this.Question = question
	this.Answer = answer
	return &this
}

// NewServiceSuperStackMetaFaqWithDefaults instantiates a new ServiceSuperStackMetaFaq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSuperStackMetaFaqWithDefaults() *ServiceSuperStackMetaFaq {
	this := ServiceSuperStackMetaFaq{}
	return &this
}

// GetQuestion returns the Question field value
func (o *ServiceSuperStackMetaFaq) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaFaq) GetQuestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *ServiceSuperStackMetaFaq) SetQuestion(v string) {
	o.Question = v
}

// GetAnswer returns the Answer field value
func (o *ServiceSuperStackMetaFaq) GetAnswer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMetaFaq) GetAnswerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Answer, true
}

// SetAnswer sets field value
func (o *ServiceSuperStackMetaFaq) SetAnswer(v string) {
	o.Answer = v
}

func (o ServiceSuperStackMetaFaq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["question"] = o.Question
	}
	if true {
		toSerialize["answer"] = o.Answer
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSuperStackMetaFaq struct {
	value *ServiceSuperStackMetaFaq
	isSet bool
}

func (v NullableServiceSuperStackMetaFaq) Get() *ServiceSuperStackMetaFaq {
	return v.value
}

func (v *NullableServiceSuperStackMetaFaq) Set(val *ServiceSuperStackMetaFaq) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSuperStackMetaFaq) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSuperStackMetaFaq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSuperStackMetaFaq(val *ServiceSuperStackMetaFaq) *NullableServiceSuperStackMetaFaq {
	return &NullableServiceSuperStackMetaFaq{value: val, isSet: true}
}

func (v NullableServiceSuperStackMetaFaq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSuperStackMetaFaq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


