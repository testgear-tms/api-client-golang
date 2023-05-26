# \ConfigurationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ConfigurationsCreateByParametersPost**](ConfigurationsApi.md#ApiV2ConfigurationsCreateByParametersPost) | **Post** /api/v2/configurations/createByParameters | Create Configurations by parameters
[**ApiV2ConfigurationsIdPatch**](ConfigurationsApi.md#ApiV2ConfigurationsIdPatch) | **Patch** /api/v2/configurations/{id} | Patch configuration
[**ApiV2ConfigurationsSearchPost**](ConfigurationsApi.md#ApiV2ConfigurationsSearchPost) | **Post** /api/v2/configurations/search | Search for configurations
[**CreateConfiguration**](ConfigurationsApi.md#CreateConfiguration) | **Post** /api/v2/configurations | Create Configuration
[**GetConfigurationById**](ConfigurationsApi.md#GetConfigurationById) | **Get** /api/v2/configurations/{id} | Get configuration by internal or global ID
[**UpdateConfiguration**](ConfigurationsApi.md#UpdateConfiguration) | **Put** /api/v2/configurations | Update Configuration



## ApiV2ConfigurationsCreateByParametersPost

> ApiV2ConfigurationsCreateByParametersPost(ctx).ConfigurationByParametersModel(configurationByParametersModel).Execute()

Create Configurations by parameters



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
    configurationByParametersModel := *openapiclient.NewConfigurationByParametersModel([]string{"ParameterIds_example"}) // ConfigurationByParametersModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigurationsApi.ApiV2ConfigurationsCreateByParametersPost(context.Background()).ConfigurationByParametersModel(configurationByParametersModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.ApiV2ConfigurationsCreateByParametersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsCreateByParametersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationByParametersModel** | [**ConfigurationByParametersModel**](ConfigurationByParametersModel.md) |  | 

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


## ApiV2ConfigurationsIdPatch

> ApiV2ConfigurationsIdPatch(ctx, id).Operation(operation).Execute()

Patch configuration



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the configuration
    operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigurationsApi.ApiV2ConfigurationsIdPatch(context.Background(), id).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.ApiV2ConfigurationsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique ID of the configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsIdPatchRequest struct via the builder pattern


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


## ApiV2ConfigurationsSearchPost

> []ConfigurationModel ApiV2ConfigurationsSearchPost(ctx).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ConfigurationSelectModel(configurationSelectModel).Execute()

Search for configurations

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
    configurationSelectModel := *openapiclient.NewConfigurationSelectModel() // ConfigurationSelectModel | Model containing all the filters (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationsApi.ApiV2ConfigurationsSearchPost(context.Background()).Skip(skip).Take(take).OrderBy(orderBy).SearchField(searchField).SearchValue(searchValue).ConfigurationSelectModel(configurationSelectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.ApiV2ConfigurationsSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ConfigurationsSearchPost`: []ConfigurationModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.ApiV2ConfigurationsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ConfigurationsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Amount of items to be skipped (offset) | 
 **take** | **int32** | Amount of items to be taken (limit) | 
 **orderBy** | **string** | SQL-like  ORDER BY statement (column1 ASC|DESC , column2 ASC|DESC) | 
 **searchField** | **string** | Property name for searching | 
 **searchValue** | **string** | Value for searching | 
 **configurationSelectModel** | [**ConfigurationSelectModel**](ConfigurationSelectModel.md) | Model containing all the filters | 

### Return type

[**[]ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfiguration

> ConfigurationModel CreateConfiguration(ctx).ConfigurationPostModel(configurationPostModel).Execute()

Create Configuration



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
    configurationPostModel := *openapiclient.NewConfigurationPostModel("ProjectId_example", "Default") // ConfigurationPostModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationsApi.CreateConfiguration(context.Background()).ConfigurationPostModel(configurationPostModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.CreateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfiguration`: ConfigurationModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.CreateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationPostModel** | [**ConfigurationPostModel**](ConfigurationPostModel.md) |  | 

### Return type

[**ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationById

> ConfigurationModel GetConfigurationById(ctx, id).Execute()

Get configuration by internal or global ID



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
    id := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Configuration internal (guid format) or global (integer format) identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationsApi.GetConfigurationById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.GetConfigurationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationById`: ConfigurationModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationsApi.GetConfigurationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Configuration internal (guid format) or global (integer format) identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationModel**](ConfigurationModel.md)

### Authorization

[Bearer or PrivateToken](../README.md#Bearer or PrivateToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> UpdateConfiguration(ctx).ConfigurationPutModel(configurationPutModel).Execute()

Update Configuration



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
    configurationPutModel := *openapiclient.NewConfigurationPutModel("d49af44b-dbd8-48b0-90e5-e065735d7229", "ProjectId_example", "Default") // ConfigurationPutModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigurationsApi.UpdateConfiguration(context.Background()).ConfigurationPutModel(configurationPutModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationPutModel** | [**ConfigurationPutModel**](ConfigurationPutModel.md) |  | 

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

