# TestRunFillByWorkItemsPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationIds** | **[]string** | Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs. | 
**WorkItemIds** | **[]string** | Specifies the work item GUIDs, from which test points are created. You can specify several GUIDs. | 
**ProjectId** | **string** | Specifies the GUID of the project, in which a test run will be created. | 
**TestPlanId** | **string** | Specifies the GUID of the test plan, within which the test run will be created. | 
**Name** | Pointer to **NullableString** | Specifies the name of the test run. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the test run description. | [optional] 
**LaunchSource** | Pointer to **NullableString** | Specifies the test run launch source. | [optional] 

## Methods

### NewTestRunFillByWorkItemsPostModel

`func NewTestRunFillByWorkItemsPostModel(configurationIds []string, workItemIds []string, projectId string, testPlanId string, ) *TestRunFillByWorkItemsPostModel`

NewTestRunFillByWorkItemsPostModel instantiates a new TestRunFillByWorkItemsPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunFillByWorkItemsPostModelWithDefaults

`func NewTestRunFillByWorkItemsPostModelWithDefaults() *TestRunFillByWorkItemsPostModel`

NewTestRunFillByWorkItemsPostModelWithDefaults instantiates a new TestRunFillByWorkItemsPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationIds

`func (o *TestRunFillByWorkItemsPostModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestRunFillByWorkItemsPostModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestRunFillByWorkItemsPostModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.


### GetWorkItemIds

`func (o *TestRunFillByWorkItemsPostModel) GetWorkItemIds() []string`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *TestRunFillByWorkItemsPostModel) GetWorkItemIdsOk() (*[]string, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *TestRunFillByWorkItemsPostModel) SetWorkItemIds(v []string)`

SetWorkItemIds sets WorkItemIds field to given value.


### GetProjectId

`func (o *TestRunFillByWorkItemsPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunFillByWorkItemsPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunFillByWorkItemsPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTestPlanId

`func (o *TestRunFillByWorkItemsPostModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestRunFillByWorkItemsPostModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestRunFillByWorkItemsPostModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *TestRunFillByWorkItemsPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunFillByWorkItemsPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunFillByWorkItemsPostModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunFillByWorkItemsPostModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunFillByWorkItemsPostModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunFillByWorkItemsPostModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TestRunFillByWorkItemsPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunFillByWorkItemsPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunFillByWorkItemsPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunFillByWorkItemsPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunFillByWorkItemsPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunFillByWorkItemsPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *TestRunFillByWorkItemsPostModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestRunFillByWorkItemsPostModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestRunFillByWorkItemsPostModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestRunFillByWorkItemsPostModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestRunFillByWorkItemsPostModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestRunFillByWorkItemsPostModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


