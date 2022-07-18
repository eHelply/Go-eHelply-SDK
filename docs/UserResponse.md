# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**GpsLocation** | Pointer to **map[string]interface{}** |  | [optional] 
**VerifiedLegalTerms** | Pointer to **bool** |  | [optional] [default to false]
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**DateUpdated** | Pointer to **time.Time** |  | [optional] 
**DateDeleted** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to [**UserEmail**](UserEmail.md) |  | [optional] 
**Missing** | Pointer to **[]string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Participants** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Flags** | Pointer to [**UserFlags**](UserFlags.md) |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 

## Methods

### NewUserResponse

`func NewUserResponse() *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UserResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountry

`func (o *UserResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPicture

`func (o *UserResponse) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *UserResponse) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *UserResponse) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *UserResponse) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetGpsLocation

`func (o *UserResponse) GetGpsLocation() map[string]interface{}`

GetGpsLocation returns the GpsLocation field if non-nil, zero value otherwise.

### GetGpsLocationOk

`func (o *UserResponse) GetGpsLocationOk() (*map[string]interface{}, bool)`

GetGpsLocationOk returns a tuple with the GpsLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsLocation

`func (o *UserResponse) SetGpsLocation(v map[string]interface{})`

SetGpsLocation sets GpsLocation field to given value.

### HasGpsLocation

`func (o *UserResponse) HasGpsLocation() bool`

HasGpsLocation returns a boolean if a field has been set.

### GetVerifiedLegalTerms

`func (o *UserResponse) GetVerifiedLegalTerms() bool`

GetVerifiedLegalTerms returns the VerifiedLegalTerms field if non-nil, zero value otherwise.

### GetVerifiedLegalTermsOk

`func (o *UserResponse) GetVerifiedLegalTermsOk() (*bool, bool)`

GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedLegalTerms

`func (o *UserResponse) SetVerifiedLegalTerms(v bool)`

SetVerifiedLegalTerms sets VerifiedLegalTerms field to given value.

### HasVerifiedLegalTerms

`func (o *UserResponse) HasVerifiedLegalTerms() bool`

HasVerifiedLegalTerms returns a boolean if a field has been set.

### GetDateCreated

`func (o *UserResponse) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UserResponse) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UserResponse) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UserResponse) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *UserResponse) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *UserResponse) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *UserResponse) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *UserResponse) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDateDeleted

`func (o *UserResponse) GetDateDeleted() time.Time`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *UserResponse) GetDateDeletedOk() (*time.Time, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *UserResponse) SetDateDeleted(v time.Time)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *UserResponse) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### GetEmail

`func (o *UserResponse) GetEmail() UserEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResponse) GetEmailOk() (*UserEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResponse) SetEmail(v UserEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMissing

`func (o *UserResponse) GetMissing() []string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *UserResponse) GetMissingOk() (*[]string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *UserResponse) SetMissing(v []string)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *UserResponse) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetUuid

`func (o *UserResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetParticipants

`func (o *UserResponse) GetParticipants() []map[string]interface{}`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *UserResponse) GetParticipantsOk() (*[]map[string]interface{}, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *UserResponse) SetParticipants(v []map[string]interface{})`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *UserResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetFlags

`func (o *UserResponse) GetFlags() UserFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *UserResponse) GetFlagsOk() (*UserFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *UserResponse) SetFlags(v UserFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *UserResponse) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserResponse) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserResponse) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserResponse) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserResponse) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


