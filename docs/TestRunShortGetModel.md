# TestRunShortGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the test run | [optional] 
**Name** | Pointer to **NullableString** | Name of the test run | [optional] 
**ProjectId** | Pointer to **string** | Unique ID of project where test run is located | [optional] 
**CreatedDate** | Pointer to **time.Time** | Date when the test run was created | [optional] 
**CreatedById** | Pointer to **string** | Unique ID of user who created the test run | [optional] 
**ModifiedDate** | Pointer to **NullableTime** | Date when the test run was modified last time | [optional] 
**ModifiedById** | Pointer to **NullableString** | Unique ID of user who modified the test run last time | [optional] 
**IsDeleted** | Pointer to **bool** | Is the test run is deleted | [optional] 
**State** | [**TestRunState**](TestRunState.md) |  | 
**StartedDate** | Pointer to **NullableTime** | Date when the test run was started | [optional] 
**AutotestsCount** | Pointer to **int32** | Number of autotests run in the test run | [optional] 
**Statistics** | [**TestResultsStatisticsGetModel**](TestResultsStatisticsGetModel.md) |  | 

## Methods

### NewTestRunShortGetModel

`func NewTestRunShortGetModel(state TestRunState, statistics TestResultsStatisticsGetModel, ) *TestRunShortGetModel`

NewTestRunShortGetModel instantiates a new TestRunShortGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunShortGetModelWithDefaults

`func NewTestRunShortGetModelWithDefaults() *TestRunShortGetModel`

NewTestRunShortGetModelWithDefaults instantiates a new TestRunShortGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRunShortGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunShortGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunShortGetModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestRunShortGetModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TestRunShortGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunShortGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunShortGetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunShortGetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunShortGetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunShortGetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *TestRunShortGetModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunShortGetModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunShortGetModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TestRunShortGetModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TestRunShortGetModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestRunShortGetModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestRunShortGetModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestRunShortGetModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedById

`func (o *TestRunShortGetModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestRunShortGetModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestRunShortGetModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestRunShortGetModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedDate

`func (o *TestRunShortGetModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestRunShortGetModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestRunShortGetModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestRunShortGetModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestRunShortGetModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestRunShortGetModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *TestRunShortGetModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestRunShortGetModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestRunShortGetModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestRunShortGetModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestRunShortGetModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestRunShortGetModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetIsDeleted

`func (o *TestRunShortGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunShortGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunShortGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestRunShortGetModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetState

`func (o *TestRunShortGetModel) GetState() TestRunState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestRunShortGetModel) GetStateOk() (*TestRunState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestRunShortGetModel) SetState(v TestRunState)`

SetState sets State field to given value.


### GetStartedDate

`func (o *TestRunShortGetModel) GetStartedDate() time.Time`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *TestRunShortGetModel) GetStartedDateOk() (*time.Time, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *TestRunShortGetModel) SetStartedDate(v time.Time)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *TestRunShortGetModel) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *TestRunShortGetModel) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *TestRunShortGetModel) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetAutotestsCount

`func (o *TestRunShortGetModel) GetAutotestsCount() int32`

GetAutotestsCount returns the AutotestsCount field if non-nil, zero value otherwise.

### GetAutotestsCountOk

`func (o *TestRunShortGetModel) GetAutotestsCountOk() (*int32, bool)`

GetAutotestsCountOk returns a tuple with the AutotestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestsCount

`func (o *TestRunShortGetModel) SetAutotestsCount(v int32)`

SetAutotestsCount sets AutotestsCount field to given value.

### HasAutotestsCount

`func (o *TestRunShortGetModel) HasAutotestsCount() bool`

HasAutotestsCount returns a boolean if a field has been set.

### GetStatistics

`func (o *TestRunShortGetModel) GetStatistics() TestResultsStatisticsGetModel`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *TestRunShortGetModel) GetStatisticsOk() (*TestResultsStatisticsGetModel, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *TestRunShortGetModel) SetStatistics(v TestResultsStatisticsGetModel)`

SetStatistics sets Statistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


