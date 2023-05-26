# AutoTestResultsForTestRunModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | **string** | Specifies the GUID of the autotest configuration, which was specified when the test run was created. | 
**Links** | Pointer to [**[]LinkPostModel**](LinkPostModel.md) | Specifies the links in the autotest. | [optional] 
**FailureReasonNames** | Pointer to [**[]FailureCategoryModel**](FailureCategoryModel.md) | Specifies the cause of autotest failure. | [optional] 
**AutoTestExternalId** | **string** | Specifies the external ID of the autotest, which was specified when the test run was created. | 
**Outcome** | [**AvailableTestResultOutcome**](AvailableTestResultOutcome.md) |  | 
**Message** | Pointer to **NullableString** | A comment for the result. | [optional] 
**Traces** | Pointer to **NullableString** | An extended comment or a stack trace. | [optional] 
**StartedOn** | Pointer to **NullableTime** | Test run start date. | [optional] 
**CompletedOn** | Pointer to **NullableTime** | Test run end date. | [optional] 
**Duration** | Pointer to **NullableInt64** | Expected or actual duration of the test run execution in milliseconds. | [optional] 
**Attachments** | Pointer to [**[]AttachmentPutModel**](AttachmentPutModel.md) | Specifies an attachment GUID. Multiple values can be sent. | [optional] 
**Parameters** | Pointer to **map[string]string** | \&quot;&lt;b&gt;parameter&lt;/b&gt;\&quot;: \&quot;&lt;b&gt;value&lt;/b&gt;\&quot; pair with arbitrary custom parameters. Multiple parameters can be sent. | [optional] 
**Properties** | Pointer to **map[string]string** | \&quot;&lt;b&gt;property&lt;/b&gt;\&quot;: \&quot;&lt;b&gt;value&lt;/b&gt;\&quot; pair with arbitrary custom properties. Multiple properties can be sent. | [optional] 
**StepResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) | Specifies the results of individual steps. | [optional] 
**SetupResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) | Specifies the results of setup steps. For information on supported values, see the &#x60;stepResults&#x60; parameter above. | [optional] 
**TeardownResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) | Specifies the results of the teardown steps. For information on supported values, see the &#x60;stepResults&#x60; parameter above. | [optional] 

## Methods

### NewAutoTestResultsForTestRunModel

`func NewAutoTestResultsForTestRunModel(configurationId string, autoTestExternalId string, outcome AvailableTestResultOutcome, ) *AutoTestResultsForTestRunModel`

NewAutoTestResultsForTestRunModel instantiates a new AutoTestResultsForTestRunModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestResultsForTestRunModelWithDefaults

`func NewAutoTestResultsForTestRunModelWithDefaults() *AutoTestResultsForTestRunModel`

NewAutoTestResultsForTestRunModelWithDefaults instantiates a new AutoTestResultsForTestRunModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationId

`func (o *AutoTestResultsForTestRunModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *AutoTestResultsForTestRunModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *AutoTestResultsForTestRunModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetLinks

`func (o *AutoTestResultsForTestRunModel) GetLinks() []LinkPostModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTestResultsForTestRunModel) GetLinksOk() (*[]LinkPostModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTestResultsForTestRunModel) SetLinks(v []LinkPostModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTestResultsForTestRunModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTestResultsForTestRunModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTestResultsForTestRunModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetFailureReasonNames

`func (o *AutoTestResultsForTestRunModel) GetFailureReasonNames() []FailureCategoryModel`

GetFailureReasonNames returns the FailureReasonNames field if non-nil, zero value otherwise.

### GetFailureReasonNamesOk

`func (o *AutoTestResultsForTestRunModel) GetFailureReasonNamesOk() (*[]FailureCategoryModel, bool)`

GetFailureReasonNamesOk returns a tuple with the FailureReasonNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasonNames

`func (o *AutoTestResultsForTestRunModel) SetFailureReasonNames(v []FailureCategoryModel)`

SetFailureReasonNames sets FailureReasonNames field to given value.

### HasFailureReasonNames

`func (o *AutoTestResultsForTestRunModel) HasFailureReasonNames() bool`

HasFailureReasonNames returns a boolean if a field has been set.

### SetFailureReasonNamesNil

`func (o *AutoTestResultsForTestRunModel) SetFailureReasonNamesNil(b bool)`

 SetFailureReasonNamesNil sets the value for FailureReasonNames to be an explicit nil

### UnsetFailureReasonNames
`func (o *AutoTestResultsForTestRunModel) UnsetFailureReasonNames()`

UnsetFailureReasonNames ensures that no value is present for FailureReasonNames, not even an explicit nil
### GetAutoTestExternalId

`func (o *AutoTestResultsForTestRunModel) GetAutoTestExternalId() string`

GetAutoTestExternalId returns the AutoTestExternalId field if non-nil, zero value otherwise.

### GetAutoTestExternalIdOk

`func (o *AutoTestResultsForTestRunModel) GetAutoTestExternalIdOk() (*string, bool)`

GetAutoTestExternalIdOk returns a tuple with the AutoTestExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestExternalId

`func (o *AutoTestResultsForTestRunModel) SetAutoTestExternalId(v string)`

SetAutoTestExternalId sets AutoTestExternalId field to given value.


### GetOutcome

`func (o *AutoTestResultsForTestRunModel) GetOutcome() AvailableTestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AutoTestResultsForTestRunModel) GetOutcomeOk() (*AvailableTestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AutoTestResultsForTestRunModel) SetOutcome(v AvailableTestResultOutcome)`

SetOutcome sets Outcome field to given value.


### GetMessage

`func (o *AutoTestResultsForTestRunModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AutoTestResultsForTestRunModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AutoTestResultsForTestRunModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AutoTestResultsForTestRunModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AutoTestResultsForTestRunModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AutoTestResultsForTestRunModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTraces

`func (o *AutoTestResultsForTestRunModel) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *AutoTestResultsForTestRunModel) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *AutoTestResultsForTestRunModel) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *AutoTestResultsForTestRunModel) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *AutoTestResultsForTestRunModel) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *AutoTestResultsForTestRunModel) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetStartedOn

`func (o *AutoTestResultsForTestRunModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *AutoTestResultsForTestRunModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *AutoTestResultsForTestRunModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *AutoTestResultsForTestRunModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *AutoTestResultsForTestRunModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *AutoTestResultsForTestRunModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *AutoTestResultsForTestRunModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *AutoTestResultsForTestRunModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *AutoTestResultsForTestRunModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *AutoTestResultsForTestRunModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *AutoTestResultsForTestRunModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *AutoTestResultsForTestRunModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *AutoTestResultsForTestRunModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutoTestResultsForTestRunModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutoTestResultsForTestRunModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutoTestResultsForTestRunModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AutoTestResultsForTestRunModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AutoTestResultsForTestRunModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetAttachments

`func (o *AutoTestResultsForTestRunModel) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AutoTestResultsForTestRunModel) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AutoTestResultsForTestRunModel) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AutoTestResultsForTestRunModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *AutoTestResultsForTestRunModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *AutoTestResultsForTestRunModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetParameters

`func (o *AutoTestResultsForTestRunModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AutoTestResultsForTestRunModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AutoTestResultsForTestRunModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AutoTestResultsForTestRunModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *AutoTestResultsForTestRunModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *AutoTestResultsForTestRunModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetProperties

`func (o *AutoTestResultsForTestRunModel) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AutoTestResultsForTestRunModel) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AutoTestResultsForTestRunModel) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AutoTestResultsForTestRunModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *AutoTestResultsForTestRunModel) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *AutoTestResultsForTestRunModel) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetStepResults

`func (o *AutoTestResultsForTestRunModel) GetStepResults() []AttachmentPutModelAutoTestStepResultsModel`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *AutoTestResultsForTestRunModel) GetStepResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *AutoTestResultsForTestRunModel) SetStepResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *AutoTestResultsForTestRunModel) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *AutoTestResultsForTestRunModel) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *AutoTestResultsForTestRunModel) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetSetupResults

`func (o *AutoTestResultsForTestRunModel) GetSetupResults() []AttachmentPutModelAutoTestStepResultsModel`

GetSetupResults returns the SetupResults field if non-nil, zero value otherwise.

### GetSetupResultsOk

`func (o *AutoTestResultsForTestRunModel) GetSetupResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetSetupResultsOk returns a tuple with the SetupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupResults

`func (o *AutoTestResultsForTestRunModel) SetSetupResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetSetupResults sets SetupResults field to given value.

### HasSetupResults

`func (o *AutoTestResultsForTestRunModel) HasSetupResults() bool`

HasSetupResults returns a boolean if a field has been set.

### SetSetupResultsNil

`func (o *AutoTestResultsForTestRunModel) SetSetupResultsNil(b bool)`

 SetSetupResultsNil sets the value for SetupResults to be an explicit nil

### UnsetSetupResults
`func (o *AutoTestResultsForTestRunModel) UnsetSetupResults()`

UnsetSetupResults ensures that no value is present for SetupResults, not even an explicit nil
### GetTeardownResults

`func (o *AutoTestResultsForTestRunModel) GetTeardownResults() []AttachmentPutModelAutoTestStepResultsModel`

GetTeardownResults returns the TeardownResults field if non-nil, zero value otherwise.

### GetTeardownResultsOk

`func (o *AutoTestResultsForTestRunModel) GetTeardownResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetTeardownResultsOk returns a tuple with the TeardownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownResults

`func (o *AutoTestResultsForTestRunModel) SetTeardownResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetTeardownResults sets TeardownResults field to given value.

### HasTeardownResults

`func (o *AutoTestResultsForTestRunModel) HasTeardownResults() bool`

HasTeardownResults returns a boolean if a field has been set.

### SetTeardownResultsNil

`func (o *AutoTestResultsForTestRunModel) SetTeardownResultsNil(b bool)`

 SetTeardownResultsNil sets the value for TeardownResults to be an explicit nil

### UnsetTeardownResults
`func (o *AutoTestResultsForTestRunModel) UnsetTeardownResults()`

UnsetTeardownResults ensures that no value is present for TeardownResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


