# WorkItemIdModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Used for search WorkItem. Internal identifier has a Guid data format. Global identifier has an integer data format | 

## Methods

### NewWorkItemIdModel

`func NewWorkItemIdModel(id string, ) *WorkItemIdModel`

NewWorkItemIdModel instantiates a new WorkItemIdModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemIdModelWithDefaults

`func NewWorkItemIdModelWithDefaults() *WorkItemIdModel`

NewWorkItemIdModelWithDefaults instantiates a new WorkItemIdModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemIdModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemIdModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemIdModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


