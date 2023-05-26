# ProjectsFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies a project name to search for | [optional] 
**IsFavorite** | Pointer to **NullableBool** | Specifies a project favorite status to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies a project deleted status to search for | [optional] 
**TestCasesCount** | Pointer to [**Int32RangeSelectorModel**](Int32RangeSelectorModel.md) |  | [optional] 
**ChecklistsCount** | Pointer to [**Int32RangeSelectorModel**](Int32RangeSelectorModel.md) |  | [optional] 
**SharedStepsCount** | Pointer to [**Int32RangeSelectorModel**](Int32RangeSelectorModel.md) |  | [optional] 
**AutotestsCount** | Pointer to [**Int32RangeSelectorModel**](Int32RangeSelectorModel.md) |  | [optional] 
**GlobalIds** | Pointer to **[]int64** | Specifies a project global IDs to search for | [optional] 
**CreatedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies an autotest creator IDs to search for | [optional] 

## Methods

### NewProjectsFilterModel

`func NewProjectsFilterModel() *ProjectsFilterModel`

NewProjectsFilterModel instantiates a new ProjectsFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsFilterModelWithDefaults

`func NewProjectsFilterModelWithDefaults() *ProjectsFilterModel`

NewProjectsFilterModelWithDefaults instantiates a new ProjectsFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectsFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectsFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsFavorite

`func (o *ProjectsFilterModel) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ProjectsFilterModel) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ProjectsFilterModel) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ProjectsFilterModel) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *ProjectsFilterModel) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *ProjectsFilterModel) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetIsDeleted

`func (o *ProjectsFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectsFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectsFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProjectsFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ProjectsFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ProjectsFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetTestCasesCount

`func (o *ProjectsFilterModel) GetTestCasesCount() Int32RangeSelectorModel`

GetTestCasesCount returns the TestCasesCount field if non-nil, zero value otherwise.

### GetTestCasesCountOk

`func (o *ProjectsFilterModel) GetTestCasesCountOk() (*Int32RangeSelectorModel, bool)`

GetTestCasesCountOk returns a tuple with the TestCasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCasesCount

`func (o *ProjectsFilterModel) SetTestCasesCount(v Int32RangeSelectorModel)`

SetTestCasesCount sets TestCasesCount field to given value.

### HasTestCasesCount

`func (o *ProjectsFilterModel) HasTestCasesCount() bool`

HasTestCasesCount returns a boolean if a field has been set.

### GetChecklistsCount

`func (o *ProjectsFilterModel) GetChecklistsCount() Int32RangeSelectorModel`

GetChecklistsCount returns the ChecklistsCount field if non-nil, zero value otherwise.

### GetChecklistsCountOk

`func (o *ProjectsFilterModel) GetChecklistsCountOk() (*Int32RangeSelectorModel, bool)`

GetChecklistsCountOk returns a tuple with the ChecklistsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklistsCount

`func (o *ProjectsFilterModel) SetChecklistsCount(v Int32RangeSelectorModel)`

SetChecklistsCount sets ChecklistsCount field to given value.

### HasChecklistsCount

`func (o *ProjectsFilterModel) HasChecklistsCount() bool`

HasChecklistsCount returns a boolean if a field has been set.

### GetSharedStepsCount

`func (o *ProjectsFilterModel) GetSharedStepsCount() Int32RangeSelectorModel`

GetSharedStepsCount returns the SharedStepsCount field if non-nil, zero value otherwise.

### GetSharedStepsCountOk

`func (o *ProjectsFilterModel) GetSharedStepsCountOk() (*Int32RangeSelectorModel, bool)`

GetSharedStepsCountOk returns a tuple with the SharedStepsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepsCount

`func (o *ProjectsFilterModel) SetSharedStepsCount(v Int32RangeSelectorModel)`

SetSharedStepsCount sets SharedStepsCount field to given value.

### HasSharedStepsCount

`func (o *ProjectsFilterModel) HasSharedStepsCount() bool`

HasSharedStepsCount returns a boolean if a field has been set.

### GetAutotestsCount

`func (o *ProjectsFilterModel) GetAutotestsCount() Int32RangeSelectorModel`

GetAutotestsCount returns the AutotestsCount field if non-nil, zero value otherwise.

### GetAutotestsCountOk

`func (o *ProjectsFilterModel) GetAutotestsCountOk() (*Int32RangeSelectorModel, bool)`

GetAutotestsCountOk returns a tuple with the AutotestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestsCount

`func (o *ProjectsFilterModel) SetAutotestsCount(v Int32RangeSelectorModel)`

SetAutotestsCount sets AutotestsCount field to given value.

### HasAutotestsCount

`func (o *ProjectsFilterModel) HasAutotestsCount() bool`

HasAutotestsCount returns a boolean if a field has been set.

### GetGlobalIds

`func (o *ProjectsFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ProjectsFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ProjectsFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ProjectsFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ProjectsFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ProjectsFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetCreatedDate

`func (o *ProjectsFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectsFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectsFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ProjectsFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedByIds

`func (o *ProjectsFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ProjectsFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ProjectsFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ProjectsFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ProjectsFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ProjectsFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


