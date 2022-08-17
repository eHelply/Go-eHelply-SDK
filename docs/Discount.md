# Discount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Rate** | **int32** |  | 
**Flat** | **int32** |  | 

## Methods

### NewDiscount

`func NewDiscount(name string, description string, rate int32, flat int32, ) *Discount`

NewDiscount instantiates a new Discount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountWithDefaults

`func NewDiscountWithDefaults() *Discount`

NewDiscountWithDefaults instantiates a new Discount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Discount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Discount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Discount) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Discount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Discount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Discount) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRate

`func (o *Discount) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Discount) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Discount) SetRate(v int32)`

SetRate sets Rate field to given value.


### GetFlat

`func (o *Discount) GetFlat() int32`

GetFlat returns the Flat field if non-nil, zero value otherwise.

### GetFlatOk

`func (o *Discount) GetFlatOk() (*int32, bool)`

GetFlatOk returns a tuple with the Flat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlat

`func (o *Discount) SetFlat(v int32)`

SetFlat sets Flat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


