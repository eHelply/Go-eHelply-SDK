# ContactBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phones** | Pointer to [**[]ContactMethod**](ContactMethod.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Socials** | Pointer to [**[]ContactMethod**](ContactMethod.md) |  | [optional] 

## Methods

### NewContactBase

`func NewContactBase() *ContactBase`

NewContactBase instantiates a new ContactBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactBaseWithDefaults

`func NewContactBaseWithDefaults() *ContactBase`

NewContactBaseWithDefaults instantiates a new ContactBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhones

`func (o *ContactBase) GetPhones() []ContactMethod`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *ContactBase) GetPhonesOk() (*[]ContactMethod, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *ContactBase) SetPhones(v []ContactMethod)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *ContactBase) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetEmail

`func (o *ContactBase) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactBase) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactBase) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactBase) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWebsite

`func (o *ContactBase) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ContactBase) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ContactBase) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ContactBase) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetSocials

`func (o *ContactBase) GetSocials() []ContactMethod`

GetSocials returns the Socials field if non-nil, zero value otherwise.

### GetSocialsOk

`func (o *ContactBase) GetSocialsOk() (*[]ContactMethod, bool)`

GetSocialsOk returns a tuple with the Socials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocials

`func (o *ContactBase) SetSocials(v []ContactMethod)`

SetSocials sets Socials field to given value.

### HasSocials

`func (o *ContactBase) HasSocials() bool`

HasSocials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


