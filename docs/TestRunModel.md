# TestRunModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTests** | Pointer to [**[]AutoTestModel**](AutoTestModel.md) |  | [optional] 
**AutoTestsCount** | Pointer to **int32** |  | [optional] 
**TestSuiteIds** | Pointer to **[]string** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**Analytic** | Pointer to [**TestRunAnalyticResultModel**](TestRunAnalyticResultModel.md) |  | [optional] 
**TestResults** | Pointer to [**[]TestResultModel**](TestResultModel.md) |  | [optional] 
**TestPlan** | Pointer to [**TestPlanModel**](TestPlanModel.md) |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**CreatedByUserName** | Pointer to **NullableString** |  | [optional] 
**StartedDate** | Pointer to **NullableTime** |  | [optional] 
**CompletedDate** | Pointer to **NullableTime** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StateName** | [**TestRunState**](TestRunState.md) |  | 
**ProjectId** | Pointer to **string** |  | [optional] 
**TestPlanId** | Pointer to **NullableString** |  | [optional] 
**RunByUserId** | Pointer to **NullableString** |  | [optional] 
**StoppedByUserId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewTestRunModel

`func NewTestRunModel(stateName TestRunState, ) *TestRunModel`

NewTestRunModel instantiates a new TestRunModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunModelWithDefaults

`func NewTestRunModelWithDefaults() *TestRunModel`

NewTestRunModelWithDefaults instantiates a new TestRunModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTests

`func (o *TestRunModel) GetAutoTests() []AutoTestModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *TestRunModel) GetAutoTestsOk() (*[]AutoTestModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *TestRunModel) SetAutoTests(v []AutoTestModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *TestRunModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *TestRunModel) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *TestRunModel) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil
### GetAutoTestsCount

`func (o *TestRunModel) GetAutoTestsCount() int32`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *TestRunModel) GetAutoTestsCountOk() (*int32, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *TestRunModel) SetAutoTestsCount(v int32)`

SetAutoTestsCount sets AutoTestsCount field to given value.

### HasAutoTestsCount

`func (o *TestRunModel) HasAutoTestsCount() bool`

HasAutoTestsCount returns a boolean if a field has been set.

### GetTestSuiteIds

`func (o *TestRunModel) GetTestSuiteIds() []string`

GetTestSuiteIds returns the TestSuiteIds field if non-nil, zero value otherwise.

### GetTestSuiteIdsOk

`func (o *TestRunModel) GetTestSuiteIdsOk() (*[]string, bool)`

GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteIds

`func (o *TestRunModel) SetTestSuiteIds(v []string)`

SetTestSuiteIds sets TestSuiteIds field to given value.

### HasTestSuiteIds

`func (o *TestRunModel) HasTestSuiteIds() bool`

HasTestSuiteIds returns a boolean if a field has been set.

### SetTestSuiteIdsNil

`func (o *TestRunModel) SetTestSuiteIdsNil(b bool)`

 SetTestSuiteIdsNil sets the value for TestSuiteIds to be an explicit nil

### UnsetTestSuiteIds
`func (o *TestRunModel) UnsetTestSuiteIds()`

UnsetTestSuiteIds ensures that no value is present for TestSuiteIds, not even an explicit nil
### GetIsAutomated

`func (o *TestRunModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestRunModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestRunModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestRunModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetAnalytic

`func (o *TestRunModel) GetAnalytic() TestRunAnalyticResultModel`

GetAnalytic returns the Analytic field if non-nil, zero value otherwise.

### GetAnalyticOk

`func (o *TestRunModel) GetAnalyticOk() (*TestRunAnalyticResultModel, bool)`

GetAnalyticOk returns a tuple with the Analytic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytic

`func (o *TestRunModel) SetAnalytic(v TestRunAnalyticResultModel)`

SetAnalytic sets Analytic field to given value.

### HasAnalytic

`func (o *TestRunModel) HasAnalytic() bool`

HasAnalytic returns a boolean if a field has been set.

### GetTestResults

`func (o *TestRunModel) GetTestResults() []TestResultModel`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestRunModel) GetTestResultsOk() (*[]TestResultModel, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestRunModel) SetTestResults(v []TestResultModel)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *TestRunModel) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.

### SetTestResultsNil

`func (o *TestRunModel) SetTestResultsNil(b bool)`

 SetTestResultsNil sets the value for TestResults to be an explicit nil

### UnsetTestResults
`func (o *TestRunModel) UnsetTestResults()`

UnsetTestResults ensures that no value is present for TestResults, not even an explicit nil
### GetTestPlan

`func (o *TestRunModel) GetTestPlan() TestPlanModel`

GetTestPlan returns the TestPlan field if non-nil, zero value otherwise.

### GetTestPlanOk

`func (o *TestRunModel) GetTestPlanOk() (*TestPlanModel, bool)`

GetTestPlanOk returns a tuple with the TestPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlan

`func (o *TestRunModel) SetTestPlan(v TestPlanModel)`

SetTestPlan sets TestPlan field to given value.

### HasTestPlan

`func (o *TestRunModel) HasTestPlan() bool`

HasTestPlan returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TestRunModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestRunModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestRunModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestRunModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *TestRunModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestRunModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestRunModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestRunModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestRunModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestRunModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestRunModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestRunModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestRunModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestRunModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *TestRunModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestRunModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestRunModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestRunModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestRunModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestRunModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedByUserName

`func (o *TestRunModel) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *TestRunModel) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *TestRunModel) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *TestRunModel) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### SetCreatedByUserNameNil

`func (o *TestRunModel) SetCreatedByUserNameNil(b bool)`

 SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil

### UnsetCreatedByUserName
`func (o *TestRunModel) UnsetCreatedByUserName()`

UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
### GetStartedDate

`func (o *TestRunModel) GetStartedDate() time.Time`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *TestRunModel) GetStartedDateOk() (*time.Time, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *TestRunModel) SetStartedDate(v time.Time)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *TestRunModel) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *TestRunModel) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *TestRunModel) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCompletedDate

`func (o *TestRunModel) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *TestRunModel) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *TestRunModel) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *TestRunModel) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *TestRunModel) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *TestRunModel) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
### GetBuild

`func (o *TestRunModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestRunModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestRunModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestRunModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestRunModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestRunModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetDescription

`func (o *TestRunModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStateName

`func (o *TestRunModel) GetStateName() TestRunState`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *TestRunModel) GetStateNameOk() (*TestRunState, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *TestRunModel) SetStateName(v TestRunState)`

SetStateName sets StateName field to given value.


### GetProjectId

`func (o *TestRunModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TestRunModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTestPlanId

`func (o *TestRunModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestRunModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestRunModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestRunModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *TestRunModel) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestRunModel) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetRunByUserId

`func (o *TestRunModel) GetRunByUserId() string`

GetRunByUserId returns the RunByUserId field if non-nil, zero value otherwise.

### GetRunByUserIdOk

`func (o *TestRunModel) GetRunByUserIdOk() (*string, bool)`

GetRunByUserIdOk returns a tuple with the RunByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByUserId

`func (o *TestRunModel) SetRunByUserId(v string)`

SetRunByUserId sets RunByUserId field to given value.

### HasRunByUserId

`func (o *TestRunModel) HasRunByUserId() bool`

HasRunByUserId returns a boolean if a field has been set.

### SetRunByUserIdNil

`func (o *TestRunModel) SetRunByUserIdNil(b bool)`

 SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil

### UnsetRunByUserId
`func (o *TestRunModel) UnsetRunByUserId()`

UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
### GetStoppedByUserId

`func (o *TestRunModel) GetStoppedByUserId() string`

GetStoppedByUserId returns the StoppedByUserId field if non-nil, zero value otherwise.

### GetStoppedByUserIdOk

`func (o *TestRunModel) GetStoppedByUserIdOk() (*string, bool)`

GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedByUserId

`func (o *TestRunModel) SetStoppedByUserId(v string)`

SetStoppedByUserId sets StoppedByUserId field to given value.

### HasStoppedByUserId

`func (o *TestRunModel) HasStoppedByUserId() bool`

HasStoppedByUserId returns a boolean if a field has been set.

### SetStoppedByUserIdNil

`func (o *TestRunModel) SetStoppedByUserIdNil(b bool)`

 SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil

### UnsetStoppedByUserId
`func (o *TestRunModel) UnsetStoppedByUserId()`

UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
### GetName

`func (o *TestRunModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLaunchSource

`func (o *TestRunModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestRunModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestRunModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestRunModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestRunModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestRunModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetId

`func (o *TestRunModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestRunModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TestRunModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestRunModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


