/*
eHelply SDK - 1.1.106

eHelply SDK for SuperStack Services

API version: 1.1.106
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// CreateReview struct for CreateReview
type CreateReview struct {
	Rating int32 `json:"rating"`
	MaxRating int32 `json:"max_rating"`
	ReviewText string `json:"review_text"`
}

// NewCreateReview instantiates a new CreateReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReview(rating int32, maxRating int32, reviewText string) *CreateReview {
	this := CreateReview{}
	this.Rating = rating
	this.MaxRating = maxRating
	this.ReviewText = reviewText
	return &this
}

// NewCreateReviewWithDefaults instantiates a new CreateReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReviewWithDefaults() *CreateReview {
	this := CreateReview{}
	return &this
}

// GetRating returns the Rating field value
func (o *CreateReview) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *CreateReview) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *CreateReview) SetRating(v int32) {
	o.Rating = v
}

// GetMaxRating returns the MaxRating field value
func (o *CreateReview) GetMaxRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxRating
}

// GetMaxRatingOk returns a tuple with the MaxRating field value
// and a boolean to check if the value has been set.
func (o *CreateReview) GetMaxRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRating, true
}

// SetMaxRating sets field value
func (o *CreateReview) SetMaxRating(v int32) {
	o.MaxRating = v
}

// GetReviewText returns the ReviewText field value
func (o *CreateReview) GetReviewText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewText
}

// GetReviewTextOk returns a tuple with the ReviewText field value
// and a boolean to check if the value has been set.
func (o *CreateReview) GetReviewTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewText, true
}

// SetReviewText sets field value
func (o *CreateReview) SetReviewText(v string) {
	o.ReviewText = v
}

func (o CreateReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rating"] = o.Rating
	}
	if true {
		toSerialize["max_rating"] = o.MaxRating
	}
	if true {
		toSerialize["review_text"] = o.ReviewText
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReview struct {
	value *CreateReview
	isSet bool
}

func (v NullableCreateReview) Get() *CreateReview {
	return v.value
}

func (v *NullableCreateReview) Set(val *CreateReview) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReview) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReview(val *CreateReview) *NullableCreateReview {
	return &NullableCreateReview{value: val, isSet: true}
}

func (v NullableCreateReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


