# TestSuiteWorkItemsSearchModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of work item | [optional] 
**GlobalIds** | Pointer to **[]int64** | Collection of global (integer) identifiers | [optional] 
**SectionIds** | Pointer to **[]string** | Collection of section identifiers | [optional] 
**Priorities** | Pointer to [**[]WorkItemPriorityModel**](WorkItemPriorityModel.md) | Collection of priorities of work item | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Is result must consist of only manual/automated work items | [optional] 
**States** | Pointer to [**[]WorkItemStates**](WorkItemStates.md) | Collection of states of work item | [optional] 
**Duration** | Pointer to [**Int32RangeSelectorModel**](Int32RangeSelectorModel.md) |  | [optional] 
**CreatedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**ModifiedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** | Collection of identifiers of users who created work item | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Collection of identifiers of users who applied last modification to work item | [optional] 
**Attributes** | Pointer to **map[string][]string** | Custom attributes of work item | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Is result must consist of only actual/deleted work items | [optional] 
**TagNames** | Pointer to **[]string** | Collection of tags | [optional] 
**EntityTypes** | Pointer to **[]string** | Collection of types of work item  &lt;br&gt;Allowed values: &#x60;TestCases&#x60;, &#x60;CheckLists&#x60;, &#x60;SharedSteps&#x60; | [optional] 

## Methods

### NewTestSuiteWorkItemsSearchModel

`func NewTestSuiteWorkItemsSearchModel() *TestSuiteWorkItemsSearchModel`

NewTestSuiteWorkItemsSearchModel instantiates a new TestSuiteWorkItemsSearchModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteWorkItemsSearchModelWithDefaults

`func NewTestSuiteWorkItemsSearchModelWithDefaults() *TestSuiteWorkItemsSearchModel`

NewTestSuiteWorkItemsSearchModelWithDefaults instantiates a new TestSuiteWorkItemsSearchModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestSuiteWorkItemsSearchModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteWorkItemsSearchModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteWorkItemsSearchModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestSuiteWorkItemsSearchModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestSuiteWorkItemsSearchModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestSuiteWorkItemsSearchModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGlobalIds

`func (o *TestSuiteWorkItemsSearchModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *TestSuiteWorkItemsSearchModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *TestSuiteWorkItemsSearchModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetSectionIds

`func (o *TestSuiteWorkItemsSearchModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *TestSuiteWorkItemsSearchModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *TestSuiteWorkItemsSearchModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetPriorities

`func (o *TestSuiteWorkItemsSearchModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *TestSuiteWorkItemsSearchModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *TestSuiteWorkItemsSearchModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *TestSuiteWorkItemsSearchModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *TestSuiteWorkItemsSearchModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *TestSuiteWorkItemsSearchModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetIsAutomated

`func (o *TestSuiteWorkItemsSearchModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestSuiteWorkItemsSearchModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestSuiteWorkItemsSearchModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestSuiteWorkItemsSearchModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *TestSuiteWorkItemsSearchModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *TestSuiteWorkItemsSearchModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetStates

`func (o *TestSuiteWorkItemsSearchModel) GetStates() []WorkItemStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *TestSuiteWorkItemsSearchModel) GetStatesOk() (*[]WorkItemStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *TestSuiteWorkItemsSearchModel) SetStates(v []WorkItemStates)`

SetStates sets States field to given value.

### HasStates

`func (o *TestSuiteWorkItemsSearchModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *TestSuiteWorkItemsSearchModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *TestSuiteWorkItemsSearchModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetDuration

`func (o *TestSuiteWorkItemsSearchModel) GetDuration() Int32RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestSuiteWorkItemsSearchModel) GetDurationOk() (*Int32RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestSuiteWorkItemsSearchModel) SetDuration(v Int32RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestSuiteWorkItemsSearchModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestSuiteWorkItemsSearchModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestSuiteWorkItemsSearchModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestSuiteWorkItemsSearchModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestSuiteWorkItemsSearchModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetCreatedByIds

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestSuiteWorkItemsSearchModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestSuiteWorkItemsSearchModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestSuiteWorkItemsSearchModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestSuiteWorkItemsSearchModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetAttributes

`func (o *TestSuiteWorkItemsSearchModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestSuiteWorkItemsSearchModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestSuiteWorkItemsSearchModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestSuiteWorkItemsSearchModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestSuiteWorkItemsSearchModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestSuiteWorkItemsSearchModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *TestSuiteWorkItemsSearchModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestSuiteWorkItemsSearchModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestSuiteWorkItemsSearchModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestSuiteWorkItemsSearchModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *TestSuiteWorkItemsSearchModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *TestSuiteWorkItemsSearchModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetTagNames

`func (o *TestSuiteWorkItemsSearchModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *TestSuiteWorkItemsSearchModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *TestSuiteWorkItemsSearchModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *TestSuiteWorkItemsSearchModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *TestSuiteWorkItemsSearchModel) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *TestSuiteWorkItemsSearchModel) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetEntityTypes

`func (o *TestSuiteWorkItemsSearchModel) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *TestSuiteWorkItemsSearchModel) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *TestSuiteWorkItemsSearchModel) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *TestSuiteWorkItemsSearchModel) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### SetEntityTypesNil

`func (o *TestSuiteWorkItemsSearchModel) SetEntityTypesNil(b bool)`

 SetEntityTypesNil sets the value for EntityTypes to be an explicit nil

### UnsetEntityTypes
`func (o *TestSuiteWorkItemsSearchModel) UnsetEntityTypes()`

UnsetEntityTypes ensures that no value is present for EntityTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


