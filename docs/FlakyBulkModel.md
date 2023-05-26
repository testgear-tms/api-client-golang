# FlakyBulkModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutotestSelect** | Pointer to [**AutotestSelectModel**](AutotestSelectModel.md) |  | [optional] 
**Value** | **bool** | Are autotests flaky | 

## Methods

### NewFlakyBulkModel

`func NewFlakyBulkModel(value bool, ) *FlakyBulkModel`

NewFlakyBulkModel instantiates a new FlakyBulkModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlakyBulkModelWithDefaults

`func NewFlakyBulkModelWithDefaults() *FlakyBulkModel`

NewFlakyBulkModelWithDefaults instantiates a new FlakyBulkModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutotestSelect

`func (o *FlakyBulkModel) GetAutotestSelect() AutotestSelectModel`

GetAutotestSelect returns the AutotestSelect field if non-nil, zero value otherwise.

### GetAutotestSelectOk

`func (o *FlakyBulkModel) GetAutotestSelectOk() (*AutotestSelectModel, bool)`

GetAutotestSelectOk returns a tuple with the AutotestSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestSelect

`func (o *FlakyBulkModel) SetAutotestSelect(v AutotestSelectModel)`

SetAutotestSelect sets AutotestSelect field to given value.

### HasAutotestSelect

`func (o *FlakyBulkModel) HasAutotestSelect() bool`

HasAutotestSelect returns a boolean if a field has been set.

### GetValue

`func (o *FlakyBulkModel) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FlakyBulkModel) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FlakyBulkModel) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


