# SecurityEncryptionKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Key** | ***os.File** |  | 
**Category** | **string** |  | 
**DeletedAt** | **string** |  | 
**CreatedAt** | **string** |  | 
**RetrievedAt** | **string** |  | 

## Methods

### NewSecurityEncryptionKeyResponse

`func NewSecurityEncryptionKeyResponse(uuid string, key *os.File, category string, deletedAt string, createdAt string, retrievedAt string, ) *SecurityEncryptionKeyResponse`

NewSecurityEncryptionKeyResponse instantiates a new SecurityEncryptionKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEncryptionKeyResponseWithDefaults

`func NewSecurityEncryptionKeyResponseWithDefaults() *SecurityEncryptionKeyResponse`

NewSecurityEncryptionKeyResponseWithDefaults instantiates a new SecurityEncryptionKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SecurityEncryptionKeyResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SecurityEncryptionKeyResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SecurityEncryptionKeyResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetKey

`func (o *SecurityEncryptionKeyResponse) GetKey() *os.File`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SecurityEncryptionKeyResponse) GetKeyOk() (**os.File, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SecurityEncryptionKeyResponse) SetKey(v *os.File)`

SetKey sets Key field to given value.


### GetCategory

`func (o *SecurityEncryptionKeyResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SecurityEncryptionKeyResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SecurityEncryptionKeyResponse) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDeletedAt

`func (o *SecurityEncryptionKeyResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *SecurityEncryptionKeyResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *SecurityEncryptionKeyResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.


### GetCreatedAt

`func (o *SecurityEncryptionKeyResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityEncryptionKeyResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityEncryptionKeyResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetRetrievedAt

`func (o *SecurityEncryptionKeyResponse) GetRetrievedAt() string`

GetRetrievedAt returns the RetrievedAt field if non-nil, zero value otherwise.

### GetRetrievedAtOk

`func (o *SecurityEncryptionKeyResponse) GetRetrievedAtOk() (*string, bool)`

GetRetrievedAtOk returns a tuple with the RetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedAt

`func (o *SecurityEncryptionKeyResponse) SetRetrievedAt(v string)`

SetRetrievedAt sets RetrievedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


