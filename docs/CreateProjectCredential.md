# CreateProjectCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**Secrets** | [**[]Credential**](Credential.md) |  | 

## Methods

### NewCreateProjectCredential

`func NewCreateProjectCredential(serviceName string, secrets []Credential, ) *CreateProjectCredential`

NewCreateProjectCredential instantiates a new CreateProjectCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectCredentialWithDefaults

`func NewCreateProjectCredentialWithDefaults() *CreateProjectCredential`

NewCreateProjectCredentialWithDefaults instantiates a new CreateProjectCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *CreateProjectCredential) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CreateProjectCredential) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CreateProjectCredential) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetSecrets

`func (o *CreateProjectCredential) GetSecrets() []Credential`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *CreateProjectCredential) GetSecretsOk() (*[]Credential, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *CreateProjectCredential) SetSecrets(v []Credential)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


