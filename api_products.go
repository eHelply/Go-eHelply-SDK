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
	"reflect"
)


// ProductsApiService ProductsApi service
type ProductsApiService service

type ApiCreateProductRequest struct {
	ctx context.Context
	ApiService *ProductsApiService
	productBase *ProductBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiCreateProductRequest) ProductBase(productBase ProductBase) ApiCreateProductRequest {
	r.productBase = &productBase
	return r
}

func (r ApiCreateProductRequest) XAccessToken(xAccessToken string) ApiCreateProductRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiCreateProductRequest) XSecretToken(xSecretToken string) ApiCreateProductRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiCreateProductRequest) Authorization(authorization string) ApiCreateProductRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateProductRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiCreateProductRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiCreateProductRequest) EhelplyProject(ehelplyProject string) ApiCreateProductRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiCreateProductRequest) EhelplyData(ehelplyData string) ApiCreateProductRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiCreateProductRequest) Execute() (*ProductReturn, *http.Response, error) {
	return r.ApiService.CreateProductExecute(r)
}

/*
CreateProduct Createproduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateProductRequest
*/
func (a *ProductsApiService) CreateProduct(ctx context.Context) ApiCreateProductRequest {
	return ApiCreateProductRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductReturn
func (a *ProductsApiService) CreateProductExecute(r ApiCreateProductRequest) (*ProductReturn, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductReturn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsApiService.CreateProduct")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productBase == nil {
		return localVarReturnValue, nil, reportError("productBase is required and must be specified")
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
	localVarPostBody = r.productBase
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

type ApiDeleteProductRequest struct {
	ctx context.Context
	ApiService *ProductsApiService
	productUuid string
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiDeleteProductRequest) XAccessToken(xAccessToken string) ApiDeleteProductRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiDeleteProductRequest) XSecretToken(xSecretToken string) ApiDeleteProductRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiDeleteProductRequest) Authorization(authorization string) ApiDeleteProductRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteProductRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiDeleteProductRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiDeleteProductRequest) EhelplyProject(ehelplyProject string) ApiDeleteProductRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiDeleteProductRequest) EhelplyData(ehelplyData string) ApiDeleteProductRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiDeleteProductRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.DeleteProductExecute(r)
}

/*
DeleteProduct Deleteproduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productUuid
 @return ApiDeleteProductRequest
*/
func (a *ProductsApiService) DeleteProduct(ctx context.Context, productUuid string) ApiDeleteProductRequest {
	return ApiDeleteProductRequest{
		ApiService: a,
		ctx: ctx,
		productUuid: productUuid,
	}
}

// Execute executes the request
//  @return bool
func (a *ProductsApiService) DeleteProductExecute(r ApiDeleteProductRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsApiService.DeleteProduct")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/products/{product_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"product_uuid"+"}", url.PathEscape(parameterToString(r.productUuid, "")), -1)

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

type ApiGetProductRequest struct {
	ctx context.Context
	ApiService *ProductsApiService
	productUuid string
	withAddons *bool
	withMeta *bool
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiGetProductRequest) WithAddons(withAddons bool) ApiGetProductRequest {
	r.withAddons = &withAddons
	return r
}

func (r ApiGetProductRequest) WithMeta(withMeta bool) ApiGetProductRequest {
	r.withMeta = &withMeta
	return r
}

func (r ApiGetProductRequest) XAccessToken(xAccessToken string) ApiGetProductRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiGetProductRequest) XSecretToken(xSecretToken string) ApiGetProductRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiGetProductRequest) Authorization(authorization string) ApiGetProductRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetProductRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiGetProductRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiGetProductRequest) EhelplyProject(ehelplyProject string) ApiGetProductRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiGetProductRequest) EhelplyData(ehelplyData string) ApiGetProductRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiGetProductRequest) Execute() (*ProductReturn, *http.Response, error) {
	return r.ApiService.GetProductExecute(r)
}

/*
GetProduct Getproduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productUuid
 @return ApiGetProductRequest
*/
func (a *ProductsApiService) GetProduct(ctx context.Context, productUuid string) ApiGetProductRequest {
	return ApiGetProductRequest{
		ApiService: a,
		ctx: ctx,
		productUuid: productUuid,
	}
}

// Execute executes the request
//  @return ProductReturn
func (a *ProductsApiService) GetProductExecute(r ApiGetProductRequest) (*ProductReturn, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductReturn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsApiService.GetProduct")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/products/{product_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"product_uuid"+"}", url.PathEscape(parameterToString(r.productUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withAddons != nil {
		localVarQueryParams.Add("with_addons", parameterToString(*r.withAddons, ""))
	}
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

type ApiSearchProductCatalogRequest struct {
	ctx context.Context
	ApiService *ProductsApiService
	productUuid string
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

func (r ApiSearchProductCatalogRequest) WithMeta(withMeta bool) ApiSearchProductCatalogRequest {
	r.withMeta = &withMeta
	return r
}

func (r ApiSearchProductCatalogRequest) Page(page int32) ApiSearchProductCatalogRequest {
	r.page = &page
	return r
}

func (r ApiSearchProductCatalogRequest) PageSize(pageSize int32) ApiSearchProductCatalogRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchProductCatalogRequest) SortOn(sortOn string) ApiSearchProductCatalogRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchProductCatalogRequest) SortDesc(sortDesc bool) ApiSearchProductCatalogRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchProductCatalogRequest) XAccessToken(xAccessToken string) ApiSearchProductCatalogRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchProductCatalogRequest) XSecretToken(xSecretToken string) ApiSearchProductCatalogRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchProductCatalogRequest) Authorization(authorization string) ApiSearchProductCatalogRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchProductCatalogRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchProductCatalogRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchProductCatalogRequest) EhelplyProject(ehelplyProject string) ApiSearchProductCatalogRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchProductCatalogRequest) EhelplyData(ehelplyData string) ApiSearchProductCatalogRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchProductCatalogRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchProductCatalogExecute(r)
}

/*
SearchProductCatalog Searchproductcatalog

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productUuid
 @return ApiSearchProductCatalogRequest
*/
func (a *ProductsApiService) SearchProductCatalog(ctx context.Context, productUuid string) ApiSearchProductCatalogRequest {
	return ApiSearchProductCatalogRequest{
		ApiService: a,
		ctx: ctx,
		productUuid: productUuid,
	}
}

// Execute executes the request
//  @return Page
func (a *ProductsApiService) SearchProductCatalogExecute(r ApiSearchProductCatalogRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsApiService.SearchProductCatalog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/products/{product_uuid}/catalogs"
	localVarPath = strings.Replace(localVarPath, "{"+"product_uuid"+"}", url.PathEscape(parameterToString(r.productUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiSearchProductsRequest struct {
	ctx context.Context
	ApiService *ProductsApiService
	withMeta *bool
	name *string
	addons *[]string
	priceMax *int32
	priceMin *int32
	quantityAvailable *bool
	isDeleted *bool
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

func (r ApiSearchProductsRequest) WithMeta(withMeta bool) ApiSearchProductsRequest {
	r.withMeta = &withMeta
	return r
}

func (r ApiSearchProductsRequest) Name(name string) ApiSearchProductsRequest {
	r.name = &name
	return r
}

func (r ApiSearchProductsRequest) Addons(addons []string) ApiSearchProductsRequest {
	r.addons = &addons
	return r
}

func (r ApiSearchProductsRequest) PriceMax(priceMax int32) ApiSearchProductsRequest {
	r.priceMax = &priceMax
	return r
}

func (r ApiSearchProductsRequest) PriceMin(priceMin int32) ApiSearchProductsRequest {
	r.priceMin = &priceMin
	return r
}

func (r ApiSearchProductsRequest) QuantityAvailable(quantityAvailable bool) ApiSearchProductsRequest {
	r.quantityAvailable = &quantityAvailable
	return r
}

func (r ApiSearchProductsRequest) IsDeleted(isDeleted bool) ApiSearchProductsRequest {
	r.isDeleted = &isDeleted
	return r
}

func (r ApiSearchProductsRequest) Page(page int32) ApiSearchProductsRequest {
	r.page = &page
	return r
}

func (r ApiSearchProductsRequest) PageSize(pageSize int32) ApiSearchProductsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSearchProductsRequest) SortOn(sortOn string) ApiSearchProductsRequest {
	r.sortOn = &sortOn
	return r
}

func (r ApiSearchProductsRequest) SortDesc(sortDesc bool) ApiSearchProductsRequest {
	r.sortDesc = &sortDesc
	return r
}

func (r ApiSearchProductsRequest) XAccessToken(xAccessToken string) ApiSearchProductsRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiSearchProductsRequest) XSecretToken(xSecretToken string) ApiSearchProductsRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiSearchProductsRequest) Authorization(authorization string) ApiSearchProductsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSearchProductsRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiSearchProductsRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiSearchProductsRequest) EhelplyProject(ehelplyProject string) ApiSearchProductsRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiSearchProductsRequest) EhelplyData(ehelplyData string) ApiSearchProductsRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiSearchProductsRequest) Execute() (*Page, *http.Response, error) {
	return r.ApiService.SearchProductsExecute(r)
}

/*
SearchProducts Searchproducts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchProductsRequest
*/
func (a *ProductsApiService) SearchProducts(ctx context.Context) ApiSearchProductsRequest {
	return ApiSearchProductsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Page
func (a *ProductsApiService) SearchProductsExecute(r ApiSearchProductsRequest) (*Page, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Page
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsApiService.SearchProducts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withMeta != nil {
		localVarQueryParams.Add("with_meta", parameterToString(*r.withMeta, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.addons != nil {
		t := *r.addons
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("addons", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("addons", parameterToString(t, "multi"))
		}
	}
	if r.priceMax != nil {
		localVarQueryParams.Add("price_max", parameterToString(*r.priceMax, ""))
	}
	if r.priceMin != nil {
		localVarQueryParams.Add("price_min", parameterToString(*r.priceMin, ""))
	}
	if r.quantityAvailable != nil {
		localVarQueryParams.Add("quantity_available", parameterToString(*r.quantityAvailable, ""))
	}
	if r.isDeleted != nil {
		localVarQueryParams.Add("is_deleted", parameterToString(*r.isDeleted, ""))
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

type ApiUpdateProductRequest struct {
	ctx context.Context
	ApiService *ProductsApiService
	productUuid string
	productBase *ProductBase
	xAccessToken *string
	xSecretToken *string
	authorization *string
	ehelplyActiveParticipant *string
	ehelplyProject *string
	ehelplyData *string
}

func (r ApiUpdateProductRequest) ProductBase(productBase ProductBase) ApiUpdateProductRequest {
	r.productBase = &productBase
	return r
}

func (r ApiUpdateProductRequest) XAccessToken(xAccessToken string) ApiUpdateProductRequest {
	r.xAccessToken = &xAccessToken
	return r
}

func (r ApiUpdateProductRequest) XSecretToken(xSecretToken string) ApiUpdateProductRequest {
	r.xSecretToken = &xSecretToken
	return r
}

func (r ApiUpdateProductRequest) Authorization(authorization string) ApiUpdateProductRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateProductRequest) EhelplyActiveParticipant(ehelplyActiveParticipant string) ApiUpdateProductRequest {
	r.ehelplyActiveParticipant = &ehelplyActiveParticipant
	return r
}

func (r ApiUpdateProductRequest) EhelplyProject(ehelplyProject string) ApiUpdateProductRequest {
	r.ehelplyProject = &ehelplyProject
	return r
}

func (r ApiUpdateProductRequest) EhelplyData(ehelplyData string) ApiUpdateProductRequest {
	r.ehelplyData = &ehelplyData
	return r
}

func (r ApiUpdateProductRequest) Execute() (*ProductReturn, *http.Response, error) {
	return r.ApiService.UpdateProductExecute(r)
}

/*
UpdateProduct Updateproduct

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productUuid
 @return ApiUpdateProductRequest
*/
func (a *ProductsApiService) UpdateProduct(ctx context.Context, productUuid string) ApiUpdateProductRequest {
	return ApiUpdateProductRequest{
		ApiService: a,
		ctx: ctx,
		productUuid: productUuid,
	}
}

// Execute executes the request
//  @return ProductReturn
func (a *ProductsApiService) UpdateProductExecute(r ApiUpdateProductRequest) (*ProductReturn, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductReturn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsApiService.UpdateProduct")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/products/{product_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"product_uuid"+"}", url.PathEscape(parameterToString(r.productUuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productBase == nil {
		return localVarReturnValue, nil, reportError("productBase is required and must be specified")
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
	localVarPostBody = r.productBase
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
