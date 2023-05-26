# ProjectShortestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of project | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates whether the project is deleted | [optional] 
**GlobalId** | Pointer to **int64** | Global ID of project | [optional] 
**Name** | Pointer to **NullableString** | Name of project | [optional] 

## Methods

### NewProjectShortestModel

`func NewProjectShortestModel() *ProjectShortestModel`

NewProjectShortestModel instantiates a new ProjectShortestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectShortestModelWithDefaults

`func NewProjectShortestModelWithDefaults() *ProjectShortestModel`

NewProjectShortestModelWithDefaults instantiates a new ProjectShortestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectShortestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectShortestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectShortestModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectShortestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ProjectShortestModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectShortestModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectShortestModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ProjectShortestModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetGlobalId

`func (o *ProjectShortestModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ProjectShortestModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ProjectShortestModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *ProjectShortestModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetName

`func (o *ProjectShortestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectShortestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectShortestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectShortestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectShortestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectShortestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


