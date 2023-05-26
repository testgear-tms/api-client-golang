# SharedStepReferenceSectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**HasThisSharedStepAsPrecondition** | Pointer to **bool** |  | [optional] 
**HasThisSharedStepAsPostcondition** | Pointer to **bool** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewSharedStepReferenceSectionModel

`func NewSharedStepReferenceSectionModel() *SharedStepReferenceSectionModel`

NewSharedStepReferenceSectionModel instantiates a new SharedStepReferenceSectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepReferenceSectionModelWithDefaults

`func NewSharedStepReferenceSectionModelWithDefaults() *SharedStepReferenceSectionModel`

NewSharedStepReferenceSectionModelWithDefaults instantiates a new SharedStepReferenceSectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedStepReferenceSectionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedStepReferenceSectionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedStepReferenceSectionModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SharedStepReferenceSectionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SharedStepReferenceSectionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedStepReferenceSectionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedStepReferenceSectionModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedStepReferenceSectionModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SharedStepReferenceSectionModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SharedStepReferenceSectionModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHasThisSharedStepAsPrecondition

`func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPrecondition() bool`

GetHasThisSharedStepAsPrecondition returns the HasThisSharedStepAsPrecondition field if non-nil, zero value otherwise.

### GetHasThisSharedStepAsPreconditionOk

`func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPreconditionOk() (*bool, bool)`

GetHasThisSharedStepAsPreconditionOk returns a tuple with the HasThisSharedStepAsPrecondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThisSharedStepAsPrecondition

`func (o *SharedStepReferenceSectionModel) SetHasThisSharedStepAsPrecondition(v bool)`

SetHasThisSharedStepAsPrecondition sets HasThisSharedStepAsPrecondition field to given value.

### HasHasThisSharedStepAsPrecondition

`func (o *SharedStepReferenceSectionModel) HasHasThisSharedStepAsPrecondition() bool`

HasHasThisSharedStepAsPrecondition returns a boolean if a field has been set.

### GetHasThisSharedStepAsPostcondition

`func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPostcondition() bool`

GetHasThisSharedStepAsPostcondition returns the HasThisSharedStepAsPostcondition field if non-nil, zero value otherwise.

### GetHasThisSharedStepAsPostconditionOk

`func (o *SharedStepReferenceSectionModel) GetHasThisSharedStepAsPostconditionOk() (*bool, bool)`

GetHasThisSharedStepAsPostconditionOk returns a tuple with the HasThisSharedStepAsPostcondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasThisSharedStepAsPostcondition

`func (o *SharedStepReferenceSectionModel) SetHasThisSharedStepAsPostcondition(v bool)`

SetHasThisSharedStepAsPostcondition sets HasThisSharedStepAsPostcondition field to given value.

### HasHasThisSharedStepAsPostcondition

`func (o *SharedStepReferenceSectionModel) HasHasThisSharedStepAsPostcondition() bool`

HasHasThisSharedStepAsPostcondition returns a boolean if a field has been set.

### GetCreatedById

`func (o *SharedStepReferenceSectionModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SharedStepReferenceSectionModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SharedStepReferenceSectionModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *SharedStepReferenceSectionModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *SharedStepReferenceSectionModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *SharedStepReferenceSectionModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *SharedStepReferenceSectionModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *SharedStepReferenceSectionModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *SharedStepReferenceSectionModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *SharedStepReferenceSectionModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedDate

`func (o *SharedStepReferenceSectionModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *SharedStepReferenceSectionModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *SharedStepReferenceSectionModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *SharedStepReferenceSectionModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *SharedStepReferenceSectionModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *SharedStepReferenceSectionModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *SharedStepReferenceSectionModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SharedStepReferenceSectionModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SharedStepReferenceSectionModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SharedStepReferenceSectionModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *SharedStepReferenceSectionModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *SharedStepReferenceSectionModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetIsDeleted

`func (o *SharedStepReferenceSectionModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SharedStepReferenceSectionModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SharedStepReferenceSectionModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SharedStepReferenceSectionModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


