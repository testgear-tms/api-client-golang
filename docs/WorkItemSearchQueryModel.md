# WorkItemSearchQueryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkItemSearchQueryModel

`func NewWorkItemSearchQueryModel() *WorkItemSearchQueryModel`

NewWorkItemSearchQueryModel instantiates a new WorkItemSearchQueryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemSearchQueryModelWithDefaults

`func NewWorkItemSearchQueryModelWithDefaults() *WorkItemSearchQueryModel`

NewWorkItemSearchQueryModelWithDefaults instantiates a new WorkItemSearchQueryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkItemSearchQueryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemSearchQueryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemSearchQueryModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkItemSearchQueryModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkItemSearchQueryModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkItemSearchQueryModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIds

`func (o *WorkItemSearchQueryModel) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *WorkItemSearchQueryModel) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *WorkItemSearchQueryModel) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *WorkItemSearchQueryModel) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *WorkItemSearchQueryModel) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *WorkItemSearchQueryModel) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetGlobalIds

`func (o *WorkItemSearchQueryModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *WorkItemSearchQueryModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *WorkItemSearchQueryModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *WorkItemSearchQueryModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *WorkItemSearchQueryModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *WorkItemSearchQueryModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetAttributes

`func (o *WorkItemSearchQueryModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemSearchQueryModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemSearchQueryModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkItemSearchQueryModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WorkItemSearchQueryModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WorkItemSearchQueryModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *WorkItemSearchQueryModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemSearchQueryModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemSearchQueryModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *WorkItemSearchQueryModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *WorkItemSearchQueryModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *WorkItemSearchQueryModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetProjectIds

`func (o *WorkItemSearchQueryModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WorkItemSearchQueryModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WorkItemSearchQueryModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *WorkItemSearchQueryModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *WorkItemSearchQueryModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *WorkItemSearchQueryModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetSectionIds

`func (o *WorkItemSearchQueryModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *WorkItemSearchQueryModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *WorkItemSearchQueryModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *WorkItemSearchQueryModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *WorkItemSearchQueryModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *WorkItemSearchQueryModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedByIds

`func (o *WorkItemSearchQueryModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *WorkItemSearchQueryModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *WorkItemSearchQueryModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *WorkItemSearchQueryModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *WorkItemSearchQueryModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *WorkItemSearchQueryModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *WorkItemSearchQueryModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *WorkItemSearchQueryModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *WorkItemSearchQueryModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *WorkItemSearchQueryModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *WorkItemSearchQueryModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *WorkItemSearchQueryModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetStates

`func (o *WorkItemSearchQueryModel) GetStates() []WorkItemStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *WorkItemSearchQueryModel) GetStatesOk() (*[]WorkItemStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *WorkItemSearchQueryModel) SetStates(v []WorkItemStates)`

SetStates sets States field to given value.

### HasStates

`func (o *WorkItemSearchQueryModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *WorkItemSearchQueryModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *WorkItemSearchQueryModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetPriorities

`func (o *WorkItemSearchQueryModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *WorkItemSearchQueryModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *WorkItemSearchQueryModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *WorkItemSearchQueryModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *WorkItemSearchQueryModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *WorkItemSearchQueryModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetTypes

`func (o *WorkItemSearchQueryModel) GetTypes() []WorkItemEntityTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WorkItemSearchQueryModel) GetTypesOk() (*[]WorkItemEntityTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WorkItemSearchQueryModel) SetTypes(v []WorkItemEntityTypes)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WorkItemSearchQueryModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *WorkItemSearchQueryModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *WorkItemSearchQueryModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetCreatedDate

`func (o *WorkItemSearchQueryModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemSearchQueryModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemSearchQueryModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WorkItemSearchQueryModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *WorkItemSearchQueryModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemSearchQueryModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemSearchQueryModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemSearchQueryModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetDuration

`func (o *WorkItemSearchQueryModel) GetDuration() Int32RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemSearchQueryModel) GetDurationOk() (*Int32RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemSearchQueryModel) SetDuration(v Int32RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkItemSearchQueryModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetIsAutomated

`func (o *WorkItemSearchQueryModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemSearchQueryModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemSearchQueryModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *WorkItemSearchQueryModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *WorkItemSearchQueryModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *WorkItemSearchQueryModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetTags

`func (o *WorkItemSearchQueryModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemSearchQueryModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemSearchQueryModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkItemSearchQueryModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *WorkItemSearchQueryModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *WorkItemSearchQueryModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAutoTestIds

`func (o *WorkItemSearchQueryModel) GetAutoTestIds() []string`

GetAutoTestIds returns the AutoTestIds field if non-nil, zero value otherwise.

### GetAutoTestIdsOk

`func (o *WorkItemSearchQueryModel) GetAutoTestIdsOk() (*[]string, bool)`

GetAutoTestIdsOk returns a tuple with the AutoTestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestIds

`func (o *WorkItemSearchQueryModel) SetAutoTestIds(v []string)`

SetAutoTestIds sets AutoTestIds field to given value.

### HasAutoTestIds

`func (o *WorkItemSearchQueryModel) HasAutoTestIds() bool`

HasAutoTestIds returns a boolean if a field has been set.

### SetAutoTestIdsNil

`func (o *WorkItemSearchQueryModel) SetAutoTestIdsNil(b bool)`

 SetAutoTestIdsNil sets the value for AutoTestIds to be an explicit nil

### UnsetAutoTestIds
`func (o *WorkItemSearchQueryModel) UnsetAutoTestIds()`

UnsetAutoTestIds ensures that no value is present for AutoTestIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


