# SharedStepReferencesQueryFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of work item | [optional] 
**GlobalIds** | Pointer to **[]int64** | Collection of global (integer) identifiers | [optional] 
**SectionIds** | Pointer to **[]string** | Collection of section identifiers | [optional] 
**CreatedByIds** | Pointer to **[]string** | Collection of identifiers of users who created work item | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Collection of identifiers of users who applied last modification to work item | [optional] 
**States** | Pointer to [**[]WorkItemStates**](WorkItemStates.md) | Collection of states of work item | [optional] 
**Priorities** | Pointer to [**[]WorkItemPriorityModel**](WorkItemPriorityModel.md) | Collection of priorities of work item | [optional] 
**EntityTypes** | Pointer to **[]string** | Collection of types of work item  &lt;br&gt;Allowed values: &#x60;TestCases&#x60;, &#x60;CheckLists&#x60;, &#x60;SharedSteps&#x60; | [optional] 
**CreatedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**ModifiedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Is result must consist of only manual/automated work items | [optional] 
**Tags** | Pointer to **[]string** | Collection of tags | [optional] 

## Methods

### NewSharedStepReferencesQueryFilterModel

`func NewSharedStepReferencesQueryFilterModel() *SharedStepReferencesQueryFilterModel`

NewSharedStepReferencesQueryFilterModel instantiates a new SharedStepReferencesQueryFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepReferencesQueryFilterModelWithDefaults

`func NewSharedStepReferencesQueryFilterModelWithDefaults() *SharedStepReferencesQueryFilterModel`

NewSharedStepReferencesQueryFilterModelWithDefaults instantiates a new SharedStepReferencesQueryFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SharedStepReferencesQueryFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedStepReferencesQueryFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedStepReferencesQueryFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedStepReferencesQueryFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SharedStepReferencesQueryFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SharedStepReferencesQueryFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGlobalIds

`func (o *SharedStepReferencesQueryFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *SharedStepReferencesQueryFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *SharedStepReferencesQueryFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *SharedStepReferencesQueryFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *SharedStepReferencesQueryFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *SharedStepReferencesQueryFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetSectionIds

`func (o *SharedStepReferencesQueryFilterModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *SharedStepReferencesQueryFilterModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *SharedStepReferencesQueryFilterModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *SharedStepReferencesQueryFilterModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *SharedStepReferencesQueryFilterModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *SharedStepReferencesQueryFilterModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedByIds

`func (o *SharedStepReferencesQueryFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *SharedStepReferencesQueryFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *SharedStepReferencesQueryFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *SharedStepReferencesQueryFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *SharedStepReferencesQueryFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *SharedStepReferencesQueryFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *SharedStepReferencesQueryFilterModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *SharedStepReferencesQueryFilterModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *SharedStepReferencesQueryFilterModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *SharedStepReferencesQueryFilterModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *SharedStepReferencesQueryFilterModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *SharedStepReferencesQueryFilterModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetStates

`func (o *SharedStepReferencesQueryFilterModel) GetStates() []WorkItemStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *SharedStepReferencesQueryFilterModel) GetStatesOk() (*[]WorkItemStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *SharedStepReferencesQueryFilterModel) SetStates(v []WorkItemStates)`

SetStates sets States field to given value.

### HasStates

`func (o *SharedStepReferencesQueryFilterModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *SharedStepReferencesQueryFilterModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *SharedStepReferencesQueryFilterModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetPriorities

`func (o *SharedStepReferencesQueryFilterModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *SharedStepReferencesQueryFilterModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *SharedStepReferencesQueryFilterModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *SharedStepReferencesQueryFilterModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *SharedStepReferencesQueryFilterModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *SharedStepReferencesQueryFilterModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetEntityTypes

`func (o *SharedStepReferencesQueryFilterModel) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *SharedStepReferencesQueryFilterModel) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *SharedStepReferencesQueryFilterModel) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *SharedStepReferencesQueryFilterModel) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### SetEntityTypesNil

`func (o *SharedStepReferencesQueryFilterModel) SetEntityTypesNil(b bool)`

 SetEntityTypesNil sets the value for EntityTypes to be an explicit nil

### UnsetEntityTypes
`func (o *SharedStepReferencesQueryFilterModel) UnsetEntityTypes()`

UnsetEntityTypes ensures that no value is present for EntityTypes, not even an explicit nil
### GetCreatedDate

`func (o *SharedStepReferencesQueryFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *SharedStepReferencesQueryFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *SharedStepReferencesQueryFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *SharedStepReferencesQueryFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *SharedStepReferencesQueryFilterModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SharedStepReferencesQueryFilterModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SharedStepReferencesQueryFilterModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SharedStepReferencesQueryFilterModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetIsAutomated

`func (o *SharedStepReferencesQueryFilterModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *SharedStepReferencesQueryFilterModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *SharedStepReferencesQueryFilterModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *SharedStepReferencesQueryFilterModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *SharedStepReferencesQueryFilterModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *SharedStepReferencesQueryFilterModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetTags

`func (o *SharedStepReferencesQueryFilterModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SharedStepReferencesQueryFilterModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SharedStepReferencesQueryFilterModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SharedStepReferencesQueryFilterModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SharedStepReferencesQueryFilterModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SharedStepReferencesQueryFilterModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


