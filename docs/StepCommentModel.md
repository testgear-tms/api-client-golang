# StepCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **NullableString** |  | [optional] 
**StepId** | **string** |  | 
**ParentStepId** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 
**TestResultId** | **string** |  | 
**CreatedById** | Pointer to **string** |  | [optional] 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewStepCommentModel

`func NewStepCommentModel(stepId string, testResultId string, ) *StepCommentModel`

NewStepCommentModel instantiates a new StepCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepCommentModelWithDefaults

`func NewStepCommentModelWithDefaults() *StepCommentModel`

NewStepCommentModelWithDefaults instantiates a new StepCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StepCommentModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StepCommentModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StepCommentModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StepCommentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetText

`func (o *StepCommentModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StepCommentModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StepCommentModel) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StepCommentModel) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *StepCommentModel) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *StepCommentModel) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetStepId

`func (o *StepCommentModel) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *StepCommentModel) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *StepCommentModel) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetParentStepId

`func (o *StepCommentModel) GetParentStepId() string`

GetParentStepId returns the ParentStepId field if non-nil, zero value otherwise.

### GetParentStepIdOk

`func (o *StepCommentModel) GetParentStepIdOk() (*string, bool)`

GetParentStepIdOk returns a tuple with the ParentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStepId

`func (o *StepCommentModel) SetParentStepId(v string)`

SetParentStepId sets ParentStepId field to given value.

### HasParentStepId

`func (o *StepCommentModel) HasParentStepId() bool`

HasParentStepId returns a boolean if a field has been set.

### SetParentStepIdNil

`func (o *StepCommentModel) SetParentStepIdNil(b bool)`

 SetParentStepIdNil sets the value for ParentStepId to be an explicit nil

### UnsetParentStepId
`func (o *StepCommentModel) UnsetParentStepId()`

UnsetParentStepId ensures that no value is present for ParentStepId, not even an explicit nil
### GetAttachments

`func (o *StepCommentModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *StepCommentModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *StepCommentModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *StepCommentModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *StepCommentModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *StepCommentModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetTestResultId

`func (o *StepCommentModel) GetTestResultId() string`

GetTestResultId returns the TestResultId field if non-nil, zero value otherwise.

### GetTestResultIdOk

`func (o *StepCommentModel) GetTestResultIdOk() (*string, bool)`

GetTestResultIdOk returns a tuple with the TestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultId

`func (o *StepCommentModel) SetTestResultId(v string)`

SetTestResultId sets TestResultId field to given value.


### GetCreatedById

`func (o *StepCommentModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *StepCommentModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *StepCommentModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *StepCommentModel) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetModifiedById

`func (o *StepCommentModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *StepCommentModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *StepCommentModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *StepCommentModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *StepCommentModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *StepCommentModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedDate

`func (o *StepCommentModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *StepCommentModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *StepCommentModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *StepCommentModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetModifiedDate

`func (o *StepCommentModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *StepCommentModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *StepCommentModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *StepCommentModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *StepCommentModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *StepCommentModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


