# UpdateProjectCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | [**[]Credential**](Credential.md) |  | 

## Methods

### NewUpdateProjectCredentialRequest

`func NewUpdateProjectCredentialRequest(secrets []Credential, ) *UpdateProjectCredentialRequest`

NewUpdateProjectCredentialRequest instantiates a new UpdateProjectCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectCredentialRequestWithDefaults

`func NewUpdateProjectCredentialRequestWithDefaults() *UpdateProjectCredentialRequest`

NewUpdateProjectCredentialRequestWithDefaults instantiates a new UpdateProjectCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *UpdateProjectCredentialRequest) GetSecrets() []Credential`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *UpdateProjectCredentialRequest) GetSecretsOk() (*[]Credential, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *UpdateProjectCredentialRequest) SetSecrets(v []Credential)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


