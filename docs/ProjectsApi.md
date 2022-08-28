# \ProjectsApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemberToProject**](ProjectsApi.md#AddMemberToProject) | **Post** /sam/projects/projects/{project_uuid}/members/{entity_uuid} | Addmembertoproject
[**ArchiveProject**](ProjectsApi.md#ArchiveProject) | **Delete** /sam/projects/projects/{project_uuid} | Archiveproject
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /sam/projects/projects | Createproject
[**CreateProjectCredential**](ProjectsApi.md#CreateProjectCredential) | **Post** /sam/projects/projects/{project_uuid}/credentials | Createprojectcredential
[**CreateProjectCredit**](ProjectsApi.md#CreateProjectCredit) | **Post** /sam/projects/projects/{project_uuid}/credits | Createprojectcredit
[**CreateProjectInvoice**](ProjectsApi.md#CreateProjectInvoice) | **Post** /sam/projects/projects/{project_uuid}/invoices | Createprojectinvoice
[**CreateProjectKey**](ProjectsApi.md#CreateProjectKey) | **Post** /sam/projects/projects/{project_uuid}/keys | Createprojectkey
[**CreateUsageType**](ProjectsApi.md#CreateUsageType) | **Post** /sam/projects/usage/types | Createusagetype
[**DeleteProjectCredential**](ProjectsApi.md#DeleteProjectCredential) | **Delete** /sam/projects/projects/{project_uuid}/credentials/{service_name} | Deleteprojectcredential
[**DeleteProjectKey**](ProjectsApi.md#DeleteProjectKey) | **Delete** /sam/projects/projects/{project_uuid}/keys | Deleteprojectkey
[**DeleteUsageType**](ProjectsApi.md#DeleteUsageType) | **Delete** /sam/projects/usage/types/{usage_type_key} | Deleteusagetype
[**GetAllProjectCredentials**](ProjectsApi.md#GetAllProjectCredentials) | **Get** /sam/projects/projects/{project_uuid}/credentials | Getallprojectcredentials
[**GetAllProjectCredits**](ProjectsApi.md#GetAllProjectCredits) | **Get** /sam/projects/projects/{project_uuid}/credits | Getallprojectcredits
[**GetAllProjectUsage**](ProjectsApi.md#GetAllProjectUsage) | **Get** /sam/projects/projects/{project_uuid}/usage | Getallprojectusage
[**GetMemberProjects**](ProjectsApi.md#GetMemberProjects) | **Get** /sam/projects/members/{entity_uuid}/projects | Getmemberprojects
[**GetProject**](ProjectsApi.md#GetProject) | **Get** /sam/projects/projects/{project_uuid} | Getproject
[**GetProjectCreditTransactions**](ProjectsApi.md#GetProjectCreditTransactions) | **Get** /sam/projects/projects/{project_uuid}/credits/{credit_uuid}/transactions | Getprojectcredittransactions
[**GetProjectInvoice**](ProjectsApi.md#GetProjectInvoice) | **Get** /sam/projects/projects/{project_uuid}/invoices | Getprojectinvoice
[**GetProjectInvoiceHistory**](ProjectsApi.md#GetProjectInvoiceHistory) | **Get** /sam/projects/projects/{project_uuid}/invoices/history | Getprojectinvoicehistory
[**GetProjectKeys**](ProjectsApi.md#GetProjectKeys) | **Get** /sam/projects/projects/{project_uuid}/keys | Getprojectkeys
[**GetProjectMembers**](ProjectsApi.md#GetProjectMembers) | **Get** /sam/projects/projects/{project_uuid}/members | Getprojectmembers
[**GetSpecificProjectCredential**](ProjectsApi.md#GetSpecificProjectCredential) | **Get** /sam/projects/projects/{project_uuid}/credentials/{service_name} | Getspecificprojectcredential
[**GetSpecificProjectUsage**](ProjectsApi.md#GetSpecificProjectUsage) | **Get** /sam/projects/projects/{project_uuid}/usage/{usage_type_key} | Getspecificprojectusage
[**GetUsageType**](ProjectsApi.md#GetUsageType) | **Get** /sam/projects/usage/types/{usage_type_key} | Getusagetype
[**RemoveMemberFromProject**](ProjectsApi.md#RemoveMemberFromProject) | **Delete** /sam/projects/projects/{project_uuid}/members/{entity_uuid} | Removememberfromproject
[**RevokeProjectCredit**](ProjectsApi.md#RevokeProjectCredit) | **Delete** /sam/projects/projects/{project_uuid}/credits/{credit_uuid} | Revokeprojectcredit
[**SearchProjects**](ProjectsApi.md#SearchProjects) | **Get** /sam/projects/projects | Searchprojects
[**SearchUsageType**](ProjectsApi.md#SearchUsageType) | **Get** /sam/projects/usage/types | Searchusagetype
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /sam/projects/projects/{project_uuid} | Updateproject
[**UpdateProjectCredential**](ProjectsApi.md#UpdateProjectCredential) | **Put** /sam/projects/projects/{project_uuid}/credentials/{service_name} | Updateprojectcredential
[**UpdateUsageType**](ProjectsApi.md#UpdateUsageType) | **Put** /sam/projects/usage/types/{usage_type_key} | Updateusagetype



## AddMemberToProject

> ResponseAddmembertoproject AddMemberToProject(ctx, projectUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Addmembertoproject

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
    projectUuid := "projectUuid_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.AddMemberToProject(context.Background(), projectUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddMemberToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMemberToProject`: ResponseAddmembertoproject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.AddMemberToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseAddmembertoproject**](ResponseAddmembertoproject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveProject

> ResponseArchiveproject ArchiveProject(ctx, projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Archiveproject

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
    projectUuid := "projectUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ArchiveProject(context.Background(), projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ArchiveProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveProject`: ResponseArchiveproject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ArchiveProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseArchiveproject**](ResponseArchiveproject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> ProjectDB CreateProject(ctx).ProjectsProjectCreate(projectsProjectCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createproject



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
    projectsProjectCreate := *openapiclient.NewProjectsProjectCreate("Name_example") // ProjectsProjectCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProject(context.Background()).ProjectsProjectCreate(projectsProjectCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsProjectCreate** | [**ProjectsProjectCreate**](ProjectsProjectCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ProjectDB**](ProjectDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectCredential

> ResponseCreateprojectcredential CreateProjectCredential(ctx, projectUuid).CreateProjectCredential(createProjectCredential).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createprojectcredential

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
    projectUuid := "projectUuid_example" // string | 
    createProjectCredential := *openapiclient.NewCreateProjectCredential("ServiceName_example", []openapiclient.Credential{*openapiclient.NewCredential("Name_example", "Value_example")}) // CreateProjectCredential | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProjectCredential(context.Background(), projectUuid).CreateProjectCredential(createProjectCredential).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProjectCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectCredential`: ResponseCreateprojectcredential
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProjectCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProjectCredential** | [**CreateProjectCredential**](CreateProjectCredential.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseCreateprojectcredential**](ResponseCreateprojectcredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectCredit

> ProjectCreditResponse CreateProjectCredit(ctx, projectUuid).CreateProjectCredit(createProjectCredit).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createprojectcredit

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
    projectUuid := "projectUuid_example" // string | 
    createProjectCredit := *openapiclient.NewCreateProjectCredit(int32(123), "GrantedReason_example") // CreateProjectCredit | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProjectCredit(context.Background(), projectUuid).CreateProjectCredit(createProjectCredit).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProjectCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectCredit`: ProjectCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProjectCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProjectCredit** | [**CreateProjectCredit**](CreateProjectCredit.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ProjectCreditResponse**](ProjectCreditResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectInvoice

> ResponseCreateprojectinvoice CreateProjectInvoice(ctx, projectUuid).CreateProjectInvoice(createProjectInvoice).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createprojectinvoice

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
    projectUuid := "projectUuid_example" // string | 
    createProjectInvoice := *openapiclient.NewCreateProjectInvoice(int32(123), int32(123)) // CreateProjectInvoice | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProjectInvoice(context.Background(), projectUuid).CreateProjectInvoice(createProjectInvoice).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProjectInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectInvoice`: ResponseCreateprojectinvoice
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProjectInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProjectInvoice** | [**CreateProjectInvoice**](CreateProjectInvoice.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseCreateprojectinvoice**](ResponseCreateprojectinvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectKey

> CreateKeyResponse CreateProjectKey(ctx, projectUuid).SecurityKeyCreate(securityKeyCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createprojectkey

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
    projectUuid := "projectUuid_example" // string | 
    securityKeyCreate := *openapiclient.NewSecurityKeyCreate("Name_example", "Summary_example") // SecurityKeyCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProjectKey(context.Background(), projectUuid).SecurityKeyCreate(securityKeyCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectKey`: CreateKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityKeyCreate** | [**SecurityKeyCreate**](SecurityKeyCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**CreateKeyResponse**](CreateKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsageType

> ProjectsUsageTypeDB CreateUsageType(ctx).ProjectsUsageTypeCreate(projectsUsageTypeCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createusagetype

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
    projectsUsageTypeCreate := *openapiclient.NewProjectsUsageTypeCreate("Key_example", "Name_example", "Summary_example", "Category_example", "Service_example", []openapiclient.ProjectsUsageTypeUnitPrice{*openapiclient.NewProjectsUsageTypeUnitPrice(int32(123), int32(123), int32(123))}) // ProjectsUsageTypeCreate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateUsageType(context.Background()).ProjectsUsageTypeCreate(projectsUsageTypeCreate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateUsageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsageType`: ProjectsUsageTypeDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateUsageType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsUsageTypeCreate** | [**ProjectsUsageTypeCreate**](ProjectsUsageTypeCreate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ProjectsUsageTypeDB**](ProjectsUsageTypeDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectCredential

> ResponseDeleteprojectcredential DeleteProjectCredential(ctx, projectUuid, serviceName).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Deleteprojectcredential

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
    projectUuid := "projectUuid_example" // string | 
    serviceName := "serviceName_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.DeleteProjectCredential(context.Background(), projectUuid, serviceName).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProjectCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProjectCredential`: ResponseDeleteprojectcredential
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.DeleteProjectCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 
**serviceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseDeleteprojectcredential**](ResponseDeleteprojectcredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectKey

> string DeleteProjectKey(ctx, projectUuid).AccessToken(accessToken).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Deleteprojectkey

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
    projectUuid := "projectUuid_example" // string | 
    accessToken := "accessToken_example" // string |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.DeleteProjectKey(context.Background(), projectUuid).AccessToken(accessToken).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProjectKey`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.DeleteProjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessToken** | **string** |  | 
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


## DeleteUsageType

> ResponseDeleteusagetype DeleteUsageType(ctx, usageTypeKey).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Deleteusagetype

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
    usageTypeKey := "usageTypeKey_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.DeleteUsageType(context.Background(), usageTypeKey).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteUsageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUsageType`: ResponseDeleteusagetype
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.DeleteUsageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageTypeKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseDeleteusagetype**](ResponseDeleteusagetype.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectCredentials

> []GetProjectCredential GetAllProjectCredentials(ctx, projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getallprojectcredentials

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
    projectUuid := "projectUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetAllProjectCredentials(context.Background(), projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAllProjectCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProjectCredentials`: []GetProjectCredential
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAllProjectCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**[]GetProjectCredential**](GetProjectCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectCredits

> Page GetAllProjectCredits(ctx, projectUuid).FullyConsumed(fullyConsumed).Revoked(revoked).Page(page).PageSize(pageSize).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getallprojectcredits

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
    projectUuid := "projectUuid_example" // string | 
    fullyConsumed := true // bool |  (optional) (default to false)
    revoked := true // bool |  (optional) (default to false)
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetAllProjectCredits(context.Background(), projectUuid).FullyConsumed(fullyConsumed).Revoked(revoked).Page(page).PageSize(pageSize).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAllProjectCredits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProjectCredits`: Page
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAllProjectCredits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullyConsumed** | **bool** |  | [default to false]
 **revoked** | **bool** |  | [default to false]
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
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


## GetAllProjectUsage

> []ProjectsProjectUsageDB GetAllProjectUsage(ctx, projectUuid).Year(year).Month(month).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getallprojectusage

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
    projectUuid := "projectUuid_example" // string | 
    year := int32(56) // int32 |  (optional)
    month := int32(56) // int32 |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetAllProjectUsage(context.Background(), projectUuid).Year(year).Month(month).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAllProjectUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProjectUsage`: []ProjectsProjectUsageDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAllProjectUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **year** | **int32** |  | 
 **month** | **int32** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**[]ProjectsProjectUsageDB**](ProjectsProjectUsageDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemberProjects

> []ProjectsProjectGet GetMemberProjects(ctx, entityUuid).Role(role).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getmemberprojects

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
    role := "role_example" // string |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetMemberProjects(context.Background(), entityUuid).Role(role).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetMemberProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemberProjects`: []ProjectsProjectGet
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetMemberProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**[]ProjectsProjectGet**](ProjectsProjectGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> ProjectDB GetProject(ctx, projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getproject



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
    projectUuid := "projectUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProject(context.Background(), projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: ProjectDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ProjectDB**](ProjectDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectCreditTransactions

> Page GetProjectCreditTransactions(ctx, projectUuid, creditUuid).Page(page).PageSize(pageSize).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getprojectcredittransactions

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
    projectUuid := "projectUuid_example" // string | 
    creditUuid := "creditUuid_example" // string | 
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProjectCreditTransactions(context.Background(), projectUuid, creditUuid).Page(page).PageSize(pageSize).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectCreditTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectCreditTransactions`: Page
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectCreditTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 
**creditUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectCreditTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
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


## GetProjectInvoice

> GetProjectInvoiceResponse GetProjectInvoice(ctx, projectUuid).WithInvoice(withInvoice).Year(year).Month(month).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getprojectinvoice

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
    projectUuid := "projectUuid_example" // string | 
    withInvoice := true // bool |  (optional) (default to false)
    year := int32(56) // int32 |  (optional)
    month := int32(56) // int32 |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProjectInvoice(context.Background(), projectUuid).WithInvoice(withInvoice).Year(year).Month(month).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectInvoice`: GetProjectInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withInvoice** | **bool** |  | [default to false]
 **year** | **int32** |  | 
 **month** | **int32** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**GetProjectInvoiceResponse**](GetProjectInvoiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectInvoiceHistory

> GetProjectInvoiceHistory GetProjectInvoiceHistory(ctx, projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getprojectinvoicehistory

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
    projectUuid := "projectUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProjectInvoiceHistory(context.Background(), projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectInvoiceHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectInvoiceHistory`: GetProjectInvoiceHistory
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectInvoiceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectInvoiceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**GetProjectInvoiceHistory**](GetProjectInvoiceHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectKeys

> []ProjectsProjectMemberDB GetProjectKeys(ctx, projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getprojectkeys

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
    projectUuid := "projectUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProjectKeys(context.Background(), projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectKeys`: []ProjectsProjectMemberDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**[]ProjectsProjectMemberDB**](ProjectsProjectMemberDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectMembers

> []ProjectsProjectMemberDB GetProjectMembers(ctx, projectUuid).Role(role).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getprojectmembers

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
    projectUuid := "projectUuid_example" // string | 
    role := "role_example" // string |  (optional)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProjectMembers(context.Background(), projectUuid).Role(role).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectMembers`: []ProjectsProjectMemberDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**[]ProjectsProjectMemberDB**](ProjectsProjectMemberDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificProjectCredential

> GetProjectCredential GetSpecificProjectCredential(ctx, projectUuid, serviceName).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getspecificprojectcredential

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
    projectUuid := "projectUuid_example" // string | 
    serviceName := "serviceName_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetSpecificProjectCredential(context.Background(), projectUuid, serviceName).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetSpecificProjectCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificProjectCredential`: GetProjectCredential
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetSpecificProjectCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 
**serviceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificProjectCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**GetProjectCredential**](GetProjectCredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificProjectUsage

> ProjectsProjectUsageDB GetSpecificProjectUsage(ctx, usageTypeKey, projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getspecificprojectusage

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
    usageTypeKey := "usageTypeKey_example" // string | 
    projectUuid := "projectUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetSpecificProjectUsage(context.Background(), usageTypeKey, projectUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetSpecificProjectUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpecificProjectUsage`: ProjectsProjectUsageDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetSpecificProjectUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageTypeKey** | **string** |  | 
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificProjectUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ProjectsProjectUsageDB**](ProjectsProjectUsageDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageType

> ProjectsUsageTypeGet GetUsageType(ctx, usageTypeKey).Execute()

Getusagetype



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
    usageTypeKey := "usageTypeKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetUsageType(context.Background(), usageTypeKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetUsageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsageType`: ProjectsUsageTypeGet
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetUsageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageTypeKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectsUsageTypeGet**](ProjectsUsageTypeGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMemberFromProject

> ResponseRemovememberfromproject RemoveMemberFromProject(ctx, projectUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Removememberfromproject

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
    projectUuid := "projectUuid_example" // string | 
    entityUuid := "entityUuid_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.RemoveMemberFromProject(context.Background(), projectUuid, entityUuid).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RemoveMemberFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMemberFromProject`: ResponseRemovememberfromproject
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.RemoveMemberFromProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 
**entityUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMemberFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseRemovememberfromproject**](ResponseRemovememberfromproject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeProjectCredit

> ResponseRevokeprojectcredit RevokeProjectCredit(ctx, projectUuid, creditUuid).RevokedReason(revokedReason).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Revokeprojectcredit

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
    projectUuid := "projectUuid_example" // string | 
    creditUuid := "creditUuid_example" // string | 
    revokedReason := "revokedReason_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.RevokeProjectCredit(context.Background(), projectUuid, creditUuid).RevokedReason(revokedReason).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RevokeProjectCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeProjectCredit`: ResponseRevokeprojectcredit
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.RevokeProjectCredit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 
**creditUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeProjectCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revokedReason** | **string** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseRevokeprojectcredit**](ResponseRevokeprojectcredit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProjects

> Page SearchProjects(ctx).IsActive(isActive).Page(page).PageSize(pageSize).Search(search).SearchOn(searchOn).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Searchprojects



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
    isActive := true // bool |  (optional) (default to false)
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    search := "search_example" // string |  (optional)
    searchOn := "searchOn_example" // string |  (optional)
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
    resp, r, err := apiClient.ProjectsApi.SearchProjects(context.Background()).IsActive(isActive).Page(page).PageSize(pageSize).Search(search).SearchOn(searchOn).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SearchProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProjects`: Page
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SearchProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isActive** | **bool** |  | [default to false]
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **search** | **string** |  | 
 **searchOn** | **string** |  | 
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


## SearchUsageType

> Page SearchUsageType(ctx).Page(page).PageSize(pageSize).Search(search).SearchOn(searchOn).SortOn(sortOn).SortDesc(sortDesc).Execute()

Searchusagetype



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
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    search := "search_example" // string |  (optional)
    searchOn := "searchOn_example" // string |  (optional)
    sortOn := "sortOn_example" // string |  (optional)
    sortDesc := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.SearchUsageType(context.Background()).Page(page).PageSize(pageSize).Search(search).SearchOn(searchOn).SortOn(sortOn).SortDesc(sortDesc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SearchUsageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsageType`: Page
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SearchUsageType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsageTypeRequest struct via the builder pattern


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


## UpdateProject

> map[string]interface{} UpdateProject(ctx, projectUuid).ProjectsProjectUpdate(projectsProjectUpdate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Updateproject

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
    projectUuid := "projectUuid_example" // string | 
    projectsProjectUpdate := *openapiclient.NewProjectsProjectUpdate() // ProjectsProjectUpdate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.UpdateProject(context.Background(), projectUuid).ProjectsProjectUpdate(projectsProjectUpdate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsProjectUpdate** | [**ProjectsProjectUpdate**](ProjectsProjectUpdate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectCredential

> ResponseUpdateprojectcredential UpdateProjectCredential(ctx, projectUuid, serviceName).UpdateProjectCredentialRequest(updateProjectCredentialRequest).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Updateprojectcredential

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
    projectUuid := "projectUuid_example" // string | 
    serviceName := "serviceName_example" // string | 
    updateProjectCredentialRequest := *openapiclient.NewUpdateProjectCredentialRequest([]openapiclient.Credential{*openapiclient.NewCredential("Name_example", "Value_example")}) // UpdateProjectCredentialRequest | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.UpdateProjectCredential(context.Background(), projectUuid, serviceName).UpdateProjectCredentialRequest(updateProjectCredentialRequest).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateProjectCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProjectCredential`: ResponseUpdateprojectcredential
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateProjectCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectUuid** | **string** |  | 
**serviceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateProjectCredentialRequest** | [**UpdateProjectCredentialRequest**](UpdateProjectCredentialRequest.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ResponseUpdateprojectcredential**](ResponseUpdateprojectcredential.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsageType

> ProjectsUsageTypeDB UpdateUsageType(ctx, usageTypeKey).ProjectsUsageTypeUpdate(projectsUsageTypeUpdate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Updateusagetype

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
    usageTypeKey := "usageTypeKey_example" // string | 
    projectsUsageTypeUpdate := *openapiclient.NewProjectsUsageTypeUpdate() // ProjectsUsageTypeUpdate | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.UpdateUsageType(context.Background(), usageTypeKey).ProjectsUsageTypeUpdate(projectsUsageTypeUpdate).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateUsageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUsageType`: ProjectsUsageTypeDB
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UpdateUsageType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usageTypeKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsageTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsUsageTypeUpdate** | [**ProjectsUsageTypeUpdate**](ProjectsUsageTypeUpdate.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**ProjectsUsageTypeDB**](ProjectsUsageTypeDB.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

