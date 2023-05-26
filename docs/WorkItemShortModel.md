# WorkItemShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** | used for versioning changes in workitem | [optional] 
**Name** | **string** |  | 
**EntityTypeName** | **string** | Property can have one of these values: CheckLists, SharedSteps, TestCases | 
**ProjectId** | **string** | This property is used to link autotest with project | 
**SectionId** | **string** | This property links workitem with section | 
**SectionName** | **string** | Name of the section where work item is located | 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**TagNames** | Pointer to **[]string** |  | [optional] 
**Iterations** | Pointer to [**[]IterationModel**](IterationModel.md) |  | [optional] 

## Methods

### NewWorkItemShortModel

`func NewWorkItemShortModel(name string, entityTypeName string, projectId string, sectionId string, sectionName string, state WorkItemStates, priority WorkItemPriorityModel, ) *WorkItemShortModel`

NewWorkItemShortModel instantiates a new WorkItemShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemShortModelWithDefaults

`func NewWorkItemShortModelWithDefaults() *WorkItemShortModel`

NewWorkItemShortModelWithDefaults instantiates a new WorkItemShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemShortModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkItemShortModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionId

`func (o *WorkItemShortModel) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkItemShortModel) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkItemShortModel) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *WorkItemShortModel) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetName

`func (o *WorkItemShortModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemShortModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemShortModel) SetName(v string)`

SetName sets Name field to given value.


### GetEntityTypeName

`func (o *WorkItemShortModel) GetEntityTypeName() string`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemShortModel) GetEntityTypeNameOk() (*string, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemShortModel) SetEntityTypeName(v string)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetProjectId

`func (o *WorkItemShortModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemShortModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemShortModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSectionId

`func (o *WorkItemShortModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemShortModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemShortModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetSectionName

`func (o *WorkItemShortModel) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *WorkItemShortModel) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *WorkItemShortModel) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.


### GetIsAutomated

`func (o *WorkItemShortModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemShortModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemShortModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *WorkItemShortModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetGlobalId

`func (o *WorkItemShortModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemShortModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemShortModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *WorkItemShortModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetDuration

`func (o *WorkItemShortModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemShortModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemShortModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkItemShortModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAttributes

`func (o *WorkItemShortModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemShortModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemShortModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkItemShortModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WorkItemShortModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WorkItemShortModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCreatedById

`func (o *WorkItemShortModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemShortModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemShortModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *WorkItemShortModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *WorkItemShortModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkItemShortModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkItemShortModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WorkItemShortModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WorkItemShortModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WorkItemShortModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedDate

`func (o *WorkItemShortModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemShortModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemShortModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WorkItemShortModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *WorkItemShortModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *WorkItemShortModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *WorkItemShortModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemShortModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemShortModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemShortModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WorkItemShortModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WorkItemShortModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetState

`func (o *WorkItemShortModel) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemShortModel) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemShortModel) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemShortModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemShortModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemShortModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetIsDeleted

`func (o *WorkItemShortModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemShortModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemShortModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *WorkItemShortModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetTagNames

`func (o *WorkItemShortModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *WorkItemShortModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *WorkItemShortModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *WorkItemShortModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *WorkItemShortModel) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *WorkItemShortModel) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetIterations

`func (o *WorkItemShortModel) GetIterations() []IterationModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *WorkItemShortModel) GetIterationsOk() (*[]IterationModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *WorkItemShortModel) SetIterations(v []IterationModel)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *WorkItemShortModel) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *WorkItemShortModel) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *WorkItemShortModel) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


