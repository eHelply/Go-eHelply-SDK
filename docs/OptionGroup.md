# OptionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Selections** | Pointer to [**[]Selection**](Selection.md) |  | [optional] 

## Methods

### NewOptionGroup

`func NewOptionGroup() *OptionGroup`

NewOptionGroup instantiates a new OptionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionGroupWithDefaults

`func NewOptionGroupWithDefaults() *OptionGroup`

NewOptionGroupWithDefaults instantiates a new OptionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OptionGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *OptionGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSelections

`func (o *OptionGroup) GetSelections() []Selection`

GetSelections returns the Selections field if non-nil, zero value otherwise.

### GetSelectionsOk

`func (o *OptionGroup) GetSelectionsOk() (*[]Selection, bool)`

GetSelectionsOk returns a tuple with the Selections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelections

`func (o *OptionGroup) SetSelections(v []Selection)`

SetSelections sets Selections field to given value.

### HasSelections

`func (o *OptionGroup) HasSelections() bool`

HasSelections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


