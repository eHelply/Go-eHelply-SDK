# FieldDynamo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Type** | Pointer to **int32** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Validations** | Pointer to [**Validations**](Validations.md) |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**Options**](Options.md) |  | [optional] 

## Methods

### NewFieldDynamo

`func NewFieldDynamo(uuid string, ) *FieldDynamo`

NewFieldDynamo instantiates a new FieldDynamo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDynamoWithDefaults

`func NewFieldDynamoWithDefaults() *FieldDynamo`

NewFieldDynamoWithDefaults instantiates a new FieldDynamo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *FieldDynamo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FieldDynamo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FieldDynamo) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *FieldDynamo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldDynamo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldDynamo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *FieldDynamo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPlaceholder

`func (o *FieldDynamo) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *FieldDynamo) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *FieldDynamo) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *FieldDynamo) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetValidations

`func (o *FieldDynamo) GetValidations() Validations`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *FieldDynamo) GetValidationsOk() (*Validations, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *FieldDynamo) SetValidations(v Validations)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *FieldDynamo) HasValidations() bool`

HasValidations returns a boolean if a field has been set.

### GetHint

`func (o *FieldDynamo) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *FieldDynamo) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *FieldDynamo) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *FieldDynamo) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetIcon

`func (o *FieldDynamo) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FieldDynamo) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FieldDynamo) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FieldDynamo) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLabel

`func (o *FieldDynamo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldDynamo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldDynamo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FieldDynamo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptions

`func (o *FieldDynamo) GetOptions() Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FieldDynamo) GetOptionsOk() (*Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FieldDynamo) SetOptions(v Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FieldDynamo) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


