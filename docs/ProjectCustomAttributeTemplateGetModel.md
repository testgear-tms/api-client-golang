# ProjectCustomAttributeTemplateGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the custom attributes template | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the custom attribute template is deleted | [optional] 
**Name** | Pointer to **NullableString** | Name of the custom attribute template | [optional] 
**CustomAttributeModels** | Pointer to [**[]CustomAttributeModel**](CustomAttributeModel.md) | Attributes of the template | [optional] 

## Methods

### NewProjectCustomAttributeTemplateGetModel

`func NewProjectCustomAttributeTemplateGetModel() *ProjectCustomAttributeTemplateGetModel`

NewProjectCustomAttributeTemplateGetModel instantiates a new ProjectCustomAttributeTemplateGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCustomAttributeTemplateGetModelWithDefaults

`func NewProjectCustomAttributeTemplateGetModelWithDefaults() *ProjectCustomAttributeTemplateGetModel`

NewProjectCustomAttributeTemplateGetModelWithDefaults instantiates a new ProjectCustomAttributeTemplateGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCustomAttributeTemplateGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCustomAttributeTemplateGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCustomAttributeTemplateGetModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectCustomAttributeTemplateGetModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ProjectCustomAttributeTemplateGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectCustomAttributeTemplateGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectCustomAttributeTemplateGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProjectCustomAttributeTemplateGetModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *ProjectCustomAttributeTemplateGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCustomAttributeTemplateGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCustomAttributeTemplateGetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectCustomAttributeTemplateGetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectCustomAttributeTemplateGetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectCustomAttributeTemplateGetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomAttributeModels

`func (o *ProjectCustomAttributeTemplateGetModel) GetCustomAttributeModels() []CustomAttributeModel`

GetCustomAttributeModels returns the CustomAttributeModels field if non-nil, zero value otherwise.

### GetCustomAttributeModelsOk

`func (o *ProjectCustomAttributeTemplateGetModel) GetCustomAttributeModelsOk() (*[]CustomAttributeModel, bool)`

GetCustomAttributeModelsOk returns a tuple with the CustomAttributeModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeModels

`func (o *ProjectCustomAttributeTemplateGetModel) SetCustomAttributeModels(v []CustomAttributeModel)`

SetCustomAttributeModels sets CustomAttributeModels field to given value.

### HasCustomAttributeModels

`func (o *ProjectCustomAttributeTemplateGetModel) HasCustomAttributeModels() bool`

HasCustomAttributeModels returns a boolean if a field has been set.

### SetCustomAttributeModelsNil

`func (o *ProjectCustomAttributeTemplateGetModel) SetCustomAttributeModelsNil(b bool)`

 SetCustomAttributeModelsNil sets the value for CustomAttributeModels to be an explicit nil

### UnsetCustomAttributeModels
`func (o *ProjectCustomAttributeTemplateGetModel) UnsetCustomAttributeModels()`

UnsetCustomAttributeModels ensures that no value is present for CustomAttributeModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


