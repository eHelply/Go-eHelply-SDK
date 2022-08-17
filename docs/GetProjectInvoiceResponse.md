# GetProjectInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectUuid** | **string** |  | 
**InvoiceUuid** | **string** |  | 
**Year** | **int32** |  | 
**Month** | **int32** |  | 
**CreatedBy** | **string** |  | 
**Invoice** | Pointer to [**GetInvoiceResponse**](GetInvoiceResponse.md) |  | [optional] 

## Methods

### NewGetProjectInvoiceResponse

`func NewGetProjectInvoiceResponse(projectUuid string, invoiceUuid string, year int32, month int32, createdBy string, ) *GetProjectInvoiceResponse`

NewGetProjectInvoiceResponse instantiates a new GetProjectInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectInvoiceResponseWithDefaults

`func NewGetProjectInvoiceResponseWithDefaults() *GetProjectInvoiceResponse`

NewGetProjectInvoiceResponseWithDefaults instantiates a new GetProjectInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectUuid

`func (o *GetProjectInvoiceResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *GetProjectInvoiceResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *GetProjectInvoiceResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.


### GetInvoiceUuid

`func (o *GetProjectInvoiceResponse) GetInvoiceUuid() string`

GetInvoiceUuid returns the InvoiceUuid field if non-nil, zero value otherwise.

### GetInvoiceUuidOk

`func (o *GetProjectInvoiceResponse) GetInvoiceUuidOk() (*string, bool)`

GetInvoiceUuidOk returns a tuple with the InvoiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUuid

`func (o *GetProjectInvoiceResponse) SetInvoiceUuid(v string)`

SetInvoiceUuid sets InvoiceUuid field to given value.


### GetYear

`func (o *GetProjectInvoiceResponse) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetProjectInvoiceResponse) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetProjectInvoiceResponse) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *GetProjectInvoiceResponse) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *GetProjectInvoiceResponse) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *GetProjectInvoiceResponse) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetCreatedBy

`func (o *GetProjectInvoiceResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetProjectInvoiceResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetProjectInvoiceResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetInvoice

`func (o *GetProjectInvoiceResponse) GetInvoice() GetInvoiceResponse`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *GetProjectInvoiceResponse) GetInvoiceOk() (*GetInvoiceResponse, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *GetProjectInvoiceResponse) SetInvoice(v GetInvoiceResponse)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *GetProjectInvoiceResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


