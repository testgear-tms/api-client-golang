# TestRunSearchQueryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**States** | Pointer to [**[]TestRunState**](TestRunState.md) |  | [optional] 
**StartedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CompletedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** |  | [optional] 
**ModifiedByIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTestRunSearchQueryModel

`func NewTestRunSearchQueryModel() *TestRunSearchQueryModel`

NewTestRunSearchQueryModel instantiates a new TestRunSearchQueryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunSearchQueryModelWithDefaults

`func NewTestRunSearchQueryModelWithDefaults() *TestRunSearchQueryModel`

NewTestRunSearchQueryModelWithDefaults instantiates a new TestRunSearchQueryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestRunSearchQueryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunSearchQueryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunSearchQueryModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunSearchQueryModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunSearchQueryModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunSearchQueryModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStates

`func (o *TestRunSearchQueryModel) GetStates() []TestRunState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *TestRunSearchQueryModel) GetStatesOk() (*[]TestRunState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *TestRunSearchQueryModel) SetStates(v []TestRunState)`

SetStates sets States field to given value.

### HasStates

`func (o *TestRunSearchQueryModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *TestRunSearchQueryModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *TestRunSearchQueryModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetStartedDate

`func (o *TestRunSearchQueryModel) GetStartedDate() DateTimeRangeSelectorModel`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *TestRunSearchQueryModel) GetStartedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *TestRunSearchQueryModel) SetStartedDate(v DateTimeRangeSelectorModel)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *TestRunSearchQueryModel) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### GetCompletedDate

`func (o *TestRunSearchQueryModel) GetCompletedDate() DateTimeRangeSelectorModel`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *TestRunSearchQueryModel) GetCompletedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *TestRunSearchQueryModel) SetCompletedDate(v DateTimeRangeSelectorModel)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *TestRunSearchQueryModel) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetCreatedByIds

`func (o *TestRunSearchQueryModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestRunSearchQueryModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestRunSearchQueryModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestRunSearchQueryModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestRunSearchQueryModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestRunSearchQueryModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *TestRunSearchQueryModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestRunSearchQueryModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestRunSearchQueryModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestRunSearchQueryModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestRunSearchQueryModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestRunSearchQueryModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


