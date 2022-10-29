/*
eHelply SDK - 1.1.115

eHelply SDK for SuperStack Services

API version: 1.1.115
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


// FieldsApiService FieldsApi service
type FieldsApiService service

type ApiCreateFieldRequest struct {
	ctx context.Context
	ApiService *FieldsApiService
	field *Field
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateFieldRequest) Field(field Field) ApiCreateFieldRequest {
	r.field = &field
	return r
}

func (r ApiCreateFieldRequest) XAccessToken(xAccessToken string) ApiCreateFieldRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateFieldRequest) XSecretToken(xSecretToken string) ApiCreateFieldRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateFieldRequest) Authorization(authorization string) ApiCreateFieldRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateFieldRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateFieldRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateFieldRequest) EhelplyProject(ehelplyProject string) ApiCreateFieldRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateFieldRequest) EhelplyData(ehelplyData string) ApiCreateFieldRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateFieldRequest) Execute() (*CreateField200Response, *http.Response, error) {
	return r.ApiService.CreateFieldExecute(r)
}

/*
CreateField Createfield

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFieldRequest
*/
func (a *FieldsApiService) CreateField(ctx context.Context) ApiCreateFieldRequest {
	return ApiCreateFieldRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateField200Response
func (a *FieldsApiService) CreateFieldExecute(r ApiCreateFieldRequest) (*CreateField200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateField200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FieldsApiService.CreateField")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fields/fields"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.field == nil {
		return localVarReturnValue, nil, reportError("field is required and must be specified")
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
	localVarPostBody = r.field
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

type ApiDeleteFieldRequest struct {
	ctx context.Context
	ApiService *FieldsApiService
	fieldUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteFieldRequest) XAccessToken(xAccessToken string) ApiDeleteFieldRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteFieldRequest) XSecretToken(xSecretToken string) ApiDeleteFieldRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteFieldRequest) Authorization(authorization string) ApiDeleteFieldRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteFieldRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteFieldRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteFieldRequest) EhelplyProject(ehelplyProject string) ApiDeleteFieldRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteFieldRequest) EhelplyData(ehelplyData string) ApiDeleteFieldRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteFieldRequest) Execute() (*DeleteField200Response, *http.Response, error) {
	return r.ApiService.DeleteFieldExecute(r)
}

/*
DeleteField Deletefield

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fieldUuid
 @return ApiDeleteFieldRequest
*/
func (a *FieldsApiService) DeleteField(ctx context.Context, fieldUuid string) ApiDeleteFieldRequest {
	return ApiDeleteFieldRequest{
		ApiService: a,
		ctx: ctx,
		fieldUuid: fieldUuid,
	}
}

// Execute executes the request
//  @return DeleteField200Response
func (a *FieldsApiService) DeleteFieldExecute(r ApiDeleteFieldRequest) (*DeleteField200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteField200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FieldsApiService.DeleteField")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fields/fields/{field_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"field_uuid"+"}", url.PathEscape(parameterToString(r.fieldUuid, "")), -1)

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

type ApiGetFieldRequest struct {
	ctx context.Context
	ApiService *FieldsApiService
	fieldUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetFieldRequest) XAccessToken(xAccessToken string) ApiGetFieldRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetFieldRequest) XSecretToken(xSecretToken string) ApiGetFieldRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetFieldRequest) Authorization(authorization string) ApiGetFieldRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetFieldRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetFieldRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetFieldRequest) EhelplyProject(ehelplyProject string) ApiGetFieldRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetFieldRequest) EhelplyData(ehelplyData string) ApiGetFieldRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetFieldRequest) Execute() (*FieldDynamo, *http.Response, error) {
	return r.ApiService.GetFieldExecute(r)
}

/*
GetField Getfield

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fieldUuid
 @return ApiGetFieldRequest
*/
func (a *FieldsApiService) GetField(ctx context.Context, fieldUuid string) ApiGetFieldRequest {
	return ApiGetFieldRequest{
		ApiService: a,
		ctx: ctx,
		fieldUuid: fieldUuid,
	}
}

// Execute executes the request
//  @return FieldDynamo
func (a *FieldsApiService) GetFieldExecute(r ApiGetFieldRequest) (*FieldDynamo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FieldDynamo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FieldsApiService.GetField")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fields/fields/{field_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"field_uuid"+"}", url.PathEscape(parameterToString(r.fieldUuid, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiUpdateFieldRequest struct {
	ctx context.Context
	ApiService *FieldsApiService
	fieldUuid string
	field *Field
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateFieldRequest) Field(field Field) ApiUpdateFieldRequest {
	r.field = &field
	return r
}

func (r ApiUpdateFieldRequest) XAccessToken(xAccessToken string) ApiUpdateFieldRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateFieldRequest) XSecretToken(xSecretToken string) ApiUpdateFieldRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateFieldRequest) Authorization(authorization string) ApiUpdateFieldRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateFieldRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateFieldRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateFieldRequest) EhelplyProject(ehelplyProject string) ApiUpdateFieldRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateFieldRequest) EhelplyData(ehelplyData string) ApiUpdateFieldRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateFieldRequest) Execute() (*UpdateField200Response, *http.Response, error) {
	return r.ApiService.UpdateFieldExecute(r)
}

/*
UpdateField Updatefield

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fieldUuid
 @return ApiUpdateFieldRequest
*/
func (a *FieldsApiService) UpdateField(ctx context.Context, fieldUuid string) ApiUpdateFieldRequest {
	return ApiUpdateFieldRequest{
		ApiService: a,
		ctx: ctx,
		fieldUuid: fieldUuid,
	}
}

// Execute executes the request
//  @return UpdateField200Response
func (a *FieldsApiService) UpdateFieldExecute(r ApiUpdateFieldRequest) (*UpdateField200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateField200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FieldsApiService.UpdateField")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fields/fields/{field_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"field_uuid"+"}", url.PathEscape(parameterToString(r.fieldUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.field == nil {
		return localVarReturnValue, nil, reportError("field is required and must be specified")
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
	localVarPostBody = r.field
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
		if localVarHTTPResponse.StatusCode == 404 {
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
