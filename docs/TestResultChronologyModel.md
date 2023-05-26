# TestResultChronologyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | Pointer to **NullableString** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewTestResultChronologyModel

`func NewTestResultChronologyModel() *TestResultChronologyModel`

NewTestResultChronologyModel instantiates a new TestResultChronologyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultChronologyModelWithDefaults

`func NewTestResultChronologyModelWithDefaults() *TestResultChronologyModel`

NewTestResultChronologyModelWithDefaults instantiates a new TestResultChronologyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *TestResultChronologyModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultChronologyModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultChronologyModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultChronologyModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultChronologyModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultChronologyModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetCount

`func (o *TestResultChronologyModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TestResultChronologyModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TestResultChronologyModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TestResultChronologyModel) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


