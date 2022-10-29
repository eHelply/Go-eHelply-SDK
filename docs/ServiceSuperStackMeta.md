# ServiceSuperStackMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HumanName** | **string** |  | 
**RouteName** | **string** |  | 
**ServiceName** | **string** |  | 
**Advertise** | **bool** |  | 
**Featured** | **bool** |  | 
**Picture** | **string** |  | 
**BackgroundPicture** | **string** |  | 
**TagLine** | **string** |  | 
**Summary** | **string** |  | 
**Description** | **string** |  | 
**Features** | [**[]ServiceSuperStackMetaFeature**](ServiceSuperStackMetaFeature.md) |  | 
**UseCases** | [**[]ServiceSuperStackMetaUseCase**](ServiceSuperStackMetaUseCase.md) |  | 
**GettingStarted** | [**ServiceSuperStackMetaGettingStarted**](ServiceSuperStackMetaGettingStarted.md) |  | 
**Faqs** | [**[]ServiceSuperStackMetaFaq**](ServiceSuperStackMetaFaq.md) |  | 

## Methods

### NewServiceSuperStackMeta

`func NewServiceSuperStackMeta(humanName string, routeName string, serviceName string, advertise bool, featured bool, picture string, backgroundPicture string, tagLine string, summary string, description string, features []ServiceSuperStackMetaFeature, useCases []ServiceSuperStackMetaUseCase, gettingStarted ServiceSuperStackMetaGettingStarted, faqs []ServiceSuperStackMetaFaq, ) *ServiceSuperStackMeta`

NewServiceSuperStackMeta instantiates a new ServiceSuperStackMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSuperStackMetaWithDefaults

`func NewServiceSuperStackMetaWithDefaults() *ServiceSuperStackMeta`

NewServiceSuperStackMetaWithDefaults instantiates a new ServiceSuperStackMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHumanName

`func (o *ServiceSuperStackMeta) GetHumanName() string`

GetHumanName returns the HumanName field if non-nil, zero value otherwise.

### GetHumanNameOk

`func (o *ServiceSuperStackMeta) GetHumanNameOk() (*string, bool)`

GetHumanNameOk returns a tuple with the HumanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanName

`func (o *ServiceSuperStackMeta) SetHumanName(v string)`

SetHumanName sets HumanName field to given value.


### GetRouteName

`func (o *ServiceSuperStackMeta) GetRouteName() string`

GetRouteName returns the RouteName field if non-nil, zero value otherwise.

### GetRouteNameOk

`func (o *ServiceSuperStackMeta) GetRouteNameOk() (*string, bool)`

GetRouteNameOk returns a tuple with the RouteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteName

`func (o *ServiceSuperStackMeta) SetRouteName(v string)`

SetRouteName sets RouteName field to given value.


### GetServiceName

`func (o *ServiceSuperStackMeta) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceSuperStackMeta) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceSuperStackMeta) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetAdvertise

`func (o *ServiceSuperStackMeta) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *ServiceSuperStackMeta) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *ServiceSuperStackMeta) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.


### GetFeatured

`func (o *ServiceSuperStackMeta) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ServiceSuperStackMeta) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ServiceSuperStackMeta) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.


### GetPicture

`func (o *ServiceSuperStackMeta) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *ServiceSuperStackMeta) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *ServiceSuperStackMeta) SetPicture(v string)`

SetPicture sets Picture field to given value.


### GetBackgroundPicture

`func (o *ServiceSuperStackMeta) GetBackgroundPicture() string`

GetBackgroundPicture returns the BackgroundPicture field if non-nil, zero value otherwise.

### GetBackgroundPictureOk

`func (o *ServiceSuperStackMeta) GetBackgroundPictureOk() (*string, bool)`

GetBackgroundPictureOk returns a tuple with the BackgroundPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPicture

`func (o *ServiceSuperStackMeta) SetBackgroundPicture(v string)`

SetBackgroundPicture sets BackgroundPicture field to given value.


### GetTagLine

`func (o *ServiceSuperStackMeta) GetTagLine() string`

GetTagLine returns the TagLine field if non-nil, zero value otherwise.

### GetTagLineOk

`func (o *ServiceSuperStackMeta) GetTagLineOk() (*string, bool)`

GetTagLineOk returns a tuple with the TagLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagLine

`func (o *ServiceSuperStackMeta) SetTagLine(v string)`

SetTagLine sets TagLine field to given value.


### GetSummary

`func (o *ServiceSuperStackMeta) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ServiceSuperStackMeta) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ServiceSuperStackMeta) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *ServiceSuperStackMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceSuperStackMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceSuperStackMeta) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *ServiceSuperStackMeta) GetFeatures() []ServiceSuperStackMetaFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceSuperStackMeta) GetFeaturesOk() (*[]ServiceSuperStackMetaFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceSuperStackMeta) SetFeatures(v []ServiceSuperStackMetaFeature)`

SetFeatures sets Features field to given value.


### GetUseCases

`func (o *ServiceSuperStackMeta) GetUseCases() []ServiceSuperStackMetaUseCase`

GetUseCases returns the UseCases field if non-nil, zero value otherwise.

### GetUseCasesOk

`func (o *ServiceSuperStackMeta) GetUseCasesOk() (*[]ServiceSuperStackMetaUseCase, bool)`

GetUseCasesOk returns a tuple with the UseCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCases

`func (o *ServiceSuperStackMeta) SetUseCases(v []ServiceSuperStackMetaUseCase)`

SetUseCases sets UseCases field to given value.


### GetGettingStarted

`func (o *ServiceSuperStackMeta) GetGettingStarted() ServiceSuperStackMetaGettingStarted`

GetGettingStarted returns the GettingStarted field if non-nil, zero value otherwise.

### GetGettingStartedOk

`func (o *ServiceSuperStackMeta) GetGettingStartedOk() (*ServiceSuperStackMetaGettingStarted, bool)`

GetGettingStartedOk returns a tuple with the GettingStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettingStarted

`func (o *ServiceSuperStackMeta) SetGettingStarted(v ServiceSuperStackMetaGettingStarted)`

SetGettingStarted sets GettingStarted field to given value.


### GetFaqs

`func (o *ServiceSuperStackMeta) GetFaqs() []ServiceSuperStackMetaFaq`

GetFaqs returns the Faqs field if non-nil, zero value otherwise.

### GetFaqsOk

`func (o *ServiceSuperStackMeta) GetFaqsOk() (*[]ServiceSuperStackMetaFaq, bool)`

GetFaqsOk returns a tuple with the Faqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaqs

`func (o *ServiceSuperStackMeta) SetFaqs(v []ServiceSuperStackMetaFaq)`

SetFaqs sets Faqs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


