# ProjectTestPlansFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**[]TestPlanStatusModel**](TestPlanStatusModel.md) |  | [optional] 
**GlobalIds** | Pointer to **[]int64** |  | [optional] 
**IsLocked** | Pointer to **NullableBool** |  | [optional] 
**LockedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**AutomaticDurationTimer** | Pointer to **[]bool** |  | [optional] 
**CreatedByIds** | Pointer to **[]string** |  | [optional] 
**CreatedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**StartDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**EndDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to **map[string][]string** |  | [optional] 
**IsDeleted** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProjectTestPlansFilterModel

`func NewProjectTestPlansFilterModel() *ProjectTestPlansFilterModel`

NewProjectTestPlansFilterModel instantiates a new ProjectTestPlansFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTestPlansFilterModelWithDefaults

`func NewProjectTestPlansFilterModelWithDefaults() *ProjectTestPlansFilterModel`

NewProjectTestPlansFilterModelWithDefaults instantiates a new ProjectTestPlansFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectTestPlansFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectTestPlansFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectTestPlansFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectTestPlansFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectTestPlansFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectTestPlansFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ProjectTestPlansFilterModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectTestPlansFilterModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectTestPlansFilterModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectTestPlansFilterModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProjectTestPlansFilterModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProjectTestPlansFilterModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *ProjectTestPlansFilterModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ProjectTestPlansFilterModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ProjectTestPlansFilterModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ProjectTestPlansFilterModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *ProjectTestPlansFilterModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *ProjectTestPlansFilterModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProductName

`func (o *ProjectTestPlansFilterModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ProjectTestPlansFilterModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ProjectTestPlansFilterModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ProjectTestPlansFilterModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *ProjectTestPlansFilterModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *ProjectTestPlansFilterModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetStatus

`func (o *ProjectTestPlansFilterModel) GetStatus() []TestPlanStatusModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectTestPlansFilterModel) GetStatusOk() (*[]TestPlanStatusModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectTestPlansFilterModel) SetStatus(v []TestPlanStatusModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectTestPlansFilterModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProjectTestPlansFilterModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProjectTestPlansFilterModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetGlobalIds

`func (o *ProjectTestPlansFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ProjectTestPlansFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ProjectTestPlansFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ProjectTestPlansFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ProjectTestPlansFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ProjectTestPlansFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetIsLocked

`func (o *ProjectTestPlansFilterModel) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectTestPlansFilterModel) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectTestPlansFilterModel) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ProjectTestPlansFilterModel) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLockedNil

`func (o *ProjectTestPlansFilterModel) SetIsLockedNil(b bool)`

 SetIsLockedNil sets the value for IsLocked to be an explicit nil

### UnsetIsLocked
`func (o *ProjectTestPlansFilterModel) UnsetIsLocked()`

UnsetIsLocked ensures that no value is present for IsLocked, not even an explicit nil
### GetLockedDate

`func (o *ProjectTestPlansFilterModel) GetLockedDate() DateTimeRangeSelectorModel`

GetLockedDate returns the LockedDate field if non-nil, zero value otherwise.

### GetLockedDateOk

`func (o *ProjectTestPlansFilterModel) GetLockedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetLockedDateOk returns a tuple with the LockedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedDate

`func (o *ProjectTestPlansFilterModel) SetLockedDate(v DateTimeRangeSelectorModel)`

SetLockedDate sets LockedDate field to given value.

### HasLockedDate

`func (o *ProjectTestPlansFilterModel) HasLockedDate() bool`

HasLockedDate returns a boolean if a field has been set.

### GetAutomaticDurationTimer

`func (o *ProjectTestPlansFilterModel) GetAutomaticDurationTimer() []bool`

GetAutomaticDurationTimer returns the AutomaticDurationTimer field if non-nil, zero value otherwise.

### GetAutomaticDurationTimerOk

`func (o *ProjectTestPlansFilterModel) GetAutomaticDurationTimerOk() (*[]bool, bool)`

GetAutomaticDurationTimerOk returns a tuple with the AutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDurationTimer

`func (o *ProjectTestPlansFilterModel) SetAutomaticDurationTimer(v []bool)`

SetAutomaticDurationTimer sets AutomaticDurationTimer field to given value.

### HasAutomaticDurationTimer

`func (o *ProjectTestPlansFilterModel) HasAutomaticDurationTimer() bool`

HasAutomaticDurationTimer returns a boolean if a field has been set.

### SetAutomaticDurationTimerNil

`func (o *ProjectTestPlansFilterModel) SetAutomaticDurationTimerNil(b bool)`

 SetAutomaticDurationTimerNil sets the value for AutomaticDurationTimer to be an explicit nil

### UnsetAutomaticDurationTimer
`func (o *ProjectTestPlansFilterModel) UnsetAutomaticDurationTimer()`

UnsetAutomaticDurationTimer ensures that no value is present for AutomaticDurationTimer, not even an explicit nil
### GetCreatedByIds

`func (o *ProjectTestPlansFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ProjectTestPlansFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ProjectTestPlansFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ProjectTestPlansFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ProjectTestPlansFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ProjectTestPlansFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetCreatedDate

`func (o *ProjectTestPlansFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectTestPlansFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectTestPlansFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ProjectTestPlansFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetStartDate

`func (o *ProjectTestPlansFilterModel) GetStartDate() DateTimeRangeSelectorModel`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProjectTestPlansFilterModel) GetStartDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProjectTestPlansFilterModel) SetStartDate(v DateTimeRangeSelectorModel)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProjectTestPlansFilterModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProjectTestPlansFilterModel) GetEndDate() DateTimeRangeSelectorModel`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProjectTestPlansFilterModel) GetEndDateOk() (*DateTimeRangeSelectorModel, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProjectTestPlansFilterModel) SetEndDate(v DateTimeRangeSelectorModel)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProjectTestPlansFilterModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTagNames

`func (o *ProjectTestPlansFilterModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *ProjectTestPlansFilterModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *ProjectTestPlansFilterModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *ProjectTestPlansFilterModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *ProjectTestPlansFilterModel) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *ProjectTestPlansFilterModel) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetAttributes

`func (o *ProjectTestPlansFilterModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProjectTestPlansFilterModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProjectTestPlansFilterModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProjectTestPlansFilterModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ProjectTestPlansFilterModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ProjectTestPlansFilterModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *ProjectTestPlansFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectTestPlansFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectTestPlansFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProjectTestPlansFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ProjectTestPlansFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ProjectTestPlansFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


