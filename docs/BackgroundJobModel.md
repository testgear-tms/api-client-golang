# BackgroundJobModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **NullableString** |  | [optional] 
**JobType** | Pointer to [**BackgroundJobType**](BackgroundJobType.md) |  | [optional] 
**State** | Pointer to [**BackgroundJobState**](BackgroundJobState.md) |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**Progress** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to [**[]BackgroundJobAttachmentModel**](BackgroundJobAttachmentModel.md) |  | [optional] 

## Methods

### NewBackgroundJobModel

`func NewBackgroundJobModel() *BackgroundJobModel`

NewBackgroundJobModel instantiates a new BackgroundJobModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundJobModelWithDefaults

`func NewBackgroundJobModelWithDefaults() *BackgroundJobModel`

NewBackgroundJobModelWithDefaults instantiates a new BackgroundJobModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackgroundJobModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackgroundJobModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackgroundJobModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackgroundJobModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *BackgroundJobModel) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackgroundJobModel) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackgroundJobModel) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BackgroundJobModel) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *BackgroundJobModel) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *BackgroundJobModel) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobType

`func (o *BackgroundJobModel) GetJobType() BackgroundJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BackgroundJobModel) GetJobTypeOk() (*BackgroundJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BackgroundJobModel) SetJobType(v BackgroundJobType)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BackgroundJobModel) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetState

`func (o *BackgroundJobModel) GetState() BackgroundJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackgroundJobModel) GetStateOk() (*BackgroundJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackgroundJobModel) SetState(v BackgroundJobState)`

SetState sets State field to given value.

### HasState

`func (o *BackgroundJobModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIsDeleted

`func (o *BackgroundJobModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *BackgroundJobModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *BackgroundJobModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *BackgroundJobModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetProgress

`func (o *BackgroundJobModel) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BackgroundJobModel) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BackgroundJobModel) SetProgress(v int64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *BackgroundJobModel) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStartDate

`func (o *BackgroundJobModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackgroundJobModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackgroundJobModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackgroundJobModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BackgroundJobModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BackgroundJobModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BackgroundJobModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BackgroundJobModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BackgroundJobModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BackgroundJobModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BackgroundJobModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BackgroundJobModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetError

`func (o *BackgroundJobModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BackgroundJobModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BackgroundJobModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BackgroundJobModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BackgroundJobModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BackgroundJobModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAttachments

`func (o *BackgroundJobModel) GetAttachments() []BackgroundJobAttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BackgroundJobModel) GetAttachmentsOk() (*[]BackgroundJobAttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BackgroundJobModel) SetAttachments(v []BackgroundJobAttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *BackgroundJobModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *BackgroundJobModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *BackgroundJobModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


