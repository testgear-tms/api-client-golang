# TestResultV2ShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewTestResultV2ShortModel

`func NewTestResultV2ShortModel() *TestResultV2ShortModel`

NewTestResultV2ShortModel instantiates a new TestResultV2ShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultV2ShortModelWithDefaults

`func NewTestResultV2ShortModelWithDefaults() *TestResultV2ShortModel`

NewTestResultV2ShortModelWithDefaults instantiates a new TestResultV2ShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultV2ShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultV2ShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultV2ShortModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestResultV2ShortModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationId

`func (o *TestResultV2ShortModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestResultV2ShortModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestResultV2ShortModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestResultV2ShortModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetWorkItemVersionId

`func (o *TestResultV2ShortModel) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultV2ShortModel) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultV2ShortModel) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestResultV2ShortModel) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### GetAutoTestId

`func (o *TestResultV2ShortModel) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *TestResultV2ShortModel) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *TestResultV2ShortModel) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *TestResultV2ShortModel) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *TestResultV2ShortModel) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *TestResultV2ShortModel) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetMessage

`func (o *TestResultV2ShortModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultV2ShortModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultV2ShortModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultV2ShortModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultV2ShortModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultV2ShortModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTraces

`func (o *TestResultV2ShortModel) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *TestResultV2ShortModel) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *TestResultV2ShortModel) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *TestResultV2ShortModel) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *TestResultV2ShortModel) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *TestResultV2ShortModel) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetStartedOn

`func (o *TestResultV2ShortModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultV2ShortModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultV2ShortModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultV2ShortModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultV2ShortModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultV2ShortModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultV2ShortModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultV2ShortModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultV2ShortModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultV2ShortModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultV2ShortModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultV2ShortModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetRunByUserId

`func (o *TestResultV2ShortModel) GetRunByUserId() string`

GetRunByUserId returns the RunByUserId field if non-nil, zero value otherwise.

### GetRunByUserIdOk

`func (o *TestResultV2ShortModel) GetRunByUserIdOk() (*string, bool)`

GetRunByUserIdOk returns a tuple with the RunByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByUserId

`func (o *TestResultV2ShortModel) SetRunByUserId(v string)`

SetRunByUserId sets RunByUserId field to given value.

### HasRunByUserId

`func (o *TestResultV2ShortModel) HasRunByUserId() bool`

HasRunByUserId returns a boolean if a field has been set.

### SetRunByUserIdNil

`func (o *TestResultV2ShortModel) SetRunByUserIdNil(b bool)`

 SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil

### UnsetRunByUserId
`func (o *TestResultV2ShortModel) UnsetRunByUserId()`

UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
### GetStoppedByUserId

`func (o *TestResultV2ShortModel) GetStoppedByUserId() string`

GetStoppedByUserId returns the StoppedByUserId field if non-nil, zero value otherwise.

### GetStoppedByUserIdOk

`func (o *TestResultV2ShortModel) GetStoppedByUserIdOk() (*string, bool)`

GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedByUserId

`func (o *TestResultV2ShortModel) SetStoppedByUserId(v string)`

SetStoppedByUserId sets StoppedByUserId field to given value.

### HasStoppedByUserId

`func (o *TestResultV2ShortModel) HasStoppedByUserId() bool`

HasStoppedByUserId returns a boolean if a field has been set.

### SetStoppedByUserIdNil

`func (o *TestResultV2ShortModel) SetStoppedByUserIdNil(b bool)`

 SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil

### UnsetStoppedByUserId
`func (o *TestResultV2ShortModel) UnsetStoppedByUserId()`

UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
### GetTestPointId

`func (o *TestResultV2ShortModel) GetTestPointId() string`

GetTestPointId returns the TestPointId field if non-nil, zero value otherwise.

### GetTestPointIdOk

`func (o *TestResultV2ShortModel) GetTestPointIdOk() (*string, bool)`

GetTestPointIdOk returns a tuple with the TestPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointId

`func (o *TestResultV2ShortModel) SetTestPointId(v string)`

SetTestPointId sets TestPointId field to given value.

### HasTestPointId

`func (o *TestResultV2ShortModel) HasTestPointId() bool`

HasTestPointId returns a boolean if a field has been set.

### SetTestPointIdNil

`func (o *TestResultV2ShortModel) SetTestPointIdNil(b bool)`

 SetTestPointIdNil sets the value for TestPointId to be an explicit nil

### UnsetTestPointId
`func (o *TestResultV2ShortModel) UnsetTestPointId()`

UnsetTestPointId ensures that no value is present for TestPointId, not even an explicit nil
### GetTestPoint

`func (o *TestResultV2ShortModel) GetTestPoint() TestPointShortModel`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *TestResultV2ShortModel) GetTestPointOk() (*TestPointShortModel, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *TestResultV2ShortModel) SetTestPoint(v TestPointShortModel)`

SetTestPoint sets TestPoint field to given value.

### HasTestPoint

`func (o *TestResultV2ShortModel) HasTestPoint() bool`

HasTestPoint returns a boolean if a field has been set.

### GetTestRunId

`func (o *TestResultV2ShortModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultV2ShortModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultV2ShortModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *TestResultV2ShortModel) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### GetOutcome

`func (o *TestResultV2ShortModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultV2ShortModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultV2ShortModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultV2ShortModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultV2ShortModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultV2ShortModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetComment

`func (o *TestResultV2ShortModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultV2ShortModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultV2ShortModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultV2ShortModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultV2ShortModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultV2ShortModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultV2ShortModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultV2ShortModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultV2ShortModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultV2ShortModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultV2ShortModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultV2ShortModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetAttachments

`func (o *TestResultV2ShortModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultV2ShortModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultV2ShortModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultV2ShortModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultV2ShortModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultV2ShortModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetParameters

`func (o *TestResultV2ShortModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultV2ShortModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultV2ShortModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultV2ShortModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultV2ShortModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultV2ShortModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetProperties

`func (o *TestResultV2ShortModel) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TestResultV2ShortModel) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TestResultV2ShortModel) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TestResultV2ShortModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *TestResultV2ShortModel) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *TestResultV2ShortModel) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


