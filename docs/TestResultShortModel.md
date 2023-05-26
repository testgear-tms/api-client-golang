# TestResultShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **NullableString** |  | [optional] 
**Traces** | Pointer to **NullableString** |  | [optional] 
**FailureType** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**TestPoint** | Pointer to [**TestPointPutModel**](TestPointPutModel.md) |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**AutoTest** | Pointer to [**AutoTestShortModel**](AutoTestShortModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 

## Methods

### NewTestResultShortModel

`func NewTestResultShortModel() *TestResultShortModel`

NewTestResultShortModel instantiates a new TestResultShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultShortModelWithDefaults

`func NewTestResultShortModelWithDefaults() *TestResultShortModel`

NewTestResultShortModelWithDefaults instantiates a new TestResultShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultShortModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestResultShortModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOutcome

`func (o *TestResultShortModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultShortModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultShortModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultShortModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultShortModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultShortModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetTraces

`func (o *TestResultShortModel) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *TestResultShortModel) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *TestResultShortModel) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *TestResultShortModel) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *TestResultShortModel) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *TestResultShortModel) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetFailureType

`func (o *TestResultShortModel) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *TestResultShortModel) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *TestResultShortModel) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *TestResultShortModel) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### SetFailureTypeNil

`func (o *TestResultShortModel) SetFailureTypeNil(b bool)`

 SetFailureTypeNil sets the value for FailureType to be an explicit nil

### UnsetFailureType
`func (o *TestResultShortModel) UnsetFailureType()`

UnsetFailureType ensures that no value is present for FailureType, not even an explicit nil
### GetMessage

`func (o *TestResultShortModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultShortModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultShortModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultShortModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultShortModel) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultShortModel) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTestPoint

`func (o *TestResultShortModel) GetTestPoint() TestPointPutModel`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *TestResultShortModel) GetTestPointOk() (*TestPointPutModel, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *TestResultShortModel) SetTestPoint(v TestPointPutModel)`

SetTestPoint sets TestPoint field to given value.

### HasTestPoint

`func (o *TestResultShortModel) HasTestPoint() bool`

HasTestPoint returns a boolean if a field has been set.

### GetCreatedDate

`func (o *TestResultShortModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultShortModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultShortModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestResultShortModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestResultShortModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestResultShortModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetAutoTest

`func (o *TestResultShortModel) GetAutoTest() AutoTestShortModel`

GetAutoTest returns the AutoTest field if non-nil, zero value otherwise.

### GetAutoTestOk

`func (o *TestResultShortModel) GetAutoTestOk() (*AutoTestShortModel, bool)`

GetAutoTestOk returns a tuple with the AutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTest

`func (o *TestResultShortModel) SetAutoTest(v AutoTestShortModel)`

SetAutoTest sets AutoTest field to given value.

### HasAutoTest

`func (o *TestResultShortModel) HasAutoTest() bool`

HasAutoTest returns a boolean if a field has been set.

### GetAttachments

`func (o *TestResultShortModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultShortModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultShortModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultShortModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultShortModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultShortModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


