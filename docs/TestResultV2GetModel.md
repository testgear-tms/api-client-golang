# TestResultV2GetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**ConfigurationModel**](ConfigurationModel.md) |  | [optional] 
**AutoTest** | Pointer to [**AutoTestModelV2GetModel**](AutoTestModelV2GetModel.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**WorkItemVersionId** | Pointer to **string** |  | [optional] 
**AutoTestId** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Traces** | Pointer to **NullableString** |  | [optional] 
**StartedOn** | Pointer to **NullableTime** |  | [optional] 
**CompletedOn** | Pointer to **NullableTime** |  | [optional] 
**RunByUserId** | Pointer to **NullableString** |  | [optional] 
**StoppedByUserId** | Pointer to **NullableString** |  | [optional] 
**TestPointId** | Pointer to **NullableString** |  | [optional] 
**TestPoint** | Pointer to [**TestPointShortModel**](TestPointShortModel.md) |  | [optional] 
**TestRunId** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **NullableString** | Property can contain one of these values: Passed, Failed, InProgress, Blocked, Skipped | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTestResultV2GetModel

`func NewTestResultV2GetModel() *TestResultV2GetModel`

NewTestResultV2GetModel instantiates a new TestResultV2GetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultV2GetModelWithDefaults

`func NewTestResultV2GetModelWithDefaults() *TestResultV2GetModel`

NewTestResultV2GetModelWithDefaults instantiates a new TestResultV2GetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *TestResultV2GetModel) GetConfiguration() ConfigurationModel`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TestResultV2GetModel) GetConfigurationOk() (*ConfigurationModel, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TestResultV2GetModel) SetConfiguration(v ConfigurationModel)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *TestResultV2GetModel) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAutoTest

`func (o *TestResultV2GetModel) GetAutoTest() AutoTestModelV2GetModel`

GetAutoTest returns the AutoTest field if non-nil, zero value otherwise.

### GetAutoTestOk

`func (o *TestResultV2GetModel) GetAutoTestOk() (*AutoTestModelV2GetModel, bool)`

GetAutoTestOk returns a tuple with the AutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTest

`func (o *TestResultV2GetModel) SetAutoTest(v AutoTestModelV2GetModel)`

SetAutoTest sets AutoTest field to given value.

### HasAutoTest

`func (o *TestResultV2GetModel) HasAutoTest() bool`

HasAutoTest returns a boolean if a field has been set.

### GetId

`func (o *TestResultV2GetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultV2GetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultV2GetModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestResultV2GetModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationId

`func (o *TestResultV2GetModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestResultV2GetModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestResultV2GetModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestResultV2GetModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetWorkItemVersionId

`func (o *TestResultV2GetModel) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultV2GetModel) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultV2GetModel) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestResultV2GetModel) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### GetAutoTestId

`func (o *TestResultV2GetModel) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *TestResultV2GetModel) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *TestResultV2GetModel) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *TestResultV2GetModel) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *TestResultV2GetModel) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *TestResultV2GetModel) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetMessage

`func (o *TestResultV2GetModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultV2GetModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultV2GetModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultV2GetModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultV2GetModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultV2GetModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTraces

`func (o *TestResultV2GetModel) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *TestResultV2GetModel) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *TestResultV2GetModel) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *TestResultV2GetModel) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *TestResultV2GetModel) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *TestResultV2GetModel) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetStartedOn

`func (o *TestResultV2GetModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultV2GetModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultV2GetModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultV2GetModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultV2GetModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultV2GetModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultV2GetModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultV2GetModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultV2GetModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultV2GetModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultV2GetModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultV2GetModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetRunByUserId

`func (o *TestResultV2GetModel) GetRunByUserId() string`

GetRunByUserId returns the RunByUserId field if non-nil, zero value otherwise.

### GetRunByUserIdOk

`func (o *TestResultV2GetModel) GetRunByUserIdOk() (*string, bool)`

GetRunByUserIdOk returns a tuple with the RunByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByUserId

`func (o *TestResultV2GetModel) SetRunByUserId(v string)`

SetRunByUserId sets RunByUserId field to given value.

### HasRunByUserId

`func (o *TestResultV2GetModel) HasRunByUserId() bool`

HasRunByUserId returns a boolean if a field has been set.

### SetRunByUserIdNil

`func (o *TestResultV2GetModel) SetRunByUserIdNil(b bool)`

 SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil

### UnsetRunByUserId
`func (o *TestResultV2GetModel) UnsetRunByUserId()`

UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
### GetStoppedByUserId

`func (o *TestResultV2GetModel) GetStoppedByUserId() string`

GetStoppedByUserId returns the StoppedByUserId field if non-nil, zero value otherwise.

### GetStoppedByUserIdOk

`func (o *TestResultV2GetModel) GetStoppedByUserIdOk() (*string, bool)`

GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedByUserId

`func (o *TestResultV2GetModel) SetStoppedByUserId(v string)`

SetStoppedByUserId sets StoppedByUserId field to given value.

### HasStoppedByUserId

`func (o *TestResultV2GetModel) HasStoppedByUserId() bool`

HasStoppedByUserId returns a boolean if a field has been set.

### SetStoppedByUserIdNil

`func (o *TestResultV2GetModel) SetStoppedByUserIdNil(b bool)`

 SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil

### UnsetStoppedByUserId
`func (o *TestResultV2GetModel) UnsetStoppedByUserId()`

UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
### GetTestPointId

`func (o *TestResultV2GetModel) GetTestPointId() string`

GetTestPointId returns the TestPointId field if non-nil, zero value otherwise.

### GetTestPointIdOk

`func (o *TestResultV2GetModel) GetTestPointIdOk() (*string, bool)`

GetTestPointIdOk returns a tuple with the TestPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointId

`func (o *TestResultV2GetModel) SetTestPointId(v string)`

SetTestPointId sets TestPointId field to given value.

### HasTestPointId

`func (o *TestResultV2GetModel) HasTestPointId() bool`

HasTestPointId returns a boolean if a field has been set.

### SetTestPointIdNil

`func (o *TestResultV2GetModel) SetTestPointIdNil(b bool)`

 SetTestPointIdNil sets the value for TestPointId to be an explicit nil

### UnsetTestPointId
`func (o *TestResultV2GetModel) UnsetTestPointId()`

UnsetTestPointId ensures that no value is present for TestPointId, not even an explicit nil
### GetTestPoint

`func (o *TestResultV2GetModel) GetTestPoint() TestPointShortModel`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *TestResultV2GetModel) GetTestPointOk() (*TestPointShortModel, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *TestResultV2GetModel) SetTestPoint(v TestPointShortModel)`

SetTestPoint sets TestPoint field to given value.

### HasTestPoint

`func (o *TestResultV2GetModel) HasTestPoint() bool`

HasTestPoint returns a boolean if a field has been set.

### GetTestRunId

`func (o *TestResultV2GetModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultV2GetModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultV2GetModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *TestResultV2GetModel) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### GetOutcome

`func (o *TestResultV2GetModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultV2GetModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultV2GetModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultV2GetModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultV2GetModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultV2GetModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetComment

`func (o *TestResultV2GetModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultV2GetModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultV2GetModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultV2GetModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultV2GetModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultV2GetModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultV2GetModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultV2GetModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultV2GetModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultV2GetModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultV2GetModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultV2GetModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetAttachments

`func (o *TestResultV2GetModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultV2GetModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultV2GetModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultV2GetModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultV2GetModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultV2GetModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetParameters

`func (o *TestResultV2GetModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultV2GetModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultV2GetModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultV2GetModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultV2GetModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultV2GetModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetProperties

`func (o *TestResultV2GetModel) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TestResultV2GetModel) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TestResultV2GetModel) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TestResultV2GetModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *TestResultV2GetModel) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *TestResultV2GetModel) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


