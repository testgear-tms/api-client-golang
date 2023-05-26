# SharedStepModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | Pointer to **string** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Steps** | Pointer to [**[]StepModel**](StepModel.md) |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewSharedStepModel

`func NewSharedStepModel() *SharedStepModel`

NewSharedStepModel instantiates a new SharedStepModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepModelWithDefaults

`func NewSharedStepModelWithDefaults() *SharedStepModel`

NewSharedStepModelWithDefaults instantiates a new SharedStepModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *SharedStepModel) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SharedStepModel) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SharedStepModel) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *SharedStepModel) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetGlobalId

`func (o *SharedStepModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SharedStepModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SharedStepModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *SharedStepModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetName

`func (o *SharedStepModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedStepModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedStepModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedStepModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SharedStepModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SharedStepModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSteps

`func (o *SharedStepModel) GetSteps() []StepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SharedStepModel) GetStepsOk() (*[]StepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SharedStepModel) SetSteps(v []StepModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SharedStepModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *SharedStepModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *SharedStepModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetIsDeleted

`func (o *SharedStepModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SharedStepModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SharedStepModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SharedStepModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


