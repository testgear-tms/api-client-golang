# PublicTestPointModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | Pointer to **string** |  | [optional] 
**ConfigurationGlobalId** | Pointer to **int64** |  | [optional] 
**AutoTestIds** | Pointer to **[]string** |  | [optional] 
**IterationId** | Pointer to **string** |  | [optional] 
**ParameterModels** | Pointer to [**[]ParameterShortModel**](ParameterShortModel.md) |  | [optional] 

## Methods

### NewPublicTestPointModel

`func NewPublicTestPointModel() *PublicTestPointModel`

NewPublicTestPointModel instantiates a new PublicTestPointModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTestPointModelWithDefaults

`func NewPublicTestPointModelWithDefaults() *PublicTestPointModel`

NewPublicTestPointModelWithDefaults instantiates a new PublicTestPointModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationId

`func (o *PublicTestPointModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *PublicTestPointModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *PublicTestPointModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *PublicTestPointModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetConfigurationGlobalId

`func (o *PublicTestPointModel) GetConfigurationGlobalId() int64`

GetConfigurationGlobalId returns the ConfigurationGlobalId field if non-nil, zero value otherwise.

### GetConfigurationGlobalIdOk

`func (o *PublicTestPointModel) GetConfigurationGlobalIdOk() (*int64, bool)`

GetConfigurationGlobalIdOk returns a tuple with the ConfigurationGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationGlobalId

`func (o *PublicTestPointModel) SetConfigurationGlobalId(v int64)`

SetConfigurationGlobalId sets ConfigurationGlobalId field to given value.

### HasConfigurationGlobalId

`func (o *PublicTestPointModel) HasConfigurationGlobalId() bool`

HasConfigurationGlobalId returns a boolean if a field has been set.

### GetAutoTestIds

`func (o *PublicTestPointModel) GetAutoTestIds() []string`

GetAutoTestIds returns the AutoTestIds field if non-nil, zero value otherwise.

### GetAutoTestIdsOk

`func (o *PublicTestPointModel) GetAutoTestIdsOk() (*[]string, bool)`

GetAutoTestIdsOk returns a tuple with the AutoTestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestIds

`func (o *PublicTestPointModel) SetAutoTestIds(v []string)`

SetAutoTestIds sets AutoTestIds field to given value.

### HasAutoTestIds

`func (o *PublicTestPointModel) HasAutoTestIds() bool`

HasAutoTestIds returns a boolean if a field has been set.

### SetAutoTestIdsNil

`func (o *PublicTestPointModel) SetAutoTestIdsNil(b bool)`

 SetAutoTestIdsNil sets the value for AutoTestIds to be an explicit nil

### UnsetAutoTestIds
`func (o *PublicTestPointModel) UnsetAutoTestIds()`

UnsetAutoTestIds ensures that no value is present for AutoTestIds, not even an explicit nil
### GetIterationId

`func (o *PublicTestPointModel) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *PublicTestPointModel) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *PublicTestPointModel) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *PublicTestPointModel) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### GetParameterModels

`func (o *PublicTestPointModel) GetParameterModels() []ParameterShortModel`

GetParameterModels returns the ParameterModels field if non-nil, zero value otherwise.

### GetParameterModelsOk

`func (o *PublicTestPointModel) GetParameterModelsOk() (*[]ParameterShortModel, bool)`

GetParameterModelsOk returns a tuple with the ParameterModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterModels

`func (o *PublicTestPointModel) SetParameterModels(v []ParameterShortModel)`

SetParameterModels sets ParameterModels field to given value.

### HasParameterModels

`func (o *PublicTestPointModel) HasParameterModels() bool`

HasParameterModels returns a boolean if a field has been set.

### SetParameterModelsNil

`func (o *PublicTestPointModel) SetParameterModelsNil(b bool)`

 SetParameterModelsNil sets the value for ParameterModels to be an explicit nil

### UnsetParameterModels
`func (o *PublicTestPointModel) UnsetParameterModels()`

UnsetParameterModels ensures that no value is present for ParameterModels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


