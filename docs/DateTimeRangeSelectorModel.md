# DateTimeRangeSelectorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **NullableTime** |  | [optional] 
**To** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDateTimeRangeSelectorModel

`func NewDateTimeRangeSelectorModel() *DateTimeRangeSelectorModel`

NewDateTimeRangeSelectorModel instantiates a new DateTimeRangeSelectorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeRangeSelectorModelWithDefaults

`func NewDateTimeRangeSelectorModelWithDefaults() *DateTimeRangeSelectorModel`

NewDateTimeRangeSelectorModelWithDefaults instantiates a new DateTimeRangeSelectorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *DateTimeRangeSelectorModel) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DateTimeRangeSelectorModel) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DateTimeRangeSelectorModel) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DateTimeRangeSelectorModel) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *DateTimeRangeSelectorModel) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *DateTimeRangeSelectorModel) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetTo

`func (o *DateTimeRangeSelectorModel) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DateTimeRangeSelectorModel) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DateTimeRangeSelectorModel) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *DateTimeRangeSelectorModel) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *DateTimeRangeSelectorModel) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *DateTimeRangeSelectorModel) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


