# UserEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUserEmail

`func NewUserEmail() *UserEmail`

NewUserEmail instantiates a new UserEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEmailWithDefaults

`func NewUserEmailWithDefaults() *UserEmail`

NewUserEmailWithDefaults instantiates a new UserEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UserEmail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UserEmail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UserEmail) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UserEmail) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVerified

`func (o *UserEmail) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserEmail) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserEmail) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserEmail) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


