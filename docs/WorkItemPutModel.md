# WorkItemPutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]AttachmentPutModel**](AttachmentPutModel.md) |  | 
**Iterations** | Pointer to [**[]IterationPutModel**](IterationPutModel.md) |  | [optional] 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) |  | [optional] 
**Id** | **string** |  | 
**SectionId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**Steps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**PreconditionSteps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**PostconditionSteps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**Duration** | Pointer to **int32** |  | [optional] 
**Attributes** | **map[string]interface{}** |  | 
**Tags** | [**[]TagShortModel**](TagShortModel.md) |  | 
**Links** | [**[]LinkPutModel**](LinkPutModel.md) |  | 
**Name** | **string** |  | 

## Methods

### NewWorkItemPutModel

`func NewWorkItemPutModel(attachments []AttachmentPutModel, id string, state WorkItemStates, priority WorkItemPriorityModel, steps []StepPutModel, preconditionSteps []StepPutModel, postconditionSteps []StepPutModel, attributes map[string]interface{}, tags []TagShortModel, links []LinkPutModel, name string, ) *WorkItemPutModel`

NewWorkItemPutModel instantiates a new WorkItemPutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemPutModelWithDefaults

`func NewWorkItemPutModelWithDefaults() *WorkItemPutModel`

NewWorkItemPutModelWithDefaults instantiates a new WorkItemPutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *WorkItemPutModel) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkItemPutModel) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkItemPutModel) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.


### GetIterations

`func (o *WorkItemPutModel) GetIterations() []IterationPutModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *WorkItemPutModel) GetIterationsOk() (*[]IterationPutModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *WorkItemPutModel) SetIterations(v []IterationPutModel)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *WorkItemPutModel) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *WorkItemPutModel) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *WorkItemPutModel) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil
### GetAutoTests

`func (o *WorkItemPutModel) GetAutoTests() []AutoTestIdModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *WorkItemPutModel) GetAutoTestsOk() (*[]AutoTestIdModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *WorkItemPutModel) SetAutoTests(v []AutoTestIdModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *WorkItemPutModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *WorkItemPutModel) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *WorkItemPutModel) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil
### GetId

`func (o *WorkItemPutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemPutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemPutModel) SetId(v string)`

SetId sets Id field to given value.


### GetSectionId

`func (o *WorkItemPutModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemPutModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemPutModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *WorkItemPutModel) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetDescription

`func (o *WorkItemPutModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkItemPutModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkItemPutModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkItemPutModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkItemPutModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkItemPutModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *WorkItemPutModel) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemPutModel) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemPutModel) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemPutModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemPutModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemPutModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSteps

`func (o *WorkItemPutModel) GetSteps() []StepPutModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemPutModel) GetStepsOk() (*[]StepPutModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemPutModel) SetSteps(v []StepPutModel)`

SetSteps sets Steps field to given value.


### GetPreconditionSteps

`func (o *WorkItemPutModel) GetPreconditionSteps() []StepPutModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *WorkItemPutModel) GetPreconditionStepsOk() (*[]StepPutModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *WorkItemPutModel) SetPreconditionSteps(v []StepPutModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetPostconditionSteps

`func (o *WorkItemPutModel) GetPostconditionSteps() []StepPutModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *WorkItemPutModel) GetPostconditionStepsOk() (*[]StepPutModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *WorkItemPutModel) SetPostconditionSteps(v []StepPutModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetDuration

`func (o *WorkItemPutModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemPutModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemPutModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkItemPutModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAttributes

`func (o *WorkItemPutModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemPutModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemPutModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTags

`func (o *WorkItemPutModel) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemPutModel) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemPutModel) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *WorkItemPutModel) GetLinks() []LinkPutModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemPutModel) GetLinksOk() (*[]LinkPutModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemPutModel) SetLinks(v []LinkPutModel)`

SetLinks sets Links field to given value.


### GetName

`func (o *WorkItemPutModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemPutModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemPutModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


