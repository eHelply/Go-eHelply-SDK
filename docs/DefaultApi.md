# \DefaultApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachEntityToAppointment**](DefaultApi.md#AttachEntityToAppointment) | **Post** /appointments/appointments/{appointment_uuid}/entities/{entity_uuid} | Attach Entity To Appointment
[**CreateAppointment**](DefaultApi.md#CreateAppointment) | **Post** /appointments/appointments | Create Appointment
[**DeleteAppointment**](DefaultApi.md#DeleteAppointment) | **Delete** /appointments/appointments/{appointment_uuid} | Delete Appointment
[**DetachEntityFromAppointment**](DefaultApi.md#DetachEntityFromAppointment) | **Delete** /appointments/appointments/{appointment_uuid}/entities/{entity_uuid} | Detach Entity From Appointment
[**GetAppointment**](DefaultApi.md#GetAppointment) | **Get** /appointments/appointments/{appointment_uuid} | Get Appointment
[**SearchAppointment**](DefaultApi.md#SearchAppointment) | **Get** /appointments/appointments | Search Appointment
[**SearchAppointmentEntities**](DefaultApi.md#SearchAppointmentEntities) | **Get** /appointments/appointments/{appointment_uuid}/entities | Search Appointment Entities
[**SearchEntityAppointments**](DefaultApi.md#SearchEntityAppointments) | **Get** /appointments/appointments/entities/{entity_uuid}/appointments | Get Entities Appointments
[**UpdateAppointment**](DefaultApi.md#UpdateAppointment) | **Put** /appointments/appointments/{appointment_uuid} | Update Appointment



## AttachEntityToAppointment

> bool AttachEntityToAppointment(ctx, appointmentUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Attach Entity To Appointment

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
    appointmentUuid := "appointmentUuid_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AttachEntityToAppointment(context.Background(), appointmentUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AttachEntityToAppointment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachEntityToAppointment`: bool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AttachEntityToAppointment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentUuid** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachEntityToAppointmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppointment

> AppointmentResponse CreateAppointment(ctx).AppointmentBase(appointmentBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Create Appointment

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
    appointmentBase := *openapiclient.NewAppointmentBase() // AppointmentBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAppointment(context.Background()).AppointmentBase(appointmentBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAppointment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppointment`: AppointmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAppointment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppointmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appointmentBase** | [**AppointmentBase**](AppointmentBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**AppointmentResponse**](AppointmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppointment

> bool DeleteAppointment(ctx, appointmentUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Delete Appointment

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
    appointmentUuid := "appointmentUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteAppointment(context.Background(), appointmentUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAppointment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppointment`: bool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAppointment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppointmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachEntityFromAppointment

> bool DetachEntityFromAppointment(ctx, appointmentUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Detach Entity From Appointment

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
    appointmentUuid := "appointmentUuid_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DetachEntityFromAppointment(context.Background(), appointmentUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DetachEntityFromAppointment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachEntityFromAppointment`: bool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DetachEntityFromAppointment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentUuid** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachEntityFromAppointmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppointment

> AppointmentResponse GetAppointment(ctx, appointmentUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Get Appointment

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
    appointmentUuid := "appointmentUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAppointment(context.Background(), appointmentUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAppointment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppointment`: AppointmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAppointment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppointmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**AppointmentResponse**](AppointmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppointment

> interface{} SearchAppointment(ctx).PlaceUuid(placeUuid).ExcludeCancelled(excludeCancelled).IsDeleted(isDeleted).StartRange(startRange).EndRange(endRange).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).Search(search).SearchOn(searchOn).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Search Appointment

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
    placeUuid := "placeUuid_example" // string |  (optional)
    excludeCancelled := true // bool |  (optional) (default to false)
    isDeleted := true // bool |  (optional) (default to false)
    startRange := "startRange_example" // string |  (optional)
    endRange := "endRange_example" // string |  (optional)
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    sortOn := "sortOn_example" // string |  (optional)
    sortDesc := true // bool |  (optional) (default to false)
    search := "search_example" // string |  (optional)
    searchOn := "searchOn_example" // string |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchAppointment(context.Background()).PlaceUuid(placeUuid).ExcludeCancelled(excludeCancelled).IsDeleted(isDeleted).StartRange(startRange).EndRange(endRange).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).Search(search).SearchOn(searchOn).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchAppointment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAppointment`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchAppointment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppointmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placeUuid** | **string** |  | 
 **excludeCancelled** | **bool** |  | [default to false]
 **isDeleted** | **bool** |  | [default to false]
 **startRange** | **string** |  | 
 **endRange** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **sortOn** | **string** |  | 
 **sortDesc** | **bool** |  | [default to false]
 **search** | **string** |  | 
 **searchOn** | **string** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppointmentEntities

> interface{} SearchAppointmentEntities(ctx, appointmentUuid).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).Search(search).SearchOn(searchOn).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Search Appointment Entities

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
    appointmentUuid := "appointmentUuid_example" // string | 
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    sortOn := "sortOn_example" // string |  (optional)
    sortDesc := true // bool |  (optional) (default to false)
    search := "search_example" // string |  (optional)
    searchOn := "searchOn_example" // string |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchAppointmentEntities(context.Background(), appointmentUuid).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).Search(search).SearchOn(searchOn).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchAppointmentEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAppointmentEntities`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchAppointmentEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppointmentEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **sortOn** | **string** |  | 
 **sortDesc** | **bool** |  | [default to false]
 **search** | **string** |  | 
 **searchOn** | **string** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEntityAppointments

> interface{} SearchEntityAppointments(ctx, entityUuid).StartDate(startDate).EndDate(endDate).IncludeCancelled(includeCancelled).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Get Entities Appointments

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
    entityUuid := "entityUuid_example" // string | 
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    includeCancelled := true // bool |  (optional) (default to false)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchEntityAppointments(context.Background(), entityUuid).StartDate(startDate).EndDate(endDate).IncludeCancelled(includeCancelled).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchEntityAppointments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEntityAppointments`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchEntityAppointments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchEntityAppointmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **includeCancelled** | **bool** |  | [default to false]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppointment

> AppointmentResponse UpdateAppointment(ctx, appointmentUuid).AppointmentBase(appointmentBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Update Appointment

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
    appointmentUuid := "appointmentUuid_example" // string | 
    appointmentBase := *openapiclient.NewAppointmentBase() // AppointmentBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAppointment(context.Background(), appointmentUuid).AppointmentBase(appointmentBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAppointment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppointment`: AppointmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAppointment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppointmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appointmentBase** | [**AppointmentBase**](AppointmentBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**AppointmentResponse**](AppointmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

