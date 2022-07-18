# AttachPaymentToProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentType** | **string** |  | 
**Number** | **string** |  | 
**ExpMonth** | **int32** |  | 
**ExpYear** | **int32** |  | 
**Cvc** | **string** |  | 

## Methods

### NewAttachPaymentToProject

`func NewAttachPaymentToProject(paymentType string, number string, expMonth int32, expYear int32, cvc string, ) *AttachPaymentToProject`

NewAttachPaymentToProject instantiates a new AttachPaymentToProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachPaymentToProjectWithDefaults

`func NewAttachPaymentToProjectWithDefaults() *AttachPaymentToProject`

NewAttachPaymentToProjectWithDefaults instantiates a new AttachPaymentToProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentType

`func (o *AttachPaymentToProject) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *AttachPaymentToProject) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *AttachPaymentToProject) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetNumber

`func (o *AttachPaymentToProject) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AttachPaymentToProject) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AttachPaymentToProject) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetExpMonth

`func (o *AttachPaymentToProject) GetExpMonth() int32`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *AttachPaymentToProject) GetExpMonthOk() (*int32, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *AttachPaymentToProject) SetExpMonth(v int32)`

SetExpMonth sets ExpMonth field to given value.


### GetExpYear

`func (o *AttachPaymentToProject) GetExpYear() int32`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *AttachPaymentToProject) GetExpYearOk() (*int32, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *AttachPaymentToProject) SetExpYear(v int32)`

SetExpYear sets ExpYear field to given value.


### GetCvc

`func (o *AttachPaymentToProject) GetCvc() string`

GetCvc returns the Cvc field if non-nil, zero value otherwise.

### GetCvcOk

`func (o *AttachPaymentToProject) GetCvcOk() (*string, bool)`

GetCvcOk returns a tuple with the Cvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvc

`func (o *AttachPaymentToProject) SetCvc(v string)`

SetCvc sets Cvc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


