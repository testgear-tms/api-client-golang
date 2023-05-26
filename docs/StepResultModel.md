# StepResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepId** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **NullableString** |  | [optional] 
**SharedStepVersionId** | Pointer to **NullableString** |  | [optional] 
**SharedStepResults** | Pointer to [**[]SharedStepResultModel**](SharedStepResultModel.md) |  | [optional] 
**Comment** | Pointer to [**StepCommentModel**](StepCommentModel.md) |  | [optional] 

## Methods

### NewStepResultModel

`func NewStepResultModel() *StepResultModel`

NewStepResultModel instantiates a new StepResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepResultModelWithDefaults

`func NewStepResultModelWithDefaults() *StepResultModel`

NewStepResultModelWithDefaults instantiates a new StepResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepId

`func (o *StepResultModel) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *StepResultModel) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *StepResultModel) SetStepId(v string)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *StepResultModel) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetOutcome

`func (o *StepResultModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *StepResultModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *StepResultModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *StepResultModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *StepResultModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *StepResultModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetSharedStepVersionId

`func (o *StepResultModel) GetSharedStepVersionId() string`

GetSharedStepVersionId returns the SharedStepVersionId field if non-nil, zero value otherwise.

### GetSharedStepVersionIdOk

`func (o *StepResultModel) GetSharedStepVersionIdOk() (*string, bool)`

GetSharedStepVersionIdOk returns a tuple with the SharedStepVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepVersionId

`func (o *StepResultModel) SetSharedStepVersionId(v string)`

SetSharedStepVersionId sets SharedStepVersionId field to given value.

### HasSharedStepVersionId

`func (o *StepResultModel) HasSharedStepVersionId() bool`

HasSharedStepVersionId returns a boolean if a field has been set.

### SetSharedStepVersionIdNil

`func (o *StepResultModel) SetSharedStepVersionIdNil(b bool)`

 SetSharedStepVersionIdNil sets the value for SharedStepVersionId to be an explicit nil

### UnsetSharedStepVersionId
`func (o *StepResultModel) UnsetSharedStepVersionId()`

UnsetSharedStepVersionId ensures that no value is present for SharedStepVersionId, not even an explicit nil
### GetSharedStepResults

`func (o *StepResultModel) GetSharedStepResults() []SharedStepResultModel`

GetSharedStepResults returns the SharedStepResults field if non-nil, zero value otherwise.

### GetSharedStepResultsOk

`func (o *StepResultModel) GetSharedStepResultsOk() (*[]SharedStepResultModel, bool)`

GetSharedStepResultsOk returns a tuple with the SharedStepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepResults

`func (o *StepResultModel) SetSharedStepResults(v []SharedStepResultModel)`

SetSharedStepResults sets SharedStepResults field to given value.

### HasSharedStepResults

`func (o *StepResultModel) HasSharedStepResults() bool`

HasSharedStepResults returns a boolean if a field has been set.

### SetSharedStepResultsNil

`func (o *StepResultModel) SetSharedStepResultsNil(b bool)`

 SetSharedStepResultsNil sets the value for SharedStepResults to be an explicit nil

### UnsetSharedStepResults
`func (o *StepResultModel) UnsetSharedStepResults()`

UnsetSharedStepResults ensures that no value is present for SharedStepResults, not even an explicit nil
### GetComment

`func (o *StepResultModel) GetComment() StepCommentModel`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StepResultModel) GetCommentOk() (*StepCommentModel, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StepResultModel) SetComment(v StepCommentModel)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StepResultModel) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


