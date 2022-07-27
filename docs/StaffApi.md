# \StaffApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStaff**](StaffApi.md#CreateStaff) | **Post** /places/staff | Createstaff
[**DeleteStaff**](StaffApi.md#DeleteStaff) | **Delete** /places/staff/{staff_uuid} | Deletestaff
[**GetStaff**](StaffApi.md#GetStaff) | **Get** /places/staff/{staff_uuid} | Getstaff
[**SearchStaff**](StaffApi.md#SearchStaff) | **Get** /places/staff | Searchstaff
[**UpdateStaff**](StaffApi.md#UpdateStaff) | **Put** /places/staff/{staff_uuid} | Updatestaff



## CreateStaff

> StaffDb CreateStaff(ctx).StaffCreate(staffCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createstaff



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
    staffCreate := *openapiclient.NewStaffCreate("entity_uuid_1234") // StaffCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaffApi.CreateStaff(context.Background()).StaffCreate(staffCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaffApi.CreateStaff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStaff`: StaffDb
    fmt.Fprintf(os.Stdout, "Response from `StaffApi.CreateStaff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStaffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **staffCreate** | [**StaffCreate**](StaffCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**StaffDb**](StaffDb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStaff

> interface{} DeleteStaff(ctx, staffUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Deletestaff



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
    staffUuid := "staffUuid_example" // string | 
    softDelete := true // bool |  (optional) (default to true)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaffApi.DeleteStaff(context.Background(), staffUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaffApi.DeleteStaff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStaff`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaffApi.DeleteStaff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**staffUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStaffRequest struct via the builder pattern


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


## GetStaff

> StaffResponse GetStaff(ctx, staffUuid).WithPlaces(withPlaces).WithCompanies(withCompanies).WithCatalog(withCatalog).WithSchedule(withSchedule).WithRoles(withRoles).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getstaff



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
    staffUuid := "staffUuid_example" // string | 
    withPlaces := true // bool |  (optional) (default to false)
    withCompanies := true // bool |  (optional) (default to false)
    withCatalog := true // bool |  (optional) (default to false)
    withSchedule := true // bool |  (optional) (default to false)
    withRoles := true // bool |  (optional) (default to false)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaffApi.GetStaff(context.Background(), staffUuid).WithPlaces(withPlaces).WithCompanies(withCompanies).WithCatalog(withCatalog).WithSchedule(withSchedule).WithRoles(withRoles).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaffApi.GetStaff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStaff`: StaffResponse
    fmt.Fprintf(os.Stdout, "Response from `StaffApi.GetStaff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**staffUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withPlaces** | **bool** |  | [default to false]
 **withCompanies** | **bool** |  | [default to false]
 **withCatalog** | **bool** |  | [default to false]
 **withSchedule** | **bool** |  | [default to false]
 **withRoles** | **bool** |  | [default to false]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**StaffResponse**](StaffResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStaff

> Page SearchStaff(ctx).ProjectUuid(projectUuid).FirstName(firstName).LastName(lastName).IsDeleted(isDeleted).WithCompanies(withCompanies).WithPlaces(withPlaces).WithSchedule(withSchedule).WithCatalog(withCatalog).WithReviews(withReviews).WithRoles(withRoles).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Searchstaff



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
    firstName := "firstName_example" // string |  (optional)
    lastName := "lastName_example" // string |  (optional)
    isDeleted := true // bool |  (optional) (default to false)
    withCompanies := true // bool |  (optional) (default to false)
    withPlaces := true // bool |  (optional) (default to false)
    withSchedule := true // bool |  (optional) (default to false)
    withCatalog := true // bool |  (optional) (default to false)
    withReviews := true // bool |  (optional) (default to false)
    withRoles := true // bool |  (optional) (default to false)
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
    resp, r, err := apiClient.StaffApi.SearchStaff(context.Background()).ProjectUuid(projectUuid).FirstName(firstName).LastName(lastName).IsDeleted(isDeleted).WithCompanies(withCompanies).WithPlaces(withPlaces).WithSchedule(withSchedule).WithCatalog(withCatalog).WithReviews(withReviews).WithRoles(withRoles).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaffApi.SearchStaff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchStaff`: Page
    fmt.Fprintf(os.Stdout, "Response from `StaffApi.SearchStaff`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStaffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectUuid** | **string** |  | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **isDeleted** | **bool** |  | [default to false]
 **withCompanies** | **bool** |  | [default to false]
 **withPlaces** | **bool** |  | [default to false]
 **withSchedule** | **bool** |  | [default to false]
 **withCatalog** | **bool** |  | [default to false]
 **withReviews** | **bool** |  | [default to false]
 **withRoles** | **bool** |  | [default to false]
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


## UpdateStaff

> StaffResponse UpdateStaff(ctx, staffUuid).StaffCreate(staffCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Updatestaff



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
    staffUuid := "staffUuid_example" // string | 
    staffCreate := *openapiclient.NewStaffCreate("entity_uuid_1234") // StaffCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaffApi.UpdateStaff(context.Background(), staffUuid).StaffCreate(staffCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaffApi.UpdateStaff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStaff`: StaffResponse
    fmt.Fprintf(os.Stdout, "Response from `StaffApi.UpdateStaff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**staffUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStaffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **staffCreate** | [**StaffCreate**](StaffCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**StaffResponse**](StaffResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

