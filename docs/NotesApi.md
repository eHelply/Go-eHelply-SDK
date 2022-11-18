# \NotesApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNote**](NotesApi.md#CreateNote) | **Post** /notes/notes | Createnote
[**DeleteNote**](NotesApi.md#DeleteNote) | **Delete** /notes/notes/{note_id} | Deletenote
[**GetNote**](NotesApi.md#GetNote) | **Get** /notes/notes/{note_id} | Getnote
[**UpdateNote**](NotesApi.md#UpdateNote) | **Put** /notes/notes/{note_id} | Updatenote



## CreateNote

> CreateNoteResponse CreateNote(ctx).NoteBase(noteBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Createnote

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
    noteBase := *openapiclient.NewNoteBase("Content_example", "Time_example", *openapiclient.NewNoteMeta("Author_example")) // NoteBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotesApi.CreateNote(context.Background()).NoteBase(noteBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotesApi.CreateNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNote`: CreateNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `NotesApi.CreateNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noteBase** | [**NoteBase**](NoteBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**CreateNoteResponse**](CreateNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNote

> DeleteNote200Response DeleteNote(ctx, noteId).Method(method).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Deletenote

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
    noteId := "noteId_example" // string | 
    method := "method_example" // string |  (optional) (default to "previous")
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotesApi.DeleteNote(context.Background(), noteId).Method(method).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotesApi.DeleteNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNote`: DeleteNote200Response
    fmt.Fprintf(os.Stdout, "Response from `NotesApi.DeleteNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **method** | **string** |  | [default to &quot;previous&quot;]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**DeleteNote200Response**](DeleteNote200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNote

> NoteDynamoHistoryResponse GetNote(ctx, noteId).History(history).HistoryContent(historyContent).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getnote

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
    noteId := "noteId_example" // string | 
    history := int32(56) // int32 |  (optional) (default to 0)
    historyContent := true // bool |  (optional) (default to true)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotesApi.GetNote(context.Background(), noteId).History(history).HistoryContent(historyContent).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotesApi.GetNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNote`: NoteDynamoHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `NotesApi.GetNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **history** | **int32** |  | [default to 0]
 **historyContent** | **bool** |  | [default to true]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**NoteDynamoHistoryResponse**](NoteDynamoHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNote

> UpdateNote200Response UpdateNote(ctx, noteId).NoteBase(noteBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Updatenote

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
    noteId := "noteId_example" // string | 
    noteBase := *openapiclient.NewNoteBase("Content_example", "Time_example", *openapiclient.NewNoteMeta("Author_example")) // NoteBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotesApi.UpdateNote(context.Background(), noteId).NoteBase(noteBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotesApi.UpdateNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNote`: UpdateNote200Response
    fmt.Fprintf(os.Stdout, "Response from `NotesApi.UpdateNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteBase** | [**NoteBase**](NoteBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**UpdateNote200Response**](UpdateNote200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

