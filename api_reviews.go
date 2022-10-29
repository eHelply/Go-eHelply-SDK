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


// ReviewsApiService ReviewsApi service
type ReviewsApiService service

type ApiCreateReviewRequest struct {
	ctx context.Context
	ApiService *ReviewsApiService
	entityType string
	entityUuid string
	createReview *CreateReview
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateReviewRequest) CreateReview(createReview CreateReview) ApiCreateReviewRequest {
	r.createReview = &createReview
	return r
}

func (r ApiCreateReviewRequest) XAccessToken(xAccessToken string) ApiCreateReviewRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateReviewRequest) XSecretToken(xSecretToken string) ApiCreateReviewRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateReviewRequest) Authorization(authorization string) ApiCreateReviewRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateReviewRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateReviewRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateReviewRequest) EhelplyProject(ehelplyProject string) ApiCreateReviewRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateReviewRequest) EhelplyData(ehelplyData string) ApiCreateReviewRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateReviewRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.CreateReviewExecute(r)
}

/*
CreateReview Create

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityType
 @param entityUuid
 @return ApiCreateReviewRequest
*/
func (a *ReviewsApiService) CreateReview(ctx context.Context, entityType string, entityUuid string) ApiCreateReviewRequest {
	return ApiCreateReviewRequest{
		ApiService: a,
		ctx: ctx,
		entityType: entityType,
		entityUuid: entityUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ReviewsApiService) CreateReviewExecute(r ApiCreateReviewRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsApiService.CreateReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/reviews/types/{entity_type}/entities/{entity_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_type"+"}", url.PathEscape(parameterToString(r.entityType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"entity_uuid"+"}", url.PathEscape(parameterToString(r.entityUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createReview == nil {
		return localVarReturnValue, nil, reportError("createReview is required and must be specified")
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
	localVarPostBody = r.createReview
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

type ApiDeleteReviewRequest struct {
	ctx context.Context
	ApiService *ReviewsApiService
	entityType string
	entityUuid string
	reviewUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteReviewRequest) XAccessToken(xAccessToken string) ApiDeleteReviewRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteReviewRequest) XSecretToken(xSecretToken string) ApiDeleteReviewRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteReviewRequest) Authorization(authorization string) ApiDeleteReviewRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteReviewRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteReviewRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteReviewRequest) EhelplyProject(ehelplyProject string) ApiDeleteReviewRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteReviewRequest) EhelplyData(ehelplyData string) ApiDeleteReviewRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteReviewRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteReviewExecute(r)
}

/*
DeleteReview Deletereview

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityType
 @param entityUuid
 @param reviewUuid
 @return ApiDeleteReviewRequest
*/
func (a *ReviewsApiService) DeleteReview(ctx context.Context, entityType string, entityUuid string, reviewUuid string) ApiDeleteReviewRequest {
	return ApiDeleteReviewRequest{
		ApiService: a,
		ctx: ctx,
		entityType: entityType,
		entityUuid: entityUuid,
		reviewUuid: reviewUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ReviewsApiService) DeleteReviewExecute(r ApiDeleteReviewRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsApiService.DeleteReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/reviews/types/{entity_type}/entities/{entity_uuid}/reviews/{review_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_type"+"}", url.PathEscape(parameterToString(r.entityType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"entity_uuid"+"}", url.PathEscape(parameterToString(r.entityUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"review_uuid"+"}", url.PathEscape(parameterToString(r.reviewUuid, "")), -1)

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

type ApiGetReviewRequest struct {
	ctx context.Context
	ApiService *ReviewsApiService
	entityType string
	entityUuid string
	reviewUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetReviewRequest) XAccessToken(xAccessToken string) ApiGetReviewRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetReviewRequest) XSecretToken(xSecretToken string) ApiGetReviewRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetReviewRequest) Authorization(authorization string) ApiGetReviewRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetReviewRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetReviewRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetReviewRequest) EhelplyProject(ehelplyProject string) ApiGetReviewRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetReviewRequest) EhelplyData(ehelplyData string) ApiGetReviewRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetReviewRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.GetReviewExecute(r)
}

/*
GetReview Getreview

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityType
 @param entityUuid
 @param reviewUuid
 @return ApiGetReviewRequest
*/
func (a *ReviewsApiService) GetReview(ctx context.Context, entityType string, entityUuid string, reviewUuid string) ApiGetReviewRequest {
	return ApiGetReviewRequest{
		ApiService: a,
		ctx: ctx,
		entityType: entityType,
		entityUuid: entityUuid,
		reviewUuid: reviewUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ReviewsApiService) GetReviewExecute(r ApiGetReviewRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsApiService.GetReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/reviews/types/{entity_type}/entities/{entity_uuid}/reviews/{review_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_type"+"}", url.PathEscape(parameterToString(r.entityType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"entity_uuid"+"}", url.PathEscape(parameterToString(r.entityUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"review_uuid"+"}", url.PathEscape(parameterToString(r.reviewUuid, "")), -1)

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

type ApiSearchReviewsRequest struct {
	ctx context.Context
	ApiService *ReviewsApiService
	entityType string
	entityUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiSearchReviewsRequest) XAccessToken(xAccessToken string) ApiSearchReviewsRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchReviewsRequest) XSecretToken(xSecretToken string) ApiSearchReviewsRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchReviewsRequest) Authorization(authorization string) ApiSearchReviewsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchReviewsRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchReviewsRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchReviewsRequest) EhelplyProject(ehelplyProject string) ApiSearchReviewsRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchReviewsRequest) EhelplyData(ehelplyData string) ApiSearchReviewsRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchReviewsRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.SearchReviewsExecute(r)
}

/*
SearchReviews Searchreview

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityType
 @param entityUuid
 @return ApiSearchReviewsRequest
*/
func (a *ReviewsApiService) SearchReviews(ctx context.Context, entityType string, entityUuid string) ApiSearchReviewsRequest {
	return ApiSearchReviewsRequest{
		ApiService: a,
		ctx: ctx,
		entityType: entityType,
		entityUuid: entityUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ReviewsApiService) SearchReviewsExecute(r ApiSearchReviewsRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsApiService.SearchReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/reviews/types/{entity_type}/entities/{entity_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_type"+"}", url.PathEscape(parameterToString(r.entityType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"entity_uuid"+"}", url.PathEscape(parameterToString(r.entityUuid, "")), -1)

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

type ApiUpdateReviewRequest struct {
	ctx context.Context
	ApiService *ReviewsApiService
	entityType string
	entityUuid string
	reviewUuid string
	updateReview *UpdateReview
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateReviewRequest) UpdateReview(updateReview UpdateReview) ApiUpdateReviewRequest {
	r.updateReview = &updateReview
	return r
}

func (r ApiUpdateReviewRequest) XAccessToken(xAccessToken string) ApiUpdateReviewRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateReviewRequest) XSecretToken(xSecretToken string) ApiUpdateReviewRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateReviewRequest) Authorization(authorization string) ApiUpdateReviewRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateReviewRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateReviewRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateReviewRequest) EhelplyProject(ehelplyProject string) ApiUpdateReviewRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateReviewRequest) EhelplyData(ehelplyData string) ApiUpdateReviewRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateReviewRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.UpdateReviewExecute(r)
}

/*
UpdateReview Updatereview

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityType
 @param entityUuid
 @param reviewUuid
 @return ApiUpdateReviewRequest
*/
func (a *ReviewsApiService) UpdateReview(ctx context.Context, entityType string, entityUuid string, reviewUuid string) ApiUpdateReviewRequest {
	return ApiUpdateReviewRequest{
		ApiService: a,
		ctx: ctx,
		entityType: entityType,
		entityUuid: entityUuid,
		reviewUuid: reviewUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *ReviewsApiService) UpdateReviewExecute(r ApiUpdateReviewRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsApiService.UpdateReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/reviews/types/{entity_type}/entities/{entity_uuid}/reviews/{review_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_type"+"}", url.PathEscape(parameterToString(r.entityType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"entity_uuid"+"}", url.PathEscape(parameterToString(r.entityUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"review_uuid"+"}", url.PathEscape(parameterToString(r.reviewUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateReview == nil {
		return localVarReturnValue, nil, reportError("updateReview is required and must be specified")
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
	localVarPostBody = r.updateReview
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
