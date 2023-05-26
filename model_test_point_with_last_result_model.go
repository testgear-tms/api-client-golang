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

// checks if the TestPointWithLastResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointWithLastResultModel{}

// TestPointWithLastResultModel struct for TestPointWithLastResultModel
type TestPointWithLastResultModel struct {
	Id *string `json:"id,omitempty"`
	WorkItemName NullableString `json:"workItemName,omitempty"`
	IsAutomated *bool `json:"isAutomated,omitempty"`
	TesterId NullableString `json:"testerId,omitempty"`
	WorkItemId *string `json:"workItemId,omitempty"`
	ConfigurationId NullableString `json:"configurationId,omitempty"`
	TestSuiteId *string `json:"testSuiteId,omitempty"`
	LastTestResult *LastTestResultModel `json:"lastTestResult,omitempty"`
	Status NullableString `json:"status,omitempty"`
	WorkItemGlobalId NullableInt64 `json:"workItemGlobalId,omitempty"`
	WorkItemEntityTypeName NullableString `json:"workItemEntityTypeName,omitempty"`
	SectionId *string `json:"sectionId,omitempty"`
	SectionName NullableString `json:"sectionName,omitempty"`
	CreatedDate NullableTime `json:"createdDate,omitempty"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	TagNames []string `json:"tagNames,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	Priority WorkItemPriorityModel `json:"priority"`
	TestSuiteNameBreadCrumbs []string `json:"testSuiteNameBreadCrumbs,omitempty"`
	GroupCount NullableInt32 `json:"groupCount,omitempty"`
	Iteration *IterationModel `json:"iteration,omitempty"`
}

// NewTestPointWithLastResultModel instantiates a new TestPointWithLastResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointWithLastResultModel(priority WorkItemPriorityModel) *TestPointWithLastResultModel {
	this := TestPointWithLastResultModel{}
	this.Priority = priority
	return &this
}

// NewTestPointWithLastResultModelWithDefaults instantiates a new TestPointWithLastResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointWithLastResultModelWithDefaults() *TestPointWithLastResultModel {
	this := TestPointWithLastResultModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestPointWithLastResultModel) SetId(v string) {
	o.Id = &v
}

// GetWorkItemName returns the WorkItemName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetWorkItemName() string {
	if o == nil || IsNil(o.WorkItemName.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemName.Get()
}

// GetWorkItemNameOk returns a tuple with the WorkItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetWorkItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemName.Get(), o.WorkItemName.IsSet()
}

// HasWorkItemName returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasWorkItemName() bool {
	if o != nil && o.WorkItemName.IsSet() {
		return true
	}

	return false
}

// SetWorkItemName gets a reference to the given NullableString and assigns it to the WorkItemName field.
func (o *TestPointWithLastResultModel) SetWorkItemName(v string) {
	o.WorkItemName.Set(&v)
}
// SetWorkItemNameNil sets the value for WorkItemName to be an explicit nil
func (o *TestPointWithLastResultModel) SetWorkItemNameNil() {
	o.WorkItemName.Set(nil)
}

// UnsetWorkItemName ensures that no value is present for WorkItemName, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetWorkItemName() {
	o.WorkItemName.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated) {
		var ret bool
		return ret
	}
	return *o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomated) {
		return nil, false
	}
	return o.IsAutomated, true
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasIsAutomated() bool {
	if o != nil && !IsNil(o.IsAutomated) {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given bool and assigns it to the IsAutomated field.
func (o *TestPointWithLastResultModel) SetIsAutomated(v bool) {
	o.IsAutomated = &v
}

// GetTesterId returns the TesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetTesterId() string {
	if o == nil || IsNil(o.TesterId.Get()) {
		var ret string
		return ret
	}
	return *o.TesterId.Get()
}

// GetTesterIdOk returns a tuple with the TesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetTesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TesterId.Get(), o.TesterId.IsSet()
}

// HasTesterId returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasTesterId() bool {
	if o != nil && o.TesterId.IsSet() {
		return true
	}

	return false
}

// SetTesterId gets a reference to the given NullableString and assigns it to the TesterId field.
func (o *TestPointWithLastResultModel) SetTesterId(v string) {
	o.TesterId.Set(&v)
}
// SetTesterIdNil sets the value for TesterId to be an explicit nil
func (o *TestPointWithLastResultModel) SetTesterIdNil() {
	o.TesterId.Set(nil)
}

// UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetTesterId() {
	o.TesterId.Unset()
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId) {
		var ret string
		return ret
	}
	return *o.WorkItemId
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkItemId) {
		return nil, false
	}
	return o.WorkItemId, true
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasWorkItemId() bool {
	if o != nil && !IsNil(o.WorkItemId) {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given string and assigns it to the WorkItemId field.
func (o *TestPointWithLastResultModel) SetWorkItemId(v string) {
	o.WorkItemId = &v
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigurationId.Get()
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationId.Get(), o.ConfigurationId.IsSet()
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given NullableString and assigns it to the ConfigurationId field.
func (o *TestPointWithLastResultModel) SetConfigurationId(v string) {
	o.ConfigurationId.Set(&v)
}
// SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil
func (o *TestPointWithLastResultModel) SetConfigurationIdNil() {
	o.ConfigurationId.Set(nil)
}

// UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetConfigurationId() {
	o.ConfigurationId.Unset()
}

// GetTestSuiteId returns the TestSuiteId field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetTestSuiteId() string {
	if o == nil || IsNil(o.TestSuiteId) {
		var ret string
		return ret
	}
	return *o.TestSuiteId
}

// GetTestSuiteIdOk returns a tuple with the TestSuiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetTestSuiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestSuiteId) {
		return nil, false
	}
	return o.TestSuiteId, true
}

// HasTestSuiteId returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasTestSuiteId() bool {
	if o != nil && !IsNil(o.TestSuiteId) {
		return true
	}

	return false
}

// SetTestSuiteId gets a reference to the given string and assigns it to the TestSuiteId field.
func (o *TestPointWithLastResultModel) SetTestSuiteId(v string) {
	o.TestSuiteId = &v
}

// GetLastTestResult returns the LastTestResult field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetLastTestResult() LastTestResultModel {
	if o == nil || IsNil(o.LastTestResult) {
		var ret LastTestResultModel
		return ret
	}
	return *o.LastTestResult
}

// GetLastTestResultOk returns a tuple with the LastTestResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetLastTestResultOk() (*LastTestResultModel, bool) {
	if o == nil || IsNil(o.LastTestResult) {
		return nil, false
	}
	return o.LastTestResult, true
}

// HasLastTestResult returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasLastTestResult() bool {
	if o != nil && !IsNil(o.LastTestResult) {
		return true
	}

	return false
}

// SetLastTestResult gets a reference to the given LastTestResultModel and assigns it to the LastTestResult field.
func (o *TestPointWithLastResultModel) SetLastTestResult(v LastTestResultModel) {
	o.LastTestResult = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TestPointWithLastResultModel) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TestPointWithLastResultModel) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetStatus() {
	o.Status.Unset()
}

// GetWorkItemGlobalId returns the WorkItemGlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetWorkItemGlobalId() int64 {
	if o == nil || IsNil(o.WorkItemGlobalId.Get()) {
		var ret int64
		return ret
	}
	return *o.WorkItemGlobalId.Get()
}

// GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetWorkItemGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemGlobalId.Get(), o.WorkItemGlobalId.IsSet()
}

// HasWorkItemGlobalId returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasWorkItemGlobalId() bool {
	if o != nil && o.WorkItemGlobalId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemGlobalId gets a reference to the given NullableInt64 and assigns it to the WorkItemGlobalId field.
func (o *TestPointWithLastResultModel) SetWorkItemGlobalId(v int64) {
	o.WorkItemGlobalId.Set(&v)
}
// SetWorkItemGlobalIdNil sets the value for WorkItemGlobalId to be an explicit nil
func (o *TestPointWithLastResultModel) SetWorkItemGlobalIdNil() {
	o.WorkItemGlobalId.Set(nil)
}

// UnsetWorkItemGlobalId ensures that no value is present for WorkItemGlobalId, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetWorkItemGlobalId() {
	o.WorkItemGlobalId.Unset()
}

// GetWorkItemEntityTypeName returns the WorkItemEntityTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetWorkItemEntityTypeName() string {
	if o == nil || IsNil(o.WorkItemEntityTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemEntityTypeName.Get()
}

// GetWorkItemEntityTypeNameOk returns a tuple with the WorkItemEntityTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetWorkItemEntityTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemEntityTypeName.Get(), o.WorkItemEntityTypeName.IsSet()
}

// HasWorkItemEntityTypeName returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasWorkItemEntityTypeName() bool {
	if o != nil && o.WorkItemEntityTypeName.IsSet() {
		return true
	}

	return false
}

// SetWorkItemEntityTypeName gets a reference to the given NullableString and assigns it to the WorkItemEntityTypeName field.
func (o *TestPointWithLastResultModel) SetWorkItemEntityTypeName(v string) {
	o.WorkItemEntityTypeName.Set(&v)
}
// SetWorkItemEntityTypeNameNil sets the value for WorkItemEntityTypeName to be an explicit nil
func (o *TestPointWithLastResultModel) SetWorkItemEntityTypeNameNil() {
	o.WorkItemEntityTypeName.Set(nil)
}

// UnsetWorkItemEntityTypeName ensures that no value is present for WorkItemEntityTypeName, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetWorkItemEntityTypeName() {
	o.WorkItemEntityTypeName.Unset()
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetSectionId() string {
	if o == nil || IsNil(o.SectionId) {
		var ret string
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetSectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SectionId) {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasSectionId() bool {
	if o != nil && !IsNil(o.SectionId) {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given string and assigns it to the SectionId field.
func (o *TestPointWithLastResultModel) SetSectionId(v string) {
	o.SectionId = &v
}

// GetSectionName returns the SectionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetSectionName() string {
	if o == nil || IsNil(o.SectionName.Get()) {
		var ret string
		return ret
	}
	return *o.SectionName.Get()
}

// GetSectionNameOk returns a tuple with the SectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetSectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionName.Get(), o.SectionName.IsSet()
}

// HasSectionName returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasSectionName() bool {
	if o != nil && o.SectionName.IsSet() {
		return true
	}

	return false
}

// SetSectionName gets a reference to the given NullableString and assigns it to the SectionName field.
func (o *TestPointWithLastResultModel) SetSectionName(v string) {
	o.SectionName.Set(&v)
}
// SetSectionNameNil sets the value for SectionName to be an explicit nil
func (o *TestPointWithLastResultModel) SetSectionNameNil() {
	o.SectionName.Set(nil)
}

// UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetSectionName() {
	o.SectionName.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *TestPointWithLastResultModel) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *TestPointWithLastResultModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *TestPointWithLastResultModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *TestPointWithLastResultModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *TestPointWithLastResultModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *TestPointWithLastResultModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *TestPointWithLastResultModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *TestPointWithLastResultModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTagNames returns the TagNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetTagNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetTagNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.TagNames) {
		return nil, false
	}
	return o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasTagNames() bool {
	if o != nil && IsNil(o.TagNames) {
		return true
	}

	return false
}

// SetTagNames gets a reference to the given []string and assigns it to the TagNames field.
func (o *TestPointWithLastResultModel) SetTagNames(v []string) {
	o.TagNames = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *TestPointWithLastResultModel) SetDuration(v int32) {
	o.Duration = &v
}

// GetPriority returns the Priority field value
func (o *TestPointWithLastResultModel) GetPriority() WorkItemPriorityModel {
	if o == nil {
		var ret WorkItemPriorityModel
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetPriorityOk() (*WorkItemPriorityModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *TestPointWithLastResultModel) SetPriority(v WorkItemPriorityModel) {
	o.Priority = v
}

// GetTestSuiteNameBreadCrumbs returns the TestSuiteNameBreadCrumbs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetTestSuiteNameBreadCrumbs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestSuiteNameBreadCrumbs
}

// GetTestSuiteNameBreadCrumbsOk returns a tuple with the TestSuiteNameBreadCrumbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetTestSuiteNameBreadCrumbsOk() ([]string, bool) {
	if o == nil || IsNil(o.TestSuiteNameBreadCrumbs) {
		return nil, false
	}
	return o.TestSuiteNameBreadCrumbs, true
}

// HasTestSuiteNameBreadCrumbs returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasTestSuiteNameBreadCrumbs() bool {
	if o != nil && IsNil(o.TestSuiteNameBreadCrumbs) {
		return true
	}

	return false
}

// SetTestSuiteNameBreadCrumbs gets a reference to the given []string and assigns it to the TestSuiteNameBreadCrumbs field.
func (o *TestPointWithLastResultModel) SetTestSuiteNameBreadCrumbs(v []string) {
	o.TestSuiteNameBreadCrumbs = v
}

// GetGroupCount returns the GroupCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointWithLastResultModel) GetGroupCount() int32 {
	if o == nil || IsNil(o.GroupCount.Get()) {
		var ret int32
		return ret
	}
	return *o.GroupCount.Get()
}

// GetGroupCountOk returns a tuple with the GroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointWithLastResultModel) GetGroupCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupCount.Get(), o.GroupCount.IsSet()
}

// HasGroupCount returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasGroupCount() bool {
	if o != nil && o.GroupCount.IsSet() {
		return true
	}

	return false
}

// SetGroupCount gets a reference to the given NullableInt32 and assigns it to the GroupCount field.
func (o *TestPointWithLastResultModel) SetGroupCount(v int32) {
	o.GroupCount.Set(&v)
}
// SetGroupCountNil sets the value for GroupCount to be an explicit nil
func (o *TestPointWithLastResultModel) SetGroupCountNil() {
	o.GroupCount.Set(nil)
}

// UnsetGroupCount ensures that no value is present for GroupCount, not even an explicit nil
func (o *TestPointWithLastResultModel) UnsetGroupCount() {
	o.GroupCount.Unset()
}

// GetIteration returns the Iteration field value if set, zero value otherwise.
func (o *TestPointWithLastResultModel) GetIteration() IterationModel {
	if o == nil || IsNil(o.Iteration) {
		var ret IterationModel
		return ret
	}
	return *o.Iteration
}

// GetIterationOk returns a tuple with the Iteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointWithLastResultModel) GetIterationOk() (*IterationModel, bool) {
	if o == nil || IsNil(o.Iteration) {
		return nil, false
	}
	return o.Iteration, true
}

// HasIteration returns a boolean if a field has been set.
func (o *TestPointWithLastResultModel) HasIteration() bool {
	if o != nil && !IsNil(o.Iteration) {
		return true
	}

	return false
}

// SetIteration gets a reference to the given IterationModel and assigns it to the Iteration field.
func (o *TestPointWithLastResultModel) SetIteration(v IterationModel) {
	o.Iteration = &v
}

func (o TestPointWithLastResultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointWithLastResultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.WorkItemName.IsSet() {
		toSerialize["workItemName"] = o.WorkItemName.Get()
	}
	if !IsNil(o.IsAutomated) {
		toSerialize["isAutomated"] = o.IsAutomated
	}
	if o.TesterId.IsSet() {
		toSerialize["testerId"] = o.TesterId.Get()
	}
	if !IsNil(o.WorkItemId) {
		toSerialize["workItemId"] = o.WorkItemId
	}
	if o.ConfigurationId.IsSet() {
		toSerialize["configurationId"] = o.ConfigurationId.Get()
	}
	if !IsNil(o.TestSuiteId) {
		toSerialize["testSuiteId"] = o.TestSuiteId
	}
	if !IsNil(o.LastTestResult) {
		toSerialize["lastTestResult"] = o.LastTestResult
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.WorkItemGlobalId.IsSet() {
		toSerialize["workItemGlobalId"] = o.WorkItemGlobalId.Get()
	}
	if o.WorkItemEntityTypeName.IsSet() {
		toSerialize["workItemEntityTypeName"] = o.WorkItemEntityTypeName.Get()
	}
	if !IsNil(o.SectionId) {
		toSerialize["sectionId"] = o.SectionId
	}
	if o.SectionName.IsSet() {
		toSerialize["sectionName"] = o.SectionName.Get()
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.TagNames != nil {
		toSerialize["tagNames"] = o.TagNames
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["priority"] = o.Priority
	if o.TestSuiteNameBreadCrumbs != nil {
		toSerialize["testSuiteNameBreadCrumbs"] = o.TestSuiteNameBreadCrumbs
	}
	if o.GroupCount.IsSet() {
		toSerialize["groupCount"] = o.GroupCount.Get()
	}
	if !IsNil(o.Iteration) {
		toSerialize["iteration"] = o.Iteration
	}
	return toSerialize, nil
}

type NullableTestPointWithLastResultModel struct {
	value *TestPointWithLastResultModel
	isSet bool
}

func (v NullableTestPointWithLastResultModel) Get() *TestPointWithLastResultModel {
	return v.value
}

func (v *NullableTestPointWithLastResultModel) Set(val *TestPointWithLastResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointWithLastResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointWithLastResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointWithLastResultModel(val *TestPointWithLastResultModel) *NullableTestPointWithLastResultModel {
	return &NullableTestPointWithLastResultModel{value: val, isSet: true}
}

func (v NullableTestPointWithLastResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointWithLastResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

