# TestPlanWithAnalyticModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analytic** | Pointer to [**TestPointAnalyticResult**](TestPointAnalyticResult.md) |  | [optional] 
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

### NewTestPlanWithAnalyticModel

`func NewTestPlanWithAnalyticModel(status TestPlanStatusModel, id string, name string, projectId string, ) *TestPlanWithAnalyticModel`

NewTestPlanWithAnalyticModel instantiates a new TestPlanWithAnalyticModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanWithAnalyticModelWithDefaults

`func NewTestPlanWithAnalyticModelWithDefaults() *TestPlanWithAnalyticModel`

NewTestPlanWithAnalyticModelWithDefaults instantiates a new TestPlanWithAnalyticModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalytic

`func (o *TestPlanWithAnalyticModel) GetAnalytic() TestPointAnalyticResult`

GetAnalytic returns the Analytic field if non-nil, zero value otherwise.

### GetAnalyticOk

`func (o *TestPlanWithAnalyticModel) GetAnalyticOk() (*TestPointAnalyticResult, bool)`

GetAnalyticOk returns a tuple with the Analytic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytic

`func (o *TestPlanWithAnalyticModel) SetAnalytic(v TestPointAnalyticResult)`

SetAnalytic sets Analytic field to given value.

### HasAnalytic

`func (o *TestPlanWithAnalyticModel) HasAnalytic() bool`

HasAnalytic returns a boolean if a field has been set.

### GetStatus

`func (o *TestPlanWithAnalyticModel) GetStatus() TestPlanStatusModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPlanWithAnalyticModel) GetStatusOk() (*TestPlanStatusModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPlanWithAnalyticModel) SetStatus(v TestPlanStatusModel)`

SetStatus sets Status field to given value.


### GetStartedOn

`func (o *TestPlanWithAnalyticModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestPlanWithAnalyticModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestPlanWithAnalyticModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestPlanWithAnalyticModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestPlanWithAnalyticModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestPlanWithAnalyticModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestPlanWithAnalyticModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestPlanWithAnalyticModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestPlanWithAnalyticModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestPlanWithAnalyticModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestPlanWithAnalyticModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestPlanWithAnalyticModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetCreatedDate

`func (o *TestPlanWithAnalyticModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPlanWithAnalyticModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPlanWithAnalyticModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPlanWithAnalyticModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPlanWithAnalyticModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPlanWithAnalyticModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *TestPlanWithAnalyticModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPlanWithAnalyticModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPlanWithAnalyticModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPlanWithAnalyticModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPlanWithAnalyticModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPlanWithAnalyticModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestPlanWithAnalyticModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPlanWithAnalyticModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPlanWithAnalyticModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestPlanWithAnalyticModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *TestPlanWithAnalyticModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestPlanWithAnalyticModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestPlanWithAnalyticModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestPlanWithAnalyticModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestPlanWithAnalyticModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestPlanWithAnalyticModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetGlobalId

`func (o *TestPlanWithAnalyticModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *TestPlanWithAnalyticModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *TestPlanWithAnalyticModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *TestPlanWithAnalyticModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TestPlanWithAnalyticModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestPlanWithAnalyticModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestPlanWithAnalyticModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestPlanWithAnalyticModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLockedDate

`func (o *TestPlanWithAnalyticModel) GetLockedDate() time.Time`

GetLockedDate returns the LockedDate field if non-nil, zero value otherwise.

### GetLockedDateOk

`func (o *TestPlanWithAnalyticModel) GetLockedDateOk() (*time.Time, bool)`

GetLockedDateOk returns a tuple with the LockedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedDate

`func (o *TestPlanWithAnalyticModel) SetLockedDate(v time.Time)`

SetLockedDate sets LockedDate field to given value.

### HasLockedDate

`func (o *TestPlanWithAnalyticModel) HasLockedDate() bool`

HasLockedDate returns a boolean if a field has been set.

### SetLockedDateNil

`func (o *TestPlanWithAnalyticModel) SetLockedDateNil(b bool)`

 SetLockedDateNil sets the value for LockedDate to be an explicit nil

### UnsetLockedDate
`func (o *TestPlanWithAnalyticModel) UnsetLockedDate()`

UnsetLockedDate ensures that no value is present for LockedDate, not even an explicit nil
### GetId

`func (o *TestPlanWithAnalyticModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPlanWithAnalyticModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPlanWithAnalyticModel) SetId(v string)`

SetId sets Id field to given value.


### GetLockedById

`func (o *TestPlanWithAnalyticModel) GetLockedById() string`

GetLockedById returns the LockedById field if non-nil, zero value otherwise.

### GetLockedByIdOk

`func (o *TestPlanWithAnalyticModel) GetLockedByIdOk() (*string, bool)`

GetLockedByIdOk returns a tuple with the LockedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedById

`func (o *TestPlanWithAnalyticModel) SetLockedById(v string)`

SetLockedById sets LockedById field to given value.

### HasLockedById

`func (o *TestPlanWithAnalyticModel) HasLockedById() bool`

HasLockedById returns a boolean if a field has been set.

### SetLockedByIdNil

`func (o *TestPlanWithAnalyticModel) SetLockedByIdNil(b bool)`

 SetLockedByIdNil sets the value for LockedById to be an explicit nil

### UnsetLockedById
`func (o *TestPlanWithAnalyticModel) UnsetLockedById()`

UnsetLockedById ensures that no value is present for LockedById, not even an explicit nil
### GetTags

`func (o *TestPlanWithAnalyticModel) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanWithAnalyticModel) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanWithAnalyticModel) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanWithAnalyticModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPlanWithAnalyticModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPlanWithAnalyticModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetName

`func (o *TestPlanWithAnalyticModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanWithAnalyticModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanWithAnalyticModel) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *TestPlanWithAnalyticModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TestPlanWithAnalyticModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TestPlanWithAnalyticModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TestPlanWithAnalyticModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TestPlanWithAnalyticModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TestPlanWithAnalyticModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *TestPlanWithAnalyticModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TestPlanWithAnalyticModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TestPlanWithAnalyticModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TestPlanWithAnalyticModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TestPlanWithAnalyticModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TestPlanWithAnalyticModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDescription

`func (o *TestPlanWithAnalyticModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestPlanWithAnalyticModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestPlanWithAnalyticModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestPlanWithAnalyticModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestPlanWithAnalyticModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestPlanWithAnalyticModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *TestPlanWithAnalyticModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanWithAnalyticModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanWithAnalyticModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanWithAnalyticModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestPlanWithAnalyticModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestPlanWithAnalyticModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProjectId

`func (o *TestPlanWithAnalyticModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestPlanWithAnalyticModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestPlanWithAnalyticModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProductName

`func (o *TestPlanWithAnalyticModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *TestPlanWithAnalyticModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *TestPlanWithAnalyticModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *TestPlanWithAnalyticModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *TestPlanWithAnalyticModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *TestPlanWithAnalyticModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *TestPlanWithAnalyticModel) GetHasAutomaticDurationTimer() bool`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *TestPlanWithAnalyticModel) GetHasAutomaticDurationTimerOk() (*bool, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *TestPlanWithAnalyticModel) SetHasAutomaticDurationTimer(v bool)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *TestPlanWithAnalyticModel) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *TestPlanWithAnalyticModel) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *TestPlanWithAnalyticModel) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *TestPlanWithAnalyticModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanWithAnalyticModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanWithAnalyticModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanWithAnalyticModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanWithAnalyticModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanWithAnalyticModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


