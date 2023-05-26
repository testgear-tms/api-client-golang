# TestRunTestResultsPartialBulkSetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to [**TestRunTestResultsSelectModel**](TestRunTestResultsSelectModel.md) |  | [optional] 
**ResultReasonIds** | Pointer to **[]string** | Unique IDs of result reasons to be assigned to test results | [optional] 
**Links** | Pointer to [**[]LinkPostModel**](LinkPostModel.md) | Collection of links to be assigned to test results | [optional] 
**Comment** | Pointer to **NullableString** | Comment to be added to test results | [optional] 
**AttachmentIds** | Pointer to **[]string** | Unique IDs of files to be attached to test results | [optional] 

## Methods

### NewTestRunTestResultsPartialBulkSetModel

`func NewTestRunTestResultsPartialBulkSetModel() *TestRunTestResultsPartialBulkSetModel`

NewTestRunTestResultsPartialBulkSetModel instantiates a new TestRunTestResultsPartialBulkSetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunTestResultsPartialBulkSetModelWithDefaults

`func NewTestRunTestResultsPartialBulkSetModelWithDefaults() *TestRunTestResultsPartialBulkSetModel`

NewTestRunTestResultsPartialBulkSetModelWithDefaults instantiates a new TestRunTestResultsPartialBulkSetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *TestRunTestResultsPartialBulkSetModel) GetSelector() TestRunTestResultsSelectModel`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *TestRunTestResultsPartialBulkSetModel) GetSelectorOk() (*TestRunTestResultsSelectModel, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *TestRunTestResultsPartialBulkSetModel) SetSelector(v TestRunTestResultsSelectModel)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *TestRunTestResultsPartialBulkSetModel) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetResultReasonIds

`func (o *TestRunTestResultsPartialBulkSetModel) GetResultReasonIds() []string`

GetResultReasonIds returns the ResultReasonIds field if non-nil, zero value otherwise.

### GetResultReasonIdsOk

`func (o *TestRunTestResultsPartialBulkSetModel) GetResultReasonIdsOk() (*[]string, bool)`

GetResultReasonIdsOk returns a tuple with the ResultReasonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReasonIds

`func (o *TestRunTestResultsPartialBulkSetModel) SetResultReasonIds(v []string)`

SetResultReasonIds sets ResultReasonIds field to given value.

### HasResultReasonIds

`func (o *TestRunTestResultsPartialBulkSetModel) HasResultReasonIds() bool`

HasResultReasonIds returns a boolean if a field has been set.

### SetResultReasonIdsNil

`func (o *TestRunTestResultsPartialBulkSetModel) SetResultReasonIdsNil(b bool)`

 SetResultReasonIdsNil sets the value for ResultReasonIds to be an explicit nil

### UnsetResultReasonIds
`func (o *TestRunTestResultsPartialBulkSetModel) UnsetResultReasonIds()`

UnsetResultReasonIds ensures that no value is present for ResultReasonIds, not even an explicit nil
### GetLinks

`func (o *TestRunTestResultsPartialBulkSetModel) GetLinks() []LinkPostModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestRunTestResultsPartialBulkSetModel) GetLinksOk() (*[]LinkPostModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestRunTestResultsPartialBulkSetModel) SetLinks(v []LinkPostModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestRunTestResultsPartialBulkSetModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestRunTestResultsPartialBulkSetModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestRunTestResultsPartialBulkSetModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetComment

`func (o *TestRunTestResultsPartialBulkSetModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestRunTestResultsPartialBulkSetModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestRunTestResultsPartialBulkSetModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestRunTestResultsPartialBulkSetModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestRunTestResultsPartialBulkSetModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestRunTestResultsPartialBulkSetModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetAttachmentIds

`func (o *TestRunTestResultsPartialBulkSetModel) GetAttachmentIds() []string`

GetAttachmentIds returns the AttachmentIds field if non-nil, zero value otherwise.

### GetAttachmentIdsOk

`func (o *TestRunTestResultsPartialBulkSetModel) GetAttachmentIdsOk() (*[]string, bool)`

GetAttachmentIdsOk returns a tuple with the AttachmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentIds

`func (o *TestRunTestResultsPartialBulkSetModel) SetAttachmentIds(v []string)`

SetAttachmentIds sets AttachmentIds field to given value.

### HasAttachmentIds

`func (o *TestRunTestResultsPartialBulkSetModel) HasAttachmentIds() bool`

HasAttachmentIds returns a boolean if a field has been set.

### SetAttachmentIdsNil

`func (o *TestRunTestResultsPartialBulkSetModel) SetAttachmentIdsNil(b bool)`

 SetAttachmentIdsNil sets the value for AttachmentIds to be an explicit nil

### UnsetAttachmentIds
`func (o *TestRunTestResultsPartialBulkSetModel) UnsetAttachmentIds()`

UnsetAttachmentIds ensures that no value is present for AttachmentIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


