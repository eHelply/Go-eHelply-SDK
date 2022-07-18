# \SecurityApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEncryptionKey**](SecurityApi.md#CreateEncryptionKey) | **Post** /sam/security/encryption/categories/{category}/keys | Createencryptionkey
[**CreateKey**](SecurityApi.md#CreateKey) | **Post** /sam/security/keys | Createkey
[**DeleteKey**](SecurityApi.md#DeleteKey) | **Delete** /sam/security/keys/{key_uuid} | Deletekey
[**GenerateToken**](SecurityApi.md#GenerateToken) | **Post** /sam/security/tokens | Generatetoken
[**GetEncryptionKey**](SecurityApi.md#GetEncryptionKey) | **Get** /sam/security/encryption/categories/{category}/keys | Getencryptionkey
[**GetKey**](SecurityApi.md#GetKey) | **Get** /sam/security/keys/{key_uuid} | Getkey
[**SearchKeys**](SecurityApi.md#SearchKeys) | **Get** /sam/security/keys | Searchkeys
[**VerifyKey**](SecurityApi.md#VerifyKey) | **Post** /sam/security/keys/verify | Verifykey



## CreateEncryptionKey

> SecurityEncryptionKeyResponse CreateEncryptionKey(ctx, category).EhelplySecuritySecretKey(ehelplySecuritySecretKey).Execute()

Createencryptionkey

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
    category := "category_example" // string | 
    ehelplySecuritySecretKey := "ehelplySecuritySecretKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.CreateEncryptionKey(context.Background(), category).EhelplySecuritySecretKey(ehelplySecuritySecretKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.CreateEncryptionKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEncryptionKey`: SecurityEncryptionKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.CreateEncryptionKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEncryptionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ehelplySecuritySecretKey** | **string** |  | 

### Return type

[**SecurityEncryptionKeyResponse**](SecurityEncryptionKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKey

> ResponseCreatekey CreateKey(ctx).SecurityKeyCreate(securityKeyCreate).AccessLength(accessLength).SecretLength(secretLength).Execute()

Createkey

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
    securityKeyCreate := *openapiclient.NewSecurityKeyCreate("Name_example", "Summary_example") // SecurityKeyCreate | 
    accessLength := int32(56) // int32 |  (optional) (default to 64)
    secretLength := int32(56) // int32 |  (optional) (default to 64)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.CreateKey(context.Background()).SecurityKeyCreate(securityKeyCreate).AccessLength(accessLength).SecretLength(secretLength).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.CreateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKey`: ResponseCreatekey
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.CreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityKeyCreate** | [**SecurityKeyCreate**](SecurityKeyCreate.md) |  | 
 **accessLength** | **int32** |  | [default to 64]
 **secretLength** | **int32** |  | [default to 64]

### Return type

[**ResponseCreatekey**](ResponseCreatekey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> ResponseDeletekey DeleteKey(ctx, keyUuid).Execute()

Deletekey

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
    keyUuid := "keyUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.DeleteKey(context.Background(), keyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.DeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKey`: ResponseDeletekey
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.DeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeletekey**](ResponseDeletekey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateToken

> ResponseGeneratetoken GenerateToken(ctx).SecurityCreateToken(securityCreateToken).Execute()

Generatetoken

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
    securityCreateToken := *openapiclient.NewSecurityCreateToken() // SecurityCreateToken | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GenerateToken(context.Background()).SecurityCreateToken(securityCreateToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GenerateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateToken`: ResponseGeneratetoken
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityCreateToken** | [**SecurityCreateToken**](SecurityCreateToken.md) |  | 

### Return type

[**ResponseGeneratetoken**](ResponseGeneratetoken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionKey

> []SecurityEncryptionKeyGet GetEncryptionKey(ctx, category).EhelplySecuritySecretKey(ehelplySecuritySecretKey).Execute()

Getencryptionkey

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
    category := "category_example" // string | 
    ehelplySecuritySecretKey := "ehelplySecuritySecretKey_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetEncryptionKey(context.Background(), category).EhelplySecuritySecretKey(ehelplySecuritySecretKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetEncryptionKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEncryptionKey`: []SecurityEncryptionKeyGet
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetEncryptionKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ehelplySecuritySecretKey** | **string** |  | 

### Return type

[**[]SecurityEncryptionKeyGet**](SecurityEncryptionKeyGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKey

> SecurityKeyGet GetKey(ctx, keyUuid).Execute()

Getkey

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
    keyUuid := "keyUuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetKey(context.Background(), keyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKey`: SecurityKeyGet
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityKeyGet**](SecurityKeyGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchKeys

> []SecurityKeyGet SearchKeys(ctx).Execute()

Searchkeys

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
    resp, r, err := apiClient.SecurityApi.SearchKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.SearchKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchKeys`: []SecurityKeyGet
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.SearchKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchKeysRequest struct via the builder pattern


### Return type

[**[]SecurityKeyGet**](SecurityKeyGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyKey

> SecurityKeyGet VerifyKey(ctx).SecurityKeyVerify(securityKeyVerify).Execute()

Verifykey

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
    securityKeyVerify := *openapiclient.NewSecurityKeyVerify("Access_example", "Secret_example") // SecurityKeyVerify | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.VerifyKey(context.Background()).SecurityKeyVerify(securityKeyVerify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.VerifyKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyKey`: SecurityKeyGet
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.VerifyKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityKeyVerify** | [**SecurityKeyVerify**](SecurityKeyVerify.md) |  | 

### Return type

[**SecurityKeyGet**](SecurityKeyGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

