# NotificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**IsRead** | Pointer to **bool** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**NotificationType** | [**NotificationTypeModel**](NotificationTypeModel.md) |  | 
**ProjectGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**TestPlanGlobalId** | Pointer to **int64** |  | [optional] 
**TestPlanName** | Pointer to **NullableString** |  | [optional] 
**WorkitemGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**WorkItemName** | Pointer to **NullableString** |  | [optional] 
**AttributeName** | Pointer to **NullableString** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationModel

`func NewNotificationModel(notificationType NotificationTypeModel, ) *NotificationModel`

NewNotificationModel instantiates a new NotificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModelWithDefaults

`func NewNotificationModelWithDefaults() *NotificationModel`

NewNotificationModelWithDefaults instantiates a new NotificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *NotificationModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *NotificationModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *NotificationModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *NotificationModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *NotificationModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *NotificationModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetIsRead

`func (o *NotificationModel) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *NotificationModel) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *NotificationModel) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *NotificationModel) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetEntityId

`func (o *NotificationModel) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *NotificationModel) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *NotificationModel) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *NotificationModel) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetNotificationType

`func (o *NotificationModel) GetNotificationType() NotificationTypeModel`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotificationModel) GetNotificationTypeOk() (*NotificationTypeModel, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotificationModel) SetNotificationType(v NotificationTypeModel)`

SetNotificationType sets NotificationType field to given value.


### GetProjectGlobalId

`func (o *NotificationModel) GetProjectGlobalId() int64`

GetProjectGlobalId returns the ProjectGlobalId field if non-nil, zero value otherwise.

### GetProjectGlobalIdOk

`func (o *NotificationModel) GetProjectGlobalIdOk() (*int64, bool)`

GetProjectGlobalIdOk returns a tuple with the ProjectGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGlobalId

`func (o *NotificationModel) SetProjectGlobalId(v int64)`

SetProjectGlobalId sets ProjectGlobalId field to given value.

### HasProjectGlobalId

`func (o *NotificationModel) HasProjectGlobalId() bool`

HasProjectGlobalId returns a boolean if a field has been set.

### SetProjectGlobalIdNil

`func (o *NotificationModel) SetProjectGlobalIdNil(b bool)`

 SetProjectGlobalIdNil sets the value for ProjectGlobalId to be an explicit nil

### UnsetProjectGlobalId
`func (o *NotificationModel) UnsetProjectGlobalId()`

UnsetProjectGlobalId ensures that no value is present for ProjectGlobalId, not even an explicit nil
### GetProjectName

`func (o *NotificationModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *NotificationModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *NotificationModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *NotificationModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *NotificationModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *NotificationModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetTestPlanGlobalId

`func (o *NotificationModel) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *NotificationModel) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *NotificationModel) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.

### HasTestPlanGlobalId

`func (o *NotificationModel) HasTestPlanGlobalId() bool`

HasTestPlanGlobalId returns a boolean if a field has been set.

### GetTestPlanName

`func (o *NotificationModel) GetTestPlanName() string`

GetTestPlanName returns the TestPlanName field if non-nil, zero value otherwise.

### GetTestPlanNameOk

`func (o *NotificationModel) GetTestPlanNameOk() (*string, bool)`

GetTestPlanNameOk returns a tuple with the TestPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanName

`func (o *NotificationModel) SetTestPlanName(v string)`

SetTestPlanName sets TestPlanName field to given value.

### HasTestPlanName

`func (o *NotificationModel) HasTestPlanName() bool`

HasTestPlanName returns a boolean if a field has been set.

### SetTestPlanNameNil

`func (o *NotificationModel) SetTestPlanNameNil(b bool)`

 SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil

### UnsetTestPlanName
`func (o *NotificationModel) UnsetTestPlanName()`

UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
### GetWorkitemGlobalId

`func (o *NotificationModel) GetWorkitemGlobalId() int64`

GetWorkitemGlobalId returns the WorkitemGlobalId field if non-nil, zero value otherwise.

### GetWorkitemGlobalIdOk

`func (o *NotificationModel) GetWorkitemGlobalIdOk() (*int64, bool)`

GetWorkitemGlobalIdOk returns a tuple with the WorkitemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkitemGlobalId

`func (o *NotificationModel) SetWorkitemGlobalId(v int64)`

SetWorkitemGlobalId sets WorkitemGlobalId field to given value.

### HasWorkitemGlobalId

`func (o *NotificationModel) HasWorkitemGlobalId() bool`

HasWorkitemGlobalId returns a boolean if a field has been set.

### SetWorkitemGlobalIdNil

`func (o *NotificationModel) SetWorkitemGlobalIdNil(b bool)`

 SetWorkitemGlobalIdNil sets the value for WorkitemGlobalId to be an explicit nil

### UnsetWorkitemGlobalId
`func (o *NotificationModel) UnsetWorkitemGlobalId()`

UnsetWorkitemGlobalId ensures that no value is present for WorkitemGlobalId, not even an explicit nil
### GetComment

`func (o *NotificationModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NotificationModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NotificationModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NotificationModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *NotificationModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *NotificationModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetWorkItemName

`func (o *NotificationModel) GetWorkItemName() string`

GetWorkItemName returns the WorkItemName field if non-nil, zero value otherwise.

### GetWorkItemNameOk

`func (o *NotificationModel) GetWorkItemNameOk() (*string, bool)`

GetWorkItemNameOk returns a tuple with the WorkItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemName

`func (o *NotificationModel) SetWorkItemName(v string)`

SetWorkItemName sets WorkItemName field to given value.

### HasWorkItemName

`func (o *NotificationModel) HasWorkItemName() bool`

HasWorkItemName returns a boolean if a field has been set.

### SetWorkItemNameNil

`func (o *NotificationModel) SetWorkItemNameNil(b bool)`

 SetWorkItemNameNil sets the value for WorkItemName to be an explicit nil

### UnsetWorkItemName
`func (o *NotificationModel) UnsetWorkItemName()`

UnsetWorkItemName ensures that no value is present for WorkItemName, not even an explicit nil
### GetAttributeName

`func (o *NotificationModel) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *NotificationModel) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *NotificationModel) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *NotificationModel) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### SetAttributeNameNil

`func (o *NotificationModel) SetAttributeNameNil(b bool)`

 SetAttributeNameNil sets the value for AttributeName to be an explicit nil

### UnsetAttributeName
`func (o *NotificationModel) UnsetAttributeName()`

UnsetAttributeName ensures that no value is present for AttributeName, not even an explicit nil
### GetCreatedById

`func (o *NotificationModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *NotificationModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *NotificationModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *NotificationModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


