# AutotestsSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**AutotestFilterModel**](AutotestFilterModel.md) |  | [optional] 
**Includes** | Pointer to [**SearchAutoTestsQueryIncludesModel**](SearchAutoTestsQueryIncludesModel.md) |  | [optional] 

## Methods

### NewAutotestsSelectModel

`func NewAutotestsSelectModel() *AutotestsSelectModel`

NewAutotestsSelectModel instantiates a new AutotestsSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotestsSelectModelWithDefaults

`func NewAutotestsSelectModelWithDefaults() *AutotestsSelectModel`

NewAutotestsSelectModelWithDefaults instantiates a new AutotestsSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AutotestsSelectModel) GetFilter() AutotestFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AutotestsSelectModel) GetFilterOk() (*AutotestFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AutotestsSelectModel) SetFilter(v AutotestFilterModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AutotestsSelectModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetIncludes

`func (o *AutotestsSelectModel) GetIncludes() SearchAutoTestsQueryIncludesModel`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *AutotestsSelectModel) GetIncludesOk() (*SearchAutoTestsQueryIncludesModel, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *AutotestsSelectModel) SetIncludes(v SearchAutoTestsQueryIncludesModel)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *AutotestsSelectModel) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


