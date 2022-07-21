# GetProjectCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**Secrets** | [**[]GetSecret**](GetSecret.md) |  | 

## Methods

### NewGetProjectCredential

`func NewGetProjectCredential(serviceName string, secrets []GetSecret, ) *GetProjectCredential`

NewGetProjectCredential instantiates a new GetProjectCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectCredentialWithDefaults

`func NewGetProjectCredentialWithDefaults() *GetProjectCredential`

NewGetProjectCredentialWithDefaults instantiates a new GetProjectCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *GetProjectCredential) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *GetProjectCredential) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *GetProjectCredential) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetSecrets

`func (o *GetProjectCredential) GetSecrets() []GetSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *GetProjectCredential) GetSecretsOk() (*[]GetSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *GetProjectCredential) SetSecrets(v []GetSecret)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


