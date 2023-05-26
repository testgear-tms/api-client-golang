# NotificationQueryFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to [**[]NotificationTypeModel**](NotificationTypeModel.md) |  | [optional] 
**IsRead** | Pointer to **NullableBool** |  | [optional] 
**CreatedDate** | Pointer to [**DateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) |  | [optional] 

## Methods

### NewNotificationQueryFilterModel

`func NewNotificationQueryFilterModel() *NotificationQueryFilterModel`

NewNotificationQueryFilterModel instantiates a new NotificationQueryFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationQueryFilterModelWithDefaults

`func NewNotificationQueryFilterModelWithDefaults() *NotificationQueryFilterModel`

NewNotificationQueryFilterModelWithDefaults instantiates a new NotificationQueryFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *NotificationQueryFilterModel) GetTypes() []NotificationTypeModel`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *NotificationQueryFilterModel) GetTypesOk() (*[]NotificationTypeModel, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *NotificationQueryFilterModel) SetTypes(v []NotificationTypeModel)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *NotificationQueryFilterModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *NotificationQueryFilterModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *NotificationQueryFilterModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetIsRead

`func (o *NotificationQueryFilterModel) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *NotificationQueryFilterModel) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *NotificationQueryFilterModel) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *NotificationQueryFilterModel) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### SetIsReadNil

`func (o *NotificationQueryFilterModel) SetIsReadNil(b bool)`

 SetIsReadNil sets the value for IsRead to be an explicit nil

### UnsetIsRead
`func (o *NotificationQueryFilterModel) UnsetIsRead()`

UnsetIsRead ensures that no value is present for IsRead, not even an explicit nil
### GetCreatedDate

`func (o *NotificationQueryFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *NotificationQueryFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *NotificationQueryFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *NotificationQueryFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


