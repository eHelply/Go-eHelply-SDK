# \CompaniesApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompanyPlacesCompaniesPost**](CompaniesApi.md#CreateCompanyPlacesCompaniesPost) | **Post** /places/companies | Create Company
[**DeleteCompanyPlacesCompaniesCompanyUuidDelete**](CompaniesApi.md#DeleteCompanyPlacesCompaniesCompanyUuidDelete) | **Delete** /places/companies/{company_uuid} | Delete Company
[**GetCompanyPlacesCompaniesCompanyUuidGet**](CompaniesApi.md#GetCompanyPlacesCompaniesCompanyUuidGet) | **Get** /places/companies/{company_uuid} | Get Company
[**SearchCompaniesPlacesCompaniesGet**](CompaniesApi.md#SearchCompaniesPlacesCompaniesGet) | **Get** /places/companies | Search Companies
[**UpdateCompanyPlacesCompaniesCompanyUuidPut**](CompaniesApi.md#UpdateCompanyPlacesCompaniesCompanyUuidPut) | **Put** /places/companies/{company_uuid} | Update Company



## CreateCompanyPlacesCompaniesPost

> CompanyResponse CreateCompanyPlacesCompaniesPost(ctx).CompanyBase(companyBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Create Company



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
    companyBase := *openapiclient.NewCompanyBase() // CompanyBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.CreateCompanyPlacesCompaniesPost(context.Background()).CompanyBase(companyBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.CreateCompanyPlacesCompaniesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompanyPlacesCompaniesPost`: CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.CreateCompanyPlacesCompaniesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyPlacesCompaniesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyBase** | [**CompanyBase**](CompanyBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**CompanyResponse**](CompanyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyPlacesCompaniesCompanyUuidDelete

> interface{} DeleteCompanyPlacesCompaniesCompanyUuidDelete(ctx, companyUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Delete Company



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
    companyUuid := "companyUuid_example" // string | 
    softDelete := true // bool |  (optional) (default to true)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.DeleteCompanyPlacesCompaniesCompanyUuidDelete(context.Background(), companyUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.DeleteCompanyPlacesCompaniesCompanyUuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCompanyPlacesCompaniesCompanyUuidDelete`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.DeleteCompanyPlacesCompaniesCompanyUuidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyPlacesCompaniesCompanyUuidDeleteRequest struct via the builder pattern


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


## GetCompanyPlacesCompaniesCompanyUuidGet

> CompanyResponse GetCompanyPlacesCompaniesCompanyUuidGet(ctx, companyUuid).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).WithPlaces(withPlaces).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Get Company



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
    companyUuid := "companyUuid_example" // string | 
    withMeta := true // bool |  (optional) (default to false)
    withCatalog := true // bool |  (optional) (default to false)
    withReviews := true // bool |  (optional) (default to false)
    withSchedule := true // bool |  (optional) (default to false)
    withBlog := true // bool |  (optional) (default to false)
    withTags := true // bool |  (optional) (default to false)
    withCategories := true // bool |  (optional) (default to false)
    withPlaces := true // bool |  (optional) (default to false)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.GetCompanyPlacesCompaniesCompanyUuidGet(context.Background(), companyUuid).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).WithPlaces(withPlaces).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetCompanyPlacesCompaniesCompanyUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyPlacesCompaniesCompanyUuidGet`: CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetCompanyPlacesCompaniesCompanyUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPlacesCompaniesCompanyUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withMeta** | **bool** |  | [default to false]
 **withCatalog** | **bool** |  | [default to false]
 **withReviews** | **bool** |  | [default to false]
 **withSchedule** | **bool** |  | [default to false]
 **withBlog** | **bool** |  | [default to false]
 **withTags** | **bool** |  | [default to false]
 **withCategories** | **bool** |  | [default to false]
 **withPlaces** | **bool** |  | [default to false]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**CompanyResponse**](CompanyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCompaniesPlacesCompaniesGet

> Page SearchCompaniesPlacesCompaniesGet(ctx).ProjectUuid(projectUuid).Name(name).Email(email).IsPublic(isPublic).IsDeleted(isDeleted).WithPlaces(withPlaces).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Search Companies



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
    projectUuid := "projectUuid_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    isPublic := true // bool |  (optional) (default to true)
    isDeleted := true // bool |  (optional) (default to false)
    withPlaces := true // bool |  (optional) (default to false)
    withMeta := true // bool |  (optional) (default to false)
    withCatalog := true // bool |  (optional) (default to false)
    withReviews := true // bool |  (optional) (default to false)
    withSchedule := true // bool |  (optional) (default to false)
    withBlog := true // bool |  (optional) (default to false)
    withTags := true // bool |  (optional) (default to false)
    withCategories := true // bool |  (optional) (default to false)
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    sortOn := "sortOn_example" // string |  (optional)
    sortDesc := true // bool |  (optional) (default to false)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.SearchCompaniesPlacesCompaniesGet(context.Background()).ProjectUuid(projectUuid).Name(name).Email(email).IsPublic(isPublic).IsDeleted(isDeleted).WithPlaces(withPlaces).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.SearchCompaniesPlacesCompaniesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCompaniesPlacesCompaniesGet`: Page
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.SearchCompaniesPlacesCompaniesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCompaniesPlacesCompaniesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectUuid** | **string** |  | 
 **name** | **string** |  | 
 **email** | **string** |  | 
 **isPublic** | **bool** |  | [default to true]
 **isDeleted** | **bool** |  | [default to false]
 **withPlaces** | **bool** |  | [default to false]
 **withMeta** | **bool** |  | [default to false]
 **withCatalog** | **bool** |  | [default to false]
 **withReviews** | **bool** |  | [default to false]
 **withSchedule** | **bool** |  | [default to false]
 **withBlog** | **bool** |  | [default to false]
 **withTags** | **bool** |  | [default to false]
 **withCategories** | **bool** |  | [default to false]
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **sortOn** | **string** |  | 
 **sortDesc** | **bool** |  | [default to false]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

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


## UpdateCompanyPlacesCompaniesCompanyUuidPut

> CompanyResponse UpdateCompanyPlacesCompaniesCompanyUuidPut(ctx, companyUuid).CompanyBase(companyBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Update Company



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
    companyUuid := "companyUuid_example" // string | 
    companyBase := *openapiclient.NewCompanyBase() // CompanyBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.UpdateCompanyPlacesCompaniesCompanyUuidPut(context.Background(), companyUuid).CompanyBase(companyBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.UpdateCompanyPlacesCompaniesCompanyUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompanyPlacesCompaniesCompanyUuidPut`: CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.UpdateCompanyPlacesCompaniesCompanyUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyBase** | [**CompanyBase**](CompanyBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**CompanyResponse**](CompanyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

