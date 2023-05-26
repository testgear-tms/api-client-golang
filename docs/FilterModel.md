# FilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to [**WorkItemSearchQueryModel**](WorkItemSearchQueryModel.md) |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**FieldsToShow** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the entity | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the entity is deleted | [optional] 

## Methods

### NewFilterModel

`func NewFilterModel() *FilterModel`

NewFilterModel instantiates a new FilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterModelWithDefaults

`func NewFilterModelWithDefaults() *FilterModel`

NewFilterModelWithDefaults instantiates a new FilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *FilterModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *FilterModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *FilterModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *FilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *FilterModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *FilterModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *FilterModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *FilterModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *FilterModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *FilterModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *FilterModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *FilterModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *FilterModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *FilterModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *FilterModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *FilterModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *FilterModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *FilterModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *FilterModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *FilterModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetData

`func (o *FilterModel) GetData() WorkItemSearchQueryModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FilterModel) GetDataOk() (*WorkItemSearchQueryModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FilterModel) SetData(v WorkItemSearchQueryModel)`

SetData sets Data field to given value.

### HasData

`func (o *FilterModel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetProjectId

`func (o *FilterModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FilterModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FilterModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FilterModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetFieldsToShow

`func (o *FilterModel) GetFieldsToShow() interface{}`

GetFieldsToShow returns the FieldsToShow field if non-nil, zero value otherwise.

### GetFieldsToShowOk

`func (o *FilterModel) GetFieldsToShowOk() (*interface{}, bool)`

GetFieldsToShowOk returns a tuple with the FieldsToShow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsToShow

`func (o *FilterModel) SetFieldsToShow(v interface{})`

SetFieldsToShow sets FieldsToShow field to given value.

### HasFieldsToShow

`func (o *FilterModel) HasFieldsToShow() bool`

HasFieldsToShow returns a boolean if a field has been set.

### SetFieldsToShowNil

`func (o *FilterModel) SetFieldsToShowNil(b bool)`

 SetFieldsToShowNil sets the value for FieldsToShow to be an explicit nil

### UnsetFieldsToShow
`func (o *FilterModel) UnsetFieldsToShow()`

UnsetFieldsToShow ensures that no value is present for FieldsToShow, not even an explicit nil
### GetName

`func (o *FilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *FilterModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FilterModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *FilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *FilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *FilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *FilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


