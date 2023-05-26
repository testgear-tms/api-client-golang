# AutotestFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | Specifies an autotest projects IDs to search for | [optional] 
**ExternalIds** | Pointer to **[]string** | Specifies an autotest external IDs to search for | [optional] 
**GlobalIds** | Pointer to **[]int64** | Specifies an autotest global IDs to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies an autotest name to search for | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Specifies an autotest flaky status to search for | [optional] 
**MustBeApproved** | Pointer to **NullableBool** | Specifies an autotest unapproved changes status to search for | [optional] 
**StabilityPercentage** | Pointer to [**Int64RangeSelectorModel**](Int64RangeSelectorModel.md) |  | [optional] 
**CreatedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies an autotest creator IDs to search for | [optional] 
**ModifiedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies an autotest last editor IDs to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies an autotest deleted status to search for | [optional] 
**Namespace** | Pointer to **NullableString** | Specifies an autotest namespace to search for | [optional] 
**IsEmptyNamespace** | Pointer to **NullableBool** | Specifies an autotest namespace name presence status to search for | [optional] 
**ClassName** | Pointer to **NullableString** | Specifies an autotest class name to search for | [optional] 
**IsEmptyClassName** | Pointer to **NullableBool** | Specifies an autotest class name presence status to search for | [optional] 
**LastTestResultOutcome** | Pointer to [**AutotestResultOutcome**](AutotestResultOutcome.md) |  | [optional] 

## Methods

### NewAutotestFilterModel

`func NewAutotestFilterModel() *AutotestFilterModel`

NewAutotestFilterModel instantiates a new AutotestFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestFilterModelWithDefaults

`func NewAutotestFilterModelWithDefaults() *AutotestFilterModel`

NewAutotestFilterModelWithDefaults instantiates a new AutotestFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *AutotestFilterModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *AutotestFilterModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *AutotestFilterModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *AutotestFilterModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *AutotestFilterModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *AutotestFilterModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetExternalIds

`func (o *AutotestFilterModel) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AutotestFilterModel) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AutotestFilterModel) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *AutotestFilterModel) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### SetExternalIdsNil

`func (o *AutotestFilterModel) SetExternalIdsNil(b bool)`

 SetExternalIdsNil sets the value for ExternalIds to be an explicit nil

### UnsetExternalIds
`func (o *AutotestFilterModel) UnsetExternalIds()`

UnsetExternalIds ensures that no value is present for ExternalIds, not even an explicit nil
### GetGlobalIds

`func (o *AutotestFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *AutotestFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *AutotestFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *AutotestFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *AutotestFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *AutotestFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetName

`func (o *AutotestFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotestFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotestFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutotestFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutotestFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutotestFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFlaky

`func (o *AutotestFilterModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutotestFilterModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutotestFilterModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutotestFilterModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutotestFilterModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutotestFilterModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetMustBeApproved

`func (o *AutotestFilterModel) GetMustBeApproved() bool`

GetMustBeApproved returns the MustBeApproved field if non-nil, zero value otherwise.

### GetMustBeApprovedOk

`func (o *AutotestFilterModel) GetMustBeApprovedOk() (*bool, bool)`

GetMustBeApprovedOk returns a tuple with the MustBeApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeApproved

`func (o *AutotestFilterModel) SetMustBeApproved(v bool)`

SetMustBeApproved sets MustBeApproved field to given value.

### HasMustBeApproved

`func (o *AutotestFilterModel) HasMustBeApproved() bool`

HasMustBeApproved returns a boolean if a field has been set.

### SetMustBeApprovedNil

`func (o *AutotestFilterModel) SetMustBeApprovedNil(b bool)`

 SetMustBeApprovedNil sets the value for MustBeApproved to be an explicit nil

### UnsetMustBeApproved
`func (o *AutotestFilterModel) UnsetMustBeApproved()`

UnsetMustBeApproved ensures that no value is present for MustBeApproved, not even an explicit nil
### GetStabilityPercentage

`func (o *AutotestFilterModel) GetStabilityPercentage() Int64RangeSelectorModel`

GetStabilityPercentage returns the StabilityPercentage field if non-nil, zero value otherwise.

### GetStabilityPercentageOk

`func (o *AutotestFilterModel) GetStabilityPercentageOk() (*Int64RangeSelectorModel, bool)`

GetStabilityPercentageOk returns a tuple with the StabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStabilityPercentage

`func (o *AutotestFilterModel) SetStabilityPercentage(v Int64RangeSelectorModel)`

SetStabilityPercentage sets StabilityPercentage field to given value.

### HasStabilityPercentage

`func (o *AutotestFilterModel) HasStabilityPercentage() bool`

HasStabilityPercentage returns a boolean if a field has been set.

### GetCreatedDate

`func (o *AutotestFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutotestFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutotestFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *AutotestFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedByIds

`func (o *AutotestFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *AutotestFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *AutotestFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *AutotestFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *AutotestFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *AutotestFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *AutotestFilterModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutotestFilterModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutotestFilterModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutotestFilterModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetModifiedByIds

`func (o *AutotestFilterModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *AutotestFilterModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *AutotestFilterModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *AutotestFilterModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *AutotestFilterModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *AutotestFilterModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *AutotestFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutotestFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutotestFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *AutotestFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *AutotestFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *AutotestFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetNamespace

`func (o *AutotestFilterModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutotestFilterModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutotestFilterModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutotestFilterModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutotestFilterModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutotestFilterModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetIsEmptyNamespace

`func (o *AutotestFilterModel) GetIsEmptyNamespace() bool`

GetIsEmptyNamespace returns the IsEmptyNamespace field if non-nil, zero value otherwise.

### GetIsEmptyNamespaceOk

`func (o *AutotestFilterModel) GetIsEmptyNamespaceOk() (*bool, bool)`

GetIsEmptyNamespaceOk returns a tuple with the IsEmptyNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyNamespace

`func (o *AutotestFilterModel) SetIsEmptyNamespace(v bool)`

SetIsEmptyNamespace sets IsEmptyNamespace field to given value.

### HasIsEmptyNamespace

`func (o *AutotestFilterModel) HasIsEmptyNamespace() bool`

HasIsEmptyNamespace returns a boolean if a field has been set.

### SetIsEmptyNamespaceNil

`func (o *AutotestFilterModel) SetIsEmptyNamespaceNil(b bool)`

 SetIsEmptyNamespaceNil sets the value for IsEmptyNamespace to be an explicit nil

### UnsetIsEmptyNamespace
`func (o *AutotestFilterModel) UnsetIsEmptyNamespace()`

UnsetIsEmptyNamespace ensures that no value is present for IsEmptyNamespace, not even an explicit nil
### GetClassName

`func (o *AutotestFilterModel) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AutotestFilterModel) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AutotestFilterModel) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *AutotestFilterModel) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *AutotestFilterModel) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *AutotestFilterModel) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil
### GetIsEmptyClassName

`func (o *AutotestFilterModel) GetIsEmptyClassName() bool`

GetIsEmptyClassName returns the IsEmptyClassName field if non-nil, zero value otherwise.

### GetIsEmptyClassNameOk

`func (o *AutotestFilterModel) GetIsEmptyClassNameOk() (*bool, bool)`

GetIsEmptyClassNameOk returns a tuple with the IsEmptyClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmptyClassName

`func (o *AutotestFilterModel) SetIsEmptyClassName(v bool)`

SetIsEmptyClassName sets IsEmptyClassName field to given value.

### HasIsEmptyClassName

`func (o *AutotestFilterModel) HasIsEmptyClassName() bool`

HasIsEmptyClassName returns a boolean if a field has been set.

### SetIsEmptyClassNameNil

`func (o *AutotestFilterModel) SetIsEmptyClassNameNil(b bool)`

 SetIsEmptyClassNameNil sets the value for IsEmptyClassName to be an explicit nil

### UnsetIsEmptyClassName
`func (o *AutotestFilterModel) UnsetIsEmptyClassName()`

UnsetIsEmptyClassName ensures that no value is present for IsEmptyClassName, not even an explicit nil
### GetLastTestResultOutcome

`func (o *AutotestFilterModel) GetLastTestResultOutcome() AutotestResultOutcome`

GetLastTestResultOutcome returns the LastTestResultOutcome field if non-nil, zero value otherwise.

### GetLastTestResultOutcomeOk

`func (o *AutotestFilterModel) GetLastTestResultOutcomeOk() (*AutotestResultOutcome, bool)`

GetLastTestResultOutcomeOk returns a tuple with the LastTestResultOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultOutcome

`func (o *AutotestFilterModel) SetLastTestResultOutcome(v AutotestResultOutcome)`

SetLastTestResultOutcome sets LastTestResultOutcome field to given value.

### HasLastTestResultOutcome

`func (o *AutotestFilterModel) HasLastTestResultOutcome() bool`

HasLastTestResultOutcome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


