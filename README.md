# Go API client for ehelply

eHelply SDK for SuperStack Services

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.116
- Package version: 1.1.116
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://superstack.ehelply.com/support](https://superstack.ehelply.com/support)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import ehelply "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), ehelply.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), ehelply.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), ehelply.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), ehelply.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.prod.ehelply.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppointmentsApi* | [**AddEntityToAppointment**](docs/AppointmentsApi.md#addentitytoappointment) | **Post** /appointments/appointments/{appointment_uuid}/entities/{entity_uuid} | Addentitytoappointment
*AppointmentsApi* | [**CreateAppointment**](docs/AppointmentsApi.md#createappointment) | **Post** /appointments/appointments | Createappointment
*AppointmentsApi* | [**DeleteAppointment**](docs/AppointmentsApi.md#deleteappointment) | **Delete** /appointments/appointments/{appointment_uuid} | Deleteappointment
*AppointmentsApi* | [**DetachEntityFromAppointment**](docs/AppointmentsApi.md#detachentityfromappointment) | **Delete** /appointments/appointments/{appointment_uuid}/entities/{entity_uuid} | Removeentityfromappointment
*AppointmentsApi* | [**GetAppointment**](docs/AppointmentsApi.md#getappointment) | **Get** /appointments/appointments/{appointment_uuid} | Getappointment
*AppointmentsApi* | [**SearchAppointment**](docs/AppointmentsApi.md#searchappointment) | **Get** /appointments/appointments | Searchappointments
*AppointmentsApi* | [**SearchAppointmentEntities**](docs/AppointmentsApi.md#searchappointmententities) | **Get** /appointments/appointments/{appointment_uuid}/entities | Searchappointmententities
*AppointmentsApi* | [**SearchEntityAppointments**](docs/AppointmentsApi.md#searchentityappointments) | **Get** /appointments/appointments/entities/{entity_uuid}/appointments | Getentityappointments
*AppointmentsApi* | [**UpdateAppointment**](docs/AppointmentsApi.md#updateappointment) | **Put** /appointments/appointments/{appointment_uuid} | Updateappointment
*BillingApi* | [**CreateBillingAccount**](docs/BillingApi.md#createbillingaccount) | **Post** /sam/billing/create_billing_account | Createbillingaccount
*BillingApi* | [**GetClientSecret**](docs/BillingApi.md#getclientsecret) | **Get** /sam/billing/retrieve_secret | Getclientsecret
*BillingApi* | [**HasPayment**](docs/BillingApi.md#haspayment) | **Get** /sam/billing/has_payment | Haspayment
*BillingApi* | [**ListPaymentMethods**](docs/BillingApi.md#listpaymentmethods) | **Get** /sam/billing/view_payment_method | Listpaymentmethods
*BillingApi* | [**ProcessPayment**](docs/BillingApi.md#processpayment) | **Post** /sam/billing/process_payment | Processpayment
*BillingApi* | [**ReconcilePaymentMethod**](docs/BillingApi.md#reconcilepaymentmethod) | **Get** /sam/billing/reconcile_payment | Reconcilepaymentmethod
*BillingApi* | [**RemovePaymentMethod**](docs/BillingApi.md#removepaymentmethod) | **Delete** /sam/billing/remove_payment_method | Removepaymentmethod
*CatalogsApi* | [**AttachProductToCatalog**](docs/CatalogsApi.md#attachproducttocatalog) | **Post** /products/catalogs/{catalog_uuid}/products/{product_uuid} | Addproducttocatalog
*CatalogsApi* | [**CreateCatalog**](docs/CatalogsApi.md#createcatalog) | **Post** /products/catalogs | Createcatalog
*CatalogsApi* | [**DeleteCatalog**](docs/CatalogsApi.md#deletecatalog) | **Delete** /products/catalogs/{catalog_uuid} | Deletecatalog
*CatalogsApi* | [**DetachProductFromCatalog**](docs/CatalogsApi.md#detachproductfromcatalog) | **Delete** /products/catalogs/{catalog_uuid}/products/{product_uuid} | Removeproductfromcatalog
*CatalogsApi* | [**GetCatalog**](docs/CatalogsApi.md#getcatalog) | **Get** /products/catalogs/{catalog_uuid} | Getcatalog
*CatalogsApi* | [**SearchCatalogProducts**](docs/CatalogsApi.md#searchcatalogproducts) | **Get** /products/catalogs/{catalog_uuid}/products | Searchcatalogproducts
*CatalogsApi* | [**SearchCatalogs**](docs/CatalogsApi.md#searchcatalogs) | **Get** /products/catalogs | Searchcatalogs
*CatalogsApi* | [**UpdateCatalog**](docs/CatalogsApi.md#updatecatalog) | **Put** /products/catalogs/{catalog_uuid} | Updatecatalog
*CategoryApi* | [**CreateCategoryPlacesCategoriesPost**](docs/CategoryApi.md#createcategoryplacescategoriespost) | **Post** /places/categories | Create Category
*CategoryApi* | [**DeleteCategoryPlacesCategoriesCategoryUuidDelete**](docs/CategoryApi.md#deletecategoryplacescategoriescategoryuuiddelete) | **Delete** /places/categories/{category_uuid} | Delete Category
*CategoryApi* | [**GetCategoryPlacesCategoriesCategoryUuidGet**](docs/CategoryApi.md#getcategoryplacescategoriescategoryuuidget) | **Get** /places/categories/{category_uuid} | Get Category
*CategoryApi* | [**SearchCategoriesPlacesCategoriesGet**](docs/CategoryApi.md#searchcategoriesplacescategoriesget) | **Get** /places/categories | Search Categories
*CategoryApi* | [**UpdateCategoryPlacesCategoriesCategoryUuidPut**](docs/CategoryApi.md#updatecategoryplacescategoriescategoryuuidput) | **Put** /places/categories/{category_uuid} | Update Category
*CompaniesApi* | [**CreateCompanyPlacesCompaniesPost**](docs/CompaniesApi.md#createcompanyplacescompaniespost) | **Post** /places/companies | Create Company
*CompaniesApi* | [**DeleteCompanyPlacesCompaniesCompanyUuidDelete**](docs/CompaniesApi.md#deletecompanyplacescompaniescompanyuuiddelete) | **Delete** /places/companies/{company_uuid} | Delete Company
*CompaniesApi* | [**GetCompanyPlacesCompaniesCompanyUuidGet**](docs/CompaniesApi.md#getcompanyplacescompaniescompanyuuidget) | **Get** /places/companies/{company_uuid} | Get Company
*CompaniesApi* | [**SearchCompaniesPlacesCompaniesGet**](docs/CompaniesApi.md#searchcompaniesplacescompaniesget) | **Get** /places/companies | Search Companies
*CompaniesApi* | [**UpdateCompanyPlacesCompaniesCompanyUuidPut**](docs/CompaniesApi.md#updatecompanyplacescompaniescompanyuuidput) | **Put** /places/companies/{company_uuid} | Update Company
*ContentApi* | [**CreateFile**](docs/ContentApi.md#createfile) | **Post** /files/files | Createfile
*ContentApi* | [**DeleteFile**](docs/ContentApi.md#deletefile) | **Delete** /files/files/{file_uuid} | Deletefile
*ContentApi* | [**GetFile**](docs/ContentApi.md#getfile) | **Get** /files/files/{file_uuid} | Getfile
*ContentApi* | [**UpdateFile**](docs/ContentApi.md#updatefile) | **Put** /files/files/{file_uuid} | Updatefile
*LoggingApi* | [**GetSubjectLogs**](docs/LoggingApi.md#getsubjectlogs) | **Get** /sam/logging/logs/services/{service}/subjects/{subject} | Getsubjectlogs
*MetaApi* | [**CreateField**](docs/MetaApi.md#createfield) | **Post** /meta/field | Create Field
*MetaApi* | [**CreateMeta**](docs/MetaApi.md#createmeta) | **Post** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Create Meta
*MetaApi* | [**DeleteField**](docs/MetaApi.md#deletefield) | **Delete** /meta/field/{field_uuid} | Delete Field
*MetaApi* | [**DeleteMeta**](docs/MetaApi.md#deletemeta) | **Delete** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Delete Meta
*MetaApi* | [**DeleteMetaFromUuid**](docs/MetaApi.md#deletemetafromuuid) | **Delete** /meta/meta/{meta_uuid} | Delete Meta From Uuid
*MetaApi* | [**GetField**](docs/MetaApi.md#getfield) | **Get** /meta/field/{field_uuid} | Get Field
*MetaApi* | [**GetMeta**](docs/MetaApi.md#getmeta) | **Get** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Get Meta
*MetaApi* | [**GetMetaFromUuid**](docs/MetaApi.md#getmetafromuuid) | **Get** /meta/meta/{meta_uuid} | Get Meta From Uuid
*MetaApi* | [**MakeSlug**](docs/MetaApi.md#makeslug) | **Post** /meta/meta/slug | Make Slug
*MetaApi* | [**TouchMeta**](docs/MetaApi.md#touchmeta) | **Post** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid}/touch | Touch Meta
*MetaApi* | [**UpdateField**](docs/MetaApi.md#updatefield) | **Put** /meta/field/{field_uuid} | Update Field
*MetaApi* | [**UpdateMeta**](docs/MetaApi.md#updatemeta) | **Put** /meta/meta/service/{service}/type/{type_str}/entity/{entity_uuid} | Update Meta
*MetaApi* | [**UpdateMetaFromUuid**](docs/MetaApi.md#updatemetafromuuid) | **Put** /meta/meta/{meta_uuid} | Update Meta From Uuid
*MonitorApi* | [**AcknowledgeAlarm**](docs/MonitorApi.md#acknowledgealarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/acknowledge | Acknowledgealarm
*MonitorApi* | [**AssignAlarm**](docs/MonitorApi.md#assignalarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/assign | Assignalarm
*MonitorApi* | [**AttachAlarmNote**](docs/MonitorApi.md#attachalarmnote) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/note | Attachalarmnote
*MonitorApi* | [**AttachAlarmTicket**](docs/MonitorApi.md#attachalarmticket) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/ticket | Attachalarmticket
*MonitorApi* | [**ClearAlarm**](docs/MonitorApi.md#clearalarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/clear | Clearalarm
*MonitorApi* | [**GetService**](docs/MonitorApi.md#getservice) | **Get** /sam/monitor/services/{service} | Getservice
*MonitorApi* | [**GetServiceAlarm**](docs/MonitorApi.md#getservicealarm) | **Get** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid} | Getservicealarm
*MonitorApi* | [**GetServiceAlarms**](docs/MonitorApi.md#getservicealarms) | **Get** /sam/monitor/services/{service}/stages/{stage}/alarms | Getservicealarms
*MonitorApi* | [**GetServiceHeartbeat**](docs/MonitorApi.md#getserviceheartbeat) | **Get** /sam/monitor/services/{service}/stages/{stage}/heartbeats | Getserviceheartbeat
*MonitorApi* | [**GetServiceKpis**](docs/MonitorApi.md#getservicekpis) | **Get** /sam/monitor/services/{service}/kpis | Getservicekpis
*MonitorApi* | [**GetServiceSpec**](docs/MonitorApi.md#getservicespec) | **Get** /sam/monitor/services/{service}/specs/{spec} | Getservicespec
*MonitorApi* | [**GetServiceSpecs**](docs/MonitorApi.md#getservicespecs) | **Get** /sam/monitor/services/{service}/specs | Getservicespecs
*MonitorApi* | [**GetServiceVitals**](docs/MonitorApi.md#getservicevitals) | **Get** /sam/monitor/services/{service}/stages/{stage}/vitals | Getservicevitals
*MonitorApi* | [**GetServices**](docs/MonitorApi.md#getservices) | **Get** /sam/monitor/services | Getservices
*MonitorApi* | [**GetServicesWithSpecs**](docs/MonitorApi.md#getserviceswithspecs) | **Get** /sam/monitor/specs/services | Getserviceswithspecs
*MonitorApi* | [**HideService**](docs/MonitorApi.md#hideservice) | **Post** /sam/monitor/services/{service}/stages/{stage}/hide | Hideservice
*MonitorApi* | [**IgnoreAlarm**](docs/MonitorApi.md#ignorealarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/ignore | Ignorealarm
*MonitorApi* | [**RegisterService**](docs/MonitorApi.md#registerservice) | **Post** /sam/monitor/services | Registerservice
*MonitorApi* | [**SearchAlarms**](docs/MonitorApi.md#searchalarms) | **Get** /sam/monitor/services/{service}/alarms | Searchalarms
*MonitorApi* | [**ShowService**](docs/MonitorApi.md#showservice) | **Post** /sam/monitor/services/{service}/stages/{stage}/show | Showservice
*MonitorApi* | [**TerminateAlarm**](docs/MonitorApi.md#terminatealarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms/{alarm_uuid}/terminate | Terminatealarm
*MonitorApi* | [**TriggerAlarm**](docs/MonitorApi.md#triggeralarm) | **Post** /sam/monitor/services/{service}/stages/{stage}/alarms | Triggeralarm
*NotesApi* | [**CreateNote**](docs/NotesApi.md#createnote) | **Post** /notes/notes | Create Note
*NotesApi* | [**DeleteNote**](docs/NotesApi.md#deletenote) | **Delete** /notes/notes/{note_id} | Delete Note
*NotesApi* | [**GetNote**](docs/NotesApi.md#getnote) | **Get** /notes/notes/{note_id} | Get Note
*NotesApi* | [**UpdateNote**](docs/NotesApi.md#updatenote) | **Put** /notes/notes/{note_id} | Update Note
*PlacesApi* | [**AdvancedSearchPlaces**](docs/PlacesApi.md#advancedsearchplaces) | **Get** /places/search/places/string | Advancedsearchplaces
*PlacesApi* | [**CreatePlacePlacesPlacesPost**](docs/PlacesApi.md#createplaceplacesplacespost) | **Post** /places/places | Create Place
*PlacesApi* | [**DeletePlace**](docs/PlacesApi.md#deleteplace) | **Delete** /places/places/{place_uuid} | Deleteplace
*PlacesApi* | [**ForwardGeocodingPlacesGeocodingForwardGet**](docs/PlacesApi.md#forwardgeocodingplacesgeocodingforwardget) | **Get** /places/geocoding/forward | Forward Geocoding
*PlacesApi* | [**GetPlace**](docs/PlacesApi.md#getplace) | **Get** /places/places/{place_uuid} | Getplace
*PlacesApi* | [**ReverseGeocodingPlacesGeocodingReverseGet**](docs/PlacesApi.md#reversegeocodingplacesgeocodingreverseget) | **Get** /places/geocoding/reverse | Reverse Geocoding
*PlacesApi* | [**SearchPlaces**](docs/PlacesApi.md#searchplaces) | **Get** /places/places | Searchplaces
*PlacesApi* | [**UpdatePlace**](docs/PlacesApi.md#updateplace) | **Put** /places/places/{place_uuid} | Updateplace
*ProductsApi* | [**CreateProduct**](docs/ProductsApi.md#createproduct) | **Post** /products/products | Createproduct
*ProductsApi* | [**DeleteProduct**](docs/ProductsApi.md#deleteproduct) | **Delete** /products/products/{product_uuid} | Deleteproduct
*ProductsApi* | [**GetProduct**](docs/ProductsApi.md#getproduct) | **Get** /products/products/{product_uuid} | Getproduct
*ProductsApi* | [**SearchProductCatalog**](docs/ProductsApi.md#searchproductcatalog) | **Get** /products/products/{product_uuid}/catalogs | Searchproductcatalog
*ProductsApi* | [**SearchProducts**](docs/ProductsApi.md#searchproducts) | **Get** /products/products | Searchproducts
*ProductsApi* | [**UpdateProduct**](docs/ProductsApi.md#updateproduct) | **Put** /products/products/{product_uuid} | Updateproduct
*ProjectsApi* | [**AddMemberToProject**](docs/ProjectsApi.md#addmembertoproject) | **Post** /sam/projects/projects/{project_uuid}/members/{entity_uuid} | Addmembertoproject
*ProjectsApi* | [**ArchiveProject**](docs/ProjectsApi.md#archiveproject) | **Delete** /sam/projects/projects/{project_uuid} | Archiveproject
*ProjectsApi* | [**CreateProject**](docs/ProjectsApi.md#createproject) | **Post** /sam/projects/projects | Createproject
*ProjectsApi* | [**CreateProjectCredential**](docs/ProjectsApi.md#createprojectcredential) | **Post** /sam/projects/projects/{project_uuid}/credentials | Createprojectcredential
*ProjectsApi* | [**CreateProjectCredit**](docs/ProjectsApi.md#createprojectcredit) | **Post** /sam/projects/projects/{project_uuid}/credits | Createprojectcredit
*ProjectsApi* | [**CreateProjectInvoice**](docs/ProjectsApi.md#createprojectinvoice) | **Post** /sam/projects/projects/{project_uuid}/invoices | Createprojectinvoice
*ProjectsApi* | [**CreateProjectKey**](docs/ProjectsApi.md#createprojectkey) | **Post** /sam/projects/projects/{project_uuid}/keys | Createprojectkey
*ProjectsApi* | [**CreateUsageType**](docs/ProjectsApi.md#createusagetype) | **Post** /sam/projects/usage/types | Createusagetype
*ProjectsApi* | [**DeleteProjectCredential**](docs/ProjectsApi.md#deleteprojectcredential) | **Delete** /sam/projects/projects/{project_uuid}/credentials/{service_name} | Deleteprojectcredential
*ProjectsApi* | [**DeleteProjectKey**](docs/ProjectsApi.md#deleteprojectkey) | **Delete** /sam/projects/projects/{project_uuid}/keys | Deleteprojectkey
*ProjectsApi* | [**DeleteUsageType**](docs/ProjectsApi.md#deleteusagetype) | **Delete** /sam/projects/usage/types/{usage_type_key} | Deleteusagetype
*ProjectsApi* | [**GetAllProjectCredentials**](docs/ProjectsApi.md#getallprojectcredentials) | **Get** /sam/projects/projects/{project_uuid}/credentials | Getallprojectcredentials
*ProjectsApi* | [**GetAllProjectCredits**](docs/ProjectsApi.md#getallprojectcredits) | **Get** /sam/projects/projects/{project_uuid}/credits | Getallprojectcredits
*ProjectsApi* | [**GetAllProjectUsage**](docs/ProjectsApi.md#getallprojectusage) | **Get** /sam/projects/projects/{project_uuid}/usage | Getallprojectusage
*ProjectsApi* | [**GetMemberProjects**](docs/ProjectsApi.md#getmemberprojects) | **Get** /sam/projects/members/{entity_uuid}/projects | Getmemberprojects
*ProjectsApi* | [**GetProject**](docs/ProjectsApi.md#getproject) | **Get** /sam/projects/projects/{project_uuid} | Getproject
*ProjectsApi* | [**GetProjectCreditTransactions**](docs/ProjectsApi.md#getprojectcredittransactions) | **Get** /sam/projects/projects/{project_uuid}/credits/{credit_uuid}/transactions | Getprojectcredittransactions
*ProjectsApi* | [**GetProjectInvoice**](docs/ProjectsApi.md#getprojectinvoice) | **Get** /sam/projects/projects/{project_uuid}/invoices | Getprojectinvoice
*ProjectsApi* | [**GetProjectInvoiceHistory**](docs/ProjectsApi.md#getprojectinvoicehistory) | **Get** /sam/projects/projects/{project_uuid}/invoices/history | Getprojectinvoicehistory
*ProjectsApi* | [**GetProjectKeys**](docs/ProjectsApi.md#getprojectkeys) | **Get** /sam/projects/projects/{project_uuid}/keys | Getprojectkeys
*ProjectsApi* | [**GetProjectMembers**](docs/ProjectsApi.md#getprojectmembers) | **Get** /sam/projects/projects/{project_uuid}/members | Getprojectmembers
*ProjectsApi* | [**GetSpecificProjectCredential**](docs/ProjectsApi.md#getspecificprojectcredential) | **Get** /sam/projects/projects/{project_uuid}/credentials/{service_name} | Getspecificprojectcredential
*ProjectsApi* | [**GetSpecificProjectUsage**](docs/ProjectsApi.md#getspecificprojectusage) | **Get** /sam/projects/projects/{project_uuid}/usage/{usage_type_key} | Getspecificprojectusage
*ProjectsApi* | [**GetUsageType**](docs/ProjectsApi.md#getusagetype) | **Get** /sam/projects/usage/types/{usage_type_key} | Getusagetype
*ProjectsApi* | [**RemoveMemberFromProject**](docs/ProjectsApi.md#removememberfromproject) | **Delete** /sam/projects/projects/{project_uuid}/members/{entity_uuid} | Removememberfromproject
*ProjectsApi* | [**RevokeProjectCredit**](docs/ProjectsApi.md#revokeprojectcredit) | **Delete** /sam/projects/projects/{project_uuid}/credits/{credit_uuid} | Revokeprojectcredit
*ProjectsApi* | [**SearchProjects**](docs/ProjectsApi.md#searchprojects) | **Get** /sam/projects/projects | Searchprojects
*ProjectsApi* | [**SearchUsageType**](docs/ProjectsApi.md#searchusagetype) | **Get** /sam/projects/usage/types | Searchusagetype
*ProjectsApi* | [**UpdateProject**](docs/ProjectsApi.md#updateproject) | **Put** /sam/projects/projects/{project_uuid} | Updateproject
*ProjectsApi* | [**UpdateProjectCredential**](docs/ProjectsApi.md#updateprojectcredential) | **Put** /sam/projects/projects/{project_uuid}/credentials/{service_name} | Updateprojectcredential
*ProjectsApi* | [**UpdateUsageType**](docs/ProjectsApi.md#updateusagetype) | **Put** /sam/projects/usage/types/{usage_type_key} | Updateusagetype
*ReviewsApi* | [**CreateReview**](docs/ReviewsApi.md#createreview) | **Post** /products/reviews/types/{entity_type}/entities/{entity_uuid} | Create
*ReviewsApi* | [**DeleteReview**](docs/ReviewsApi.md#deletereview) | **Delete** /products/reviews/types/{entity_type}/entities/{entity_uuid}/reviews/{review_uuid} | Deletereview
*ReviewsApi* | [**GetReview**](docs/ReviewsApi.md#getreview) | **Get** /products/reviews/types/{entity_type}/entities/{entity_uuid}/reviews/{review_uuid} | Getreview
*ReviewsApi* | [**SearchReviews**](docs/ReviewsApi.md#searchreviews) | **Get** /products/reviews/types/{entity_type}/entities/{entity_uuid} | Searchreview
*ReviewsApi* | [**UpdateReview**](docs/ReviewsApi.md#updatereview) | **Put** /products/reviews/types/{entity_type}/entities/{entity_uuid}/reviews/{review_uuid} | Updatereview
*SecurityApi* | [**CreateEncryptionKey**](docs/SecurityApi.md#createencryptionkey) | **Post** /sam/security/encryption/categories/{category}/keys | Createencryptionkey
*SecurityApi* | [**CreateKey**](docs/SecurityApi.md#createkey) | **Post** /sam/security/keys | Createkey
*SecurityApi* | [**DeleteKey**](docs/SecurityApi.md#deletekey) | **Delete** /sam/security/keys/{key_uuid} | Deletekey
*SecurityApi* | [**GenerateToken**](docs/SecurityApi.md#generatetoken) | **Post** /sam/security/tokens | Generatetoken
*SecurityApi* | [**GetEncryptionKey**](docs/SecurityApi.md#getencryptionkey) | **Get** /sam/security/encryption/categories/{category}/keys | Getencryptionkey
*SecurityApi* | [**GetKey**](docs/SecurityApi.md#getkey) | **Get** /sam/security/keys/{key_uuid} | Getkey
*SecurityApi* | [**SearchKeys**](docs/SecurityApi.md#searchkeys) | **Get** /sam/security/keys | Searchkeys
*SecurityApi* | [**VerifyKey**](docs/SecurityApi.md#verifykey) | **Post** /sam/security/keys/verify | Verifykey
*StaffApi* | [**CreateStaff**](docs/StaffApi.md#createstaff) | **Post** /places/staff | Createstaff
*StaffApi* | [**DeleteStaff**](docs/StaffApi.md#deletestaff) | **Delete** /places/staff/{staff_uuid} | Deletestaff
*StaffApi* | [**GetStaff**](docs/StaffApi.md#getstaff) | **Get** /places/staff/{staff_uuid} | Getstaff
*StaffApi* | [**SearchStaff**](docs/StaffApi.md#searchstaff) | **Get** /places/staff | Searchstaff
*StaffApi* | [**UpdateStaff**](docs/StaffApi.md#updatestaff) | **Put** /places/staff/{staff_uuid} | Updatestaff
*SupportApi* | [**CreateContact**](docs/SupportApi.md#createcontact) | **Post** /sam/support/contact | Createcontact
*SupportApi* | [**CreateTicket**](docs/SupportApi.md#createticket) | **Post** /sam/support/projects/{project_uuid}/members/{member_uuid}/tickets | Createticket
*SupportApi* | [**ListTickets**](docs/SupportApi.md#listtickets) | **Get** /sam/support/projects/{project_uuid}/members/{member_uuid}/tickets | Listtickets
*SupportApi* | [**UpdateTicket**](docs/SupportApi.md#updateticket) | **Put** /sam/support/projects/{project_uuid}/members/{member_uuid}/tickets/{ticket_id} | Updateticket
*SupportApi* | [**ViewTicket**](docs/SupportApi.md#viewticket) | **Get** /sam/support/projects/{project_uuid}/members/{member_uuid}/tickets/{ticket_id} | Viewticket
*TagApi* | [**DeleteTag**](docs/TagApi.md#deletetag) | **Delete** /places/tags/{tag_uuid} | Deletetag
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /places/tags | Createtag
*TagsApi* | [**GetTag**](docs/TagsApi.md#gettag) | **Get** /places/tags/{tag_uuid} | Gettag
*TagsApi* | [**SearchTag**](docs/TagsApi.md#searchtag) | **Get** /places/tags | Searchtag
*TagsApi* | [**UpdateTag**](docs/TagsApi.md#updatetag) | **Put** /places/tags/{tag_uuid} | Updatetag
*UsersApi* | [**ConfirmSignup**](docs/UsersApi.md#confirmsignup) | **Post** /sam/users/auth/signup/confirm | Confirmsignup
*UsersApi* | [**CreateParticipant**](docs/UsersApi.md#createparticipant) | **Post** /sam/users/participants | Createparticipant
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /sam/users | Createuser
*UsersApi* | [**DeleteParticipant**](docs/UsersApi.md#deleteparticipant) | **Delete** /sam/users/participants/{participant_id} | Deleteparticipant
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /sam/users/{user_id} | Deleteuser
*UsersApi* | [**GetParticipant**](docs/UsersApi.md#getparticipant) | **Get** /sam/users/participants/{participant_id} | Getparticipant
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /sam/users/{user_id} | Getuser
*UsersApi* | [**Login**](docs/UsersApi.md#login) | **Post** /sam/users/auth/login | Login
*UsersApi* | [**RefreshToken**](docs/UsersApi.md#refreshtoken) | **Post** /sam/users/auth/{app_client}/refresh-token | Refreshtoken
*UsersApi* | [**ResetPassword**](docs/UsersApi.md#resetpassword) | **Post** /sam/users/auth/password/reset | Resetpassword
*UsersApi* | [**ResetPasswordConfirmation**](docs/UsersApi.md#resetpasswordconfirmation) | **Post** /sam/users/auth/password/reset/confirm | Resetpasswordconfirmation
*UsersApi* | [**SearchParticipants**](docs/UsersApi.md#searchparticipants) | **Get** /sam/users/participants | Searchparticipants
*UsersApi* | [**Signup**](docs/UsersApi.md#signup) | **Post** /sam/users/auth/signup | Signup
*UsersApi* | [**UpdateParticipant**](docs/UsersApi.md#updateparticipant) | **Put** /sam/users/participants/{participant_id} | Updateparticipant
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /sam/users/{user_id} | Updateuser
*UsersApi* | [**UserValidations**](docs/UsersApi.md#uservalidations) | **Post** /sam/users/validations/{field} | Uservalidations


## Documentation For Models

 - [AddressBase](docs/AddressBase.md)
 - [AlarmAcknowledge](docs/AlarmAcknowledge.md)
 - [AlarmAssign](docs/AlarmAssign.md)
 - [AlarmCreate](docs/AlarmCreate.md)
 - [AlarmIgnore](docs/AlarmIgnore.md)
 - [AlarmNote](docs/AlarmNote.md)
 - [AlarmResponse](docs/AlarmResponse.md)
 - [AlarmTerminate](docs/AlarmTerminate.md)
 - [AlarmTicket](docs/AlarmTicket.md)
 - [AppointmentBase](docs/AppointmentBase.md)
 - [AppointmentResponse](docs/AppointmentResponse.md)
 - [AttachPaymentToProject](docs/AttachPaymentToProject.md)
 - [Basic](docs/Basic.md)
 - [BasicMeta](docs/BasicMeta.md)
 - [BasicMetaCreate](docs/BasicMetaCreate.md)
 - [CatalogBase](docs/CatalogBase.md)
 - [CatalogReturn](docs/CatalogReturn.md)
 - [CategoryBase](docs/CategoryBase.md)
 - [CategoryResponse](docs/CategoryResponse.md)
 - [Company](docs/Company.md)
 - [CompanyBase](docs/CompanyBase.md)
 - [CompanyResponse](docs/CompanyResponse.md)
 - [Contact](docs/Contact.md)
 - [ContactBase](docs/ContactBase.md)
 - [ContactMethod](docs/ContactMethod.md)
 - [ContactResponse](docs/ContactResponse.md)
 - [CreateFile200Response](docs/CreateFile200Response.md)
 - [CreateKeyResponse](docs/CreateKeyResponse.md)
 - [CreateProjectCredential](docs/CreateProjectCredential.md)
 - [CreateProjectCredit](docs/CreateProjectCredit.md)
 - [CreateProjectInvoice](docs/CreateProjectInvoice.md)
 - [CreateReview](docs/CreateReview.md)
 - [CreateTicket](docs/CreateTicket.md)
 - [Credential](docs/Credential.md)
 - [CustomList](docs/CustomList.md)
 - [DatesMeta](docs/DatesMeta.md)
 - [DeleteFile200Response](docs/DeleteFile200Response.md)
 - [Detailed](docs/Detailed.md)
 - [DetailedMeta](docs/DetailedMeta.md)
 - [DetailedMetaCreate](docs/DetailedMetaCreate.md)
 - [DetailedMetaGet](docs/DetailedMetaGet.md)
 - [Discount](docs/Discount.md)
 - [Email](docs/Email.md)
 - [Field](docs/Field.md)
 - [FieldDynamo](docs/FieldDynamo.md)
 - [GetAppointment403Response](docs/GetAppointment403Response.md)
 - [GetInvoiceResponse](docs/GetInvoiceResponse.md)
 - [GetProjectCredential](docs/GetProjectCredential.md)
 - [GetProjectInvoiceHistory](docs/GetProjectInvoiceHistory.md)
 - [GetProjectInvoiceResponse](docs/GetProjectInvoiceResponse.md)
 - [GetSecret](docs/GetSecret.md)
 - [GetServiceServiceWithSpecsResponse](docs/GetServiceServiceWithSpecsResponse.md)
 - [GetServiceSpecResponse](docs/GetServiceSpecResponse.md)
 - [GetServiceSpecsResponse](docs/GetServiceSpecsResponse.md)
 - [GetTransactionResponse](docs/GetTransactionResponse.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [HeartbeatResponse](docs/HeartbeatResponse.md)
 - [History](docs/History.md)
 - [KpiResponse](docs/KpiResponse.md)
 - [LineItem](docs/LineItem.md)
 - [LoggingDynamo](docs/LoggingDynamo.md)
 - [MetaChildren](docs/MetaChildren.md)
 - [MetaCreate](docs/MetaCreate.md)
 - [MetaCustom](docs/MetaCustom.md)
 - [MetaDynamo](docs/MetaDynamo.md)
 - [MetaGet](docs/MetaGet.md)
 - [MetaSlugger](docs/MetaSlugger.md)
 - [Note](docs/Note.md)
 - [NoteBase](docs/NoteBase.md)
 - [NoteDynamoHistoryResponse](docs/NoteDynamoHistoryResponse.md)
 - [NoteDynamoResponse](docs/NoteDynamoResponse.md)
 - [NoteMeta](docs/NoteMeta.md)
 - [OptionGroup](docs/OptionGroup.md)
 - [Options](docs/Options.md)
 - [Page](docs/Page.md)
 - [Pagination](docs/Pagination.md)
 - [ParticipantCreate](docs/ParticipantCreate.md)
 - [ParticipantUpdate](docs/ParticipantUpdate.md)
 - [ParticipantUserReturn](docs/ParticipantUserReturn.md)
 - [Payment](docs/Payment.md)
 - [PaymentMethodResponse](docs/PaymentMethodResponse.md)
 - [PlaceBase](docs/PlaceBase.md)
 - [PlaceResponse](docs/PlaceResponse.md)
 - [ProductBase](docs/ProductBase.md)
 - [ProductReturn](docs/ProductReturn.md)
 - [ProjectCreditResponse](docs/ProjectCreditResponse.md)
 - [ProjectDB](docs/ProjectDB.md)
 - [ProjectsProjectCreate](docs/ProjectsProjectCreate.md)
 - [ProjectsProjectGet](docs/ProjectsProjectGet.md)
 - [ProjectsProjectMemberDB](docs/ProjectsProjectMemberDB.md)
 - [ProjectsProjectUpdate](docs/ProjectsProjectUpdate.md)
 - [ProjectsProjectUsageDB](docs/ProjectsProjectUsageDB.md)
 - [ProjectsUsageTypeCreate](docs/ProjectsUsageTypeCreate.md)
 - [ProjectsUsageTypeDB](docs/ProjectsUsageTypeDB.md)
 - [ProjectsUsageTypeGet](docs/ProjectsUsageTypeGet.md)
 - [ProjectsUsageTypeUnitPrice](docs/ProjectsUsageTypeUnitPrice.md)
 - [ProjectsUsageTypeUpdate](docs/ProjectsUsageTypeUpdate.md)
 - [ResponseAddmembertoproject](docs/ResponseAddmembertoproject.md)
 - [ResponseArchiveproject](docs/ResponseArchiveproject.md)
 - [ResponseCreatekey](docs/ResponseCreatekey.md)
 - [ResponseCreateprojectcredential](docs/ResponseCreateprojectcredential.md)
 - [ResponseCreateprojectinvoice](docs/ResponseCreateprojectinvoice.md)
 - [ResponseDeletekey](docs/ResponseDeletekey.md)
 - [ResponseDeleteprojectcredential](docs/ResponseDeleteprojectcredential.md)
 - [ResponseDeleteusagetype](docs/ResponseDeleteusagetype.md)
 - [ResponseGeneratetoken](docs/ResponseGeneratetoken.md)
 - [ResponseRemovememberfromproject](docs/ResponseRemovememberfromproject.md)
 - [ResponseRevokeprojectcredit](docs/ResponseRevokeprojectcredit.md)
 - [ResponseUpdateprojectcredential](docs/ResponseUpdateprojectcredential.md)
 - [SecurityCreateToken](docs/SecurityCreateToken.md)
 - [SecurityEncryptionKeyGet](docs/SecurityEncryptionKeyGet.md)
 - [SecurityEncryptionKeyResponse](docs/SecurityEncryptionKeyResponse.md)
 - [SecurityKeyCreate](docs/SecurityKeyCreate.md)
 - [SecurityKeyGet](docs/SecurityKeyGet.md)
 - [SecurityKeyVerify](docs/SecurityKeyVerify.md)
 - [Selection](docs/Selection.md)
 - [ServiceCreate](docs/ServiceCreate.md)
 - [ServiceMessageResponse](docs/ServiceMessageResponse.md)
 - [ServiceResponse](docs/ServiceResponse.md)
 - [StaffBase](docs/StaffBase.md)
 - [StaffResponse](docs/StaffResponse.md)
 - [StatsVitalsResponse](docs/StatsVitalsResponse.md)
 - [StripeAccountResponse](docs/StripeAccountResponse.md)
 - [StripeCustomerSecretResponse](docs/StripeCustomerSecretResponse.md)
 - [TagBase](docs/TagBase.md)
 - [TagResponse](docs/TagResponse.md)
 - [Tax](docs/Tax.md)
 - [TicketResponse](docs/TicketResponse.md)
 - [TicketsResponse](docs/TicketsResponse.md)
 - [UpdateFile200Response](docs/UpdateFile200Response.md)
 - [UpdateProjectCredentialRequest](docs/UpdateProjectCredentialRequest.md)
 - [UpdateReview](docs/UpdateReview.md)
 - [User](docs/User.md)
 - [UserConfirmation](docs/UserConfirmation.md)
 - [UserEmail](docs/UserEmail.md)
 - [UserFlags](docs/UserFlags.md)
 - [UserLogin](docs/UserLogin.md)
 - [UserLoginReturn](docs/UserLoginReturn.md)
 - [UserPasswordReset](docs/UserPasswordReset.md)
 - [UserPasswordResetConfirmation](docs/UserPasswordResetConfirmation.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserSignup](docs/UserSignup.md)
 - [UserSignupReturn](docs/UserSignupReturn.md)
 - [UserTokenReturn](docs/UserTokenReturn.md)
 - [UserValidations](docs/UserValidations.md)
 - [ValidationError](docs/ValidationError.md)
 - [Validations](docs/Validations.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@ehelply.com

