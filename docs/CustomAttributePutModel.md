# CustomAttributePutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the attribute | [optional] 
**Options** | Pointer to [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of the attribute options  &lt;br /&gt;  Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) |  | 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | Pointer to **bool** | Indicates if the attribute is enabled | [optional] 
**IsRequired** | Pointer to **bool** | Indicates if the attribute value is mandatory to specify | [optional] 
**IsGlobal** | Pointer to **bool** | Indicates if the attribute is available across all projects | [optional] 

## Methods

### NewCustomAttributePutModel

`func NewCustomAttributePutModel(type_ CustomAttributeTypesEnum, name string, ) *CustomAttributePutModel`

NewCustomAttributePutModel instantiates a new CustomAttributePutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributePutModelWithDefaults

`func NewCustomAttributePutModelWithDefaults() *CustomAttributePutModel`

NewCustomAttributePutModelWithDefaults instantiates a new CustomAttributePutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAttributePutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributePutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributePutModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomAttributePutModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptions

`func (o *CustomAttributePutModel) GetOptions() []CustomAttributeOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomAttributePutModel) GetOptionsOk() (*[]CustomAttributeOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomAttributePutModel) SetOptions(v []CustomAttributeOptionModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CustomAttributePutModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CustomAttributePutModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CustomAttributePutModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetType

`func (o *CustomAttributePutModel) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAttributePutModel) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAttributePutModel) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetIsDeleted

`func (o *CustomAttributePutModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributePutModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributePutModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CustomAttributePutModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetName

`func (o *CustomAttributePutModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAttributePutModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAttributePutModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *CustomAttributePutModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomAttributePutModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomAttributePutModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CustomAttributePutModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsRequired

`func (o *CustomAttributePutModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomAttributePutModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomAttributePutModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *CustomAttributePutModel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsGlobal

`func (o *CustomAttributePutModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CustomAttributePutModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CustomAttributePutModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *CustomAttributePutModel) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


