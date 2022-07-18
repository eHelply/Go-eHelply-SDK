# UserConfirmation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**VerificationCode** | **string** |  | 

## Methods

### NewUserConfirmation

`func NewUserConfirmation(email string, verificationCode string, ) *UserConfirmation`

NewUserConfirmation instantiates a new UserConfirmation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfirmationWithDefaults

`func NewUserConfirmationWithDefaults() *UserConfirmation`

NewUserConfirmationWithDefaults instantiates a new UserConfirmation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserConfirmation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserConfirmation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserConfirmation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetVerificationCode

`func (o *UserConfirmation) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *UserConfirmation) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *UserConfirmation) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


