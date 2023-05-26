# AutoTestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalId** | Pointer to **int64** | Global ID of the autotest | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the autotest is deleted | [optional] 
**MustBeApproved** | Pointer to **bool** | Indicates if the autotest has unapproved changes from linked work items | [optional] 
**Id** | Pointer to **string** | Unique ID of the autotest | [optional] 
**CreatedDate** | Pointer to **time.Time** | Creation date of the autotest | [optional] 
**ModifiedDate** | Pointer to **NullableTime** | Last modification date of the project | [optional] 
**CreatedById** | Pointer to **string** | Unique ID of the project creator | [optional] 
**ModifiedById** | Pointer to **NullableString** | Unique ID of the project last editor | [optional] 
**LastTestRunId** | Pointer to **NullableString** | Unique ID of the autotest last test run | [optional] 
**LastTestRunName** | Pointer to **NullableString** | Name of the autotest last test run | [optional] 
**LastTestResultId** | Pointer to **NullableString** | Unique ID of the autotest last test result | [optional] 
**LastTestResultOutcome** | Pointer to **NullableString** | Outcome of the autotest last test result | [optional] 
**StabilityPercentage** | Pointer to **NullableInt32** | Stability percentage of the autotest | [optional] 
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
**Labels** | Pointer to [**[]LabelShortModel**](LabelShortModel.md) | Collection of the autotest labels | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Indicates if the autotest is marked as flaky | [optional] 

## Methods

### NewAutoTestModel

`func NewAutoTestModel(externalId string, projectId string, name string, ) *AutoTestModel`

NewAutoTestModel instantiates a new AutoTestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestModelWithDefaults

`func NewAutoTestModelWithDefaults() *AutoTestModel`

NewAutoTestModelWithDefaults instantiates a new AutoTestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalId

`func (o *AutoTestModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *AutoTestModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *AutoTestModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *AutoTestModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *AutoTestModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutoTestModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutoTestModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AutoTestModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetMustBeApproved

`func (o *AutoTestModel) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutoTestModel) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutoTestModel) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.

### HasMustBeApproved

`func (o *AutoTestModel) HasMustBeApproved() bool`

HasMustBeApproved returns a boolean if a field has been set.

### GetId

`func (o *AutoTestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTestModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoTestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *AutoTestModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTestModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTestModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AutoTestModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *AutoTestModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTestModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTestModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTestModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTestModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTestModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *AutoTestModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutoTestModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutoTestModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *AutoTestModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *AutoTestModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutoTestModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutoTestModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutoTestModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutoTestModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutoTestModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetLastTestRunId

`func (o *AutoTestModel) GetLastTestRunId() string`

GetLastTestRunId returns the LastTestRunId field if non-nil, zero value otherwise.

### GetLastTestRunIdOk

`func (o *AutoTestModel) GetLastTestRunIdOk() (*string, bool)`

GetLastTestRunIdOk returns a tuple with the LastTestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestRunId

`func (o *AutoTestModel) SetLastTestRunId(v string)`

SetLastTestRunId sets LastTestRunId field to given value.

### HasLastTestRunId

`func (o *AutoTestModel) HasLastTestRunId() bool`

HasLastTestRunId returns a boolean if a field has been set.

### SetLastTestRunIdNil

`func (o *AutoTestModel) SetLastTestRunIdNil(b bool)`

 SetLastTestRunIdNil sets the value for LastTestRunId to be an explicit nil

### UnsetLastTestRunId
`func (o *AutoTestModel) UnsetLastTestRunId()`

UnsetLastTestRunId ensures that no value is present for LastTestRunId, not even an explicit nil
### GetLastTestRunName

`func (o *AutoTestModel) GetLastTestRunName() string`

GetLastTestRunName returns the LastTestRunName field if non-nil, zero value otherwise.

### GetLastTestRunNameOk

`func (o *AutoTestModel) GetLastTestRunNameOk() (*string, bool)`

GetLastTestRunNameOk returns a tuple with the LastTestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestRunName

`func (o *AutoTestModel) SetLastTestRunName(v string)`

SetLastTestRunName sets LastTestRunName field to given value.

### HasLastTestRunName

`func (o *AutoTestModel) HasLastTestRunName() bool`

HasLastTestRunName returns a boolean if a field has been set.

### SetLastTestRunNameNil

`func (o *AutoTestModel) SetLastTestRunNameNil(b bool)`

 SetLastTestRunNameNil sets the value for LastTestRunName to be an explicit nil

### UnsetLastTestRunName
`func (o *AutoTestModel) UnsetLastTestRunName()`

UnsetLastTestRunName ensures that no value is present for LastTestRunName, not even an explicit nil
### GetLastTestResultId

`func (o *AutoTestModel) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *AutoTestModel) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *AutoTestModel) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *AutoTestModel) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *AutoTestModel) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *AutoTestModel) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutoTestModel) GetLastTestResultOutcome() string`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutoTestModel) GetLastTestResultOutcomeOk() (*string, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutoTestModel) SetLastTestResultOutcome(v string)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutoTestModel) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.

### SetLastTestResultOutcomeNil

`func (o *AutoTestModel) SetLastTestResultOutcomeNil(b bool)`

 SetLastTestResultOutcomeNil sets the value for LastTestResultOutcome to be an explicit nil

### UnsetLastTestResultOutcome
`func (o *AutoTestModel) UnsetLastTestResultOutcome()`

UnsetLastTestResultOutcome ensures that no value is present for LastTestResultOutcome, not even an explicit nil
### GetStabilityPercentage

`func (o *AutoTestModel) GetStabilityPercentage() int32`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutoTestModel) GetStabilityPercentageOk() (*int32, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutoTestModel) SetStabilityPercentage(v int32)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutoTestModel) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### SetStabilityPercentageNil

`func (o *AutoTestModel) SetStabilityPercentageNil(b bool)`

 SetStabilityPercentageNil sets the value for StabilityPercentage to be an explicit nil

### UnsetStabilityPercentage
`func (o *AutoTestModel) UnsetStabilityPercentage()`

UnsetStabilityPercentage ensures that no value is present for StabilityPercentage, not even an explicit nil
### GetExternalId

`func (o *AutoTestModel) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoTestModel) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoTestModel) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinks

`func (o *AutoTestModel) GetLinks() []LinkPutModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTestModel) GetLinksOk() (*[]LinkPutModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTestModel) SetLinks(v []LinkPutModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTestModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTestModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTestModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetProjectId

`func (o *AutoTestModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *AutoTestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestModel) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AutoTestModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassname

`func (o *AutoTestModel) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *AutoTestModel) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *AutoTestModel) SetClassname(v string)`

SetClassname sets Classname field to given value.

### HasClassname

`func (o *AutoTestModel) HasClassname() bool`

HasClassname returns a boolean if a field has been set.

### SetClassnameNil

`func (o *AutoTestModel) SetClassnameNil(b bool)`

 SetClassnameNil sets the value for Classname to be an explicit nil

### UnsetClassname
`func (o *AutoTestModel) UnsetClassname()`

UnsetClassname ensures that no value is present for Classname, not even an explicit nil
### GetSteps

`func (o *AutoTestModel) GetSteps() []AutoTestStepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestModel) GetStepsOk() (*[]AutoTestStepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestModel) SetSteps(v []AutoTestStepModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetSetup

`func (o *AutoTestModel) GetSetup() []AutoTestStepModel`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AutoTestModel) GetSetupOk() (*[]AutoTestStepModel, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AutoTestModel) SetSetup(v []AutoTestStepModel)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AutoTestModel) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *AutoTestModel) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *AutoTestModel) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetTeardown

`func (o *AutoTestModel) GetTeardown() []AutoTestStepModel`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *AutoTestModel) GetTeardownOk() (*[]AutoTestStepModel, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *AutoTestModel) SetTeardown(v []AutoTestStepModel)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *AutoTestModel) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *AutoTestModel) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *AutoTestModel) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetTitle

`func (o *AutoTestModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *AutoTestModel) GetLabels() []LabelShortModel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AutoTestModel) GetLabelsOk() (*[]LabelShortModel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AutoTestModel) SetLabels(v []LabelShortModel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AutoTestModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AutoTestModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AutoTestModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTestModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTestModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTestModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


