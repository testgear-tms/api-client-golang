# AutoTestPutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Used for search autotest. If value equals Guid mask filled with zeros, search will be executed using ExternalId | [optional] 
**WorkItemIdsForLinkWithAutoTest** | Pointer to **[]string** |  | [optional] 
**ExternalId** | **string** | External ID of the autotest | 
**Links** | Pointer to [**[]LinkPutModel**](LinkPutModel.md) | Collection of the autotest links | [optional] 
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

### NewAutoTestPutModel

`func NewAutoTestPutModel(externalId string, projectId string, name string, ) *AutoTestPutModel`

NewAutoTestPutModel instantiates a new AutoTestPutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestPutModelWithDefaults

`func NewAutoTestPutModelWithDefaults() *AutoTestPutModel`

NewAutoTestPutModelWithDefaults instantiates a new AutoTestPutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTestPutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTestPutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTestPutModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoTestPutModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestPutModel) GetWorkItemIdsForLinkWithAutoTest() []string`

GetWorkItemIdsForLinkWithAutoTest returns the WorkItemIdsForLinkWithAutoTest field if non-nil, zero value otherwise.

### GetWorkItemIdsForLinkWithAutoTestOk

`func (o *AutoTestPutModel) GetWorkItemIdsForLinkWithAutoTestOk() (*[]string, bool)`

GetWorkItemIdsForLinkWithAutoTestOk returns a tuple with the WorkItemIdsForLinkWithAutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestPutModel) SetWorkItemIdsForLinkWithAutoTest(v []string)`

SetWorkItemIdsForLinkWithAutoTest sets WorkItemIdsForLinkWithAutoTest field to given value.

### HasWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestPutModel) HasWorkItemIdsForLinkWithAutoTest() bool`

HasWorkItemIdsForLinkWithAutoTest returns a boolean if a field has been set.

### SetWorkItemIdsForLinkWithAutoTestNil

`func (o *AutoTestPutModel) SetWorkItemIdsForLinkWithAutoTestNil(b bool)`

 SetWorkItemIdsForLinkWithAutoTestNil sets the value for WorkItemIdsForLinkWithAutoTest to be an explicit nil

### UnsetWorkItemIdsForLinkWithAutoTest
`func (o *AutoTestPutModel) UnsetWorkItemIdsForLinkWithAutoTest()`

UnsetWorkItemIdsForLinkWithAutoTest ensures that no value is present for WorkItemIdsForLinkWithAutoTest, not even an explicit nil
### GetExternalId

`func (o *AutoTestPutModel) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoTestPutModel) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoTestPutModel) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinks

`func (o *AutoTestPutModel) GetLinks() []LinkPutModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTestPutModel) GetLinksOk() (*[]LinkPutModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTestPutModel) SetLinks(v []LinkPutModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTestPutModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTestPutModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTestPutModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetProjectId

`func (o *AutoTestPutModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestPutModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestPutModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *AutoTestPutModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestPutModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestPutModel) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AutoTestPutModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestPutModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestPutModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestPutModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestPutModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestPutModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassname

`func (o *AutoTestPutModel) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *AutoTestPutModel) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *AutoTestPutModel) SetClassname(v string)`

SetClassname sets Classname field to given value.

### HasClassname

`func (o *AutoTestPutModel) HasClassname() bool`

HasClassname returns a boolean if a field has been set.

### SetClassnameNil

`func (o *AutoTestPutModel) SetClassnameNil(b bool)`

 SetClassnameNil sets the value for Classname to be an explicit nil

### UnsetClassname
`func (o *AutoTestPutModel) UnsetClassname()`

UnsetClassname ensures that no value is present for Classname, not even an explicit nil
### GetSteps

`func (o *AutoTestPutModel) GetSteps() []AutoTestStepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestPutModel) GetStepsOk() (*[]AutoTestStepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestPutModel) SetSteps(v []AutoTestStepModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestPutModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestPutModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestPutModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetSetup

`func (o *AutoTestPutModel) GetSetup() []AutoTestStepModel`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AutoTestPutModel) GetSetupOk() (*[]AutoTestStepModel, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AutoTestPutModel) SetSetup(v []AutoTestStepModel)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AutoTestPutModel) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *AutoTestPutModel) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *AutoTestPutModel) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetTeardown

`func (o *AutoTestPutModel) GetTeardown() []AutoTestStepModel`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *AutoTestPutModel) GetTeardownOk() (*[]AutoTestStepModel, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *AutoTestPutModel) SetTeardown(v []AutoTestStepModel)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *AutoTestPutModel) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *AutoTestPutModel) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *AutoTestPutModel) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetTitle

`func (o *AutoTestPutModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestPutModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestPutModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestPutModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestPutModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestPutModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestPutModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestPutModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestPutModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestPutModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestPutModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestPutModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *AutoTestPutModel) GetLabels() []LabelPostModel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AutoTestPutModel) GetLabelsOk() (*[]LabelPostModel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AutoTestPutModel) SetLabels(v []LabelPostModel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AutoTestPutModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AutoTestPutModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AutoTestPutModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestPutModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestPutModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestPutModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTestPutModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTestPutModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTestPutModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


