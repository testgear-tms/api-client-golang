# WorkItemFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameOrId** | Pointer to **NullableString** | Name or identifier (UUID) of work item | [optional] 
**IncludeIds** | Pointer to **[]string** | Collection of identifiers of work items which need to be included in result regardless of filtering | [optional] 
**ExcludeIds** | Pointer to **[]string** | Collection of identifiers of work items which need to be excluded from result regardless of filtering | [optional] 
**Name** | Pointer to **NullableString** | Name of work item | [optional] 
**Ids** | Pointer to **[]string** | Specifies a work item unique IDs to search for | [optional] 
**GlobalIds** | Pointer to **[]int64** | Collection of global (integer) identifiers | [optional] 
**Attributes** | Pointer to **map[string][]string** | Custom attributes of work item | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Is result must consist of only actual/deleted work items | [optional] 
**ProjectIds** | Pointer to **[]string** | Collection of project identifiers | [optional] 
**SectionIds** | Pointer to **[]string** | Collection of section identifiers | [optional] 
**CreatedByIds** | Pointer to **[]string** | Collection of identifiers of users who created work item | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Collection of identifiers of users who applied last modification to work item | [optional] 
**States** | Pointer to [**[]WorkItemStates**](WorkItemStates.md) | Collection of states of work item | [optional] 
**Priorities** | Pointer to [**[]WorkItemPriorityModel**](WorkItemPriorityModel.md) | Collection of priorities of work item | [optional] 
**Types** | Pointer to [**[]WorkItemEntityTypes**](WorkItemEntityTypes.md) | Collection of types of work item | [optional] 
**CreatedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**ModifiedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**Duration** | Pointer to [**Int32RangeSelectorModel**](Int32RangeSelectorModel.md) |  | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Is result must consist of only manual/automated work items | [optional] 
**Tags** | Pointer to **[]string** | Collection of tags | [optional] 
**AutoTestIds** | Pointer to **[]string** | Collection of identifiers of linked autotests | [optional] 

## Methods

### NewWorkItemFilterModel

`func NewWorkItemFilterModel() *WorkItemFilterModel`

NewWorkItemFilterModel instantiates a new WorkItemFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemFilterModelWithDefaults

`func NewWorkItemFilterModelWithDefaults() *WorkItemFilterModel`

NewWorkItemFilterModelWithDefaults instantiates a new WorkItemFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameOrId

`func (o *WorkItemFilterModel) GetNameOrId() string`

GetNameOrId returns the NameOrId field if non-nil, zero value otherwise.

### GetNameOrIdOk

`func (o *WorkItemFilterModel) GetNameOrIdOk() (*string, bool)`

GetNameOrIdOk returns a tuple with the NameOrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOrId

`func (o *WorkItemFilterModel) SetNameOrId(v string)`

SetNameOrId sets NameOrId field to given value.

### HasNameOrId

`func (o *WorkItemFilterModel) HasNameOrId() bool`

HasNameOrId returns a boolean if a field has been set.

### SetNameOrIdNil

`func (o *WorkItemFilterModel) SetNameOrIdNil(b bool)`

 SetNameOrIdNil sets the value for NameOrId to be an explicit nil

### UnsetNameOrId
`func (o *WorkItemFilterModel) UnsetNameOrId()`

UnsetNameOrId ensures that no value is present for NameOrId, not even an explicit nil
### GetIncludeIds

`func (o *WorkItemFilterModel) GetIncludeIds() []string`

GetIncludeIds returns the IncludeIds field if non-nil, zero value otherwise.

### GetIncludeIdsOk

`func (o *WorkItemFilterModel) GetIncludeIdsOk() (*[]string, bool)`

GetIncludeIdsOk returns a tuple with the IncludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIds

`func (o *WorkItemFilterModel) SetIncludeIds(v []string)`

SetIncludeIds sets IncludeIds field to given value.

### HasIncludeIds

`func (o *WorkItemFilterModel) HasIncludeIds() bool`

HasIncludeIds returns a boolean if a field has been set.

### SetIncludeIdsNil

`func (o *WorkItemFilterModel) SetIncludeIdsNil(b bool)`

 SetIncludeIdsNil sets the value for IncludeIds to be an explicit nil

### UnsetIncludeIds
`func (o *WorkItemFilterModel) UnsetIncludeIds()`

UnsetIncludeIds ensures that no value is present for IncludeIds, not even an explicit nil
### GetExcludeIds

`func (o *WorkItemFilterModel) GetExcludeIds() []string`

GetExcludeIds returns the ExcludeIds field if non-nil, zero value otherwise.

### GetExcludeIdsOk

`func (o *WorkItemFilterModel) GetExcludeIdsOk() (*[]string, bool)`

GetExcludeIdsOk returns a tuple with the ExcludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeIds

`func (o *WorkItemFilterModel) SetExcludeIds(v []string)`

SetExcludeIds sets ExcludeIds field to given value.

### HasExcludeIds

`func (o *WorkItemFilterModel) HasExcludeIds() bool`

HasExcludeIds returns a boolean if a field has been set.

### SetExcludeIdsNil

`func (o *WorkItemFilterModel) SetExcludeIdsNil(b bool)`

 SetExcludeIdsNil sets the value for ExcludeIds to be an explicit nil

### UnsetExcludeIds
`func (o *WorkItemFilterModel) UnsetExcludeIds()`

UnsetExcludeIds ensures that no value is present for ExcludeIds, not even an explicit nil
### GetName

`func (o *WorkItemFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkItemFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkItemFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkItemFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIds

`func (o *WorkItemFilterModel) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *WorkItemFilterModel) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *WorkItemFilterModel) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *WorkItemFilterModel) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *WorkItemFilterModel) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *WorkItemFilterModel) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetGlobalIds

`func (o *WorkItemFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *WorkItemFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *WorkItemFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *WorkItemFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *WorkItemFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *WorkItemFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetAttributes

`func (o *WorkItemFilterModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemFilterModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemFilterModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkItemFilterModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WorkItemFilterModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WorkItemFilterModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *WorkItemFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *WorkItemFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *WorkItemFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *WorkItemFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetProjectIds

`func (o *WorkItemFilterModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WorkItemFilterModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WorkItemFilterModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *WorkItemFilterModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *WorkItemFilterModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *WorkItemFilterModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetSectionIds

`func (o *WorkItemFilterModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *WorkItemFilterModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *WorkItemFilterModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *WorkItemFilterModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *WorkItemFilterModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *WorkItemFilterModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedByIds

`func (o *WorkItemFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *WorkItemFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *WorkItemFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *WorkItemFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *WorkItemFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *WorkItemFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *WorkItemFilterModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *WorkItemFilterModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *WorkItemFilterModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *WorkItemFilterModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *WorkItemFilterModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *WorkItemFilterModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetStates

`func (o *WorkItemFilterModel) GetStates() []WorkItemStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *WorkItemFilterModel) GetStatesOk() (*[]WorkItemStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *WorkItemFilterModel) SetStates(v []WorkItemStates)`

SetStates sets States field to given value.

### HasStates

`func (o *WorkItemFilterModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *WorkItemFilterModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *WorkItemFilterModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetPriorities

`func (o *WorkItemFilterModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *WorkItemFilterModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *WorkItemFilterModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *WorkItemFilterModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *WorkItemFilterModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *WorkItemFilterModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetTypes

`func (o *WorkItemFilterModel) GetTypes() []WorkItemEntityTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WorkItemFilterModel) GetTypesOk() (*[]WorkItemEntityTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WorkItemFilterModel) SetTypes(v []WorkItemEntityTypes)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WorkItemFilterModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *WorkItemFilterModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *WorkItemFilterModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetCreatedDate

`func (o *WorkItemFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WorkItemFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *WorkItemFilterModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemFilterModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemFilterModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemFilterModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetDuration

`func (o *WorkItemFilterModel) GetDuration() Int32RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemFilterModel) GetDurationOk() (*Int32RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemFilterModel) SetDuration(v Int32RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkItemFilterModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetIsAutomated

`func (o *WorkItemFilterModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemFilterModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemFilterModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *WorkItemFilterModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *WorkItemFilterModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *WorkItemFilterModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetTags

`func (o *WorkItemFilterModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemFilterModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemFilterModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkItemFilterModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *WorkItemFilterModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *WorkItemFilterModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAutoTestIds

`func (o *WorkItemFilterModel) GetAutoTestIds() []string`

GetAutoTestIds returns the AutoTestIds field if non-nil, zero value otherwise.

### GetAutoTestIdsOk

`func (o *WorkItemFilterModel) GetAutoTestIdsOk() (*[]string, bool)`

GetAutoTestIdsOk returns a tuple with the AutoTestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestIds

`func (o *WorkItemFilterModel) SetAutoTestIds(v []string)`

SetAutoTestIds sets AutoTestIds field to given value.

### HasAutoTestIds

`func (o *WorkItemFilterModel) HasAutoTestIds() bool`

HasAutoTestIds returns a boolean if a field has been set.

### SetAutoTestIdsNil

`func (o *WorkItemFilterModel) SetAutoTestIdsNil(b bool)`

 SetAutoTestIdsNil sets the value for AutoTestIds to be an explicit nil

### UnsetAutoTestIds
`func (o *WorkItemFilterModel) UnsetAutoTestIds()`

UnsetAutoTestIds ensures that no value is present for AutoTestIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


