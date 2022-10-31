# FactCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Data** | **map[string]interface{}** |  | 
**Public** | **bool** |  | 

## Methods

### NewFactCreate

`func NewFactCreate(name string, data map[string]interface{}, public bool, ) *FactCreate`

NewFactCreate instantiates a new FactCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactCreateWithDefaults

`func NewFactCreateWithDefaults() *FactCreate`

NewFactCreateWithDefaults instantiates a new FactCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FactCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FactCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FactCreate) SetName(v string)`

SetName sets Name field to given value.


### GetData

`func (o *FactCreate) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FactCreate) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FactCreate) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetPublic

`func (o *FactCreate) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FactCreate) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FactCreate) SetPublic(v bool)`

SetPublic sets Public field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


