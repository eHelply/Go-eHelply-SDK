# PaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**Last4Digits** | **string** |  | 
**CardBrand** | **string** |  | 
**ProjectUuid** | **string** |  | 

## Methods

### NewPaymentMethodResponse

`func NewPaymentMethodResponse(paymentId string, last4Digits string, cardBrand string, projectUuid string, ) *PaymentMethodResponse`

NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodResponseWithDefaults

`func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse`

NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentMethodResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentMethodResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentMethodResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetLast4Digits

`func (o *PaymentMethodResponse) GetLast4Digits() string`

GetLast4Digits returns the Last4Digits field if non-nil, zero value otherwise.

### GetLast4DigitsOk

`func (o *PaymentMethodResponse) GetLast4DigitsOk() (*string, bool)`

GetLast4DigitsOk returns a tuple with the Last4Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4Digits

`func (o *PaymentMethodResponse) SetLast4Digits(v string)`

SetLast4Digits sets Last4Digits field to given value.


### GetCardBrand

`func (o *PaymentMethodResponse) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *PaymentMethodResponse) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *PaymentMethodResponse) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.


### GetProjectUuid

`func (o *PaymentMethodResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *PaymentMethodResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *PaymentMethodResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


