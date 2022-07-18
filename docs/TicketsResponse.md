# TicketsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** |  | 
**Priority** | **string** |  | 
**TicketId** | **string** |  | 

## Methods

### NewTicketsResponse

`func NewTicketsResponse(subject string, priority string, ticketId string, ) *TicketsResponse`

NewTicketsResponse instantiates a new TicketsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketsResponseWithDefaults

`func NewTicketsResponseWithDefaults() *TicketsResponse`

NewTicketsResponseWithDefaults instantiates a new TicketsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *TicketsResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketsResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketsResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetPriority

`func (o *TicketsResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketsResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketsResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetTicketId

`func (o *TicketsResponse) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *TicketsResponse) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *TicketsResponse) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


