# \MetaApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateField**](MetaApi.md#CreateField) | **Post** /meta/field | Create Field
[**CreateMeta**](MetaApi.md#CreateMeta) | **Post** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Create Meta
[**DeleteField**](MetaApi.md#DeleteField) | **Delete** /meta/field/{field_uuid} | Delete Field
[**DeleteMeta**](MetaApi.md#DeleteMeta) | **Delete** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Delete Meta
[**DeleteMetaFromUuid**](MetaApi.md#DeleteMetaFromUuid) | **Delete** /meta/meta/{meta_uuid} | Delete Meta From Uuid
[**GetField**](MetaApi.md#GetField) | **Get** /meta/field/{field_uuid} | Get Field
[**GetMeta**](MetaApi.md#GetMeta) | **Get** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Get Meta
[**GetMetaFromUuid**](MetaApi.md#GetMetaFromUuid) | **Get** /meta/meta/{meta_uuid} | Get Meta From Uuid
[**MakeSlug**](MetaApi.md#MakeSlug) | **Post** /meta/meta/slug | Make Slug
[**TouchMeta**](MetaApi.md#TouchMeta) | **Post** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid}/touch | Touch Meta
[**UpdateField**](MetaApi.md#UpdateField) | **Put** /meta/field/{field_uuid} | Update Field
[**UpdateMeta**](MetaApi.md#UpdateMeta) | **Put** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Update Meta
[**UpdateMetaFromUuid**](MetaApi.md#UpdateMetaFromUuid) | **Put** /meta/meta/{meta_uuid} | Update Meta From Uuid



## CreateField

> FieldDynamo CreateField(ctx).Field(field).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Create Field

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
    field := *openapiclient.NewField() // Field | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.CreateField(context.Background()).Field(field).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.CreateField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateField`: FieldDynamo
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.CreateField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | [**Field**](Field.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**FieldDynamo**](FieldDynamo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMeta

> MetaDynamo CreateMeta(ctx, service, typeStr, entityUuid).MetaCreate(metaCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Create Meta

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
    typeStr := "typeStr_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    metaCreate := *openapiclient.NewMetaCreate() // MetaCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.CreateMeta(context.Background(), service, typeStr, entityUuid).MetaCreate(metaCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.CreateMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMeta`: MetaDynamo
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.CreateMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**typeStr** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metaCreate** | [**MetaCreate**](MetaCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**MetaDynamo**](MetaDynamo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteField

> interface{} DeleteField(ctx, fieldUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Delete Field

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
    fieldUuid := "fieldUuid_example" // string | 
    softDelete := true // bool |  (optional) (default to true)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.DeleteField(context.Background(), fieldUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.DeleteField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteField`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.DeleteField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softDelete** | **bool** |  | [default to true]
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


## DeleteMeta

> interface{} DeleteMeta(ctx, service, typeStr, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Delete Meta

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
    typeStr := "typeStr_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.DeleteMeta(context.Background(), service, typeStr, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.DeleteMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMeta`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.DeleteMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**typeStr** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteMetaFromUuid

> interface{} DeleteMetaFromUuid(ctx, metaUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Delete Meta From Uuid

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
    metaUuid := "metaUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.DeleteMetaFromUuid(context.Background(), metaUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.DeleteMetaFromUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMetaFromUuid`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.DeleteMetaFromUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metaUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetaFromUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetField

> FieldDynamo GetField(ctx, fieldUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Get Field

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
    fieldUuid := "fieldUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.GetField(context.Background(), fieldUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.GetField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetField`: FieldDynamo
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.GetField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**FieldDynamo**](FieldDynamo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeta

> MetaGet GetMeta(ctx, service, typeStr, entityUuid).Detailed(detailed).Custom(custom).Dates(dates).History(history).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Get Meta

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
    typeStr := "typeStr_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    detailed := true // bool |  (optional) (default to false)
    custom := true // bool |  (optional) (default to false)
    dates := true // bool |  (optional) (default to false)
    history := int32(56) // int32 |  (optional) (default to 0)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.GetMeta(context.Background(), service, typeStr, entityUuid).Detailed(detailed).Custom(custom).Dates(dates).History(history).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.GetMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMeta`: MetaGet
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.GetMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**typeStr** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **detailed** | **bool** |  | [default to false]
 **custom** | **bool** |  | [default to false]
 **dates** | **bool** |  | [default to false]
 **history** | **int32** |  | [default to 0]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**MetaGet**](MetaGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetaFromUuid

> MetaGet GetMetaFromUuid(ctx, metaUuid).Detailed(detailed).Custom(custom).Dates(dates).History(history).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Get Meta From Uuid

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
    metaUuid := "metaUuid_example" // string | 
    detailed := true // bool |  (optional) (default to false)
    custom := true // bool |  (optional) (default to false)
    dates := true // bool |  (optional) (default to false)
    history := int32(56) // int32 |  (optional) (default to 0)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.GetMetaFromUuid(context.Background(), metaUuid).Detailed(detailed).Custom(custom).Dates(dates).History(history).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.GetMetaFromUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaFromUuid`: MetaGet
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.GetMetaFromUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metaUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaFromUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detailed** | **bool** |  | [default to false]
 **custom** | **bool** |  | [default to false]
 **dates** | **bool** |  | [default to false]
 **history** | **int32** |  | [default to 0]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**MetaGet**](MetaGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeSlug

> interface{} MakeSlug(ctx).MetaSlugger(metaSlugger).Execute()

Make Slug

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
    metaSlugger := *openapiclient.NewMetaSlugger("Name_example") // MetaSlugger | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.MakeSlug(context.Background()).MetaSlugger(metaSlugger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.MakeSlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeSlug`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.MakeSlug`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metaSlugger** | [**MetaSlugger**](MetaSlugger.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TouchMeta

> MetaDynamo TouchMeta(ctx, service, typeStr, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Touch Meta

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
    typeStr := "typeStr_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.TouchMeta(context.Background(), service, typeStr, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.TouchMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TouchMeta`: MetaDynamo
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.TouchMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**typeStr** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTouchMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**MetaDynamo**](MetaDynamo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateField

> interface{} UpdateField(ctx, fieldUuid).Field(field).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Update Field

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
    fieldUuid := "fieldUuid_example" // string | 
    field := *openapiclient.NewField() // Field | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.UpdateField(context.Background(), fieldUuid).Field(field).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.UpdateField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateField`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.UpdateField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | [**Field**](Field.md) |  | 
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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeta

> MetaDynamo UpdateMeta(ctx, service, typeStr, entityUuid).MetaCreate(metaCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Update Meta

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
    typeStr := "typeStr_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    metaCreate := *openapiclient.NewMetaCreate() // MetaCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.UpdateMeta(context.Background(), service, typeStr, entityUuid).MetaCreate(metaCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.UpdateMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMeta`: MetaDynamo
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.UpdateMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** |  | 
**typeStr** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metaCreate** | [**MetaCreate**](MetaCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**MetaDynamo**](MetaDynamo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetaFromUuid

> MetaDynamo UpdateMetaFromUuid(ctx, metaUuid).MetaCreate(metaCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Update Meta From Uuid

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
    metaUuid := "metaUuid_example" // string | 
    metaCreate := *openapiclient.NewMetaCreate() // MetaCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaApi.UpdateMetaFromUuid(context.Background(), metaUuid).MetaCreate(metaCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaApi.UpdateMetaFromUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetaFromUuid`: MetaDynamo
    fmt.Fprintf(os.Stdout, "Response from `MetaApi.UpdateMetaFromUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metaUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetaFromUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metaCreate** | [**MetaCreate**](MetaCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**MetaDynamo**](MetaDynamo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

