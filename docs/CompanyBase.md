# CompanyBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] [default to true]
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Contact** | Pointer to [**ContactBase**](ContactBase.md) |  | [optional] 

## Methods

### NewCompanyBase

`func NewCompanyBase() *CompanyBase`

NewCompanyBase instantiates a new CompanyBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyBaseWithDefaults

`func NewCompanyBaseWithDefaults() *CompanyBase`

NewCompanyBaseWithDefaults instantiates a new CompanyBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompanyBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *CompanyBase) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CompanyBase) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CompanyBase) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CompanyBase) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPublic

`func (o *CompanyBase) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CompanyBase) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CompanyBase) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CompanyBase) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetMeta

`func (o *CompanyBase) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CompanyBase) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CompanyBase) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CompanyBase) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetContact

`func (o *CompanyBase) GetContact() ContactBase`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CompanyBase) GetContactOk() (*ContactBase, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CompanyBase) SetContact(v ContactBase)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CompanyBase) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


