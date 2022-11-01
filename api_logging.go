/*
eHelply SDK - 1.1.117

eHelply SDK for SuperStack Services

API version: 1.1.117
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


// LoggingApiService LoggingApi service
type LoggingApiService service

type ApiGetSubjectLogsRequest struct {
	ctx context.Context
	ApiService *LoggingApiService
	service string
	subject string
	dateStart *string
	dateEnd *string
	desc *bool
	limit *int32
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetSubjectLogsRequest) DateStart(dateStart string) ApiGetSubjectLogsRequest {
	r.dateStart = &dateStart
	return r
}

func (r ApiGetSubjectLogsRequest) DateEnd(dateEnd string) ApiGetSubjectLogsRequest {
	r.dateEnd = &dateEnd
	return r
}

func (r ApiGetSubjectLogsRequest) Desc(desc bool) ApiGetSubjectLogsRequest {
	r.desc = &desc
	return r
}

func (r ApiGetSubjectLogsRequest) Limit(limit int32) ApiGetSubjectLogsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetSubjectLogsRequest) XAccessToken(xAccessToken string) ApiGetSubjectLogsRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetSubjectLogsRequest) XSecretToken(xSecretToken string) ApiGetSubjectLogsRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetSubjectLogsRequest) Authorization(authorization string) ApiGetSubjectLogsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetSubjectLogsRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetSubjectLogsRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetSubjectLogsRequest) EhelplyProject(ehelplyProject string) ApiGetSubjectLogsRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetSubjectLogsRequest) EhelplyData(ehelplyData string) ApiGetSubjectLogsRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetSubjectLogsRequest) Execute() ([]LoggingDynamo, *http.Response, error) {
	return r.ApiService.GetSubjectLogsExecute(r)
}

/*
GetSubjectLogs Getsubjectlogs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param service
 @param subject
 @return ApiGetSubjectLogsRequest
*/
func (a *LoggingApiService) GetSubjectLogs(ctx context.Context, service string, subject string) ApiGetSubjectLogsRequest {
	return ApiGetSubjectLogsRequest{
		ApiService: a,
		ctx: ctx,
		service: service,
		subject: subject,
	}
}

// Execute executes the request
//  @return []LoggingDynamo
func (a *LoggingApiService) GetSubjectLogsExecute(r ApiGetSubjectLogsRequest) ([]LoggingDynamo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LoggingDynamo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LoggingApiService.GetSubjectLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sam/logging/logs/services/{service}/subjects/{subject}"
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", url.PathEscape(parameterToString(r.service, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subject"+"}", url.PathEscape(parameterToString(r.subject, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.dateStart != nil {
		localVarQueryParams.Add("date_start", parameterToString(*r.dateStart, ""))
	}
	if r.dateEnd != nil {
		localVarQueryParams.Add("date_end", parameterToString(*r.dateEnd, ""))
	}
	if r.desc != nil {
		localVarQueryParams.Add("desc", parameterToString(*r.desc, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
