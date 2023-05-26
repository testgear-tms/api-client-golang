# AutoTestPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkItemIdsForLinkWithAutoTest** | Pointer to **[]string** | Specifies the IDs of work items to link your autotest to. You can specify several IDs. | [optional] 
**ShouldCreateWorkItem** | Pointer to **NullableBool** | Creates a test case linked to the autotest. | [optional] 
**ExternalId** | **string** | External ID of the autotest | 
**Links** | Pointer to [**[]LinkPostModel**](LinkPostModel.md) | Collection of the autotest links | [optional] 
**ProjectId** | **string** | Unique ID of the autotest project | 
**Name** | **string** | Name of the autotest | 
**Namespace** | Pointer to **NullableString** | Name of the autotest namespace | [optional] 
**Classname** | Pointer to **NullableString** | Name of the autotest class | [optional] 
**Steps** | Pointer to [**[]AutoTestStepModel**](AutoTestStepModel.md) | Collection of the autotest steps | [optional] 
**Setup** | Pointer to [**[]AutoTestStepModel**](AutoTestStepModel.md) | Collection of the autotest setup steps | [optional] 
**Teardown** | Pointer to [**[]AutoTestStepModel**](AutoTestStepModel.md) | Collection of the autotest teardown steps | [optional] 
**Title** | Pointer to **NullableString** | Name of the autotest in autotest&#39;s card | [optional] 
**Description** | Pointer to **NullableString** | Description of the autotest in autotest&#39;s card | [optional] 
**Labels** | Pointer to [**[]LabelPostModel**](LabelPostModel.md) | Collection of the autotest labels | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Indicates if the autotest is marked as flaky | [optional] 

## Methods

### NewAutoTestPostModel

`func NewAutoTestPostModel(externalId string, projectId string, name string, ) *AutoTestPostModel`

NewAutoTestPostModel instantiates a new AutoTestPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestPostModelWithDefaults

`func NewAutoTestPostModelWithDefaults() *AutoTestPostModel`

NewAutoTestPostModelWithDefaults instantiates a new AutoTestPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestPostModel) GetWorkItemIdsForLinkWithAutoTest() []string`

GetWorkItemIdsForLinkWithAutoTest returns the WorkItemIdsForLinkWithAutoTest field if non-nil, zero value otherwise.

### GetWorkItemIdsForLinkWithAutoTestOk

`func (o *AutoTestPostModel) GetWorkItemIdsForLinkWithAutoTestOk() (*[]string, bool)`

GetWorkItemIdsForLinkWithAutoTestOk returns a tuple with the WorkItemIdsForLinkWithAutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestPostModel) SetWorkItemIdsForLinkWithAutoTest(v []string)`

SetWorkItemIdsForLinkWithAutoTest sets WorkItemIdsForLinkWithAutoTest field to given value.

### HasWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestPostModel) HasWorkItemIdsForLinkWithAutoTest() bool`

HasWorkItemIdsForLinkWithAutoTest returns a boolean if a field has been set.

### SetWorkItemIdsForLinkWithAutoTestNil

`func (o *AutoTestPostModel) SetWorkItemIdsForLinkWithAutoTestNil(b bool)`

 SetWorkItemIdsForLinkWithAutoTestNil sets the value for WorkItemIdsForLinkWithAutoTest to be an explicit nil

### UnsetWorkItemIdsForLinkWithAutoTest
`func (o *AutoTestPostModel) UnsetWorkItemIdsForLinkWithAutoTest()`

UnsetWorkItemIdsForLinkWithAutoTest ensures that no value is present for WorkItemIdsForLinkWithAutoTest, not even an explicit nil
### GetShouldCreateWorkItem

`func (o *AutoTestPostModel) GetShouldCreateWorkItem() bool`

GetShouldCreateWorkItem returns the ShouldCreateWorkItem field if non-nil, zero value otherwise.

### GetShouldCreateWorkItemOk

`func (o *AutoTestPostModel) GetShouldCreateWorkItemOk() (*bool, bool)`

GetShouldCreateWorkItemOk returns a tuple with the ShouldCreateWorkItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateWorkItem

`func (o *AutoTestPostModel) SetShouldCreateWorkItem(v bool)`

SetShouldCreateWorkItem sets ShouldCreateWorkItem field to given value.

### HasShouldCreateWorkItem

`func (o *AutoTestPostModel) HasShouldCreateWorkItem() bool`

HasShouldCreateWorkItem returns a boolean if a field has been set.

### SetShouldCreateWorkItemNil

`func (o *AutoTestPostModel) SetShouldCreateWorkItemNil(b bool)`

 SetShouldCreateWorkItemNil sets the value for ShouldCreateWorkItem to be an explicit nil

### UnsetShouldCreateWorkItem
`func (o *AutoTestPostModel) UnsetShouldCreateWorkItem()`

UnsetShouldCreateWorkItem ensures that no value is present for ShouldCreateWorkItem, not even an explicit nil
### GetExternalId

`func (o *AutoTestPostModel) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoTestPostModel) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoTestPostModel) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinks

`func (o *AutoTestPostModel) GetLinks() []LinkPostModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTestPostModel) GetLinksOk() (*[]LinkPostModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTestPostModel) SetLinks(v []LinkPostModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTestPostModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTestPostModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTestPostModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetProjectId

`func (o *AutoTestPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *AutoTestPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestPostModel) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AutoTestPostModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestPostModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestPostModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestPostModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestPostModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestPostModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassname

`func (o *AutoTestPostModel) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *AutoTestPostModel) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *AutoTestPostModel) SetClassname(v string)`

SetClassname sets Classname field to given value.

### HasClassname

`func (o *AutoTestPostModel) HasClassname() bool`

HasClassname returns a boolean if a field has been set.

### SetClassnameNil

`func (o *AutoTestPostModel) SetClassnameNil(b bool)`

 SetClassnameNil sets the value for Classname to be an explicit nil

### UnsetClassname
`func (o *AutoTestPostModel) UnsetClassname()`

UnsetClassname ensures that no value is present for Classname, not even an explicit nil
### GetSteps

`func (o *AutoTestPostModel) GetSteps() []AutoTestStepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestPostModel) GetStepsOk() (*[]AutoTestStepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestPostModel) SetSteps(v []AutoTestStepModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestPostModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestPostModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestPostModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetSetup

`func (o *AutoTestPostModel) GetSetup() []AutoTestStepModel`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AutoTestPostModel) GetSetupOk() (*[]AutoTestStepModel, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AutoTestPostModel) SetSetup(v []AutoTestStepModel)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AutoTestPostModel) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *AutoTestPostModel) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *AutoTestPostModel) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetTeardown

`func (o *AutoTestPostModel) GetTeardown() []AutoTestStepModel`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *AutoTestPostModel) GetTeardownOk() (*[]AutoTestStepModel, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *AutoTestPostModel) SetTeardown(v []AutoTestStepModel)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *AutoTestPostModel) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *AutoTestPostModel) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *AutoTestPostModel) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetTitle

`func (o *AutoTestPostModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestPostModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestPostModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestPostModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestPostModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestPostModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *AutoTestPostModel) GetLabels() []LabelPostModel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AutoTestPostModel) GetLabelsOk() (*[]LabelPostModel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AutoTestPostModel) SetLabels(v []LabelPostModel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AutoTestPostModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AutoTestPostModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AutoTestPostModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestPostModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestPostModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestPostModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTestPostModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTestPostModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTestPostModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


