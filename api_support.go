/*
eHelply SDK - 1.1.110

eHelply SDK for SuperStack Services

API version: 1.1.110
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


// SupportApiService SupportApi service
type SupportApiService service

type ApiCreateContactRequest struct {
	ctx context.Context
	ApiService *SupportApiService
	contact *Contact
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateContactRequest) Contact(contact Contact) ApiCreateContactRequest {
	r.contact = &contact
	return r
}

func (r ApiCreateContactRequest) XAccessToken(xAccessToken string) ApiCreateContactRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateContactRequest) XSecretToken(xSecretToken string) ApiCreateContactRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateContactRequest) Authorization(authorization string) ApiCreateContactRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateContactRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateContactRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateContactRequest) EhelplyProject(ehelplyProject string) ApiCreateContactRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateContactRequest) EhelplyData(ehelplyData string) ApiCreateContactRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateContactRequest) Execute() (*ContactResponse, *http.Response, error) {
	return r.ApiService.CreateContactExecute(r)
}

/*
CreateContact Createcontact

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateContactRequest
*/
func (a *SupportApiService) CreateContact(ctx context.Context) ApiCreateContactRequest {
	return ApiCreateContactRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ContactResponse
func (a *SupportApiService) CreateContactExecute(r ApiCreateContactRequest) (*ContactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportApiService.CreateContact")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/support/contact"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contact == nil {
		return localVarReturnValue, nil, reportError("contact is required and must be specified")
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
	localVarPostBody = r.contact
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

type ApiCreateTicketRequest struct {
	ctx context.Context
	ApiService *SupportApiService
	projectUuid string
	memberUuid string
	createTicket *CreateTicket
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateTicketRequest) CreateTicket(createTicket CreateTicket) ApiCreateTicketRequest {
	r.createTicket = &createTicket
	return r
}

func (r ApiCreateTicketRequest) XAccessToken(xAccessToken string) ApiCreateTicketRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateTicketRequest) XSecretToken(xSecretToken string) ApiCreateTicketRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateTicketRequest) Authorization(authorization string) ApiCreateTicketRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateTicketRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateTicketRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateTicketRequest) EhelplyProject(ehelplyProject string) ApiCreateTicketRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateTicketRequest) EhelplyData(ehelplyData string) ApiCreateTicketRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateTicketRequest) Execute() (*TicketResponse, *http.Response, error) {
	return r.ApiService.CreateTicketExecute(r)
}

/*
CreateTicket Createticket

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @param memberUuid
 @return ApiCreateTicketRequest
*/
func (a *SupportApiService) CreateTicket(ctx context.Context, projectUuid string, memberUuid string) ApiCreateTicketRequest {
	return ApiCreateTicketRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
		memberUuid: memberUuid,
	}
}

// Execute executes the request
//  @return TicketResponse
func (a *SupportApiService) CreateTicketExecute(r ApiCreateTicketRequest) (*TicketResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TicketResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportApiService.CreateTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/support/projects/{project_uuid}/members/{member_uuid}/tickets"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"member_uuid"+"}", url.PathEscape(parameterToString(r.memberUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createTicket == nil {
		return localVarReturnValue, nil, reportError("createTicket is required and must be specified")
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
	localVarPostBody = r.createTicket
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

type ApiListTicketsRequest struct {
	ctx context.Context
	ApiService *SupportApiService
	projectUuid string
	memberUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiListTicketsRequest) XAccessToken(xAccessToken string) ApiListTicketsRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiListTicketsRequest) XSecretToken(xSecretToken string) ApiListTicketsRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiListTicketsRequest) Authorization(authorization string) ApiListTicketsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiListTicketsRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiListTicketsRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiListTicketsRequest) EhelplyProject(ehelplyProject string) ApiListTicketsRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiListTicketsRequest) EhelplyData(ehelplyData string) ApiListTicketsRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiListTicketsRequest) Execute() ([]TicketsResponse, *http.Response, error) {
	return r.ApiService.ListTicketsExecute(r)
}

/*
ListTickets Listtickets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @param memberUuid
 @return ApiListTicketsRequest
*/
func (a *SupportApiService) ListTickets(ctx context.Context, projectUuid string, memberUuid string) ApiListTicketsRequest {
	return ApiListTicketsRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
		memberUuid: memberUuid,
	}
}

// Execute executes the request
//  @return []TicketsResponse
func (a *SupportApiService) ListTicketsExecute(r ApiListTicketsRequest) ([]TicketsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TicketsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportApiService.ListTickets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/support/projects/{project_uuid}/members/{member_uuid}/tickets"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"member_uuid"+"}", url.PathEscape(parameterToString(r.memberUuid, "")), -1)

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

type ApiUpdateTicketRequest struct {
	ctx context.Context
	ApiService *SupportApiService
	projectUuid string
	memberUuid string
	ticketId string
	createTicket *CreateTicket
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateTicketRequest) CreateTicket(createTicket CreateTicket) ApiUpdateTicketRequest {
	r.createTicket = &createTicket
	return r
}

func (r ApiUpdateTicketRequest) XAccessToken(xAccessToken string) ApiUpdateTicketRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateTicketRequest) XSecretToken(xSecretToken string) ApiUpdateTicketRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateTicketRequest) Authorization(authorization string) ApiUpdateTicketRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateTicketRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateTicketRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateTicketRequest) EhelplyProject(ehelplyProject string) ApiUpdateTicketRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateTicketRequest) EhelplyData(ehelplyData string) ApiUpdateTicketRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateTicketRequest) Execute() (*TicketResponse, *http.Response, error) {
	return r.ApiService.UpdateTicketExecute(r)
}

/*
UpdateTicket Updateticket

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @param memberUuid
 @param ticketId
 @return ApiUpdateTicketRequest
*/
func (a *SupportApiService) UpdateTicket(ctx context.Context, projectUuid string, memberUuid string, ticketId string) ApiUpdateTicketRequest {
	return ApiUpdateTicketRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
		memberUuid: memberUuid,
		ticketId: ticketId,
	}
}

// Execute executes the request
//  @return TicketResponse
func (a *SupportApiService) UpdateTicketExecute(r ApiUpdateTicketRequest) (*TicketResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TicketResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportApiService.UpdateTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/support/projects/{project_uuid}/members/{member_uuid}/tickets/{ticket_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"member_uuid"+"}", url.PathEscape(parameterToString(r.memberUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ticket_id"+"}", url.PathEscape(parameterToString(r.ticketId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createTicket == nil {
		return localVarReturnValue, nil, reportError("createTicket is required and must be specified")
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
	localVarPostBody = r.createTicket
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

type ApiViewTicketRequest struct {
	ctx context.Context
	ApiService *SupportApiService
	projectUuid string
	memberUuid string
	ticketId string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiViewTicketRequest) XAccessToken(xAccessToken string) ApiViewTicketRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiViewTicketRequest) XSecretToken(xSecretToken string) ApiViewTicketRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiViewTicketRequest) Authorization(authorization string) ApiViewTicketRequest {
	r.authorization = &authorization
	return r
}

func (r ApiViewTicketRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiViewTicketRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiViewTicketRequest) EhelplyProject(ehelplyProject string) ApiViewTicketRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiViewTicketRequest) EhelplyData(ehelplyData string) ApiViewTicketRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiViewTicketRequest) Execute() (*TicketResponse, *http.Response, error) {
	return r.ApiService.ViewTicketExecute(r)
}

/*
ViewTicket Viewticket

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @param memberUuid
 @param ticketId
 @return ApiViewTicketRequest
*/
func (a *SupportApiService) ViewTicket(ctx context.Context, projectUuid string, memberUuid string, ticketId string) ApiViewTicketRequest {
	return ApiViewTicketRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
		memberUuid: memberUuid,
		ticketId: ticketId,
	}
}

// Execute executes the request
//  @return TicketResponse
func (a *SupportApiService) ViewTicketExecute(r ApiViewTicketRequest) (*TicketResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TicketResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupportApiService.ViewTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/support/projects/{project_uuid}/members/{member_uuid}/tickets/{ticket_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"member_uuid"+"}", url.PathEscape(parameterToString(r.memberUuid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ticket_id"+"}", url.PathEscape(parameterToString(r.ticketId, "")), -1)

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
