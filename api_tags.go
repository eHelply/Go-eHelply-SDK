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


// TagsApiService TagsApi service
type TagsApiService service

type ApiCreateTagRequest struct {
	ctx context.Context
	ApiService *TagsApiService
	tagBase *TagBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateTagRequest) TagBase(tagBase TagBase) ApiCreateTagRequest {
	r.tagBase = &tagBase
	return r
}

func (r ApiCreateTagRequest) XAccessToken(xAccessToken string) ApiCreateTagRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateTagRequest) XSecretToken(xSecretToken string) ApiCreateTagRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateTagRequest) Authorization(authorization string) ApiCreateTagRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateTagRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateTagRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateTagRequest) EhelplyProject(ehelplyProject string) ApiCreateTagRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateTagRequest) EhelplyData(ehelplyData string) ApiCreateTagRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateTagRequest) Execute() (*TagDb, *http.Response, error) {
	return r.ApiService.CreateTagExecute(r)
}

/*
CreateTag Createtag

Creates a tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTagRequest
*/
func (a *TagsApiService) CreateTag(ctx context.Context) ApiCreateTagRequest {
	return ApiCreateTagRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TagDb
func (a *TagsApiService) CreateTagExecute(r ApiCreateTagRequest) (*TagDb, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagDb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.CreateTag")
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

type ApiGetTagRequest struct {
	ctx context.Context
	ApiService *TagsApiService
	tagUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetTagRequest) XAccessToken(xAccessToken string) ApiGetTagRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetTagRequest) XSecretToken(xSecretToken string) ApiGetTagRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetTagRequest) Authorization(authorization string) ApiGetTagRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetTagRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetTagRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetTagRequest) EhelplyProject(ehelplyProject string) ApiGetTagRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetTagRequest) EhelplyData(ehelplyData string) ApiGetTagRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetTagRequest) Execute() (*TagBase, *http.Response, error) {
	return r.ApiService.GetTagExecute(r)
}

/*
GetTag Gettag

Gets the tag member information given the tag ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagUuid
 @return ApiGetTagRequest
*/
func (a *TagsApiService) GetTag(ctx context.Context, tagUuid string) ApiGetTagRequest {
	return ApiGetTagRequest{
		ApiService: a,
		ctx: ctx,
		tagUuid: tagUuid,
	}
}

// Execute executes the request
//  @return TagBase
func (a *TagsApiService) GetTagExecute(r ApiGetTagRequest) (*TagBase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.GetTag")
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

type ApiSearchTagRequest struct {
	ctx context.Context
	ApiService *TagsApiService
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

func (r ApiSearchTagRequest) ProjectUuid(projectUuid string) ApiSearchTagRequest {
	r.projectUuid = &projectUuid
	return r
}

func (r ApiSearchTagRequest) Name(name string) ApiSearchTagRequest {
	r.name = &name
	return r
}

func (r ApiSearchTagRequest) Page(page int32) ApiSearchTagRequest {
	r.page = &page
	return r
}

func (r ApiSearchTagRequest) PageSize(pageSize int32) ApiSearchTagRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchTagRequest) SortOn(sortOn string) ApiSearchTagRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchTagRequest) SortDesc(sortDesc bool) ApiSearchTagRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchTagRequest) XAccessToken(xAccessToken string) ApiSearchTagRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchTagRequest) XSecretToken(xSecretToken string) ApiSearchTagRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchTagRequest) Authorization(authorization string) ApiSearchTagRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchTagRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchTagRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchTagRequest) EhelplyProject(ehelplyProject string) ApiSearchTagRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchTagRequest) EhelplyData(ehelplyData string) ApiSearchTagRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchTagRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchTagExecute(r)
}

/*
SearchTag Searchtag

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
 @return ApiSearchTagRequest
*/
func (a *TagsApiService) SearchTag(ctx context.Context) ApiSearchTagRequest {
	return ApiSearchTagRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Page
func (a *TagsApiService) SearchTagExecute(r ApiSearchTagRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.SearchTag")
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

type ApiUpdateTagRequest struct {
	ctx context.Context
	ApiService *TagsApiService
	tagUuid string
	tagBase *TagBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateTagRequest) TagBase(tagBase TagBase) ApiUpdateTagRequest {
	r.tagBase = &tagBase
	return r
}

func (r ApiUpdateTagRequest) XAccessToken(xAccessToken string) ApiUpdateTagRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateTagRequest) XSecretToken(xSecretToken string) ApiUpdateTagRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateTagRequest) Authorization(authorization string) ApiUpdateTagRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateTagRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateTagRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateTagRequest) EhelplyProject(ehelplyProject string) ApiUpdateTagRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateTagRequest) EhelplyData(ehelplyData string) ApiUpdateTagRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateTagRequest) Execute() (*TagBase, *http.Response, error) {
	return r.ApiService.UpdateTagExecute(r)
}

/*
UpdateTag Updatetag

Update tag with given info, only updating the fields supplied. Tag Uuid must be sent however.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagUuid
 @return ApiUpdateTagRequest
*/
func (a *TagsApiService) UpdateTag(ctx context.Context, tagUuid string) ApiUpdateTagRequest {
	return ApiUpdateTagRequest{
		ApiService: a,
		ctx: ctx,
		tagUuid: tagUuid,
	}
}

// Execute executes the request
//  @return TagBase
func (a *TagsApiService) UpdateTagExecute(r ApiUpdateTagRequest) (*TagBase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.UpdateTag")
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
