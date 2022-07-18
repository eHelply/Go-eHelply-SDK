# AppointmentBase

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

## Methods

### NewAppointmentBase

`func NewAppointmentBase() *AppointmentBase`

NewAppointmentBase instantiates a new AppointmentBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppointmentBaseWithDefaults

`func NewAppointmentBaseWithDefaults() *AppointmentBase`

NewAppointmentBaseWithDefaults instantiates a new AppointmentBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectUuid

`func (o *AppointmentBase) GetProjectUuid() string`

GetProjectUuid returns the ProjectUuid field if non-nil, zero value otherwise.

### GetProjectUuidOk

`func (o *AppointmentBase) GetProjectUuidOk() (*string, bool)`

GetProjectUuidOk returns a tuple with the ProjectUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUuid

`func (o *AppointmentBase) SetProjectUuid(v string)`

SetProjectUuid sets ProjectUuid field to given value.

### HasProjectUuid

`func (o *AppointmentBase) HasProjectUuid() bool`

HasProjectUuid returns a boolean if a field has been set.

### GetPlaceUuid

`func (o *AppointmentBase) GetPlaceUuid() string`

GetPlaceUuid returns the PlaceUuid field if non-nil, zero value otherwise.

### GetPlaceUuidOk

`func (o *AppointmentBase) GetPlaceUuidOk() (*string, bool)`

GetPlaceUuidOk returns a tuple with the PlaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceUuid

`func (o *AppointmentBase) SetPlaceUuid(v string)`

SetPlaceUuid sets PlaceUuid field to given value.

### HasPlaceUuid

`func (o *AppointmentBase) HasPlaceUuid() bool`

HasPlaceUuid returns a boolean if a field has been set.

### GetReviewGroupUuid

`func (o *AppointmentBase) GetReviewGroupUuid() string`

GetReviewGroupUuid returns the ReviewGroupUuid field if non-nil, zero value otherwise.

### GetReviewGroupUuidOk

`func (o *AppointmentBase) GetReviewGroupUuidOk() (*string, bool)`

GetReviewGroupUuidOk returns a tuple with the ReviewGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewGroupUuid

`func (o *AppointmentBase) SetReviewGroupUuid(v string)`

SetReviewGroupUuid sets ReviewGroupUuid field to given value.

### HasReviewGroupUuid

`func (o *AppointmentBase) HasReviewGroupUuid() bool`

HasReviewGroupUuid returns a boolean if a field has been set.

### GetExpectedFinishAt

`func (o *AppointmentBase) GetExpectedFinishAt() string`

GetExpectedFinishAt returns the ExpectedFinishAt field if non-nil, zero value otherwise.

### GetExpectedFinishAtOk

`func (o *AppointmentBase) GetExpectedFinishAtOk() (*string, bool)`

GetExpectedFinishAtOk returns a tuple with the ExpectedFinishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedFinishAt

`func (o *AppointmentBase) SetExpectedFinishAt(v string)`

SetExpectedFinishAt sets ExpectedFinishAt field to given value.

### HasExpectedFinishAt

`func (o *AppointmentBase) HasExpectedFinishAt() bool`

HasExpectedFinishAt returns a boolean if a field has been set.

### GetExpectedStartAt

`func (o *AppointmentBase) GetExpectedStartAt() string`

GetExpectedStartAt returns the ExpectedStartAt field if non-nil, zero value otherwise.

### GetExpectedStartAtOk

`func (o *AppointmentBase) GetExpectedStartAtOk() (*string, bool)`

GetExpectedStartAtOk returns a tuple with the ExpectedStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedStartAt

`func (o *AppointmentBase) SetExpectedStartAt(v string)`

SetExpectedStartAt sets ExpectedStartAt field to given value.

### HasExpectedStartAt

`func (o *AppointmentBase) HasExpectedStartAt() bool`

HasExpectedStartAt returns a boolean if a field has been set.

### GetActualStartAt

`func (o *AppointmentBase) GetActualStartAt() string`

GetActualStartAt returns the ActualStartAt field if non-nil, zero value otherwise.

### GetActualStartAtOk

`func (o *AppointmentBase) GetActualStartAtOk() (*string, bool)`

GetActualStartAtOk returns a tuple with the ActualStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualStartAt

`func (o *AppointmentBase) SetActualStartAt(v string)`

SetActualStartAt sets ActualStartAt field to given value.

### HasActualStartAt

`func (o *AppointmentBase) HasActualStartAt() bool`

HasActualStartAt returns a boolean if a field has been set.

### GetActualFinishAt

`func (o *AppointmentBase) GetActualFinishAt() string`

GetActualFinishAt returns the ActualFinishAt field if non-nil, zero value otherwise.

### GetActualFinishAtOk

`func (o *AppointmentBase) GetActualFinishAtOk() (*string, bool)`

GetActualFinishAtOk returns a tuple with the ActualFinishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualFinishAt

`func (o *AppointmentBase) SetActualFinishAt(v string)`

SetActualFinishAt sets ActualFinishAt field to given value.

### HasActualFinishAt

`func (o *AppointmentBase) HasActualFinishAt() bool`

HasActualFinishAt returns a boolean if a field has been set.

### GetProducts

`func (o *AppointmentBase) GetProducts() map[string]interface{}`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AppointmentBase) GetProductsOk() (*map[string]interface{}, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AppointmentBase) SetProducts(v map[string]interface{})`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AppointmentBase) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetMetaUuid

`func (o *AppointmentBase) GetMetaUuid() string`

GetMetaUuid returns the MetaUuid field if non-nil, zero value otherwise.

### GetMetaUuidOk

`func (o *AppointmentBase) GetMetaUuidOk() (*string, bool)`

GetMetaUuidOk returns a tuple with the MetaUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaUuid

`func (o *AppointmentBase) SetMetaUuid(v string)`

SetMetaUuid sets MetaUuid field to given value.

### HasMetaUuid

`func (o *AppointmentBase) HasMetaUuid() bool`

HasMetaUuid returns a boolean if a field has been set.

### GetStatus

`func (o *AppointmentBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppointmentBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppointmentBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppointmentBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCancellationAt

`func (o *AppointmentBase) GetCancellationAt() string`

GetCancellationAt returns the CancellationAt field if non-nil, zero value otherwise.

### GetCancellationAtOk

`func (o *AppointmentBase) GetCancellationAtOk() (*string, bool)`

GetCancellationAtOk returns a tuple with the CancellationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationAt

`func (o *AppointmentBase) SetCancellationAt(v string)`

SetCancellationAt sets CancellationAt field to given value.

### HasCancellationAt

`func (o *AppointmentBase) HasCancellationAt() bool`

HasCancellationAt returns a boolean if a field has been set.

### GetCancellationReason

`func (o *AppointmentBase) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *AppointmentBase) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *AppointmentBase) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.

### HasCancellationReason

`func (o *AppointmentBase) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### GetCancellationEntityUuid

`func (o *AppointmentBase) GetCancellationEntityUuid() string`

GetCancellationEntityUuid returns the CancellationEntityUuid field if non-nil, zero value otherwise.

### GetCancellationEntityUuidOk

`func (o *AppointmentBase) GetCancellationEntityUuidOk() (*string, bool)`

GetCancellationEntityUuidOk returns a tuple with the CancellationEntityUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationEntityUuid

`func (o *AppointmentBase) SetCancellationEntityUuid(v string)`

SetCancellationEntityUuid sets CancellationEntityUuid field to given value.

### HasCancellationEntityUuid

`func (o *AppointmentBase) HasCancellationEntityUuid() bool`

HasCancellationEntityUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


