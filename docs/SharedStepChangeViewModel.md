# SharedStepChangeViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Steps** | Pointer to [**[]WorkItemStepChangeViewModel**](WorkItemStepChangeViewModel.md) |  | [optional] 

## Methods

### NewSharedStepChangeViewModel

`func NewSharedStepChangeViewModel() *SharedStepChangeViewModel`

NewSharedStepChangeViewModel instantiates a new SharedStepChangeViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepChangeViewModelWithDefaults

`func NewSharedStepChangeViewModelWithDefaults() *SharedStepChangeViewModel`

NewSharedStepChangeViewModelWithDefaults instantiates a new SharedStepChangeViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedStepChangeViewModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedStepChangeViewModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedStepChangeViewModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SharedStepChangeViewModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGlobalId

`func (o *SharedStepChangeViewModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SharedStepChangeViewModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SharedStepChangeViewModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *SharedStepChangeViewModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetName

`func (o *SharedStepChangeViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedStepChangeViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedStepChangeViewModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedStepChangeViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SharedStepChangeViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SharedStepChangeViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSteps

`func (o *SharedStepChangeViewModel) GetSteps() []WorkItemStepChangeViewModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SharedStepChangeViewModel) GetStepsOk() (*[]WorkItemStepChangeViewModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SharedStepChangeViewModel) SetSteps(v []WorkItemStepChangeViewModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SharedStepChangeViewModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *SharedStepChangeViewModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *SharedStepChangeViewModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


