# \TestRunsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TestRunsIdStatisticsFilterPost**](TestRunsApi.md#ApiV2TestRunsIdStatisticsFilterPost) | **Post** /api/v2/testRuns/{id}/statistics/filter | Search for the test run test results and build statistics
[**ApiV2TestRunsIdTestPointsResultsGet**](TestRunsApi.md#ApiV2TestRunsIdTestPointsResultsGet) | **Get** /api/v2/testRuns/{id}/testPoints/results | Get test results from the test run grouped by test points
[**ApiV2TestRunsIdTestResultsBulkPut**](TestRunsApi.md#ApiV2TestRunsIdTestResultsBulkPut) | **Put** /api/v2/testRuns/{id}/testResults/bulk | Partial edit of multiple test results in the test run
[**ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet**](TestRunsApi.md#ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet) | **Get** /api/v2/testRuns/{id}/testResults/lastModified/modificationDate | Get modification date of last test result of the test run
[**ApiV2TestRunsSearchPost**](TestRunsApi.md#ApiV2TestRunsSearchPost) | **Post** /api/v2/testRuns/search | Search for test runs
[**CompleteTestRun**](TestRunsApi.md#CompleteTestRun) | **Post** /api/v2/testRuns/{id}/complete | Complete TestRun
[**CreateAndFillByAutoTests**](TestRunsApi.md#CreateAndFillByAutoTests) | **Post** /api/v2/testRuns/byAutoTests | Create test runs based on autotests and configurations
[**CreateAndFillByConfigurations**](TestRunsApi.md#CreateAndFillByConfigurations) | **Post** /api/v2/testRuns/byConfigurations | Create test runs picking the needed test points
[**CreateAndFillByWorkItems**](TestRunsApi.md#CreateAndFillByWorkItems) | **Post** /api/v2/testRuns/byWorkItems | Create test run based on configurations and work items
[**CreateEmpty**](TestRunsApi.md#CreateEmpty) | **Post** /api/v2/testRuns | Create empty TestRun
[**GetTestRunById**](TestRunsApi.md#GetTestRunById) | **Get** /api/v2/testRuns/{id} | Get TestRun by Id
[**SetAutoTestResultsForTestRun**](TestRunsApi.md#SetAutoTestResultsForTestRun) | **Post** /api/v2/testRuns/{id}/testResults | Send test results to the test runs in the system
[**StartTestRun**](TestRunsApi.md#StartTestRun) | **Post** /api/v2/testRuns/{id}/start | Start TestRun
[**StopTestRun**](TestRunsApi.md#StopTestRun) | **Post** /api/v2/testRuns/{id}/stop | Stop TestRun
[**UpdateEmpty**](TestRunsApi.md#UpdateEmpty) | **Put** /api/v2/testRuns | Update empty TestRun



## ApiV2TestRunsIdStatisticsFilterPost

> TestResultsStatisticsGetModel ApiV2TestRunsIdStatisticsFilterPost(ctx, id).TestResultsLocalFilterModel(testResultsLocalFilterModel).Execute()

Search for the test run test results and build statistics

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID
    testResultsLocalFilterModel := *openapiclient.NewTestResultsLocalFilterModel() // TestResultsLocalFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.ApiV2TestRunsIdStatisticsFilterPost(context.Background(), id).TestResultsLocalFilterModel(testResultsLocalFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.ApiV2TestRunsIdStatisticsFilterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestRunsIdStatisticsFilterPost`: TestResultsStatisticsGetModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.ApiV2TestRunsIdStatisticsFilterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdStatisticsFilterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testResultsLocalFilterModel** | [**TestResultsLocalFilterModel**](TestResultsLocalFilterModel.md) |  | 

### Return type

[**TestResultsStatisticsGetModel**](TestResultsStatisticsGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsIdTestPointsResultsGet

> []TestPointResultModel ApiV2TestRunsIdTestPointsResultsGet(ctx, id).Execute()

Get test results from the test run grouped by test points

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.ApiV2TestRunsIdTestPointsResultsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.ApiV2TestRunsIdTestPointsResultsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestRunsIdTestPointsResultsGet`: []TestPointResultModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.ApiV2TestRunsIdTestPointsResultsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdTestPointsResultsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TestPointResultModel**](TestPointResultModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestRunsIdTestResultsBulkPut

> ApiV2TestRunsIdTestResultsBulkPut(ctx, id).TestRunTestResultsPartialBulkSetModel(testRunTestResultsPartialBulkSetModel).Execute()

Partial edit of multiple test results in the test run

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID
    testRunTestResultsPartialBulkSetModel := *openapiclient.NewTestRunTestResultsPartialBulkSetModel() // TestRunTestResultsPartialBulkSetModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestRunsApi.ApiV2TestRunsIdTestResultsBulkPut(context.Background(), id).TestRunTestResultsPartialBulkSetModel(testRunTestResultsPartialBulkSetModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.ApiV2TestRunsIdTestResultsBulkPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdTestResultsBulkPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testRunTestResultsPartialBulkSetModel** | [**TestRunTestResultsPartialBulkSetModel**](TestRunTestResultsPartialBulkSetModel.md) |  | 

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


## ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet

> time.Time ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet(ctx, id).Execute()

Get modification date of last test result of the test run

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test run unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet`: time.Time
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.ApiV2TestRunsIdTestResultsLastModifiedModificationDateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test run unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsIdTestResultsLastModifiedModificationDateGetRequest struct via the builder pattern


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


## ApiV2TestRunsSearchPost

> []TestRunShortGetModel ApiV2TestRunsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestRunFilterModel(testRunFilterModel).Execute()

Search for test runs

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
    testRunFilterModel := *openapiclient.NewTestRunFilterModel() // TestRunFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.ApiV2TestRunsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestRunFilterModel(testRunFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.ApiV2TestRunsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestRunsSearchPost`: []TestRunShortGetModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.ApiV2TestRunsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestRunsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **testRunFilterModel** | [**TestRunFilterModel**](TestRunFilterModel.md) |  | 

### Return type

[**[]TestRunShortGetModel**](TestRunShortGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteTestRun

> CompleteTestRun(ctx, id).Execute()

Complete TestRun



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestRunsApi.CompleteTestRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.CompleteTestRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTestRunRequest struct via the builder pattern


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


## CreateAndFillByAutoTests

> TestRunV2GetModel CreateAndFillByAutoTests(ctx).TestRunFillByAutoTestsPostModel(testRunFillByAutoTestsPostModel).Execute()

Create test runs based on autotests and configurations



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
    testRunFillByAutoTestsPostModel := *openapiclient.NewTestRunFillByAutoTestsPostModel("ProjectId_example", []string{"ConfigurationIds_example"}, []string{"AutoTestExternalIds_example"}) // TestRunFillByAutoTestsPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.CreateAndFillByAutoTests(context.Background()).TestRunFillByAutoTestsPostModel(testRunFillByAutoTestsPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.CreateAndFillByAutoTests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAndFillByAutoTests`: TestRunV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.CreateAndFillByAutoTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndFillByAutoTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunFillByAutoTestsPostModel** | [**TestRunFillByAutoTestsPostModel**](TestRunFillByAutoTestsPostModel.md) |  | 

### Return type

[**TestRunV2GetModel**](TestRunV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAndFillByConfigurations

> TestRunV2GetModel CreateAndFillByConfigurations(ctx).TestRunFillByConfigurationsPostModel(testRunFillByConfigurationsPostModel).Execute()

Create test runs picking the needed test points



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
    testRunFillByConfigurationsPostModel := *openapiclient.NewTestRunFillByConfigurationsPostModel([]openapiclient.TestPointSelector{*openapiclient.NewTestPointSelector("ConfigurationId_example", []string{"WorkItemIds_example"})}, "ProjectId_example", "TestPlanId_example") // TestRunFillByConfigurationsPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.CreateAndFillByConfigurations(context.Background()).TestRunFillByConfigurationsPostModel(testRunFillByConfigurationsPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.CreateAndFillByConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAndFillByConfigurations`: TestRunV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.CreateAndFillByConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndFillByConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunFillByConfigurationsPostModel** | [**TestRunFillByConfigurationsPostModel**](TestRunFillByConfigurationsPostModel.md) |  | 

### Return type

[**TestRunV2GetModel**](TestRunV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAndFillByWorkItems

> TestRunV2GetModel CreateAndFillByWorkItems(ctx).TestRunFillByWorkItemsPostModel(testRunFillByWorkItemsPostModel).Execute()

Create test run based on configurations and work items



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
    testRunFillByWorkItemsPostModel := *openapiclient.NewTestRunFillByWorkItemsPostModel([]string{"ConfigurationIds_example"}, []string{"WorkItemIds_example"}, "ProjectId_example", "TestPlanId_example") // TestRunFillByWorkItemsPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.CreateAndFillByWorkItems(context.Background()).TestRunFillByWorkItemsPostModel(testRunFillByWorkItemsPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.CreateAndFillByWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAndFillByWorkItems`: TestRunV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.CreateAndFillByWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAndFillByWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunFillByWorkItemsPostModel** | [**TestRunFillByWorkItemsPostModel**](TestRunFillByWorkItemsPostModel.md) |  | 

### Return type

[**TestRunV2GetModel**](TestRunV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmpty

> TestRunV2GetModel CreateEmpty(ctx).TestRunV2PostShortModel(testRunV2PostShortModel).Execute()

Create empty TestRun



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
    testRunV2PostShortModel := *openapiclient.NewTestRunV2PostShortModel("d49af44b-dbd8-48b0-90e5-e065735d7229") // TestRunV2PostShortModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.CreateEmpty(context.Background()).TestRunV2PostShortModel(testRunV2PostShortModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.CreateEmpty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmpty`: TestRunV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.CreateEmpty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmptyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunV2PostShortModel** | [**TestRunV2PostShortModel**](TestRunV2PostShortModel.md) |  | 

### Return type

[**TestRunV2GetModel**](TestRunV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestRunById

> TestRunV2GetModel GetTestRunById(ctx, id).Execute()

Get TestRun by Id



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.GetTestRunById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.GetTestRunById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTestRunById`: TestRunV2GetModel
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.GetTestRunById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestRunByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestRunV2GetModel**](TestRunV2GetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAutoTestResultsForTestRun

> []string SetAutoTestResultsForTestRun(ctx, id).AutoTestResultsForTestRunModel(autoTestResultsForTestRunModel).Execute()

Send test results to the test runs in the system



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test Run internal identifier (GUID format)
    autoTestResultsForTestRunModel := []openapiclient.AutoTestResultsForTestRunModel{*openapiclient.NewAutoTestResultsForTestRunModel("ConfigurationId_example", "AutoTestExternalId_example", openapiclient.AvailableTestResultOutcome("Passed"))} // []AutoTestResultsForTestRunModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestRunsApi.SetAutoTestResultsForTestRun(context.Background(), id).AutoTestResultsForTestRunModel(autoTestResultsForTestRunModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.SetAutoTestResultsForTestRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAutoTestResultsForTestRun`: []string
    fmt.Fprintf(os.Stdout, "Response from `TestRunsApi.SetAutoTestResultsForTestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAutoTestResultsForTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTestResultsForTestRunModel** | [**[]AutoTestResultsForTestRunModel**](AutoTestResultsForTestRunModel.md) |  | 

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


## StartTestRun

> StartTestRun(ctx, id).Execute()

Start TestRun



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestRunsApi.StartTestRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.StartTestRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTestRunRequest struct via the builder pattern


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


## StopTestRun

> StopTestRun(ctx, id).Execute()

Stop TestRun



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test Run internal identifier (GUID format)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestRunsApi.StopTestRun(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.StopTestRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test Run internal identifier (GUID format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTestRunRequest struct via the builder pattern


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


## UpdateEmpty

> UpdateEmpty(ctx).TestRunV2PutModel(testRunV2PutModel).Execute()

Update empty TestRun



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
    testRunV2PutModel := *openapiclient.NewTestRunV2PutModel("d49af44b-dbd8-48b0-90e5-e065735d7229", "First run") // TestRunV2PutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestRunsApi.UpdateEmpty(context.Background()).TestRunV2PutModel(testRunV2PutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestRunsApi.UpdateEmpty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmptyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRunV2PutModel** | [**TestRunV2PutModel**](TestRunV2PutModel.md) |  | 

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

