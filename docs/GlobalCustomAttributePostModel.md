# GlobalCustomAttributePostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of attribute | 
**IsEnabled** | Pointer to **NullableBool** | Indicates whether the attribute is available | [optional] 
**IsRequired** | Pointer to **NullableBool** | Indicates whether the attribute value is mandatory to specify | [optional] 
**Options** | Pointer to [**[]CustomAttributeOptionPostModel**](CustomAttributeOptionPostModel.md) | Collection of attribute options  &lt;br /&gt;  Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) |  | 

## Methods

### NewGlobalCustomAttributePostModel

`func NewGlobalCustomAttributePostModel(name string, type_ CustomAttributeTypesEnum, ) *GlobalCustomAttributePostModel`

NewGlobalCustomAttributePostModel instantiates a new GlobalCustomAttributePostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalCustomAttributePostModelWithDefaults

`func NewGlobalCustomAttributePostModelWithDefaults() *GlobalCustomAttributePostModel`

NewGlobalCustomAttributePostModelWithDefaults instantiates a new GlobalCustomAttributePostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GlobalCustomAttributePostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalCustomAttributePostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalCustomAttributePostModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *GlobalCustomAttributePostModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GlobalCustomAttributePostModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GlobalCustomAttributePostModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *GlobalCustomAttributePostModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *GlobalCustomAttributePostModel) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *GlobalCustomAttributePostModel) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIsRequired

`func (o *GlobalCustomAttributePostModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *GlobalCustomAttributePostModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *GlobalCustomAttributePostModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *GlobalCustomAttributePostModel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *GlobalCustomAttributePostModel) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *GlobalCustomAttributePostModel) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetOptions

`func (o *GlobalCustomAttributePostModel) GetOptions() []CustomAttributeOptionPostModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GlobalCustomAttributePostModel) GetOptionsOk() (*[]CustomAttributeOptionPostModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GlobalCustomAttributePostModel) SetOptions(v []CustomAttributeOptionPostModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GlobalCustomAttributePostModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GlobalCustomAttributePostModel) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GlobalCustomAttributePostModel) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetType

`func (o *GlobalCustomAttributePostModel) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalCustomAttributePostModel) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalCustomAttributePostModel) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


