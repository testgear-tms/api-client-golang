# TestPointShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestSuiteId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**TesterId** | Pointer to **NullableString** |  | [optional] 
**WorkItemId** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** | Applies one of these values: Blocked, NoResults, Failed, Passed | [optional] 
**LastTestResultId** | Pointer to **NullableString** |  | [optional] 
**IterationId** | Pointer to **string** |  | [optional] 

## Methods

### NewTestPointShortModel

`func NewTestPointShortModel() *TestPointShortModel`

NewTestPointShortModel instantiates a new TestPointShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointShortModelWithDefaults

`func NewTestPointShortModelWithDefaults() *TestPointShortModel`

NewTestPointShortModelWithDefaults instantiates a new TestPointShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestSuiteId

`func (o *TestPointShortModel) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointShortModel) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointShortModel) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.

### HasTestSuiteId

`func (o *TestPointShortModel) HasTestSuiteId() bool`

HasTestSuiteId returns a boolean if a field has been set.

### GetId

`func (o *TestPointShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointShortModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestPointShortModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTesterId

`func (o *TestPointShortModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointShortModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointShortModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointShortModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointShortModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointShortModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetWorkItemId

`func (o *TestPointShortModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointShortModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointShortModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *TestPointShortModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *TestPointShortModel) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *TestPointShortModel) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
### GetConfigurationId

`func (o *TestPointShortModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointShortModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointShortModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestPointShortModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *TestPointShortModel) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *TestPointShortModel) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetStatus

`func (o *TestPointShortModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointShortModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointShortModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPointShortModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPointShortModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPointShortModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastTestResultId

`func (o *TestPointShortModel) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *TestPointShortModel) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *TestPointShortModel) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *TestPointShortModel) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *TestPointShortModel) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *TestPointShortModel) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetIterationId

`func (o *TestPointShortModel) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *TestPointShortModel) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *TestPointShortModel) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *TestPointShortModel) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


