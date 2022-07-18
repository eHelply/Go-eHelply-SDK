# CreateReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | **int32** |  | 
**MaxRating** | **int32** |  | 
**ReviewText** | **string** |  | 

## Methods

### NewCreateReview

`func NewCreateReview(rating int32, maxRating int32, reviewText string, ) *CreateReview`

NewCreateReview instantiates a new CreateReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReviewWithDefaults

`func NewCreateReviewWithDefaults() *CreateReview`

NewCreateReviewWithDefaults instantiates a new CreateReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *CreateReview) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CreateReview) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CreateReview) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetMaxRating

`func (o *CreateReview) GetMaxRating() int32`

GetMaxRating returns the MaxRating field if non-nil, zero value otherwise.

### GetMaxRatingOk

`func (o *CreateReview) GetMaxRatingOk() (*int32, bool)`

GetMaxRatingOk returns a tuple with the MaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRating

`func (o *CreateReview) SetMaxRating(v int32)`

SetMaxRating sets MaxRating field to given value.


### GetReviewText

`func (o *CreateReview) GetReviewText() string`

GetReviewText returns the ReviewText field if non-nil, zero value otherwise.

### GetReviewTextOk

`func (o *CreateReview) GetReviewTextOk() (*string, bool)`

GetReviewTextOk returns a tuple with the ReviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewText

`func (o *CreateReview) SetReviewText(v string)`

SetReviewText sets ReviewText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


