# \TestPointsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TestPointsIdTestRunsGet**](TestPointsApi.md#ApiV2TestPointsIdTestRunsGet) | **Get** /api/v2/testPoints/{id}/testRuns | Get all test runs which use test point
[**ApiV2TestPointsIdWorkItemGet**](TestPointsApi.md#ApiV2TestPointsIdWorkItemGet) | **Get** /api/v2/testPoints/{id}/workItem | Get work item represented by test point
[**ApiV2TestPointsSearchIdPost**](TestPointsApi.md#ApiV2TestPointsSearchIdPost) | **Post** /api/v2/testPoints/search/id | Search for test points and extract IDs only
[**ApiV2TestPointsSearchPost**](TestPointsApi.md#ApiV2TestPointsSearchPost) | **Post** /api/v2/testPoints/search | Search for test points



## ApiV2TestPointsIdTestRunsGet

> []TestRunModel ApiV2TestPointsIdTestRunsGet(ctx, id).Execute()

Get all test runs which use test point

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test point unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPointsApi.ApiV2TestPointsIdTestRunsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPointsApi.ApiV2TestPointsIdTestRunsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPointsIdTestRunsGet`: []TestRunModel
    fmt.Fprintf(os.Stdout, "Response from `TestPointsApi.ApiV2TestPointsIdTestRunsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test point unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPointsIdTestRunsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiV2TestPointsIdWorkItemGet

> WorkItemModel ApiV2TestPointsIdWorkItemGet(ctx, id).Execute()

Get work item represented by test point

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test point unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPointsApi.ApiV2TestPointsIdWorkItemGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPointsApi.ApiV2TestPointsIdWorkItemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPointsIdWorkItemGet`: WorkItemModel
    fmt.Fprintf(os.Stdout, "Response from `TestPointsApi.ApiV2TestPointsIdWorkItemGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test point unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPointsIdWorkItemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItemModel**](WorkItemModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestPointsSearchIdPost

> []string ApiV2TestPointsSearchIdPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestPointFilterModel(testPointFilterModel).Execute()

Search for test points and extract IDs only

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
    testPointFilterModel := *openapiclient.NewTestPointFilterModel() // TestPointFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPointsApi.ApiV2TestPointsSearchIdPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestPointFilterModel(testPointFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPointsApi.ApiV2TestPointsSearchIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPointsSearchIdPost`: []string
    fmt.Fprintf(os.Stdout, "Response from `TestPointsApi.ApiV2TestPointsSearchIdPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPointsSearchIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **testPointFilterModel** | [**TestPointFilterModel**](TestPointFilterModel.md) |  | 

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


## ApiV2TestPointsSearchPost

> []TestPointShortGetModel ApiV2TestPointsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestPointFilterModel(testPointFilterModel).Execute()

Search for test points

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
    testPointFilterModel := *openapiclient.NewTestPointFilterModel() // TestPointFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestPointsApi.ApiV2TestPointsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestPointFilterModel(testPointFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestPointsApi.ApiV2TestPointsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestPointsSearchPost`: []TestPointShortGetModel
    fmt.Fprintf(os.Stdout, "Response from `TestPointsApi.ApiV2TestPointsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestPointsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **testPointFilterModel** | [**TestPointFilterModel**](TestPointFilterModel.md) |  | 

### Return type

[**[]TestPointShortGetModel**](TestPointShortGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

