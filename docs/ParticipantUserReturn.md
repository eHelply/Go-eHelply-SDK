# ParticipantUserReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**UserUuid** | Pointer to **string** |  | [optional] 
**ParticipantMeta** | Pointer to **map[string]interface{}** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to [**Email**](Email.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**GpsLocation** | Pointer to **map[string]interface{}** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **time.Time** |  | [optional] 
**VerifiedLegalTerms** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**DateUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewParticipantUserReturn

`func NewParticipantUserReturn() *ParticipantUserReturn`

NewParticipantUserReturn instantiates a new ParticipantUserReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantUserReturnWithDefaults

`func NewParticipantUserReturnWithDefaults() *ParticipantUserReturn`

NewParticipantUserReturnWithDefaults instantiates a new ParticipantUserReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ParticipantUserReturn) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ParticipantUserReturn) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ParticipantUserReturn) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ParticipantUserReturn) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUserUuid

`func (o *ParticipantUserReturn) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *ParticipantUserReturn) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *ParticipantUserReturn) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *ParticipantUserReturn) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

### GetParticipantMeta

`func (o *ParticipantUserReturn) GetParticipantMeta() map[string]interface{}`

GetParticipantMeta returns the ParticipantMeta field if non-nil, zero value otherwise.

### GetParticipantMetaOk

`func (o *ParticipantUserReturn) GetParticipantMetaOk() (*map[string]interface{}, bool)`

GetParticipantMetaOk returns a tuple with the ParticipantMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantMeta

`func (o *ParticipantUserReturn) SetParticipantMeta(v map[string]interface{})`

SetParticipantMeta sets ParticipantMeta field to given value.

### HasParticipantMeta

`func (o *ParticipantUserReturn) HasParticipantMeta() bool`

HasParticipantMeta returns a boolean if a field has been set.

### GetFirstName

`func (o *ParticipantUserReturn) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ParticipantUserReturn) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ParticipantUserReturn) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ParticipantUserReturn) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ParticipantUserReturn) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ParticipantUserReturn) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ParticipantUserReturn) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ParticipantUserReturn) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *ParticipantUserReturn) GetEmail() Email`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ParticipantUserReturn) GetEmailOk() (*Email, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ParticipantUserReturn) SetEmail(v Email)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ParticipantUserReturn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ParticipantUserReturn) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ParticipantUserReturn) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ParticipantUserReturn) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ParticipantUserReturn) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountry

`func (o *ParticipantUserReturn) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ParticipantUserReturn) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ParticipantUserReturn) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ParticipantUserReturn) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetGpsLocation

`func (o *ParticipantUserReturn) GetGpsLocation() map[string]interface{}`

GetGpsLocation returns the GpsLocation field if non-nil, zero value otherwise.

### GetGpsLocationOk

`func (o *ParticipantUserReturn) GetGpsLocationOk() (*map[string]interface{}, bool)`

GetGpsLocationOk returns a tuple with the GpsLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsLocation

`func (o *ParticipantUserReturn) SetGpsLocation(v map[string]interface{})`

SetGpsLocation sets GpsLocation field to given value.

### HasGpsLocation

`func (o *ParticipantUserReturn) HasGpsLocation() bool`

HasGpsLocation returns a boolean if a field has been set.

### GetPicture

`func (o *ParticipantUserReturn) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *ParticipantUserReturn) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *ParticipantUserReturn) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *ParticipantUserReturn) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetLastLogin

`func (o *ParticipantUserReturn) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *ParticipantUserReturn) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *ParticipantUserReturn) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *ParticipantUserReturn) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetVerifiedLegalTerms

`func (o *ParticipantUserReturn) GetVerifiedLegalTerms() bool`

GetVerifiedLegalTerms returns the VerifiedLegalTerms field if non-nil, zero value otherwise.

### GetVerifiedLegalTermsOk

`func (o *ParticipantUserReturn) GetVerifiedLegalTermsOk() (*bool, bool)`

GetVerifiedLegalTermsOk returns a tuple with the VerifiedLegalTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedLegalTerms

`func (o *ParticipantUserReturn) SetVerifiedLegalTerms(v bool)`

SetVerifiedLegalTerms sets VerifiedLegalTerms field to given value.

### HasVerifiedLegalTerms

`func (o *ParticipantUserReturn) HasVerifiedLegalTerms() bool`

HasVerifiedLegalTerms returns a boolean if a field has been set.

### GetDateCreated

`func (o *ParticipantUserReturn) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ParticipantUserReturn) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ParticipantUserReturn) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ParticipantUserReturn) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ParticipantUserReturn) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ParticipantUserReturn) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ParticipantUserReturn) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ParticipantUserReturn) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


