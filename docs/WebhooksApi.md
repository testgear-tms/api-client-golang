# \WebhooksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2WebhooksGet**](WebhooksApi.md#ApiV2WebhooksGet) | **Get** /api/v2/webhooks | Get all webhooks
[**ApiV2WebhooksIdDelete**](WebhooksApi.md#ApiV2WebhooksIdDelete) | **Delete** /api/v2/webhooks/{id} | Delete webhook by ID
[**ApiV2WebhooksIdGet**](WebhooksApi.md#ApiV2WebhooksIdGet) | **Get** /api/v2/webhooks/{id} | Get webhook by ID
[**ApiV2WebhooksIdPut**](WebhooksApi.md#ApiV2WebhooksIdPut) | **Put** /api/v2/webhooks/{id} | Edit webhook by ID
[**ApiV2WebhooksPost**](WebhooksApi.md#ApiV2WebhooksPost) | **Post** /api/v2/webhooks | Create webhook
[**ApiV2WebhooksSearchPost**](WebhooksApi.md#ApiV2WebhooksSearchPost) | **Post** /api/v2/webhooks/search | Search for webhooks
[**ApiV2WebhooksSpecialVariablesGet**](WebhooksApi.md#ApiV2WebhooksSpecialVariablesGet) | **Get** /api/v2/webhooks/specialVariables | Get special variables for webhook event type



## ApiV2WebhooksGet

> []WebHookModel ApiV2WebhooksGet(ctx).ProjectId(projectId).Execute()

Get all webhooks

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project unique ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ApiV2WebhooksGet(context.Background()).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ApiV2WebhooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WebhooksGet`: []WebHookModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ApiV2WebhooksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | Project unique ID | 

### Return type

[**[]WebHookModel**](WebHookModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WebhooksIdDelete

> ApiV2WebhooksIdDelete(ctx, id).Execute()

Delete webhook by ID

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.ApiV2WebhooksIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ApiV2WebhooksIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WebhooksIdDeleteRequest struct via the builder pattern


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


## ApiV2WebhooksIdGet

> WebHookModel ApiV2WebhooksIdGet(ctx, id).Execute()

Get webhook by ID

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook unique ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ApiV2WebhooksIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ApiV2WebhooksIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WebhooksIdGet`: WebHookModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ApiV2WebhooksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WebhooksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebHookModel**](WebHookModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WebhooksIdPut

> WebHookModel ApiV2WebhooksIdPut(ctx, id).WebHookPostModel(webHookPostModel).Execute()

Edit webhook by ID

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook unique ID
    webHookPostModel := *openapiclient.NewWebHookPostModel("ProjectId_example", openapiclient.WebHookEventTypeModel("AutomatedTestRunCreated"), "Url_example", openapiclient.RequestTypeModel("Post"), "Name_example") // WebHookPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ApiV2WebhooksIdPut(context.Background(), id).WebHookPostModel(webHookPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ApiV2WebhooksIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WebhooksIdPut`: WebHookModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ApiV2WebhooksIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook unique ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WebhooksIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webHookPostModel** | [**WebHookPostModel**](WebHookPostModel.md) |  | 

### Return type

[**WebHookModel**](WebHookModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WebhooksPost

> WebHookModel ApiV2WebhooksPost(ctx).WebHookPostModel(webHookPostModel).Execute()

Create webhook

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
    webHookPostModel := *openapiclient.NewWebHookPostModel("ProjectId_example", openapiclient.WebHookEventTypeModel("AutomatedTestRunCreated"), "Url_example", openapiclient.RequestTypeModel("Post"), "Name_example") // WebHookPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ApiV2WebhooksPost(context.Background()).WebHookPostModel(webHookPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ApiV2WebhooksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WebhooksPost`: WebHookModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ApiV2WebhooksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WebhooksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webHookPostModel** | [**WebHookPostModel**](WebHookPostModel.md) |  | 

### Return type

[**WebHookModel**](WebHookModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WebhooksSearchPost

> []WebHookModel ApiV2WebhooksSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SearchWebhooksQueryModel(searchWebhooksQueryModel).Execute()

Search for webhooks

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
    searchWebhooksQueryModel := *openapiclient.NewSearchWebhooksQueryModel() // SearchWebhooksQueryModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ApiV2WebhooksSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).SearchWebhooksQueryModel(searchWebhooksQueryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ApiV2WebhooksSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WebhooksSearchPost`: []WebHookModel
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ApiV2WebhooksSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WebhooksSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **searchWebhooksQueryModel** | [**SearchWebhooksQueryModel**](SearchWebhooksQueryModel.md) |  | 

### Return type

[**[]WebHookModel**](WebHookModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2WebhooksSpecialVariablesGet

> []string ApiV2WebhooksSpecialVariablesGet(ctx).EventType(eventType).Execute()

Get special variables for webhook event type

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
    eventType := openapiclient.WebHookEventType("AutomatedTestRunCreated") // WebHookEventType | Webhook event type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ApiV2WebhooksSpecialVariablesGet(context.Background()).EventType(eventType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ApiV2WebhooksSpecialVariablesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2WebhooksSpecialVariablesGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ApiV2WebhooksSpecialVariablesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2WebhooksSpecialVariablesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventType** | [**WebHookEventType**](WebHookEventType.md) | Webhook event type | 

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

