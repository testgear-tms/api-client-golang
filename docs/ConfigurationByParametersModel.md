# ConfigurationByParametersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** | This property is used to link configuration with project | [optional] 
**ParameterIds** | **[]string** |  | 

## Methods

### NewConfigurationByParametersModel

`func NewConfigurationByParametersModel(parameterIds []string, ) *ConfigurationByParametersModel`

NewConfigurationByParametersModel instantiates a new ConfigurationByParametersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationByParametersModelWithDefaults

`func NewConfigurationByParametersModelWithDefaults() *ConfigurationByParametersModel`

NewConfigurationByParametersModelWithDefaults instantiates a new ConfigurationByParametersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ConfigurationByParametersModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ConfigurationByParametersModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ConfigurationByParametersModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ConfigurationByParametersModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParameterIds

`func (o *ConfigurationByParametersModel) GetParameterIds() []string`

GetParameterIds returns the ParameterIds field if non-nil, zero value otherwise.

### GetParameterIdsOk

`func (o *ConfigurationByParametersModel) GetParameterIdsOk() (*[]string, bool)`

GetParameterIdsOk returns a tuple with the ParameterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterIds

`func (o *ConfigurationByParametersModel) SetParameterIds(v []string)`

SetParameterIds sets ParameterIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


