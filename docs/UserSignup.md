# UserSignup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**PhoneNumber** | **string** |  | 
**Country** | **string** |  | 
**Lat** | Pointer to **string** |  | [optional] 
**Lng** | Pointer to **string** |  | [optional] 
**VerifiedLegalTerms** | Pointer to **bool** |  | [optional] [default to false]
**Picture** | Pointer to **string** |  | [optional] 
**Newsletters** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserSignup

`func NewUserSignup(username string, password string, email string, firstName string, lastName string, phoneNumber string, country string, ) *UserSignup`

NewUserSignup instantiates a new UserSignup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSignupWithDefaults

`func NewUserSignupWithDefaults() *UserSignup`

NewUserSignupWithDefaults instantiates a new UserSignup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserSignup) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSignup) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSignup) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserSignup) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserSignup) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserSignup) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *UserSignup) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSignup) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSignup) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UserSignup) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserSignup) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserSignup) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserSignup) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserSignup) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserSignup) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPhoneNumber

`func (o *UserSignup) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserSignup) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserSignup) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetCountry

`func (o *UserSignup) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserSignup) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserSignup) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetLat

`func (o *UserSignup) GetLat() string`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *UserSignup) GetLatOk() (*string, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *UserSignup) SetLat(v string)`

SetLat sets Lat field to given value.

### HasLat

`func (o *UserSignup) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *UserSignup) GetLng() string`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *UserSignup) GetLngOk() (*string, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *UserSignup) SetLng(v string)`

SetLng sets Lng field to given value.

### HasLng

`func (o *UserSignup) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetVerifiedLegalTerms

`func (o *UserSignup) GetVerifiedLegalTerms() bool`

GetVerifiedLegalTerms returns the VerifiedLegalTerms field if non-nil, zero value otherwise.

### GetVerifiedLegalTermsOk

`func (o *UserSignup) GetVerifiedLegalTermsOk() (*bool, bool)`

GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedLegalTerms

`func (o *UserSignup) SetVerifiedLegalTerms(v bool)`

SetVerifiedLegalTerms sets VerifiedLegalTerms field to given value.

### HasVerifiedLegalTerms

`func (o *UserSignup) HasVerifiedLegalTerms() bool`

HasVerifiedLegalTerms returns a boolean if a field has been set.

### GetPicture

`func (o *UserSignup) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *UserSignup) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *UserSignup) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *UserSignup) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetNewsletters

`func (o *UserSignup) GetNewsletters() bool`

GetNewsletters returns the Newsletters field if non-nil, zero value otherwise.

### GetNewslettersOk

`func (o *UserSignup) GetNewslettersOk() (*bool, bool)`

GetNewslettersOk returns a tuple with the Newsletters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsletters

`func (o *UserSignup) SetNewsletters(v bool)`

SetNewsletters sets Newsletters field to given value.

### HasNewsletters

`func (o *UserSignup) HasNewsletters() bool`

HasNewsletters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


