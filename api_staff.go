/*
eHelply SDK - 1.1.103

eHelply SDK for SuperStack Services

API version: 1.1.103
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

type ApiCreateStaffRequest struct {
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

func (r ApiCreateStaffRequest) StaffCreate(staffCreate StaffCreate) ApiCreateStaffRequest {
	r.staffCreate = &staffCreate
	return r
}

func (r ApiCreateStaffRequest) XAccessToken(xAccessToken string) ApiCreateStaffRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateStaffRequest) XSecretToken(xSecretToken string) ApiCreateStaffRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateStaffRequest) Authorization(authorization string) ApiCreateStaffRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateStaffRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateStaffRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateStaffRequest) EhelplyProject(ehelplyProject string) ApiCreateStaffRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateStaffRequest) EhelplyData(ehelplyData string) ApiCreateStaffRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateStaffRequest) Execute() (*StaffDb, *http.Response, error) {
	return r.ApiService.CreateStaffExecute(r)
}

/*
CreateStaff Createstaff

Creates a staff member

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateStaffRequest
*/
func (a *StaffApiService) CreateStaff(ctx context.Context) ApiCreateStaffRequest {
	return ApiCreateStaffRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StaffDb
func (a *StaffApiService) CreateStaffExecute(r ApiCreateStaffRequest) (*StaffDb, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StaffDb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.CreateStaff")
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

type ApiDeleteStaffRequest struct {
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

func (r ApiDeleteStaffRequest) SoftDelete(softDelete bool) ApiDeleteStaffRequest {
	r.softDelete = &softDelete
	return r
}

func (r ApiDeleteStaffRequest) XAccessToken(xAccessToken string) ApiDeleteStaffRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteStaffRequest) XSecretToken(xSecretToken string) ApiDeleteStaffRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteStaffRequest) Authorization(authorization string) ApiDeleteStaffRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteStaffRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteStaffRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteStaffRequest) EhelplyProject(ehelplyProject string) ApiDeleteStaffRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteStaffRequest) EhelplyData(ehelplyData string) ApiDeleteStaffRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteStaffRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteStaffExecute(r)
}

/*
DeleteStaff Deletestaff

Deletes the staff member with the given ID and returns True if successful

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param staffUuid
 @return ApiDeleteStaffRequest
*/
func (a *StaffApiService) DeleteStaff(ctx context.Context, staffUuid string) ApiDeleteStaffRequest {
	return ApiDeleteStaffRequest{
		ApiService: a,
		ctx: ctx,
		staffUuid: staffUuid,
	}
}

// Execute executes the request
//  @return interface{}
func (a *StaffApiService) DeleteStaffExecute(r ApiDeleteStaffRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.DeleteStaff")
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

type ApiGetStaffRequest struct {
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

func (r ApiGetStaffRequest) WithPlaces(withPlaces bool) ApiGetStaffRequest {
	r.withPlaces = &withPlaces
	return r
}

func (r ApiGetStaffRequest) WithCompanies(withCompanies bool) ApiGetStaffRequest {
	r.withCompanies = &withCompanies
	return r
}

func (r ApiGetStaffRequest) WithCatalog(withCatalog bool) ApiGetStaffRequest {
	r.withCatalog = &withCatalog
	return r
}

func (r ApiGetStaffRequest) WithSchedule(withSchedule bool) ApiGetStaffRequest {
	r.withSchedule = &withSchedule
	return r
}

func (r ApiGetStaffRequest) WithRoles(withRoles bool) ApiGetStaffRequest {
	r.withRoles = &withRoles
	return r
}

func (r ApiGetStaffRequest) XAccessToken(xAccessToken string) ApiGetStaffRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetStaffRequest) XSecretToken(xSecretToken string) ApiGetStaffRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetStaffRequest) Authorization(authorization string) ApiGetStaffRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetStaffRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetStaffRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetStaffRequest) EhelplyProject(ehelplyProject string) ApiGetStaffRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetStaffRequest) EhelplyData(ehelplyData string) ApiGetStaffRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetStaffRequest) Execute() (*StaffResponse, *http.Response, error) {
	return r.ApiService.GetStaffExecute(r)
}

/*
GetStaff Getstaff

Gets the staff member information given the staff ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param staffUuid
 @return ApiGetStaffRequest
*/
func (a *StaffApiService) GetStaff(ctx context.Context, staffUuid string) ApiGetStaffRequest {
	return ApiGetStaffRequest{
		ApiService: a,
		ctx: ctx,
		staffUuid: staffUuid,
	}
}

// Execute executes the request
//  @return StaffResponse
func (a *StaffApiService) GetStaffExecute(r ApiGetStaffRequest) (*StaffResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StaffResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.GetStaff")
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

type ApiSearchStaffRequest struct {
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

func (r ApiSearchStaffRequest) ProjectUuid(projectUuid string) ApiSearchStaffRequest {
	r.projectUuid = &projectUuid
	return r
}

func (r ApiSearchStaffRequest) FirstName(firstName string) ApiSearchStaffRequest {
	r.firstName = &firstName
	return r
}

func (r ApiSearchStaffRequest) LastName(lastName string) ApiSearchStaffRequest {
	r.lastName = &lastName
	return r
}

func (r ApiSearchStaffRequest) IsDeleted(isDeleted bool) ApiSearchStaffRequest {
	r.isDeleted = &isDeleted
	return r
}

func (r ApiSearchStaffRequest) WithCompanies(withCompanies bool) ApiSearchStaffRequest {
	r.withCompanies = &withCompanies
	return r
}

func (r ApiSearchStaffRequest) WithPlaces(withPlaces bool) ApiSearchStaffRequest {
	r.withPlaces = &withPlaces
	return r
}

func (r ApiSearchStaffRequest) WithSchedule(withSchedule bool) ApiSearchStaffRequest {
	r.withSchedule = &withSchedule
	return r
}

func (r ApiSearchStaffRequest) WithCatalog(withCatalog bool) ApiSearchStaffRequest {
	r.withCatalog = &withCatalog
	return r
}

func (r ApiSearchStaffRequest) WithReviews(withReviews bool) ApiSearchStaffRequest {
	r.withReviews = &withReviews
	return r
}

func (r ApiSearchStaffRequest) WithRoles(withRoles bool) ApiSearchStaffRequest {
	r.withRoles = &withRoles
	return r
}

func (r ApiSearchStaffRequest) Page(page int32) ApiSearchStaffRequest {
	r.page = &page
	return r
}

func (r ApiSearchStaffRequest) PageSize(pageSize int32) ApiSearchStaffRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchStaffRequest) SortOn(sortOn string) ApiSearchStaffRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchStaffRequest) SortDesc(sortDesc bool) ApiSearchStaffRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchStaffRequest) XAccessToken(xAccessToken string) ApiSearchStaffRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchStaffRequest) XSecretToken(xSecretToken string) ApiSearchStaffRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchStaffRequest) Authorization(authorization string) ApiSearchStaffRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchStaffRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchStaffRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchStaffRequest) EhelplyProject(ehelplyProject string) ApiSearchStaffRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchStaffRequest) EhelplyData(ehelplyData string) ApiSearchStaffRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchStaffRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchStaffExecute(r)
}

/*
SearchStaff Searchstaff

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
 @return ApiSearchStaffRequest
*/
func (a *StaffApiService) SearchStaff(ctx context.Context) ApiSearchStaffRequest {
	return ApiSearchStaffRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Page
func (a *StaffApiService) SearchStaffExecute(r ApiSearchStaffRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.SearchStaff")
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

type ApiUpdateStaffRequest struct {
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

func (r ApiUpdateStaffRequest) StaffCreate(staffCreate StaffCreate) ApiUpdateStaffRequest {
	r.staffCreate = &staffCreate
	return r
}

func (r ApiUpdateStaffRequest) XAccessToken(xAccessToken string) ApiUpdateStaffRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateStaffRequest) XSecretToken(xSecretToken string) ApiUpdateStaffRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateStaffRequest) Authorization(authorization string) ApiUpdateStaffRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateStaffRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateStaffRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateStaffRequest) EhelplyProject(ehelplyProject string) ApiUpdateStaffRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateStaffRequest) EhelplyData(ehelplyData string) ApiUpdateStaffRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateStaffRequest) Execute() (*StaffResponse, *http.Response, error) {
	return r.ApiService.UpdateStaffExecute(r)
}

/*
UpdateStaff Updatestaff

Update staff with given info, only updating the fields supplied. Staff Uuid must be sent however.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param staffUuid
 @return ApiUpdateStaffRequest
*/
func (a *StaffApiService) UpdateStaff(ctx context.Context, staffUuid string) ApiUpdateStaffRequest {
	return ApiUpdateStaffRequest{
		ApiService: a,
		ctx: ctx,
		staffUuid: staffUuid,
	}
}

// Execute executes the request
//  @return StaffResponse
func (a *StaffApiService) UpdateStaffExecute(r ApiUpdateStaffRequest) (*StaffResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StaffResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaffApiService.UpdateStaff")
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
