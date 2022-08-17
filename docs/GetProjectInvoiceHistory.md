# GetProjectInvoiceHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectUuid** | **string** |  | 
**InvoiceHistory** | Pointer to [**[]History**](History.md) |  | [optional] [default to []]

## Methods

### NewGetProjectInvoiceHistory

`func NewGetProjectInvoiceHistory(projectUuid string, ) *GetProjectInvoiceHistory`

NewGetProjectInvoiceHistory instantiates a new GetProjectInvoiceHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectInvoiceHistoryWithDefaults

`func NewGetProjectInvoiceHistoryWithDefaults() *GetProjectInvoiceHistory`

NewGetProjectInvoiceHistoryWithDefaults instantiates a new GetProjectInvoiceHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectUuid

`func (o *GetProjectInvoiceHistory) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *GetProjectInvoiceHistory) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *GetProjectInvoiceHistory) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.


### GetInvoiceHistory

`func (o *GetProjectInvoiceHistory) GetInvoiceHistory() []History`

GetInvoiceHistory returns the InvoiceHistory field if non-nil, zero value otherwise.

### GetInvoiceHistoryOk

`func (o *GetProjectInvoiceHistory) GetInvoiceHistoryOk() (*[]History, bool)`

GetInvoiceHistoryOk returns a tuple with the InvoiceHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceHistory

`func (o *GetProjectInvoiceHistory) SetInvoiceHistory(v []History)`

SetInvoiceHistory sets InvoiceHistory field to given value.

### HasInvoiceHistory

`func (o *GetProjectInvoiceHistory) HasInvoiceHistory() bool`

HasInvoiceHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


