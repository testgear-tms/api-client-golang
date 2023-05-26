# TestPlanChangeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TestPlanId** | Pointer to **string** |  | [optional] 
**TestPlanChangedFields** | Pointer to [**TestPlanChangedFieldsViewModel**](TestPlanChangedFieldsViewModel.md) |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTestPlanChangeModel

`func NewTestPlanChangeModel() *TestPlanChangeModel`

NewTestPlanChangeModel instantiates a new TestPlanChangeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanChangeModelWithDefaults

`func NewTestPlanChangeModelWithDefaults() *TestPlanChangeModel`

NewTestPlanChangeModelWithDefaults instantiates a new TestPlanChangeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPlanChangeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPlanChangeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPlanChangeModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestPlanChangeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTestPlanId

`func (o *TestPlanChangeModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestPlanChangeModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestPlanChangeModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestPlanChangeModel) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### GetTestPlanChangedFields

`func (o *TestPlanChangeModel) GetTestPlanChangedFields() TestPlanChangedFieldsViewModel`

GetTestPlanChangedFields returns the TestPlanChangedFields field if non-nil, zero value otherwise.

### GetTestPlanChangedFieldsOk

`func (o *TestPlanChangeModel) GetTestPlanChangedFieldsOk() (*TestPlanChangedFieldsViewModel, bool)`

GetTestPlanChangedFieldsOk returns a tuple with the TestPlanChangedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanChangedFields

`func (o *TestPlanChangeModel) SetTestPlanChangedFields(v TestPlanChangedFieldsViewModel)`

SetTestPlanChangedFields sets TestPlanChangedFields field to given value.

### HasTestPlanChangedFields

`func (o *TestPlanChangeModel) HasTestPlanChangedFields() bool`

HasTestPlanChangedFields returns a boolean if a field has been set.

### GetCreatedById

`func (o *TestPlanChangeModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPlanChangeModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPlanChangeModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *TestPlanChangeModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TestPlanChangeModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPlanChangeModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPlanChangeModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPlanChangeModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPlanChangeModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPlanChangeModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


