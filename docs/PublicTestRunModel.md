# PublicTestRunModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestRunId** | Pointer to **string** |  | [optional] 
**TestPlanId** | Pointer to **NullableString** |  | [optional] 
**TestPlanGlobalId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**Configurations** | Pointer to [**[]ConfigurationModel**](ConfigurationModel.md) |  | [optional] 
**AutoTests** | Pointer to [**[]AutoTestModel**](AutoTestModel.md) |  | [optional] 
**TestPoints** | Pointer to [**[]PublicTestPointModel**](PublicTestPointModel.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPublicTestRunModel

`func NewPublicTestRunModel() *PublicTestRunModel`

NewPublicTestRunModel instantiates a new PublicTestRunModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTestRunModelWithDefaults

`func NewPublicTestRunModelWithDefaults() *PublicTestRunModel`

NewPublicTestRunModelWithDefaults instantiates a new PublicTestRunModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestRunId

`func (o *PublicTestRunModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *PublicTestRunModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *PublicTestRunModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *PublicTestRunModel) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### GetTestPlanId

`func (o *PublicTestRunModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *PublicTestRunModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *PublicTestRunModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *PublicTestRunModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *PublicTestRunModel) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *PublicTestRunModel) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestPlanGlobalId

`func (o *PublicTestRunModel) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *PublicTestRunModel) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *PublicTestRunModel) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.

### HasTestPlanGlobalId

`func (o *PublicTestRunModel) HasTestPlanGlobalId() bool`

HasTestPlanGlobalId returns a boolean if a field has been set.

### GetName

`func (o *PublicTestRunModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicTestRunModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicTestRunModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicTestRunModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PublicTestRunModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PublicTestRunModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProductName

`func (o *PublicTestRunModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PublicTestRunModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PublicTestRunModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *PublicTestRunModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *PublicTestRunModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *PublicTestRunModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetBuild

`func (o *PublicTestRunModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *PublicTestRunModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *PublicTestRunModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *PublicTestRunModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *PublicTestRunModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *PublicTestRunModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetConfigurations

`func (o *PublicTestRunModel) GetConfigurations() []ConfigurationModel`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *PublicTestRunModel) GetConfigurationsOk() (*[]ConfigurationModel, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *PublicTestRunModel) SetConfigurations(v []ConfigurationModel)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *PublicTestRunModel) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### SetConfigurationsNil

`func (o *PublicTestRunModel) SetConfigurationsNil(b bool)`

 SetConfigurationsNil sets the value for Configurations to be an explicit nil

### UnsetConfigurations
`func (o *PublicTestRunModel) UnsetConfigurations()`

UnsetConfigurations ensures that no value is present for Configurations, not even an explicit nil
### GetAutoTests

`func (o *PublicTestRunModel) GetAutoTests() []AutoTestModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *PublicTestRunModel) GetAutoTestsOk() (*[]AutoTestModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *PublicTestRunModel) SetAutoTests(v []AutoTestModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *PublicTestRunModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *PublicTestRunModel) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *PublicTestRunModel) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil
### GetTestPoints

`func (o *PublicTestRunModel) GetTestPoints() []PublicTestPointModel`

GetTestPoints returns the TestPoints field if non-nil, zero value otherwise.

### GetTestPointsOk

`func (o *PublicTestRunModel) GetTestPointsOk() (*[]PublicTestPointModel, bool)`

GetTestPointsOk returns a tuple with the TestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoints

`func (o *PublicTestRunModel) SetTestPoints(v []PublicTestPointModel)`

SetTestPoints sets TestPoints field to given value.

### HasTestPoints

`func (o *PublicTestRunModel) HasTestPoints() bool`

HasTestPoints returns a boolean if a field has been set.

### SetTestPointsNil

`func (o *PublicTestRunModel) SetTestPointsNil(b bool)`

 SetTestPointsNil sets the value for TestPoints to be an explicit nil

### UnsetTestPoints
`func (o *PublicTestRunModel) UnsetTestPoints()`

UnsetTestPoints ensures that no value is present for TestPoints, not even an explicit nil
### GetStatus

`func (o *PublicTestRunModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicTestRunModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicTestRunModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicTestRunModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PublicTestRunModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PublicTestRunModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


