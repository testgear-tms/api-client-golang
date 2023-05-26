# WorkItemPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityTypeName** | [**WorkItemEntityTypes**](WorkItemEntityTypes.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**Steps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**PreconditionSteps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**PostconditionSteps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**Duration** | **int32** | Must be 0 for shared steps and greater than 0 for the other types of work items | 
**Attributes** | **map[string]interface{}** |  | 
**Tags** | [**[]TagShortModel**](TagShortModel.md) |  | 
**Attachments** | Pointer to [**[]AttachmentPutModel**](AttachmentPutModel.md) |  | [optional] 
**Iterations** | Pointer to [**[]IterationPutModel**](IterationPutModel.md) |  | [optional] 
**Links** | [**[]LinkPostModel**](LinkPostModel.md) |  | 
**Name** | **string** |  | 
**ProjectId** | **string** | This property is used to link workitem with project | 
**SectionId** | **string** |  | 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) |  | [optional] 

## Methods

### NewWorkItemPostModel

`func NewWorkItemPostModel(entityTypeName WorkItemEntityTypes, state WorkItemStates, priority WorkItemPriorityModel, steps []StepPutModel, preconditionSteps []StepPutModel, postconditionSteps []StepPutModel, duration int32, attributes map[string]interface{}, tags []TagShortModel, links []LinkPostModel, name string, projectId string, sectionId string, ) *WorkItemPostModel`

NewWorkItemPostModel instantiates a new WorkItemPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemPostModelWithDefaults

`func NewWorkItemPostModelWithDefaults() *WorkItemPostModel`

NewWorkItemPostModelWithDefaults instantiates a new WorkItemPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityTypeName

`func (o *WorkItemPostModel) GetEntityTypeName() WorkItemEntityTypes`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemPostModel) GetEntityTypeNameOk() (*WorkItemEntityTypes, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemPostModel) SetEntityTypeName(v WorkItemEntityTypes)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetDescription

`func (o *WorkItemPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkItemPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkItemPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkItemPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkItemPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkItemPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *WorkItemPostModel) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemPostModel) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemPostModel) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemPostModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemPostModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemPostModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSteps

`func (o *WorkItemPostModel) GetSteps() []StepPutModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemPostModel) GetStepsOk() (*[]StepPutModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemPostModel) SetSteps(v []StepPutModel)`

SetSteps sets Steps field to given value.


### GetPreconditionSteps

`func (o *WorkItemPostModel) GetPreconditionSteps() []StepPutModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *WorkItemPostModel) GetPreconditionStepsOk() (*[]StepPutModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *WorkItemPostModel) SetPreconditionSteps(v []StepPutModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetPostconditionSteps

`func (o *WorkItemPostModel) GetPostconditionSteps() []StepPutModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *WorkItemPostModel) GetPostconditionStepsOk() (*[]StepPutModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *WorkItemPostModel) SetPostconditionSteps(v []StepPutModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetDuration

`func (o *WorkItemPostModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemPostModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemPostModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetAttributes

`func (o *WorkItemPostModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemPostModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemPostModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTags

`func (o *WorkItemPostModel) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemPostModel) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemPostModel) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.


### GetAttachments

`func (o *WorkItemPostModel) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkItemPostModel) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkItemPostModel) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *WorkItemPostModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *WorkItemPostModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *WorkItemPostModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetIterations

`func (o *WorkItemPostModel) GetIterations() []IterationPutModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *WorkItemPostModel) GetIterationsOk() (*[]IterationPutModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *WorkItemPostModel) SetIterations(v []IterationPutModel)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *WorkItemPostModel) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *WorkItemPostModel) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *WorkItemPostModel) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil
### GetLinks

`func (o *WorkItemPostModel) GetLinks() []LinkPostModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemPostModel) GetLinksOk() (*[]LinkPostModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemPostModel) SetLinks(v []LinkPostModel)`

SetLinks sets Links field to given value.


### GetName

`func (o *WorkItemPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemPostModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *WorkItemPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSectionId

`func (o *WorkItemPostModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemPostModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemPostModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetAutoTests

`func (o *WorkItemPostModel) GetAutoTests() []AutoTestIdModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *WorkItemPostModel) GetAutoTestsOk() (*[]AutoTestIdModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *WorkItemPostModel) SetAutoTests(v []AutoTestIdModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *WorkItemPostModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *WorkItemPostModel) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *WorkItemPostModel) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


