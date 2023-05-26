# ProjectAttributesFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies an attribute name to search for | [optional] 
**IsRequired** | Pointer to **NullableBool** | Specifies an attribute mandatory status to search for | [optional] 
**IsGlobal** | Pointer to **NullableBool** | Specifies an attribute global status to search for | [optional] 
**Types** | Pointer to [**[]CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Specifies an attribute types to search for | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Specifies an attribute enabled status to search for | [optional] 

## Methods

### NewProjectAttributesFilterModel

`func NewProjectAttributesFilterModel() *ProjectAttributesFilterModel`

NewProjectAttributesFilterModel instantiates a new ProjectAttributesFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAttributesFilterModelWithDefaults

`func NewProjectAttributesFilterModelWithDefaults() *ProjectAttributesFilterModel`

NewProjectAttributesFilterModelWithDefaults instantiates a new ProjectAttributesFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectAttributesFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectAttributesFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectAttributesFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectAttributesFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectAttributesFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectAttributesFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsRequired

`func (o *ProjectAttributesFilterModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ProjectAttributesFilterModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ProjectAttributesFilterModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ProjectAttributesFilterModel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *ProjectAttributesFilterModel) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *ProjectAttributesFilterModel) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetIsGlobal

`func (o *ProjectAttributesFilterModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *ProjectAttributesFilterModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *ProjectAttributesFilterModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *ProjectAttributesFilterModel) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### SetIsGlobalNil

`func (o *ProjectAttributesFilterModel) SetIsGlobalNil(b bool)`

 SetIsGlobalNil sets the value for IsGlobal to be an explicit nil

### UnsetIsGlobal
`func (o *ProjectAttributesFilterModel) UnsetIsGlobal()`

UnsetIsGlobal ensures that no value is present for IsGlobal, not even an explicit nil
### GetTypes

`func (o *ProjectAttributesFilterModel) GetTypes() []CustomAttributeTypesEnum`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ProjectAttributesFilterModel) GetTypesOk() (*[]CustomAttributeTypesEnum, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ProjectAttributesFilterModel) SetTypes(v []CustomAttributeTypesEnum)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ProjectAttributesFilterModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *ProjectAttributesFilterModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *ProjectAttributesFilterModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetIsEnabled

`func (o *ProjectAttributesFilterModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ProjectAttributesFilterModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ProjectAttributesFilterModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ProjectAttributesFilterModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *ProjectAttributesFilterModel) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ProjectAttributesFilterModel) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


