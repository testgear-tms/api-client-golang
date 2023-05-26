# TestSuiteWithChildrenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]TestSuiteWithChildrenModel**](TestSuiteWithChildrenModel.md) |  | [optional] 
**TesterId** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**TestPlanId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewTestSuiteWithChildrenModel

`func NewTestSuiteWithChildrenModel() *TestSuiteWithChildrenModel`

NewTestSuiteWithChildrenModel instantiates a new TestSuiteWithChildrenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteWithChildrenModelWithDefaults

`func NewTestSuiteWithChildrenModelWithDefaults() *TestSuiteWithChildrenModel`

NewTestSuiteWithChildrenModelWithDefaults instantiates a new TestSuiteWithChildrenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *TestSuiteWithChildrenModel) GetChildren() []TestSuiteWithChildrenModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TestSuiteWithChildrenModel) GetChildrenOk() (*[]TestSuiteWithChildrenModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TestSuiteWithChildrenModel) SetChildren(v []TestSuiteWithChildrenModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TestSuiteWithChildrenModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *TestSuiteWithChildrenModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *TestSuiteWithChildrenModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetTesterId

`func (o *TestSuiteWithChildrenModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestSuiteWithChildrenModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestSuiteWithChildrenModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestSuiteWithChildrenModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestSuiteWithChildrenModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestSuiteWithChildrenModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetParentId

`func (o *TestSuiteWithChildrenModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteWithChildrenModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteWithChildrenModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TestSuiteWithChildrenModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TestSuiteWithChildrenModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteWithChildrenModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *TestSuiteWithChildrenModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestSuiteWithChildrenModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestSuiteWithChildrenModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestSuiteWithChildrenModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### GetName

`func (o *TestSuiteWithChildrenModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteWithChildrenModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteWithChildrenModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestSuiteWithChildrenModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestSuiteWithChildrenModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestSuiteWithChildrenModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *TestSuiteWithChildrenModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteWithChildrenModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteWithChildrenModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestSuiteWithChildrenModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TestSuiteWithChildrenModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestSuiteWithChildrenModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestSuiteWithChildrenModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestSuiteWithChildrenModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


