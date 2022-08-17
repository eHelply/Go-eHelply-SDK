# GetTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | Pointer to [**GetInvoiceResponse**](GetInvoiceResponse.md) |  | [optional] 
**TransactionUuid** | **string** |  | 
**StripeId** | **string** |  | 
**Credit** | **int32** |  | 
**Debit** | **int32** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewGetTransactionResponse

`func NewGetTransactionResponse(transactionUuid string, stripeId string, credit int32, debit int32, createdAt string, ) *GetTransactionResponse`

NewGetTransactionResponse instantiates a new GetTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionResponseWithDefaults

`func NewGetTransactionResponseWithDefaults() *GetTransactionResponse`

NewGetTransactionResponseWithDefaults instantiates a new GetTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *GetTransactionResponse) GetInvoice() GetInvoiceResponse`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *GetTransactionResponse) GetInvoiceOk() (*GetInvoiceResponse, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *GetTransactionResponse) SetInvoice(v GetInvoiceResponse)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *GetTransactionResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetTransactionUuid

`func (o *GetTransactionResponse) GetTransactionUuid() string`

GetTransactionUuid returns the TransactionUuid field if non-nil, zero value otherwise.

### GetTransactionUuidOk

`func (o *GetTransactionResponse) GetTransactionUuidOk() (*string, bool)`

GetTransactionUuidOk returns a tuple with the TransactionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUuid

`func (o *GetTransactionResponse) SetTransactionUuid(v string)`

SetTransactionUuid sets TransactionUuid field to given value.


### GetStripeId

`func (o *GetTransactionResponse) GetStripeId() string`

GetStripeId returns the StripeId field if non-nil, zero value otherwise.

### GetStripeIdOk

`func (o *GetTransactionResponse) GetStripeIdOk() (*string, bool)`

GetStripeIdOk returns a tuple with the StripeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeId

`func (o *GetTransactionResponse) SetStripeId(v string)`

SetStripeId sets StripeId field to given value.


### GetCredit

`func (o *GetTransactionResponse) GetCredit() int32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *GetTransactionResponse) GetCreditOk() (*int32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *GetTransactionResponse) SetCredit(v int32)`

SetCredit sets Credit field to given value.


### GetDebit

`func (o *GetTransactionResponse) GetDebit() int32`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *GetTransactionResponse) GetDebitOk() (*int32, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *GetTransactionResponse) SetDebit(v int32)`

SetDebit sets Debit field to given value.


### GetCreatedAt

`func (o *GetTransactionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTransactionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTransactionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


