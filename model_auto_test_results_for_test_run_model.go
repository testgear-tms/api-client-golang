/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"time"
)

// checks if the AutoTestResultsForTestRunModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestResultsForTestRunModel{}

// AutoTestResultsForTestRunModel struct for AutoTestResultsForTestRunModel
type AutoTestResultsForTestRunModel struct {
	// Specifies the GUID of the autotest configuration, which was specified when the test run was created.
	ConfigurationId string `json:"configurationId"`
	// Specifies the links in the autotest.
	Links []LinkPostModel `json:"links,omitempty"`
	// Specifies the cause of autotest failure.
	FailureReasonNames []FailureCategoryModel `json:"failureReasonNames,omitempty"`
	// Specifies the external ID of the autotest, which was specified when the test run was created.
	AutoTestExternalId string `json:"autoTestExternalId"`
	Outcome AvailableTestResultOutcome `json:"outcome"`
	// A comment for the result.
	Message NullableString `json:"message,omitempty"`
	// An extended comment or a stack trace.
	Traces NullableString `json:"traces,omitempty"`
	// Test run start date.
	StartedOn NullableTime `json:"startedOn,omitempty"`
	// Test run end date.
	CompletedOn NullableTime `json:"completedOn,omitempty"`
	// Expected or actual duration of the test run execution in milliseconds.
	Duration NullableInt64 `json:"duration,omitempty"`
	// Specifies an attachment GUID. Multiple values can be sent.
	Attachments []AttachmentPutModel `json:"attachments,omitempty"`
	// \"<b>parameter</b>\": \"<b>value</b>\" pair with arbitrary custom parameters. Multiple parameters can be sent.
	Parameters map[string]string `json:"parameters,omitempty"`
	// \"<b>property</b>\": \"<b>value</b>\" pair with arbitrary custom properties. Multiple properties can be sent.
	Properties map[string]string `json:"properties,omitempty"`
	// Specifies the results of individual steps.
	StepResults []AttachmentPutModelAutoTestStepResultsModel `json:"stepResults,omitempty"`
	// Specifies the results of setup steps. For information on supported values, see the `stepResults` parameter above.
	SetupResults []AttachmentPutModelAutoTestStepResultsModel `json:"setupResults,omitempty"`
	// Specifies the results of the teardown steps. For information on supported values, see the `stepResults` parameter above.
	TeardownResults []AttachmentPutModelAutoTestStepResultsModel `json:"teardownResults,omitempty"`
}

// NewAutoTestResultsForTestRunModel instantiates a new AutoTestResultsForTestRunModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestResultsForTestRunModel(configurationId string, autoTestExternalId string, outcome AvailableTestResultOutcome) *AutoTestResultsForTestRunModel {
	this := AutoTestResultsForTestRunModel{}
	this.ConfigurationId = configurationId
	this.AutoTestExternalId = autoTestExternalId
	this.Outcome = outcome
	return &this
}

// NewAutoTestResultsForTestRunModelWithDefaults instantiates a new AutoTestResultsForTestRunModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestResultsForTestRunModelWithDefaults() *AutoTestResultsForTestRunModel {
	this := AutoTestResultsForTestRunModel{}
	return &this
}

// GetConfigurationId returns the ConfigurationId field value
func (o *AutoTestResultsForTestRunModel) GetConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value
// and a boolean to check if the value has been set.
func (o *AutoTestResultsForTestRunModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationId, true
}

// SetConfigurationId sets field value
func (o *AutoTestResultsForTestRunModel) SetConfigurationId(v string) {
	o.ConfigurationId = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetLinks() []LinkPostModel {
	if o == nil {
		var ret []LinkPostModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetLinksOk() ([]LinkPostModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkPostModel and assigns it to the Links field.
func (o *AutoTestResultsForTestRunModel) SetLinks(v []LinkPostModel) {
	o.Links = v
}

// GetFailureReasonNames returns the FailureReasonNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetFailureReasonNames() []FailureCategoryModel {
	if o == nil {
		var ret []FailureCategoryModel
		return ret
	}
	return o.FailureReasonNames
}

// GetFailureReasonNamesOk returns a tuple with the FailureReasonNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetFailureReasonNamesOk() ([]FailureCategoryModel, bool) {
	if o == nil || IsNil(o.FailureReasonNames) {
		return nil, false
	}
	return o.FailureReasonNames, true
}

// HasFailureReasonNames returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasFailureReasonNames() bool {
	if o != nil && IsNil(o.FailureReasonNames) {
		return true
	}

	return false
}

// SetFailureReasonNames gets a reference to the given []FailureCategoryModel and assigns it to the FailureReasonNames field.
func (o *AutoTestResultsForTestRunModel) SetFailureReasonNames(v []FailureCategoryModel) {
	o.FailureReasonNames = v
}

// GetAutoTestExternalId returns the AutoTestExternalId field value
func (o *AutoTestResultsForTestRunModel) GetAutoTestExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoTestExternalId
}

// GetAutoTestExternalIdOk returns a tuple with the AutoTestExternalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestResultsForTestRunModel) GetAutoTestExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoTestExternalId, true
}

// SetAutoTestExternalId sets field value
func (o *AutoTestResultsForTestRunModel) SetAutoTestExternalId(v string) {
	o.AutoTestExternalId = v
}

// GetOutcome returns the Outcome field value
func (o *AutoTestResultsForTestRunModel) GetOutcome() AvailableTestResultOutcome {
	if o == nil {
		var ret AvailableTestResultOutcome
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *AutoTestResultsForTestRunModel) GetOutcomeOk() (*AvailableTestResultOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *AutoTestResultsForTestRunModel) SetOutcome(v AvailableTestResultOutcome) {
	o.Outcome = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *AutoTestResultsForTestRunModel) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *AutoTestResultsForTestRunModel) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *AutoTestResultsForTestRunModel) UnsetMessage() {
	o.Message.Unset()
}

// GetTraces returns the Traces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetTraces() string {
	if o == nil || IsNil(o.Traces.Get()) {
		var ret string
		return ret
	}
	return *o.Traces.Get()
}

// GetTracesOk returns a tuple with the Traces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetTracesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traces.Get(), o.Traces.IsSet()
}

// HasTraces returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasTraces() bool {
	if o != nil && o.Traces.IsSet() {
		return true
	}

	return false
}

// SetTraces gets a reference to the given NullableString and assigns it to the Traces field.
func (o *AutoTestResultsForTestRunModel) SetTraces(v string) {
	o.Traces.Set(&v)
}
// SetTracesNil sets the value for Traces to be an explicit nil
func (o *AutoTestResultsForTestRunModel) SetTracesNil() {
	o.Traces.Set(nil)
}

// UnsetTraces ensures that no value is present for Traces, not even an explicit nil
func (o *AutoTestResultsForTestRunModel) UnsetTraces() {
	o.Traces.Unset()
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *AutoTestResultsForTestRunModel) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *AutoTestResultsForTestRunModel) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *AutoTestResultsForTestRunModel) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *AutoTestResultsForTestRunModel) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *AutoTestResultsForTestRunModel) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *AutoTestResultsForTestRunModel) UnsetCompletedOn() {
	o.CompletedOn.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetDuration() int64 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *AutoTestResultsForTestRunModel) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *AutoTestResultsForTestRunModel) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *AutoTestResultsForTestRunModel) UnsetDuration() {
	o.Duration.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetAttachments() []AttachmentPutModel {
	if o == nil {
		var ret []AttachmentPutModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetAttachmentsOk() ([]AttachmentPutModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentPutModel and assigns it to the Attachments field.
func (o *AutoTestResultsForTestRunModel) SetAttachments(v []AttachmentPutModel) {
	o.Attachments = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *AutoTestResultsForTestRunModel) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasProperties() bool {
	if o != nil && IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *AutoTestResultsForTestRunModel) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetStepResults returns the StepResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetStepResults() []AttachmentPutModelAutoTestStepResultsModel {
	if o == nil {
		var ret []AttachmentPutModelAutoTestStepResultsModel
		return ret
	}
	return o.StepResults
}

// GetStepResultsOk returns a tuple with the StepResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetStepResultsOk() ([]AttachmentPutModelAutoTestStepResultsModel, bool) {
	if o == nil || IsNil(o.StepResults) {
		return nil, false
	}
	return o.StepResults, true
}

// HasStepResults returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasStepResults() bool {
	if o != nil && IsNil(o.StepResults) {
		return true
	}

	return false
}

// SetStepResults gets a reference to the given []AttachmentPutModelAutoTestStepResultsModel and assigns it to the StepResults field.
func (o *AutoTestResultsForTestRunModel) SetStepResults(v []AttachmentPutModelAutoTestStepResultsModel) {
	o.StepResults = v
}

// GetSetupResults returns the SetupResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetSetupResults() []AttachmentPutModelAutoTestStepResultsModel {
	if o == nil {
		var ret []AttachmentPutModelAutoTestStepResultsModel
		return ret
	}
	return o.SetupResults
}

// GetSetupResultsOk returns a tuple with the SetupResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetSetupResultsOk() ([]AttachmentPutModelAutoTestStepResultsModel, bool) {
	if o == nil || IsNil(o.SetupResults) {
		return nil, false
	}
	return o.SetupResults, true
}

// HasSetupResults returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasSetupResults() bool {
	if o != nil && IsNil(o.SetupResults) {
		return true
	}

	return false
}

// SetSetupResults gets a reference to the given []AttachmentPutModelAutoTestStepResultsModel and assigns it to the SetupResults field.
func (o *AutoTestResultsForTestRunModel) SetSetupResults(v []AttachmentPutModelAutoTestStepResultsModel) {
	o.SetupResults = v
}

// GetTeardownResults returns the TeardownResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestResultsForTestRunModel) GetTeardownResults() []AttachmentPutModelAutoTestStepResultsModel {
	if o == nil {
		var ret []AttachmentPutModelAutoTestStepResultsModel
		return ret
	}
	return o.TeardownResults
}

// GetTeardownResultsOk returns a tuple with the TeardownResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestResultsForTestRunModel) GetTeardownResultsOk() ([]AttachmentPutModelAutoTestStepResultsModel, bool) {
	if o == nil || IsNil(o.TeardownResults) {
		return nil, false
	}
	return o.TeardownResults, true
}

// HasTeardownResults returns a boolean if a field has been set.
func (o *AutoTestResultsForTestRunModel) HasTeardownResults() bool {
	if o != nil && IsNil(o.TeardownResults) {
		return true
	}

	return false
}

// SetTeardownResults gets a reference to the given []AttachmentPutModelAutoTestStepResultsModel and assigns it to the TeardownResults field.
func (o *AutoTestResultsForTestRunModel) SetTeardownResults(v []AttachmentPutModelAutoTestStepResultsModel) {
	o.TeardownResults = v
}

func (o AutoTestResultsForTestRunModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestResultsForTestRunModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configurationId"] = o.ConfigurationId
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.FailureReasonNames != nil {
		toSerialize["failureReasonNames"] = o.FailureReasonNames
	}
	toSerialize["autoTestExternalId"] = o.AutoTestExternalId
	toSerialize["outcome"] = o.Outcome
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Traces.IsSet() {
		toSerialize["traces"] = o.Traces.Get()
	}
	if o.StartedOn.IsSet() {
		toSerialize["startedOn"] = o.StartedOn.Get()
	}
	if o.CompletedOn.IsSet() {
		toSerialize["completedOn"] = o.CompletedOn.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.StepResults != nil {
		toSerialize["stepResults"] = o.StepResults
	}
	if o.SetupResults != nil {
		toSerialize["setupResults"] = o.SetupResults
	}
	if o.TeardownResults != nil {
		toSerialize["teardownResults"] = o.TeardownResults
	}
	return toSerialize, nil
}

type NullableAutoTestResultsForTestRunModel struct {
	value *AutoTestResultsForTestRunModel
	isSet bool
}

func (v NullableAutoTestResultsForTestRunModel) Get() *AutoTestResultsForTestRunModel {
	return v.value
}

func (v *NullableAutoTestResultsForTestRunModel) Set(val *AutoTestResultsForTestRunModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestResultsForTestRunModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestResultsForTestRunModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestResultsForTestRunModel(val *AutoTestResultsForTestRunModel) *NullableAutoTestResultsForTestRunModel {
	return &NullableAutoTestResultsForTestRunModel{value: val, isSet: true}
}

func (v NullableAutoTestResultsForTestRunModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestResultsForTestRunModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

