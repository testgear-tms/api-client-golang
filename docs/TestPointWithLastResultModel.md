# TestPointWithLastResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**WorkItemName** | Pointer to **NullableString** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**TesterId** | Pointer to **NullableString** |  | [optional] 
**WorkItemId** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **NullableString** |  | [optional] 
**TestSuiteId** | Pointer to **string** |  | [optional] 
**LastTestResult** | Pointer to [**LastTestResultModel**](LastTestResultModel.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**WorkItemGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**WorkItemEntityTypeName** | Pointer to **NullableString** |  | [optional] 
**SectionId** | Pointer to **string** |  | [optional] 
**SectionName** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**TagNames** | Pointer to **[]string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**TestSuiteNameBreadCrumbs** | Pointer to **[]string** |  | [optional] 
**GroupCount** | Pointer to **NullableInt32** |  | [optional] 
**Iteration** | Pointer to [**IterationModel**](IterationModel.md) |  | [optional] 

## Methods

### NewTestPointWithLastResultModel

`func NewTestPointWithLastResultModel(priority WorkItemPriorityModel, ) *TestPointWithLastResultModel`

NewTestPointWithLastResultModel instantiates a new TestPointWithLastResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointWithLastResultModelWithDefaults

`func NewTestPointWithLastResultModelWithDefaults() *TestPointWithLastResultModel`

NewTestPointWithLastResultModelWithDefaults instantiates a new TestPointWithLastResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPointWithLastResultModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointWithLastResultModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointWithLastResultModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestPointWithLastResultModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkItemName

`func (o *TestPointWithLastResultModel) GetWorkItemName() string`

GetWorkItemName returns the WorkItemName field if non-nil, zero value otherwise.

### GetWorkItemNameOk

`func (o *TestPointWithLastResultModel) GetWorkItemNameOk() (*string, bool)`

GetWorkItemNameOk returns a tuple with the WorkItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemName

`func (o *TestPointWithLastResultModel) SetWorkItemName(v string)`

SetWorkItemName sets WorkItemName field to given value.

### HasWorkItemName

`func (o *TestPointWithLastResultModel) HasWorkItemName() bool`

HasWorkItemName returns a boolean if a field has been set.

### SetWorkItemNameNil

`func (o *TestPointWithLastResultModel) SetWorkItemNameNil(b bool)`

 SetWorkItemNameNil sets the value for WorkItemName to be an explicit nil

### UnsetWorkItemName
`func (o *TestPointWithLastResultModel) UnsetWorkItemName()`

UnsetWorkItemName ensures that no value is present for WorkItemName, not even an explicit nil
### GetIsAutomated

`func (o *TestPointWithLastResultModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPointWithLastResultModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPointWithLastResultModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestPointWithLastResultModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetTesterId

`func (o *TestPointWithLastResultModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointWithLastResultModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointWithLastResultModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointWithLastResultModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointWithLastResultModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointWithLastResultModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetWorkItemId

`func (o *TestPointWithLastResultModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointWithLastResultModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointWithLastResultModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *TestPointWithLastResultModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### GetConfigurationId

`func (o *TestPointWithLastResultModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointWithLastResultModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointWithLastResultModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestPointWithLastResultModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *TestPointWithLastResultModel) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *TestPointWithLastResultModel) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetTestSuiteId

`func (o *TestPointWithLastResultModel) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointWithLastResultModel) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointWithLastResultModel) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.

### HasTestSuiteId

`func (o *TestPointWithLastResultModel) HasTestSuiteId() bool`

HasTestSuiteId returns a boolean if a field has been set.

### GetLastTestResult

`func (o *TestPointWithLastResultModel) GetLastTestResult() LastTestResultModel`

GetLastTestResult returns the LastTestResult field if non-nil, zero value otherwise.

### GetLastTestResultOk

`func (o *TestPointWithLastResultModel) GetLastTestResultOk() (*LastTestResultModel, bool)`

GetLastTestResultOk returns a tuple with the LastTestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResult

`func (o *TestPointWithLastResultModel) SetLastTestResult(v LastTestResultModel)`

SetLastTestResult sets LastTestResult field to given value.

### HasLastTestResult

`func (o *TestPointWithLastResultModel) HasLastTestResult() bool`

HasLastTestResult returns a boolean if a field has been set.

### GetStatus

`func (o *TestPointWithLastResultModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointWithLastResultModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointWithLastResultModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPointWithLastResultModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPointWithLastResultModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPointWithLastResultModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetWorkItemGlobalId

`func (o *TestPointWithLastResultModel) GetWorkItemGlobalId() int64`

GetWorkItemGlobalId returns the WorkItemGlobalId field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdOk

`func (o *TestPointWithLastResultModel) GetWorkItemGlobalIdOk() (*int64, bool)`

GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalId

`func (o *TestPointWithLastResultModel) SetWorkItemGlobalId(v int64)`

SetWorkItemGlobalId sets WorkItemGlobalId field to given value.

### HasWorkItemGlobalId

`func (o *TestPointWithLastResultModel) HasWorkItemGlobalId() bool`

HasWorkItemGlobalId returns a boolean if a field has been set.

### SetWorkItemGlobalIdNil

`func (o *TestPointWithLastResultModel) SetWorkItemGlobalIdNil(b bool)`

 SetWorkItemGlobalIdNil sets the value for WorkItemGlobalId to be an explicit nil

### UnsetWorkItemGlobalId
`func (o *TestPointWithLastResultModel) UnsetWorkItemGlobalId()`

UnsetWorkItemGlobalId ensures that no value is present for WorkItemGlobalId, not even an explicit nil
### GetWorkItemEntityTypeName

`func (o *TestPointWithLastResultModel) GetWorkItemEntityTypeName() string`

GetWorkItemEntityTypeName returns the WorkItemEntityTypeName field if non-nil, zero value otherwise.

### GetWorkItemEntityTypeNameOk

`func (o *TestPointWithLastResultModel) GetWorkItemEntityTypeNameOk() (*string, bool)`

GetWorkItemEntityTypeNameOk returns a tuple with the WorkItemEntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemEntityTypeName

`func (o *TestPointWithLastResultModel) SetWorkItemEntityTypeName(v string)`

SetWorkItemEntityTypeName sets WorkItemEntityTypeName field to given value.

### HasWorkItemEntityTypeName

`func (o *TestPointWithLastResultModel) HasWorkItemEntityTypeName() bool`

HasWorkItemEntityTypeName returns a boolean if a field has been set.

### SetWorkItemEntityTypeNameNil

`func (o *TestPointWithLastResultModel) SetWorkItemEntityTypeNameNil(b bool)`

 SetWorkItemEntityTypeNameNil sets the value for WorkItemEntityTypeName to be an explicit nil

### UnsetWorkItemEntityTypeName
`func (o *TestPointWithLastResultModel) UnsetWorkItemEntityTypeName()`

UnsetWorkItemEntityTypeName ensures that no value is present for WorkItemEntityTypeName, not even an explicit nil
### GetSectionId

`func (o *TestPointWithLastResultModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TestPointWithLastResultModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TestPointWithLastResultModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *TestPointWithLastResultModel) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetSectionName

`func (o *TestPointWithLastResultModel) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *TestPointWithLastResultModel) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *TestPointWithLastResultModel) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.

### HasSectionName

`func (o *TestPointWithLastResultModel) HasSectionName() bool`

HasSectionName returns a boolean if a field has been set.

### SetSectionNameNil

`func (o *TestPointWithLastResultModel) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *TestPointWithLastResultModel) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetCreatedDate

`func (o *TestPointWithLastResultModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPointWithLastResultModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPointWithLastResultModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPointWithLastResultModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPointWithLastResultModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPointWithLastResultModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *TestPointWithLastResultModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPointWithLastResultModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPointWithLastResultModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPointWithLastResultModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPointWithLastResultModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPointWithLastResultModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestPointWithLastResultModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPointWithLastResultModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPointWithLastResultModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestPointWithLastResultModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *TestPointWithLastResultModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestPointWithLastResultModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestPointWithLastResultModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestPointWithLastResultModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestPointWithLastResultModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestPointWithLastResultModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetAttributes

`func (o *TestPointWithLastResultModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPointWithLastResultModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPointWithLastResultModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPointWithLastResultModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPointWithLastResultModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPointWithLastResultModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetTagNames

`func (o *TestPointWithLastResultModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *TestPointWithLastResultModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *TestPointWithLastResultModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *TestPointWithLastResultModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *TestPointWithLastResultModel) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *TestPointWithLastResultModel) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetDuration

`func (o *TestPointWithLastResultModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPointWithLastResultModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPointWithLastResultModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestPointWithLastResultModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetPriority

`func (o *TestPointWithLastResultModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestPointWithLastResultModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestPointWithLastResultModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetTestSuiteNameBreadCrumbs

`func (o *TestPointWithLastResultModel) GetTestSuiteNameBreadCrumbs() []string`

GetTestSuiteNameBreadCrumbs returns the TestSuiteNameBreadCrumbs field if non-nil, zero value otherwise.

### GetTestSuiteNameBreadCrumbsOk

`func (o *TestPointWithLastResultModel) GetTestSuiteNameBreadCrumbsOk() (*[]string, bool)`

GetTestSuiteNameBreadCrumbsOk returns a tuple with the TestSuiteNameBreadCrumbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteNameBreadCrumbs

`func (o *TestPointWithLastResultModel) SetTestSuiteNameBreadCrumbs(v []string)`

SetTestSuiteNameBreadCrumbs sets TestSuiteNameBreadCrumbs field to given value.

### HasTestSuiteNameBreadCrumbs

`func (o *TestPointWithLastResultModel) HasTestSuiteNameBreadCrumbs() bool`

HasTestSuiteNameBreadCrumbs returns a boolean if a field has been set.

### SetTestSuiteNameBreadCrumbsNil

`func (o *TestPointWithLastResultModel) SetTestSuiteNameBreadCrumbsNil(b bool)`

 SetTestSuiteNameBreadCrumbsNil sets the value for TestSuiteNameBreadCrumbs to be an explicit nil

### UnsetTestSuiteNameBreadCrumbs
`func (o *TestPointWithLastResultModel) UnsetTestSuiteNameBreadCrumbs()`

UnsetTestSuiteNameBreadCrumbs ensures that no value is present for TestSuiteNameBreadCrumbs, not even an explicit nil
### GetGroupCount

`func (o *TestPointWithLastResultModel) GetGroupCount() int32`

GetGroupCount returns the GroupCount field if non-nil, zero value otherwise.

### GetGroupCountOk

`func (o *TestPointWithLastResultModel) GetGroupCountOk() (*int32, bool)`

GetGroupCountOk returns a tuple with the GroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCount

`func (o *TestPointWithLastResultModel) SetGroupCount(v int32)`

SetGroupCount sets GroupCount field to given value.

### HasGroupCount

`func (o *TestPointWithLastResultModel) HasGroupCount() bool`

HasGroupCount returns a boolean if a field has been set.

### SetGroupCountNil

`func (o *TestPointWithLastResultModel) SetGroupCountNil(b bool)`

 SetGroupCountNil sets the value for GroupCount to be an explicit nil

### UnsetGroupCount
`func (o *TestPointWithLastResultModel) UnsetGroupCount()`

UnsetGroupCount ensures that no value is present for GroupCount, not even an explicit nil
### GetIteration

`func (o *TestPointWithLastResultModel) GetIteration() IterationModel`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *TestPointWithLastResultModel) GetIterationOk() (*IterationModel, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *TestPointWithLastResultModel) SetIteration(v IterationModel)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *TestPointWithLastResultModel) HasIteration() bool`

HasIteration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


