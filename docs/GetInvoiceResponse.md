# GetInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceUuid** | **string** |  | 
**LineItems** | [**[]LineItem**](LineItem.md) |  | 
**Subtotal** | **int32** |  | 
**Discounts** | [**[]Discount**](Discount.md) |  | 
**Taxes** | [**[]Tax**](Tax.md) |  | 
**Total** | **int32** |  | 
**Notes** | [**[]Note**](Note.md) |  | 
**Paid** | Pointer to **bool** |  | [optional] [default to false]
**CreatedAt** | **string** |  | 
**UpdateAt** | **string** |  | 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGetInvoiceResponse

`func NewGetInvoiceResponse(invoiceUuid string, lineItems []LineItem, subtotal int32, discounts []Discount, taxes []Tax, total int32, notes []Note, createdAt string, updateAt string, ) *GetInvoiceResponse`

NewGetInvoiceResponse instantiates a new GetInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceResponseWithDefaults

`func NewGetInvoiceResponseWithDefaults() *GetInvoiceResponse`

NewGetInvoiceResponseWithDefaults instantiates a new GetInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceUuid

`func (o *GetInvoiceResponse) GetInvoiceUuid() string`

GetInvoiceUuid returns the InvoiceUuid field if non-nil, zero value otherwise.

### GetInvoiceUuidOk

`func (o *GetInvoiceResponse) GetInvoiceUuidOk() (*string, bool)`

GetInvoiceUuidOk returns a tuple with the InvoiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUuid

`func (o *GetInvoiceResponse) SetInvoiceUuid(v string)`

SetInvoiceUuid sets InvoiceUuid field to given value.


### GetLineItems

`func (o *GetInvoiceResponse) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *GetInvoiceResponse) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *GetInvoiceResponse) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.


### GetSubtotal

`func (o *GetInvoiceResponse) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *GetInvoiceResponse) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *GetInvoiceResponse) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.


### GetDiscounts

`func (o *GetInvoiceResponse) GetDiscounts() []Discount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *GetInvoiceResponse) GetDiscountsOk() (*[]Discount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *GetInvoiceResponse) SetDiscounts(v []Discount)`

SetDiscounts sets Discounts field to given value.


### GetTaxes

`func (o *GetInvoiceResponse) GetTaxes() []Tax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *GetInvoiceResponse) GetTaxesOk() (*[]Tax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *GetInvoiceResponse) SetTaxes(v []Tax)`

SetTaxes sets Taxes field to given value.


### GetTotal

`func (o *GetInvoiceResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInvoiceResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInvoiceResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetNotes

`func (o *GetInvoiceResponse) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetInvoiceResponse) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetInvoiceResponse) SetNotes(v []Note)`

SetNotes sets Notes field to given value.


### GetPaid

`func (o *GetInvoiceResponse) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *GetInvoiceResponse) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *GetInvoiceResponse) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *GetInvoiceResponse) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetInvoiceResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetInvoiceResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetInvoiceResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdateAt

`func (o *GetInvoiceResponse) GetUpdateAt() string`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *GetInvoiceResponse) GetUpdateAtOk() (*string, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *GetInvoiceResponse) SetUpdateAt(v string)`

SetUpdateAt sets UpdateAt field to given value.


### GetDeletedAt

`func (o *GetInvoiceResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GetInvoiceResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GetInvoiceResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GetInvoiceResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


