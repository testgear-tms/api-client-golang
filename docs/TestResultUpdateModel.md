# TestResultUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetupResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) |  | [optional] 
**TeardownResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) |  | [optional] 
**DurationInMs** | Pointer to **NullableInt64** |  | [optional] 
**StepComments** | Pointer to [**[]TestResultStepCommentPutModel**](TestResultStepCommentPutModel.md) |  | [optional] 
**FailureClassIds** | Pointer to **[]string** |  | [optional] 
**Outcome** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**StepResults** | Pointer to [**[]StepResultModel**](StepResultModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentPutModel**](AttachmentPutModel.md) |  | [optional] 

## Methods

### NewTestResultUpdateModel

`func NewTestResultUpdateModel() *TestResultUpdateModel`

NewTestResultUpdateModel instantiates a new TestResultUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultUpdateModelWithDefaults

`func NewTestResultUpdateModelWithDefaults() *TestResultUpdateModel`

NewTestResultUpdateModelWithDefaults instantiates a new TestResultUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetupResults

`func (o *TestResultUpdateModel) GetSetupResults() []AttachmentPutModelAutoTestStepResultsModel`

GetSetupResults returns the SetupResults field if non-nil, zero value otherwise.

### GetSetupResultsOk

`func (o *TestResultUpdateModel) GetSetupResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetSetupResultsOk returns a tuple with the SetupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupResults

`func (o *TestResultUpdateModel) SetSetupResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetSetupResults sets SetupResults field to given value.

### HasSetupResults

`func (o *TestResultUpdateModel) HasSetupResults() bool`

HasSetupResults returns a boolean if a field has been set.

### SetSetupResultsNil

`func (o *TestResultUpdateModel) SetSetupResultsNil(b bool)`

 SetSetupResultsNil sets the value for SetupResults to be an explicit nil

### UnsetSetupResults
`func (o *TestResultUpdateModel) UnsetSetupResults()`

UnsetSetupResults ensures that no value is present for SetupResults, not even an explicit nil
### GetTeardownResults

`func (o *TestResultUpdateModel) GetTeardownResults() []AttachmentPutModelAutoTestStepResultsModel`

GetTeardownResults returns the TeardownResults field if non-nil, zero value otherwise.

### GetTeardownResultsOk

`func (o *TestResultUpdateModel) GetTeardownResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetTeardownResultsOk returns a tuple with the TeardownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownResults

`func (o *TestResultUpdateModel) SetTeardownResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetTeardownResults sets TeardownResults field to given value.

### HasTeardownResults

`func (o *TestResultUpdateModel) HasTeardownResults() bool`

HasTeardownResults returns a boolean if a field has been set.

### SetTeardownResultsNil

`func (o *TestResultUpdateModel) SetTeardownResultsNil(b bool)`

 SetTeardownResultsNil sets the value for TeardownResults to be an explicit nil

### UnsetTeardownResults
`func (o *TestResultUpdateModel) UnsetTeardownResults()`

UnsetTeardownResults ensures that no value is present for TeardownResults, not even an explicit nil
### GetDurationInMs

`func (o *TestResultUpdateModel) GetDurationInMs() int64`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *TestResultUpdateModel) GetDurationInMsOk() (*int64, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *TestResultUpdateModel) SetDurationInMs(v int64)`

SetDurationInMs sets DurationInMs field to given value.

### HasDurationInMs

`func (o *TestResultUpdateModel) HasDurationInMs() bool`

HasDurationInMs returns a boolean if a field has been set.

### SetDurationInMsNil

`func (o *TestResultUpdateModel) SetDurationInMsNil(b bool)`

 SetDurationInMsNil sets the value for DurationInMs to be an explicit nil

### UnsetDurationInMs
`func (o *TestResultUpdateModel) UnsetDurationInMs()`

UnsetDurationInMs ensures that no value is present for DurationInMs, not even an explicit nil
### GetStepComments

`func (o *TestResultUpdateModel) GetStepComments() []TestResultStepCommentPutModel`

GetStepComments returns the StepComments field if non-nil, zero value otherwise.

### GetStepCommentsOk

`func (o *TestResultUpdateModel) GetStepCommentsOk() (*[]TestResultStepCommentPutModel, bool)`

GetStepCommentsOk returns a tuple with the StepComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepComments

`func (o *TestResultUpdateModel) SetStepComments(v []TestResultStepCommentPutModel)`

SetStepComments sets StepComments field to given value.

### HasStepComments

`func (o *TestResultUpdateModel) HasStepComments() bool`

HasStepComments returns a boolean if a field has been set.

### SetStepCommentsNil

`func (o *TestResultUpdateModel) SetStepCommentsNil(b bool)`

 SetStepCommentsNil sets the value for StepComments to be an explicit nil

### UnsetStepComments
`func (o *TestResultUpdateModel) UnsetStepComments()`

UnsetStepComments ensures that no value is present for StepComments, not even an explicit nil
### GetFailureClassIds

`func (o *TestResultUpdateModel) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *TestResultUpdateModel) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *TestResultUpdateModel) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.

### HasFailureClassIds

`func (o *TestResultUpdateModel) HasFailureClassIds() bool`

HasFailureClassIds returns a boolean if a field has been set.

### SetFailureClassIdsNil

`func (o *TestResultUpdateModel) SetFailureClassIdsNil(b bool)`

 SetFailureClassIdsNil sets the value for FailureClassIds to be an explicit nil

### UnsetFailureClassIds
`func (o *TestResultUpdateModel) UnsetFailureClassIds()`

UnsetFailureClassIds ensures that no value is present for FailureClassIds, not even an explicit nil
### GetOutcome

`func (o *TestResultUpdateModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultUpdateModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultUpdateModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultUpdateModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultUpdateModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultUpdateModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetComment

`func (o *TestResultUpdateModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultUpdateModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultUpdateModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultUpdateModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultUpdateModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultUpdateModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultUpdateModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultUpdateModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultUpdateModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultUpdateModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultUpdateModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultUpdateModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStepResults

`func (o *TestResultUpdateModel) GetStepResults() []StepResultModel`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *TestResultUpdateModel) GetStepResultsOk() (*[]StepResultModel, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *TestResultUpdateModel) SetStepResults(v []StepResultModel)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *TestResultUpdateModel) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *TestResultUpdateModel) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *TestResultUpdateModel) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *TestResultUpdateModel) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultUpdateModel) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultUpdateModel) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultUpdateModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultUpdateModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultUpdateModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


