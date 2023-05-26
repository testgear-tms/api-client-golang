# WorkItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | Pointer to **string** | used for versioning changes in workitem | [optional] 
**MedianDuration** | Pointer to **int64** | used for getting a median duration of all autotests related to this workitem | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**EntityTypeName** | [**WorkItemEntityTypes**](WorkItemEntityTypes.md) |  | 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**AutoTests** | Pointer to [**[]AutoTestModel**](AutoTestModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 
**SectionPreconditionSteps** | Pointer to [**[]StepModel**](StepModel.md) |  | [optional] 
**SectionPostconditionSteps** | Pointer to [**[]StepModel**](StepModel.md) |  | [optional] 
**VersionNumber** | Pointer to **int32** | used for define chronology of workitem state in each version | [optional] 
**Iterations** | Pointer to [**[]IterationModel**](IterationModel.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**Id** | **string** |  | 
**SectionId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**Steps** | [**[]StepModel**](StepModel.md) |  | 
**PreconditionSteps** | [**[]StepModel**](StepModel.md) |  | 
**PostconditionSteps** | [**[]StepModel**](StepModel.md) |  | 
**Duration** | Pointer to **int32** |  | [optional] 
**Attributes** | **map[string]interface{}** |  | 
**Tags** | [**[]TagShortModel**](TagShortModel.md) |  | 
**Links** | [**[]LinkModel**](LinkModel.md) |  | 
**Name** | **string** |  | 

## Methods

### NewWorkItemModel

`func NewWorkItemModel(entityTypeName WorkItemEntityTypes, id string, state WorkItemStates, priority WorkItemPriorityModel, steps []StepModel, preconditionSteps []StepModel, postconditionSteps []StepModel, attributes map[string]interface{}, tags []TagShortModel, links []LinkModel, name string, ) *WorkItemModel`

NewWorkItemModel instantiates a new WorkItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemModelWithDefaults

`func NewWorkItemModelWithDefaults() *WorkItemModel`

NewWorkItemModelWithDefaults instantiates a new WorkItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *WorkItemModel) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkItemModel) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkItemModel) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *WorkItemModel) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetMedianDuration

`func (o *WorkItemModel) GetMedianDuration() int64`

GetMedianDuration returns the MedianDuration field if non-nil, zero value otherwise.

### GetMedianDurationOk

`func (o *WorkItemModel) GetMedianDurationOk() (*int64, bool)`

GetMedianDurationOk returns a tuple with the MedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDuration

`func (o *WorkItemModel) SetMedianDuration(v int64)`

SetMedianDuration sets MedianDuration field to given value.

### HasMedianDuration

`func (o *WorkItemModel) HasMedianDuration() bool`

HasMedianDuration returns a boolean if a field has been set.

### GetIsDeleted

`func (o *WorkItemModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *WorkItemModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetProjectId

`func (o *WorkItemModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WorkItemModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetEntityTypeName

`func (o *WorkItemModel) GetEntityTypeName() WorkItemEntityTypes`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemModel) GetEntityTypeNameOk() (*WorkItemEntityTypes, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemModel) SetEntityTypeName(v WorkItemEntityTypes)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetIsAutomated

`func (o *WorkItemModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *WorkItemModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetAutoTests

`func (o *WorkItemModel) GetAutoTests() []AutoTestModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *WorkItemModel) GetAutoTestsOk() (*[]AutoTestModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *WorkItemModel) SetAutoTests(v []AutoTestModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *WorkItemModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *WorkItemModel) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *WorkItemModel) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil
### GetAttachments

`func (o *WorkItemModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkItemModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkItemModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *WorkItemModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *WorkItemModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *WorkItemModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetSectionPreconditionSteps

`func (o *WorkItemModel) GetSectionPreconditionSteps() []StepModel`

GetSectionPreconditionSteps returns the SectionPreconditionSteps field if non-nil, zero value otherwise.

### GetSectionPreconditionStepsOk

`func (o *WorkItemModel) GetSectionPreconditionStepsOk() (*[]StepModel, bool)`

GetSectionPreconditionStepsOk returns a tuple with the SectionPreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionPreconditionSteps

`func (o *WorkItemModel) SetSectionPreconditionSteps(v []StepModel)`

SetSectionPreconditionSteps sets SectionPreconditionSteps field to given value.

### HasSectionPreconditionSteps

`func (o *WorkItemModel) HasSectionPreconditionSteps() bool`

HasSectionPreconditionSteps returns a boolean if a field has been set.

### SetSectionPreconditionStepsNil

`func (o *WorkItemModel) SetSectionPreconditionStepsNil(b bool)`

 SetSectionPreconditionStepsNil sets the value for SectionPreconditionSteps to be an explicit nil

### UnsetSectionPreconditionSteps
`func (o *WorkItemModel) UnsetSectionPreconditionSteps()`

UnsetSectionPreconditionSteps ensures that no value is present for SectionPreconditionSteps, not even an explicit nil
### GetSectionPostconditionSteps

`func (o *WorkItemModel) GetSectionPostconditionSteps() []StepModel`

GetSectionPostconditionSteps returns the SectionPostconditionSteps field if non-nil, zero value otherwise.

### GetSectionPostconditionStepsOk

`func (o *WorkItemModel) GetSectionPostconditionStepsOk() (*[]StepModel, bool)`

GetSectionPostconditionStepsOk returns a tuple with the SectionPostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionPostconditionSteps

`func (o *WorkItemModel) SetSectionPostconditionSteps(v []StepModel)`

SetSectionPostconditionSteps sets SectionPostconditionSteps field to given value.

### HasSectionPostconditionSteps

`func (o *WorkItemModel) HasSectionPostconditionSteps() bool`

HasSectionPostconditionSteps returns a boolean if a field has been set.

### SetSectionPostconditionStepsNil

`func (o *WorkItemModel) SetSectionPostconditionStepsNil(b bool)`

 SetSectionPostconditionStepsNil sets the value for SectionPostconditionSteps to be an explicit nil

### UnsetSectionPostconditionSteps
`func (o *WorkItemModel) UnsetSectionPostconditionSteps()`

UnsetSectionPostconditionSteps ensures that no value is present for SectionPostconditionSteps, not even an explicit nil
### GetVersionNumber

`func (o *WorkItemModel) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *WorkItemModel) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *WorkItemModel) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *WorkItemModel) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetIterations

`func (o *WorkItemModel) GetIterations() []IterationModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *WorkItemModel) GetIterationsOk() (*[]IterationModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *WorkItemModel) SetIterations(v []IterationModel)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *WorkItemModel) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *WorkItemModel) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *WorkItemModel) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil
### GetCreatedDate

`func (o *WorkItemModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WorkItemModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *WorkItemModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WorkItemModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WorkItemModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *WorkItemModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *WorkItemModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *WorkItemModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkItemModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkItemModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WorkItemModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WorkItemModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WorkItemModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetGlobalId

`func (o *WorkItemModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *WorkItemModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetId

`func (o *WorkItemModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemModel) SetId(v string)`

SetId sets Id field to given value.


### GetSectionId

`func (o *WorkItemModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *WorkItemModel) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetDescription

`func (o *WorkItemModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkItemModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkItemModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkItemModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkItemModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkItemModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *WorkItemModel) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemModel) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemModel) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSteps

`func (o *WorkItemModel) GetSteps() []StepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemModel) GetStepsOk() (*[]StepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemModel) SetSteps(v []StepModel)`

SetSteps sets Steps field to given value.


### GetPreconditionSteps

`func (o *WorkItemModel) GetPreconditionSteps() []StepModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *WorkItemModel) GetPreconditionStepsOk() (*[]StepModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *WorkItemModel) SetPreconditionSteps(v []StepModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetPostconditionSteps

`func (o *WorkItemModel) GetPostconditionSteps() []StepModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *WorkItemModel) GetPostconditionStepsOk() (*[]StepModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *WorkItemModel) SetPostconditionSteps(v []StepModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetDuration

`func (o *WorkItemModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkItemModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAttributes

`func (o *WorkItemModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTags

`func (o *WorkItemModel) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemModel) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemModel) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *WorkItemModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.


### GetName

`func (o *WorkItemModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


