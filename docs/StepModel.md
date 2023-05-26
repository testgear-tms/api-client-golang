# StepModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkItem** | Pointer to [**SharedStepModel**](SharedStepModel.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **NullableString** |  | [optional] 
**Expected** | Pointer to **NullableString** |  | [optional] 
**TestData** | Pointer to **NullableString** |  | [optional] 
**Comments** | Pointer to **NullableString** |  | [optional] 
**WorkItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStepModel

`func NewStepModel() *StepModel`

NewStepModel instantiates a new StepModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepModelWithDefaults

`func NewStepModelWithDefaults() *StepModel`

NewStepModelWithDefaults instantiates a new StepModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkItem

`func (o *StepModel) GetWorkItem() SharedStepModel`

GetWorkItem returns the WorkItem field if non-nil, zero value otherwise.

### GetWorkItemOk

`func (o *StepModel) GetWorkItemOk() (*SharedStepModel, bool)`

GetWorkItemOk returns a tuple with the WorkItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItem

`func (o *StepModel) SetWorkItem(v SharedStepModel)`

SetWorkItem sets WorkItem field to given value.

### HasWorkItem

`func (o *StepModel) HasWorkItem() bool`

HasWorkItem returns a boolean if a field has been set.

### GetId

`func (o *StepModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StepModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StepModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StepModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *StepModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StepModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StepModel) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *StepModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *StepModel) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *StepModel) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetExpected

`func (o *StepModel) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *StepModel) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *StepModel) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *StepModel) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### SetExpectedNil

`func (o *StepModel) SetExpectedNil(b bool)`

 SetExpectedNil sets the value for Expected to be an explicit nil

### UnsetExpected
`func (o *StepModel) UnsetExpected()`

UnsetExpected ensures that no value is present for Expected, not even an explicit nil
### GetTestData

`func (o *StepModel) GetTestData() string`

GetTestData returns the TestData field if non-nil, zero value otherwise.

### GetTestDataOk

`func (o *StepModel) GetTestDataOk() (*string, bool)`

GetTestDataOk returns a tuple with the TestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestData

`func (o *StepModel) SetTestData(v string)`

SetTestData sets TestData field to given value.

### HasTestData

`func (o *StepModel) HasTestData() bool`

HasTestData returns a boolean if a field has been set.

### SetTestDataNil

`func (o *StepModel) SetTestDataNil(b bool)`

 SetTestDataNil sets the value for TestData to be an explicit nil

### UnsetTestData
`func (o *StepModel) UnsetTestData()`

UnsetTestData ensures that no value is present for TestData, not even an explicit nil
### GetComments

`func (o *StepModel) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *StepModel) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *StepModel) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *StepModel) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *StepModel) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *StepModel) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetWorkItemId

`func (o *StepModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *StepModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *StepModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *StepModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *StepModel) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *StepModel) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


