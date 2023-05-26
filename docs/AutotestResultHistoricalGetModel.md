# AutotestResultHistoricalGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**TestRunId** | Pointer to **string** |  | [optional] 
**TestRunName** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**Outcome** | [**AutotestResultOutcome**](AutotestResultOutcome.md) |  | 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] [readonly] 
**TestPlanId** | Pointer to **NullableString** |  | [optional] 
**TestPlanGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**TestPlanName** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewAutotestResultHistoricalGetModel

`func NewAutotestResultHistoricalGetModel(outcome AutotestResultOutcome, ) *AutotestResultHistoricalGetModel`

NewAutotestResultHistoricalGetModel instantiates a new AutotestResultHistoricalGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestResultHistoricalGetModelWithDefaults

`func NewAutotestResultHistoricalGetModelWithDefaults() *AutotestResultHistoricalGetModel`

NewAutotestResultHistoricalGetModelWithDefaults instantiates a new AutotestResultHistoricalGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutotestResultHistoricalGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutotestResultHistoricalGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutotestResultHistoricalGetModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutotestResultHistoricalGetModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *AutotestResultHistoricalGetModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutotestResultHistoricalGetModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutotestResultHistoricalGetModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AutotestResultHistoricalGetModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedById

`func (o *AutotestResultHistoricalGetModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutotestResultHistoricalGetModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutotestResultHistoricalGetModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *AutotestResultHistoricalGetModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetTestRunId

`func (o *AutotestResultHistoricalGetModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *AutotestResultHistoricalGetModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *AutotestResultHistoricalGetModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *AutotestResultHistoricalGetModel) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### GetTestRunName

`func (o *AutotestResultHistoricalGetModel) GetTestRunName() string`

GetTestRunName returns the TestRunName field if non-nil, zero value otherwise.

### GetTestRunNameOk

`func (o *AutotestResultHistoricalGetModel) GetTestRunNameOk() (*string, bool)`

GetTestRunNameOk returns a tuple with the TestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunName

`func (o *AutotestResultHistoricalGetModel) SetTestRunName(v string)`

SetTestRunName sets TestRunName field to given value.

### HasTestRunName

`func (o *AutotestResultHistoricalGetModel) HasTestRunName() bool`

HasTestRunName returns a boolean if a field has been set.

### SetTestRunNameNil

`func (o *AutotestResultHistoricalGetModel) SetTestRunNameNil(b bool)`

 SetTestRunNameNil sets the value for TestRunName to be an explicit nil

### UnsetTestRunName
`func (o *AutotestResultHistoricalGetModel) UnsetTestRunName()`

UnsetTestRunName ensures that no value is present for TestRunName, not even an explicit nil
### GetConfigurationId

`func (o *AutotestResultHistoricalGetModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *AutotestResultHistoricalGetModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *AutotestResultHistoricalGetModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *AutotestResultHistoricalGetModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetOutcome

`func (o *AutotestResultHistoricalGetModel) GetOutcome() AutotestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AutotestResultHistoricalGetModel) GetOutcomeOk() (*AutotestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AutotestResultHistoricalGetModel) SetOutcome(v AutotestResultOutcome)`

SetOutcome sets Outcome field to given value.


### GetLaunchSource

`func (o *AutotestResultHistoricalGetModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *AutotestResultHistoricalGetModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *AutotestResultHistoricalGetModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *AutotestResultHistoricalGetModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *AutotestResultHistoricalGetModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *AutotestResultHistoricalGetModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetModifiedDate

`func (o *AutotestResultHistoricalGetModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutotestResultHistoricalGetModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutotestResultHistoricalGetModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutotestResultHistoricalGetModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutotestResultHistoricalGetModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutotestResultHistoricalGetModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *AutotestResultHistoricalGetModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutotestResultHistoricalGetModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutotestResultHistoricalGetModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutotestResultHistoricalGetModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutotestResultHistoricalGetModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutotestResultHistoricalGetModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetTestPlanId

`func (o *AutotestResultHistoricalGetModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *AutotestResultHistoricalGetModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *AutotestResultHistoricalGetModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *AutotestResultHistoricalGetModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *AutotestResultHistoricalGetModel) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *AutotestResultHistoricalGetModel) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestPlanGlobalId

`func (o *AutotestResultHistoricalGetModel) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *AutotestResultHistoricalGetModel) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *AutotestResultHistoricalGetModel) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.

### HasTestPlanGlobalId

`func (o *AutotestResultHistoricalGetModel) HasTestPlanGlobalId() bool`

HasTestPlanGlobalId returns a boolean if a field has been set.

### SetTestPlanGlobalIdNil

`func (o *AutotestResultHistoricalGetModel) SetTestPlanGlobalIdNil(b bool)`

 SetTestPlanGlobalIdNil sets the value for TestPlanGlobalId to be an explicit nil

### UnsetTestPlanGlobalId
`func (o *AutotestResultHistoricalGetModel) UnsetTestPlanGlobalId()`

UnsetTestPlanGlobalId ensures that no value is present for TestPlanGlobalId, not even an explicit nil
### GetTestPlanName

`func (o *AutotestResultHistoricalGetModel) GetTestPlanName() string`

GetTestPlanName returns the TestPlanName field if non-nil, zero value otherwise.

### GetTestPlanNameOk

`func (o *AutotestResultHistoricalGetModel) GetTestPlanNameOk() (*string, bool)`

GetTestPlanNameOk returns a tuple with the TestPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanName

`func (o *AutotestResultHistoricalGetModel) SetTestPlanName(v string)`

SetTestPlanName sets TestPlanName field to given value.

### HasTestPlanName

`func (o *AutotestResultHistoricalGetModel) HasTestPlanName() bool`

HasTestPlanName returns a boolean if a field has been set.

### SetTestPlanNameNil

`func (o *AutotestResultHistoricalGetModel) SetTestPlanNameNil(b bool)`

 SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil

### UnsetTestPlanName
`func (o *AutotestResultHistoricalGetModel) UnsetTestPlanName()`

UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
### GetDuration

`func (o *AutotestResultHistoricalGetModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutotestResultHistoricalGetModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutotestResultHistoricalGetModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutotestResultHistoricalGetModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AutotestResultHistoricalGetModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AutotestResultHistoricalGetModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


