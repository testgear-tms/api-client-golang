# TestPlanPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewTestPlanPostModel

`func NewTestPlanPostModel(name string, projectId string, ) *TestPlanPostModel`

NewTestPlanPostModel instantiates a new TestPlanPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanPostModelWithDefaults

`func NewTestPlanPostModelWithDefaults() *TestPlanPostModel`

NewTestPlanPostModelWithDefaults instantiates a new TestPlanPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TestPlanPostModel) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanPostModel) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanPostModel) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanPostModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPlanPostModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPlanPostModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetName

`func (o *TestPlanPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanPostModel) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *TestPlanPostModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TestPlanPostModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TestPlanPostModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TestPlanPostModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TestPlanPostModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TestPlanPostModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *TestPlanPostModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TestPlanPostModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TestPlanPostModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TestPlanPostModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TestPlanPostModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TestPlanPostModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetDescription

`func (o *TestPlanPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestPlanPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestPlanPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestPlanPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestPlanPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestPlanPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuild

`func (o *TestPlanPostModel) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *TestPlanPostModel) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *TestPlanPostModel) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *TestPlanPostModel) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### SetBuildNil

`func (o *TestPlanPostModel) SetBuildNil(b bool)`

 SetBuildNil sets the value for Build to be an explicit nil

### UnsetBuild
`func (o *TestPlanPostModel) UnsetBuild()`

UnsetBuild ensures that no value is present for Build, not even an explicit nil
### GetProjectId

`func (o *TestPlanPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestPlanPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestPlanPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProductName

`func (o *TestPlanPostModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *TestPlanPostModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *TestPlanPostModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *TestPlanPostModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *TestPlanPostModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *TestPlanPostModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetHasAutomaticDurationTimer

`func (o *TestPlanPostModel) GetHasAutomaticDurationTimer() bool`

GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field if non-nil, zero value otherwise.

### GetHasAutomaticDurationTimerOk

`func (o *TestPlanPostModel) GetHasAutomaticDurationTimerOk() (*bool, bool)`

GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomaticDurationTimer

`func (o *TestPlanPostModel) SetHasAutomaticDurationTimer(v bool)`

SetHasAutomaticDurationTimer sets HasAutomaticDurationTimer field to given value.

### HasHasAutomaticDurationTimer

`func (o *TestPlanPostModel) HasHasAutomaticDurationTimer() bool`

HasHasAutomaticDurationTimer returns a boolean if a field has been set.

### SetHasAutomaticDurationTimerNil

`func (o *TestPlanPostModel) SetHasAutomaticDurationTimerNil(b bool)`

 SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil

### UnsetHasAutomaticDurationTimer
`func (o *TestPlanPostModel) UnsetHasAutomaticDurationTimer()`

UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
### GetAttributes

`func (o *TestPlanPostModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanPostModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanPostModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanPostModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanPostModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanPostModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


