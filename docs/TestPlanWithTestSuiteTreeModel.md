# TestPlanWithTestSuiteTreeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestSuites** | Pointer to [**[]TestSuiteWithChildrenModel**](TestSuiteWithChildrenModel.md) |  | [optional] 
**Status** | [**TestPlanStatusModel**](TestPlanStatusModel.md) |  | 
**StartedOn** | Pointer to **NullableTime** | Set when test plan is starter (status changed to: In Progress) | [optional] 
**CompletedOn** | Pointer to **NullableTime** | set when test plan status is completed (status changed to: Completed) | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**GlobalId** | Pointer to **int64** | Used for search Test plan | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**LockedDate** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** |  | 
**LockedById** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**[]TagShortModel**](TagShortModel.md) |  | [optional] 
**Name** | **string** |  | 
**StartDate** | Pointer to **NullableTime** | Used for analytics | [optional] 
**EndDate** | Pointer to **NullableTime** | Used for analytics | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Build** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | **string** |  | 
**ProductName** | Pointer to **NullableString** |  | [optional] 
**HasAutomaticDurationTimer** | Pointer to **NullableBool** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTestPlanWithTestSuiteTreeModel

`func NewTestPlanWithTestSuiteTreeModel(status TestPlanStatusModel, id string, name string, projectId string, ) *TestPlanWithTestSuiteTreeModel`

NewTestPlanWithTestSuiteTreeModel instantiates a new TestPlanWithTestSuiteTreeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanWithTestSuiteTreeModelWithDefaults

`func NewTestPlanWithTestSuiteTreeModelWithDefaults() *TestPlanWithTestSuiteTreeModel`

NewTestPlanWithTestSuiteTreeModelWithDefaults instantiates a new TestPlanWithTestSuiteTreeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestSuites

`func (o *TestPlanWithTestSuiteTreeModel) GetTestSuites() []TestSuiteWithChildrenModel`

GetTestSuites returns the TestSuites field if non-nil, zero value otherwise.

### GetTestSuitesOk

`func (o *TestPlanWithTestSuiteTreeModel) GetTestSuitesOk() (*[]TestSuiteWithChildrenModel, bool)`

GetTestSuitesOk returns a tuple with the TestSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuites

`func (o *TestPlanWithTestSuiteTreeModel) SetTestSuites(v []TestSuiteWithChildrenModel)`

SetTestSuites sets TestSuites field to given value.

### HasTestSuites

`func (o *TestPlanWithTestSuiteTreeModel) HasTestSuites() bool`

HasTestSuites returns a boolean if a field has been set.

### SetTestSuitesNil

`func (o *TestPlanWithTestSuiteTreeModel) SetTestSuitesNil(b bool)`

 SetTestSuitesNil sets the value for TestSuites to be an explicit nil

### UnsetTestSuites
`func (o *TestPlanWithTestSuiteTreeModel) UnsetTestSuites()`

UnsetTestSuites ensures that no value is present for TestSuites, not even an explicit nil
### GetStatus

`func (o *TestPlanWithTestSuiteTreeModel) GetStatus() TestPlanStatusModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPlanWithTestSuiteTreeModel) GetStatusOk() (*TestPlanStatusModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPlanWithTestSuiteTreeModel) SetStatus(v TestPlanStatusModel)`

SetStatus sets Status field to given value.


### GetStartedOn

`func (o *TestPlanWithTestSuiteTreeModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestPlanWithTestSuiteTreeModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestPlanWithTestSuiteTreeModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestPlanWithTestSuiteTreeModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestPlanWithTestSuiteTreeModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestPlanWithTestSuiteTreeModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestPlanWithTestSuiteTreeModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestPlanWithTestSuiteTreeModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestPlanWithTestSuiteTreeModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestPlanWithTestSuiteTreeModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestPlanWithTestSuiteTreeModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestPlanWithTestSuiteTreeModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetCreatedDate

`func (o *TestPlanWithTestSuiteTreeModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPlanWithTestSuiteTreeModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPlanWithTestSuiteTreeModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPlanWithTestSuiteTreeModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPlanWithTestSuiteTreeModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPlanWithTestSuiteTreeModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *TestPlanWithTestSuiteTreeModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPlanWithTestSuiteTreeModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPlanWithTestSuiteTreeModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPlanWithTestSuiteTreeModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPlanWithTestSuiteTreeModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPlanWithTestSuiteTreeModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestPlanWithTestSuiteTreeModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPlanWithTestSuiteTreeModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPlanWithTestSuiteTreeModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestPlanWithTestSuiteTreeModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *TestPlanWithTestSuiteTreeModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestPlanWithTestSuiteTreeModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestPlanWithTestSuiteTreeModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestPlanWithTestSuiteTreeModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestPlanWithTestSuiteTreeModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestPlanWithTestSuiteTreeModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetGlobalId

`func (o *TestPlanWithTestSuiteTreeModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *TestPlanWithTestSuiteTreeModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *TestPlanWithTestSuiteTreeModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *TestPlanWithTestSuiteTreeModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TestPlanWithTestSuiteTreeModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestPlanWithTestSuiteTreeModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestPlanWithTestSuiteTreeModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestPlanWithTestSuiteTreeModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLockedDate

`func (o *TestPlanWithTestSuiteTreeModel) GetLockedDate() time.Time`

GetLockedDate returns the LockedDate field if non-nil, zero value otherwise.

### GetLockedDateOk

`func (o *TestPlanWithTestSuiteTreeModel) GetLockedDateOk() (*time.Time, bool)`

GetLockedDateOk returns a tuple with the LockedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedDate

`func (o *TestPlanWithTestSuiteTreeModel) SetLockedDate(v time.Time)`

SetLockedDate sets LockedDate field to given value.

### HasLockedDate

`func (o *TestPlanWithTestSuiteTreeModel) HasLockedDate() bool`

HasLockedDate returns a boolean if a field has been set.

### SetLockedDateNil

`func (o *TestPlanWithTestSuiteTreeModel) SetLockedDateNil(b bool)`

 SetLockedDateNil sets the value for LockedDate to be an explicit nil

### UnsetLockedDate
`func (o *TestPlanWithTestSuiteTreeModel) UnsetLockedDate()`

UnsetLockedDate ensures that no value is present for LockedDate, not even an explicit nil
### GetId

`func (o *TestPlanWithTestSuiteTreeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPlanWithTestSuiteTreeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPlanWithTestSuiteTreeModel) SetId(v string)`

SetId sets Id field to given value.


### GetLockedById

`func (o *TestPlanWithTestSuiteTreeModel) GetLockedById() string`

GetLockedById returns the LockedById field if non-nil, zero value otherwise.

### GetLockedByIdOk

`func (o *TestPlanWithTestSuiteTreeModel) GetLockedByIdOk() (*string, bool)`

GetLockedByIdOk returns a tuple with the LockedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedById

`func (o *TestPlanWithTestSuiteTreeModel) SetLockedById(v string)`

SetLockedById sets LockedById field to given value.

### HasLockedById

`func (o *TestPlanWithTestSuiteTreeModel) HasLockedById() bool`

HasLockedById returns a boolean if a field has been set.

### SetLockedByIdNil

`func (o *TestPlanWithTestSuiteTreeModel) SetLockedByIdNil(b bool)`

 SetLockedByIdNil sets the value for LockedById to be an explicit nil

### UnsetLockedById
`func (o *TestPlanWithTestSuiteTreeModel) UnsetLockedById()`

UnsetLockedById ensures that no value is present for LockedById, not even an explicit nil
### GetTags

`func (o *TestPlanWithTestSuiteTreeModel) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanWithTestSuiteTreeModel) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanWithTestSuiteTreeModel) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanWithTestSuiteTreeModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPlanWithTestSuiteTreeModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPlanWithTestSuiteTreeModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetName

`func (o *TestPlanWithTestSuiteTreeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanWithTestSuiteTreeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanWithTestSuiteTreeModel) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *TestPlanWithTestSuiteTreeModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TestPlanWithTestSuiteTreeModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TestPlanWithTestSuiteTreeModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TestPlanWithTestSuiteTreeModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TestPlanWithTestSuiteTreeModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TestPlanWithTestSuiteTreeModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *TestPlanWithTestSuiteTreeModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TestPlanWithTestSuiteTreeModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TestPlanWithTestSuiteTreeModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TestPlanWithTestSuiteTreeModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TestPlanWithTestSuiteTreeModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TestPlanWithTestSuiteTreeModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDescription

`func (o *TestPlanWithTestSuiteTreeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestPlanWithTestSuiteTreeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestPlanWithTestSuiteTreeModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestPlanWithTestSuiteTreeModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestPlanWithTestSuiteTreeModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestPlanWithTestSuiteTreeModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *TestPlanWithTestSuiteTreeModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanWithTestSuiteTreeModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanWithTestSuiteTreeModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanWithTestSuiteTreeModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestPlanWithTestSuiteTreeModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestPlanWithTestSuiteTreeModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProjectId

`func (o *TestPlanWithTestSuiteTreeModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestPlanWithTestSuiteTreeModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestPlanWithTestSuiteTreeModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProductName

`func (o *TestPlanWithTestSuiteTreeModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *TestPlanWithTestSuiteTreeModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *TestPlanWithTestSuiteTreeModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *TestPlanWithTestSuiteTreeModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *TestPlanWithTestSuiteTreeModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *TestPlanWithTestSuiteTreeModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *TestPlanWithTestSuiteTreeModel) GetHasAutomaticDurationTimer() bool`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *TestPlanWithTestSuiteTreeModel) GetHasAutomaticDurationTimerOk() (*bool, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *TestPlanWithTestSuiteTreeModel) SetHasAutomaticDurationTimer(v bool)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *TestPlanWithTestSuiteTreeModel) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *TestPlanWithTestSuiteTreeModel) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *TestPlanWithTestSuiteTreeModel) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *TestPlanWithTestSuiteTreeModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanWithTestSuiteTreeModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanWithTestSuiteTreeModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanWithTestSuiteTreeModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanWithTestSuiteTreeModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanWithTestSuiteTreeModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


