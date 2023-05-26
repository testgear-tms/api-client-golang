# TestPointResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestPointId** | Pointer to **NullableString** |  | [optional] 
**AggregatedOutcome** | Pointer to **NullableString** |  | [optional] 
**WorkItemGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**WorkItemName** | Pointer to **NullableString** |  | [optional] 
**ConfigurationName** | Pointer to **NullableString** |  | [optional] 
**TestResults** | Pointer to [**[]TestResultShortModel**](TestResultShortModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 

## Methods

### NewTestPointResultModel

`func NewTestPointResultModel() *TestPointResultModel`

NewTestPointResultModel instantiates a new TestPointResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointResultModelWithDefaults

`func NewTestPointResultModelWithDefaults() *TestPointResultModel`

NewTestPointResultModelWithDefaults instantiates a new TestPointResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestPointId

`func (o *TestPointResultModel) GetTestPointId() string`

GetTestPointId returns the TestPointId field if non-nil, zero value otherwise.

### GetTestPointIdOk

`func (o *TestPointResultModel) GetTestPointIdOk() (*string, bool)`

GetTestPointIdOk returns a tuple with the TestPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointId

`func (o *TestPointResultModel) SetTestPointId(v string)`

SetTestPointId sets TestPointId field to given value.

### HasTestPointId

`func (o *TestPointResultModel) HasTestPointId() bool`

HasTestPointId returns a boolean if a field has been set.

### SetTestPointIdNil

`func (o *TestPointResultModel) SetTestPointIdNil(b bool)`

 SetTestPointIdNil sets the value for TestPointId to be an explicit nil

### UnsetTestPointId
`func (o *TestPointResultModel) UnsetTestPointId()`

UnsetTestPointId ensures that no value is present for TestPointId, not even an explicit nil
### GetAggregatedOutcome

`func (o *TestPointResultModel) GetAggregatedOutcome() string`

GetAggregatedOutcome returns the AggregatedOutcome field if non-nil, zero value otherwise.

### GetAggregatedOutcomeOk

`func (o *TestPointResultModel) GetAggregatedOutcomeOk() (*string, bool)`

GetAggregatedOutcomeOk returns a tuple with the AggregatedOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedOutcome

`func (o *TestPointResultModel) SetAggregatedOutcome(v string)`

SetAggregatedOutcome sets AggregatedOutcome field to given value.

### HasAggregatedOutcome

`func (o *TestPointResultModel) HasAggregatedOutcome() bool`

HasAggregatedOutcome returns a boolean if a field has been set.

### SetAggregatedOutcomeNil

`func (o *TestPointResultModel) SetAggregatedOutcomeNil(b bool)`

 SetAggregatedOutcomeNil sets the value for AggregatedOutcome to be an explicit nil

### UnsetAggregatedOutcome
`func (o *TestPointResultModel) UnsetAggregatedOutcome()`

UnsetAggregatedOutcome ensures that no value is present for AggregatedOutcome, not even an explicit nil
### GetWorkItemGlobalId

`func (o *TestPointResultModel) GetWorkItemGlobalId() int64`

GetWorkItemGlobalId returns the WorkItemGlobalId field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdOk

`func (o *TestPointResultModel) GetWorkItemGlobalIdOk() (*int64, bool)`

GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalId

`func (o *TestPointResultModel) SetWorkItemGlobalId(v int64)`

SetWorkItemGlobalId sets WorkItemGlobalId field to given value.

### HasWorkItemGlobalId

`func (o *TestPointResultModel) HasWorkItemGlobalId() bool`

HasWorkItemGlobalId returns a boolean if a field has been set.

### SetWorkItemGlobalIdNil

`func (o *TestPointResultModel) SetWorkItemGlobalIdNil(b bool)`

 SetWorkItemGlobalIdNil sets the value for WorkItemGlobalId to be an explicit nil

### UnsetWorkItemGlobalId
`func (o *TestPointResultModel) UnsetWorkItemGlobalId()`

UnsetWorkItemGlobalId ensures that no value is present for WorkItemGlobalId, not even an explicit nil
### GetWorkItemName

`func (o *TestPointResultModel) GetWorkItemName() string`

GetWorkItemName returns the WorkItemName field if non-nil, zero value otherwise.

### GetWorkItemNameOk

`func (o *TestPointResultModel) GetWorkItemNameOk() (*string, bool)`

GetWorkItemNameOk returns a tuple with the WorkItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemName

`func (o *TestPointResultModel) SetWorkItemName(v string)`

SetWorkItemName sets WorkItemName field to given value.

### HasWorkItemName

`func (o *TestPointResultModel) HasWorkItemName() bool`

HasWorkItemName returns a boolean if a field has been set.

### SetWorkItemNameNil

`func (o *TestPointResultModel) SetWorkItemNameNil(b bool)`

 SetWorkItemNameNil sets the value for WorkItemName to be an explicit nil

### UnsetWorkItemName
`func (o *TestPointResultModel) UnsetWorkItemName()`

UnsetWorkItemName ensures that no value is present for WorkItemName, not even an explicit nil
### GetConfigurationName

`func (o *TestPointResultModel) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestPointResultModel) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestPointResultModel) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *TestPointResultModel) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### SetConfigurationNameNil

`func (o *TestPointResultModel) SetConfigurationNameNil(b bool)`

 SetConfigurationNameNil sets the value for ConfigurationName to be an explicit nil

### UnsetConfigurationName
`func (o *TestPointResultModel) UnsetConfigurationName()`

UnsetConfigurationName ensures that no value is present for ConfigurationName, not even an explicit nil
### GetTestResults

`func (o *TestPointResultModel) GetTestResults() []TestResultShortModel`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestPointResultModel) GetTestResultsOk() (*[]TestResultShortModel, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestPointResultModel) SetTestResults(v []TestResultShortModel)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *TestPointResultModel) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.

### SetTestResultsNil

`func (o *TestPointResultModel) SetTestResultsNil(b bool)`

 SetTestResultsNil sets the value for TestResults to be an explicit nil

### UnsetTestResults
`func (o *TestPointResultModel) UnsetTestResults()`

UnsetTestResults ensures that no value is present for TestResults, not even an explicit nil
### GetAttachments

`func (o *TestPointResultModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestPointResultModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestPointResultModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestPointResultModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestPointResultModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestPointResultModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


