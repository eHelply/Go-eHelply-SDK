# CreateTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **string** |  | 
**Subject** | **string** |  | 

## Methods

### NewCreateTicket

`func NewCreateTicket(priority string, subject string, ) *CreateTicket`

NewCreateTicket instantiates a new CreateTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketWithDefaults

`func NewCreateTicketWithDefaults() *CreateTicket`

NewCreateTicketWithDefaults instantiates a new CreateTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *CreateTicket) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateTicket) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateTicket) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetSubject

`func (o *CreateTicket) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateTicket) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateTicket) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


