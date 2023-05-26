# TestPointPutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TesterId** | Pointer to **NullableString** |  | [optional] 
**IterationId** | Pointer to **string** |  | [optional] 
**WorkItemId** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | Pointer to **NullableString** |  | [optional] 
**TestSuiteId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**LastTestResultId** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewTestPointPutModel

`func NewTestPointPutModel() *TestPointPutModel`

NewTestPointPutModel instantiates a new TestPointPutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointPutModelWithDefaults

`func NewTestPointPutModelWithDefaults() *TestPointPutModel`

NewTestPointPutModelWithDefaults instantiates a new TestPointPutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTesterId

`func (o *TestPointPutModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointPutModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointPutModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointPutModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointPutModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointPutModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetIterationId

`func (o *TestPointPutModel) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *TestPointPutModel) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *TestPointPutModel) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *TestPointPutModel) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### GetWorkItemId

`func (o *TestPointPutModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointPutModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointPutModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *TestPointPutModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *TestPointPutModel) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *TestPointPutModel) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
### GetConfigurationId

`func (o *TestPointPutModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointPutModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointPutModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestPointPutModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *TestPointPutModel) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *TestPointPutModel) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetTestSuiteId

`func (o *TestPointPutModel) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointPutModel) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointPutModel) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.

### HasTestSuiteId

`func (o *TestPointPutModel) HasTestSuiteId() bool`

HasTestSuiteId returns a boolean if a field has been set.

### GetStatus

`func (o *TestPointPutModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointPutModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointPutModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPointPutModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPointPutModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPointPutModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastTestResultId

`func (o *TestPointPutModel) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *TestPointPutModel) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *TestPointPutModel) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *TestPointPutModel) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *TestPointPutModel) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *TestPointPutModel) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetId

`func (o *TestPointPutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointPutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointPutModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestPointPutModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TestPointPutModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestPointPutModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestPointPutModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestPointPutModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


