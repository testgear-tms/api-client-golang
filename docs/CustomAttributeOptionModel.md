# CustomAttributeOptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the attribute option | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates if the attributes option is deleted | [optional] 
**Value** | Pointer to **NullableString** | Value of the attribute option | [optional] 
**IsDefault** | Pointer to **bool** | Indicates if the attribute option is used by default | [optional] 

## Methods

### NewCustomAttributeOptionModel

`func NewCustomAttributeOptionModel() *CustomAttributeOptionModel`

NewCustomAttributeOptionModel instantiates a new CustomAttributeOptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeOptionModelWithDefaults

`func NewCustomAttributeOptionModelWithDefaults() *CustomAttributeOptionModel`

NewCustomAttributeOptionModelWithDefaults instantiates a new CustomAttributeOptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAttributeOptionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributeOptionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributeOptionModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomAttributeOptionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CustomAttributeOptionModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributeOptionModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributeOptionModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CustomAttributeOptionModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetValue

`func (o *CustomAttributeOptionModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomAttributeOptionModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomAttributeOptionModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomAttributeOptionModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomAttributeOptionModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomAttributeOptionModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *CustomAttributeOptionModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CustomAttributeOptionModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CustomAttributeOptionModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CustomAttributeOptionModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


