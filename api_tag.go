/*
eHelply SDK - 1.1.88

eHelply SDK for SuperStack Services

API version: 1.1.88
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

type ApiCreateTagPlacesTagsPostRequest struct {
	ctx context.Context
	ApiService *TagApiService
	tagBase *TagBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateTagPlacesTagsPostRequest) TagBase(tagBase TagBase) ApiCreateTagPlacesTagsPostRequest {
	r.tagBase = &tagBase
	return r
}

func (r ApiCreateTagPlacesTagsPostRequest) XAccessToken(xAccessToken string) ApiCreateTagPlacesTagsPostRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateTagPlacesTagsPostRequest) XSecretToken(xSecretToken string) ApiCreateTagPlacesTagsPostRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateTagPlacesTagsPostRequest) Authorization(authorization string) ApiCreateTagPlacesTagsPostRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateTagPlacesTagsPostRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateTagPlacesTagsPostRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateTagPlacesTagsPostRequest) EhelplyProject(ehelplyProject string) ApiCreateTagPlacesTagsPostRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateTagPlacesTagsPostRequest) EhelplyData(ehelplyData string) ApiCreateTagPlacesTagsPostRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateTagPlacesTagsPostRequest) Execute() (*TagDb, *http.Response, error) {
	return r.ApiService.CreateTagPlacesTagsPostExecute(r)
}

/*
CreateTagPlacesTagsPost Create Tag

Creates a tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTagPlacesTagsPostRequest
*/
func (a *TagApiService) CreateTagPlacesTagsPost(ctx context.Context) ApiCreateTagPlacesTagsPostRequest {
	return ApiCreateTagPlacesTagsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TagDb
func (a *TagApiService) CreateTagPlacesTagsPostExecute(r ApiCreateTagPlacesTagsPostRequest) (*TagDb, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagDb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.CreateTagPlacesTagsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tagBase == nil {
		return localVarReturnValue, nil, reportError("tagBase is required and must be specified")
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
	localVarPostBody = r.tagBase
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

type ApiDeleteTagPlacesTagsTagUuidDeleteRequest struct {
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

func (r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) XAccessToken(xAccessToken string) ApiDeleteTagPlacesTagsTagUuidDeleteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) XSecretToken(xSecretToken string) ApiDeleteTagPlacesTagsTagUuidDeleteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) Authorization(authorization string) ApiDeleteTagPlacesTagsTagUuidDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteTagPlacesTagsTagUuidDeleteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) EhelplyProject(ehelplyProject string) ApiDeleteTagPlacesTagsTagUuidDeleteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) EhelplyData(ehelplyData string) ApiDeleteTagPlacesTagsTagUuidDeleteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteTagPlacesTagsTagUuidDeleteExecute(r)
}

/*
DeleteTagPlacesTagsTagUuidDelete Delete Tag

Deletes the tag member with the given ID and returns True if successful

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagUuid
 @return ApiDeleteTagPlacesTagsTagUuidDeleteRequest
*/
func (a *TagApiService) DeleteTagPlacesTagsTagUuidDelete(ctx context.Context, tagUuid string) ApiDeleteTagPlacesTagsTagUuidDeleteRequest {
	return ApiDeleteTagPlacesTagsTagUuidDeleteRequest{
		ApiService: a,
		ctx: ctx,
		tagUuid: tagUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *TagApiService) DeleteTagPlacesTagsTagUuidDeleteExecute(r ApiDeleteTagPlacesTagsTagUuidDeleteRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.DeleteTagPlacesTagsTagUuidDelete")
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

type ApiGetTagPlacesTagsTagUuidGetRequest struct {
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

func (r ApiGetTagPlacesTagsTagUuidGetRequest) XAccessToken(xAccessToken string) ApiGetTagPlacesTagsTagUuidGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetTagPlacesTagsTagUuidGetRequest) XSecretToken(xSecretToken string) ApiGetTagPlacesTagsTagUuidGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetTagPlacesTagsTagUuidGetRequest) Authorization(authorization string) ApiGetTagPlacesTagsTagUuidGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetTagPlacesTagsTagUuidGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetTagPlacesTagsTagUuidGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetTagPlacesTagsTagUuidGetRequest) EhelplyProject(ehelplyProject string) ApiGetTagPlacesTagsTagUuidGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetTagPlacesTagsTagUuidGetRequest) EhelplyData(ehelplyData string) ApiGetTagPlacesTagsTagUuidGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetTagPlacesTagsTagUuidGetRequest) Execute() (*TagBase, *http.Response, error) {
	return r.ApiService.GetTagPlacesTagsTagUuidGetExecute(r)
}

/*
GetTagPlacesTagsTagUuidGet Get Tag

Gets the tag member information given the tag ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagUuid
 @return ApiGetTagPlacesTagsTagUuidGetRequest
*/
func (a *TagApiService) GetTagPlacesTagsTagUuidGet(ctx context.Context, tagUuid string) ApiGetTagPlacesTagsTagUuidGetRequest {
	return ApiGetTagPlacesTagsTagUuidGetRequest{
		ApiService: a,
		ctx: ctx,
		tagUuid: tagUuid,
	}
}

// Execute executes the request
//  @return TagBase
func (a *TagApiService) GetTagPlacesTagsTagUuidGetExecute(r ApiGetTagPlacesTagsTagUuidGetRequest) (*TagBase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.GetTagPlacesTagsTagUuidGet")
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

type ApiSearchTagsPlacesTagsGetRequest struct {
	ctx context.Context
	ApiService *TagApiService
	projectUuid *string
	name *string
	page *int32
	pageSize *int32
	sortOn *string
	sortDesc *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiSearchTagsPlacesTagsGetRequest) ProjectUuid(projectUuid string) ApiSearchTagsPlacesTagsGetRequest {
	r.projectUuid = &projectUuid
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) Name(name string) ApiSearchTagsPlacesTagsGetRequest {
	r.name = &name
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) Page(page int32) ApiSearchTagsPlacesTagsGetRequest {
	r.page = &page
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) PageSize(pageSize int32) ApiSearchTagsPlacesTagsGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) SortOn(sortOn string) ApiSearchTagsPlacesTagsGetRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) SortDesc(sortDesc bool) ApiSearchTagsPlacesTagsGetRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) XAccessToken(xAccessToken string) ApiSearchTagsPlacesTagsGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) XSecretToken(xSecretToken string) ApiSearchTagsPlacesTagsGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) Authorization(authorization string) ApiSearchTagsPlacesTagsGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchTagsPlacesTagsGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) EhelplyProject(ehelplyProject string) ApiSearchTagsPlacesTagsGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) EhelplyData(ehelplyData string) ApiSearchTagsPlacesTagsGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchTagsPlacesTagsGetRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchTagsPlacesTagsGetExecute(r)
}

/*
SearchTagsPlacesTagsGet Search Tags

TODO
Item return format:
```
{
    uuid                                **type:** string
    project_uuid                        **type:** string or None

    name                                **type:** string or None

    meta                                **type:** dict or None

    created_at                          **type:** string or None

    updated_at                          **type:** string or None

    deleted_at                          **type:** string or None

}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchTagsPlacesTagsGetRequest
*/
func (a *TagApiService) SearchTagsPlacesTagsGet(ctx context.Context) ApiSearchTagsPlacesTagsGetRequest {
	return ApiSearchTagsPlacesTagsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Page
func (a *TagApiService) SearchTagsPlacesTagsGetExecute(r ApiSearchTagsPlacesTagsGetRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.SearchTagsPlacesTagsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectUuid != nil {
		localVarQueryParams.Add("project_uuid", parameterToString(*r.projectUuid, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.sortOn != nil {
		localVarQueryParams.Add("sort_on", parameterToString(*r.sortOn, ""))
	}
	if r.sortDesc != nil {
		localVarQueryParams.Add("sort_desc", parameterToString(*r.sortDesc, ""))
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

type ApiUpdateTagPlacesTagsTagUuidPutRequest struct {
	ctx context.Context
	ApiService *TagApiService
	tagUuid string
	tagBase *TagBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) TagBase(tagBase TagBase) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	r.tagBase = &tagBase
	return r
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) XAccessToken(xAccessToken string) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) XSecretToken(xSecretToken string) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) Authorization(authorization string) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) EhelplyProject(ehelplyProject string) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) EhelplyData(ehelplyData string) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateTagPlacesTagsTagUuidPutRequest) Execute() (*TagBase, *http.Response, error) {
	return r.ApiService.UpdateTagPlacesTagsTagUuidPutExecute(r)
}

/*
UpdateTagPlacesTagsTagUuidPut Update Tag

Update tag with given info, only updating the fields supplied. Tag Uuid must be sent however.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagUuid
 @return ApiUpdateTagPlacesTagsTagUuidPutRequest
*/
func (a *TagApiService) UpdateTagPlacesTagsTagUuidPut(ctx context.Context, tagUuid string) ApiUpdateTagPlacesTagsTagUuidPutRequest {
	return ApiUpdateTagPlacesTagsTagUuidPutRequest{
		ApiService: a,
		ctx: ctx,
		tagUuid: tagUuid,
	}
}

// Execute executes the request
//  @return TagBase
func (a *TagApiService) UpdateTagPlacesTagsTagUuidPutExecute(r ApiUpdateTagPlacesTagsTagUuidPutRequest) (*TagBase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.UpdateTagPlacesTagsTagUuidPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/tags/{tag_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tag_uuid"+"}", url.PathEscape(parameterToString(r.tagUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tagBase == nil {
		return localVarReturnValue, nil, reportError("tagBase is required and must be specified")
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
	localVarPostBody = r.tagBase
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