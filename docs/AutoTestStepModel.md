# AutoTestStepModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Step name. | 
**Description** | Pointer to **NullableString** | Detailed step description. It appears when the step is unfolded. | [optional] 
**Steps** | Pointer to [**[]AutoTestStepModel**](AutoTestStepModel.md) | Includes a nested step inside another step. The maximum nesting level is 15. | [optional] 

## Methods

### NewAutoTestStepModel

`func NewAutoTestStepModel(title string, ) *AutoTestStepModel`

NewAutoTestStepModel instantiates a new AutoTestStepModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestStepModelWithDefaults

`func NewAutoTestStepModelWithDefaults() *AutoTestStepModel`

NewAutoTestStepModelWithDefaults instantiates a new AutoTestStepModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AutoTestStepModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestStepModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestStepModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AutoTestStepModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestStepModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestStepModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestStepModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestStepModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestStepModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSteps

`func (o *AutoTestStepModel) GetSteps() []AutoTestStepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestStepModel) GetStepsOk() (*[]AutoTestStepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestStepModel) SetSteps(v []AutoTestStepModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestStepModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestStepModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestStepModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


