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


// CompaniesApiService CompaniesApi service
type CompaniesApiService service

type ApiCreateCompanyPlacesCompaniesPostRequest struct {
	ctx context.Context
	ApiService *CompaniesApiService
	companyBase *CompanyBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) CompanyBase(companyBase CompanyBase) ApiCreateCompanyPlacesCompaniesPostRequest {
	r.companyBase = &companyBase
	return r
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) XAccessToken(xAccessToken string) ApiCreateCompanyPlacesCompaniesPostRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) XSecretToken(xSecretToken string) ApiCreateCompanyPlacesCompaniesPostRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) Authorization(authorization string) ApiCreateCompanyPlacesCompaniesPostRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateCompanyPlacesCompaniesPostRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) EhelplyProject(ehelplyProject string) ApiCreateCompanyPlacesCompaniesPostRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) EhelplyData(ehelplyData string) ApiCreateCompanyPlacesCompaniesPostRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateCompanyPlacesCompaniesPostRequest) Execute() (*CompanyResponse, *http.Response, error) {
	return r.ApiService.CreateCompanyPlacesCompaniesPostExecute(r)
}

/*
CreateCompanyPlacesCompaniesPost Create Company

Creates a company

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCompanyPlacesCompaniesPostRequest
*/
func (a *CompaniesApiService) CreateCompanyPlacesCompaniesPost(ctx context.Context) ApiCreateCompanyPlacesCompaniesPostRequest {
	return ApiCreateCompanyPlacesCompaniesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CompanyResponse
func (a *CompaniesApiService) CreateCompanyPlacesCompaniesPostExecute(r ApiCreateCompanyPlacesCompaniesPostRequest) (*CompanyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompaniesApiService.CreateCompanyPlacesCompaniesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/companies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.companyBase == nil {
		return localVarReturnValue, nil, reportError("companyBase is required and must be specified")
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
	localVarPostBody = r.companyBase
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

type ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest struct {
	ctx context.Context
	ApiService *CompaniesApiService
	companyUuid string
	softDelete *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) SoftDelete(softDelete bool) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	r.softDelete = &softDelete
	return r
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) XAccessToken(xAccessToken string) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) XSecretToken(xSecretToken string) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) Authorization(authorization string) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) EhelplyProject(ehelplyProject string) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) EhelplyData(ehelplyData string) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeletePlacePlacesCompaniesCompanyUuidDeleteExecute(r)
}

/*
DeletePlacePlacesCompaniesCompanyUuidDelete Delete Place

Deletes the company with the given ID and returns True if successful

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyUuid
 @return ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest
*/
func (a *CompaniesApiService) DeletePlacePlacesCompaniesCompanyUuidDelete(ctx context.Context, companyUuid string) ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest {
	return ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest{
		ApiService: a,
		ctx: ctx,
		companyUuid: companyUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *CompaniesApiService) DeletePlacePlacesCompaniesCompanyUuidDeleteExecute(r ApiDeletePlacePlacesCompaniesCompanyUuidDeleteRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompaniesApiService.DeletePlacePlacesCompaniesCompanyUuidDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/companies/{company_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_uuid"+"}", url.PathEscape(parameterToString(r.companyUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.softDelete != nil {
		localVarQueryParams.Add("soft_delete", parameterToString(*r.softDelete, ""))
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

type ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest struct {
	ctx context.Context
	ApiService *CompaniesApiService
	companyUuid string
	withMeta *bool
	withCatalog *bool
	withReviews *bool
	withSchedule *bool
	withBlog *bool
	withTags *bool
	withCategories *bool
	withPlaces *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithMeta(withMeta bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withMeta = &withMeta
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithCatalog(withCatalog bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withCatalog = &withCatalog
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithReviews(withReviews bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withReviews = &withReviews
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithSchedule(withSchedule bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withSchedule = &withSchedule
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithBlog(withBlog bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withBlog = &withBlog
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithTags(withTags bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withTags = &withTags
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithCategories(withCategories bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withCategories = &withCategories
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) WithPlaces(withPlaces bool) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.withPlaces = &withPlaces
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) XAccessToken(xAccessToken string) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) XSecretToken(xSecretToken string) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) Authorization(authorization string) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) EhelplyProject(ehelplyProject string) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) EhelplyData(ehelplyData string) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) Execute() (*CompanyResponse, *http.Response, error) {
	return r.ApiService.GetCompanyPlacesCompaniesCompanyUuidGetExecute(r)
}

/*
GetCompanyPlacesCompaniesCompanyUuidGet Get Company

Gets the company information given the Place ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyUuid
 @return ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest
*/
func (a *CompaniesApiService) GetCompanyPlacesCompaniesCompanyUuidGet(ctx context.Context, companyUuid string) ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest {
	return ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest{
		ApiService: a,
		ctx: ctx,
		companyUuid: companyUuid,
	}
}

// Execute executes the request
//  @return CompanyResponse
func (a *CompaniesApiService) GetCompanyPlacesCompaniesCompanyUuidGetExecute(r ApiGetCompanyPlacesCompaniesCompanyUuidGetRequest) (*CompanyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompaniesApiService.GetCompanyPlacesCompaniesCompanyUuidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/companies/{company_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_uuid"+"}", url.PathEscape(parameterToString(r.companyUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withMeta != nil {
		localVarQueryParams.Add("with_meta", parameterToString(*r.withMeta, ""))
	}
	if r.withCatalog != nil {
		localVarQueryParams.Add("with_catalog", parameterToString(*r.withCatalog, ""))
	}
	if r.withReviews != nil {
		localVarQueryParams.Add("with_reviews", parameterToString(*r.withReviews, ""))
	}
	if r.withSchedule != nil {
		localVarQueryParams.Add("with_schedule", parameterToString(*r.withSchedule, ""))
	}
	if r.withBlog != nil {
		localVarQueryParams.Add("with_blog", parameterToString(*r.withBlog, ""))
	}
	if r.withTags != nil {
		localVarQueryParams.Add("with_tags", parameterToString(*r.withTags, ""))
	}
	if r.withCategories != nil {
		localVarQueryParams.Add("with_categories", parameterToString(*r.withCategories, ""))
	}
	if r.withPlaces != nil {
		localVarQueryParams.Add("with_places", parameterToString(*r.withPlaces, ""))
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

type ApiSearchCompaniesPlacesCompaniesGetRequest struct {
	ctx context.Context
	ApiService *CompaniesApiService
	projectUuid *string
	name *string
	email *string
	isPublic *bool
	isDeleted *bool
	withPlaces *bool
	withMeta *bool
	withCatalog *bool
	withReviews *bool
	withSchedule *bool
	withBlog *bool
	withTags *bool
	withCategories *bool
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

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) ProjectUuid(projectUuid string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.projectUuid = &projectUuid
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) Name(name string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.name = &name
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) Email(email string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.email = &email
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) IsPublic(isPublic bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.isPublic = &isPublic
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) IsDeleted(isDeleted bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.isDeleted = &isDeleted
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithPlaces(withPlaces bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withPlaces = &withPlaces
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithMeta(withMeta bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withMeta = &withMeta
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithCatalog(withCatalog bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withCatalog = &withCatalog
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithReviews(withReviews bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withReviews = &withReviews
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithSchedule(withSchedule bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withSchedule = &withSchedule
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithBlog(withBlog bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withBlog = &withBlog
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithTags(withTags bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withTags = &withTags
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) WithCategories(withCategories bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.withCategories = &withCategories
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) Page(page int32) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.page = &page
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) PageSize(pageSize int32) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) SortOn(sortOn string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) SortDesc(sortDesc bool) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) XAccessToken(xAccessToken string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) XSecretToken(xSecretToken string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) Authorization(authorization string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) EhelplyProject(ehelplyProject string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) EhelplyData(ehelplyData string) ApiSearchCompaniesPlacesCompaniesGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchCompaniesPlacesCompaniesGetRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchCompaniesPlacesCompaniesGetExecute(r)
}

/*
SearchCompaniesPlacesCompaniesGet Search Companies

Search all companies and returns paginated results with Companies being stored in items field.
Can search by `project_uuid, name, email` string fields or the `is_public and is_deleted` boolean fields.
To search with these fields use query params with string values. For sorting fill out "sort_desc" field with either
true/false and the "sort_on" query parameter
with column you want to sort on (ex: name). Max pagination items per page is 50.
Item return format:
```
{
    uuid                                **type:** string
    project_uuid                        **type:** string or None

    meta_uuid                           **type:** string or None

    catalog_data                        **type:** dict or None

    review_group_data                   **type:** dict or None

    schedule_data                       **type:** dict or None

    blog_data                           **type:** dict or None

    tags                                **type:** [TagBase] or None

    categories                          **type:** [CategoryBase] or None

    places                              **type:** PlaceBase or None

    created_at                          **type:** string or None

    updated_at                          **type:** string or None

    deleted_at                          **type:** string or None

}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchCompaniesPlacesCompaniesGetRequest
*/
func (a *CompaniesApiService) SearchCompaniesPlacesCompaniesGet(ctx context.Context) ApiSearchCompaniesPlacesCompaniesGetRequest {
	return ApiSearchCompaniesPlacesCompaniesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Page
func (a *CompaniesApiService) SearchCompaniesPlacesCompaniesGetExecute(r ApiSearchCompaniesPlacesCompaniesGetRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompaniesApiService.SearchCompaniesPlacesCompaniesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/companies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectUuid != nil {
		localVarQueryParams.Add("project_uuid", parameterToString(*r.projectUuid, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.email != nil {
		localVarQueryParams.Add("email", parameterToString(*r.email, ""))
	}
	if r.isPublic != nil {
		localVarQueryParams.Add("is_public", parameterToString(*r.isPublic, ""))
	}
	if r.isDeleted != nil {
		localVarQueryParams.Add("is_deleted", parameterToString(*r.isDeleted, ""))
	}
	if r.withPlaces != nil {
		localVarQueryParams.Add("with_places", parameterToString(*r.withPlaces, ""))
	}
	if r.withMeta != nil {
		localVarQueryParams.Add("with_meta", parameterToString(*r.withMeta, ""))
	}
	if r.withCatalog != nil {
		localVarQueryParams.Add("with_catalog", parameterToString(*r.withCatalog, ""))
	}
	if r.withReviews != nil {
		localVarQueryParams.Add("with_reviews", parameterToString(*r.withReviews, ""))
	}
	if r.withSchedule != nil {
		localVarQueryParams.Add("with_schedule", parameterToString(*r.withSchedule, ""))
	}
	if r.withBlog != nil {
		localVarQueryParams.Add("with_blog", parameterToString(*r.withBlog, ""))
	}
	if r.withTags != nil {
		localVarQueryParams.Add("with_tags", parameterToString(*r.withTags, ""))
	}
	if r.withCategories != nil {
		localVarQueryParams.Add("with_categories", parameterToString(*r.withCategories, ""))
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

type ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest struct {
	ctx context.Context
	ApiService *CompaniesApiService
	companyUuid string
	companyBase *CompanyBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) CompanyBase(companyBase CompanyBase) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	r.companyBase = &companyBase
	return r
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) XAccessToken(xAccessToken string) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) XSecretToken(xSecretToken string) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) Authorization(authorization string) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) EhelplyProject(ehelplyProject string) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) EhelplyData(ehelplyData string) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) Execute() (*CompanyResponse, *http.Response, error) {
	return r.ApiService.UpdateCompanyPlacesCompaniesCompanyUuidPutExecute(r)
}

/*
UpdateCompanyPlacesCompaniesCompanyUuidPut Update Company

Update company with given info, only updating the fields supplied. Company Uuid must be sent however.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyUuid
 @return ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest
*/
func (a *CompaniesApiService) UpdateCompanyPlacesCompaniesCompanyUuidPut(ctx context.Context, companyUuid string) ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest {
	return ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest{
		ApiService: a,
		ctx: ctx,
		companyUuid: companyUuid,
	}
}

// Execute executes the request
//  @return CompanyResponse
func (a *CompaniesApiService) UpdateCompanyPlacesCompaniesCompanyUuidPutExecute(r ApiUpdateCompanyPlacesCompaniesCompanyUuidPutRequest) (*CompanyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompaniesApiService.UpdateCompanyPlacesCompaniesCompanyUuidPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/companies/{company_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"company_uuid"+"}", url.PathEscape(parameterToString(r.companyUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.companyBase == nil {
		return localVarReturnValue, nil, reportError("companyBase is required and must be specified")
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
	localVarPostBody = r.companyBase
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
