# TestPlanLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BugLink** | Pointer to [**LinkModel**](LinkModel.md) |  | [optional] 
**WorkItemGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**WorkItemName** | Pointer to **NullableString** |  | [optional] 
**ConfigurationName** | Pointer to **NullableString** |  | [optional] 
**CreatedById** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to [**ExternalLinkModel**](ExternalLinkModel.md) |  | [optional] 

## Methods

### NewTestPlanLink

`func NewTestPlanLink() *TestPlanLink`

NewTestPlanLink instantiates a new TestPlanLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanLinkWithDefaults

`func NewTestPlanLinkWithDefaults() *TestPlanLink`

NewTestPlanLinkWithDefaults instantiates a new TestPlanLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBugLink

`func (o *TestPlanLink) GetBugLink() LinkModel`

GetBugLink returns the BugLink field if non-nil, zero value otherwise.

### GetBugLinkOk

`func (o *TestPlanLink) GetBugLinkOk() (*LinkModel, bool)`

GetBugLinkOk returns a tuple with the BugLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugLink

`func (o *TestPlanLink) SetBugLink(v LinkModel)`

SetBugLink sets BugLink field to given value.

### HasBugLink

`func (o *TestPlanLink) HasBugLink() bool`

HasBugLink returns a boolean if a field has been set.

### GetWorkItemGlobalId

`func (o *TestPlanLink) GetWorkItemGlobalId() int64`

GetWorkItemGlobalId returns the WorkItemGlobalId field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdOk

`func (o *TestPlanLink) GetWorkItemGlobalIdOk() (*int64, bool)`

GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalId

`func (o *TestPlanLink) SetWorkItemGlobalId(v int64)`

SetWorkItemGlobalId sets WorkItemGlobalId field to given value.

### HasWorkItemGlobalId

`func (o *TestPlanLink) HasWorkItemGlobalId() bool`

HasWorkItemGlobalId returns a boolean if a field has been set.

### SetWorkItemGlobalIdNil

`func (o *TestPlanLink) SetWorkItemGlobalIdNil(b bool)`

 SetWorkItemGlobalIdNil sets the value for WorkItemGlobalId to be an explicit nil

### UnsetWorkItemGlobalId
`func (o *TestPlanLink) UnsetWorkItemGlobalId()`

UnsetWorkItemGlobalId ensures that no value is present for WorkItemGlobalId, not even an explicit nil
### GetWorkItemName

`func (o *TestPlanLink) GetWorkItemName() string`

GetWorkItemName returns the WorkItemName field if non-nil, zero value otherwise.

### GetWorkItemNameOk

`func (o *TestPlanLink) GetWorkItemNameOk() (*string, bool)`

GetWorkItemNameOk returns a tuple with the WorkItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemName

`func (o *TestPlanLink) SetWorkItemName(v string)`

SetWorkItemName sets WorkItemName field to given value.

### HasWorkItemName

`func (o *TestPlanLink) HasWorkItemName() bool`

HasWorkItemName returns a boolean if a field has been set.

### SetWorkItemNameNil

`func (o *TestPlanLink) SetWorkItemNameNil(b bool)`

 SetWorkItemNameNil sets the value for WorkItemName to be an explicit nil

### UnsetWorkItemName
`func (o *TestPlanLink) UnsetWorkItemName()`

UnsetWorkItemName ensures that no value is present for WorkItemName, not even an explicit nil
### GetConfigurationName

`func (o *TestPlanLink) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestPlanLink) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestPlanLink) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *TestPlanLink) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### SetConfigurationNameNil

`func (o *TestPlanLink) SetConfigurationNameNil(b bool)`

 SetConfigurationNameNil sets the value for ConfigurationName to be an explicit nil

### UnsetConfigurationName
`func (o *TestPlanLink) UnsetConfigurationName()`

UnsetConfigurationName ensures that no value is present for ConfigurationName, not even an explicit nil
### GetCreatedById

`func (o *TestPlanLink) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPlanLink) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPlanLink) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestPlanLink) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### SetCreatedByIdNil

`func (o *TestPlanLink) SetCreatedByIdNil(b bool)`

 SetCreatedByIdNil sets the value for CreatedById to be an explicit nil

### UnsetCreatedById
`func (o *TestPlanLink) UnsetCreatedById()`

UnsetCreatedById ensures that no value is present for CreatedById, not even an explicit nil
### GetComment

`func (o *TestPlanLink) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestPlanLink) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestPlanLink) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestPlanLink) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestPlanLink) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestPlanLink) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetInfo

`func (o *TestPlanLink) GetInfo() ExternalLinkModel`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TestPlanLink) GetInfoOk() (*ExternalLinkModel, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TestPlanLink) SetInfo(v ExternalLinkModel)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TestPlanLink) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


