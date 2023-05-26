# \CustomAttributeTemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost) | **Post** /api/v2/customAttributes/templates/{id}/customAttributes/exclude | Exclude CustomAttributes from CustomAttributeTemplate
[**ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost) | **Post** /api/v2/customAttributes/templates/{id}/customAttributes/include | Include CustomAttributes to CustomAttributeTemplate
[**ApiV2CustomAttributesTemplatesIdDelete**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesIdDelete) | **Delete** /api/v2/customAttributes/templates/{id} | Delete CustomAttributeTemplate
[**ApiV2CustomAttributesTemplatesIdGet**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesIdGet) | **Get** /api/v2/customAttributes/templates/{id} | Get CustomAttributeTemplate by ID
[**ApiV2CustomAttributesTemplatesNameGet**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesNameGet) | **Get** /api/v2/customAttributes/templates/{name} | Get CustomAttributeTemplate by name
[**ApiV2CustomAttributesTemplatesPost**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesPost) | **Post** /api/v2/customAttributes/templates | Create CustomAttributeTemplate
[**ApiV2CustomAttributesTemplatesPut**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesPut) | **Put** /api/v2/customAttributes/templates | Update custom attributes template
[**ApiV2CustomAttributesTemplatesSearchPost**](CustomAttributeTemplatesApi.md#ApiV2CustomAttributesTemplatesSearchPost) | **Post** /api/v2/customAttributes/templates/search | Search CustomAttributeTemplates



## ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost

> ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost(ctx, id).RequestBody(requestBody).Execute()

Exclude CustomAttributes from CustomAttributeTemplate



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Attribute template internal (UUID) identifier
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdCustomAttributesExcludePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Attribute template internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesIdCustomAttributesExcludePostRequest struct via the builder pattern


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


## ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost

> ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost(ctx, id).RequestBody(requestBody).Execute()

Include CustomAttributes to CustomAttributeTemplate



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Attribute template internal (UUID) identifier
    requestBody := []string{"Property_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdCustomAttributesIncludePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Attribute template internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesIdCustomAttributesIncludePostRequest struct via the builder pattern


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


## ApiV2CustomAttributesTemplatesIdDelete

> NoContentResult ApiV2CustomAttributesTemplatesIdDelete(ctx, id).Execute()

Delete CustomAttributeTemplate



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Attribute template internal (UUID) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesTemplatesIdDelete`: NoContentResult
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Attribute template internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NoContentResult**](NoContentResult.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2CustomAttributesTemplatesIdGet

> CustomAttributeTemplateModel ApiV2CustomAttributesTemplatesIdGet(ctx, id).Execute()

Get CustomAttributeTemplate by ID



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | CustomAttributeTemplate internal (UUID) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesTemplatesIdGet`: CustomAttributeTemplateModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CustomAttributeTemplate internal (UUID) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomAttributeTemplateModel**](CustomAttributeTemplateModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2CustomAttributesTemplatesNameGet

> CustomAttributeTemplateModel ApiV2CustomAttributesTemplatesNameGet(ctx, name).Execute()

Get CustomAttributeTemplate by name



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
    name := "name_example" // string | CustomAttributeTemplate name for search

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesTemplatesNameGet`: CustomAttributeTemplateModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | CustomAttributeTemplate name for search | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomAttributeTemplateModel**](CustomAttributeTemplateModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2CustomAttributesTemplatesPost

> CustomAttributeTemplateModel ApiV2CustomAttributesTemplatesPost(ctx).CustomAttributeTemplatePostModel(customAttributeTemplatePostModel).Execute()

Create CustomAttributeTemplate



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
    customAttributeTemplatePostModel := *openapiclient.NewCustomAttributeTemplatePostModel("Name_example") // CustomAttributeTemplatePostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesPost(context.Background()).CustomAttributeTemplatePostModel(customAttributeTemplatePostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesTemplatesPost`: CustomAttributeTemplateModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customAttributeTemplatePostModel** | [**CustomAttributeTemplatePostModel**](CustomAttributeTemplatePostModel.md) |  | 

### Return type

[**CustomAttributeTemplateModel**](CustomAttributeTemplateModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2CustomAttributesTemplatesPut

> ApiV2CustomAttributesTemplatesPut(ctx).CustomAttributeTemplatePutModel(customAttributeTemplatePutModel).Execute()

Update custom attributes template

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
    customAttributeTemplatePutModel := *openapiclient.NewCustomAttributeTemplatePutModel("Id_example", "Name_example") // CustomAttributeTemplatePutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesPut(context.Background()).CustomAttributeTemplatePutModel(customAttributeTemplatePutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customAttributeTemplatePutModel** | [**CustomAttributeTemplatePutModel**](CustomAttributeTemplatePutModel.md) |  | 

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


## ApiV2CustomAttributesTemplatesSearchPost

> []SearchCustomAttributeTemplateGetModel ApiV2CustomAttributesTemplatesSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).CustomAttributeTemplateSearchQueryModel(customAttributeTemplateSearchQueryModel).Execute()

Search CustomAttributeTemplates



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
    customAttributeTemplateSearchQueryModel := *openapiclient.NewCustomAttributeTemplateSearchQueryModel() // CustomAttributeTemplateSearchQueryModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).CustomAttributeTemplateSearchQueryModel(customAttributeTemplateSearchQueryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2CustomAttributesTemplatesSearchPost`: []SearchCustomAttributeTemplateGetModel
    fmt.Fprintf(os.Stdout, "Response from `CustomAttributeTemplatesApi.ApiV2CustomAttributesTemplatesSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2CustomAttributesTemplatesSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **customAttributeTemplateSearchQueryModel** | [**CustomAttributeTemplateSearchQueryModel**](CustomAttributeTemplateSearchQueryModel.md) |  | 

### Return type

[**[]SearchCustomAttributeTemplateGetModel**](SearchCustomAttributeTemplateGetModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

