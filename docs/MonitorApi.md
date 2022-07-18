# \MonitorApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcknowledgeAlarm**](MonitorApi.md#AcknowledgeAlarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/acknowledge | Acknowledgealarm
[**AssignAlarm**](MonitorApi.md#AssignAlarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/assign | Assignalarm
[**AttachAlarmNote**](MonitorApi.md#AttachAlarmNote) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/note | Attachalarmnote
[**AttachAlarmTicket**](MonitorApi.md#AttachAlarmTicket) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/ticket | Attachalarmticket
[**ClearAlarm**](MonitorApi.md#ClearAlarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/clear | Clearalarm
[**GetService**](MonitorApi.md#GetService) | **Get** /sam/monitor/services/{service} | Getservice
[**GetServiceAlarm**](MonitorApi.md#GetServiceAlarm) | **Get** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid} | Getservicealarm
[**GetServiceAlarms**](MonitorApi.md#GetServiceAlarms) | **Get** /sam/monitor/services/{service}/stages/{stage}/alarms | Getservicealarms
[**GetServiceHeartbeat**](MonitorApi.md#GetServiceHeartbeat) | **Get** /sam/monitor/services/{service}/stages/{stage}/heartbeats | Getserviceheartbeat
[**GetServiceKpis**](MonitorApi.md#GetServiceKpis) | **Get** /sam/monitor/services/{service}/kpis | Getservicekpis
[**GetServiceSpec**](MonitorApi.md#GetServiceSpec) | **Get** /sam/monitor/services/{service}/specs/{spec} | Getservicespec
[**GetServiceSpecs**](MonitorApi.md#GetServiceSpecs) | **Get** /sam/monitor/services/{service}/specs | Getservicespecs
[**GetServiceVitals**](MonitorApi.md#GetServiceVitals) | **Get** /sam/monitor/services/{service}/stages/{stage}/vitals | Getservicevitals
[**GetServices**](MonitorApi.md#GetServices) | **Get** /sam/monitor/services | Getservices
[**GetServicesWithSpecs**](MonitorApi.md#GetServicesWithSpecs) | **Get** /sam/monitor/specs/services | Getserviceswithspecs
[**HideService**](MonitorApi.md#HideService) | **Post** /sam/monitor/services/{service}/stages/{stage}/hide | Hideservice
[**IgnoreAlarm**](MonitorApi.md#IgnoreAlarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/ignore | Ignorealarm
[**RegisterService**](MonitorApi.md#RegisterService) | **Post** /sam/monitor/services | Registerservice
[**SearchAlarms**](MonitorApi.md#SearchAlarms) | **Get** /sam/monitor/services/{service}/alarms | Searchalarms
[**ShowService**](MonitorApi.md#ShowService) | **Post** /sam/monitor/services/{service}/stages/{stage}/show | Showservice
[**TerminateAlarm**](MonitorApi.md#TerminateAlarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/terminate | Terminatealarm
[**TriggerAlarm**](MonitorApi.md#TriggerAlarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms | Triggeralarm



## AcknowledgeAlarm

> AlarmResponse AcknowledgeAlarm(ctx, service, stage, alarmUuid).AlarmAcknowledge(alarmAcknowledge).Execute()

Acknowledgealarm

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 
    alarmAcknowledge := *openapiclient.NewAlarmAcknowledge("AcknowledgerUuid_example") // AlarmAcknowledge | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.AcknowledgeAlarm(context.Background(), service, stage, alarmUuid).AlarmAcknowledge(alarmAcknowledge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.AcknowledgeAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcknowledgeAlarm`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.AcknowledgeAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alarmAcknowledge** | [**AlarmAcknowledge**](AlarmAcknowledge.md) |  | 

### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignAlarm

> AlarmResponse AssignAlarm(ctx, service, stage, alarmUuid).AlarmAssign(alarmAssign).Execute()

Assignalarm

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 
    alarmAssign := *openapiclient.NewAlarmAssign("AssigneeUuid_example") // AlarmAssign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.AssignAlarm(context.Background(), service, stage, alarmUuid).AlarmAssign(alarmAssign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.AssignAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignAlarm`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.AssignAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alarmAssign** | [**AlarmAssign**](AlarmAssign.md) |  | 

### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachAlarmNote

> AlarmResponse AttachAlarmNote(ctx, service, stage, alarmUuid).AlarmNote(alarmNote).Execute()

Attachalarmnote

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 
    alarmNote := *openapiclient.NewAlarmNote("AuthorUuid_example", "Message_example") // AlarmNote | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.AttachAlarmNote(context.Background(), service, stage, alarmUuid).AlarmNote(alarmNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.AttachAlarmNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachAlarmNote`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.AttachAlarmNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachAlarmNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alarmNote** | [**AlarmNote**](AlarmNote.md) |  | 

### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachAlarmTicket

> AlarmResponse AttachAlarmTicket(ctx, service, stage, alarmUuid).AlarmTicket(alarmTicket).Execute()

Attachalarmticket

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 
    alarmTicket := *openapiclient.NewAlarmTicket("TicketUuid_example") // AlarmTicket | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.AttachAlarmTicket(context.Background(), service, stage, alarmUuid).AlarmTicket(alarmTicket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.AttachAlarmTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachAlarmTicket`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.AttachAlarmTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachAlarmTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alarmTicket** | [**AlarmTicket**](AlarmTicket.md) |  | 

### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearAlarm

> AlarmResponse ClearAlarm(ctx, service, stage, alarmUuid).Execute()

Clearalarm

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.ClearAlarm(context.Background(), service, stage, alarmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ClearAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearAlarm`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ClearAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> ServiceResponse GetService(ctx, service).Heartbeats(heartbeats).HeartbeatLimit(heartbeatLimit).Alarms(alarms).AlarmLimit(alarmLimit).Stage(stage).Execute()

Getservice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    heartbeats := true // bool |  (optional) (default to false)
    heartbeatLimit := int32(56) // int32 |  (optional) (default to 5)
    alarms := true // bool |  (optional) (default to false)
    alarmLimit := int32(56) // int32 |  (optional) (default to 5)
    stage := "stage_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetService(context.Background(), service).Heartbeats(heartbeats).HeartbeatLimit(heartbeatLimit).Alarms(alarms).AlarmLimit(alarmLimit).Stage(stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: ServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **heartbeats** | **bool** |  | [default to false]
 **heartbeatLimit** | **int32** |  | [default to 5]
 **alarms** | **bool** |  | [default to false]
 **alarmLimit** | **int32** |  | [default to 5]
 **stage** | **string** |  | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAlarm

> AlarmResponse GetServiceAlarm(ctx, service, stage, alarmUuid).Execute()

Getservicealarm

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServiceAlarm(context.Background(), service, stage, alarmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServiceAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAlarm`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServiceAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAlarms

> []AlarmResponse GetServiceAlarms(ctx, service, stage).History(history).IncludeTerminated(includeTerminated).IncludeCleared(includeCleared).Execute()

Getservicealarms

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    history := int32(56) // int32 |  (optional) (default to 5)
    includeTerminated := true // bool |  (optional) (default to false)
    includeCleared := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServiceAlarms(context.Background(), service, stage).History(history).IncludeTerminated(includeTerminated).IncludeCleared(includeCleared).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServiceAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAlarms`: []AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServiceAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **history** | **int32** |  | [default to 5]
 **includeTerminated** | **bool** |  | [default to false]
 **includeCleared** | **bool** |  | [default to false]

### Return type

[**[]AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceHeartbeat

> []HeartbeatResponse GetServiceHeartbeat(ctx, service, stage).History(history).Execute()

Getserviceheartbeat

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    history := int32(56) // int32 |  (optional) (default to 5)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServiceHeartbeat(context.Background(), service, stage).History(history).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServiceHeartbeat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceHeartbeat`: []HeartbeatResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServiceHeartbeat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceHeartbeatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **history** | **int32** |  | [default to 5]

### Return type

[**[]HeartbeatResponse**](HeartbeatResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKpis

> []KpiResponse GetServiceKpis(ctx, service).History(history).Execute()

Getservicekpis

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    history := int32(56) // int32 |  (optional) (default to 5)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServiceKpis(context.Background(), service).History(history).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServiceKpis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceKpis`: []KpiResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServiceKpis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKpisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **history** | **int32** |  | [default to 5]

### Return type

[**[]KpiResponse**](KpiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSpec

> GetServiceSpecResponse GetServiceSpec(ctx, service, spec).Execute()

Getservicespec

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    spec := "spec_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServiceSpec(context.Background(), service, spec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServiceSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceSpec`: GetServiceSpecResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServiceSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**spec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetServiceSpecResponse**](GetServiceSpecResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSpecs

> GetServiceSpecsResponse GetServiceSpecs(ctx, service).Execute()

Getservicespecs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServiceSpecs(context.Background(), service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServiceSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceSpecs`: GetServiceSpecsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServiceSpecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceSpecsResponse**](GetServiceSpecsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceVitals

> []StatsVitalsResponse GetServiceVitals(ctx, service, stage).History(history).Execute()

Getservicevitals

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    history := int32(56) // int32 |  (optional) (default to 5)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServiceVitals(context.Background(), service, stage).History(history).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServiceVitals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceVitals`: []StatsVitalsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServiceVitals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceVitalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **history** | **int32** |  | [default to 5]

### Return type

[**[]StatsVitalsResponse**](StatsVitalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> []ServiceResponse GetServices(ctx).Heartbeats(heartbeats).HeartbeatLimit(heartbeatLimit).Alarms(alarms).AlarmLimit(alarmLimit).IncludeHidden(includeHidden).Stage(stage).Key(key).Execute()

Getservices

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    heartbeats := true // bool |  (optional) (default to false)
    heartbeatLimit := int32(56) // int32 |  (optional) (default to 5)
    alarms := true // bool |  (optional) (default to false)
    alarmLimit := int32(56) // int32 |  (optional) (default to 5)
    includeHidden := true // bool |  (optional) (default to false)
    stage := "stage_example" // string |  (optional)
    key := "key_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServices(context.Background()).Heartbeats(heartbeats).HeartbeatLimit(heartbeatLimit).Alarms(alarms).AlarmLimit(alarmLimit).IncludeHidden(includeHidden).Stage(stage).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: []ServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **heartbeats** | **bool** |  | [default to false]
 **heartbeatLimit** | **int32** |  | [default to 5]
 **alarms** | **bool** |  | [default to false]
 **alarmLimit** | **int32** |  | [default to 5]
 **includeHidden** | **bool** |  | [default to false]
 **stage** | **string** |  | 
 **key** | **string** |  | 

### Return type

[**[]ServiceResponse**](ServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesWithSpecs

> GetServiceServiceWithSpecsResponse GetServicesWithSpecs(ctx).Execute()

Getserviceswithspecs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.GetServicesWithSpecs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetServicesWithSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServicesWithSpecs`: GetServiceServiceWithSpecsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetServicesWithSpecs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesWithSpecsRequest struct via the builder pattern


### Return type

[**GetServiceServiceWithSpecsResponse**](GetServiceServiceWithSpecsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideService

> ServiceMessageResponse HideService(ctx, service, stage).Execute()

Hideservice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.HideService(context.Background(), service, stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.HideService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HideService`: ServiceMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.HideService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceMessageResponse**](ServiceMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IgnoreAlarm

> AlarmResponse IgnoreAlarm(ctx, service, stage, alarmUuid).AlarmIgnore(alarmIgnore).Execute()

Ignorealarm

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 
    alarmIgnore := *openapiclient.NewAlarmIgnore("IgnorerUuid_example") // AlarmIgnore | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.IgnoreAlarm(context.Background(), service, stage, alarmUuid).AlarmIgnore(alarmIgnore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.IgnoreAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IgnoreAlarm`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.IgnoreAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIgnoreAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alarmIgnore** | [**AlarmIgnore**](AlarmIgnore.md) |  | 

### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterService

> ServiceResponse RegisterService(ctx).ServiceCreate(serviceCreate).Execute()

Registerservice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceCreate := *openapiclient.NewServiceCreate("Name_example", "Key_example", "Version_example", "Summary_example", []string{"Authors_example"}, []string{"AuthorEmails_example"}) // ServiceCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.RegisterService(context.Background()).ServiceCreate(serviceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.RegisterService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterService`: ServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.RegisterService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceCreate** | [**ServiceCreate**](ServiceCreate.md) |  | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAlarms

> Page SearchAlarms(ctx, service).Page(page).PageSize(pageSize).Search(search).SearchOn(searchOn).SortOn(sortOn).SortDesc(sortDesc).Execute()

Searchalarms

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    search := "search_example" // string |  (optional)
    searchOn := "searchOn_example" // string |  (optional)
    sortOn := "sortOn_example" // string |  (optional)
    sortDesc := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.SearchAlarms(context.Background(), service).Page(page).PageSize(pageSize).Search(search).SearchOn(searchOn).SortOn(sortOn).SortDesc(sortDesc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.SearchAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAlarms`: Page
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.SearchAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **search** | **string** |  | 
 **searchOn** | **string** |  | 
 **sortOn** | **string** |  | 
 **sortDesc** | **bool** |  | [default to false]

### Return type

[**Page**](Page.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowService

> ServiceMessageResponse ShowService(ctx, service, stage).Execute()

Showservice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.ShowService(context.Background(), service, stage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ShowService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowService`: ServiceMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ShowService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceMessageResponse**](ServiceMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateAlarm

> AlarmResponse TerminateAlarm(ctx, service, stage, alarmUuid).AlarmTerminate(alarmTerminate).Execute()

Terminatealarm

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmUuid := "alarmUuid_example" // string | 
    alarmTerminate := *openapiclient.NewAlarmTerminate("TerminaterUuid_example") // AlarmTerminate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.TerminateAlarm(context.Background(), service, stage, alarmUuid).AlarmTerminate(alarmTerminate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.TerminateAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateAlarm`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.TerminateAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 
**alarmUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **alarmTerminate** | [**AlarmTerminate**](AlarmTerminate.md) |  | 

### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerAlarm

> AlarmResponse TriggerAlarm(ctx, service, stage).AlarmCreate(alarmCreate).Execute()

Triggeralarm

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    service := "service_example" // string | 
    stage := "stage_example" // string | 
    alarmCreate := *openapiclient.NewAlarmCreate() // AlarmCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorApi.TriggerAlarm(context.Background(), service, stage).AlarmCreate(alarmCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.TriggerAlarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerAlarm`: AlarmResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.TriggerAlarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**stage** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alarmCreate** | [**AlarmCreate**](AlarmCreate.md) |  | 

### Return type

[**AlarmResponse**](AlarmResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

