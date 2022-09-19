# \LoggingApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubjectLogs**](LoggingApi.md#GetSubjectLogs) | **Get** /sam/logging/logs/services/{service}/subjects/{subject} | Getsubjectlogs



## GetSubjectLogs

> []LoggingDynamo GetSubjectLogs(ctx, service, subject).DateStart(dateStart).DateEnd(dateEnd).Desc(desc).Limit(limit).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getsubjectlogs

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
    subject := "subject_example" // string | 
    dateStart := "dateStart_example" // string |  (optional)
    dateEnd := "dateEnd_example" // string |  (optional)
    desc := true // bool |  (optional) (default to true)
    limit := int32(56) // int32 |  (optional) (default to 50)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoggingApi.GetSubjectLogs(context.Background(), service, subject).DateStart(dateStart).DateEnd(dateEnd).Desc(desc).Limit(limit).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingApi.GetSubjectLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjectLogs`: []LoggingDynamo
    fmt.Fprintf(os.Stdout, "Response from `LoggingApi.GetSubjectLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**subject** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateStart** | **string** |  | 
 **dateEnd** | **string** |  | 
 **desc** | **bool** |  | [default to true]
 **limit** | **int32** |  | [default to 50]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**[]LoggingDynamo**](LoggingDynamo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

