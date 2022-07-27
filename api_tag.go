/*
eHelply SDK - 1.1.98

eHelply SDK for SuperStack Services

API version: 1.1.98
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


// TagApiService TagApi service
type TagApiService service

type ApiDeleteTagRequest struct {
	ctx context.Context
	ApiService *TagApiService
	tagUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteTagRequest) XAccessToken(xAccessToken string) ApiDeleteTagRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteTagRequest) XSecretToken(xSecretToken string) ApiDeleteTagRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteTagRequest) Authorization(authorization string) ApiDeleteTagRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteTagRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteTagRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteTagRequest) EhelplyProject(ehelplyProject string) ApiDeleteTagRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteTagRequest) EhelplyData(ehelplyData string) ApiDeleteTagRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteTagRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteTagExecute(r)
}

/*
DeleteTag Deletetag

Deletes the tag member with the given ID and returns True if successful

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagUuid
 @return ApiDeleteTagRequest
*/
func (a *TagApiService) DeleteTag(ctx context.Context, tagUuid string) ApiDeleteTagRequest {
	return ApiDeleteTagRequest{
		ApiService: a,
		ctx: ctx,
		tagUuid: tagUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *TagApiService) DeleteTagExecute(r ApiDeleteTagRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.DeleteTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/tags/{tag_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tag_uuid"+"}", url.PathEscape(parameterToString(r.tagUuid, "")), -1)

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
