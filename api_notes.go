/*
eHelply SDK - 1.1.113

eHelply SDK for SuperStack Services

API version: 1.1.113
Contact: support@ehelply.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ehelply

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// NotesApiService NotesApi service
type NotesApiService service

type ApiCreateNoteRequest struct {
	ctx context.Context
	ApiService *NotesApiService
	noteBase *NoteBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateNoteRequest) NoteBase(noteBase NoteBase) ApiCreateNoteRequest {
	r.noteBase = &noteBase
	return r
}

func (r ApiCreateNoteRequest) XAccessToken(xAccessToken string) ApiCreateNoteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateNoteRequest) XSecretToken(xSecretToken string) ApiCreateNoteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateNoteRequest) Authorization(authorization string) ApiCreateNoteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateNoteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateNoteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateNoteRequest) EhelplyProject(ehelplyProject string) ApiCreateNoteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateNoteRequest) EhelplyData(ehelplyData string) ApiCreateNoteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateNoteRequest) Execute() (*NoteDynamoResponse, *http.Response, error) {
	return r.ApiService.CreateNoteExecute(r)
}

/*
CreateNote Create Note

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNoteRequest
*/
func (a *NotesApiService) CreateNote(ctx context.Context) ApiCreateNoteRequest {
	return ApiCreateNoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NoteDynamoResponse
func (a *NotesApiService) CreateNoteExecute(r ApiCreateNoteRequest) (*NoteDynamoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NoteDynamoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.CreateNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/notes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.noteBase == nil {
		return localVarReturnValue, nil, reportError("noteBase is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAccessToken != nil {
		localVarHeaderParams["x-access-token"] = parameterToString(*r.xAccessToken, "")
	}
	if r.xSecretToken != nil {
		localVarHeaderParams["x-secret-token"] = parameterToString(*r.xSecretToken, "")
	}
	if r.authorization != nil {
		localVarHeaderParams["authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ehelplyActiveParticipant != nil {
		localVarHeaderParams["ehelply-active-participant"] = parameterToString(*r.ehelplyActiveParticipant, "")
	}
	if r.ehelplyProject != nil {
		localVarHeaderParams["ehelply-project"] = parameterToString(*r.ehelplyProject, "")
	}
	if r.ehelplyData != nil {
		localVarHeaderParams["ehelply-data"] = parameterToString(*r.ehelplyData, "")
	}
	// body params
	localVarPostBody = r.noteBase
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteNoteRequest struct {
	ctx context.Context
	ApiService *NotesApiService
	noteId string
	method *string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteNoteRequest) Method(method string) ApiDeleteNoteRequest {
	r.method = &method
	return r
}

func (r ApiDeleteNoteRequest) XAccessToken(xAccessToken string) ApiDeleteNoteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteNoteRequest) XSecretToken(xSecretToken string) ApiDeleteNoteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteNoteRequest) Authorization(authorization string) ApiDeleteNoteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteNoteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteNoteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteNoteRequest) EhelplyProject(ehelplyProject string) ApiDeleteNoteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteNoteRequest) EhelplyData(ehelplyData string) ApiDeleteNoteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteNoteRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteNoteExecute(r)
}

/*
DeleteNote Delete Note

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param noteId
 @return ApiDeleteNoteRequest
*/
func (a *NotesApiService) DeleteNote(ctx context.Context, noteId string) ApiDeleteNoteRequest {
	return ApiDeleteNoteRequest{
		ApiService: a,
		ctx: ctx,
		noteId: noteId,
	}
}

// Execute executes the request
//  @return interface{}
func (a *NotesApiService) DeleteNoteExecute(r ApiDeleteNoteRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.DeleteNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/notes/{note_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"note_id"+"}", url.PathEscape(parameterToString(r.noteId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.method != nil {
		localVarQueryParams.Add("method", parameterToString(*r.method, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAccessToken != nil {
		localVarHeaderParams["x-access-token"] = parameterToString(*r.xAccessToken, "")
	}
	if r.xSecretToken != nil {
		localVarHeaderParams["x-secret-token"] = parameterToString(*r.xSecretToken, "")
	}
	if r.authorization != nil {
		localVarHeaderParams["authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ehelplyActiveParticipant != nil {
		localVarHeaderParams["ehelply-active-participant"] = parameterToString(*r.ehelplyActiveParticipant, "")
	}
	if r.ehelplyProject != nil {
		localVarHeaderParams["ehelply-project"] = parameterToString(*r.ehelplyProject, "")
	}
	if r.ehelplyData != nil {
		localVarHeaderParams["ehelply-data"] = parameterToString(*r.ehelplyData, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetNoteRequest struct {
	ctx context.Context
	ApiService *NotesApiService
	noteId string
	history *int32
	historyContent *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetNoteRequest) History(history int32) ApiGetNoteRequest {
	r.history = &history
	return r
}

func (r ApiGetNoteRequest) HistoryContent(historyContent bool) ApiGetNoteRequest {
	r.historyContent = &historyContent
	return r
}

func (r ApiGetNoteRequest) XAccessToken(xAccessToken string) ApiGetNoteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetNoteRequest) XSecretToken(xSecretToken string) ApiGetNoteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetNoteRequest) Authorization(authorization string) ApiGetNoteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetNoteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetNoteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetNoteRequest) EhelplyProject(ehelplyProject string) ApiGetNoteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetNoteRequest) EhelplyData(ehelplyData string) ApiGetNoteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetNoteRequest) Execute() (*NoteDynamoHistoryResponse, *http.Response, error) {
	return r.ApiService.GetNoteExecute(r)
}

/*
GetNote Get Note

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param noteId
 @return ApiGetNoteRequest
*/
func (a *NotesApiService) GetNote(ctx context.Context, noteId string) ApiGetNoteRequest {
	return ApiGetNoteRequest{
		ApiService: a,
		ctx: ctx,
		noteId: noteId,
	}
}

// Execute executes the request
//  @return NoteDynamoHistoryResponse
func (a *NotesApiService) GetNoteExecute(r ApiGetNoteRequest) (*NoteDynamoHistoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NoteDynamoHistoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.GetNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/notes/{note_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"note_id"+"}", url.PathEscape(parameterToString(r.noteId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.history != nil {
		localVarQueryParams.Add("history", parameterToString(*r.history, ""))
	}
	if r.historyContent != nil {
		localVarQueryParams.Add("history_content", parameterToString(*r.historyContent, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAccessToken != nil {
		localVarHeaderParams["x-access-token"] = parameterToString(*r.xAccessToken, "")
	}
	if r.xSecretToken != nil {
		localVarHeaderParams["x-secret-token"] = parameterToString(*r.xSecretToken, "")
	}
	if r.authorization != nil {
		localVarHeaderParams["authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ehelplyActiveParticipant != nil {
		localVarHeaderParams["ehelply-active-participant"] = parameterToString(*r.ehelplyActiveParticipant, "")
	}
	if r.ehelplyProject != nil {
		localVarHeaderParams["ehelply-project"] = parameterToString(*r.ehelplyProject, "")
	}
	if r.ehelplyData != nil {
		localVarHeaderParams["ehelply-data"] = parameterToString(*r.ehelplyData, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateNoteRequest struct {
	ctx context.Context
	ApiService *NotesApiService
	noteId string
	noteBase *NoteBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateNoteRequest) NoteBase(noteBase NoteBase) ApiUpdateNoteRequest {
	r.noteBase = &noteBase
	return r
}

func (r ApiUpdateNoteRequest) XAccessToken(xAccessToken string) ApiUpdateNoteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateNoteRequest) XSecretToken(xSecretToken string) ApiUpdateNoteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateNoteRequest) Authorization(authorization string) ApiUpdateNoteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateNoteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateNoteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateNoteRequest) EhelplyProject(ehelplyProject string) ApiUpdateNoteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateNoteRequest) EhelplyData(ehelplyData string) ApiUpdateNoteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateNoteRequest) Execute() (*NoteDynamoResponse, *http.Response, error) {
	return r.ApiService.UpdateNoteExecute(r)
}

/*
UpdateNote Update Note

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param noteId
 @return ApiUpdateNoteRequest
*/
func (a *NotesApiService) UpdateNote(ctx context.Context, noteId string) ApiUpdateNoteRequest {
	return ApiUpdateNoteRequest{
		ApiService: a,
		ctx: ctx,
		noteId: noteId,
	}
}

// Execute executes the request
//  @return NoteDynamoResponse
func (a *NotesApiService) UpdateNoteExecute(r ApiUpdateNoteRequest) (*NoteDynamoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NoteDynamoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesApiService.UpdateNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/notes/{note_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"note_id"+"}", url.PathEscape(parameterToString(r.noteId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.noteBase == nil {
		return localVarReturnValue, nil, reportError("noteBase is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAccessToken != nil {
		localVarHeaderParams["x-access-token"] = parameterToString(*r.xAccessToken, "")
	}
	if r.xSecretToken != nil {
		localVarHeaderParams["x-secret-token"] = parameterToString(*r.xSecretToken, "")
	}
	if r.authorization != nil {
		localVarHeaderParams["authorization"] = parameterToString(*r.authorization, "")
	}
	if r.ehelplyActiveParticipant != nil {
		localVarHeaderParams["ehelply-active-participant"] = parameterToString(*r.ehelplyActiveParticipant, "")
	}
	if r.ehelplyProject != nil {
		localVarHeaderParams["ehelply-project"] = parameterToString(*r.ehelplyProject, "")
	}
	if r.ehelplyData != nil {
		localVarHeaderParams["ehelply-data"] = parameterToString(*r.ehelplyData, "")
	}
	// body params
	localVarPostBody = r.noteBase
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
