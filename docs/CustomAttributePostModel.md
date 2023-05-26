# CustomAttributePostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**[]CustomAttributeOptionPostModel**](CustomAttributeOptionPostModel.md) | Collection of attribute options  &lt;br /&gt;  Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) |  | 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | Pointer to **bool** | Indicates if the attribute is enabled | [optional] 
**IsRequired** | Pointer to **bool** | Indicates if the attribute value is mandatory to specify | [optional] 
**IsGlobal** | Pointer to **bool** | Indicates if the attribute is available across all projects | [optional] 

## Methods

### NewCustomAttributePostModel

`func NewCustomAttributePostModel(type_ CustomAttributeTypesEnum, name string, ) *CustomAttributePostModel`

NewCustomAttributePostModel instantiates a new CustomAttributePostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributePostModelWithDefaults

`func NewCustomAttributePostModelWithDefaults() *CustomAttributePostModel`

NewCustomAttributePostModelWithDefaults instantiates a new CustomAttributePostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *CustomAttributePostModel) GetOptions() []CustomAttributeOptionPostModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomAttributePostModel) GetOptionsOk() (*[]CustomAttributeOptionPostModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomAttributePostModel) SetOptions(v []CustomAttributeOptionPostModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CustomAttributePostModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CustomAttributePostModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CustomAttributePostModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetType

`func (o *CustomAttributePostModel) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAttributePostModel) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAttributePostModel) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetName

`func (o *CustomAttributePostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAttributePostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAttributePostModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *CustomAttributePostModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomAttributePostModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomAttributePostModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CustomAttributePostModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsRequired

`func (o *CustomAttributePostModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomAttributePostModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomAttributePostModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *CustomAttributePostModel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsGlobal

`func (o *CustomAttributePostModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CustomAttributePostModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CustomAttributePostModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *CustomAttributePostModel) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


