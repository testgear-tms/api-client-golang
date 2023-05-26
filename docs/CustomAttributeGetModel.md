# CustomAttributeGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the attribute | [optional] 
**Options** | Pointer to [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of the attribute options | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) |  | 
**IsDeleted** | Pointer to **bool** | Indicates if the attribute is deleted | [optional] 
**Name** | Pointer to **NullableString** | Name of the attribute | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates if the attribute is enabled | [optional] 
**IsRequired** | Pointer to **bool** | Indicates if the attribute is mandatory to specify | [optional] 
**IsGlobal** | Pointer to **bool** | Indicates if the attribute is available across all projects | [optional] 

## Methods

### NewCustomAttributeGetModel

`func NewCustomAttributeGetModel(type_ CustomAttributeTypesEnum, ) *CustomAttributeGetModel`

NewCustomAttributeGetModel instantiates a new CustomAttributeGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeGetModelWithDefaults

`func NewCustomAttributeGetModelWithDefaults() *CustomAttributeGetModel`

NewCustomAttributeGetModelWithDefaults instantiates a new CustomAttributeGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAttributeGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributeGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributeGetModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomAttributeGetModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptions

`func (o *CustomAttributeGetModel) GetOptions() []CustomAttributeOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomAttributeGetModel) GetOptionsOk() (*[]CustomAttributeOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomAttributeGetModel) SetOptions(v []CustomAttributeOptionModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CustomAttributeGetModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CustomAttributeGetModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CustomAttributeGetModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetType

`func (o *CustomAttributeGetModel) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAttributeGetModel) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAttributeGetModel) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetIsDeleted

`func (o *CustomAttributeGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributeGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributeGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CustomAttributeGetModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *CustomAttributeGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAttributeGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAttributeGetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomAttributeGetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomAttributeGetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomAttributeGetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsEnabled

`func (o *CustomAttributeGetModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomAttributeGetModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomAttributeGetModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CustomAttributeGetModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsRequired

`func (o *CustomAttributeGetModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomAttributeGetModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomAttributeGetModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *CustomAttributeGetModel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsGlobal

`func (o *CustomAttributeGetModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CustomAttributeGetModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CustomAttributeGetModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *CustomAttributeGetModel) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


