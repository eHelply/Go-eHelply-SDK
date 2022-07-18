# UserPasswordResetConfirmation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**ConfirmationCode** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewUserPasswordResetConfirmation

`func NewUserPasswordResetConfirmation(email string, confirmationCode string, password string, ) *UserPasswordResetConfirmation`

NewUserPasswordResetConfirmation instantiates a new UserPasswordResetConfirmation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordResetConfirmationWithDefaults

`func NewUserPasswordResetConfirmationWithDefaults() *UserPasswordResetConfirmation`

NewUserPasswordResetConfirmationWithDefaults instantiates a new UserPasswordResetConfirmation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserPasswordResetConfirmation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPasswordResetConfirmation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPasswordResetConfirmation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetConfirmationCode

`func (o *UserPasswordResetConfirmation) GetConfirmationCode() string`

GetConfirmationCode returns the ConfirmationCode field if non-nil, zero value otherwise.

### GetConfirmationCodeOk

`func (o *UserPasswordResetConfirmation) GetConfirmationCodeOk() (*string, bool)`

GetConfirmationCodeOk returns a tuple with the ConfirmationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationCode

`func (o *UserPasswordResetConfirmation) SetConfirmationCode(v string)`

SetConfirmationCode sets ConfirmationCode field to given value.


### GetPassword

`func (o *UserPasswordResetConfirmation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPasswordResetConfirmation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPasswordResetConfirmation) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


