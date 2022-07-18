# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CognitoId** | Pointer to **string** |  | [optional] 
**CognitoData** | Pointer to **map[string]interface{}** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**GpsLocation** | Pointer to **map[string]interface{}** |  | [optional] 
**VerifiedLegalTerms** | Pointer to **bool** |  | [optional] [default to false]
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**DateUpdated** | Pointer to **time.Time** |  | [optional] 
**DateDeleted** | Pointer to **time.Time** |  | [optional] 
**Uuid** | **string** |  | 
**LastLogin** | Pointer to **time.Time** |  | [optional] 
**FirstLogin** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewUser

`func NewUser(uuid string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCognitoId

`func (o *User) GetCognitoId() string`

GetCognitoId returns the CognitoId field if non-nil, zero value otherwise.

### GetCognitoIdOk

`func (o *User) GetCognitoIdOk() (*string, bool)`

GetCognitoIdOk returns a tuple with the CognitoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitoId

`func (o *User) SetCognitoId(v string)`

SetCognitoId sets CognitoId field to given value.

### HasCognitoId

`func (o *User) HasCognitoId() bool`

HasCognitoId returns a boolean if a field has been set.

### GetCognitoData

`func (o *User) GetCognitoData() map[string]interface{}`

GetCognitoData returns the CognitoData field if non-nil, zero value otherwise.

### GetCognitoDataOk

`func (o *User) GetCognitoDataOk() (*map[string]interface{}, bool)`

GetCognitoDataOk returns a tuple with the CognitoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitoData

`func (o *User) SetCognitoData(v map[string]interface{})`

SetCognitoData sets CognitoData field to given value.

### HasCognitoData

`func (o *User) HasCognitoData() bool`

HasCognitoData returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *User) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *User) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *User) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *User) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountry

`func (o *User) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *User) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *User) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *User) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPicture

`func (o *User) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *User) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *User) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *User) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetGpsLocation

`func (o *User) GetGpsLocation() map[string]interface{}`

GetGpsLocation returns the GpsLocation field if non-nil, zero value otherwise.

### GetGpsLocationOk

`func (o *User) GetGpsLocationOk() (*map[string]interface{}, bool)`

GetGpsLocationOk returns a tuple with the GpsLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsLocation

`func (o *User) SetGpsLocation(v map[string]interface{})`

SetGpsLocation sets GpsLocation field to given value.

### HasGpsLocation

`func (o *User) HasGpsLocation() bool`

HasGpsLocation returns a boolean if a field has been set.

### GetVerifiedLegalTerms

`func (o *User) GetVerifiedLegalTerms() bool`

GetVerifiedLegalTerms returns the VerifiedLegalTerms field if non-nil, zero value otherwise.

### GetVerifiedLegalTermsOk

`func (o *User) GetVerifiedLegalTermsOk() (*bool, bool)`

GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedLegalTerms

`func (o *User) SetVerifiedLegalTerms(v bool)`

SetVerifiedLegalTerms sets VerifiedLegalTerms field to given value.

### HasVerifiedLegalTerms

`func (o *User) HasVerifiedLegalTerms() bool`

HasVerifiedLegalTerms returns a boolean if a field has been set.

### GetDateCreated

`func (o *User) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *User) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *User) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *User) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *User) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *User) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *User) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *User) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDateDeleted

`func (o *User) GetDateDeleted() time.Time`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *User) GetDateDeletedOk() (*time.Time, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *User) SetDateDeleted(v time.Time)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *User) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetUuid

`func (o *User) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *User) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *User) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetFirstLogin

`func (o *User) GetFirstLogin() bool`

GetFirstLogin returns the FirstLogin field if non-nil, zero value otherwise.

### GetFirstLoginOk

`func (o *User) GetFirstLoginOk() (*bool, bool)`

GetFirstLoginOk returns a tuple with the FirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLogin

`func (o *User) SetFirstLogin(v bool)`

SetFirstLogin sets FirstLogin field to given value.

### HasFirstLogin

`func (o *User) HasFirstLogin() bool`

HasFirstLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


