# \AttachmentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2AttachmentsIdDelete**](AttachmentsApi.md#ApiV2AttachmentsIdDelete) | **Delete** /api/v2/attachments/{id} | Delete attachment file
[**ApiV2AttachmentsIdGet**](AttachmentsApi.md#ApiV2AttachmentsIdGet) | **Get** /api/v2/attachments/{id} | Download attachment file
[**ApiV2AttachmentsOccupiedFileStorageSizeGet**](AttachmentsApi.md#ApiV2AttachmentsOccupiedFileStorageSizeGet) | **Get** /api/v2/attachments/occupiedFileStorageSize | Get size of attachments storage in bytes
[**ApiV2AttachmentsPost**](AttachmentsApi.md#ApiV2AttachmentsPost) | **Post** /api/v2/attachments | Upload new attachment file



## ApiV2AttachmentsIdDelete

> ApiV2AttachmentsIdDelete(ctx, id).Execute()

Delete attachment file

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
    r, err := apiClient.AttachmentsApi.ApiV2AttachmentsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.ApiV2AttachmentsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AttachmentsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AttachmentsIdGet

> *os.File ApiV2AttachmentsIdGet(ctx, id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()

Download attachment file

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
    width := int32(56) // int32 | Width of the result image (optional)
    height := int32(56) // int32 | Height of the result image (optional)
    resizeType := openapiclient.ImageResizeType("Crop") // ImageResizeType | Type of resizing to apply to the result image (optional)
    backgroundColor := "backgroundColor_example" // string | Color of the background if the `resizeType` is `AddBackgroundStripes` (optional)
    preview := true // bool | If image must be converted to a preview (lower quality, no animation) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentsApi.ApiV2AttachmentsIdGet(context.Background(), id).Width(width).Height(height).ResizeType(resizeType).BackgroundColor(backgroundColor).Preview(preview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.ApiV2AttachmentsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AttachmentsIdGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.ApiV2AttachmentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AttachmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **width** | **int32** | Width of the result image | 
 **height** | **int32** | Height of the result image | 
 **resizeType** | [**ImageResizeType**](ImageResizeType.md) | Type of resizing to apply to the result image | 
 **backgroundColor** | **string** | Color of the background if the &#x60;resizeType&#x60; is &#x60;AddBackgroundStripes&#x60; | 
 **preview** | **bool** | If image must be converted to a preview (lower quality, no animation) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AttachmentsOccupiedFileStorageSizeGet

> int64 ApiV2AttachmentsOccupiedFileStorageSizeGet(ctx).Execute()

Get size of attachments storage in bytes

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentsApi.ApiV2AttachmentsOccupiedFileStorageSizeGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.ApiV2AttachmentsOccupiedFileStorageSizeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AttachmentsOccupiedFileStorageSizeGet`: int64
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.ApiV2AttachmentsOccupiedFileStorageSizeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AttachmentsOccupiedFileStorageSizeGetRequest struct via the builder pattern


### Return type

**int64**

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2AttachmentsPost

> AttachmentModel ApiV2AttachmentsPost(ctx).File(file).Execute()

Upload new attachment file



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
    resp, r, err := apiClient.AttachmentsApi.ApiV2AttachmentsPost(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsApi.ApiV2AttachmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2AttachmentsPost`: AttachmentModel
    fmt.Fprintf(os.Stdout, "Response from `AttachmentsApi.ApiV2AttachmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2AttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**AttachmentModel**](AttachmentModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

