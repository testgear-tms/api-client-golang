# TestPlanPutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewTestPlanPutModel

`func NewTestPlanPutModel(id string, name string, projectId string, ) *TestPlanPutModel`

NewTestPlanPutModel instantiates a new TestPlanPutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanPutModelWithDefaults

`func NewTestPlanPutModelWithDefaults() *TestPlanPutModel`

NewTestPlanPutModelWithDefaults instantiates a new TestPlanPutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPlanPutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPlanPutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPlanPutModel) SetId(v string)`

SetId sets Id field to given value.


### GetLockedById

`func (o *TestPlanPutModel) GetLockedById() string`

GetLockedById returns the LockedById field if non-nil, zero value otherwise.

### GetLockedByIdOk

`func (o *TestPlanPutModel) GetLockedByIdOk() (*string, bool)`

GetLockedByIdOk returns a tuple with the LockedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedById

`func (o *TestPlanPutModel) SetLockedById(v string)`

SetLockedById sets LockedById field to given value.

### HasLockedById

`func (o *TestPlanPutModel) HasLockedById() bool`

HasLockedById returns a boolean if a field has been set.

### SetLockedByIdNil

`func (o *TestPlanPutModel) SetLockedByIdNil(b bool)`

 SetLockedByIdNil sets the value for LockedById to be an explicit nil

### UnsetLockedById
`func (o *TestPlanPutModel) UnsetLockedById()`

UnsetLockedById ensures that no value is present for LockedById, not even an explicit nil
### GetTags

`func (o *TestPlanPutModel) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanPutModel) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanPutModel) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanPutModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPlanPutModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPlanPutModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetName

`func (o *TestPlanPutModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanPutModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanPutModel) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *TestPlanPutModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TestPlanPutModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TestPlanPutModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TestPlanPutModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TestPlanPutModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TestPlanPutModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *TestPlanPutModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TestPlanPutModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TestPlanPutModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TestPlanPutModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TestPlanPutModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TestPlanPutModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDescription

`func (o *TestPlanPutModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestPlanPutModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestPlanPutModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestPlanPutModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestPlanPutModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestPlanPutModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *TestPlanPutModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanPutModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanPutModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanPutModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestPlanPutModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestPlanPutModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProjectId

`func (o *TestPlanPutModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestPlanPutModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestPlanPutModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProductName

`func (o *TestPlanPutModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *TestPlanPutModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *TestPlanPutModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *TestPlanPutModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *TestPlanPutModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *TestPlanPutModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *TestPlanPutModel) GetHasAutomaticDurationTimer() bool`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *TestPlanPutModel) GetHasAutomaticDurationTimerOk() (*bool, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *TestPlanPutModel) SetHasAutomaticDurationTimer(v bool)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *TestPlanPutModel) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *TestPlanPutModel) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *TestPlanPutModel) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *TestPlanPutModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanPutModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanPutModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanPutModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanPutModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanPutModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


