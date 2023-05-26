# TestSuiteV2PutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**IsDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewTestSuiteV2PutModel

`func NewTestSuiteV2PutModel(id string, name string, ) *TestSuiteV2PutModel`

NewTestSuiteV2PutModel instantiates a new TestSuiteV2PutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteV2PutModelWithDefaults

`func NewTestSuiteV2PutModelWithDefaults() *TestSuiteV2PutModel`

NewTestSuiteV2PutModelWithDefaults instantiates a new TestSuiteV2PutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestSuiteV2PutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteV2PutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteV2PutModel) SetId(v string)`

SetId sets Id field to given value.


### GetParentId

`func (o *TestSuiteV2PutModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteV2PutModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteV2PutModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TestSuiteV2PutModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TestSuiteV2PutModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteV2PutModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetName

`func (o *TestSuiteV2PutModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteV2PutModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteV2PutModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsDeleted

`func (o *TestSuiteV2PutModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestSuiteV2PutModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestSuiteV2PutModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestSuiteV2PutModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


