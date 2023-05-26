# TestSuiteV2TreeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]TestSuiteV2TreeModel**](TestSuiteV2TreeModel.md) | nested enumeration of children is allowed | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**TestPlanId** | **string** |  | 
**Name** | **string** |  | 
**Type** | Pointer to [**TestSuiteType**](TestSuiteType.md) |  | [optional] 
**SaveStructure** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewTestSuiteV2TreeModel

`func NewTestSuiteV2TreeModel(testPlanId string, name string, ) *TestSuiteV2TreeModel`

NewTestSuiteV2TreeModel instantiates a new TestSuiteV2TreeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteV2TreeModelWithDefaults

`func NewTestSuiteV2TreeModelWithDefaults() *TestSuiteV2TreeModel`

NewTestSuiteV2TreeModelWithDefaults instantiates a new TestSuiteV2TreeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *TestSuiteV2TreeModel) GetChildren() []TestSuiteV2TreeModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TestSuiteV2TreeModel) GetChildrenOk() (*[]TestSuiteV2TreeModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TestSuiteV2TreeModel) SetChildren(v []TestSuiteV2TreeModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TestSuiteV2TreeModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *TestSuiteV2TreeModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *TestSuiteV2TreeModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetId

`func (o *TestSuiteV2TreeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteV2TreeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteV2TreeModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestSuiteV2TreeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *TestSuiteV2TreeModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteV2TreeModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteV2TreeModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TestSuiteV2TreeModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TestSuiteV2TreeModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteV2TreeModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *TestSuiteV2TreeModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestSuiteV2TreeModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestSuiteV2TreeModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *TestSuiteV2TreeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteV2TreeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteV2TreeModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestSuiteV2TreeModel) GetType() TestSuiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestSuiteV2TreeModel) GetTypeOk() (*TestSuiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestSuiteV2TreeModel) SetType(v TestSuiteType)`

SetType sets Type field to given value.

### HasType

`func (o *TestSuiteV2TreeModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSaveStructure

`func (o *TestSuiteV2TreeModel) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *TestSuiteV2TreeModel) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *TestSuiteV2TreeModel) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.

### HasSaveStructure

`func (o *TestSuiteV2TreeModel) HasSaveStructure() bool`

HasSaveStructure returns a boolean if a field has been set.

### SetSaveStructureNil

`func (o *TestSuiteV2TreeModel) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *TestSuiteV2TreeModel) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


