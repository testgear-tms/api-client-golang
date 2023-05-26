# SectionWithStepsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreconditionSteps** | Pointer to [**[]StepModel**](StepModel.md) |  | [optional] 
**PostconditionSteps** | Pointer to [**[]StepModel**](StepModel.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewSectionWithStepsModel

`func NewSectionWithStepsModel(name string, ) *SectionWithStepsModel`

NewSectionWithStepsModel instantiates a new SectionWithStepsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithStepsModelWithDefaults

`func NewSectionWithStepsModelWithDefaults() *SectionWithStepsModel`

NewSectionWithStepsModelWithDefaults instantiates a new SectionWithStepsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreconditionSteps

`func (o *SectionWithStepsModel) GetPreconditionSteps() []StepModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *SectionWithStepsModel) GetPreconditionStepsOk() (*[]StepModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *SectionWithStepsModel) SetPreconditionSteps(v []StepModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.

### HasPreconditionSteps

`func (o *SectionWithStepsModel) HasPreconditionSteps() bool`

HasPreconditionSteps returns a boolean if a field has been set.

### SetPreconditionStepsNil

`func (o *SectionWithStepsModel) SetPreconditionStepsNil(b bool)`

 SetPreconditionStepsNil sets the value for PreconditionSteps to be an explicit nil

### UnsetPreconditionSteps
`func (o *SectionWithStepsModel) UnsetPreconditionSteps()`

UnsetPreconditionSteps ensures that no value is present for PreconditionSteps, not even an explicit nil
### GetPostconditionSteps

`func (o *SectionWithStepsModel) GetPostconditionSteps() []StepModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *SectionWithStepsModel) GetPostconditionStepsOk() (*[]StepModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *SectionWithStepsModel) SetPostconditionSteps(v []StepModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.

### HasPostconditionSteps

`func (o *SectionWithStepsModel) HasPostconditionSteps() bool`

HasPostconditionSteps returns a boolean if a field has been set.

### SetPostconditionStepsNil

`func (o *SectionWithStepsModel) SetPostconditionStepsNil(b bool)`

 SetPostconditionStepsNil sets the value for PostconditionSteps to be an explicit nil

### UnsetPostconditionSteps
`func (o *SectionWithStepsModel) UnsetPostconditionSteps()`

UnsetPostconditionSteps ensures that no value is present for PostconditionSteps, not even an explicit nil
### GetProjectId

`func (o *SectionWithStepsModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SectionWithStepsModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SectionWithStepsModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SectionWithStepsModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *SectionWithStepsModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *SectionWithStepsModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetParentId

`func (o *SectionWithStepsModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SectionWithStepsModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SectionWithStepsModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SectionWithStepsModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *SectionWithStepsModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *SectionWithStepsModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetIsDeleted

`func (o *SectionWithStepsModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SectionWithStepsModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SectionWithStepsModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SectionWithStepsModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetId

`func (o *SectionWithStepsModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectionWithStepsModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectionWithStepsModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SectionWithStepsModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *SectionWithStepsModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *SectionWithStepsModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *SectionWithStepsModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *SectionWithStepsModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *SectionWithStepsModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SectionWithStepsModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SectionWithStepsModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SectionWithStepsModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *SectionWithStepsModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *SectionWithStepsModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *SectionWithStepsModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SectionWithStepsModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SectionWithStepsModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SectionWithStepsModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *SectionWithStepsModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *SectionWithStepsModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *SectionWithStepsModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *SectionWithStepsModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *SectionWithStepsModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *SectionWithStepsModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetName

`func (o *SectionWithStepsModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectionWithStepsModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectionWithStepsModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


