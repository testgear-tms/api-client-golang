# WebHookModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the webhook | [optional] 
**EventType** | [**WebHookEventTypeModel**](WebHookEventTypeModel.md) |  | 
**Description** | Pointer to **NullableString** | Description of the webhook | [optional] 
**Url** | Pointer to **NullableString** | Url to which the webhook sends request | [optional] 
**RequestType** | [**RequestTypeModel**](RequestTypeModel.md) |  | 
**ShouldSendBody** | Pointer to **bool** | Indicates if the webhook sends body | [optional] 
**Headers** | Pointer to **map[string]string** | Collection of headers which the webhook sends | [optional] 
**QueryParameters** | Pointer to **map[string]string** | Collection of query parameters which the webhook sends | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates if the webhook is active | [optional] 
**ShouldSendCustomBody** | Pointer to **bool** | Indicates if the webhook sends custom body | [optional] 
**CustomBody** | Pointer to **NullableString** | Custom body of the webhook | [optional] 
**CustomBodyMediaType** | Pointer to **NullableString** | MIME type of body of the webhook | [optional] 
**ShouldReplaceParameters** | Pointer to **bool** | Indicates if the webhook injects parameters | [optional] 
**ShouldEscapeParameters** | Pointer to **bool** | Indicates if the webhook escapes invalid characters in parameters | [optional] 
**CreatedDate** | Pointer to **time.Time** | Creation date of the webhook | [optional] 
**CreatedById** | Pointer to **string** | Unique ID of user who created the webhook | [optional] 
**ModifiedDate** | Pointer to **NullableTime** | Last modification date of the webhook | [optional] 
**ModifiedById** | Pointer to **NullableString** | Unique ID of user who modified the webhook last time | [optional] 
**ProjectId** | Pointer to **string** | Unique ID of project where the webhook is located | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewWebHookModel

`func NewWebHookModel(eventType WebHookEventTypeModel, requestType RequestTypeModel, ) *WebHookModel`

NewWebHookModel instantiates a new WebHookModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookModelWithDefaults

`func NewWebHookModelWithDefaults() *WebHookModel`

NewWebHookModelWithDefaults instantiates a new WebHookModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebHookModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebHookModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebHookModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebHookModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebHookModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebHookModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEventType

`func (o *WebHookModel) GetEventType() WebHookEventTypeModel`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebHookModel) GetEventTypeOk() (*WebHookEventTypeModel, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebHookModel) SetEventType(v WebHookEventTypeModel)`

SetEventType sets EventType field to given value.


### GetDescription

`func (o *WebHookModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebHookModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebHookModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebHookModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebHookModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebHookModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *WebHookModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebHookModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WebHookModel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebHookModel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetRequestType

`func (o *WebHookModel) GetRequestType() RequestTypeModel`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WebHookModel) GetRequestTypeOk() (*RequestTypeModel, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WebHookModel) SetRequestType(v RequestTypeModel)`

SetRequestType sets RequestType field to given value.


### GetShouldSendBody

`func (o *WebHookModel) GetShouldSendBody() bool`

GetShouldSendBody returns the ShouldSendBody field if non-nil, zero value otherwise.

### GetShouldSendBodyOk

`func (o *WebHookModel) GetShouldSendBodyOk() (*bool, bool)`

GetShouldSendBodyOk returns a tuple with the ShouldSendBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendBody

`func (o *WebHookModel) SetShouldSendBody(v bool)`

SetShouldSendBody sets ShouldSendBody field to given value.

### HasShouldSendBody

`func (o *WebHookModel) HasShouldSendBody() bool`

HasShouldSendBody returns a boolean if a field has been set.

### GetHeaders

`func (o *WebHookModel) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebHookModel) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebHookModel) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebHookModel) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *WebHookModel) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *WebHookModel) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetQueryParameters

`func (o *WebHookModel) GetQueryParameters() map[string]string`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *WebHookModel) GetQueryParametersOk() (*map[string]string, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *WebHookModel) SetQueryParameters(v map[string]string)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *WebHookModel) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### SetQueryParametersNil

`func (o *WebHookModel) SetQueryParametersNil(b bool)`

 SetQueryParametersNil sets the value for QueryParameters to be an explicit nil

### UnsetQueryParameters
`func (o *WebHookModel) UnsetQueryParameters()`

UnsetQueryParameters ensures that no value is present for QueryParameters, not even an explicit nil
### GetIsEnabled

`func (o *WebHookModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WebHookModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WebHookModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *WebHookModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetShouldSendCustomBody

`func (o *WebHookModel) GetShouldSendCustomBody() bool`

GetShouldSendCustomBody returns the ShouldSendCustomBody field if non-nil, zero value otherwise.

### GetShouldSendCustomBodyOk

`func (o *WebHookModel) GetShouldSendCustomBodyOk() (*bool, bool)`

GetShouldSendCustomBodyOk returns a tuple with the ShouldSendCustomBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendCustomBody

`func (o *WebHookModel) SetShouldSendCustomBody(v bool)`

SetShouldSendCustomBody sets ShouldSendCustomBody field to given value.

### HasShouldSendCustomBody

`func (o *WebHookModel) HasShouldSendCustomBody() bool`

HasShouldSendCustomBody returns a boolean if a field has been set.

### GetCustomBody

`func (o *WebHookModel) GetCustomBody() string`

GetCustomBody returns the CustomBody field if non-nil, zero value otherwise.

### GetCustomBodyOk

`func (o *WebHookModel) GetCustomBodyOk() (*string, bool)`

GetCustomBodyOk returns a tuple with the CustomBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBody

`func (o *WebHookModel) SetCustomBody(v string)`

SetCustomBody sets CustomBody field to given value.

### HasCustomBody

`func (o *WebHookModel) HasCustomBody() bool`

HasCustomBody returns a boolean if a field has been set.

### SetCustomBodyNil

`func (o *WebHookModel) SetCustomBodyNil(b bool)`

 SetCustomBodyNil sets the value for CustomBody to be an explicit nil

### UnsetCustomBody
`func (o *WebHookModel) UnsetCustomBody()`

UnsetCustomBody ensures that no value is present for CustomBody, not even an explicit nil
### GetCustomBodyMediaType

`func (o *WebHookModel) GetCustomBodyMediaType() string`

GetCustomBodyMediaType returns the CustomBodyMediaType field if non-nil, zero value otherwise.

### GetCustomBodyMediaTypeOk

`func (o *WebHookModel) GetCustomBodyMediaTypeOk() (*string, bool)`

GetCustomBodyMediaTypeOk returns a tuple with the CustomBodyMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBodyMediaType

`func (o *WebHookModel) SetCustomBodyMediaType(v string)`

SetCustomBodyMediaType sets CustomBodyMediaType field to given value.

### HasCustomBodyMediaType

`func (o *WebHookModel) HasCustomBodyMediaType() bool`

HasCustomBodyMediaType returns a boolean if a field has been set.

### SetCustomBodyMediaTypeNil

`func (o *WebHookModel) SetCustomBodyMediaTypeNil(b bool)`

 SetCustomBodyMediaTypeNil sets the value for CustomBodyMediaType to be an explicit nil

### UnsetCustomBodyMediaType
`func (o *WebHookModel) UnsetCustomBodyMediaType()`

UnsetCustomBodyMediaType ensures that no value is present for CustomBodyMediaType, not even an explicit nil
### GetShouldReplaceParameters

`func (o *WebHookModel) GetShouldReplaceParameters() bool`

GetShouldReplaceParameters returns the ShouldReplaceParameters field if non-nil, zero value otherwise.

### GetShouldReplaceParametersOk

`func (o *WebHookModel) GetShouldReplaceParametersOk() (*bool, bool)`

GetShouldReplaceParametersOk returns a tuple with the ShouldReplaceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReplaceParameters

`func (o *WebHookModel) SetShouldReplaceParameters(v bool)`

SetShouldReplaceParameters sets ShouldReplaceParameters field to given value.

### HasShouldReplaceParameters

`func (o *WebHookModel) HasShouldReplaceParameters() bool`

HasShouldReplaceParameters returns a boolean if a field has been set.

### GetShouldEscapeParameters

`func (o *WebHookModel) GetShouldEscapeParameters() bool`

GetShouldEscapeParameters returns the ShouldEscapeParameters field if non-nil, zero value otherwise.

### GetShouldEscapeParametersOk

`func (o *WebHookModel) GetShouldEscapeParametersOk() (*bool, bool)`

GetShouldEscapeParametersOk returns a tuple with the ShouldEscapeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldEscapeParameters

`func (o *WebHookModel) SetShouldEscapeParameters(v bool)`

SetShouldEscapeParameters sets ShouldEscapeParameters field to given value.

### HasShouldEscapeParameters

`func (o *WebHookModel) HasShouldEscapeParameters() bool`

HasShouldEscapeParameters returns a boolean if a field has been set.

### GetCreatedDate

`func (o *WebHookModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WebHookModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WebHookModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WebHookModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedById

`func (o *WebHookModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WebHookModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WebHookModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *WebHookModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedDate

`func (o *WebHookModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WebHookModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WebHookModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WebHookModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WebHookModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WebHookModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *WebHookModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WebHookModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WebHookModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WebHookModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WebHookModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WebHookModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetProjectId

`func (o *WebHookModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebHookModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebHookModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WebHookModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetId

`func (o *WebHookModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebHookModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebHookModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebHookModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *WebHookModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WebHookModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WebHookModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *WebHookModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


