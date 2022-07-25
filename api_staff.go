/*
eHelply SDK - 1.1.92

eHelply SDK for SuperStack Services

API version: 1.1.92
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


// StaffApiService StaffApi service
type StaffApiService service

type ApiCreateStaffPlacesStaffPostRequest struct {
	ctx context.Context
	ApiService *StaffApiService
	staffCreate *StaffCreate
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateStaffPlacesStaffPostRequest) StaffCreate(staffCreate StaffCreate) ApiCreateStaffPlacesStaffPostRequest {
	r.staffCreate = &staffCreate
	return r
}

func (r ApiCreateStaffPlacesStaffPostRequest) XAccessToken(xAccessToken string) ApiCreateStaffPlacesStaffPostRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateStaffPlacesStaffPostRequest) XSecretToken(xSecretToken string) ApiCreateStaffPlacesStaffPostRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateStaffPlacesStaffPostRequest) Authorization(authorization string) ApiCreateStaffPlacesStaffPostRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateStaffPlacesStaffPostRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateStaffPlacesStaffPostRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateStaffPlacesStaffPostRequest) EhelplyProject(ehelplyProject string) ApiCreateStaffPlacesStaffPostRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateStaffPlacesStaffPostRequest) EhelplyData(ehelplyData string) ApiCreateStaffPlacesStaffPostRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateStaffPlacesStaffPostRequest) Execute() (*StaffDb, *http.Response, error) {
	return r.ApiService.CreateStaffPlacesStaffPostExecute(r)
}

/*
CreateStaffPlacesStaffPost Create Staff

Creates a staff member

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateStaffPlacesStaffPostRequest
*/
func (a *StaffApiService) CreateStaffPlacesStaffPost(ctx context.Context) ApiCreateStaffPlacesStaffPostRequest {
	return ApiCreateStaffPlacesStaffPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StaffDb
func (a *StaffApiService) CreateStaffPlacesStaffPostExecute(r ApiCreateStaffPlacesStaffPostRequest) (*StaffDb, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StaffDb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.CreateStaffPlacesStaffPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/staff"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.staffCreate == nil {
		return localVarReturnValue, nil, reportError("staffCreate is required and must be specified")
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
	localVarPostBody = r.staffCreate
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

type ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest struct {
	ctx context.Context
	ApiService *StaffApiService
	staffUuid string
	softDelete *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) SoftDelete(softDelete bool) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	r.softDelete = &softDelete
	return r
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) XAccessToken(xAccessToken string) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) XSecretToken(xSecretToken string) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) Authorization(authorization string) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) EhelplyProject(ehelplyProject string) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) EhelplyData(ehelplyData string) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteStaffPlacesStaffStaffUuidDeleteExecute(r)
}

/*
DeleteStaffPlacesStaffStaffUuidDelete Delete Staff

Deletes the staff member with the given ID and returns True if successful

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param staffUuid
 @return ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest
*/
func (a *StaffApiService) DeleteStaffPlacesStaffStaffUuidDelete(ctx context.Context, staffUuid string) ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest {
	return ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest{
		ApiService: a,
		ctx: ctx,
		staffUuid: staffUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *StaffApiService) DeleteStaffPlacesStaffStaffUuidDeleteExecute(r ApiDeleteStaffPlacesStaffStaffUuidDeleteRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.DeleteStaffPlacesStaffStaffUuidDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/staff/{staff_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"staff_uuid"+"}", url.PathEscape(parameterToString(r.staffUuid, "")), -1)

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

type ApiGetStaffPlacesStaffStaffUuidGetRequest struct {
	ctx context.Context
	ApiService *StaffApiService
	staffUuid string
	withPlaces *bool
	withCompanies *bool
	withCatalog *bool
	withSchedule *bool
	withRoles *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) WithPlaces(withPlaces bool) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.withPlaces = &withPlaces
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) WithCompanies(withCompanies bool) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.withCompanies = &withCompanies
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) WithCatalog(withCatalog bool) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.withCatalog = &withCatalog
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) WithSchedule(withSchedule bool) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.withSchedule = &withSchedule
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) WithRoles(withRoles bool) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.withRoles = &withRoles
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) XAccessToken(xAccessToken string) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) XSecretToken(xSecretToken string) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) Authorization(authorization string) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) EhelplyProject(ehelplyProject string) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) EhelplyData(ehelplyData string) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetStaffPlacesStaffStaffUuidGetRequest) Execute() (*StaffResponse, *http.Response, error) {
	return r.ApiService.GetStaffPlacesStaffStaffUuidGetExecute(r)
}

/*
GetStaffPlacesStaffStaffUuidGet Get Staff

Gets the staff member information given the staff ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param staffUuid
 @return ApiGetStaffPlacesStaffStaffUuidGetRequest
*/
func (a *StaffApiService) GetStaffPlacesStaffStaffUuidGet(ctx context.Context, staffUuid string) ApiGetStaffPlacesStaffStaffUuidGetRequest {
	return ApiGetStaffPlacesStaffStaffUuidGetRequest{
		ApiService: a,
		ctx: ctx,
		staffUuid: staffUuid,
	}
}

// Execute executes the request
//  @return StaffResponse
func (a *StaffApiService) GetStaffPlacesStaffStaffUuidGetExecute(r ApiGetStaffPlacesStaffStaffUuidGetRequest) (*StaffResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StaffResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.GetStaffPlacesStaffStaffUuidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/staff/{staff_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"staff_uuid"+"}", url.PathEscape(parameterToString(r.staffUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withPlaces != nil {
		localVarQueryParams.Add("with_places", parameterToString(*r.withPlaces, ""))
	}
	if r.withCompanies != nil {
		localVarQueryParams.Add("with_companies", parameterToString(*r.withCompanies, ""))
	}
	if r.withCatalog != nil {
		localVarQueryParams.Add("with_catalog", parameterToString(*r.withCatalog, ""))
	}
	if r.withSchedule != nil {
		localVarQueryParams.Add("with_schedule", parameterToString(*r.withSchedule, ""))
	}
	if r.withRoles != nil {
		localVarQueryParams.Add("with_roles", parameterToString(*r.withRoles, ""))
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

type ApiSearchStaffPlacesStaffGetRequest struct {
	ctx context.Context
	ApiService *StaffApiService
	projectUuid *string
	firstName *string
	lastName *string
	isDeleted *bool
	withCompanies *bool
	withPlaces *bool
	withSchedule *bool
	withCatalog *bool
	withReviews *bool
	withRoles *bool
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

func (r ApiSearchStaffPlacesStaffGetRequest) ProjectUuid(projectUuid string) ApiSearchStaffPlacesStaffGetRequest {
	r.projectUuid = &projectUuid
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) FirstName(firstName string) ApiSearchStaffPlacesStaffGetRequest {
	r.firstName = &firstName
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) LastName(lastName string) ApiSearchStaffPlacesStaffGetRequest {
	r.lastName = &lastName
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) IsDeleted(isDeleted bool) ApiSearchStaffPlacesStaffGetRequest {
	r.isDeleted = &isDeleted
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) WithCompanies(withCompanies bool) ApiSearchStaffPlacesStaffGetRequest {
	r.withCompanies = &withCompanies
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) WithPlaces(withPlaces bool) ApiSearchStaffPlacesStaffGetRequest {
	r.withPlaces = &withPlaces
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) WithSchedule(withSchedule bool) ApiSearchStaffPlacesStaffGetRequest {
	r.withSchedule = &withSchedule
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) WithCatalog(withCatalog bool) ApiSearchStaffPlacesStaffGetRequest {
	r.withCatalog = &withCatalog
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) WithReviews(withReviews bool) ApiSearchStaffPlacesStaffGetRequest {
	r.withReviews = &withReviews
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) WithRoles(withRoles bool) ApiSearchStaffPlacesStaffGetRequest {
	r.withRoles = &withRoles
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) Page(page int32) ApiSearchStaffPlacesStaffGetRequest {
	r.page = &page
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) PageSize(pageSize int32) ApiSearchStaffPlacesStaffGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) SortOn(sortOn string) ApiSearchStaffPlacesStaffGetRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) SortDesc(sortDesc bool) ApiSearchStaffPlacesStaffGetRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) XAccessToken(xAccessToken string) ApiSearchStaffPlacesStaffGetRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) XSecretToken(xSecretToken string) ApiSearchStaffPlacesStaffGetRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) Authorization(authorization string) ApiSearchStaffPlacesStaffGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchStaffPlacesStaffGetRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) EhelplyProject(ehelplyProject string) ApiSearchStaffPlacesStaffGetRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) EhelplyData(ehelplyData string) ApiSearchStaffPlacesStaffGetRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchStaffPlacesStaffGetRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchStaffPlacesStaffGetExecute(r)
}

/*
SearchStaffPlacesStaffGet Search Staff

TODO
Item return format:
```
{
    uuid                                **type:** string
    project_uuid                        **type:** string or None

    entity                              **type:** string or None

    place                               **type:** dict or None

    company                             **type:** dict or None

    schedule                            **type:** dict or None

    catalog                             **type:** dict or None

    reviews                             **type:** dict or None

    created_at                          **type:** string or None

    updated_at                          **type:** string or None

    deleted_at                          **type:** string or None

}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchStaffPlacesStaffGetRequest
*/
func (a *StaffApiService) SearchStaffPlacesStaffGet(ctx context.Context) ApiSearchStaffPlacesStaffGetRequest {
	return ApiSearchStaffPlacesStaffGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Page
func (a *StaffApiService) SearchStaffPlacesStaffGetExecute(r ApiSearchStaffPlacesStaffGetRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.SearchStaffPlacesStaffGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/staff"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.projectUuid != nil {
		localVarQueryParams.Add("project_uuid", parameterToString(*r.projectUuid, ""))
	}
	if r.firstName != nil {
		localVarQueryParams.Add("first_name", parameterToString(*r.firstName, ""))
	}
	if r.lastName != nil {
		localVarQueryParams.Add("last_name", parameterToString(*r.lastName, ""))
	}
	if r.isDeleted != nil {
		localVarQueryParams.Add("is_deleted", parameterToString(*r.isDeleted, ""))
	}
	if r.withCompanies != nil {
		localVarQueryParams.Add("with_companies", parameterToString(*r.withCompanies, ""))
	}
	if r.withPlaces != nil {
		localVarQueryParams.Add("with_places", parameterToString(*r.withPlaces, ""))
	}
	if r.withSchedule != nil {
		localVarQueryParams.Add("with_schedule", parameterToString(*r.withSchedule, ""))
	}
	if r.withCatalog != nil {
		localVarQueryParams.Add("with_catalog", parameterToString(*r.withCatalog, ""))
	}
	if r.withReviews != nil {
		localVarQueryParams.Add("with_reviews", parameterToString(*r.withReviews, ""))
	}
	if r.withRoles != nil {
		localVarQueryParams.Add("with_roles", parameterToString(*r.withRoles, ""))
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

type ApiUpdateStaffPlacesStaffStaffUuidPutRequest struct {
	ctx context.Context
	ApiService *StaffApiService
	staffUuid string
	staffCreate *StaffCreate
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) StaffCreate(staffCreate StaffCreate) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	r.staffCreate = &staffCreate
	return r
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) XAccessToken(xAccessToken string) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) XSecretToken(xSecretToken string) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) Authorization(authorization string) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) EhelplyProject(ehelplyProject string) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) EhelplyData(ehelplyData string) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) Execute() (*StaffResponse, *http.Response, error) {
	return r.ApiService.UpdateStaffPlacesStaffStaffUuidPutExecute(r)
}

/*
UpdateStaffPlacesStaffStaffUuidPut Update Staff

Update staff with given info, only updating the fields supplied. Staff Uuid must be sent however.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param staffUuid
 @return ApiUpdateStaffPlacesStaffStaffUuidPutRequest
*/
func (a *StaffApiService) UpdateStaffPlacesStaffStaffUuidPut(ctx context.Context, staffUuid string) ApiUpdateStaffPlacesStaffStaffUuidPutRequest {
	return ApiUpdateStaffPlacesStaffStaffUuidPutRequest{
		ApiService: a,
		ctx: ctx,
		staffUuid: staffUuid,
	}
}

// Execute executes the request
//  @return StaffResponse
func (a *StaffApiService) UpdateStaffPlacesStaffStaffUuidPutExecute(r ApiUpdateStaffPlacesStaffStaffUuidPutRequest) (*StaffResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StaffResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.UpdateStaffPlacesStaffStaffUuidPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/places/staff/{staff_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"staff_uuid"+"}", url.PathEscape(parameterToString(r.staffUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.staffCreate == nil {
		return localVarReturnValue, nil, reportError("staffCreate is required and must be specified")
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
	localVarPostBody = r.staffCreate
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
