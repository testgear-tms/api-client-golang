# SectionPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ProjectId** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**PreconditionSteps** | Pointer to [**[]StepPutModel**](StepPutModel.md) |  | [optional] 
**PostconditionSteps** | Pointer to [**[]StepPutModel**](StepPutModel.md) |  | [optional] 

## Methods

### NewSectionPostModel

`func NewSectionPostModel(name string, projectId string, ) *SectionPostModel`

NewSectionPostModel instantiates a new SectionPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionPostModelWithDefaults

`func NewSectionPostModelWithDefaults() *SectionPostModel`

NewSectionPostModelWithDefaults instantiates a new SectionPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SectionPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectionPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectionPostModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *SectionPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SectionPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SectionPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetParentId

`func (o *SectionPostModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SectionPostModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SectionPostModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SectionPostModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *SectionPostModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *SectionPostModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPreconditionSteps

`func (o *SectionPostModel) GetPreconditionSteps() []StepPutModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *SectionPostModel) GetPreconditionStepsOk() (*[]StepPutModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *SectionPostModel) SetPreconditionSteps(v []StepPutModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.

### HasPreconditionSteps

`func (o *SectionPostModel) HasPreconditionSteps() bool`

HasPreconditionSteps returns a boolean if a field has been set.

### SetPreconditionStepsNil

`func (o *SectionPostModel) SetPreconditionStepsNil(b bool)`

 SetPreconditionStepsNil sets the value for PreconditionSteps to be an explicit nil

### UnsetPreconditionSteps
`func (o *SectionPostModel) UnsetPreconditionSteps()`

UnsetPreconditionSteps ensures that no value is present for PreconditionSteps, not even an explicit nil
### GetPostconditionSteps

`func (o *SectionPostModel) GetPostconditionSteps() []StepPutModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *SectionPostModel) GetPostconditionStepsOk() (*[]StepPutModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *SectionPostModel) SetPostconditionSteps(v []StepPutModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.

### HasPostconditionSteps

`func (o *SectionPostModel) HasPostconditionSteps() bool`

HasPostconditionSteps returns a boolean if a field has been set.

### SetPostconditionStepsNil

`func (o *SectionPostModel) SetPostconditionStepsNil(b bool)`

 SetPostconditionStepsNil sets the value for PostconditionSteps to be an explicit nil

### UnsetPostconditionSteps
`func (o *SectionPostModel) UnsetPostconditionSteps()`

UnsetPostconditionSteps ensures that no value is present for PostconditionSteps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


