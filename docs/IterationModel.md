# IterationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ParameterShortModel**](ParameterShortModel.md) |  | [optional] 

## Methods

### NewIterationModel

`func NewIterationModel() *IterationModel`

NewIterationModel instantiates a new IterationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationModelWithDefaults

`func NewIterationModelWithDefaults() *IterationModel`

NewIterationModelWithDefaults instantiates a new IterationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IterationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IterationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IterationModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IterationModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameters

`func (o *IterationModel) GetParameters() []ParameterShortModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IterationModel) GetParametersOk() (*[]ParameterShortModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IterationModel) SetParameters(v []ParameterShortModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IterationModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *IterationModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *IterationModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


