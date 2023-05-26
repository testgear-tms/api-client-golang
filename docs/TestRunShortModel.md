# TestRunShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateName** | [**TestRunState**](TestRunState.md) |  | 
**ProjectId** | Pointer to **string** |  | [optional] 
**TestPlanId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewTestRunShortModel

`func NewTestRunShortModel(stateName TestRunState, ) *TestRunShortModel`

NewTestRunShortModel instantiates a new TestRunShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunShortModelWithDefaults

`func NewTestRunShortModelWithDefaults() *TestRunShortModel`

NewTestRunShortModelWithDefaults instantiates a new TestRunShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateName

`func (o *TestRunShortModel) GetStateName() TestRunState`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *TestRunShortModel) GetStateNameOk() (*TestRunState, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *TestRunShortModel) SetStateName(v TestRunState)`

SetStateName sets StateName field to given value.


### GetProjectId

`func (o *TestRunShortModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunShortModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunShortModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TestRunShortModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTestPlanId

`func (o *TestRunShortModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestRunShortModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestRunShortModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestRunShortModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *TestRunShortModel) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestRunShortModel) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetName

`func (o *TestRunShortModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunShortModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunShortModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunShortModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunShortModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunShortModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TestRunShortModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunShortModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunShortModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunShortModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunShortModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunShortModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *TestRunShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunShortModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestRunShortModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TestRunShortModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunShortModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunShortModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestRunShortModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


