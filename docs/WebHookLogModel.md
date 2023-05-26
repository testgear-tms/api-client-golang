# WebHookLogModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebHookName** | Pointer to **NullableString** |  | [optional] 
**EventType** | [**WebHookEventTypeModel**](WebHookEventTypeModel.md) |  | 
**WebHookId** | Pointer to **string** |  | [optional] 
**RequestBody** | Pointer to **NullableString** |  | [optional] 
**RequestMeta** | Pointer to **NullableString** |  | [optional] 
**ResponseStatusCode** | Pointer to **int32** |  | [optional] 
**ResponseBody** | Pointer to **NullableString** |  | [optional] 
**ResponseMeta** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**RequestType** | [**RequestTypeModel**](RequestTypeModel.md) |  | 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewWebHookLogModel

`func NewWebHookLogModel(eventType WebHookEventTypeModel, requestType RequestTypeModel, ) *WebHookLogModel`

NewWebHookLogModel instantiates a new WebHookLogModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookLogModelWithDefaults

`func NewWebHookLogModelWithDefaults() *WebHookLogModel`

NewWebHookLogModelWithDefaults instantiates a new WebHookLogModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebHookName

`func (o *WebHookLogModel) GetWebHookName() string`

GetWebHookName returns the WebHookName field if non-nil, zero value otherwise.

### GetWebHookNameOk

`func (o *WebHookLogModel) GetWebHookNameOk() (*string, bool)`

GetWebHookNameOk returns a tuple with the WebHookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookName

`func (o *WebHookLogModel) SetWebHookName(v string)`

SetWebHookName sets WebHookName field to given value.

### HasWebHookName

`func (o *WebHookLogModel) HasWebHookName() bool`

HasWebHookName returns a boolean if a field has been set.

### SetWebHookNameNil

`func (o *WebHookLogModel) SetWebHookNameNil(b bool)`

 SetWebHookNameNil sets the value for WebHookName to be an explicit nil

### UnsetWebHookName
`func (o *WebHookLogModel) UnsetWebHookName()`

UnsetWebHookName ensures that no value is present for WebHookName, not even an explicit nil
### GetEventType

`func (o *WebHookLogModel) GetEventType() WebHookEventTypeModel`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebHookLogModel) GetEventTypeOk() (*WebHookEventTypeModel, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebHookLogModel) SetEventType(v WebHookEventTypeModel)`

SetEventType sets EventType field to given value.


### GetWebHookId

`func (o *WebHookLogModel) GetWebHookId() string`

GetWebHookId returns the WebHookId field if non-nil, zero value otherwise.

### GetWebHookIdOk

`func (o *WebHookLogModel) GetWebHookIdOk() (*string, bool)`

GetWebHookIdOk returns a tuple with the WebHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookId

`func (o *WebHookLogModel) SetWebHookId(v string)`

SetWebHookId sets WebHookId field to given value.

### HasWebHookId

`func (o *WebHookLogModel) HasWebHookId() bool`

HasWebHookId returns a boolean if a field has been set.

### GetRequestBody

`func (o *WebHookLogModel) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *WebHookLogModel) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *WebHookLogModel) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *WebHookLogModel) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### SetRequestBodyNil

`func (o *WebHookLogModel) SetRequestBodyNil(b bool)`

 SetRequestBodyNil sets the value for RequestBody to be an explicit nil

### UnsetRequestBody
`func (o *WebHookLogModel) UnsetRequestBody()`

UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
### GetRequestMeta

`func (o *WebHookLogModel) GetRequestMeta() string`

GetRequestMeta returns the RequestMeta field if non-nil, zero value otherwise.

### GetRequestMetaOk

`func (o *WebHookLogModel) GetRequestMetaOk() (*string, bool)`

GetRequestMetaOk returns a tuple with the RequestMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMeta

`func (o *WebHookLogModel) SetRequestMeta(v string)`

SetRequestMeta sets RequestMeta field to given value.

### HasRequestMeta

`func (o *WebHookLogModel) HasRequestMeta() bool`

HasRequestMeta returns a boolean if a field has been set.

### SetRequestMetaNil

`func (o *WebHookLogModel) SetRequestMetaNil(b bool)`

 SetRequestMetaNil sets the value for RequestMeta to be an explicit nil

### UnsetRequestMeta
`func (o *WebHookLogModel) UnsetRequestMeta()`

UnsetRequestMeta ensures that no value is present for RequestMeta, not even an explicit nil
### GetResponseStatusCode

`func (o *WebHookLogModel) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *WebHookLogModel) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *WebHookLogModel) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *WebHookLogModel) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetResponseBody

`func (o *WebHookLogModel) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *WebHookLogModel) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *WebHookLogModel) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *WebHookLogModel) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *WebHookLogModel) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *WebHookLogModel) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseMeta

`func (o *WebHookLogModel) GetResponseMeta() string`

GetResponseMeta returns the ResponseMeta field if non-nil, zero value otherwise.

### GetResponseMetaOk

`func (o *WebHookLogModel) GetResponseMetaOk() (*string, bool)`

GetResponseMetaOk returns a tuple with the ResponseMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMeta

`func (o *WebHookLogModel) SetResponseMeta(v string)`

SetResponseMeta sets ResponseMeta field to given value.

### HasResponseMeta

`func (o *WebHookLogModel) HasResponseMeta() bool`

HasResponseMeta returns a boolean if a field has been set.

### SetResponseMetaNil

`func (o *WebHookLogModel) SetResponseMetaNil(b bool)`

 SetResponseMetaNil sets the value for ResponseMeta to be an explicit nil

### UnsetResponseMeta
`func (o *WebHookLogModel) UnsetResponseMeta()`

UnsetResponseMeta ensures that no value is present for ResponseMeta, not even an explicit nil
### GetProjectId

`func (o *WebHookLogModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebHookLogModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebHookLogModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WebHookLogModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetUrl

`func (o *WebHookLogModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookLogModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookLogModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebHookLogModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WebHookLogModel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebHookLogModel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRequestType

`func (o *WebHookLogModel) GetRequestType() RequestTypeModel`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WebHookLogModel) GetRequestTypeOk() (*RequestTypeModel, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WebHookLogModel) SetRequestType(v RequestTypeModel)`

SetRequestType sets RequestType field to given value.


### GetCreatedDate

`func (o *WebHookLogModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WebHookLogModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WebHookLogModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WebHookLogModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *WebHookLogModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *WebHookLogModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *WebHookLogModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WebHookLogModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WebHookLogModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WebHookLogModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WebHookLogModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WebHookLogModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *WebHookLogModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WebHookLogModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WebHookLogModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *WebHookLogModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *WebHookLogModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WebHookLogModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WebHookLogModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WebHookLogModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WebHookLogModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WebHookLogModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetId

`func (o *WebHookLogModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebHookLogModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebHookLogModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebHookLogModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *WebHookLogModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WebHookLogModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WebHookLogModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *WebHookLogModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


