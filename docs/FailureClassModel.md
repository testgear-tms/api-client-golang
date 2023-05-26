# FailureClassModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**FailureCategory** | [**FailureCategoryModel**](FailureCategoryModel.md) |  | 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**FailureClassRegexes** | Pointer to [**[]FailureClassRegexModel**](FailureClassRegexModel.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewFailureClassModel

`func NewFailureClassModel(failureCategory FailureCategoryModel, ) *FailureClassModel`

NewFailureClassModel instantiates a new FailureClassModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureClassModelWithDefaults

`func NewFailureClassModelWithDefaults() *FailureClassModel`

NewFailureClassModelWithDefaults instantiates a new FailureClassModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FailureClassModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FailureClassModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FailureClassModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FailureClassModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FailureClassModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FailureClassModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFailureCategory

`func (o *FailureClassModel) GetFailureCategory() FailureCategoryModel`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *FailureClassModel) GetFailureCategoryOk() (*FailureCategoryModel, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *FailureClassModel) SetFailureCategory(v FailureCategoryModel)`

SetFailureCategory sets FailureCategory field to given value.


### GetCreatedDate

`func (o *FailureClassModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *FailureClassModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *FailureClassModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *FailureClassModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *FailureClassModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *FailureClassModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *FailureClassModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *FailureClassModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *FailureClassModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *FailureClassModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *FailureClassModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *FailureClassModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *FailureClassModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *FailureClassModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *FailureClassModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *FailureClassModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *FailureClassModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *FailureClassModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *FailureClassModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *FailureClassModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetFailureClassRegexes

`func (o *FailureClassModel) GetFailureClassRegexes() []FailureClassRegexModel`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *FailureClassModel) GetFailureClassRegexesOk() (*[]FailureClassRegexModel, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *FailureClassModel) SetFailureClassRegexes(v []FailureClassRegexModel)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.

### HasFailureClassRegexes

`func (o *FailureClassModel) HasFailureClassRegexes() bool`

HasFailureClassRegexes returns a boolean if a field has been set.

### SetFailureClassRegexesNil

`func (o *FailureClassModel) SetFailureClassRegexesNil(b bool)`

 SetFailureClassRegexesNil sets the value for FailureClassRegexes to be an explicit nil

### UnsetFailureClassRegexes
`func (o *FailureClassModel) UnsetFailureClassRegexes()`

UnsetFailureClassRegexes ensures that no value is present for FailureClassRegexes, not even an explicit nil
### GetId

`func (o *FailureClassModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailureClassModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailureClassModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FailureClassModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *FailureClassModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FailureClassModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FailureClassModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *FailureClassModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


