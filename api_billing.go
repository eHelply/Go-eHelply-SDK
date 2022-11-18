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


// BillingApiService BillingApi service
type BillingApiService service

type ApiCreateBillingAccountRequest struct {
	ctx context.Context
	ApiService *BillingApiService
	projectUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateBillingAccountRequest) XAccessToken(xAccessToken string) ApiCreateBillingAccountRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateBillingAccountRequest) XSecretToken(xSecretToken string) ApiCreateBillingAccountRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateBillingAccountRequest) Authorization(authorization string) ApiCreateBillingAccountRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateBillingAccountRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateBillingAccountRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateBillingAccountRequest) EhelplyProject(ehelplyProject string) ApiCreateBillingAccountRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateBillingAccountRequest) EhelplyData(ehelplyData string) ApiCreateBillingAccountRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateBillingAccountRequest) Execute() (*StripeAccountResponse, *http.Response, error) {
	return r.ApiService.CreateBillingAccountExecute(r)
}

/*
CreateBillingAccount Createbillingaccount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @return ApiCreateBillingAccountRequest
*/
func (a *BillingApiService) CreateBillingAccount(ctx context.Context, projectUuid string) ApiCreateBillingAccountRequest {
	return ApiCreateBillingAccountRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
	}
}

// Execute executes the request
//  @return StripeAccountResponse
func (a *BillingApiService) CreateBillingAccountExecute(r ApiCreateBillingAccountRequest) (*StripeAccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StripeAccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingApiService.CreateBillingAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/billing/projects/{project_uuid}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)

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

type ApiGetClientSecretRequest struct {
	ctx context.Context
	ApiService *BillingApiService
	projectUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetClientSecretRequest) XAccessToken(xAccessToken string) ApiGetClientSecretRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetClientSecretRequest) XSecretToken(xSecretToken string) ApiGetClientSecretRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetClientSecretRequest) Authorization(authorization string) ApiGetClientSecretRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetClientSecretRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetClientSecretRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetClientSecretRequest) EhelplyProject(ehelplyProject string) ApiGetClientSecretRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetClientSecretRequest) EhelplyData(ehelplyData string) ApiGetClientSecretRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetClientSecretRequest) Execute() (*StripeCustomerSecretResponse, *http.Response, error) {
	return r.ApiService.GetClientSecretExecute(r)
}

/*
GetClientSecret Getclientsecret

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @return ApiGetClientSecretRequest
*/
func (a *BillingApiService) GetClientSecret(ctx context.Context, projectUuid string) ApiGetClientSecretRequest {
	return ApiGetClientSecretRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
	}
}

// Execute executes the request
//  @return StripeCustomerSecretResponse
func (a *BillingApiService) GetClientSecretExecute(r ApiGetClientSecretRequest) (*StripeCustomerSecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StripeCustomerSecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingApiService.GetClientSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/billing/projects/{project_uuid}/secrets"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)

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

type ApiHasPaymentRequest struct {
	ctx context.Context
	ApiService *BillingApiService
	projectUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiHasPaymentRequest) XAccessToken(xAccessToken string) ApiHasPaymentRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiHasPaymentRequest) XSecretToken(xSecretToken string) ApiHasPaymentRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiHasPaymentRequest) Authorization(authorization string) ApiHasPaymentRequest {
	r.authorization = &authorization
	return r
}

func (r ApiHasPaymentRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiHasPaymentRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiHasPaymentRequest) EhelplyProject(ehelplyProject string) ApiHasPaymentRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiHasPaymentRequest) EhelplyData(ehelplyData string) ApiHasPaymentRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiHasPaymentRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.HasPaymentExecute(r)
}

/*
HasPayment Haspayment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @return ApiHasPaymentRequest
*/
func (a *BillingApiService) HasPayment(ctx context.Context, projectUuid string) ApiHasPaymentRequest {
	return ApiHasPaymentRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
	}
}

// Execute executes the request
//  @return bool
func (a *BillingApiService) HasPaymentExecute(r ApiHasPaymentRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingApiService.HasPayment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/billing/projects/{project_uuid}/payment-methods-exist"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)

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

type ApiListPaymentMethodsRequest struct {
	ctx context.Context
	ApiService *BillingApiService
	projectUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiListPaymentMethodsRequest) XAccessToken(xAccessToken string) ApiListPaymentMethodsRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiListPaymentMethodsRequest) XSecretToken(xSecretToken string) ApiListPaymentMethodsRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiListPaymentMethodsRequest) Authorization(authorization string) ApiListPaymentMethodsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiListPaymentMethodsRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiListPaymentMethodsRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiListPaymentMethodsRequest) EhelplyProject(ehelplyProject string) ApiListPaymentMethodsRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiListPaymentMethodsRequest) EhelplyData(ehelplyData string) ApiListPaymentMethodsRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiListPaymentMethodsRequest) Execute() ([]PaymentMethodResponse, *http.Response, error) {
	return r.ApiService.ListPaymentMethodsExecute(r)
}

/*
ListPaymentMethods Listpaymentmethods

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @return ApiListPaymentMethodsRequest
*/
func (a *BillingApiService) ListPaymentMethods(ctx context.Context, projectUuid string) ApiListPaymentMethodsRequest {
	return ApiListPaymentMethodsRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
	}
}

// Execute executes the request
//  @return []PaymentMethodResponse
func (a *BillingApiService) ListPaymentMethodsExecute(r ApiListPaymentMethodsRequest) ([]PaymentMethodResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PaymentMethodResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingApiService.ListPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/billing/projects/{project_uuid}/payment-methods"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)

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

type ApiProcessPaymentRequest struct {
	ctx context.Context
	ApiService *BillingApiService
	projectUuid string
	payment *Payment
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiProcessPaymentRequest) Payment(payment Payment) ApiProcessPaymentRequest {
	r.payment = &payment
	return r
}

func (r ApiProcessPaymentRequest) XAccessToken(xAccessToken string) ApiProcessPaymentRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiProcessPaymentRequest) XSecretToken(xSecretToken string) ApiProcessPaymentRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiProcessPaymentRequest) Authorization(authorization string) ApiProcessPaymentRequest {
	r.authorization = &authorization
	return r
}

func (r ApiProcessPaymentRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiProcessPaymentRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiProcessPaymentRequest) EhelplyProject(ehelplyProject string) ApiProcessPaymentRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiProcessPaymentRequest) EhelplyData(ehelplyData string) ApiProcessPaymentRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiProcessPaymentRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ProcessPaymentExecute(r)
}

/*
ProcessPayment Processpayment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @return ApiProcessPaymentRequest
*/
func (a *BillingApiService) ProcessPayment(ctx context.Context, projectUuid string) ApiProcessPaymentRequest {
	return ApiProcessPaymentRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
	}
}

// Execute executes the request
//  @return string
func (a *BillingApiService) ProcessPaymentExecute(r ApiProcessPaymentRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingApiService.ProcessPayment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/billing/projects/{project_uuid}/payments"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payment == nil {
		return localVarReturnValue, nil, reportError("payment is required and must be specified")
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
	localVarPostBody = r.payment
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

type ApiReconcilePaymentMethodRequest struct {
	ctx context.Context
	ApiService *BillingApiService
	projectUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiReconcilePaymentMethodRequest) XAccessToken(xAccessToken string) ApiReconcilePaymentMethodRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiReconcilePaymentMethodRequest) XSecretToken(xSecretToken string) ApiReconcilePaymentMethodRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiReconcilePaymentMethodRequest) Authorization(authorization string) ApiReconcilePaymentMethodRequest {
	r.authorization = &authorization
	return r
}

func (r ApiReconcilePaymentMethodRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiReconcilePaymentMethodRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiReconcilePaymentMethodRequest) EhelplyProject(ehelplyProject string) ApiReconcilePaymentMethodRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiReconcilePaymentMethodRequest) EhelplyData(ehelplyData string) ApiReconcilePaymentMethodRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiReconcilePaymentMethodRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.ReconcilePaymentMethodExecute(r)
}

/*
ReconcilePaymentMethod Reconcilepaymentmethod

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @return ApiReconcilePaymentMethodRequest
*/
func (a *BillingApiService) ReconcilePaymentMethod(ctx context.Context, projectUuid string) ApiReconcilePaymentMethodRequest {
	return ApiReconcilePaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
	}
}

// Execute executes the request
//  @return bool
func (a *BillingApiService) ReconcilePaymentMethodExecute(r ApiReconcilePaymentMethodRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingApiService.ReconcilePaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/billing/projects/{project_uuid}/payment-methods-reconciliation"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)

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

type ApiRemovePaymentMethodRequest struct {
	ctx context.Context
	ApiService *BillingApiService
	projectUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiRemovePaymentMethodRequest) XAccessToken(xAccessToken string) ApiRemovePaymentMethodRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiRemovePaymentMethodRequest) XSecretToken(xSecretToken string) ApiRemovePaymentMethodRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiRemovePaymentMethodRequest) Authorization(authorization string) ApiRemovePaymentMethodRequest {
	r.authorization = &authorization
	return r
}

func (r ApiRemovePaymentMethodRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiRemovePaymentMethodRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiRemovePaymentMethodRequest) EhelplyProject(ehelplyProject string) ApiRemovePaymentMethodRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiRemovePaymentMethodRequest) EhelplyData(ehelplyData string) ApiRemovePaymentMethodRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiRemovePaymentMethodRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.RemovePaymentMethodExecute(r)
}

/*
RemovePaymentMethod Removepaymentmethod

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectUuid
 @return ApiRemovePaymentMethodRequest
*/
func (a *BillingApiService) RemovePaymentMethod(ctx context.Context, projectUuid string) ApiRemovePaymentMethodRequest {
	return ApiRemovePaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
		projectUuid: projectUuid,
	}
}

// Execute executes the request
//  @return string
func (a *BillingApiService) RemovePaymentMethodExecute(r ApiRemovePaymentMethodRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillingApiService.RemovePaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/billing/projects/{project_uuid}/payment-methods"
	localVarPath = strings.Replace(localVarPath, "{"+"project_uuid"+"}", url.PathEscape(parameterToString(r.projectUuid, "")), -1)

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
