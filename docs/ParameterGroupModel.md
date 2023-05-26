# ParameterGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Values** | Pointer to **map[string]string** |  | [optional] 
**ParameterKeyId** | Pointer to **string** |  | [optional] 

## Methods

### NewParameterGroupModel

`func NewParameterGroupModel() *ParameterGroupModel`

NewParameterGroupModel instantiates a new ParameterGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterGroupModelWithDefaults

`func NewParameterGroupModelWithDefaults() *ParameterGroupModel`

NewParameterGroupModelWithDefaults instantiates a new ParameterGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParameterGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParameterGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ParameterGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ParameterGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValues

`func (o *ParameterGroupModel) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ParameterGroupModel) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ParameterGroupModel) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ParameterGroupModel) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ParameterGroupModel) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ParameterGroupModel) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetParameterKeyId

`func (o *ParameterGroupModel) GetParameterKeyId() string`

GetParameterKeyId returns the ParameterKeyId field if non-nil, zero value otherwise.

### GetParameterKeyIdOk

`func (o *ParameterGroupModel) GetParameterKeyIdOk() (*string, bool)`

GetParameterKeyIdOk returns a tuple with the ParameterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterKeyId

`func (o *ParameterGroupModel) SetParameterKeyId(v string)`

SetParameterKeyId sets ParameterKeyId field to given value.

### HasParameterKeyId

`func (o *ParameterGroupModel) HasParameterKeyId() bool`

HasParameterKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


