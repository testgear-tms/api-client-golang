# \CustomAttributesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2CustomAttributesGlobalIdDelete**](CustomAttributesApi.md#ApiV2CustomAttributesGlobalIdDelete) | **Delete** /api/v2/customAttributes/global/{id} | Delete global attribute
[**ApiV2CustomAttributesGlobalIdPut**](CustomAttributesApi.md#ApiV2CustomAttributesGlobalIdPut) | **Put** /api/v2/customAttributes/global/{id} | Edit global attribute
[**ApiV2CustomAttributesGlobalPost**](CustomAttributesApi.md#ApiV2CustomAttributesGlobalPost) | **Post** /api/v2/customAttributes/global | Create global attribute
[**ApiV2CustomAttributesIdGet**](CustomAttributesApi.md#ApiV2CustomAttributesIdGet) | **Get** /api/v2/customAttributes/{id} | Get attribute
[**ApiV2CustomAttributesSearchPost**](CustomAttributesApi.md#ApiV2CustomAttributesSearchPost) | **Post** /api/v2/customAttributes/search | Search for attributes



## ApiV2CustomAttributesGlobalIdDelete

> ApiV2CustomAttributesGlobalIdDelete(ctx, id).Execute()

Delete global attribute

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomAttributesApi.ApiV2CustomAttributesGlobalIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributesApi.ApiV2CustomAttributesGlobalIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesGlobalIdDeleteRequest struct via the builder pattern


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


## ApiV2CustomAttributesGlobalIdPut

> CustomAttributeModel ApiV2CustomAttributesGlobalIdPut(ctx, id).GlobalCustomAttributeUpdateModel(globalCustomAttributeUpdateModel).Execute()

Edit global attribute

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of attribute
    globalCustomAttributeUpdateModel := *openapiclient.NewGlobalCustomAttributeUpdateModel("Name_example") // GlobalCustomAttributeUpdateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributesApi.ApiV2CustomAttributesGlobalIdPut(context.Background(), id).GlobalCustomAttributeUpdateModel(globalCustomAttributeUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributesApi.ApiV2CustomAttributesGlobalIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesGlobalIdPut`: CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributesApi.ApiV2CustomAttributesGlobalIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesGlobalIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalCustomAttributeUpdateModel** | [**GlobalCustomAttributeUpdateModel**](GlobalCustomAttributeUpdateModel.md) |  | 

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


## ApiV2CustomAttributesGlobalPost

> CustomAttributeModel ApiV2CustomAttributesGlobalPost(ctx).GlobalCustomAttributePostModel(globalCustomAttributePostModel).Execute()

Create global attribute

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
    globalCustomAttributePostModel := *openapiclient.NewGlobalCustomAttributePostModel("Name_example", openapiclient.CustomAttributeTypesEnum("string")) // GlobalCustomAttributePostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributesApi.ApiV2CustomAttributesGlobalPost(context.Background()).GlobalCustomAttributePostModel(globalCustomAttributePostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributesApi.ApiV2CustomAttributesGlobalPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesGlobalPost`: CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributesApi.ApiV2CustomAttributesGlobalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesGlobalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalCustomAttributePostModel** | [**GlobalCustomAttributePostModel**](GlobalCustomAttributePostModel.md) |  | 

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


## ApiV2CustomAttributesIdGet

> CustomAttributeModel ApiV2CustomAttributesIdGet(ctx, id).Execute()

Get attribute

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributesApi.ApiV2CustomAttributesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributesApi.ApiV2CustomAttributesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesIdGet`: CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributesApi.ApiV2CustomAttributesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesIdGetRequest struct via the builder pattern


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


## ApiV2CustomAttributesSearchPost

> []CustomAttributeModel ApiV2CustomAttributesSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).CustomAttributeSearchQueryModel(customAttributeSearchQueryModel).Execute()

Search for attributes

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
    customAttributeSearchQueryModel := *openapiclient.NewCustomAttributeSearchQueryModel() // CustomAttributeSearchQueryModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributesApi.ApiV2CustomAttributesSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).CustomAttributeSearchQueryModel(customAttributeSearchQueryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributesApi.ApiV2CustomAttributesSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesSearchPost`: []CustomAttributeModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributesApi.ApiV2CustomAttributesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **customAttributeSearchQueryModel** | [**CustomAttributeSearchQueryModel**](CustomAttributeSearchQueryModel.md) |  | 

### Return type

[**[]CustomAttributeModel**](CustomAttributeModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

