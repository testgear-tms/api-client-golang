# TestResultsFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestRunIds** | Pointer to **[]string** | Specifies a test result test run IDs to search for | [optional] 
**ConfigurationIds** | Pointer to **[]string** | Specifies a test result configuration IDs to search for | [optional] 
**Outcomes** | Pointer to [**[]TestResultOutcome**](TestResultOutcome.md) | Specifies a test result outcomes to search for | [optional] 
**FailureCategories** | Pointer to [**[]FailureCategoryModel**](FailureCategoryModel.md) | Specifies a test result failure categories to search for | [optional] 
**Namespace** | Pointer to **NullableString** | Specifies a test result namespace to search for | [optional] 
**ClassName** | Pointer to **NullableString** | Specifies a test result class name to search for | [optional] 

## Methods

### NewTestResultsFilterModel

`func NewTestResultsFilterModel() *TestResultsFilterModel`

NewTestResultsFilterModel instantiates a new TestResultsFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsFilterModelWithDefaults

`func NewTestResultsFilterModelWithDefaults() *TestResultsFilterModel`

NewTestResultsFilterModelWithDefaults instantiates a new TestResultsFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestRunIds

`func (o *TestResultsFilterModel) GetTestRunIds() []string`

GetTestRunIds returns the TestRunIds field if non-nil, zero value otherwise.

### GetTestRunIdsOk

`func (o *TestResultsFilterModel) GetTestRunIdsOk() (*[]string, bool)`

GetTestRunIdsOk returns a tuple with the TestRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunIds

`func (o *TestResultsFilterModel) SetTestRunIds(v []string)`

SetTestRunIds sets TestRunIds field to given value.

### HasTestRunIds

`func (o *TestResultsFilterModel) HasTestRunIds() bool`

HasTestRunIds returns a boolean if a field has been set.

### SetTestRunIdsNil

`func (o *TestResultsFilterModel) SetTestRunIdsNil(b bool)`

 SetTestRunIdsNil sets the value for TestRunIds to be an explicit nil

### UnsetTestRunIds
`func (o *TestResultsFilterModel) UnsetTestRunIds()`

UnsetTestRunIds ensures that no value is present for TestRunIds, not even an explicit nil
### GetConfigurationIds

`func (o *TestResultsFilterModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestResultsFilterModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestResultsFilterModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *TestResultsFilterModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *TestResultsFilterModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *TestResultsFilterModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetOutcomes

`func (o *TestResultsFilterModel) GetOutcomes() []TestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *TestResultsFilterModel) GetOutcomesOk() (*[]TestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *TestResultsFilterModel) SetOutcomes(v []TestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *TestResultsFilterModel) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *TestResultsFilterModel) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *TestResultsFilterModel) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetFailureCategories

`func (o *TestResultsFilterModel) GetFailureCategories() []FailureCategoryModel`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *TestResultsFilterModel) GetFailureCategoriesOk() (*[]FailureCategoryModel, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *TestResultsFilterModel) SetFailureCategories(v []FailureCategoryModel)`

SetFailureCategories sets FailureCategories field to given value.

### HasFailureCategories

`func (o *TestResultsFilterModel) HasFailureCategories() bool`

HasFailureCategories returns a boolean if a field has been set.

### SetFailureCategoriesNil

`func (o *TestResultsFilterModel) SetFailureCategoriesNil(b bool)`

 SetFailureCategoriesNil sets the value for FailureCategories to be an explicit nil

### UnsetFailureCategories
`func (o *TestResultsFilterModel) UnsetFailureCategories()`

UnsetFailureCategories ensures that no value is present for FailureCategories, not even an explicit nil
### GetNamespace

`func (o *TestResultsFilterModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TestResultsFilterModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TestResultsFilterModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TestResultsFilterModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *TestResultsFilterModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *TestResultsFilterModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassName

`func (o *TestResultsFilterModel) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *TestResultsFilterModel) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *TestResultsFilterModel) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *TestResultsFilterModel) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *TestResultsFilterModel) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *TestResultsFilterModel) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


