# CustomAttributeSearchQueryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of attribute | [optional] 
**ProjectIds** | Pointer to **[]string** | Unique IDs of projects where attribute is in use | [optional] 
**CustomAttributeIds** | Pointer to **[]string** | Unique IDs of attributes for search restriction | [optional] 
**CustomAttributeTypes** | Pointer to [**[]CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Collection of attribute types | [optional] 
**IsGlobal** | Pointer to **NullableBool** | Indicates whether the attribute is available across all projects | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Indicates whether the attribute is deleted | [optional] 

## Methods

### NewCustomAttributeSearchQueryModel

`func NewCustomAttributeSearchQueryModel() *CustomAttributeSearchQueryModel`

NewCustomAttributeSearchQueryModel instantiates a new CustomAttributeSearchQueryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeSearchQueryModelWithDefaults

`func NewCustomAttributeSearchQueryModelWithDefaults() *CustomAttributeSearchQueryModel`

NewCustomAttributeSearchQueryModelWithDefaults instantiates a new CustomAttributeSearchQueryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomAttributeSearchQueryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAttributeSearchQueryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAttributeSearchQueryModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomAttributeSearchQueryModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomAttributeSearchQueryModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomAttributeSearchQueryModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectIds

`func (o *CustomAttributeSearchQueryModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CustomAttributeSearchQueryModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CustomAttributeSearchQueryModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CustomAttributeSearchQueryModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *CustomAttributeSearchQueryModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *CustomAttributeSearchQueryModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetCustomAttributeIds

`func (o *CustomAttributeSearchQueryModel) GetCustomAttributeIds() []string`

GetCustomAttributeIds returns the CustomAttributeIds field if non-nil, zero value otherwise.

### GetCustomAttributeIdsOk

`func (o *CustomAttributeSearchQueryModel) GetCustomAttributeIdsOk() (*[]string, bool)`

GetCustomAttributeIdsOk returns a tuple with the CustomAttributeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeIds

`func (o *CustomAttributeSearchQueryModel) SetCustomAttributeIds(v []string)`

SetCustomAttributeIds sets CustomAttributeIds field to given value.

### HasCustomAttributeIds

`func (o *CustomAttributeSearchQueryModel) HasCustomAttributeIds() bool`

HasCustomAttributeIds returns a boolean if a field has been set.

### SetCustomAttributeIdsNil

`func (o *CustomAttributeSearchQueryModel) SetCustomAttributeIdsNil(b bool)`

 SetCustomAttributeIdsNil sets the value for CustomAttributeIds to be an explicit nil

### UnsetCustomAttributeIds
`func (o *CustomAttributeSearchQueryModel) UnsetCustomAttributeIds()`

UnsetCustomAttributeIds ensures that no value is present for CustomAttributeIds, not even an explicit nil
### GetCustomAttributeTypes

`func (o *CustomAttributeSearchQueryModel) GetCustomAttributeTypes() []CustomAttributeTypesEnum`

GetCustomAttributeTypes returns the CustomAttributeTypes field if non-nil, zero value otherwise.

### GetCustomAttributeTypesOk

`func (o *CustomAttributeSearchQueryModel) GetCustomAttributeTypesOk() (*[]CustomAttributeTypesEnum, bool)`

GetCustomAttributeTypesOk returns a tuple with the CustomAttributeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeTypes

`func (o *CustomAttributeSearchQueryModel) SetCustomAttributeTypes(v []CustomAttributeTypesEnum)`

SetCustomAttributeTypes sets CustomAttributeTypes field to given value.

### HasCustomAttributeTypes

`func (o *CustomAttributeSearchQueryModel) HasCustomAttributeTypes() bool`

HasCustomAttributeTypes returns a boolean if a field has been set.

### SetCustomAttributeTypesNil

`func (o *CustomAttributeSearchQueryModel) SetCustomAttributeTypesNil(b bool)`

 SetCustomAttributeTypesNil sets the value for CustomAttributeTypes to be an explicit nil

### UnsetCustomAttributeTypes
`func (o *CustomAttributeSearchQueryModel) UnsetCustomAttributeTypes()`

UnsetCustomAttributeTypes ensures that no value is present for CustomAttributeTypes, not even an explicit nil
### GetIsGlobal

`func (o *CustomAttributeSearchQueryModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CustomAttributeSearchQueryModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CustomAttributeSearchQueryModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *CustomAttributeSearchQueryModel) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### SetIsGlobalNil

`func (o *CustomAttributeSearchQueryModel) SetIsGlobalNil(b bool)`

 SetIsGlobalNil sets the value for IsGlobal to be an explicit nil

### UnsetIsGlobal
`func (o *CustomAttributeSearchQueryModel) UnsetIsGlobal()`

UnsetIsGlobal ensures that no value is present for IsGlobal, not even an explicit nil
### GetIsDeleted

`func (o *CustomAttributeSearchQueryModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributeSearchQueryModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributeSearchQueryModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CustomAttributeSearchQueryModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *CustomAttributeSearchQueryModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *CustomAttributeSearchQueryModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


