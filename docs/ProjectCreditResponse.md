# ProjectCreditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**ProjectUuid** | **string** |  | 
**CreditsGranted** | **int32** |  | 
**CreditsConsumed** | Pointer to **int32** |  | [optional] 
**GrantedBy** | **string** |  | 
**GrantedReason** | **string** |  | 
**GrantedAt** | **string** |  | 
**FullyConsumedAt** | Pointer to **string** |  | [optional] 
**RevokedReason** | Pointer to **string** |  | [optional] 
**RevokedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectCreditResponse

`func NewProjectCreditResponse(uuid string, projectUuid string, creditsGranted int32, grantedBy string, grantedReason string, grantedAt string, ) *ProjectCreditResponse`

NewProjectCreditResponse instantiates a new ProjectCreditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreditResponseWithDefaults

`func NewProjectCreditResponseWithDefaults() *ProjectCreditResponse`

NewProjectCreditResponseWithDefaults instantiates a new ProjectCreditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ProjectCreditResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectCreditResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectCreditResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetProjectUuid

`func (o *ProjectCreditResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *ProjectCreditResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *ProjectCreditResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.


### GetCreditsGranted

`func (o *ProjectCreditResponse) GetCreditsGranted() int32`

GetCreditsGranted returns the CreditsGranted field if non-nil, zero value otherwise.

### GetCreditsGrantedOk

`func (o *ProjectCreditResponse) GetCreditsGrantedOk() (*int32, bool)`

GetCreditsGrantedOk returns a tuple with the CreditsGranted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsGranted

`func (o *ProjectCreditResponse) SetCreditsGranted(v int32)`

SetCreditsGranted sets CreditsGranted field to given value.


### GetCreditsConsumed

`func (o *ProjectCreditResponse) GetCreditsConsumed() int32`

GetCreditsConsumed returns the CreditsConsumed field if non-nil, zero value otherwise.

### GetCreditsConsumedOk

`func (o *ProjectCreditResponse) GetCreditsConsumedOk() (*int32, bool)`

GetCreditsConsumedOk returns a tuple with the CreditsConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsConsumed

`func (o *ProjectCreditResponse) SetCreditsConsumed(v int32)`

SetCreditsConsumed sets CreditsConsumed field to given value.

### HasCreditsConsumed

`func (o *ProjectCreditResponse) HasCreditsConsumed() bool`

HasCreditsConsumed returns a boolean if a field has been set.

### GetGrantedBy

`func (o *ProjectCreditResponse) GetGrantedBy() string`

GetGrantedBy returns the GrantedBy field if non-nil, zero value otherwise.

### GetGrantedByOk

`func (o *ProjectCreditResponse) GetGrantedByOk() (*string, bool)`

GetGrantedByOk returns a tuple with the GrantedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedBy

`func (o *ProjectCreditResponse) SetGrantedBy(v string)`

SetGrantedBy sets GrantedBy field to given value.


### GetGrantedReason

`func (o *ProjectCreditResponse) GetGrantedReason() string`

GetGrantedReason returns the GrantedReason field if non-nil, zero value otherwise.

### GetGrantedReasonOk

`func (o *ProjectCreditResponse) GetGrantedReasonOk() (*string, bool)`

GetGrantedReasonOk returns a tuple with the GrantedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedReason

`func (o *ProjectCreditResponse) SetGrantedReason(v string)`

SetGrantedReason sets GrantedReason field to given value.


### GetGrantedAt

`func (o *ProjectCreditResponse) GetGrantedAt() string`

GetGrantedAt returns the GrantedAt field if non-nil, zero value otherwise.

### GetGrantedAtOk

`func (o *ProjectCreditResponse) GetGrantedAtOk() (*string, bool)`

GetGrantedAtOk returns a tuple with the GrantedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAt

`func (o *ProjectCreditResponse) SetGrantedAt(v string)`

SetGrantedAt sets GrantedAt field to given value.


### GetFullyConsumedAt

`func (o *ProjectCreditResponse) GetFullyConsumedAt() string`

GetFullyConsumedAt returns the FullyConsumedAt field if non-nil, zero value otherwise.

### GetFullyConsumedAtOk

`func (o *ProjectCreditResponse) GetFullyConsumedAtOk() (*string, bool)`

GetFullyConsumedAtOk returns a tuple with the FullyConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyConsumedAt

`func (o *ProjectCreditResponse) SetFullyConsumedAt(v string)`

SetFullyConsumedAt sets FullyConsumedAt field to given value.

### HasFullyConsumedAt

`func (o *ProjectCreditResponse) HasFullyConsumedAt() bool`

HasFullyConsumedAt returns a boolean if a field has been set.

### GetRevokedReason

`func (o *ProjectCreditResponse) GetRevokedReason() string`

GetRevokedReason returns the RevokedReason field if non-nil, zero value otherwise.

### GetRevokedReasonOk

`func (o *ProjectCreditResponse) GetRevokedReasonOk() (*string, bool)`

GetRevokedReasonOk returns a tuple with the RevokedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedReason

`func (o *ProjectCreditResponse) SetRevokedReason(v string)`

SetRevokedReason sets RevokedReason field to given value.

### HasRevokedReason

`func (o *ProjectCreditResponse) HasRevokedReason() bool`

HasRevokedReason returns a boolean if a field has been set.

### GetRevokedAt

`func (o *ProjectCreditResponse) GetRevokedAt() string`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *ProjectCreditResponse) GetRevokedAtOk() (*string, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *ProjectCreditResponse) SetRevokedAt(v string)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *ProjectCreditResponse) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


