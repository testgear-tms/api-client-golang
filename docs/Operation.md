# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **interface{}** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Op** | Pointer to **NullableString** |  | [optional] 
**From** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOperation

`func NewOperation() *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Operation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Operation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Operation) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Operation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Operation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Operation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetPath

`func (o *Operation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Operation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Operation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Operation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Operation) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Operation) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetOp

`func (o *Operation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *Operation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *Operation) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *Operation) HasOp() bool`

HasOp returns a boolean if a field has been set.

### SetOpNil

`func (o *Operation) SetOpNil(b bool)`

 SetOpNil sets the value for Op to be an explicit nil

### UnsetOp
`func (o *Operation) UnsetOp()`

UnsetOp ensures that no value is present for Op, not even an explicit nil
### GetFrom

`func (o *Operation) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Operation) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Operation) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Operation) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *Operation) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *Operation) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


