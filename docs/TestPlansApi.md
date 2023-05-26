# \TestPlansApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTestPointsWithSections**](TestPlansApi.md#AddTestPointsWithSections) | **Post** /api/v2/testPlans/{id}/test-points/withSections | Add test-points to TestPlan with sections
[**AddWorkItemsWithSections**](TestPlansApi.md#AddWorkItemsWithSections) | **Post** /api/v2/testPlans/{id}/workItems/withSections | Add WorkItems to TestPlan with Sections as TestSuites
[**ApiV2TestPlansIdAnalyticsGet**](TestPlansApi.md#ApiV2TestPlansIdAnalyticsGet) | **Get** /api/v2/testPlans/{id}/analytics | Get analytics by TestPlan
[**ApiV2TestPlansIdAutobalancePost**](TestPlansApi.md#ApiV2TestPlansIdAutobalancePost) | **Post** /api/v2/testPlans/{id}/autobalance | Distribute test points between the users
[**ApiV2TestPlansIdConfigurationsGet**](TestPlansApi.md#ApiV2TestPlansIdConfigurationsGet) | **Get** /api/v2/testPlans/{id}/configurations | Get TestPlan configurations
[**ApiV2TestPlansIdExportTestPointsXlsxPost**](TestPlansApi.md#ApiV2TestPlansIdExportTestPointsXlsxPost) | **Post** /api/v2/testPlans/{id}/export/testPoints/xlsx | Export TestPoints from TestPlan in xls format
[**ApiV2TestPlansIdExportTestResultHistoryXlsxPost**](TestPlansApi.md#ApiV2TestPlansIdExportTestResultHistoryXlsxPost) | **Post** /api/v2/testPlans/{id}/export/testResultHistory/xlsx | Export TestResults history from TestPlan in xls format
[**ApiV2TestPlansIdHistoryGet**](TestPlansApi.md#ApiV2TestPlansIdHistoryGet) | **Get** /api/v2/testPlans/{id}/history | Get TestPlan history
[**ApiV2TestPlansIdLinksGet**](TestPlansApi.md#ApiV2TestPlansIdLinksGet) | **Get** /api/v2/testPlans/{id}/links | Get Links of TestPlan
[**ApiV2TestPlansIdPatch**](TestPlansApi.md#ApiV2TestPlansIdPatch) | **Patch** /api/v2/testPlans/{id} | Patch test plan
[**ApiV2TestPlansIdTestPointsLastResultsGet**](TestPlansApi.md#ApiV2TestPlansIdTestPointsLastResultsGet) | **Get** /api/v2/testPlans/{id}/testPoints/lastResults | Get TestPoints with last result from TestPlan
[**ApiV2TestPlansIdTestPointsResetPost**](TestPlansApi.md#ApiV2TestPlansIdTestPointsResetPost) | **Post** /api/v2/testPlans/{id}/testPoints/reset | Reset TestPoints status of TestPlan
[**ApiV2TestPlansIdTestPointsTesterDelete**](TestPlansApi.md#ApiV2TestPlansIdTestPointsTesterDelete) | **Delete** /api/v2/testPlans/{id}/testPoints/tester | Unassign users from multiple test points
[**ApiV2TestPlansIdTestPointsTesterUserIdPost**](TestPlansApi.md#ApiV2TestPlansIdTestPointsTesterUserIdPost) | **Post** /api/v2/testPlans/{id}/testPoints/tester/{userId} | Assign user as a tester to multiple test points
[**ApiV2TestPlansIdTestRunsGet**](TestPlansApi.md#ApiV2TestPlansIdTestRunsGet) | **Get** /api/v2/testPlans/{id}/testRuns | Get TestRuns of TestPlan
[**ApiV2TestPlansIdTestRunsSearchPost**](TestPlansApi.md#ApiV2TestPlansIdTestRunsSearchPost) | **Post** /api/v2/testPlans/{id}/testRuns/search | Search TestRuns of TestPlan
[**ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet**](TestPlansApi.md#ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet) | **Get** /api/v2/testPlans/{id}/testRuns/testResults/lastModified/modifiedDate | Get last modification date of test plan&#39;s test results
[**ApiV2TestPlansIdUnlockRequestPost**](TestPlansApi.md#ApiV2TestPlansIdUnlockRequestPost) | **Post** /api/v2/testPlans/{id}/unlock/request | Send unlock TestPlan notification
[**ApiV2TestPlansShortsPost**](TestPlansApi.md#ApiV2TestPlansShortsPost) | **Post** /api/v2/testPlans/shorts | Get TestPlans short models by Project identifiers
[**Clone**](TestPlansApi.md#Clone) | **Post** /api/v2/testPlans/{id}/clone | Clone TestPlan
[**Complete**](TestPlansApi.md#Complete) | **Post** /api/v2/testPlans/{id}/complete | Complete TestPlan
[**CreateTestPlan**](TestPlansApi.md#CreateTestPlan) | **Post** /api/v2/testPlans | Create TestPlan
[**DeleteTestPlan**](TestPlansApi.md#DeleteTestPlan) | **Delete** /api/v2/testPlans/{id} | Delete TestPlan
[**GetTestPlanById**](TestPlansApi.md#GetTestPlanById) | **Get** /api/v2/testPlans/{id} | Get TestPlan by Id
[**GetTestSuitesById**](TestPlansApi.md#GetTestSuitesById) | **Get** /api/v2/testPlans/{id}/testSuites | Get TestSuites Tree By Id
[**Pause**](TestPlansApi.md#Pause) | **Post** /api/v2/testPlans/{id}/pause | Pause TestPlan
[**RestoreTestPlan**](TestPlansApi.md#RestoreTestPlan) | **Post** /api/v2/testPlans/{id}/restore | Restore TestPlan
[**Start**](TestPlansApi.md#Start) | **Post** /api/v2/testPlans/{id}/start | Start TestPlan
[**UpdateTestPlan**](TestPlansApi.md#UpdateTestPlan) | **Put** /api/v2/testPlans | Update TestPlan



## AddTestPointsWithSections

> AddTestPointsWithSections(ctx, id).WorkItemSelectModel(workItemSelectModel).Execute()

Add test-points to TestPlan with sections

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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    workItemSelectModel := *openapiclient.NewWorkItemSelectModel() // WorkItemSelectModel | Filter object to retrieve work items for test-suite's project (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.AddTestPointsWithSections(context.Background(), id).WorkItemSelectModel(workItemSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.AddTestPointsWithSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTestPointsWithSectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workItemSelectModel** | [**WorkItemSelectModel**](WorkItemSelectModel.md) | Filter object to retrieve work items for test-suite&#39;s project | 

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


## AddWorkItemsWithSections

> AddWorkItemsWithSections(ctx, id).RequestBody(requestBody).Execute()

Add WorkItems to TestPlan with Sections as TestSuites



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.AddWorkItemsWithSections(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.AddWorkItemsWithSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkItemsWithSectionsRequest struct via the builder pattern


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


## ApiV2TestPlansIdAnalyticsGet

> TestPointAnalyticResult ApiV2TestPlansIdAnalyticsGet(ctx, id).Execute()

Get analytics by TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdAnalyticsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdAnalyticsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdAnalyticsGet`: TestPointAnalyticResult
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdAnalyticsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdAnalyticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestPointAnalyticResult**](TestPointAnalyticResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdAutobalancePost

> TestPlanWithTestSuiteTreeModel ApiV2TestPlansIdAutobalancePost(ctx, id).Testers(testers).Execute()

Distribute test points between the users

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
    id := "id_example" // string | Test plan unique or global ID
    testers := []string{"Inner_example"} // []string | Specifies a project user IDs to distribute (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdAutobalancePost(context.Background(), id).Testers(testers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdAutobalancePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdAutobalancePost`: TestPlanWithTestSuiteTreeModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdAutobalancePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan unique or global ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdAutobalancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testers** | **[]string** | Specifies a project user IDs to distribute | 

### Return type

[**TestPlanWithTestSuiteTreeModel**](TestPlanWithTestSuiteTreeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdConfigurationsGet

> []ConfigurationModel ApiV2TestPlansIdConfigurationsGet(ctx, id).Execute()

Get TestPlan configurations



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdConfigurationsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdConfigurationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdConfigurationsGet`: []ConfigurationModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdConfigurationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdConfigurationsGetRequest struct via the builder pattern


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


## ApiV2TestPlansIdExportTestPointsXlsxPost

> *os.File ApiV2TestPlansIdExportTestPointsXlsxPost(ctx, id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).GetXlsxTestPointsByTestPlanModel(getXlsxTestPointsByTestPlanModel).Execute()

Export TestPoints from TestPlan in xls format



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)
    getXlsxTestPointsByTestPlanModel := *openapiclient.NewGetXlsxTestPointsByTestPlanModel() // GetXlsxTestPointsByTestPlanModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdExportTestPointsXlsxPost(context.Background(), id).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).GetXlsxTestPointsByTestPlanModel(getXlsxTestPointsByTestPlanModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdExportTestPointsXlsxPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdExportTestPointsXlsxPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdExportTestPointsXlsxPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdExportTestPointsXlsxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeZoneOffsetInMinutes** | **int64** |  | 
 **getXlsxTestPointsByTestPlanModel** | [**GetXlsxTestPointsByTestPlanModel**](GetXlsxTestPointsByTestPlanModel.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdExportTestResultHistoryXlsxPost

> *os.File ApiV2TestPlansIdExportTestResultHistoryXlsxPost(ctx, id).MustReturnOnlyLastTestResult(mustReturnOnlyLastTestResult).IncludeSteps(includeSteps).IncludeDeletedTestSuites(includeDeletedTestSuites).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).Execute()

Export TestResults history from TestPlan in xls format



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    mustReturnOnlyLastTestResult := true // bool |  (optional)
    includeSteps := true // bool |  (optional)
    includeDeletedTestSuites := true // bool |  (optional)
    timeZoneOffsetInMinutes := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdExportTestResultHistoryXlsxPost(context.Background(), id).MustReturnOnlyLastTestResult(mustReturnOnlyLastTestResult).IncludeSteps(includeSteps).IncludeDeletedTestSuites(includeDeletedTestSuites).TimeZoneOffsetInMinutes(timeZoneOffsetInMinutes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdExportTestResultHistoryXlsxPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdExportTestResultHistoryXlsxPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdExportTestResultHistoryXlsxPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdExportTestResultHistoryXlsxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mustReturnOnlyLastTestResult** | **bool** |  | 
 **includeSteps** | **bool** |  | 
 **includeDeletedTestSuites** | **bool** |  | 
 **timeZoneOffsetInMinutes** | **int64** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdHistoryGet

> []TestPlanChangeModel ApiV2TestPlansIdHistoryGet(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get TestPlan history



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdHistoryGet(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdHistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdHistoryGet`: []TestPlanChangeModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdHistoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestPlanChangeModel**](TestPlanChangeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdLinksGet

> []TestPlanLink ApiV2TestPlansIdLinksGet(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Execute()

Get Links of TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    skip := int32(56) // int32 |  (optional)
    take := int32(56) // int32 |  (optional)
    orderBy := "orderBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdLinksGet(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdLinksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdLinksGet`: []TestPlanLink
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdLinksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdLinksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | 
 **take** | **int32** |  | 
 **orderBy** | **string** |  | 

### Return type

[**[]TestPlanLink**](TestPlanLink.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdPatch

> ApiV2TestPlansIdPatch(ctx, id).Operation(operation).Execute()

Patch test plan



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the test plan
    operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.ApiV2TestPlansIdPatch(context.Background(), id).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of the test plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdPatchRequest struct via the builder pattern


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


## ApiV2TestPlansIdTestPointsLastResultsGet

> []TestPointWithLastResultModel ApiV2TestPlansIdTestPointsLastResultsGet(ctx, id).TesterId(testerId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get TestPoints with last result from TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    testerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdTestPointsLastResultsGet(context.Background(), id).TesterId(testerId).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdTestPointsLastResultsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdTestPointsLastResultsGet`: []TestPointWithLastResultModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdTestPointsLastResultsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdTestPointsLastResultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testerId** | **string** |  | 
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 

### Return type

[**[]TestPointWithLastResultModel**](TestPointWithLastResultModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdTestPointsResetPost

> ApiV2TestPlansIdTestPointsResetPost(ctx, id).RequestBody(requestBody).Execute()

Reset TestPoints status of TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.ApiV2TestPlansIdTestPointsResetPost(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdTestPointsResetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdTestPointsResetPostRequest struct via the builder pattern


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


## ApiV2TestPlansIdTestPointsTesterDelete

> []string ApiV2TestPlansIdTestPointsTesterDelete(ctx, id).TestPointSelectModel(testPointSelectModel).Execute()

Unassign users from multiple test points

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
    id := "id_example" // string | Unique or global ID of the test plan
    testPointSelectModel := *openapiclient.NewTestPointSelectModel() // TestPointSelectModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdTestPointsTesterDelete(context.Background(), id).TestPointSelectModel(testPointSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdTestPointsTesterDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdTestPointsTesterDelete`: []string
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdTestPointsTesterDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the test plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdTestPointsTesterDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testPointSelectModel** | [**TestPointSelectModel**](TestPointSelectModel.md) |  | 

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


## ApiV2TestPlansIdTestPointsTesterUserIdPost

> []string ApiV2TestPlansIdTestPointsTesterUserIdPost(ctx, id, userId).TestPointSelectModel(testPointSelectModel).Execute()

Assign user as a tester to multiple test points

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
    id := "id_example" // string | Unique or global ID of the test plan
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the user
    testPointSelectModel := *openapiclient.NewTestPointSelectModel() // TestPointSelectModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdTestPointsTesterUserIdPost(context.Background(), id, userId).TestPointSelectModel(testPointSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdTestPointsTesterUserIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdTestPointsTesterUserIdPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdTestPointsTesterUserIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique or global ID of the test plan | 
**userId** | **string** | Unique ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdTestPointsTesterUserIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **testPointSelectModel** | [**TestPointSelectModel**](TestPointSelectModel.md) |  | 

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


## ApiV2TestPlansIdTestRunsGet

> []TestRunModel ApiV2TestPlansIdTestRunsGet(ctx, id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()

Get TestRuns of TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    notStarted := true // bool |  (optional)
    inProgress := true // bool |  (optional)
    stopped := true // bool |  (optional)
    completed := true // bool |  (optional)
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdTestRunsGet(context.Background(), id).NotStarted(notStarted).InProgress(inProgress).Stopped(stopped).Completed(completed).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdTestRunsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdTestRunsGet`: []TestRunModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdTestRunsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdTestRunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notStarted** | **bool** |  | 
 **inProgress** | **bool** |  | 
 **stopped** | **bool** |  | 
 **completed** | **bool** |  | 
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


## ApiV2TestPlansIdTestRunsSearchPost

> []TestRunModel ApiV2TestPlansIdTestRunsSearchPost(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestRunSearchQueryModel(testRunSearchQueryModel).Execute()

Search TestRuns of TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier
    skip := int32(56) // int32 | Amount of items to be skipped (offset) (optional)
    take := int32(56) // int32 | Amount of items to be taken (limit) (optional)
    orderBy := "orderBy_example" // string | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) (optional)
    searchField := "searchField_example" // string | Property name for searching (optional)
    searchValue := "searchValue_example" // string | Value for searching (optional)
    testRunSearchQueryModel := *openapiclient.NewTestRunSearchQueryModel() // TestRunSearchQueryModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdTestRunsSearchPost(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestRunSearchQueryModel(testRunSearchQueryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdTestRunsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdTestRunsSearchPost`: []TestRunModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdTestRunsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdTestRunsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **testRunSearchQueryModel** | [**TestRunSearchQueryModel**](TestRunSearchQueryModel.md) |  | 

### Return type

[**[]TestRunModel**](TestRunModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet

> time.Time ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet(ctx, id).Execute()

Get last modification date of test plan's test results

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
    id := "id_example" // string | Test plan unique or global ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet`: time.Time
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan unique or global ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdTestRunsTestResultsLastModifiedModifiedDateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**time.Time**](time.Time.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPlansIdUnlockRequestPost

> ApiV2TestPlansIdUnlockRequestPost(ctx, id).Execute()

Send unlock TestPlan notification



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.ApiV2TestPlansIdUnlockRequestPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansIdUnlockRequestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansIdUnlockRequestPostRequest struct via the builder pattern


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


## ApiV2TestPlansShortsPost

> []TestPlanShortModel ApiV2TestPlansShortsPost(ctx).IsDeleted(isDeleted).RequestBody(requestBody).Execute()

Get TestPlans short models by Project identifiers



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
    isDeleted := true // bool |  (optional)
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.ApiV2TestPlansShortsPost(context.Background()).IsDeleted(isDeleted).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.ApiV2TestPlansShortsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPlansShortsPost`: []TestPlanShortModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.ApiV2TestPlansShortsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPlansShortsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDeleted** | **bool** |  | 
 **requestBody** | **[]string** |  | 

### Return type

[**[]TestPlanShortModel**](TestPlanShortModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Clone

> TestPlanModel Clone(ctx, id).Execute()

Clone TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.Clone(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.Clone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Clone`: TestPlanModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.Clone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestPlanModel**](TestPlanModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Complete

> Complete(ctx, id).Execute()

Complete TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.Complete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.Complete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteRequest struct via the builder pattern


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


## CreateTestPlan

> TestPlanModel CreateTestPlan(ctx).TestPlanPostModel(testPlanPostModel).Execute()

Create TestPlan



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
    testPlanPostModel := *openapiclient.NewTestPlanPostModel("Base test plan", "d49af44b-dbd8-48b0-90e5-e065735d7229") // TestPlanPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.CreateTestPlan(context.Background()).TestPlanPostModel(testPlanPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.CreateTestPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTestPlan`: TestPlanModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.CreateTestPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTestPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testPlanPostModel** | [**TestPlanPostModel**](TestPlanPostModel.md) |  | 

### Return type

[**TestPlanModel**](TestPlanModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTestPlan

> DeleteTestPlan(ctx, id).Execute()

Delete TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.DeleteTestPlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.DeleteTestPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTestPlanRequest struct via the builder pattern


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


## GetTestPlanById

> TestPlanModel GetTestPlanById(ctx, id).Execute()

Get TestPlan by Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.GetTestPlanById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.GetTestPlanById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestPlanById`: TestPlanModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.GetTestPlanById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestPlanByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestPlanModel**](TestPlanModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestSuitesById

> []TestSuiteV2TreeModel GetTestSuitesById(ctx, id).Execute()

Get TestSuites Tree By Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPlansApi.GetTestSuitesById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.GetTestSuitesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestSuitesById`: []TestSuiteV2TreeModel
    fmt.Fprintf(os.Stdout, "Response from `TestPlansApi.GetTestSuitesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestSuitesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestSuiteV2TreeModel**](TestSuiteV2TreeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Pause

> Pause(ctx, id).Execute()

Pause TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.Pause(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.Pause``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseRequest struct via the builder pattern


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


## RestoreTestPlan

> RestoreTestPlan(ctx, id).Execute()

Restore TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.RestoreTestPlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.RestoreTestPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreTestPlanRequest struct via the builder pattern


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


## Start

> Start(ctx, id).Execute()

Start TestPlan



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test plan internal (guid format) or global (int  format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.Start(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.Start``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test plan internal (guid format) or global (int  format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartRequest struct via the builder pattern


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


## UpdateTestPlan

> UpdateTestPlan(ctx).TestPlanPutModel(testPlanPutModel).Execute()

Update TestPlan



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
    testPlanPutModel := *openapiclient.NewTestPlanPutModel("d49af44b-dbd8-48b0-90e5-e065735d7229", "Base test plan", "d49af44b-dbd8-48b0-90e5-e065735d7229") // TestPlanPutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestPlansApi.UpdateTestPlan(context.Background()).TestPlanPutModel(testPlanPutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPlansApi.UpdateTestPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTestPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testPlanPutModel** | [**TestPlanPutModel**](TestPlanPutModel.md) |  | 

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

