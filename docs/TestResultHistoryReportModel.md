# TestResultHistoryReportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** | If test run was stopped, this property equals identifier of user who stopped it.Otherwise, the property equals identifier of user who created the test result | [optional] 
**TestRunId** | Pointer to **NullableString** |  | [optional] 
**TestRunName** | Pointer to **NullableString** |  | [optional] 
**CreatedByUserName** | Pointer to **NullableString** |  | [optional] 
**TestPlanId** | Pointer to **NullableString** |  | [optional] 
**TestPlanGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**TestPlanName** | Pointer to **NullableString** |  | [optional] 
**ConfigurationName** | Pointer to **NullableString** | If test point related to the test result has configuration, this property will be equal to the test point configuration name. Otherwise, this property will be equal to the test result configuration name | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**Outcome** | Pointer to **NullableString** | If any test result related to the test run is linked with autotest and the run has an outcome, the outcome value equalsto the worst outcome of the last modified test result.Otherwise, the outcome equals to the outcome of first created test result in the test run | [optional] 
**Comment** | Pointer to **NullableString** | If any test result related to the test run is linked with autotest, comment will have default valueOtherwise, the comment equals to the comment of first created test result in the test run | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) | If any test result related to the test run is linked with autotest, link will be equal to the links of last modified test result.Otherwise, the links equals to the links of first created test result in the test run | [optional] 
**StartedOn** | Pointer to **NullableTime** |  | [optional] 
**CompletedOn** | Pointer to **NullableTime** |  | [optional] 
**Duration** | Pointer to **NullableInt64** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) | If any test result related to the test run is linked with autotest, attachments will be equal to the attachments of last modified test result.Otherwise, the attachments equals to the attachments of first created test result in the test run | [optional] 
**WorkItemVersionId** | Pointer to **NullableString** |  | [optional] 
**WorkItemVersionNumber** | Pointer to **NullableInt32** |  | [optional] 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**FailureClassIds** | Pointer to **[]string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTestResultHistoryReportModel

`func NewTestResultHistoryReportModel() *TestResultHistoryReportModel`

NewTestResultHistoryReportModel instantiates a new TestResultHistoryReportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultHistoryReportModelWithDefaults

`func NewTestResultHistoryReportModelWithDefaults() *TestResultHistoryReportModel`

NewTestResultHistoryReportModelWithDefaults instantiates a new TestResultHistoryReportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultHistoryReportModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultHistoryReportModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultHistoryReportModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestResultHistoryReportModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TestResultHistoryReportModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultHistoryReportModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultHistoryReportModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestResultHistoryReportModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *TestResultHistoryReportModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultHistoryReportModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultHistoryReportModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultHistoryReportModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetUserId

`func (o *TestResultHistoryReportModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TestResultHistoryReportModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TestResultHistoryReportModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TestResultHistoryReportModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTestRunId

`func (o *TestResultHistoryReportModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultHistoryReportModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultHistoryReportModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *TestResultHistoryReportModel) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### SetTestRunIdNil

`func (o *TestResultHistoryReportModel) SetTestRunIdNil(b bool)`

 SetTestRunIdNil sets the value for TestRunId to be an explicit nil

### UnsetTestRunId
`func (o *TestResultHistoryReportModel) UnsetTestRunId()`

UnsetTestRunId ensures that no value is present for TestRunId, not even an explicit nil
### GetTestRunName

`func (o *TestResultHistoryReportModel) GetTestRunName() string`

GetTestRunName returns the TestRunName field if non-nil, zero value otherwise.

### GetTestRunNameOk

`func (o *TestResultHistoryReportModel) GetTestRunNameOk() (*string, bool)`

GetTestRunNameOk returns a tuple with the TestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunName

`func (o *TestResultHistoryReportModel) SetTestRunName(v string)`

SetTestRunName sets TestRunName field to given value.

### HasTestRunName

`func (o *TestResultHistoryReportModel) HasTestRunName() bool`

HasTestRunName returns a boolean if a field has been set.

### SetTestRunNameNil

`func (o *TestResultHistoryReportModel) SetTestRunNameNil(b bool)`

 SetTestRunNameNil sets the value for TestRunName to be an explicit nil

### UnsetTestRunName
`func (o *TestResultHistoryReportModel) UnsetTestRunName()`

UnsetTestRunName ensures that no value is present for TestRunName, not even an explicit nil
### GetCreatedByUserName

`func (o *TestResultHistoryReportModel) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *TestResultHistoryReportModel) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *TestResultHistoryReportModel) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *TestResultHistoryReportModel) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### SetCreatedByUserNameNil

`func (o *TestResultHistoryReportModel) SetCreatedByUserNameNil(b bool)`

 SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil

### UnsetCreatedByUserName
`func (o *TestResultHistoryReportModel) UnsetCreatedByUserName()`

UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
### GetTestPlanId

`func (o *TestResultHistoryReportModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestResultHistoryReportModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestResultHistoryReportModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestResultHistoryReportModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *TestResultHistoryReportModel) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestResultHistoryReportModel) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestPlanGlobalId

`func (o *TestResultHistoryReportModel) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *TestResultHistoryReportModel) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *TestResultHistoryReportModel) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.

### HasTestPlanGlobalId

`func (o *TestResultHistoryReportModel) HasTestPlanGlobalId() bool`

HasTestPlanGlobalId returns a boolean if a field has been set.

### SetTestPlanGlobalIdNil

`func (o *TestResultHistoryReportModel) SetTestPlanGlobalIdNil(b bool)`

 SetTestPlanGlobalIdNil sets the value for TestPlanGlobalId to be an explicit nil

### UnsetTestPlanGlobalId
`func (o *TestResultHistoryReportModel) UnsetTestPlanGlobalId()`

UnsetTestPlanGlobalId ensures that no value is present for TestPlanGlobalId, not even an explicit nil
### GetTestPlanName

`func (o *TestResultHistoryReportModel) GetTestPlanName() string`

GetTestPlanName returns the TestPlanName field if non-nil, zero value otherwise.

### GetTestPlanNameOk

`func (o *TestResultHistoryReportModel) GetTestPlanNameOk() (*string, bool)`

GetTestPlanNameOk returns a tuple with the TestPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanName

`func (o *TestResultHistoryReportModel) SetTestPlanName(v string)`

SetTestPlanName sets TestPlanName field to given value.

### HasTestPlanName

`func (o *TestResultHistoryReportModel) HasTestPlanName() bool`

HasTestPlanName returns a boolean if a field has been set.

### SetTestPlanNameNil

`func (o *TestResultHistoryReportModel) SetTestPlanNameNil(b bool)`

 SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil

### UnsetTestPlanName
`func (o *TestResultHistoryReportModel) UnsetTestPlanName()`

UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
### GetConfigurationName

`func (o *TestResultHistoryReportModel) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestResultHistoryReportModel) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestResultHistoryReportModel) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *TestResultHistoryReportModel) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### SetConfigurationNameNil

`func (o *TestResultHistoryReportModel) SetConfigurationNameNil(b bool)`

 SetConfigurationNameNil sets the value for ConfigurationName to be an explicit nil

### UnsetConfigurationName
`func (o *TestResultHistoryReportModel) UnsetConfigurationName()`

UnsetConfigurationName ensures that no value is present for ConfigurationName, not even an explicit nil
### GetIsAutomated

`func (o *TestResultHistoryReportModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestResultHistoryReportModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestResultHistoryReportModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestResultHistoryReportModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetOutcome

`func (o *TestResultHistoryReportModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultHistoryReportModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultHistoryReportModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultHistoryReportModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultHistoryReportModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultHistoryReportModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetComment

`func (o *TestResultHistoryReportModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultHistoryReportModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultHistoryReportModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultHistoryReportModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultHistoryReportModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultHistoryReportModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultHistoryReportModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultHistoryReportModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultHistoryReportModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultHistoryReportModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultHistoryReportModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultHistoryReportModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStartedOn

`func (o *TestResultHistoryReportModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultHistoryReportModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultHistoryReportModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultHistoryReportModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultHistoryReportModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultHistoryReportModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultHistoryReportModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultHistoryReportModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultHistoryReportModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultHistoryReportModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultHistoryReportModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultHistoryReportModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *TestResultHistoryReportModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestResultHistoryReportModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestResultHistoryReportModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestResultHistoryReportModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestResultHistoryReportModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestResultHistoryReportModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetCreatedById

`func (o *TestResultHistoryReportModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestResultHistoryReportModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestResultHistoryReportModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestResultHistoryReportModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *TestResultHistoryReportModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestResultHistoryReportModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestResultHistoryReportModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestResultHistoryReportModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestResultHistoryReportModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestResultHistoryReportModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetAttachments

`func (o *TestResultHistoryReportModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultHistoryReportModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultHistoryReportModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultHistoryReportModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultHistoryReportModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultHistoryReportModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetWorkItemVersionId

`func (o *TestResultHistoryReportModel) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultHistoryReportModel) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultHistoryReportModel) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestResultHistoryReportModel) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### SetWorkItemVersionIdNil

`func (o *TestResultHistoryReportModel) SetWorkItemVersionIdNil(b bool)`

 SetWorkItemVersionIdNil sets the value for WorkItemVersionId to be an explicit nil

### UnsetWorkItemVersionId
`func (o *TestResultHistoryReportModel) UnsetWorkItemVersionId()`

UnsetWorkItemVersionId ensures that no value is present for WorkItemVersionId, not even an explicit nil
### GetWorkItemVersionNumber

`func (o *TestResultHistoryReportModel) GetWorkItemVersionNumber() int32`

GetWorkItemVersionNumber returns the WorkItemVersionNumber field if non-nil, zero value otherwise.

### GetWorkItemVersionNumberOk

`func (o *TestResultHistoryReportModel) GetWorkItemVersionNumberOk() (*int32, bool)`

GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionNumber

`func (o *TestResultHistoryReportModel) SetWorkItemVersionNumber(v int32)`

SetWorkItemVersionNumber sets WorkItemVersionNumber field to given value.

### HasWorkItemVersionNumber

`func (o *TestResultHistoryReportModel) HasWorkItemVersionNumber() bool`

HasWorkItemVersionNumber returns a boolean if a field has been set.

### SetWorkItemVersionNumberNil

`func (o *TestResultHistoryReportModel) SetWorkItemVersionNumberNil(b bool)`

 SetWorkItemVersionNumberNil sets the value for WorkItemVersionNumber to be an explicit nil

### UnsetWorkItemVersionNumber
`func (o *TestResultHistoryReportModel) UnsetWorkItemVersionNumber()`

UnsetWorkItemVersionNumber ensures that no value is present for WorkItemVersionNumber, not even an explicit nil
### GetLaunchSource

`func (o *TestResultHistoryReportModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestResultHistoryReportModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestResultHistoryReportModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestResultHistoryReportModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestResultHistoryReportModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestResultHistoryReportModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetFailureClassIds

`func (o *TestResultHistoryReportModel) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *TestResultHistoryReportModel) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *TestResultHistoryReportModel) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.

### HasFailureClassIds

`func (o *TestResultHistoryReportModel) HasFailureClassIds() bool`

HasFailureClassIds returns a boolean if a field has been set.

### SetFailureClassIdsNil

`func (o *TestResultHistoryReportModel) SetFailureClassIdsNil(b bool)`

 SetFailureClassIdsNil sets the value for FailureClassIds to be an explicit nil

### UnsetFailureClassIds
`func (o *TestResultHistoryReportModel) UnsetFailureClassIds()`

UnsetFailureClassIds ensures that no value is present for FailureClassIds, not even an explicit nil
### GetParameters

`func (o *TestResultHistoryReportModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultHistoryReportModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultHistoryReportModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultHistoryReportModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultHistoryReportModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultHistoryReportModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


