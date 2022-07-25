/*
eHelply SDK - 1.1.93

eHelply SDK for SuperStack Services

API version: 1.1.93
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// UpdateReview struct for UpdateReview
type UpdateReview struct {
	Rating *int32 `json:"rating,omitempty"`
	MaxRating *int32 `json:"max_rating,omitempty"`
	ReviewText *string `json:"review_text,omitempty"`
}

// NewUpdateReview instantiates a new UpdateReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReview() *UpdateReview {
	this := UpdateReview{}
	return &this
}

// NewUpdateReviewWithDefaults instantiates a new UpdateReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReviewWithDefaults() *UpdateReview {
	this := UpdateReview{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *UpdateReview) GetRating() int32 {
	if o == nil || o.Rating == nil {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReview) GetRatingOk() (*int32, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *UpdateReview) HasRating() bool {
	if o != nil && o.Rating != nil {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *UpdateReview) SetRating(v int32) {
	o.Rating = &v
}

// GetMaxRating returns the MaxRating field value if set, zero value otherwise.
func (o *UpdateReview) GetMaxRating() int32 {
	if o == nil || o.MaxRating == nil {
		var ret int32
		return ret
	}
	return *o.MaxRating
}

// GetMaxRatingOk returns a tuple with the MaxRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReview) GetMaxRatingOk() (*int32, bool) {
	if o == nil || o.MaxRating == nil {
		return nil, false
	}
	return o.MaxRating, true
}

// HasMaxRating returns a boolean if a field has been set.
func (o *UpdateReview) HasMaxRating() bool {
	if o != nil && o.MaxRating != nil {
		return true
	}

	return false
}

// SetMaxRating gets a reference to the given int32 and assigns it to the MaxRating field.
func (o *UpdateReview) SetMaxRating(v int32) {
	o.MaxRating = &v
}

// GetReviewText returns the ReviewText field value if set, zero value otherwise.
func (o *UpdateReview) GetReviewText() string {
	if o == nil || o.ReviewText == nil {
		var ret string
		return ret
	}
	return *o.ReviewText
}

// GetReviewTextOk returns a tuple with the ReviewText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReview) GetReviewTextOk() (*string, bool) {
	if o == nil || o.ReviewText == nil {
		return nil, false
	}
	return o.ReviewText, true
}

// HasReviewText returns a boolean if a field has been set.
func (o *UpdateReview) HasReviewText() bool {
	if o != nil && o.ReviewText != nil {
		return true
	}

	return false
}

// SetReviewText gets a reference to the given string and assigns it to the ReviewText field.
func (o *UpdateReview) SetReviewText(v string) {
	o.ReviewText = &v
}

func (o UpdateReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rating != nil {
		toSerialize["rating"] = o.Rating
	}
	if o.MaxRating != nil {
		toSerialize["max_rating"] = o.MaxRating
	}
	if o.ReviewText != nil {
		toSerialize["review_text"] = o.ReviewText
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateReview struct {
	value *UpdateReview
	isSet bool
}

func (v NullableUpdateReview) Get() *UpdateReview {
	return v.value
}

func (v *NullableUpdateReview) Set(val *UpdateReview) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReview) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReview(val *UpdateReview) *NullableUpdateReview {
	return &NullableUpdateReview{value: val, isSet: true}
}

func (v NullableUpdateReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


