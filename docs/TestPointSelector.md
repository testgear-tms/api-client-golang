# TestPointSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | **string** | Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs. | 
**WorkItemIds** | **[]string** | Specifies the work item GUIDs, from which test points are created. You can specify several GUIDs. | 

## Methods

### NewTestPointSelector

`func NewTestPointSelector(configurationId string, workItemIds []string, ) *TestPointSelector`

NewTestPointSelector instantiates a new TestPointSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointSelectorWithDefaults

`func NewTestPointSelectorWithDefaults() *TestPointSelector`

NewTestPointSelectorWithDefaults instantiates a new TestPointSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationId

`func (o *TestPointSelector) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointSelector) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointSelector) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetWorkItemIds

`func (o *TestPointSelector) GetWorkItemIds() []string`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *TestPointSelector) GetWorkItemIdsOk() (*[]string, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *TestPointSelector) SetWorkItemIds(v []string)`

SetWorkItemIds sets WorkItemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


