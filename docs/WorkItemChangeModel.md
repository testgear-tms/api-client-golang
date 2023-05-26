# WorkItemChangeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**WorkItemId** | Pointer to **string** |  | [optional] 
**OldVersionId** | Pointer to **string** |  | [optional] 
**NewVersionId** | Pointer to **string** |  | [optional] 
**WorkItemChangedFields** | Pointer to [**WorkItemChangedFieldsViewModel**](WorkItemChangedFieldsViewModel.md) |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewWorkItemChangeModel

`func NewWorkItemChangeModel() *WorkItemChangeModel`

NewWorkItemChangeModel instantiates a new WorkItemChangeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemChangeModelWithDefaults

`func NewWorkItemChangeModelWithDefaults() *WorkItemChangeModel`

NewWorkItemChangeModelWithDefaults instantiates a new WorkItemChangeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemChangeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemChangeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemChangeModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkItemChangeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkItemId

`func (o *WorkItemChangeModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *WorkItemChangeModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *WorkItemChangeModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *WorkItemChangeModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### GetOldVersionId

`func (o *WorkItemChangeModel) GetOldVersionId() string`

GetOldVersionId returns the OldVersionId field if non-nil, zero value otherwise.

### GetOldVersionIdOk

`func (o *WorkItemChangeModel) GetOldVersionIdOk() (*string, bool)`

GetOldVersionIdOk returns a tuple with the OldVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldVersionId

`func (o *WorkItemChangeModel) SetOldVersionId(v string)`

SetOldVersionId sets OldVersionId field to given value.

### HasOldVersionId

`func (o *WorkItemChangeModel) HasOldVersionId() bool`

HasOldVersionId returns a boolean if a field has been set.

### GetNewVersionId

`func (o *WorkItemChangeModel) GetNewVersionId() string`

GetNewVersionId returns the NewVersionId field if non-nil, zero value otherwise.

### GetNewVersionIdOk

`func (o *WorkItemChangeModel) GetNewVersionIdOk() (*string, bool)`

GetNewVersionIdOk returns a tuple with the NewVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersionId

`func (o *WorkItemChangeModel) SetNewVersionId(v string)`

SetNewVersionId sets NewVersionId field to given value.

### HasNewVersionId

`func (o *WorkItemChangeModel) HasNewVersionId() bool`

HasNewVersionId returns a boolean if a field has been set.

### GetWorkItemChangedFields

`func (o *WorkItemChangeModel) GetWorkItemChangedFields() WorkItemChangedFieldsViewModel`

GetWorkItemChangedFields returns the WorkItemChangedFields field if non-nil, zero value otherwise.

### GetWorkItemChangedFieldsOk

`func (o *WorkItemChangeModel) GetWorkItemChangedFieldsOk() (*WorkItemChangedFieldsViewModel, bool)`

GetWorkItemChangedFieldsOk returns a tuple with the WorkItemChangedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemChangedFields

`func (o *WorkItemChangeModel) SetWorkItemChangedFields(v WorkItemChangedFieldsViewModel)`

SetWorkItemChangedFields sets WorkItemChangedFields field to given value.

### HasWorkItemChangedFields

`func (o *WorkItemChangeModel) HasWorkItemChangedFields() bool`

HasWorkItemChangedFields returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkItemChangeModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemChangeModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemChangeModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *WorkItemChangeModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *WorkItemChangeModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemChangeModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemChangeModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WorkItemChangeModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *WorkItemChangeModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *WorkItemChangeModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


