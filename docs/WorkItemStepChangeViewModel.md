# WorkItemStepChangeViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** |  | [optional] 
**Expected** | Pointer to **NullableString** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**TestData** | Pointer to **NullableString** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**WorkItemId** | Pointer to **NullableString** |  | [optional] 
**WorkItem** | Pointer to [**SharedStepChangeViewModel**](SharedStepChangeViewModel.md) |  | [optional] 

## Methods

### NewWorkItemStepChangeViewModel

`func NewWorkItemStepChangeViewModel() *WorkItemStepChangeViewModel`

NewWorkItemStepChangeViewModel instantiates a new WorkItemStepChangeViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemStepChangeViewModelWithDefaults

`func NewWorkItemStepChangeViewModelWithDefaults() *WorkItemStepChangeViewModel`

NewWorkItemStepChangeViewModelWithDefaults instantiates a new WorkItemStepChangeViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkItemStepChangeViewModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkItemStepChangeViewModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkItemStepChangeViewModel) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkItemStepChangeViewModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *WorkItemStepChangeViewModel) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *WorkItemStepChangeViewModel) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetExpected

`func (o *WorkItemStepChangeViewModel) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *WorkItemStepChangeViewModel) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *WorkItemStepChangeViewModel) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *WorkItemStepChangeViewModel) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### SetExpectedNil

`func (o *WorkItemStepChangeViewModel) SetExpectedNil(b bool)`

 SetExpectedNil sets the value for Expected to be an explicit nil

### UnsetExpected
`func (o *WorkItemStepChangeViewModel) UnsetExpected()`

UnsetExpected ensures that no value is present for Expected, not even an explicit nil
### GetComments

`func (o *WorkItemStepChangeViewModel) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WorkItemStepChangeViewModel) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WorkItemStepChangeViewModel) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WorkItemStepChangeViewModel) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *WorkItemStepChangeViewModel) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *WorkItemStepChangeViewModel) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetTestData

`func (o *WorkItemStepChangeViewModel) GetTestData() string`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *WorkItemStepChangeViewModel) GetTestDataOk() (*string, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *WorkItemStepChangeViewModel) SetTestData(v string)`

SetTestData sets TestData field to given value.

### HasTestData

`func (o *WorkItemStepChangeViewModel) HasTestData() bool`

HasTestData returns a boolean if a field has been set.

### SetTestDataNil

`func (o *WorkItemStepChangeViewModel) SetTestDataNil(b bool)`

 SetTestDataNil sets the value for TestData to be an explicit nil

### UnsetTestData
`func (o *WorkItemStepChangeViewModel) UnsetTestData()`

UnsetTestData ensures that no value is present for TestData, not even an explicit nil
### GetIndex

`func (o *WorkItemStepChangeViewModel) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WorkItemStepChangeViewModel) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WorkItemStepChangeViewModel) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *WorkItemStepChangeViewModel) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetWorkItemId

`func (o *WorkItemStepChangeViewModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *WorkItemStepChangeViewModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *WorkItemStepChangeViewModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *WorkItemStepChangeViewModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *WorkItemStepChangeViewModel) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *WorkItemStepChangeViewModel) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
### GetWorkItem

`func (o *WorkItemStepChangeViewModel) GetWorkItem() SharedStepChangeViewModel`

GetWorkItem returns the WorkItem field if non-nil, zero value otherwise.

### GetWorkItemOk

`func (o *WorkItemStepChangeViewModel) GetWorkItemOk() (*SharedStepChangeViewModel, bool)`

GetWorkItemOk returns a tuple with the WorkItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItem

`func (o *WorkItemStepChangeViewModel) SetWorkItem(v SharedStepChangeViewModel)`

SetWorkItem sets WorkItem field to given value.

### HasWorkItem

`func (o *WorkItemStepChangeViewModel) HasWorkItem() bool`

HasWorkItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


