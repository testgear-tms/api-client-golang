# AutotestHistoricalResultSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcomes** | Pointer to [**[]AutotestResultOutcome**](AutotestResultOutcome.md) |  | [optional] 
**TestPlanIds** | Pointer to **[]string** |  | [optional] 
**TestRunIds** | Pointer to **[]string** |  | [optional] 
**ConfigurationIds** | Pointer to **[]string** |  | [optional] 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 
**Duration** | Pointer to [**Int64RangeSelectorModel**](Int64RangeSelectorModel.md) |  | [optional] 

## Methods

### NewAutotestHistoricalResultSelectModel

`func NewAutotestHistoricalResultSelectModel() *AutotestHistoricalResultSelectModel`

NewAutotestHistoricalResultSelectModel instantiates a new AutotestHistoricalResultSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestHistoricalResultSelectModelWithDefaults

`func NewAutotestHistoricalResultSelectModelWithDefaults() *AutotestHistoricalResultSelectModel`

NewAutotestHistoricalResultSelectModelWithDefaults instantiates a new AutotestHistoricalResultSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcomes

`func (o *AutotestHistoricalResultSelectModel) GetOutcomes() []AutotestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *AutotestHistoricalResultSelectModel) GetOutcomesOk() (*[]AutotestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *AutotestHistoricalResultSelectModel) SetOutcomes(v []AutotestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *AutotestHistoricalResultSelectModel) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *AutotestHistoricalResultSelectModel) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *AutotestHistoricalResultSelectModel) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetTestPlanIds

`func (o *AutotestHistoricalResultSelectModel) GetTestPlanIds() []string`

GetTestPlanIds returns the TestPlanIds field if non-nil, zero value otherwise.

### GetTestPlanIdsOk

`func (o *AutotestHistoricalResultSelectModel) GetTestPlanIdsOk() (*[]string, bool)`

GetTestPlanIdsOk returns a tuple with the TestPlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanIds

`func (o *AutotestHistoricalResultSelectModel) SetTestPlanIds(v []string)`

SetTestPlanIds sets TestPlanIds field to given value.

### HasTestPlanIds

`func (o *AutotestHistoricalResultSelectModel) HasTestPlanIds() bool`

HasTestPlanIds returns a boolean if a field has been set.

### SetTestPlanIdsNil

`func (o *AutotestHistoricalResultSelectModel) SetTestPlanIdsNil(b bool)`

 SetTestPlanIdsNil sets the value for TestPlanIds to be an explicit nil

### UnsetTestPlanIds
`func (o *AutotestHistoricalResultSelectModel) UnsetTestPlanIds()`

UnsetTestPlanIds ensures that no value is present for TestPlanIds, not even an explicit nil
### GetTestRunIds

`func (o *AutotestHistoricalResultSelectModel) GetTestRunIds() []string`

GetTestRunIds returns the TestRunIds field if non-nil, zero value otherwise.

### GetTestRunIdsOk

`func (o *AutotestHistoricalResultSelectModel) GetTestRunIdsOk() (*[]string, bool)`

GetTestRunIdsOk returns a tuple with the TestRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunIds

`func (o *AutotestHistoricalResultSelectModel) SetTestRunIds(v []string)`

SetTestRunIds sets TestRunIds field to given value.

### HasTestRunIds

`func (o *AutotestHistoricalResultSelectModel) HasTestRunIds() bool`

HasTestRunIds returns a boolean if a field has been set.

### SetTestRunIdsNil

`func (o *AutotestHistoricalResultSelectModel) SetTestRunIdsNil(b bool)`

 SetTestRunIdsNil sets the value for TestRunIds to be an explicit nil

### UnsetTestRunIds
`func (o *AutotestHistoricalResultSelectModel) UnsetTestRunIds()`

UnsetTestRunIds ensures that no value is present for TestRunIds, not even an explicit nil
### GetConfigurationIds

`func (o *AutotestHistoricalResultSelectModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *AutotestHistoricalResultSelectModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *AutotestHistoricalResultSelectModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *AutotestHistoricalResultSelectModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *AutotestHistoricalResultSelectModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *AutotestHistoricalResultSelectModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetLaunchSource

`func (o *AutotestHistoricalResultSelectModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *AutotestHistoricalResultSelectModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *AutotestHistoricalResultSelectModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *AutotestHistoricalResultSelectModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *AutotestHistoricalResultSelectModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *AutotestHistoricalResultSelectModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetUserIds

`func (o *AutotestHistoricalResultSelectModel) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *AutotestHistoricalResultSelectModel) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *AutotestHistoricalResultSelectModel) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *AutotestHistoricalResultSelectModel) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *AutotestHistoricalResultSelectModel) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *AutotestHistoricalResultSelectModel) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetDuration

`func (o *AutotestHistoricalResultSelectModel) GetDuration() Int64RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutotestHistoricalResultSelectModel) GetDurationOk() (*Int64RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutotestHistoricalResultSelectModel) SetDuration(v Int64RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutotestHistoricalResultSelectModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


