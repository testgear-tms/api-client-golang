# LastTestResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TestRunId** | Pointer to **string** |  | [optional] 
**AutoTestId** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**WorkItemVersionId** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 

## Methods

### NewLastTestResultModel

`func NewLastTestResultModel() *LastTestResultModel`

NewLastTestResultModel instantiates a new LastTestResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastTestResultModelWithDefaults

`func NewLastTestResultModelWithDefaults() *LastTestResultModel`

NewLastTestResultModelWithDefaults instantiates a new LastTestResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LastTestResultModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LastTestResultModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LastTestResultModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LastTestResultModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTestRunId

`func (o *LastTestResultModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *LastTestResultModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *LastTestResultModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *LastTestResultModel) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### GetAutoTestId

`func (o *LastTestResultModel) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *LastTestResultModel) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *LastTestResultModel) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *LastTestResultModel) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *LastTestResultModel) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *LastTestResultModel) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetComment

`func (o *LastTestResultModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LastTestResultModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LastTestResultModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LastTestResultModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *LastTestResultModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *LastTestResultModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *LastTestResultModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LastTestResultModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LastTestResultModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LastTestResultModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *LastTestResultModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *LastTestResultModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetWorkItemVersionId

`func (o *LastTestResultModel) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *LastTestResultModel) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *LastTestResultModel) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *LastTestResultModel) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### GetAttachments

`func (o *LastTestResultModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *LastTestResultModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *LastTestResultModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *LastTestResultModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *LastTestResultModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *LastTestResultModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


