# Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**InsetLabel** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**MaxLength** | Pointer to **float32** |  | [optional] 
**Counter** | Pointer to **bool** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IconPosition** | Pointer to **string** |  | [optional] 
**Selections** | Pointer to [**[]OptionGroup**](OptionGroup.md) |  | [optional] 

## Methods

### NewOptions

`func NewOptions() *Options`

NewOptions instantiates a new Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsWithDefaults

`func NewOptionsWithDefaults() *Options`

NewOptionsWithDefaults instantiates a new Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *Options) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Options) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Options) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Options) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetLabel

`func (o *Options) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Options) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Options) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Options) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetInsetLabel

`func (o *Options) GetInsetLabel() string`

GetInsetLabel returns the InsetLabel field if non-nil, zero value otherwise.

### GetInsetLabelOk

`func (o *Options) GetInsetLabelOk() (*string, bool)`

GetInsetLabelOk returns a tuple with the InsetLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsetLabel

`func (o *Options) SetInsetLabel(v string)`

SetInsetLabel sets InsetLabel field to given value.

### HasInsetLabel

`func (o *Options) HasInsetLabel() bool`

HasInsetLabel returns a boolean if a field has been set.

### GetPlaceholder

`func (o *Options) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *Options) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *Options) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *Options) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetHint

`func (o *Options) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *Options) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *Options) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *Options) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetIcon

`func (o *Options) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Options) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Options) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Options) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetMaxLength

`func (o *Options) GetMaxLength() float32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *Options) GetMaxLengthOk() (*float32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *Options) SetMaxLength(v float32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *Options) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetCounter

`func (o *Options) GetCounter() bool`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *Options) GetCounterOk() (*bool, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *Options) SetCounter(v bool)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *Options) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetCaption

`func (o *Options) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *Options) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *Options) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *Options) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetColor

`func (o *Options) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Options) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Options) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Options) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetSize

`func (o *Options) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Options) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Options) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Options) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *Options) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Options) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Options) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Options) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIconPosition

`func (o *Options) GetIconPosition() string`

GetIconPosition returns the IconPosition field if non-nil, zero value otherwise.

### GetIconPositionOk

`func (o *Options) GetIconPositionOk() (*string, bool)`

GetIconPositionOk returns a tuple with the IconPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPosition

`func (o *Options) SetIconPosition(v string)`

SetIconPosition sets IconPosition field to given value.

### HasIconPosition

`func (o *Options) HasIconPosition() bool`

HasIconPosition returns a boolean if a field has been set.

### GetSelections

`func (o *Options) GetSelections() []OptionGroup`

GetSelections returns the Selections field if non-nil, zero value otherwise.

### GetSelectionsOk

`func (o *Options) GetSelectionsOk() (*[]OptionGroup, bool)`

GetSelectionsOk returns a tuple with the Selections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelections

`func (o *Options) SetSelections(v []OptionGroup)`

SetSelections sets Selections field to given value.

### HasSelections

`func (o *Options) HasSelections() bool`

HasSelections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


