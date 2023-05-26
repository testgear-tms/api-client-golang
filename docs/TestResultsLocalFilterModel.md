# TestResultsLocalFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationIds** | Pointer to **[]string** | Specifies a test result configuration IDs to search for | [optional] 
**Outcomes** | Pointer to [**[]TestResultOutcome**](TestResultOutcome.md) | Specifies a test result outcomes to search for | [optional] 
**FailureCategories** | Pointer to [**[]FailureCategoryModel**](FailureCategoryModel.md) | Specifies a test result failure categories to search for | [optional] 
**Namespace** | Pointer to **NullableString** | Specifies a test result namespace to search for | [optional] 
**ClassName** | Pointer to **NullableString** | Specifies a test result class name to search for | [optional] 

## Methods

### NewTestResultsLocalFilterModel

`func NewTestResultsLocalFilterModel() *TestResultsLocalFilterModel`

NewTestResultsLocalFilterModel instantiates a new TestResultsLocalFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsLocalFilterModelWithDefaults

`func NewTestResultsLocalFilterModelWithDefaults() *TestResultsLocalFilterModel`

NewTestResultsLocalFilterModelWithDefaults instantiates a new TestResultsLocalFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationIds

`func (o *TestResultsLocalFilterModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestResultsLocalFilterModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestResultsLocalFilterModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *TestResultsLocalFilterModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *TestResultsLocalFilterModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *TestResultsLocalFilterModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetOutcomes

`func (o *TestResultsLocalFilterModel) GetOutcomes() []TestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *TestResultsLocalFilterModel) GetOutcomesOk() (*[]TestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *TestResultsLocalFilterModel) SetOutcomes(v []TestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *TestResultsLocalFilterModel) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *TestResultsLocalFilterModel) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *TestResultsLocalFilterModel) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetFailureCategories

`func (o *TestResultsLocalFilterModel) GetFailureCategories() []FailureCategoryModel`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *TestResultsLocalFilterModel) GetFailureCategoriesOk() (*[]FailureCategoryModel, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *TestResultsLocalFilterModel) SetFailureCategories(v []FailureCategoryModel)`

SetFailureCategories sets FailureCategories field to given value.

### HasFailureCategories

`func (o *TestResultsLocalFilterModel) HasFailureCategories() bool`

HasFailureCategories returns a boolean if a field has been set.

### SetFailureCategoriesNil

`func (o *TestResultsLocalFilterModel) SetFailureCategoriesNil(b bool)`

 SetFailureCategoriesNil sets the value for FailureCategories to be an explicit nil

### UnsetFailureCategories
`func (o *TestResultsLocalFilterModel) UnsetFailureCategories()`

UnsetFailureCategories ensures that no value is present for FailureCategories, not even an explicit nil
### GetNamespace

`func (o *TestResultsLocalFilterModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TestResultsLocalFilterModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TestResultsLocalFilterModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TestResultsLocalFilterModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *TestResultsLocalFilterModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *TestResultsLocalFilterModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassName

`func (o *TestResultsLocalFilterModel) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *TestResultsLocalFilterModel) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *TestResultsLocalFilterModel) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *TestResultsLocalFilterModel) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *TestResultsLocalFilterModel) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *TestResultsLocalFilterModel) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


