# TestPlanChangedFieldsViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**StringTestPlanChangedFieldViewModel**](StringTestPlanChangedFieldViewModel.md) |  | [optional] 
**Description** | Pointer to [**StringTestPlanChangedFieldViewModel**](StringTestPlanChangedFieldViewModel.md) |  | [optional] 
**ProductName** | Pointer to [**StringTestPlanChangedFieldViewModel**](StringTestPlanChangedFieldViewModel.md) |  | [optional] 
**Build** | Pointer to [**StringTestPlanChangedFieldViewModel**](StringTestPlanChangedFieldViewModel.md) |  | [optional] 
**Period** | Pointer to [**PeriodViewModelTestPlanChangedFieldViewModel**](PeriodViewModelTestPlanChangedFieldViewModel.md) |  | [optional] 
**Status** | Pointer to [**StringTestPlanChangedFieldViewModel**](StringTestPlanChangedFieldViewModel.md) |  | [optional] 
**Tags** | Pointer to [**StringArrayTestPlanChangedFieldViewModel**](StringArrayTestPlanChangedFieldViewModel.md) |  | [optional] 
**TestSuite** | Pointer to [**TestSuiteChangeViewModelTestPlanChangedFieldViewModel**](TestSuiteChangeViewModelTestPlanChangedFieldViewModel.md) |  | [optional] 
**TestPoints** | Pointer to [**TestPointChangeViewModelTestPlanChangedFieldViewModel**](TestPointChangeViewModelTestPlanChangedFieldViewModel.md) |  | [optional] 
**TestResults** | Pointer to [**TestResultChangeViewModelTestPlanChangedFieldViewModel**](TestResultChangeViewModelTestPlanChangedFieldViewModel.md) |  | [optional] 
**Locking** | Pointer to [**BooleanTestPlanChangedFieldViewModel**](BooleanTestPlanChangedFieldViewModel.md) |  | [optional] 
**HasAutomaticDurationTimer** | Pointer to [**BooleanNullableTestPlanChangedFieldViewModel**](BooleanNullableTestPlanChangedFieldViewModel.md) |  | [optional] 
**Attributes** | Pointer to [**map[string]CustomAttributeChangeModel**](CustomAttributeChangeModel.md) |  | [optional] 

## Methods

### NewTestPlanChangedFieldsViewModel

`func NewTestPlanChangedFieldsViewModel() *TestPlanChangedFieldsViewModel`

NewTestPlanChangedFieldsViewModel instantiates a new TestPlanChangedFieldsViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanChangedFieldsViewModelWithDefaults

`func NewTestPlanChangedFieldsViewModelWithDefaults() *TestPlanChangedFieldsViewModel`

NewTestPlanChangedFieldsViewModelWithDefaults instantiates a new TestPlanChangedFieldsViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestPlanChangedFieldsViewModel) GetName() StringTestPlanChangedFieldViewModel`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanChangedFieldsViewModel) GetNameOk() (*StringTestPlanChangedFieldViewModel, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanChangedFieldsViewModel) SetName(v StringTestPlanChangedFieldViewModel)`

SetName sets Name field to given value.

### HasName

`func (o *TestPlanChangedFieldsViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TestPlanChangedFieldsViewModel) GetDescription() StringTestPlanChangedFieldViewModel`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestPlanChangedFieldsViewModel) GetDescriptionOk() (*StringTestPlanChangedFieldViewModel, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestPlanChangedFieldsViewModel) SetDescription(v StringTestPlanChangedFieldViewModel)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestPlanChangedFieldsViewModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductName

`func (o *TestPlanChangedFieldsViewModel) GetProductName() StringTestPlanChangedFieldViewModel`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *TestPlanChangedFieldsViewModel) GetProductNameOk() (*StringTestPlanChangedFieldViewModel, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *TestPlanChangedFieldsViewModel) SetProductName(v StringTestPlanChangedFieldViewModel)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *TestPlanChangedFieldsViewModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetBuild

`func (o *TestPlanChangedFieldsViewModel) GetBuild() StringTestPlanChangedFieldViewModel`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanChangedFieldsViewModel) GetBuildOk() (*StringTestPlanChangedFieldViewModel, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanChangedFieldsViewModel) SetBuild(v StringTestPlanChangedFieldViewModel)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanChangedFieldsViewModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetPeriod

`func (o *TestPlanChangedFieldsViewModel) GetPeriod() PeriodViewModelTestPlanChangedFieldViewModel`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TestPlanChangedFieldsViewModel) GetPeriodOk() (*PeriodViewModelTestPlanChangedFieldViewModel, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TestPlanChangedFieldsViewModel) SetPeriod(v PeriodViewModelTestPlanChangedFieldViewModel)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TestPlanChangedFieldsViewModel) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetStatus

`func (o *TestPlanChangedFieldsViewModel) GetStatus() StringTestPlanChangedFieldViewModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPlanChangedFieldsViewModel) GetStatusOk() (*StringTestPlanChangedFieldViewModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPlanChangedFieldsViewModel) SetStatus(v StringTestPlanChangedFieldViewModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPlanChangedFieldsViewModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *TestPlanChangedFieldsViewModel) GetTags() StringArrayTestPlanChangedFieldViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanChangedFieldsViewModel) GetTagsOk() (*StringArrayTestPlanChangedFieldViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanChangedFieldsViewModel) SetTags(v StringArrayTestPlanChangedFieldViewModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanChangedFieldsViewModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTestSuite

`func (o *TestPlanChangedFieldsViewModel) GetTestSuite() TestSuiteChangeViewModelTestPlanChangedFieldViewModel`

GetTestSuite returns the TestSuite field if non-nil, zero value otherwise.

### GetTestSuiteOk

`func (o *TestPlanChangedFieldsViewModel) GetTestSuiteOk() (*TestSuiteChangeViewModelTestPlanChangedFieldViewModel, bool)`

GetTestSuiteOk returns a tuple with the TestSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuite

`func (o *TestPlanChangedFieldsViewModel) SetTestSuite(v TestSuiteChangeViewModelTestPlanChangedFieldViewModel)`

SetTestSuite sets TestSuite field to given value.

### HasTestSuite

`func (o *TestPlanChangedFieldsViewModel) HasTestSuite() bool`

HasTestSuite returns a boolean if a field has been set.

### GetTestPoints

`func (o *TestPlanChangedFieldsViewModel) GetTestPoints() TestPointChangeViewModelTestPlanChangedFieldViewModel`

GetTestPoints returns the TestPoints field if non-nil, zero value otherwise.

### GetTestPointsOk

`func (o *TestPlanChangedFieldsViewModel) GetTestPointsOk() (*TestPointChangeViewModelTestPlanChangedFieldViewModel, bool)`

GetTestPointsOk returns a tuple with the TestPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoints

`func (o *TestPlanChangedFieldsViewModel) SetTestPoints(v TestPointChangeViewModelTestPlanChangedFieldViewModel)`

SetTestPoints sets TestPoints field to given value.

### HasTestPoints

`func (o *TestPlanChangedFieldsViewModel) HasTestPoints() bool`

HasTestPoints returns a boolean if a field has been set.

### GetTestResults

`func (o *TestPlanChangedFieldsViewModel) GetTestResults() TestResultChangeViewModelTestPlanChangedFieldViewModel`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestPlanChangedFieldsViewModel) GetTestResultsOk() (*TestResultChangeViewModelTestPlanChangedFieldViewModel, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestPlanChangedFieldsViewModel) SetTestResults(v TestResultChangeViewModelTestPlanChangedFieldViewModel)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *TestPlanChangedFieldsViewModel) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.

### GetLocking

`func (o *TestPlanChangedFieldsViewModel) GetLocking() BooleanTestPlanChangedFieldViewModel`

GetLocking returns the Locking field if non-nil, zero value otherwise.

### GetLockingOk

`func (o *TestPlanChangedFieldsViewModel) GetLockingOk() (*BooleanTestPlanChangedFieldViewModel, bool)`

GetLockingOk returns a tuple with the Locking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocking

`func (o *TestPlanChangedFieldsViewModel) SetLocking(v BooleanTestPlanChangedFieldViewModel)`

SetLocking sets Locking field to given value.

### HasLocking

`func (o *TestPlanChangedFieldsViewModel) HasLocking() bool`

HasLocking returns a boolean if a field has been set.

### GetHasAutomaticDurationTimer

`func (o *TestPlanChangedFieldsViewModel) GetHasAutomaticDurationTimer() BooleanNullableTestPlanChangedFieldViewModel`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *TestPlanChangedFieldsViewModel) GetHasAutomaticDurationTimerOk() (*BooleanNullableTestPlanChangedFieldViewModel, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *TestPlanChangedFieldsViewModel) SetHasAutomaticDurationTimer(v BooleanNullableTestPlanChangedFieldViewModel)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *TestPlanChangedFieldsViewModel) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### GetAttributes

`func (o *TestPlanChangedFieldsViewModel) GetAttributes() map[string]CustomAttributeChangeModel`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanChangedFieldsViewModel) GetAttributesOk() (*map[string]CustomAttributeChangeModel, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanChangedFieldsViewModel) SetAttributes(v map[string]CustomAttributeChangeModel)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanChangedFieldsViewModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanChangedFieldsViewModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanChangedFieldsViewModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


