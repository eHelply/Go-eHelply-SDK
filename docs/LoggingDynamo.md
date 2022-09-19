# LoggingDynamo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** |  | 
**Time** | **string** |  | 
**Log** | **map[string]interface{}** |  | 
**Severity** | **string** |  | 
**Subject** | **string** |  | 
**ServiceName** | **string** |  | 

## Methods

### NewLoggingDynamo

`func NewLoggingDynamo(service string, time string, log map[string]interface{}, severity string, subject string, serviceName string, ) *LoggingDynamo`

NewLoggingDynamo instantiates a new LoggingDynamo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingDynamoWithDefaults

`func NewLoggingDynamoWithDefaults() *LoggingDynamo`

NewLoggingDynamoWithDefaults instantiates a new LoggingDynamo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *LoggingDynamo) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *LoggingDynamo) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *LoggingDynamo) SetService(v string)`

SetService sets Service field to given value.


### GetTime

`func (o *LoggingDynamo) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LoggingDynamo) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LoggingDynamo) SetTime(v string)`

SetTime sets Time field to given value.


### GetLog

`func (o *LoggingDynamo) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *LoggingDynamo) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *LoggingDynamo) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.


### GetSeverity

`func (o *LoggingDynamo) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *LoggingDynamo) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *LoggingDynamo) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSubject

`func (o *LoggingDynamo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LoggingDynamo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LoggingDynamo) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetServiceName

`func (o *LoggingDynamo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *LoggingDynamo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *LoggingDynamo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


