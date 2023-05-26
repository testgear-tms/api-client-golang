# \AutoTestsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2AutoTestsFlakyBulkPost**](AutoTestsApi.md#ApiV2AutoTestsFlakyBulkPost) | **Post** /api/v2/autoTests/flaky/bulk | Set \&quot;Flaky\&quot; status for multiple autotests
[**ApiV2AutoTestsIdPatch**](AutoTestsApi.md#ApiV2AutoTestsIdPatch) | **Patch** /api/v2/autoTests/{id} | Patch auto test
[**ApiV2AutoTestsIdTestResultsSearchPost**](AutoTestsApi.md#ApiV2AutoTestsIdTestResultsSearchPost) | **Post** /api/v2/autoTests/{id}/testResults/search | Get test results history for autotest
[**ApiV2AutoTestsIdWorkItemsChangedIdGet**](AutoTestsApi.md#ApiV2AutoTestsIdWorkItemsChangedIdGet) | **Get** /api/v2/autoTests/{id}/workItems/changed/id | Get identifiers of changed linked work items
[**ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost**](AutoTestsApi.md#ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost) | **Post** /api/v2/autoTests/{id}/workItems/changed/{workItemId}/approve | Approve changes to work items linked to autotest
[**ApiV2AutoTestsSearchPost**](AutoTestsApi.md#ApiV2AutoTestsSearchPost) | **Post** /api/v2/autoTests/search | Search for autotests
[**CreateAutoTest**](AutoTestsApi.md#CreateAutoTest) | **Post** /api/v2/autoTests | Create autotest
[**CreateMultiple**](AutoTestsApi.md#CreateMultiple) | **Post** /api/v2/autoTests/bulk | Create multiple autotests
[**DeleteAutoTest**](AutoTestsApi.md#DeleteAutoTest) | **Delete** /api/v2/autoTests/{id} | Delete autotest
[**DeleteAutoTestLinkFromWorkItem**](AutoTestsApi.md#DeleteAutoTestLinkFromWorkItem) | **Delete** /api/v2/autoTests/{id}/workItems | Unlink autotest from work item
[**GetAllAutoTests**](AutoTestsApi.md#GetAllAutoTests) | **Get** /api/v2/autoTests | 
[**GetAutoTestAverageDuration**](AutoTestsApi.md#GetAutoTestAverageDuration) | **Get** /api/v2/autoTests/{id}/averageDuration | Get average autotest duration
[**GetAutoTestById**](AutoTestsApi.md#GetAutoTestById) | **Get** /api/v2/autoTests/{id} | Get autotest by internal or global ID
[**GetAutoTestChronology**](AutoTestsApi.md#GetAutoTestChronology) | **Get** /api/v2/autoTests/{id}/chronology | Get autotest chronology
[**GetTestRuns**](AutoTestsApi.md#GetTestRuns) | **Get** /api/v2/autoTests/{id}/testRuns | Get completed tests runs for autotests
[**GetWorkItemResults**](AutoTestsApi.md#GetWorkItemResults) | **Get** /api/v2/autoTests/{id}/testResultHistory | 
[**GetWorkItemsLinkedToAutoTest**](AutoTestsApi.md#GetWorkItemsLinkedToAutoTest) | **Get** /api/v2/autoTests/{id}/workItems | Get work items linked to autotest
[**LinkAutoTestToWorkItem**](AutoTestsApi.md#LinkAutoTestToWorkItem) | **Post** /api/v2/autoTests/{id}/workItems | Link autotest with work items
[**UpdateAutoTest**](AutoTestsApi.md#UpdateAutoTest) | **Put** /api/v2/autoTests | Update autotest
[**UpdateMultiple**](AutoTestsApi.md#UpdateMultiple) | **Put** /api/v2/autoTests/bulk | Update multiple autotests



## ApiV2AutoTestsFlakyBulkPost

> ApiV2AutoTestsFlakyBulkPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).FlakyBulkModel(flakyBulkModel).Execute()

Set \"Flaky\" status for multiple autotests



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
    flakyBulkModel := *openapiclient.NewFlakyBulkModel(false) // FlakyBulkModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.ApiV2AutoTestsFlakyBulkPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).FlakyBulkModel(flakyBulkModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.ApiV2AutoTestsFlakyBulkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsFlakyBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **flakyBulkModel** | [**FlakyBulkModel**](FlakyBulkModel.md) |  | 

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


## ApiV2AutoTestsIdPatch

> ApiV2AutoTestsIdPatch(ctx, id).Operation(operation).Execute()

Patch auto test



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Global Id of auto test
    operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.ApiV2AutoTestsIdPatch(context.Background(), id).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.ApiV2AutoTestsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Id of auto test | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdPatchRequest struct via the builder pattern


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


## ApiV2AutoTestsIdTestResultsSearchPost

> []AutotestResultHistoricalGetModel ApiV2AutoTestsIdTestResultsSearchPost(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestHistoricalResultSelectModel(autotestHistoricalResultSelectModel).Execute()

Get test results history for autotest



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
    id := "id_example" // string | Autotest identifier
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    autotestHistoricalResultSelectModel := *openapiclient.NewAutotestHistoricalResultSelectModel() // AutotestHistoricalResultSelectModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.ApiV2AutoTestsIdTestResultsSearchPost(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestHistoricalResultSelectModel(autotestHistoricalResultSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.ApiV2AutoTestsIdTestResultsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AutoTestsIdTestResultsSearchPost`: []AutotestResultHistoricalGetModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.ApiV2AutoTestsIdTestResultsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdTestResultsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **autotestHistoricalResultSelectModel** | [**AutotestHistoricalResultSelectModel**](AutotestHistoricalResultSelectModel.md) |  | 

### Return type

[**[]AutotestResultHistoricalGetModel**](AutotestResultHistoricalGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutoTestsIdWorkItemsChangedIdGet

> []string ApiV2AutoTestsIdWorkItemsChangedIdGet(ctx, id).Execute()

Get identifiers of changed linked work items



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.ApiV2AutoTestsIdWorkItemsChangedIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.ApiV2AutoTestsIdWorkItemsChangedIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AutoTestsIdWorkItemsChangedIdGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.ApiV2AutoTestsIdWorkItemsChangedIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdWorkItemsChangedIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost

> ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost(ctx, id, workItemId).Execute()

Approve changes to work items linked to autotest



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    workItemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost(context.Background(), id, workItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.ApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**workItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsIdWorkItemsChangedWorkItemIdApprovePostRequest struct via the builder pattern


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


## ApiV2AutoTestsSearchPost

> []AutoTestModel ApiV2AutoTestsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestsSelectModel(autotestsSelectModel).Execute()

Search for autotests

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
    autotestsSelectModel := *openapiclient.NewAutotestsSelectModel() // AutotestsSelectModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.ApiV2AutoTestsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).AutotestsSelectModel(autotestsSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.ApiV2AutoTestsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AutoTestsSearchPost`: []AutoTestModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.ApiV2AutoTestsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AutoTestsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **autotestsSelectModel** | [**AutotestsSelectModel**](AutotestsSelectModel.md) |  | 

### Return type

[**[]AutoTestModel**](AutoTestModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutoTest

> AutoTestModel CreateAutoTest(ctx).AutoTestPostModel(autoTestPostModel).Execute()

Create autotest



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
    autoTestPostModel := *openapiclient.NewAutoTestPostModel("ExternalId_example", "ProjectId_example", "Name_example") // AutoTestPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.CreateAutoTest(context.Background()).AutoTestPostModel(autoTestPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.CreateAutoTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutoTest`: AutoTestModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.CreateAutoTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestPostModel** | [**AutoTestPostModel**](AutoTestPostModel.md) |  | 

### Return type

[**AutoTestModel**](AutoTestModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMultiple

> []AutoTestModel CreateMultiple(ctx).AutoTestPostModel(autoTestPostModel).Execute()

Create multiple autotests



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
    autoTestPostModel := []openapiclient.AutoTestPostModel{*openapiclient.NewAutoTestPostModel("ExternalId_example", "ProjectId_example", "Name_example")} // []AutoTestPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.CreateMultiple(context.Background()).AutoTestPostModel(autoTestPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.CreateMultiple``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultiple`: []AutoTestModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.CreateMultiple`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestPostModel** | [**[]AutoTestPostModel**](AutoTestPostModel.md) |  | 

### Return type

[**[]AutoTestModel**](AutoTestModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoTest

> DeleteAutoTest(ctx, id).Execute()

Delete autotest



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
    id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.DeleteAutoTest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.DeleteAutoTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoTestRequest struct via the builder pattern


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


## DeleteAutoTestLinkFromWorkItem

> DeleteAutoTestLinkFromWorkItem(ctx, id).WorkItemId(workItemId).Execute()

Unlink autotest from work item



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
    id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier
    workItemId := "workItemId_example" // string | Work item internal (UUID) or global (integer) identifier (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.DeleteAutoTestLinkFromWorkItem(context.Background(), id).WorkItemId(workItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.DeleteAutoTestLinkFromWorkItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoTestLinkFromWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workItemId** | **string** | Work item internal (UUID) or global (integer) identifier | 

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


## GetAllAutoTests

> []AutoTestModel GetAllAutoTests(ctx).ProjectId(projectId).ExternalId(externalId).GlobalId(globalId).Namespace(namespace).IsNamespaceNull(isNamespaceNull).IncludeEmptyNamespaces(includeEmptyNamespaces).ClassName(className).IsClassnameNull(isClassnameNull).IncludeEmptyClassNames(includeEmptyClassNames).IsDeleted(isDeleted).Deleted(deleted).Labels(labels).StabilityMinimal(stabilityMinimal).MinStability(minStability).StabilityMaximal(stabilityMaximal).MaxStability(maxStability).IsFlaky(isFlaky).Flaky(flaky).IncludeSteps(includeSteps).IncludeLabels(includeLabels).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()



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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project internal ID (optional)
    externalId := "externalId_example" // string | Autotest external ID (optional)
    globalId := int64(789) // int64 | Autotest global ID (optional)
    namespace := "namespace_example" // string | Namespace in which autotest is located (optional)
    isNamespaceNull := true // bool | OBSOLETE: Use `includeEmptyNamespaces` instead (optional)
    includeEmptyNamespaces := true // bool | If result must contain autotests without namespace (optional)
    className := "className_example" // string | Name of class in which autotest is located (optional)
    isClassnameNull := true // bool | OBSOLETE: Use `includeEmptyClassNames` instead (optional)
    includeEmptyClassNames := true // bool | If result must contain autotests without class (optional)
    isDeleted := true // bool | OBSOLETE: Use `deleted` instead (optional)
    deleted := true // bool | Is autotest deleted (optional)
    labels := []string{"Inner_example"} // []string | Include only autotests with provided labels (optional)
    stabilityMinimal := int32(56) // int32 | OBSOLETE: Use `minStability` instead (optional)
    minStability := int32(56) // int32 | Minimum stability value of autotest (optional)
    stabilityMaximal := int32(56) // int32 | OBSOLETE: Use `maxStability` instead (optional)
    maxStability := int32(56) // int32 | Maximum stability value of autotest (optional)
    isFlaky := true // bool | OBSOLETE: Use `flaky` instead (optional)
    flaky := true // bool | Is autotest marked as \"Flaky\" (optional)
    includeSteps := true // bool | If result must also include autotest steps (optional)
    includeLabels := true // bool | If result must also include autotest labels (optional)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.GetAllAutoTests(context.Background()).ProjectId(projectId).ExternalId(externalId).GlobalId(globalId).Namespace(namespace).IsNamespaceNull(isNamespaceNull).IncludeEmptyNamespaces(includeEmptyNamespaces).ClassName(className).IsClassnameNull(isClassnameNull).IncludeEmptyClassNames(includeEmptyClassNames).IsDeleted(isDeleted).Deleted(deleted).Labels(labels).StabilityMinimal(stabilityMinimal).MinStability(minStability).StabilityMaximal(stabilityMaximal).MaxStability(maxStability).IsFlaky(isFlaky).Flaky(flaky).IncludeSteps(includeSteps).IncludeLabels(includeLabels).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.GetAllAutoTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAutoTests`: []AutoTestModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.GetAllAutoTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAutoTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | Project internal ID | 
 **externalId** | **string** | Autotest external ID | 
 **globalId** | **int64** | Autotest global ID | 
 **namespace** | **string** | Namespace in which autotest is located | 
 **isNamespaceNull** | **bool** | OBSOLETE: Use &#x60;includeEmptyNamespaces&#x60; instead | 
 **includeEmptyNamespaces** | **bool** | If result must contain autotests without namespace | 
 **className** | **string** | Name of class in which autotest is located | 
 **isClassnameNull** | **bool** | OBSOLETE: Use &#x60;includeEmptyClassNames&#x60; instead | 
 **includeEmptyClassNames** | **bool** | If result must contain autotests without class | 
 **isDeleted** | **bool** | OBSOLETE: Use &#x60;deleted&#x60; instead | 
 **deleted** | **bool** | Is autotest deleted | 
 **labels** | **[]string** | Include only autotests with provided labels | 
 **stabilityMinimal** | **int32** | OBSOLETE: Use &#x60;minStability&#x60; instead | 
 **minStability** | **int32** | Minimum stability value of autotest | 
 **stabilityMaximal** | **int32** | OBSOLETE: Use &#x60;maxStability&#x60; instead | 
 **maxStability** | **int32** | Maximum stability value of autotest | 
 **isFlaky** | **bool** | OBSOLETE: Use &#x60;flaky&#x60; instead | 
 **flaky** | **bool** | Is autotest marked as \&quot;Flaky\&quot; | 
 **includeSteps** | **bool** | If result must also include autotest steps | 
 **includeLabels** | **bool** | If result must also include autotest labels | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]AutoTestModel**](AutoTestModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestAverageDuration

> AutoTestAverageDurationModel GetAutoTestAverageDuration(ctx, id).Execute()

Get average autotest duration



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
    id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.GetAutoTestAverageDuration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.GetAutoTestAverageDuration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoTestAverageDuration`: AutoTestAverageDurationModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.GetAutoTestAverageDuration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestAverageDurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoTestAverageDurationModel**](AutoTestAverageDurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestById

> AutoTestModel GetAutoTestById(ctx, id).Execute()

Get autotest by internal or global ID



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
    id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.GetAutoTestById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.GetAutoTestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoTestById`: AutoTestModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.GetAutoTestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoTestModel**](AutoTestModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoTestChronology

> []TestResultChronologyModel GetAutoTestChronology(ctx, id).Execute()

Get autotest chronology



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
    id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.GetAutoTestChronology(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.GetAutoTestChronology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoTestChronology`: []TestResultChronologyModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.GetAutoTestChronology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTestChronologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestResultChronologyModel**](TestResultChronologyModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestRuns

> []TestRunShortModel GetTestRuns(ctx, id).Execute()

Get completed tests runs for autotests



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
    id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.GetTestRuns(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.GetTestRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestRuns`: []TestRunShortModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.GetTestRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestRunShortModel**](TestRunShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemResults

> []TestResultHistoryReportModel GetWorkItemResults(ctx, id).From(from).To(to).ConfigurationIds(configurationIds).TestPlanIds(testPlanIds).UserIds(userIds).Outcomes(outcomes).IsAutomated(isAutomated).Automated(automated).TestRunIds(testRunIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()



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
    id := "id_example" // string | 
    from := time.Now() // time.Time | Take results from this date (optional)
    to := time.Now() // time.Time | Take results until this date (optional)
    configurationIds := []string{"Inner_example"} // []string | Identifiers of test result configurations (optional)
    testPlanIds := []string{"Inner_example"} // []string | Identifiers of test plans which contain test results (optional)
    userIds := []string{"Inner_example"} // []string | Identifiers of users who set test results (optional)
    outcomes := []string{"Inner_example"} // []string | List of outcomes of test results (optional)
    isAutomated := true // bool | OBSOLETE: Use `Automated` instead (optional)
    automated := true // bool | If result must consist of only manual/automated test results (optional)
    testRunIds := []string{"Inner_example"} // []string | Identifiers of test runs which contain test results (optional)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.GetWorkItemResults(context.Background(), id).From(from).To(to).ConfigurationIds(configurationIds).TestPlanIds(testPlanIds).UserIds(userIds).Outcomes(outcomes).IsAutomated(isAutomated).Automated(automated).TestRunIds(testRunIds).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.GetWorkItemResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItemResults`: []TestResultHistoryReportModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.GetWorkItemResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Take results from this date | 
 **to** | **time.Time** | Take results until this date | 
 **configurationIds** | **[]string** | Identifiers of test result configurations | 
 **testPlanIds** | **[]string** | Identifiers of test plans which contain test results | 
 **userIds** | **[]string** | Identifiers of users who set test results | 
 **outcomes** | **[]string** | List of outcomes of test results | 
 **isAutomated** | **bool** | OBSOLETE: Use &#x60;Automated&#x60; instead | 
 **automated** | **bool** | If result must consist of only manual/automated test results | 
 **testRunIds** | **[]string** | Identifiers of test runs which contain test results | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestResultHistoryReportModel**](TestResultHistoryReportModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItemsLinkedToAutoTest

> []WorkItemIdentifierModel GetWorkItemsLinkedToAutoTest(ctx, id).IsDeleted(isDeleted).IsWorkItemDeleted(isWorkItemDeleted).Execute()

Get work items linked to autotest



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
    id := "id_example" // string | Specifies the autotest entity ID.<br />  You can copy it from the address bar in your web browser or use autotest GUID.
    isDeleted := true // bool | Specifies that a test is deleted or still relevant. (optional)
    isWorkItemDeleted := true // bool | OBSOLETE: Use `isDeleted` instead (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutoTestsApi.GetWorkItemsLinkedToAutoTest(context.Background(), id).IsDeleted(isDeleted).IsWorkItemDeleted(isWorkItemDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.GetWorkItemsLinkedToAutoTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItemsLinkedToAutoTest`: []WorkItemIdentifierModel
    fmt.Fprintf(os.Stdout, "Response from `AutoTestsApi.GetWorkItemsLinkedToAutoTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the autotest entity ID.&lt;br /&gt;  You can copy it from the address bar in your web browser or use autotest GUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemsLinkedToAutoTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isDeleted** | **bool** | Specifies that a test is deleted or still relevant. | 
 **isWorkItemDeleted** | **bool** | OBSOLETE: Use &#x60;isDeleted&#x60; instead | [default to false]

### Return type

[**[]WorkItemIdentifierModel**](WorkItemIdentifierModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkAutoTestToWorkItem

> LinkAutoTestToWorkItem(ctx, id).WorkItemIdModel(workItemIdModel).Execute()

Link autotest with work items



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
    id := "id_example" // string | Autotest internal (UUID) or global (integer) identifier
    workItemIdModel := *openapiclient.NewWorkItemIdModel("d49af44b-dbd8-48b0-90e5-e065735d7229") // WorkItemIdModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.LinkAutoTestToWorkItem(context.Background(), id).WorkItemIdModel(workItemIdModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.LinkAutoTestToWorkItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Autotest internal (UUID) or global (integer) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkAutoTestToWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workItemIdModel** | [**WorkItemIdModel**](WorkItemIdModel.md) |  | 

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


## UpdateAutoTest

> UpdateAutoTest(ctx).AutoTestPutModel(autoTestPutModel).Execute()

Update autotest



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
    autoTestPutModel := *openapiclient.NewAutoTestPutModel("ExternalId_example", "ProjectId_example", "Name_example") // AutoTestPutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.UpdateAutoTest(context.Background()).AutoTestPutModel(autoTestPutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.UpdateAutoTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestPutModel** | [**AutoTestPutModel**](AutoTestPutModel.md) |  | 

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


## UpdateMultiple

> UpdateMultiple(ctx).AutoTestPutModel(autoTestPutModel).Execute()

Update multiple autotests



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
    autoTestPutModel := []openapiclient.AutoTestPutModel{*openapiclient.NewAutoTestPutModel("ExternalId_example", "ProjectId_example", "Name_example")} // []AutoTestPutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutoTestsApi.UpdateMultiple(context.Background()).AutoTestPutModel(autoTestPutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutoTestsApi.UpdateMultiple``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestPutModel** | [**[]AutoTestPutModel**](AutoTestPutModel.md) |  | 

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

