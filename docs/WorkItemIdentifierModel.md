# WorkItemIdentifierModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Used for search WorkItem. Internal identifier has a Guid data format. Global identifier has an integer data format | [optional] 
**GlobalId** | Pointer to **int64** |  | [optional] 

## Methods

### NewWorkItemIdentifierModel

`func NewWorkItemIdentifierModel() *WorkItemIdentifierModel`

NewWorkItemIdentifierModel instantiates a new WorkItemIdentifierModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemIdentifierModelWithDefaults

`func NewWorkItemIdentifierModelWithDefaults() *WorkItemIdentifierModel`

NewWorkItemIdentifierModelWithDefaults instantiates a new WorkItemIdentifierModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemIdentifierModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemIdentifierModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemIdentifierModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkItemIdentifierModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGlobalId

`func (o *WorkItemIdentifierModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemIdentifierModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemIdentifierModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *WorkItemIdentifierModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


