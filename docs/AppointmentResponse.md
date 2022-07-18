# AppointmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectUuid** | Pointer to **string** |  | [optional] 
**PlaceUuid** | Pointer to **string** |  | [optional] 
**ReviewGroupUuid** | Pointer to **string** |  | [optional] 
**ExpectedFinishAt** | Pointer to **string** |  | [optional] 
**ExpectedStartAt** | Pointer to **string** |  | [optional] 
**ActualStartAt** | Pointer to **string** |  | [optional] 
**ActualFinishAt** | Pointer to **string** |  | [optional] 
**Products** | Pointer to **map[string]interface{}** |  | [optional] 
**MetaUuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CancellationAt** | Pointer to **string** |  | [optional] 
**CancellationReason** | Pointer to **string** |  | [optional] 
**CancellationEntityUuid** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAppointmentResponse

`func NewAppointmentResponse(uuid string, createdAt string, updatedAt string, ) *AppointmentResponse`

NewAppointmentResponse instantiates a new AppointmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppointmentResponseWithDefaults

`func NewAppointmentResponseWithDefaults() *AppointmentResponse`

NewAppointmentResponseWithDefaults instantiates a new AppointmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectUuid

`func (o *AppointmentResponse) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *AppointmentResponse) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *AppointmentResponse) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *AppointmentResponse) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetPlaceUuid

`func (o *AppointmentResponse) GetPlaceUuid() string`

GetPlaceUuid returns the PlaceUuid field if non-nil, zero value otherwise.

### GetPlaceUuidOk

`func (o *AppointmentResponse) GetPlaceUuidOk() (*string, bool)`

GetPlaceUuidOk returns a tuple with the PlaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceUuid

`func (o *AppointmentResponse) SetPlaceUuid(v string)`

SetPlaceUuid sets PlaceUuid field to given value.

### HasPlaceUuid

`func (o *AppointmentResponse) HasPlaceUuid() bool`

HasPlaceUuid returns a boolean if a field has been set.

### GetReviewGroupUuid

`func (o *AppointmentResponse) GetReviewGroupUuid() string`

GetReviewGroupUuid returns the ReviewGroupUuid field if non-nil, zero value otherwise.

### GetReviewGroupUuidOk

`func (o *AppointmentResponse) GetReviewGroupUuidOk() (*string, bool)`

GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewGroupUuid

`func (o *AppointmentResponse) SetReviewGroupUuid(v string)`

SetReviewGroupUuid sets ReviewGroupUuid field to given value.

### HasReviewGroupUuid

`func (o *AppointmentResponse) HasReviewGroupUuid() bool`

HasReviewGroupUuid returns a boolean if a field has been set.

### GetExpectedFinishAt

`func (o *AppointmentResponse) GetExpectedFinishAt() string`

GetExpectedFinishAt returns the ExpectedFinishAt field if non-nil, zero value otherwise.

### GetExpectedFinishAtOk

`func (o *AppointmentResponse) GetExpectedFinishAtOk() (*string, bool)`

GetExpectedFinishAtOk returns a tuple with the ExpectedFinishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedFinishAt

`func (o *AppointmentResponse) SetExpectedFinishAt(v string)`

SetExpectedFinishAt sets ExpectedFinishAt field to given value.

### HasExpectedFinishAt

`func (o *AppointmentResponse) HasExpectedFinishAt() bool`

HasExpectedFinishAt returns a boolean if a field has been set.

### GetExpectedStartAt

`func (o *AppointmentResponse) GetExpectedStartAt() string`

GetExpectedStartAt returns the ExpectedStartAt field if non-nil, zero value otherwise.

### GetExpectedStartAtOk

`func (o *AppointmentResponse) GetExpectedStartAtOk() (*string, bool)`

GetExpectedStartAtOk returns a tuple with the ExpectedStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStartAt

`func (o *AppointmentResponse) SetExpectedStartAt(v string)`

SetExpectedStartAt sets ExpectedStartAt field to given value.

### HasExpectedStartAt

`func (o *AppointmentResponse) HasExpectedStartAt() bool`

HasExpectedStartAt returns a boolean if a field has been set.

### GetActualStartAt

`func (o *AppointmentResponse) GetActualStartAt() string`

GetActualStartAt returns the ActualStartAt field if non-nil, zero value otherwise.

### GetActualStartAtOk

`func (o *AppointmentResponse) GetActualStartAtOk() (*string, bool)`

GetActualStartAtOk returns a tuple with the ActualStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStartAt

`func (o *AppointmentResponse) SetActualStartAt(v string)`

SetActualStartAt sets ActualStartAt field to given value.

### HasActualStartAt

`func (o *AppointmentResponse) HasActualStartAt() bool`

HasActualStartAt returns a boolean if a field has been set.

### GetActualFinishAt

`func (o *AppointmentResponse) GetActualFinishAt() string`

GetActualFinishAt returns the ActualFinishAt field if non-nil, zero value otherwise.

### GetActualFinishAtOk

`func (o *AppointmentResponse) GetActualFinishAtOk() (*string, bool)`

GetActualFinishAtOk returns a tuple with the ActualFinishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualFinishAt

`func (o *AppointmentResponse) SetActualFinishAt(v string)`

SetActualFinishAt sets ActualFinishAt field to given value.

### HasActualFinishAt

`func (o *AppointmentResponse) HasActualFinishAt() bool`

HasActualFinishAt returns a boolean if a field has been set.

### GetProducts

`func (o *AppointmentResponse) GetProducts() map[string]interface{}`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AppointmentResponse) GetProductsOk() (*map[string]interface{}, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AppointmentResponse) SetProducts(v map[string]interface{})`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AppointmentResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetMetaUuid

`func (o *AppointmentResponse) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *AppointmentResponse) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *AppointmentResponse) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *AppointmentResponse) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetStatus

`func (o *AppointmentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppointmentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppointmentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppointmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCancellationAt

`func (o *AppointmentResponse) GetCancellationAt() string`

GetCancellationAt returns the CancellationAt field if non-nil, zero value otherwise.

### GetCancellationAtOk

`func (o *AppointmentResponse) GetCancellationAtOk() (*string, bool)`

GetCancellationAtOk returns a tuple with the CancellationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationAt

`func (o *AppointmentResponse) SetCancellationAt(v string)`

SetCancellationAt sets CancellationAt field to given value.

### HasCancellationAt

`func (o *AppointmentResponse) HasCancellationAt() bool`

HasCancellationAt returns a boolean if a field has been set.

### GetCancellationReason

`func (o *AppointmentResponse) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *AppointmentResponse) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *AppointmentResponse) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.

### HasCancellationReason

`func (o *AppointmentResponse) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### GetCancellationEntityUuid

`func (o *AppointmentResponse) GetCancellationEntityUuid() string`

GetCancellationEntityUuid returns the CancellationEntityUuid field if non-nil, zero value otherwise.

### GetCancellationEntityUuidOk

`func (o *AppointmentResponse) GetCancellationEntityUuidOk() (*string, bool)`

GetCancellationEntityUuidOk returns a tuple with the CancellationEntityUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationEntityUuid

`func (o *AppointmentResponse) SetCancellationEntityUuid(v string)`

SetCancellationEntityUuid sets CancellationEntityUuid field to given value.

### HasCancellationEntityUuid

`func (o *AppointmentResponse) HasCancellationEntityUuid() bool`

HasCancellationEntityUuid returns a boolean if a field has been set.

### GetUuid

`func (o *AppointmentResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppointmentResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppointmentResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCreatedAt

`func (o *AppointmentResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppointmentResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppointmentResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppointmentResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppointmentResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppointmentResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *AppointmentResponse) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AppointmentResponse) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AppointmentResponse) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AppointmentResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


