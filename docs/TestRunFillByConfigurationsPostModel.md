# TestRunFillByConfigurationsPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestPointSelectors** | [**[]TestPointSelector**](TestPointSelector.md) | Specifies an array of work items and configuration to create a test run for. | 
**ProjectId** | **string** | Specifies the GUID of the project, in which a test run will be created. | 
**TestPlanId** | **string** | Specifies the GUID of the test plan, within which the test run will be created. | 
**Name** | Pointer to **NullableString** | Specifies the name of the test run. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the test run description. | [optional] 
**LaunchSource** | Pointer to **NullableString** | Specifies the test run launch source. | [optional] 

## Methods

### NewTestRunFillByConfigurationsPostModel

`func NewTestRunFillByConfigurationsPostModel(testPointSelectors []TestPointSelector, projectId string, testPlanId string, ) *TestRunFillByConfigurationsPostModel`

NewTestRunFillByConfigurationsPostModel instantiates a new TestRunFillByConfigurationsPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunFillByConfigurationsPostModelWithDefaults

`func NewTestRunFillByConfigurationsPostModelWithDefaults() *TestRunFillByConfigurationsPostModel`

NewTestRunFillByConfigurationsPostModelWithDefaults instantiates a new TestRunFillByConfigurationsPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestPointSelectors

`func (o *TestRunFillByConfigurationsPostModel) GetTestPointSelectors() []TestPointSelector`

GetTestPointSelectors returns the TestPointSelectors field if non-nil, zero value otherwise.

### GetTestPointSelectorsOk

`func (o *TestRunFillByConfigurationsPostModel) GetTestPointSelectorsOk() (*[]TestPointSelector, bool)`

GetTestPointSelectorsOk returns a tuple with the TestPointSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointSelectors

`func (o *TestRunFillByConfigurationsPostModel) SetTestPointSelectors(v []TestPointSelector)`

SetTestPointSelectors sets TestPointSelectors field to given value.


### GetProjectId

`func (o *TestRunFillByConfigurationsPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunFillByConfigurationsPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunFillByConfigurationsPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTestPlanId

`func (o *TestRunFillByConfigurationsPostModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestRunFillByConfigurationsPostModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestRunFillByConfigurationsPostModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *TestRunFillByConfigurationsPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunFillByConfigurationsPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunFillByConfigurationsPostModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunFillByConfigurationsPostModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunFillByConfigurationsPostModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunFillByConfigurationsPostModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TestRunFillByConfigurationsPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunFillByConfigurationsPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunFillByConfigurationsPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunFillByConfigurationsPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunFillByConfigurationsPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunFillByConfigurationsPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *TestRunFillByConfigurationsPostModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestRunFillByConfigurationsPostModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestRunFillByConfigurationsPostModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestRunFillByConfigurationsPostModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestRunFillByConfigurationsPostModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestRunFillByConfigurationsPostModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


