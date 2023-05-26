# \ProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlobaAttributesToProject**](ProjectsApi.md#AddGlobaAttributesToProject) | **Post** /api/v2/projects/{id}/globalAttributes | Add global attributes to project
[**ApiV2ProjectsIdAttributesTemplatesSearchPost**](ProjectsApi.md#ApiV2ProjectsIdAttributesTemplatesSearchPost) | **Post** /api/v2/projects/{id}/attributes/templates/search | Search for custom attributes templates
[**ApiV2ProjectsIdAttributesTemplatesTemplateIdDelete**](ProjectsApi.md#ApiV2ProjectsIdAttributesTemplatesTemplateIdDelete) | **Delete** /api/v2/projects/{id}/attributes/templates/{templateId} | Delete CustomAttributeTemplate from Project
[**ApiV2ProjectsIdAttributesTemplatesTemplateIdPost**](ProjectsApi.md#ApiV2ProjectsIdAttributesTemplatesTemplateIdPost) | **Post** /api/v2/projects/{id}/attributes/templates/{templateId} | Add CustomAttributeTemplate to Project
[**ApiV2ProjectsIdFailureClassesGet**](ProjectsApi.md#ApiV2ProjectsIdFailureClassesGet) | **Get** /api/v2/projects/{id}/failureClasses | Get Project FailureClasses
[**ApiV2ProjectsIdFavoritePut**](ProjectsApi.md#ApiV2ProjectsIdFavoritePut) | **Put** /api/v2/projects/{id}/favorite | Mark Project as favorite
[**ApiV2ProjectsIdFiltersGet**](ProjectsApi.md#ApiV2ProjectsIdFiltersGet) | **Get** /api/v2/projects/{id}/filters | Get Project filters
[**ApiV2ProjectsIdPatch**](ProjectsApi.md#ApiV2ProjectsIdPatch) | **Patch** /api/v2/projects/{id} | Patch project
[**ApiV2ProjectsIdTestPlansAnalyticsGet**](ProjectsApi.md#ApiV2ProjectsIdTestPlansAnalyticsGet) | **Get** /api/v2/projects/{id}/testPlans/analytics | Get TestPlans analytics
[**ApiV2ProjectsIdTestPlansDeleteBulkPost**](ProjectsApi.md#ApiV2ProjectsIdTestPlansDeleteBulkPost) | **Post** /api/v2/projects/{id}/testPlans/delete/bulk | Delete multiple test plans
[**ApiV2ProjectsIdTestPlansNameExistsGet**](ProjectsApi.md#ApiV2ProjectsIdTestPlansNameExistsGet) | **Get** /api/v2/projects/{id}/testPlans/{name}/exists | Checks if TestPlan exists with the specified name exists for the project
[**ApiV2ProjectsIdTestPlansRestoreBulkPost**](ProjectsApi.md#ApiV2ProjectsIdTestPlansRestoreBulkPost) | **Post** /api/v2/projects/{id}/testPlans/restore/bulk | Restore multiple test plans
[**ApiV2ProjectsIdTestPlansSearchPost**](ProjectsApi.md#ApiV2ProjectsIdTestPlansSearchPost) | **Post** /api/v2/projects/{id}/testPlans/search | Get Project TestPlans with analytics
[**ApiV2ProjectsIdTestRunsActiveGet**](ProjectsApi.md#ApiV2ProjectsIdTestRunsActiveGet) | **Get** /api/v2/projects/{id}/testRuns/active | Get active Project TestRuns
[**ApiV2ProjectsIdTestRunsFullGet**](ProjectsApi.md#ApiV2ProjectsIdTestRunsFullGet) | **Get** /api/v2/projects/{id}/testRuns/full | Get Project TestRuns full models
[**ApiV2ProjectsIdWorkItemsSearchIdPost**](ProjectsApi.md#ApiV2ProjectsIdWorkItemsSearchIdPost) | **Post** /api/v2/projects/{id}/workItems/search/id | Search for work items and extract IDs only
[**ApiV2ProjectsIdWorkItemsSearchPost**](ProjectsApi.md#ApiV2ProjectsIdWorkItemsSearchPost) | **Post** /api/v2/projects/{id}/workItems/search | Search for work items
[**ApiV2ProjectsIdWorkItemsTagsGet**](ProjectsApi.md#ApiV2ProjectsIdWorkItemsTagsGet) | **Get** /api/v2/projects/{id}/workItems/tags | Get WorkItems Tags
[**ApiV2ProjectsNameNameExistsGet**](ProjectsApi.md#ApiV2ProjectsNameNameExistsGet) | **Get** /api/v2/projects/name/{name}/exists | 
[**ApiV2ProjectsSearchPost**](ProjectsApi.md#ApiV2ProjectsSearchPost) | **Post** /api/v2/projects/search | Search for projects
[**BackgroundImportProject**](ProjectsApi.md#BackgroundImportProject) | **Post** /api/v2/projects/import/json | Import project from JSON file in background job
[**BackgroundImportToExistingProject**](ProjectsApi.md#BackgroundImportToExistingProject) | **Post** /api/v2/projects/{id}/import/json | Import project from JSON file into existing project in background job
[**BackgroundImportZipProject**](ProjectsApi.md#BackgroundImportZipProject) | **Post** /api/v2/projects/import/zip | Import project from Zip file in background job
[**BackgroundImportZipToExistingProject**](ProjectsApi.md#BackgroundImportZipToExistingProject) | **Post** /api/v2/projects/{id}/import/zip | Import project from Zip file into existing project in background job
[**CallImport**](ProjectsApi.md#CallImport) | **Post** /api/v2/projects/import | Import project from JSON file
[**CreateCustomAttributeTestPlanProjectRelations**](ProjectsApi.md#CreateCustomAttributeTestPlanProjectRelations) | **Post** /api/v2/projects/{id}/testPlans/attributes | Add attributes to project&#39;s test plans
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /api/v2/projects | Create project
[**CreateProjectsAttribute**](ProjectsApi.md#CreateProjectsAttribute) | **Post** /api/v2/projects/{id}/attributes | Create project attribute
[**DeleteCustomAttributeTestPlanProjectRelations**](ProjectsApi.md#DeleteCustomAttributeTestPlanProjectRelations) | **Delete** /api/v2/projects/{id}/testPlans/attribute/{attributeId} | Delete attribute from project&#39;s test plans
[**DeleteProject**](ProjectsApi.md#DeleteProject) | **Delete** /api/v2/projects/{id} | Delete project
[**DeleteProjectAutoTests**](ProjectsApi.md#DeleteProjectAutoTests) | **Delete** /api/v2/projects/{id}/autoTests | Delete project
[**DeleteProjectsAttribute**](ProjectsApi.md#DeleteProjectsAttribute) | **Delete** /api/v2/projects/{id}/attributes/{attributeId} | Delete project attribute
[**Export**](ProjectsApi.md#Export) | **Post** /api/v2/projects/{id}/export | Export project as JSON file
[**ExportProjectJson**](ProjectsApi.md#ExportProjectJson) | **Post** /api/v2/projects/{id}/export/json | Export project as JSON file in background job
[**ExportProjectWithTestPlansJson**](ProjectsApi.md#ExportProjectWithTestPlansJson) | **Post** /api/v2/projects/{id}/export/testPlans/json | Export project as JSON file with test plans in background job
[**ExportProjectWithTestPlansZip**](ProjectsApi.md#ExportProjectWithTestPlansZip) | **Post** /api/v2/projects/{id}/export/testPlans/zip | Export project as Zip file with test plans in background job
[**ExportProjectZip**](ProjectsApi.md#ExportProjectZip) | **Post** /api/v2/projects/{id}/export/zip | Export project as Zip file in background job
[**ExportWithTestPlansAndConfigurations**](ProjectsApi.md#ExportWithTestPlansAndConfigurations) | **Post** /api/v2/projects/{id}/export-by-testPlans | Export project with test plans, test suites and test points as JSON file
[**GetAllProjects**](ProjectsApi.md#GetAllProjects) | **Get** /api/v2/projects | Get all projects
[**GetAttributeByProjectId**](ProjectsApi.md#GetAttributeByProjectId) | **Get** /api/v2/projects/{id}/attributes/{attributeId} | Get project attribute
[**GetAttributesByProjectId**](ProjectsApi.md#GetAttributesByProjectId) | **Get** /api/v2/projects/{id}/attributes | Get project attributes
[**GetAutoTestsNamespaces**](ProjectsApi.md#GetAutoTestsNamespaces) | **Get** /api/v2/projects/{id}/autoTestsNamespaces | Get namespaces of autotests in project
[**GetConfigurationsByProjectId**](ProjectsApi.md#GetConfigurationsByProjectId) | **Get** /api/v2/projects/{id}/configurations | Get project configurations
[**GetCustomAttributeTestPlanProjectRelations**](ProjectsApi.md#GetCustomAttributeTestPlanProjectRelations) | **Get** /api/v2/projects/{id}/testPlans/attributes | Get project&#39;s test plan attributes
[**GetProjectById**](ProjectsApi.md#GetProjectById) | **Get** /api/v2/projects/{id} | Get project by ID
[**GetSectionsByProjectId**](ProjectsApi.md#GetSectionsByProjectId) | **Get** /api/v2/projects/{id}/sections | Get project sections
[**GetTestPlansByProjectId**](ProjectsApi.md#GetTestPlansByProjectId) | **Get** /api/v2/projects/{id}/testPlans | Get project test plans
[**GetTestRunsByProjectId**](ProjectsApi.md#GetTestRunsByProjectId) | **Get** /api/v2/projects/{id}/testRuns | Get project test runs
[**GetWorkItemsByProjectId**](ProjectsApi.md#GetWorkItemsByProjectId) | **Get** /api/v2/projects/{id}/workItems | Get project work items
[**ImportToExistingProject**](ProjectsApi.md#ImportToExistingProject) | **Post** /api/v2/projects/{id}/import | Import project from JSON file into existing project
[**RestoreProject**](ProjectsApi.md#RestoreProject) | **Post** /api/v2/projects/{id}/restore | Restore project
[**SearchAttributesInProject**](ProjectsApi.md#SearchAttributesInProject) | **Post** /api/v2/projects/{id}/attributes/search | Search for attributes used in the project
[**SearchTestPlanAttributesInProject**](ProjectsApi.md#SearchTestPlanAttributesInProject) | **Post** /api/v2/projects/{id}/testPlans/attributes/search | Search for attributes used in the project test plans
[**UpdateCustomAttributeTestPlanProjectRelations**](ProjectsApi.md#UpdateCustomAttributeTestPlanProjectRelations) | **Put** /api/v2/projects/{id}/testPlans/attribute | Update attribute of project&#39;s test plans
[**UpdateProject**](ProjectsApi.md#UpdateProject) | **Put** /api/v2/projects | Update project
[**UpdateProjectsAttribute**](ProjectsApi.md#UpdateProjectsAttribute) | **Put** /api/v2/projects/{id}/attributes | Edit attribute of the project



## AddGlobaAttributesToProject

> AddGlobaAttributesToProject(ctx, id).RequestBody(requestBody).Execute()

Add global attributes to project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.AddGlobaAttributesToProject(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.AddGlobaAttributesToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGlobaAttributesToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdAttributesTemplatesSearchPost

> []ProjectCustomAttributeTemplateGetModel ApiV2ProjectsIdAttributesTemplatesSearchPost(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectCustomAttributesTemplatesFilterModel(projectCustomAttributesTemplatesFilterModel).Execute()

Search for custom attributes templates

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | 
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    projectCustomAttributesTemplatesFilterModel := *openapiclient.NewProjectCustomAttributesTemplatesFilterModel() // ProjectCustomAttributesTemplatesFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdAttributesTemplatesSearchPost(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectCustomAttributesTemplatesFilterModel(projectCustomAttributesTemplatesFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdAttributesTemplatesSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdAttributesTemplatesSearchPost`: []ProjectCustomAttributeTemplateGetModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdAttributesTemplatesSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdAttributesTemplatesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectCustomAttributesTemplatesFilterModel** | [**ProjectCustomAttributesTemplatesFilterModel**](ProjectCustomAttributesTemplatesFilterModel.md) |  | 

### Return type

[**[]ProjectCustomAttributeTemplateGetModel**](ProjectCustomAttributeTemplateGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdAttributesTemplatesTemplateIdDelete

> ApiV2ProjectsIdAttributesTemplatesTemplateIdDelete(ctx, id, templateId).Execute()

Delete CustomAttributeTemplate from Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | CustomAttributeTemplate internal (UUID) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdAttributesTemplatesTemplateIdDelete(context.Background(), id, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdAttributesTemplatesTemplateIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 
**templateId** | **string** | CustomAttributeTemplate internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdAttributesTemplatesTemplateIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdAttributesTemplatesTemplateIdPost

> ApiV2ProjectsIdAttributesTemplatesTemplateIdPost(ctx, id, templateId).Execute()

Add CustomAttributeTemplate to Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    templateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | CustomAttributeTemplate internal (UUID) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdAttributesTemplatesTemplateIdPost(context.Background(), id, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdAttributesTemplatesTemplateIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 
**templateId** | **string** | CustomAttributeTemplate internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdAttributesTemplatesTemplateIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdFailureClassesGet

> []FailureClassModel ApiV2ProjectsIdFailureClassesGet(ctx, id).IsDeleted(isDeleted).Execute()

Get Project FailureClasses



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    isDeleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdFailureClassesGet(context.Background(), id).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdFailureClassesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdFailureClassesGet`: []FailureClassModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdFailureClassesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdFailureClassesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 

### Return type

[**[]FailureClassModel**](FailureClassModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdFavoritePut

> ApiV2ProjectsIdFavoritePut(ctx, id).Execute()

Mark Project as favorite

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdFavoritePut(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdFavoritePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdFavoritePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdFiltersGet

> []FilterModel ApiV2ProjectsIdFiltersGet(ctx, id).Execute()

Get Project filters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdFiltersGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdFiltersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdFiltersGet`: []FilterModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdFiltersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdFiltersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FilterModel**](FilterModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdPatch

> ApiV2ProjectsIdPatch(ctx, id).Operation(operation).Execute()

Patch project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique or global Id of project
    operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdPatch(context.Background(), id).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global Id of project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operation** | [**[]Operation**](Operation.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestPlansAnalyticsGet

> []TestPlanWithAnalyticModel ApiV2ProjectsIdTestPlansAnalyticsGet(ctx, id).IsDeleted(isDeleted).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get TestPlans analytics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project internal (UUID) identifier
    isDeleted := true // bool |  (optional)
    mustUpdateCache := true // bool |  (optional) (default to false)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestPlansAnalyticsGet(context.Background(), id).IsDeleted(isDeleted).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestPlansAnalyticsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestPlansAnalyticsGet`: []TestPlanWithAnalyticModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestPlansAnalyticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestPlansAnalyticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 
 **mustUpdateCache** | **bool** |  | [default to false]
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestPlanWithAnalyticModel**](TestPlanWithAnalyticModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestPlansDeleteBulkPost

> []string ApiV2ProjectsIdTestPlansDeleteBulkPost(ctx, id).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()

Delete multiple test plans

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Unique or global ID of the project
    projectTestPlansFilterModel := *openapiclient.NewProjectTestPlansFilterModel() // ProjectTestPlansFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestPlansDeleteBulkPost(context.Background(), id).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestPlansDeleteBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestPlansDeleteBulkPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestPlansDeleteBulkPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestPlansDeleteBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectTestPlansFilterModel** | [**ProjectTestPlansFilterModel**](ProjectTestPlansFilterModel.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestPlansNameExistsGet

> bool ApiV2ProjectsIdTestPlansNameExistsGet(ctx, id, name).Execute()

Checks if TestPlan exists with the specified name exists for the project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project internal (UUID) or global (integer) identifier
    name := "name_example" // string | TestPlan name to check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestPlansNameExistsGet(context.Background(), id, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestPlansNameExistsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestPlansNameExistsGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestPlansNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 
**name** | **string** | TestPlan name to check | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestPlansNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestPlansRestoreBulkPost

> ApiV2ProjectsIdTestPlansRestoreBulkPost(ctx, id).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()

Restore multiple test plans

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Unique or global ID of the project
    projectTestPlansFilterModel := *openapiclient.NewProjectTestPlansFilterModel() // ProjectTestPlansFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestPlansRestoreBulkPost(context.Background(), id).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestPlansRestoreBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestPlansRestoreBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectTestPlansFilterModel** | [**ProjectTestPlansFilterModel**](ProjectTestPlansFilterModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestPlansSearchPost

> []TestPlanWithAnalyticModel ApiV2ProjectsIdTestPlansSearchPost(ctx, id).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()

Get Project TestPlans with analytics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    mustUpdateCache := true // bool |  (optional) (default to false)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    projectTestPlansFilterModel := *openapiclient.NewProjectTestPlansFilterModel() // ProjectTestPlansFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestPlansSearchPost(context.Background(), id).MustUpdateCache(mustUpdateCache).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectTestPlansFilterModel(projectTestPlansFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestPlansSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestPlansSearchPost`: []TestPlanWithAnalyticModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestPlansSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestPlansSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mustUpdateCache** | **bool** |  | [default to false]
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectTestPlansFilterModel** | [**ProjectTestPlansFilterModel**](ProjectTestPlansFilterModel.md) |  | 

### Return type

[**[]TestPlanWithAnalyticModel**](TestPlanWithAnalyticModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestRunsActiveGet

> []PublicTestRunModel ApiV2ProjectsIdTestRunsActiveGet(ctx, id).Execute()

Get active Project TestRuns



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestRunsActiveGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestRunsActiveGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestRunsActiveGet`: []PublicTestRunModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestRunsActiveGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestRunsActiveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PublicTestRunModel**](PublicTestRunModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdTestRunsFullGet

> []TestRunModel ApiV2ProjectsIdTestRunsFullGet(ctx, id).IncludeTestResults(includeTestResults).MustAggregateTestResults(mustAggregateTestResults).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get Project TestRuns full models



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    includeTestResults := true // bool |  (optional) (default to false)
    mustAggregateTestResults := true // bool |  (optional) (default to true)
    notStarted := true // bool |  (optional)
    inProgress := true // bool |  (optional)
    stopped := true // bool |  (optional)
    completed := true // bool |  (optional)
    createdDateFrom := time.Now() // time.Time |  (optional)
    createdDateTo := time.Now() // time.Time |  (optional)
    testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdTestRunsFullGet(context.Background(), id).IncludeTestResults(includeTestResults).MustAggregateTestResults(mustAggregateTestResults).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdTestRunsFullGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdTestRunsFullGet`: []TestRunModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdTestRunsFullGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdTestRunsFullGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTestResults** | **bool** |  | [default to false]
 **mustAggregateTestResults** | **bool** |  | [default to true]
 **notStarted** | **bool** |  | 
 **inProgress** | **bool** |  | 
 **stopped** | **bool** |  | 
 **completed** | **bool** |  | 
 **createdDateFrom** | **time.Time** |  | 
 **createdDateTo** | **time.Time** |  | 
 **testPlanId** | **string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestRunModel**](TestRunModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdWorkItemsSearchIdPost

> []string ApiV2ProjectsIdWorkItemsSearchIdPost(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemSelectModel(workItemSelectModel).Execute()

Search for work items and extract IDs only

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Unique or global ID of the project
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    workItemSelectModel := *openapiclient.NewWorkItemSelectModel() // WorkItemSelectModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdWorkItemsSearchIdPost(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemSelectModel(workItemSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdWorkItemsSearchIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdWorkItemsSearchIdPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdWorkItemsSearchIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdWorkItemsSearchIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **workItemSelectModel** | [**WorkItemSelectModel**](WorkItemSelectModel.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdWorkItemsSearchPost

> []WorkItemShortModel ApiV2ProjectsIdWorkItemsSearchPost(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemSelectModel(workItemSelectModel).Execute()

Search for work items

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Unique or global ID of the project
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    workItemSelectModel := *openapiclient.NewWorkItemSelectModel() // WorkItemSelectModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdWorkItemsSearchPost(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).WorkItemSelectModel(workItemSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdWorkItemsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdWorkItemsSearchPost`: []WorkItemShortModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdWorkItemsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdWorkItemsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **workItemSelectModel** | [**WorkItemSelectModel**](WorkItemSelectModel.md) |  | 

### Return type

[**[]WorkItemShortModel**](WorkItemShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsIdWorkItemsTagsGet

> []TagShortModel ApiV2ProjectsIdWorkItemsTagsGet(ctx, id).IsDeleted(isDeleted).Execute()

Get WorkItems Tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project internal (UUID) identifier
    isDeleted := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsIdWorkItemsTagsGet(context.Background(), id).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsIdWorkItemsTagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsIdWorkItemsTagsGet`: []TagShortModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsIdWorkItemsTagsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsIdWorkItemsTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** |  | 

### Return type

[**[]TagShortModel**](TagShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsNameNameExistsGet

> bool ApiV2ProjectsNameNameExistsGet(ctx, name).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsNameNameExistsGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsNameNameExistsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsNameNameExistsGet`: bool
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsNameNameExistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsNameNameExistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsSearchPost

> []ProjectModel ApiV2ProjectsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectsFilterModel(projectsFilterModel).Execute()

Search for projects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    projectsFilterModel := *openapiclient.NewProjectsFilterModel() // ProjectsFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ApiV2ProjectsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectsFilterModel(projectsFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ApiV2ProjectsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsSearchPost`: []ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ApiV2ProjectsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectsFilterModel** | [**ProjectsFilterModel**](ProjectsFilterModel.md) |  | 

### Return type

[**[]ProjectModel**](ProjectModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundImportProject

> string BackgroundImportProject(ctx).File(file).Execute()

Import project from JSON file in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.BackgroundImportProject(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.BackgroundImportProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.BackgroundImportProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundImportToExistingProject

> string BackgroundImportToExistingProject(ctx, id).File(file).Execute()

Import project from JSON file into existing project in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.BackgroundImportToExistingProject(context.Background(), id).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.BackgroundImportToExistingProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportToExistingProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.BackgroundImportToExistingProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportToExistingProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | Select file | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundImportZipProject

> string BackgroundImportZipProject(ctx).File(file).Execute()

Import project from Zip file in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.BackgroundImportZipProject(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.BackgroundImportZipProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportZipProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.BackgroundImportZipProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportZipProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackgroundImportZipToExistingProject

> string BackgroundImportZipToExistingProject(ctx, id).File(file).Execute()

Import project from Zip file into existing project in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.BackgroundImportZipToExistingProject(context.Background(), id).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.BackgroundImportZipToExistingProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackgroundImportZipToExistingProject`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.BackgroundImportZipToExistingProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackgroundImportZipToExistingProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | Select file | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallImport

> CallImport(ctx).IncludeAttachments(includeAttachments).File(file).Execute()

Import project from JSON file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    includeAttachments := true // bool | Enables attachment import. (optional) (default to false)
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.CallImport(context.Background()).IncludeAttachments(includeAttachments).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CallImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeAttachments** | **bool** | Enables attachment import. | [default to false]
 **file** | ***os.File** | Select file | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomAttributeTestPlanProjectRelations

> CreateCustomAttributeTestPlanProjectRelations(ctx, id).RequestBody(requestBody).Execute()

Add attributes to project's test plans



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.CreateCustomAttributeTestPlanProjectRelations(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateCustomAttributeTestPlanProjectRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> ProjectModel CreateProject(ctx).ProjectPostModel(projectPostModel).Execute()

Create project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectPostModel := *openapiclient.NewProjectPostModel("Name_example") // ProjectPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProject(context.Background()).ProjectPostModel(projectPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectPostModel** | [**ProjectPostModel**](ProjectPostModel.md) |  | 

### Return type

[**ProjectModel**](ProjectModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectsAttribute

> CustomAttributeModel CreateProjectsAttribute(ctx, id).CustomAttributePostModel(customAttributePostModel).Execute()

Create project attribute



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    customAttributePostModel := *openapiclient.NewCustomAttributePostModel(openapiclient.CustomAttributeTypesEnum("string"), "Name_example") // CustomAttributePostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProjectsAttribute(context.Background(), id).CustomAttributePostModel(customAttributePostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProjectsAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectsAttribute`: CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProjectsAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectsAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributePostModel** | [**CustomAttributePostModel**](CustomAttributePostModel.md) |  | 

### Return type

[**CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomAttributeTestPlanProjectRelations

> DeleteCustomAttributeTestPlanProjectRelations(ctx, id, attributeId).Execute()

Delete attribute from project's test plans



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    attributeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.DeleteCustomAttributeTestPlanProjectRelations(context.Background(), id, attributeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteCustomAttributeTestPlanProjectRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 
**attributeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, id).Execute()

Delete project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.DeleteProject(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectAutoTests

> DeleteProjectAutoTests(ctx, id).Execute()

Delete project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.DeleteProjectAutoTests(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProjectAutoTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectAutoTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectsAttribute

> DeleteProjectsAttribute(ctx, id, attributeId).Execute()

Delete project attribute



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    attributeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project attribute internal (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.DeleteProjectsAttribute(context.Background(), id, attributeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProjectsAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 
**attributeId** | **string** | Project attribute internal (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectsAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export

> *os.File Export(ctx, id).IncludeAttachments(includeAttachments).ProjectExportQueryModel(projectExportQueryModel).Execute()

Export project as JSON file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Specifies the ID of the project you want to export.
    includeAttachments := true // bool | Enables attachment export. (optional) (default to false)
    projectExportQueryModel := *openapiclient.NewProjectExportQueryModel() // ProjectExportQueryModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.Export(context.Background(), id).IncludeAttachments(includeAttachments).ProjectExportQueryModel(projectExportQueryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.Export`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the project you want to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAttachments** | **bool** | Enables attachment export. | [default to false]
 **projectExportQueryModel** | [**ProjectExportQueryModel**](ProjectExportQueryModel.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectJson

> string ExportProjectJson(ctx, id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()

Export project as JSON file in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
    projectExportQueryModel := *openapiclient.NewProjectExportQueryModel() // ProjectExportQueryModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ExportProjectJson(context.Background(), id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ExportProjectJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectJson`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ExportProjectJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportQueryModel** | [**ProjectExportQueryModel**](ProjectExportQueryModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectWithTestPlansJson

> string ExportProjectWithTestPlansJson(ctx, id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()

Export project as JSON file with test plans in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
    projectExportWithTestPlansPostModel := *openapiclient.NewProjectExportWithTestPlansPostModel() // ProjectExportWithTestPlansPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ExportProjectWithTestPlansJson(context.Background(), id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ExportProjectWithTestPlansJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectWithTestPlansJson`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ExportProjectWithTestPlansJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectWithTestPlansJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportWithTestPlansPostModel** | [**ProjectExportWithTestPlansPostModel**](ProjectExportWithTestPlansPostModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectWithTestPlansZip

> string ExportProjectWithTestPlansZip(ctx, id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()

Export project as Zip file with test plans in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
    projectExportWithTestPlansPostModel := *openapiclient.NewProjectExportWithTestPlansPostModel() // ProjectExportWithTestPlansPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ExportProjectWithTestPlansZip(context.Background(), id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ExportProjectWithTestPlansZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectWithTestPlansZip`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ExportProjectWithTestPlansZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectWithTestPlansZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportWithTestPlansPostModel** | [**ProjectExportWithTestPlansPostModel**](ProjectExportWithTestPlansPostModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProjectZip

> string ExportProjectZip(ctx, id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()

Export project as Zip file in background job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
    projectExportQueryModel := *openapiclient.NewProjectExportQueryModel() // ProjectExportQueryModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ExportProjectZip(context.Background(), id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).ProjectExportQueryModel(projectExportQueryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ExportProjectZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProjectZip`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ExportProjectZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportProjectZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **projectExportQueryModel** | [**ProjectExportQueryModel**](ProjectExportQueryModel.md) |  | 

### Return type

**string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWithTestPlansAndConfigurations

> *os.File ExportWithTestPlansAndConfigurations(ctx, id).IncludeAttachments(includeAttachments).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()

Export project with test plans, test suites and test points as JSON file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Specifies the ID of the project you want to export.
    includeAttachments := true // bool | Enables attachment export. (optional) (default to false)
    projectExportWithTestPlansPostModel := *openapiclient.NewProjectExportWithTestPlansPostModel() // ProjectExportWithTestPlansPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ExportWithTestPlansAndConfigurations(context.Background(), id).IncludeAttachments(includeAttachments).ProjectExportWithTestPlansPostModel(projectExportWithTestPlansPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ExportWithTestPlansAndConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportWithTestPlansAndConfigurations`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ExportWithTestPlansAndConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the project you want to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportWithTestPlansAndConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAttachments** | **bool** | Enables attachment export. | [default to false]
 **projectExportWithTestPlansPostModel** | [**ProjectExportWithTestPlansPostModel**](ProjectExportWithTestPlansPostModel.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjects

> []ProjectModel GetAllProjects(ctx).IsDeleted(isDeleted).ProjectName(projectName).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get all projects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    isDeleted := true // bool | If result must consist of only actual/deleted parameters (optional)
    projectName := "projectName_example" // string |  (optional)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetAllProjects(context.Background()).IsDeleted(isDeleted).ProjectName(projectName).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAllProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllProjects`: []ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAllProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeleted** | **bool** | If result must consist of only actual/deleted parameters | 
 **projectName** | **string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]ProjectModel**](ProjectModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributeByProjectId

> CustomAttributeModel GetAttributeByProjectId(ctx, id, attributeId).Execute()

Get project attribute



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    attributeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project attribute internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetAttributeByProjectId(context.Background(), id, attributeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAttributeByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeByProjectId`: CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAttributeByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 
**attributeId** | **string** | Project attribute internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributesByProjectId

> []CustomAttributeModel GetAttributesByProjectId(ctx, id).IsDeleted(isDeleted).Execute()

Get project attributes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    isDeleted := openapiclient.DeletionState("Any") // DeletionState |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetAttributesByProjectId(context.Background(), id).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAttributesByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributesByProjectId`: []CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAttributesByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributesByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | [**DeletionState**](DeletionState.md) |  | 

### Return type

[**[]CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestsNamespaces

> []AutoTestNamespaceModel GetAutoTestsNamespaces(ctx, id).Execute()

Get namespaces of autotests in project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetAutoTestsNamespaces(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetAutoTestsNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoTestsNamespaces`: []AutoTestNamespaceModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetAutoTestsNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestsNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AutoTestNamespaceModel**](AutoTestNamespaceModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationsByProjectId

> []ConfigurationModel GetConfigurationsByProjectId(ctx, id).Execute()

Get project configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetConfigurationsByProjectId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetConfigurationsByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationsByProjectId`: []ConfigurationModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetConfigurationsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomAttributeTestPlanProjectRelations

> []CustomAttributeModel GetCustomAttributeTestPlanProjectRelations(ctx, id).Execute()

Get project's test plan attributes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetCustomAttributeTestPlanProjectRelations(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetCustomAttributeTestPlanProjectRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomAttributeTestPlanProjectRelations`: []CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetCustomAttributeTestPlanProjectRelations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectById

> ProjectModel GetProjectById(ctx, id).Execute()

Get project by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProjectById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjectById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectById`: ProjectModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectModel**](ProjectModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSectionsByProjectId

> []SectionModel GetSectionsByProjectId(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get project sections



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetSectionsByProjectId(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetSectionsByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSectionsByProjectId`: []SectionModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetSectionsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSectionsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]SectionModel**](SectionModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestPlansByProjectId

> []TestPlanModel GetTestPlansByProjectId(ctx, id).IsDeleted(isDeleted).Execute()

Get project test plans



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    isDeleted := true // bool | If result must consist of only actual/archived test plans (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetTestPlansByProjectId(context.Background(), id).IsDeleted(isDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetTestPlansByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestPlansByProjectId`: []TestPlanModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetTestPlansByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestPlansByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** | If result must consist of only actual/archived test plans | 

### Return type

[**[]TestPlanModel**](TestPlanModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestRunsByProjectId

> []TestRunV2GetModel GetTestRunsByProjectId(ctx, id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get project test runs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    notStarted := true // bool |  (optional)
    inProgress := true // bool |  (optional)
    stopped := true // bool |  (optional)
    completed := true // bool |  (optional)
    createdDateFrom := time.Now() // time.Time |  (optional)
    createdDateTo := time.Now() // time.Time |  (optional)
    testPlanId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetTestRunsByProjectId(context.Background(), id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).CreatedDateFrom(createdDateFrom).CreatedDateTo(createdDateTo).TestPlanId(testPlanId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetTestRunsByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestRunsByProjectId`: []TestRunV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetTestRunsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestRunsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notStarted** | **bool** |  | 
 **inProgress** | **bool** |  | 
 **stopped** | **bool** |  | 
 **completed** | **bool** |  | 
 **createdDateFrom** | **time.Time** |  | 
 **createdDateTo** | **time.Time** |  | 
 **testPlanId** | **string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestRunV2GetModel**](TestRunV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemsByProjectId

> []WorkItemShortModel GetWorkItemsByProjectId(ctx, id).IsDeleted(isDeleted).TagNames(tagNames).IncludeIterations(includeIterations).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get project work items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    isDeleted := true // bool | If result must consist of only actual/deleted work items (optional) (default to false)
    tagNames := []string{"Inner_example"} // []string | List of tags to filter by (optional)
    includeIterations := true // bool |  (optional) (default to true)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetWorkItemsByProjectId(context.Background(), id).IsDeleted(isDeleted).TagNames(tagNames).IncludeIterations(includeIterations).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetWorkItemsByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItemsByProjectId`: []WorkItemShortModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetWorkItemsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** | If result must consist of only actual/deleted work items | [default to false]
 **tagNames** | **[]string** | List of tags to filter by | 
 **includeIterations** | **bool** |  | [default to true]
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]WorkItemShortModel**](WorkItemShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportToExistingProject

> ImportToExistingProject(ctx, id).IncludeAttachments(includeAttachments).File(file).Execute()

Import project from JSON file into existing project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    includeAttachments := true // bool |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.ImportToExistingProject(context.Background(), id).IncludeAttachments(includeAttachments).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ImportToExistingProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportToExistingProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAttachments** | **bool** |  | 
 **file** | ***os.File** | Select file | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreProject

> RestoreProject(ctx, id).Execute()

Restore project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.RestoreProject(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.RestoreProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAttributesInProject

> []CustomAttributeGetModel SearchAttributesInProject(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()

Search for attributes used in the project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Unique or global project ID
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    projectAttributesFilterModel := *openapiclient.NewProjectAttributesFilterModel() // ProjectAttributesFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.SearchAttributesInProject(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SearchAttributesInProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAttributesInProject`: []CustomAttributeGetModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SearchAttributesInProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAttributesInProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectAttributesFilterModel** | [**ProjectAttributesFilterModel**](ProjectAttributesFilterModel.md) |  | 

### Return type

[**[]CustomAttributeGetModel**](CustomAttributeGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTestPlanAttributesInProject

> []CustomAttributeGetModel SearchTestPlanAttributesInProject(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()

Search for attributes used in the project test plans

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Unique or global project ID
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    projectAttributesFilterModel := *openapiclient.NewProjectAttributesFilterModel() // ProjectAttributesFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.SearchTestPlanAttributesInProject(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ProjectAttributesFilterModel(projectAttributesFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.SearchTestPlanAttributesInProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTestPlanAttributesInProject`: []CustomAttributeGetModel
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.SearchTestPlanAttributesInProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchTestPlanAttributesInProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **projectAttributesFilterModel** | [**ProjectAttributesFilterModel**](ProjectAttributesFilterModel.md) |  | 

### Return type

[**[]CustomAttributeGetModel**](CustomAttributeGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomAttributeTestPlanProjectRelations

> UpdateCustomAttributeTestPlanProjectRelations(ctx, id).CustomAttributeTestPlanProjectRelationPutModel(customAttributeTestPlanProjectRelationPutModel).Execute()

Update attribute of project's test plans



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Project internal (UUID) or global (integer) identifier
    customAttributeTestPlanProjectRelationPutModel := *openapiclient.NewCustomAttributeTestPlanProjectRelationPutModel("Id_example", false, false) // CustomAttributeTestPlanProjectRelationPutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.UpdateCustomAttributeTestPlanProjectRelations(context.Background(), id).CustomAttributeTestPlanProjectRelationPutModel(customAttributeTestPlanProjectRelationPutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateCustomAttributeTestPlanProjectRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomAttributeTestPlanProjectRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributeTestPlanProjectRelationPutModel** | [**CustomAttributeTestPlanProjectRelationPutModel**](CustomAttributeTestPlanProjectRelationPutModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> UpdateProject(ctx).ProjectPutModel(projectPutModel).Execute()

Update project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectPutModel := *openapiclient.NewProjectPutModel("Id_example", "Name_example") // ProjectPutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.UpdateProject(context.Background()).ProjectPutModel(projectPutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectPutModel** | [**ProjectPutModel**](ProjectPutModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectsAttribute

> UpdateProjectsAttribute(ctx, id).CustomAttributePutModel(customAttributePutModel).Execute()

Edit attribute of the project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Unique or global project ID
    customAttributePutModel := *openapiclient.NewCustomAttributePutModel(openapiclient.CustomAttributeTypesEnum("string"), "Name_example") // CustomAttributePutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProjectsApi.UpdateProjectsAttribute(context.Background(), id).CustomAttributePutModel(customAttributePutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UpdateProjectsAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectsAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributePutModel** | [**CustomAttributePutModel**](CustomAttributePutModel.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

