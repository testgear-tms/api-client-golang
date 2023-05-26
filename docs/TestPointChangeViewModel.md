# TestPointChangeViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**TestPointCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewTestPointChangeViewModel

`func NewTestPointChangeViewModel() *TestPointChangeViewModel`

NewTestPointChangeViewModel instantiates a new TestPointChangeViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointChangeViewModelWithDefaults

`func NewTestPointChangeViewModelWithDefaults() *TestPointChangeViewModel`

NewTestPointChangeViewModelWithDefaults instantiates a new TestPointChangeViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TestPointChangeViewModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TestPointChangeViewModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TestPointChangeViewModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TestPointChangeViewModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *TestPointChangeViewModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TestPointChangeViewModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TestPointChangeViewModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *TestPointChangeViewModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *TestPointChangeViewModel) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *TestPointChangeViewModel) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetTestPointCount

`func (o *TestPointChangeViewModel) GetTestPointCount() int64`

GetTestPointCount returns the TestPointCount field if non-nil, zero value otherwise.

### GetTestPointCountOk

`func (o *TestPointChangeViewModel) GetTestPointCountOk() (*int64, bool)`

GetTestPointCountOk returns a tuple with the TestPointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointCount

`func (o *TestPointChangeViewModel) SetTestPointCount(v int64)`

SetTestPointCount sets TestPointCount field to given value.

### HasTestPointCount

`func (o *TestPointChangeViewModel) HasTestPointCount() bool`

HasTestPointCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


