/*
eHelply SDK - 1.1.122

eHelply SDK for SuperStack Services

API version: 1.1.122
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
	"os"
)


// ContentApiService ContentApi service
type ContentApiService service

type ApiCreateFileRequest struct {
	ctx context.Context
	ApiService *ContentApiService
	file **os.File
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateFileRequest) File(file *os.File) ApiCreateFileRequest {
	r.file = &file
	return r
}

func (r ApiCreateFileRequest) XAccessToken(xAccessToken string) ApiCreateFileRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateFileRequest) XSecretToken(xSecretToken string) ApiCreateFileRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateFileRequest) Authorization(authorization string) ApiCreateFileRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateFileRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateFileRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateFileRequest) EhelplyProject(ehelplyProject string) ApiCreateFileRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateFileRequest) EhelplyData(ehelplyData string) ApiCreateFileRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateFileRequest) Execute() (*CreateFile200Response, *http.Response, error) {
	return r.ApiService.CreateFileExecute(r)
}

/*
CreateFile Createfile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFileRequest
*/
func (a *ContentApiService) CreateFile(ctx context.Context) ApiCreateFileRequest {
	return ApiCreateFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateFile200Response
func (a *ContentApiService) CreateFileExecute(r ApiCreateFileRequest) (*CreateFile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateFile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.CreateFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/files"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	fileLocalVarFile := *r.file
	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiDeleteFileRequest struct {
	ctx context.Context
	ApiService *ContentApiService
	fileUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteFileRequest) XAccessToken(xAccessToken string) ApiDeleteFileRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteFileRequest) XSecretToken(xSecretToken string) ApiDeleteFileRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteFileRequest) Authorization(authorization string) ApiDeleteFileRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteFileRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteFileRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteFileRequest) EhelplyProject(ehelplyProject string) ApiDeleteFileRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteFileRequest) EhelplyData(ehelplyData string) ApiDeleteFileRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteFileRequest) Execute() (*DeleteFile200Response, *http.Response, error) {
	return r.ApiService.DeleteFileExecute(r)
}

/*
DeleteFile Deletefile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileUuid
 @return ApiDeleteFileRequest
*/
func (a *ContentApiService) DeleteFile(ctx context.Context, fileUuid string) ApiDeleteFileRequest {
	return ApiDeleteFileRequest{
		ApiService: a,
		ctx: ctx,
		fileUuid: fileUuid,
	}
}

// Execute executes the request
//  @return DeleteFile200Response
func (a *ContentApiService) DeleteFileExecute(r ApiDeleteFileRequest) (*DeleteFile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteFile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.DeleteFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/files/{file_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_uuid"+"}", url.PathEscape(parameterToString(r.fileUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiGetFileRequest struct {
	ctx context.Context
	ApiService *ContentApiService
	fileUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetFileRequest) XAccessToken(xAccessToken string) ApiGetFileRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetFileRequest) XSecretToken(xSecretToken string) ApiGetFileRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetFileRequest) Authorization(authorization string) ApiGetFileRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetFileRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetFileRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetFileRequest) EhelplyProject(ehelplyProject string) ApiGetFileRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetFileRequest) EhelplyData(ehelplyData string) ApiGetFileRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetFileRequest) Execute() (**os.File, *http.Response, error) {
	return r.ApiService.GetFileExecute(r)
}

/*
GetFile Getfile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileUuid
 @return ApiGetFileRequest
*/
func (a *ContentApiService) GetFile(ctx context.Context, fileUuid string) ApiGetFileRequest {
	return ApiGetFileRequest{
		ApiService: a,
		ctx: ctx,
		fileUuid: fileUuid,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ContentApiService) GetFileExecute(r ApiGetFileRequest) (**os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  **os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.GetFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/files/{file_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_uuid"+"}", url.PathEscape(parameterToString(r.fileUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiUpdateFileRequest struct {
	ctx context.Context
	ApiService *ContentApiService
	fileUuid string
	file **os.File
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateFileRequest) File(file *os.File) ApiUpdateFileRequest {
	r.file = &file
	return r
}

func (r ApiUpdateFileRequest) XAccessToken(xAccessToken string) ApiUpdateFileRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateFileRequest) XSecretToken(xSecretToken string) ApiUpdateFileRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateFileRequest) Authorization(authorization string) ApiUpdateFileRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateFileRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateFileRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateFileRequest) EhelplyProject(ehelplyProject string) ApiUpdateFileRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateFileRequest) EhelplyData(ehelplyData string) ApiUpdateFileRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateFileRequest) Execute() (*UpdateFile200Response, *http.Response, error) {
	return r.ApiService.UpdateFileExecute(r)
}

/*
UpdateFile Updatefile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fileUuid
 @return ApiUpdateFileRequest
*/
func (a *ContentApiService) UpdateFile(ctx context.Context, fileUuid string) ApiUpdateFileRequest {
	return ApiUpdateFileRequest{
		ApiService: a,
		ctx: ctx,
		fileUuid: fileUuid,
	}
}

// Execute executes the request
//  @return UpdateFile200Response
func (a *ContentApiService) UpdateFileExecute(r ApiUpdateFileRequest) (*UpdateFile200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateFile200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentApiService.UpdateFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/files/{file_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"file_uuid"+"}", url.PathEscape(parameterToString(r.fileUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"

	fileLocalVarFile := *r.file
	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)
		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetAppointment403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
