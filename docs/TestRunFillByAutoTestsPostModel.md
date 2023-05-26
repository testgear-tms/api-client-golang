# TestRunFillByAutoTestsPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Specifies the GUID of the project, in which a test run will be created. | 
**Name** | Pointer to **NullableString** | Specifies the name of the test run. | [optional] 
**ConfigurationIds** | **[]string** | Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs. | 
**AutoTestExternalIds** | **[]string** | Specifies the external ID of the autotest. You can specify several IDs. | 
**Description** | Pointer to **NullableString** | Specifies the test run description. | [optional] 
**LaunchSource** | Pointer to **NullableString** | Specifies the test run launch source. | [optional] 

## Methods

### NewTestRunFillByAutoTestsPostModel

`func NewTestRunFillByAutoTestsPostModel(projectId string, configurationIds []string, autoTestExternalIds []string, ) *TestRunFillByAutoTestsPostModel`

NewTestRunFillByAutoTestsPostModel instantiates a new TestRunFillByAutoTestsPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunFillByAutoTestsPostModelWithDefaults

`func NewTestRunFillByAutoTestsPostModelWithDefaults() *TestRunFillByAutoTestsPostModel`

NewTestRunFillByAutoTestsPostModelWithDefaults instantiates a new TestRunFillByAutoTestsPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *TestRunFillByAutoTestsPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunFillByAutoTestsPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunFillByAutoTestsPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *TestRunFillByAutoTestsPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunFillByAutoTestsPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunFillByAutoTestsPostModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunFillByAutoTestsPostModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunFillByAutoTestsPostModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunFillByAutoTestsPostModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationIds

`func (o *TestRunFillByAutoTestsPostModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestRunFillByAutoTestsPostModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestRunFillByAutoTestsPostModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.


### GetAutoTestExternalIds

`func (o *TestRunFillByAutoTestsPostModel) GetAutoTestExternalIds() []string`

GetAutoTestExternalIds returns the AutoTestExternalIds field if non-nil, zero value otherwise.

### GetAutoTestExternalIdsOk

`func (o *TestRunFillByAutoTestsPostModel) GetAutoTestExternalIdsOk() (*[]string, bool)`

GetAutoTestExternalIdsOk returns a tuple with the AutoTestExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestExternalIds

`func (o *TestRunFillByAutoTestsPostModel) SetAutoTestExternalIds(v []string)`

SetAutoTestExternalIds sets AutoTestExternalIds field to given value.


### GetDescription

`func (o *TestRunFillByAutoTestsPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunFillByAutoTestsPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunFillByAutoTestsPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunFillByAutoTestsPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunFillByAutoTestsPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunFillByAutoTestsPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *TestRunFillByAutoTestsPostModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestRunFillByAutoTestsPostModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestRunFillByAutoTestsPostModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestRunFillByAutoTestsPostModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestRunFillByAutoTestsPostModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestRunFillByAutoTestsPostModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


