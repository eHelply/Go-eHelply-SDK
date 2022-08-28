# \BillingApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingAccount**](BillingApi.md#CreateBillingAccount) | **Post** /sam/billing/create_billing_account | Createbillingaccount
[**GetClientSecret**](BillingApi.md#GetClientSecret) | **Get** /sam/billing/retrieve_secret | Getclientsecret
[**HasPayment**](BillingApi.md#HasPayment) | **Get** /sam/billing/has_payment | Haspayment
[**ListPaymentMethods**](BillingApi.md#ListPaymentMethods) | **Get** /sam/billing/view_payment_method | Listpaymentmethods
[**ProcessPayment**](BillingApi.md#ProcessPayment) | **Post** /sam/billing/process_payment | Processpayment
[**ReconcilePaymentMethod**](BillingApi.md#ReconcilePaymentMethod) | **Get** /sam/billing/reconcile_payment | Reconcilepaymentmethod
[**RemovePaymentMethod**](BillingApi.md#RemovePaymentMethod) | **Delete** /sam/billing/remove_payment_method | Removepaymentmethod



## CreateBillingAccount

> StripeAccountResponse CreateBillingAccount(ctx).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createbillingaccount

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
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.CreateBillingAccount(context.Background()).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.CreateBillingAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBillingAccount`: StripeAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.CreateBillingAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**StripeAccountResponse**](StripeAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientSecret

> StripeCustomerSecretResponse GetClientSecret(ctx).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getclientsecret

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
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.GetClientSecret(context.Background()).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetClientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientSecret`: StripeCustomerSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetClientSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**StripeCustomerSecretResponse**](StripeCustomerSecretResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasPayment

> bool HasPayment(ctx).ProjectUuid(projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Haspayment

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
    projectUuid := TODO // interface{} |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.HasPayment(context.Background()).ProjectUuid(projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.HasPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasPayment`: bool
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.HasPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHasPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectUuid** | [**interface{}**](interface{}.md) |  | 
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


## ListPaymentMethods

> []PaymentMethodResponse ListPaymentMethods(ctx).ProjectUuid(projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Listpaymentmethods

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
    projectUuid := TODO // interface{} |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.ListPaymentMethods(context.Background()).ProjectUuid(projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentMethods`: []PaymentMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectUuid** | [**interface{}**](interface{}.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**[]PaymentMethodResponse**](PaymentMethodResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessPayment

> string ProcessPayment(ctx).Payment(payment).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Processpayment

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
    payment := *openapiclient.NewPayment(int32(123)) // Payment | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.ProcessPayment(context.Background()).Payment(payment).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ProcessPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessPayment`: string
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ProcessPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payment** | [**Payment**](Payment.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReconcilePaymentMethod

> bool ReconcilePaymentMethod(ctx).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Reconcilepaymentmethod

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
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.ReconcilePaymentMethod(context.Background()).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ReconcilePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReconcilePaymentMethod`: bool
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ReconcilePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReconcilePaymentMethodRequest struct via the builder pattern


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


## RemovePaymentMethod

> string RemovePaymentMethod(ctx).ProjectUuid(projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Removepaymentmethod

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
    projectUuid := TODO // interface{} |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.RemovePaymentMethod(context.Background()).ProjectUuid(projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.RemovePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePaymentMethod`: string
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.RemovePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemovePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectUuid** | [**interface{}**](interface{}.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

