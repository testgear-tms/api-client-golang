# \TestResultsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TestResultsIdAggregatedGet**](TestResultsApi.md#ApiV2TestResultsIdAggregatedGet) | **Get** /api/v2/testResults/{id}/aggregated | Get test result by ID aggregated with previous results
[**ApiV2TestResultsIdAttachmentsAttachmentIdPut**](TestResultsApi.md#ApiV2TestResultsIdAttachmentsAttachmentIdPut) | **Put** /api/v2/testResults/{id}/attachments/{attachmentId} | Attach file to the test result
[**ApiV2TestResultsIdAttachmentsInfoGet**](TestResultsApi.md#ApiV2TestResultsIdAttachmentsInfoGet) | **Get** /api/v2/testResults/{id}/attachments/info | Get test result attachments meta-information
[**ApiV2TestResultsIdGet**](TestResultsApi.md#ApiV2TestResultsIdGet) | **Get** /api/v2/testResults/{id} | Get test result by ID
[**ApiV2TestResultsIdPut**](TestResultsApi.md#ApiV2TestResultsIdPut) | **Put** /api/v2/testResults/{id} | Edit test result by ID
[**ApiV2TestResultsSearchPost**](TestResultsApi.md#ApiV2TestResultsSearchPost) | **Post** /api/v2/testResults/search | Search for test results
[**ApiV2TestResultsStatisticsFilterPost**](TestResultsApi.md#ApiV2TestResultsStatisticsFilterPost) | **Post** /api/v2/testResults/statistics/filter | Search for test results and extract statistics
[**CreateAttachment**](TestResultsApi.md#CreateAttachment) | **Post** /api/v2/testResults/{id}/attachments | Upload and link attachment to TestResult
[**DeleteAttachment**](TestResultsApi.md#DeleteAttachment) | **Delete** /api/v2/testResults/{id}/attachments/{attachmentId} | Remove attachment and unlink from TestResult
[**DownloadAttachment**](TestResultsApi.md#DownloadAttachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId} | Get attachment of TestResult
[**GetAttachment**](TestResultsApi.md#GetAttachment) | **Get** /api/v2/testResults/{id}/attachments/{attachmentId}/info | Get Metadata of TestResult&#39;s attachment
[**GetAttachments**](TestResultsApi.md#GetAttachments) | **Get** /api/v2/testResults/{id}/attachments | Get all attachments of TestResult



## ApiV2TestResultsIdAggregatedGet

> TestResultModel ApiV2TestResultsIdAggregatedGet(ctx, id).Execute()

Get test result by ID aggregated with previous results

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.ApiV2TestResultsIdAggregatedGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.ApiV2TestResultsIdAggregatedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestResultsIdAggregatedGet`: TestResultModel
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.ApiV2TestResultsIdAggregatedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdAggregatedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestResultModel**](TestResultModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsIdAttachmentsAttachmentIdPut

> ApiV2TestResultsIdAttachmentsAttachmentIdPut(ctx, id, attachmentId).Execute()

Attach file to the test result

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID
    attachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Attachment unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestResultsApi.ApiV2TestResultsIdAttachmentsAttachmentIdPut(context.Background(), id, attachmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.ApiV2TestResultsIdAttachmentsAttachmentIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 
**attachmentId** | **string** | Attachment unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdAttachmentsAttachmentIdPutRequest struct via the builder pattern


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


## ApiV2TestResultsIdAttachmentsInfoGet

> []AttachmentModel ApiV2TestResultsIdAttachmentsInfoGet(ctx, id).Execute()

Get test result attachments meta-information

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.ApiV2TestResultsIdAttachmentsInfoGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.ApiV2TestResultsIdAttachmentsInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestResultsIdAttachmentsInfoGet`: []AttachmentModel
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.ApiV2TestResultsIdAttachmentsInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdAttachmentsInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsIdGet

> TestResultModel ApiV2TestResultsIdGet(ctx, id).Execute()

Get test result by ID

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.ApiV2TestResultsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.ApiV2TestResultsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestResultsIdGet`: TestResultModel
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.ApiV2TestResultsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestResultModel**](TestResultModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsIdPut

> ApiV2TestResultsIdPut(ctx, id).TestResultUpdateModel(testResultUpdateModel).Execute()

Edit test result by ID

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Test result unique ID
    testResultUpdateModel := *openapiclient.NewTestResultUpdateModel() // TestResultUpdateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestResultsApi.ApiV2TestResultsIdPut(context.Background(), id).TestResultUpdateModel(testResultUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.ApiV2TestResultsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testResultUpdateModel** | [**TestResultUpdateModel**](TestResultUpdateModel.md) |  | 

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


## ApiV2TestResultsSearchPost

> []TestResultShortGetModel ApiV2TestResultsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestResultsFilterModel(testResultsFilterModel).Execute()

Search for test results

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
    testResultsFilterModel := *openapiclient.NewTestResultsFilterModel() // TestResultsFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.ApiV2TestResultsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).TestResultsFilterModel(testResultsFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.ApiV2TestResultsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestResultsSearchPost`: []TestResultShortGetModel
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.ApiV2TestResultsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **testResultsFilterModel** | [**TestResultsFilterModel**](TestResultsFilterModel.md) |  | 

### Return type

[**[]TestResultShortGetModel**](TestResultShortGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TestResultsStatisticsFilterPost

> TestResultsStatisticsGetModel ApiV2TestResultsStatisticsFilterPost(ctx).TestResultsFilterModel(testResultsFilterModel).Execute()

Search for test results and extract statistics

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
    testResultsFilterModel := *openapiclient.NewTestResultsFilterModel() // TestResultsFilterModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.ApiV2TestResultsStatisticsFilterPost(context.Background()).TestResultsFilterModel(testResultsFilterModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.ApiV2TestResultsStatisticsFilterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2TestResultsStatisticsFilterPost`: TestResultsStatisticsGetModel
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.ApiV2TestResultsStatisticsFilterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TestResultsStatisticsFilterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testResultsFilterModel** | [**TestResultsFilterModel**](TestResultsFilterModel.md) |  | 

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


## CreateAttachment

> string CreateAttachment(ctx, id).File(file).Execute()

Upload and link attachment to TestResult



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
    file := os.NewFile(1234, "some_file") // *os.File | Select file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.CreateAttachment(context.Background(), id).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.CreateAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttachment`: string
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.CreateAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttachmentRequest struct via the builder pattern


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


## DeleteAttachment

> DeleteAttachment(ctx, id, attachmentId).Execute()

Remove attachment and unlink from TestResult



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
    attachmentId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Attachment internal identifier (guid format)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestResultsApi.DeleteAttachment(context.Background(), id, attachmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.DeleteAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 
**attachmentId** | **string** | Attachment internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## DownloadAttachment

> DownloadAttachment(ctx, attachmentId, id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()

Get attachment of TestResult



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
    attachmentId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Attachment internal identifier (guid format)
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
    width := int32(56) // int32 | Width of the result image (optional)
    height := int32(56) // int32 | Height of the result image (optional)
    resizeType := openapiclient.ImageResizeType("Crop") // ImageResizeType | Type of resizing to apply to the result image (optional)
    backgroundColor := "backgroundColor_example" // string | Color of the background if the `resizeType` is `AddBackgroundStripes` (optional)
    preview := true // bool | If image must be converted to a preview (lower quality, no animation) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TestResultsApi.DownloadAttachment(context.Background(), attachmentId, id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.DownloadAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attachmentId** | **string** | Attachment internal identifier (guid format) | 
**id** | **string** | Test result internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **width** | **int32** | Width of the result image | 
 **height** | **int32** | Height of the result image | 
 **resizeType** | [**ImageResizeType**](ImageResizeType.md) | Type of resizing to apply to the result image | 
 **backgroundColor** | **string** | Color of the background if the &#x60;resizeType&#x60; is &#x60;AddBackgroundStripes&#x60; | 
 **preview** | **bool** | If image must be converted to a preview (lower quality, no animation) | 

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


## GetAttachment

> AttachmentModel GetAttachment(ctx, id, attachmentId).Execute()

Get Metadata of TestResult's attachment



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)
    attachmentId := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Attachment internal identifier (guid format)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.GetAttachment(context.Background(), id, attachmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.GetAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttachment`: AttachmentModel
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 
**attachmentId** | **string** | Attachment internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachments

> []AttachmentModel GetAttachments(ctx, id).Execute()

Get all attachments of TestResult



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Test result internal identifier (guid format)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TestResultsApi.GetAttachments(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TestResultsApi.GetAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttachments`: []AttachmentModel
    fmt.Fprintf(os.Stdout, "Response from `TestResultsApi.GetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Test result internal identifier (guid format) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

