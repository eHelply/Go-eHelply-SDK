/*
eHelply SDK - 1.1.121

eHelply SDK for SuperStack Services

API version: 1.1.121
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"encoding/json"
)

// ServiceSuperStackMeta struct for ServiceSuperStackMeta
type ServiceSuperStackMeta struct {
	HumanName string `json:"human_name"`
	RouteName string `json:"route_name"`
	ServiceName string `json:"service_name"`
	Advertise bool `json:"advertise"`
	Featured bool `json:"featured"`
	Picture string `json:"picture"`
	BackgroundPicture string `json:"background_picture"`
	TagLine string `json:"tag_line"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	Features []ServiceSuperStackMetaFeature `json:"features"`
	UseCases []ServiceSuperStackMetaUseCase `json:"use_cases"`
	GettingStarted ServiceSuperStackMetaGettingStarted `json:"getting_started"`
	Faqs []ServiceSuperStackMetaFaq `json:"faqs"`
}

// NewServiceSuperStackMeta instantiates a new ServiceSuperStackMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSuperStackMeta(humanName string, routeName string, serviceName string, advertise bool, featured bool, picture string, backgroundPicture string, tagLine string, summary string, description string, features []ServiceSuperStackMetaFeature, useCases []ServiceSuperStackMetaUseCase, gettingStarted ServiceSuperStackMetaGettingStarted, faqs []ServiceSuperStackMetaFaq) *ServiceSuperStackMeta {
	this := ServiceSuperStackMeta{}
	this.HumanName = humanName
	this.RouteName = routeName
	this.ServiceName = serviceName
	this.Advertise = advertise
	this.Featured = featured
	this.Picture = picture
	this.BackgroundPicture = backgroundPicture
	this.TagLine = tagLine
	this.Summary = summary
	this.Description = description
	this.Features = features
	this.UseCases = useCases
	this.GettingStarted = gettingStarted
	this.Faqs = faqs
	return &this
}

// NewServiceSuperStackMetaWithDefaults instantiates a new ServiceSuperStackMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSuperStackMetaWithDefaults() *ServiceSuperStackMeta {
	this := ServiceSuperStackMeta{}
	return &this
}

// GetHumanName returns the HumanName field value
func (o *ServiceSuperStackMeta) GetHumanName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HumanName
}

// GetHumanNameOk returns a tuple with the HumanName field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetHumanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HumanName, true
}

// SetHumanName sets field value
func (o *ServiceSuperStackMeta) SetHumanName(v string) {
	o.HumanName = v
}

// GetRouteName returns the RouteName field value
func (o *ServiceSuperStackMeta) GetRouteName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteName
}

// GetRouteNameOk returns a tuple with the RouteName field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetRouteNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteName, true
}

// SetRouteName sets field value
func (o *ServiceSuperStackMeta) SetRouteName(v string) {
	o.RouteName = v
}

// GetServiceName returns the ServiceName field value
func (o *ServiceSuperStackMeta) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *ServiceSuperStackMeta) SetServiceName(v string) {
	o.ServiceName = v
}

// GetAdvertise returns the Advertise field value
func (o *ServiceSuperStackMeta) GetAdvertise() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Advertise
}

// GetAdvertiseOk returns a tuple with the Advertise field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetAdvertiseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Advertise, true
}

// SetAdvertise sets field value
func (o *ServiceSuperStackMeta) SetAdvertise(v bool) {
	o.Advertise = v
}

// GetFeatured returns the Featured field value
func (o *ServiceSuperStackMeta) GetFeatured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetFeaturedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Featured, true
}

// SetFeatured sets field value
func (o *ServiceSuperStackMeta) SetFeatured(v bool) {
	o.Featured = v
}

// GetPicture returns the Picture field value
func (o *ServiceSuperStackMeta) GetPicture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Picture
}

// GetPictureOk returns a tuple with the Picture field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetPictureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Picture, true
}

// SetPicture sets field value
func (o *ServiceSuperStackMeta) SetPicture(v string) {
	o.Picture = v
}

// GetBackgroundPicture returns the BackgroundPicture field value
func (o *ServiceSuperStackMeta) GetBackgroundPicture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackgroundPicture
}

// GetBackgroundPictureOk returns a tuple with the BackgroundPicture field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetBackgroundPictureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundPicture, true
}

// SetBackgroundPicture sets field value
func (o *ServiceSuperStackMeta) SetBackgroundPicture(v string) {
	o.BackgroundPicture = v
}

// GetTagLine returns the TagLine field value
func (o *ServiceSuperStackMeta) GetTagLine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagLine
}

// GetTagLineOk returns a tuple with the TagLine field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetTagLineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagLine, true
}

// SetTagLine sets field value
func (o *ServiceSuperStackMeta) SetTagLine(v string) {
	o.TagLine = v
}

// GetSummary returns the Summary field value
func (o *ServiceSuperStackMeta) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ServiceSuperStackMeta) SetSummary(v string) {
	o.Summary = v
}

// GetDescription returns the Description field value
func (o *ServiceSuperStackMeta) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceSuperStackMeta) SetDescription(v string) {
	o.Description = v
}

// GetFeatures returns the Features field value
func (o *ServiceSuperStackMeta) GetFeatures() []ServiceSuperStackMetaFeature {
	if o == nil {
		var ret []ServiceSuperStackMetaFeature
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetFeaturesOk() ([]ServiceSuperStackMetaFeature, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *ServiceSuperStackMeta) SetFeatures(v []ServiceSuperStackMetaFeature) {
	o.Features = v
}

// GetUseCases returns the UseCases field value
func (o *ServiceSuperStackMeta) GetUseCases() []ServiceSuperStackMetaUseCase {
	if o == nil {
		var ret []ServiceSuperStackMetaUseCase
		return ret
	}

	return o.UseCases
}

// GetUseCasesOk returns a tuple with the UseCases field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetUseCasesOk() ([]ServiceSuperStackMetaUseCase, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseCases, true
}

// SetUseCases sets field value
func (o *ServiceSuperStackMeta) SetUseCases(v []ServiceSuperStackMetaUseCase) {
	o.UseCases = v
}

// GetGettingStarted returns the GettingStarted field value
func (o *ServiceSuperStackMeta) GetGettingStarted() ServiceSuperStackMetaGettingStarted {
	if o == nil {
		var ret ServiceSuperStackMetaGettingStarted
		return ret
	}

	return o.GettingStarted
}

// GetGettingStartedOk returns a tuple with the GettingStarted field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetGettingStartedOk() (*ServiceSuperStackMetaGettingStarted, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GettingStarted, true
}

// SetGettingStarted sets field value
func (o *ServiceSuperStackMeta) SetGettingStarted(v ServiceSuperStackMetaGettingStarted) {
	o.GettingStarted = v
}

// GetFaqs returns the Faqs field value
func (o *ServiceSuperStackMeta) GetFaqs() []ServiceSuperStackMetaFaq {
	if o == nil {
		var ret []ServiceSuperStackMetaFaq
		return ret
	}

	return o.Faqs
}

// GetFaqsOk returns a tuple with the Faqs field value
// and a boolean to check if the value has been set.
func (o *ServiceSuperStackMeta) GetFaqsOk() ([]ServiceSuperStackMetaFaq, bool) {
	if o == nil {
		return nil, false
	}
	return o.Faqs, true
}

// SetFaqs sets field value
func (o *ServiceSuperStackMeta) SetFaqs(v []ServiceSuperStackMetaFaq) {
	o.Faqs = v
}

func (o ServiceSuperStackMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["human_name"] = o.HumanName
	}
	if true {
		toSerialize["route_name"] = o.RouteName
	}
	if true {
		toSerialize["service_name"] = o.ServiceName
	}
	if true {
		toSerialize["advertise"] = o.Advertise
	}
	if true {
		toSerialize["featured"] = o.Featured
	}
	if true {
		toSerialize["picture"] = o.Picture
	}
	if true {
		toSerialize["background_picture"] = o.BackgroundPicture
	}
	if true {
		toSerialize["tag_line"] = o.TagLine
	}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["features"] = o.Features
	}
	if true {
		toSerialize["use_cases"] = o.UseCases
	}
	if true {
		toSerialize["getting_started"] = o.GettingStarted
	}
	if true {
		toSerialize["faqs"] = o.Faqs
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSuperStackMeta struct {
	value *ServiceSuperStackMeta
	isSet bool
}

func (v NullableServiceSuperStackMeta) Get() *ServiceSuperStackMeta {
	return v.value
}

func (v *NullableServiceSuperStackMeta) Set(val *ServiceSuperStackMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSuperStackMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSuperStackMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSuperStackMeta(val *ServiceSuperStackMeta) *NullableServiceSuperStackMeta {
	return &NullableServiceSuperStackMeta{value: val, isSet: true}
}

func (v NullableServiceSuperStackMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSuperStackMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


