/*
eHelply SDK - 1.1.118

eHelply SDK for SuperStack Services

API version: 1.1.118
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


// CategoryApiService CategoryApi service
type CategoryApiService service

type ApiCreateCategoryPlacesCategoriesPostRequest struct {
	ctx context.Context
	ApiService *CategoryApiService
	categoryBase *CategoryBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) CategoryBase(categoryBase CategoryBase) ApiCreateCategoryPlacesCategoriesPostRequest {
	r.categoryBase = &categoryBase
	return r
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) XAccessToken(xAccessToken string) ApiCreateCategoryPlacesCategoriesPostRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) XSecretToken(xSecretToken string) ApiCreateCategoryPlacesCategoriesPostRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) Authorization(authorization string) ApiCreateCategoryPlacesCategoriesPostRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateCategoryPlacesCategoriesPostRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) EhelplyProject(ehelplyProject string) ApiCreateCategoryPlacesCategoriesPostRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) EhelplyData(ehelplyData string) ApiCreateCategoryPlacesCategoriesPostRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateCategoryPlacesCategoriesPostRequest) Execute() (*CategoryDb, *http.Response, error) {
	return r.ApiService.CreateCategoryPlacesCategoriesPostExecute(r)
}

/*
CreateCategoryPlacesCategoriesPost Create Category

Creates a category

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCategoryPlacesCategoriesPostRequest
*/
func (a *CategoryApiService) CreateCategoryPlacesCategoriesPost(ctx context.Context) ApiCreateCategoryPlacesCategoriesPostRequest {
	return ApiCreateCategoryPlacesCategoriesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CategoryDb
func (a *CategoryApiService) CreateCategoryPlacesCategoriesPostExecute(r ApiCreateCategoryPlacesCategoriesPostRequest) (*CategoryDb, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CategoryDb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoryApiService.CreateCategoryPlacesCategoriesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.categoryBase == nil {
		return localVarReturnValue, nil, reportError("categoryBase is required and must be specified")
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
	localVarPostBody = r.categoryBase
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

type ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest struct {
	ctx context.Context
	ApiService *CategoryApiService
	categoryUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) XAccessToken(xAccessToken string) ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) XSecretToken(xSecretToken string) ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) Authorization(authorization string) ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) EhelplyProject(ehelplyProject string) ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) EhelplyData(ehelplyData string) ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteCategoryPlacesCategoriesCategoryUuidDeleteExecute(r)
}

/*
DeleteCategoryPlacesCategoriesCategoryUuidDelete Delete Category

Deletes the category with the given ID and returns True if successful

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param categoryUuid
 @return ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest
*/
func (a *CategoryApiService) DeleteCategoryPlacesCategoriesCategoryUuidDelete(ctx context.Context, categoryUuid string) ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest {
	return ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest{
		ApiService: a,
		ctx: ctx,
		categoryUuid: categoryUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *CategoryApiService) DeleteCategoryPlacesCategoriesCategoryUuidDeleteExecute(r ApiDeleteCategoryPlacesCategoriesCategoryUuidDeleteRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoryApiService.DeleteCategoryPlacesCategoriesCategoryUuidDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/categories/{category_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"category_uuid"+"}", url.PathEscape(parameterToString(r.categoryUuid, "")), -1)

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

type ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest struct {
	ctx context.Context
	ApiService *CategoryApiService
	categoryUuid string
	withMeta *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) WithMeta(withMeta bool) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	r.withMeta = &withMeta
	return r
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) XAccessToken(xAccessToken string) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) XSecretToken(xSecretToken string) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) Authorization(authorization string) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) EhelplyProject(ehelplyProject string) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) EhelplyData(ehelplyData string) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) Execute() (*CategoryBase, *http.Response, error) {
	return r.ApiService.GetCategoryPlacesCategoriesCategoryUuidGetExecute(r)
}

/*
GetCategoryPlacesCategoriesCategoryUuidGet Get Category

Gets the category information given the category ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param categoryUuid
 @return ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest
*/
func (a *CategoryApiService) GetCategoryPlacesCategoriesCategoryUuidGet(ctx context.Context, categoryUuid string) ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest {
	return ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest{
		ApiService: a,
		ctx: ctx,
		categoryUuid: categoryUuid,
	}
}

// Execute executes the request
//  @return CategoryBase
func (a *CategoryApiService) GetCategoryPlacesCategoriesCategoryUuidGetExecute(r ApiGetCategoryPlacesCategoriesCategoryUuidGetRequest) (*CategoryBase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CategoryBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoryApiService.GetCategoryPlacesCategoriesCategoryUuidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/categories/{category_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"category_uuid"+"}", url.PathEscape(parameterToString(r.categoryUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withMeta != nil {
		localVarQueryParams.Add("with_meta", parameterToString(*r.withMeta, ""))
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

type ApiSearchCategoriesPlacesCategoriesGetRequest struct {
	ctx context.Context
	ApiService *CategoryApiService
	projectUuid *string
	name *string
	withMeta *bool
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

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) ProjectUuid(projectUuid string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.projectUuid = &projectUuid
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) Name(name string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.name = &name
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) WithMeta(withMeta bool) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.withMeta = &withMeta
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) Page(page int32) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.page = &page
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) PageSize(pageSize int32) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) SortOn(sortOn string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) SortDesc(sortDesc bool) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) XAccessToken(xAccessToken string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) XSecretToken(xSecretToken string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) Authorization(authorization string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) EhelplyProject(ehelplyProject string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) EhelplyData(ehelplyData string) ApiSearchCategoriesPlacesCategoriesGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchCategoriesPlacesCategoriesGetRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchCategoriesPlacesCategoriesGetExecute(r)
}

/*
SearchCategoriesPlacesCategoriesGet Search Categories

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
 @return ApiSearchCategoriesPlacesCategoriesGetRequest
*/
func (a *CategoryApiService) SearchCategoriesPlacesCategoriesGet(ctx context.Context) ApiSearchCategoriesPlacesCategoriesGetRequest {
	return ApiSearchCategoriesPlacesCategoriesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Page
func (a *CategoryApiService) SearchCategoriesPlacesCategoriesGetExecute(r ApiSearchCategoriesPlacesCategoriesGetRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoryApiService.SearchCategoriesPlacesCategoriesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectUuid != nil {
		localVarQueryParams.Add("project_uuid", parameterToString(*r.projectUuid, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.withMeta != nil {
		localVarQueryParams.Add("with_meta", parameterToString(*r.withMeta, ""))
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

type ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest struct {
	ctx context.Context
	ApiService *CategoryApiService
	categoryUuid string
	categoryBase *CategoryBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) CategoryBase(categoryBase CategoryBase) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	r.categoryBase = &categoryBase
	return r
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) XAccessToken(xAccessToken string) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) XSecretToken(xSecretToken string) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) Authorization(authorization string) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) EhelplyProject(ehelplyProject string) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) EhelplyData(ehelplyData string) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) Execute() (*CategoryBase, *http.Response, error) {
	return r.ApiService.UpdateCategoryPlacesCategoriesCategoryUuidPutExecute(r)
}

/*
UpdateCategoryPlacesCategoriesCategoryUuidPut Update Category

Update category with given info, only updating the fields supplied. Category Uuid must be sent however.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param categoryUuid
 @return ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest
*/
func (a *CategoryApiService) UpdateCategoryPlacesCategoriesCategoryUuidPut(ctx context.Context, categoryUuid string) ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest {
	return ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest{
		ApiService: a,
		ctx: ctx,
		categoryUuid: categoryUuid,
	}
}

// Execute executes the request
//  @return CategoryBase
func (a *CategoryApiService) UpdateCategoryPlacesCategoriesCategoryUuidPutExecute(r ApiUpdateCategoryPlacesCategoriesCategoryUuidPutRequest) (*CategoryBase, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CategoryBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CategoryApiService.UpdateCategoryPlacesCategoriesCategoryUuidPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/categories/{category_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"category_uuid"+"}", url.PathEscape(parameterToString(r.categoryUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.categoryBase == nil {
		return localVarReturnValue, nil, reportError("categoryBase is required and must be specified")
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
	localVarPostBody = r.categoryBase
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
