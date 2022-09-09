# \PlacesApi

All URIs are relative to *https://api.prod.ehelply.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvancedSearchPlaces**](PlacesApi.md#AdvancedSearchPlaces) | **Get** /places/search/places/string | Advancedsearchplaces
[**CreatePlacePlacesPlacesPost**](PlacesApi.md#CreatePlacePlacesPlacesPost) | **Post** /places/places | Create Place
[**DeletePlace**](PlacesApi.md#DeletePlace) | **Delete** /places/places/{place_uuid} | Deleteplace
[**ForwardGeocodingPlacesGeocodingForwardGet**](PlacesApi.md#ForwardGeocodingPlacesGeocodingForwardGet) | **Get** /places/geocoding/forward | Forward Geocoding
[**GetPlace**](PlacesApi.md#GetPlace) | **Get** /places/places/{place_uuid} | Getplace
[**ReverseGeocodingPlacesGeocodingReverseGet**](PlacesApi.md#ReverseGeocodingPlacesGeocodingReverseGet) | **Get** /places/geocoding/reverse | Reverse Geocoding
[**SearchPlaces**](PlacesApi.md#SearchPlaces) | **Get** /places/places | Searchplaces
[**UpdatePlace**](PlacesApi.md#UpdatePlace) | **Put** /places/places/{place_uuid} | Updateplace



## AdvancedSearchPlaces

> Page AdvancedSearchPlaces(ctx).SearchString(searchString).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Advancedsearchplaces



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchString := "searchString_example" // string |  (optional) (default to "")
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    sortOn := "sortOn_example" // string |  (optional)
    sortDesc := true // bool |  (optional) (default to false)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.AdvancedSearchPlaces(context.Background()).SearchString(searchString).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.AdvancedSearchPlaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdvancedSearchPlaces`: Page
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.AdvancedSearchPlaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdvancedSearchPlacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchString** | **string** |  | [default to &quot;&quot;]
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **sortOn** | **string** |  | 
 **sortDesc** | **bool** |  | [default to false]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**Page**](Page.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlacePlacesPlacesPost

> PlaceResponse CreatePlacePlacesPlacesPost(ctx).PlaceBase(placeBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Create Place



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    placeBase := *openapiclient.NewPlaceBase("Example Clinic") // PlaceBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.CreatePlacePlacesPlacesPost(context.Background()).PlaceBase(placeBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.CreatePlacePlacesPlacesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlacePlacesPlacesPost`: PlaceResponse
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.CreatePlacePlacesPlacesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlacePlacesPlacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placeBase** | [**PlaceBase**](PlaceBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**PlaceResponse**](PlaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlace

> interface{} DeletePlace(ctx, placeUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Deleteplace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    placeUuid := "placeUuid_example" // string | 
    softDelete := true // bool |  (optional) (default to true)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.DeletePlace(context.Background(), placeUuid).SoftDelete(softDelete).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.DeletePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePlace`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.DeletePlace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **softDelete** | **bool** |  | [default to true]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForwardGeocodingPlacesGeocodingForwardGet

> interface{} ForwardGeocodingPlacesGeocodingForwardGet(ctx).SearchingPlace(searchingPlace).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Forward Geocoding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchingPlace := "searchingPlace_example" // string | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.ForwardGeocodingPlacesGeocodingForwardGet(context.Background()).SearchingPlace(searchingPlace).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.ForwardGeocodingPlacesGeocodingForwardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForwardGeocodingPlacesGeocodingForwardGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.ForwardGeocodingPlacesGeocodingForwardGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForwardGeocodingPlacesGeocodingForwardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchingPlace** | **string** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlace

> PlaceResponse GetPlace(ctx, placeUuid).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithCollection(withCollection).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).WithCompany(withCompany).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Getplace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    placeUuid := "placeUuid_example" // string | 
    withMeta := true // bool |  (optional) (default to false)
    withCatalog := true // bool |  (optional) (default to false)
    withReviews := true // bool |  (optional) (default to false)
    withSchedule := true // bool |  (optional) (default to false)
    withCollection := true // bool |  (optional) (default to false)
    withBlog := true // bool |  (optional) (default to false)
    withTags := true // bool |  (optional) (default to false)
    withCategories := true // bool |  (optional) (default to false)
    withCompany := true // bool |  (optional) (default to false)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.GetPlace(context.Background(), placeUuid).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithCollection(withCollection).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).WithCompany(withCompany).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.GetPlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlace`: PlaceResponse
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.GetPlace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withMeta** | **bool** |  | [default to false]
 **withCatalog** | **bool** |  | [default to false]
 **withReviews** | **bool** |  | [default to false]
 **withSchedule** | **bool** |  | [default to false]
 **withCollection** | **bool** |  | [default to false]
 **withBlog** | **bool** |  | [default to false]
 **withTags** | **bool** |  | [default to false]
 **withCategories** | **bool** |  | [default to false]
 **withCompany** | **bool** |  | [default to false]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**PlaceResponse**](PlaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReverseGeocodingPlacesGeocodingReverseGet

> interface{} ReverseGeocodingPlacesGeocodingReverseGet(ctx).Long(long).Lat(lat).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Reverse Geocoding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    long := float32(8.14) // float32 | 
    lat := float32(8.14) // float32 | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.ReverseGeocodingPlacesGeocodingReverseGet(context.Background()).Long(long).Lat(lat).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.ReverseGeocodingPlacesGeocodingReverseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReverseGeocodingPlacesGeocodingReverseGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.ReverseGeocodingPlacesGeocodingReverseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReverseGeocodingPlacesGeocodingReverseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **long** | **float32** |  | 
 **lat** | **float32** |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPlaces

> Page SearchPlaces(ctx).Name(name).AddressLine1(addressLine1).AddressLine2(addressLine2).City(city).ProvinceState(provinceState).Country(country).PostalZipCode(postalZipCode).Lat(lat).Lng(lng).Email(email).IsPublic(isPublic).IsDeleted(isDeleted).WithCompany(withCompany).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithCollection(withCollection).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Searchplaces



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string |  (optional)
    addressLine1 := "addressLine1_example" // string |  (optional)
    addressLine2 := "addressLine2_example" // string |  (optional)
    city := "city_example" // string |  (optional)
    provinceState := "provinceState_example" // string |  (optional)
    country := "country_example" // string |  (optional)
    postalZipCode := "postalZipCode_example" // string |  (optional)
    lat := "lat_example" // string |  (optional)
    lng := "lng_example" // string |  (optional)
    email := "email_example" // string |  (optional)
    isPublic := true // bool |  (optional) (default to true)
    isDeleted := true // bool |  (optional) (default to false)
    withCompany := true // bool |  (optional) (default to false)
    withMeta := true // bool |  (optional) (default to false)
    withCatalog := true // bool |  (optional) (default to false)
    withReviews := true // bool |  (optional) (default to false)
    withSchedule := true // bool |  (optional) (default to false)
    withCollection := true // bool |  (optional) (default to false)
    withBlog := true // bool |  (optional) (default to false)
    withTags := true // bool |  (optional) (default to false)
    withCategories := true // bool |  (optional) (default to false)
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 25)
    sortOn := "sortOn_example" // string |  (optional)
    sortDesc := true // bool |  (optional) (default to false)
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.SearchPlaces(context.Background()).Name(name).AddressLine1(addressLine1).AddressLine2(addressLine2).City(city).ProvinceState(provinceState).Country(country).PostalZipCode(postalZipCode).Lat(lat).Lng(lng).Email(email).IsPublic(isPublic).IsDeleted(isDeleted).WithCompany(withCompany).WithMeta(withMeta).WithCatalog(withCatalog).WithReviews(withReviews).WithSchedule(withSchedule).WithCollection(withCollection).WithBlog(withBlog).WithTags(withTags).WithCategories(withCategories).Page(page).PageSize(pageSize).SortOn(sortOn).SortDesc(sortDesc).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.SearchPlaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPlaces`: Page
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.SearchPlaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPlacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **addressLine1** | **string** |  | 
 **addressLine2** | **string** |  | 
 **city** | **string** |  | 
 **provinceState** | **string** |  | 
 **country** | **string** |  | 
 **postalZipCode** | **string** |  | 
 **lat** | **string** |  | 
 **lng** | **string** |  | 
 **email** | **string** |  | 
 **isPublic** | **bool** |  | [default to true]
 **isDeleted** | **bool** |  | [default to false]
 **withCompany** | **bool** |  | [default to false]
 **withMeta** | **bool** |  | [default to false]
 **withCatalog** | **bool** |  | [default to false]
 **withReviews** | **bool** |  | [default to false]
 **withSchedule** | **bool** |  | [default to false]
 **withCollection** | **bool** |  | [default to false]
 **withBlog** | **bool** |  | [default to false]
 **withTags** | **bool** |  | [default to false]
 **withCategories** | **bool** |  | [default to false]
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **sortOn** | **string** |  | 
 **sortDesc** | **bool** |  | [default to false]
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**Page**](Page.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlace

> PlaceResponse UpdatePlace(ctx, placeUuid).PlaceBase(placeBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()

Updateplace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    placeUuid := "placeUuid_example" // string | 
    placeBase := *openapiclient.NewPlaceBase("Example Clinic") // PlaceBase | 
    xAccessToken := "xAccessToken_example" // string |  (optional)
    xSecretToken := "xSecretToken_example" // string |  (optional)
    authorization := "authorization_example" // string |  (optional)
    ehelplyActiveParticipant := "ehelplyActiveParticipant_example" // string |  (optional)
    ehelplyProject := "ehelplyProject_example" // string |  (optional)
    ehelplyData := "ehelplyData_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlacesApi.UpdatePlace(context.Background(), placeUuid).PlaceBase(placeBase).XAccessToken(xAccessToken).XSecretToken(xSecretToken).Authorization(authorization).EhelplyActiveParticipant(ehelplyActiveParticipant).EhelplyProject(ehelplyProject).EhelplyData(ehelplyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesApi.UpdatePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlace`: PlaceResponse
    fmt.Fprintf(os.Stdout, "Response from `PlacesApi.UpdatePlace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **placeBase** | [**PlaceBase**](PlaceBase.md) |  | 
 **xAccessToken** | **string** |  | 
 **xSecretToken** | **string** |  | 
 **authorization** | **string** |  | 
 **ehelplyActiveParticipant** | **string** |  | 
 **ehelplyProject** | **string** |  | 
 **ehelplyData** | **string** |  | 

### Return type

[**PlaceResponse**](PlaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

