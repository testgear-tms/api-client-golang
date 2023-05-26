# CustomAttributeOptionPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableString** | Value of the attribute option | [optional] 
**IsDefault** | Pointer to **bool** | Indicates if the attribute option is used by default | [optional] 

## Methods

### NewCustomAttributeOptionPostModel

`func NewCustomAttributeOptionPostModel() *CustomAttributeOptionPostModel`

NewCustomAttributeOptionPostModel instantiates a new CustomAttributeOptionPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeOptionPostModelWithDefaults

`func NewCustomAttributeOptionPostModelWithDefaults() *CustomAttributeOptionPostModel`

NewCustomAttributeOptionPostModelWithDefaults instantiates a new CustomAttributeOptionPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomAttributeOptionPostModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomAttributeOptionPostModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomAttributeOptionPostModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomAttributeOptionPostModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomAttributeOptionPostModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomAttributeOptionPostModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *CustomAttributeOptionPostModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CustomAttributeOptionPostModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CustomAttributeOptionPostModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CustomAttributeOptionPostModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


