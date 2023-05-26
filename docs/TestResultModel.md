# TestResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTestId** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**StartedOn** | Pointer to **NullableTime** |  | [optional] 
**CompletedOn** | Pointer to **NullableTime** |  | [optional] 
**DurationInMs** | Pointer to **NullableInt64** |  | [optional] 
**Traces** | Pointer to **NullableString** |  | [optional] 
**FailureType** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**RunByUserId** | Pointer to **NullableString** |  | [optional] 
**StoppedByUserId** | Pointer to **NullableString** |  | [optional] 
**TestPointId** | Pointer to **string** |  | [optional] 
**TestRunId** | Pointer to **string** |  | [optional] 
**TestPoint** | Pointer to [**TestPointPutModel**](TestPointPutModel.md) |  | [optional] 
**AutoTest** | Pointer to [**AutoTestModel**](AutoTestModel.md) |  | [optional] 
**AutoTestStepResults** | Pointer to [**[]AttachmentModelAutoTestStepResultsModel**](AttachmentModelAutoTestStepResultsModel.md) |  | [optional] 
**SetupResults** | Pointer to [**[]AttachmentModelAutoTestStepResultsModel**](AttachmentModelAutoTestStepResultsModel.md) |  | [optional] 
**TeardownResults** | Pointer to [**[]AttachmentModelAutoTestStepResultsModel**](AttachmentModelAutoTestStepResultsModel.md) |  | [optional] 
**WorkItemVersionId** | Pointer to **string** |  | [optional] 
**WorkItemVersionNumber** | Pointer to **NullableInt32** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**StepComments** | Pointer to [**[]StepCommentModel**](StepCommentModel.md) |  | [optional] 
**FailureClassIds** | Pointer to **[]string** |  | [optional] 
**Outcome** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**StepResults** | Pointer to [**[]StepResultModel**](StepResultModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 

## Methods

### NewTestResultModel

`func NewTestResultModel() *TestResultModel`

NewTestResultModel instantiates a new TestResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultModelWithDefaults

`func NewTestResultModelWithDefaults() *TestResultModel`

NewTestResultModelWithDefaults instantiates a new TestResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTestId

`func (o *TestResultModel) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *TestResultModel) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *TestResultModel) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *TestResultModel) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *TestResultModel) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *TestResultModel) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetConfigurationId

`func (o *TestResultModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestResultModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestResultModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestResultModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetStartedOn

`func (o *TestResultModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDurationInMs

`func (o *TestResultModel) GetDurationInMs() int64`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *TestResultModel) GetDurationInMsOk() (*int64, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *TestResultModel) SetDurationInMs(v int64)`

SetDurationInMs sets DurationInMs field to given value.

### HasDurationInMs

`func (o *TestResultModel) HasDurationInMs() bool`

HasDurationInMs returns a boolean if a field has been set.

### SetDurationInMsNil

`func (o *TestResultModel) SetDurationInMsNil(b bool)`

 SetDurationInMsNil sets the value for DurationInMs to be an explicit nil

### UnsetDurationInMs
`func (o *TestResultModel) UnsetDurationInMs()`

UnsetDurationInMs ensures that no value is present for DurationInMs, not even an explicit nil
### GetTraces

`func (o *TestResultModel) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *TestResultModel) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *TestResultModel) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *TestResultModel) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *TestResultModel) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *TestResultModel) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetFailureType

`func (o *TestResultModel) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *TestResultModel) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *TestResultModel) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *TestResultModel) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### SetFailureTypeNil

`func (o *TestResultModel) SetFailureTypeNil(b bool)`

 SetFailureTypeNil sets the value for FailureType to be an explicit nil

### UnsetFailureType
`func (o *TestResultModel) UnsetFailureType()`

UnsetFailureType ensures that no value is present for FailureType, not even an explicit nil
### GetMessage

`func (o *TestResultModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRunByUserId

`func (o *TestResultModel) GetRunByUserId() string`

GetRunByUserId returns the RunByUserId field if non-nil, zero value otherwise.

### GetRunByUserIdOk

`func (o *TestResultModel) GetRunByUserIdOk() (*string, bool)`

GetRunByUserIdOk returns a tuple with the RunByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByUserId

`func (o *TestResultModel) SetRunByUserId(v string)`

SetRunByUserId sets RunByUserId field to given value.

### HasRunByUserId

`func (o *TestResultModel) HasRunByUserId() bool`

HasRunByUserId returns a boolean if a field has been set.

### SetRunByUserIdNil

`func (o *TestResultModel) SetRunByUserIdNil(b bool)`

 SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil

### UnsetRunByUserId
`func (o *TestResultModel) UnsetRunByUserId()`

UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
### GetStoppedByUserId

`func (o *TestResultModel) GetStoppedByUserId() string`

GetStoppedByUserId returns the StoppedByUserId field if non-nil, zero value otherwise.

### GetStoppedByUserIdOk

`func (o *TestResultModel) GetStoppedByUserIdOk() (*string, bool)`

GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedByUserId

`func (o *TestResultModel) SetStoppedByUserId(v string)`

SetStoppedByUserId sets StoppedByUserId field to given value.

### HasStoppedByUserId

`func (o *TestResultModel) HasStoppedByUserId() bool`

HasStoppedByUserId returns a boolean if a field has been set.

### SetStoppedByUserIdNil

`func (o *TestResultModel) SetStoppedByUserIdNil(b bool)`

 SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil

### UnsetStoppedByUserId
`func (o *TestResultModel) UnsetStoppedByUserId()`

UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
### GetTestPointId

`func (o *TestResultModel) GetTestPointId() string`

GetTestPointId returns the TestPointId field if non-nil, zero value otherwise.

### GetTestPointIdOk

`func (o *TestResultModel) GetTestPointIdOk() (*string, bool)`

GetTestPointIdOk returns a tuple with the TestPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointId

`func (o *TestResultModel) SetTestPointId(v string)`

SetTestPointId sets TestPointId field to given value.

### HasTestPointId

`func (o *TestResultModel) HasTestPointId() bool`

HasTestPointId returns a boolean if a field has been set.

### GetTestRunId

`func (o *TestResultModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *TestResultModel) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### GetTestPoint

`func (o *TestResultModel) GetTestPoint() TestPointPutModel`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *TestResultModel) GetTestPointOk() (*TestPointPutModel, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *TestResultModel) SetTestPoint(v TestPointPutModel)`

SetTestPoint sets TestPoint field to given value.

### HasTestPoint

`func (o *TestResultModel) HasTestPoint() bool`

HasTestPoint returns a boolean if a field has been set.

### GetAutoTest

`func (o *TestResultModel) GetAutoTest() AutoTestModel`

GetAutoTest returns the AutoTest field if non-nil, zero value otherwise.

### GetAutoTestOk

`func (o *TestResultModel) GetAutoTestOk() (*AutoTestModel, bool)`

GetAutoTestOk returns a tuple with the AutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTest

`func (o *TestResultModel) SetAutoTest(v AutoTestModel)`

SetAutoTest sets AutoTest field to given value.

### HasAutoTest

`func (o *TestResultModel) HasAutoTest() bool`

HasAutoTest returns a boolean if a field has been set.

### GetAutoTestStepResults

`func (o *TestResultModel) GetAutoTestStepResults() []AttachmentModelAutoTestStepResultsModel`

GetAutoTestStepResults returns the AutoTestStepResults field if non-nil, zero value otherwise.

### GetAutoTestStepResultsOk

`func (o *TestResultModel) GetAutoTestStepResultsOk() (*[]AttachmentModelAutoTestStepResultsModel, bool)`

GetAutoTestStepResultsOk returns a tuple with the AutoTestStepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestStepResults

`func (o *TestResultModel) SetAutoTestStepResults(v []AttachmentModelAutoTestStepResultsModel)`

SetAutoTestStepResults sets AutoTestStepResults field to given value.

### HasAutoTestStepResults

`func (o *TestResultModel) HasAutoTestStepResults() bool`

HasAutoTestStepResults returns a boolean if a field has been set.

### SetAutoTestStepResultsNil

`func (o *TestResultModel) SetAutoTestStepResultsNil(b bool)`

 SetAutoTestStepResultsNil sets the value for AutoTestStepResults to be an explicit nil

### UnsetAutoTestStepResults
`func (o *TestResultModel) UnsetAutoTestStepResults()`

UnsetAutoTestStepResults ensures that no value is present for AutoTestStepResults, not even an explicit nil
### GetSetupResults

`func (o *TestResultModel) GetSetupResults() []AttachmentModelAutoTestStepResultsModel`

GetSetupResults returns the SetupResults field if non-nil, zero value otherwise.

### GetSetupResultsOk

`func (o *TestResultModel) GetSetupResultsOk() (*[]AttachmentModelAutoTestStepResultsModel, bool)`

GetSetupResultsOk returns a tuple with the SetupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupResults

`func (o *TestResultModel) SetSetupResults(v []AttachmentModelAutoTestStepResultsModel)`

SetSetupResults sets SetupResults field to given value.

### HasSetupResults

`func (o *TestResultModel) HasSetupResults() bool`

HasSetupResults returns a boolean if a field has been set.

### SetSetupResultsNil

`func (o *TestResultModel) SetSetupResultsNil(b bool)`

 SetSetupResultsNil sets the value for SetupResults to be an explicit nil

### UnsetSetupResults
`func (o *TestResultModel) UnsetSetupResults()`

UnsetSetupResults ensures that no value is present for SetupResults, not even an explicit nil
### GetTeardownResults

`func (o *TestResultModel) GetTeardownResults() []AttachmentModelAutoTestStepResultsModel`

GetTeardownResults returns the TeardownResults field if non-nil, zero value otherwise.

### GetTeardownResultsOk

`func (o *TestResultModel) GetTeardownResultsOk() (*[]AttachmentModelAutoTestStepResultsModel, bool)`

GetTeardownResultsOk returns a tuple with the TeardownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownResults

`func (o *TestResultModel) SetTeardownResults(v []AttachmentModelAutoTestStepResultsModel)`

SetTeardownResults sets TeardownResults field to given value.

### HasTeardownResults

`func (o *TestResultModel) HasTeardownResults() bool`

HasTeardownResults returns a boolean if a field has been set.

### SetTeardownResultsNil

`func (o *TestResultModel) SetTeardownResultsNil(b bool)`

 SetTeardownResultsNil sets the value for TeardownResults to be an explicit nil

### UnsetTeardownResults
`func (o *TestResultModel) UnsetTeardownResults()`

UnsetTeardownResults ensures that no value is present for TeardownResults, not even an explicit nil
### GetWorkItemVersionId

`func (o *TestResultModel) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultModel) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultModel) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestResultModel) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### GetWorkItemVersionNumber

`func (o *TestResultModel) GetWorkItemVersionNumber() int32`

GetWorkItemVersionNumber returns the WorkItemVersionNumber field if non-nil, zero value otherwise.

### GetWorkItemVersionNumberOk

`func (o *TestResultModel) GetWorkItemVersionNumberOk() (*int32, bool)`

GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionNumber

`func (o *TestResultModel) SetWorkItemVersionNumber(v int32)`

SetWorkItemVersionNumber sets WorkItemVersionNumber field to given value.

### HasWorkItemVersionNumber

`func (o *TestResultModel) HasWorkItemVersionNumber() bool`

HasWorkItemVersionNumber returns a boolean if a field has been set.

### SetWorkItemVersionNumberNil

`func (o *TestResultModel) SetWorkItemVersionNumberNil(b bool)`

 SetWorkItemVersionNumberNil sets the value for WorkItemVersionNumber to be an explicit nil

### UnsetWorkItemVersionNumber
`func (o *TestResultModel) UnsetWorkItemVersionNumber()`

UnsetWorkItemVersionNumber ensures that no value is present for WorkItemVersionNumber, not even an explicit nil
### GetParameters

`func (o *TestResultModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetProperties

`func (o *TestResultModel) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TestResultModel) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TestResultModel) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TestResultModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *TestResultModel) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *TestResultModel) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetId

`func (o *TestResultModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestResultModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TestResultModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestResultModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *TestResultModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestResultModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestResultModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestResultModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestResultModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestResultModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestResultModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *TestResultModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestResultModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestResultModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestResultModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestResultModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestResultModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetStepComments

`func (o *TestResultModel) GetStepComments() []StepCommentModel`

GetStepComments returns the StepComments field if non-nil, zero value otherwise.

### GetStepCommentsOk

`func (o *TestResultModel) GetStepCommentsOk() (*[]StepCommentModel, bool)`

GetStepCommentsOk returns a tuple with the StepComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepComments

`func (o *TestResultModel) SetStepComments(v []StepCommentModel)`

SetStepComments sets StepComments field to given value.

### HasStepComments

`func (o *TestResultModel) HasStepComments() bool`

HasStepComments returns a boolean if a field has been set.

### SetStepCommentsNil

`func (o *TestResultModel) SetStepCommentsNil(b bool)`

 SetStepCommentsNil sets the value for StepComments to be an explicit nil

### UnsetStepComments
`func (o *TestResultModel) UnsetStepComments()`

UnsetStepComments ensures that no value is present for StepComments, not even an explicit nil
### GetFailureClassIds

`func (o *TestResultModel) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *TestResultModel) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *TestResultModel) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.

### HasFailureClassIds

`func (o *TestResultModel) HasFailureClassIds() bool`

HasFailureClassIds returns a boolean if a field has been set.

### SetFailureClassIdsNil

`func (o *TestResultModel) SetFailureClassIdsNil(b bool)`

 SetFailureClassIdsNil sets the value for FailureClassIds to be an explicit nil

### UnsetFailureClassIds
`func (o *TestResultModel) UnsetFailureClassIds()`

UnsetFailureClassIds ensures that no value is present for FailureClassIds, not even an explicit nil
### GetOutcome

`func (o *TestResultModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetComment

`func (o *TestResultModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStepResults

`func (o *TestResultModel) GetStepResults() []StepResultModel`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *TestResultModel) GetStepResultsOk() (*[]StepResultModel, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *TestResultModel) SetStepResults(v []StepResultModel)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *TestResultModel) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *TestResultModel) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *TestResultModel) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *TestResultModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


