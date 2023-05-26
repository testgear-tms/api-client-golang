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

// checks if the TestPointShortGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointShortGetModel{}

// TestPointShortGetModel struct for TestPointShortGetModel
type TestPointShortGetModel struct {
	// Unique ID of the test point
	Id *string `json:"id,omitempty"`
	// Creation date of the test point
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Unique ID of the test point creator
	CreatedById *string `json:"createdById,omitempty"`
	// Last modification date of the test point
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	// Unique ID of the test point last editor
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	// Unique ID of the test point assigned user
	TesterId NullableString `json:"testerId,omitempty"`
	// Collection of the test point parameters
	Parameters map[string]string `json:"parameters,omitempty"`
	// Collection of attributes of work item the test point represents
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Collection of the test point tags
	Tags []string `json:"tags,omitempty"`
	// Collection of the test point links
	Links []string `json:"links,omitempty"`
	// Unique ID of test suite the test point assigned to
	TestSuiteId *string `json:"testSuiteId,omitempty"`
	// Unique ID of work item the test point represents
	WorkItemId *string `json:"workItemId,omitempty"`
	// Global ID of work item the test point represents
	WorkItemGlobalId *int64 `json:"workItemGlobalId,omitempty"`
	// Unique ID of work item version the test point represents
	WorkItemVersionId *string `json:"workItemVersionId,omitempty"`
	Status TestPointStatus `json:"status"`
	Priority WorkItemPriorityModel `json:"priority"`
	// Indicates if the test point represents an autotest
	IsAutomated *bool `json:"isAutomated,omitempty"`
	// Name of the test point
	Name NullableString `json:"name,omitempty"`
	// Unique ID of the test point configuration
	ConfigurationId *string `json:"configurationId,omitempty"`
	// Duration of the test point
	Duration *int32 `json:"duration,omitempty"`
	// Unique ID of section where work item the test point represents is located
	SectionId *string `json:"sectionId,omitempty"`
	// Name of section where work item the test point represents is located
	SectionName NullableString `json:"sectionName,omitempty"`
	// Unique ID of the test point project
	ProjectId *string `json:"projectId,omitempty"`
	LastTestResult LastTestResultModel `json:"lastTestResult"`
	// Unique ID of work item iteration the test point represents
	IterationId *string `json:"iterationId,omitempty"`
}

// NewTestPointShortGetModel instantiates a new TestPointShortGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointShortGetModel(status TestPointStatus, priority WorkItemPriorityModel, lastTestResult LastTestResultModel) *TestPointShortGetModel {
	this := TestPointShortGetModel{}
	this.Status = status
	this.Priority = priority
	this.LastTestResult = lastTestResult
	return &this
}

// NewTestPointShortGetModelWithDefaults instantiates a new TestPointShortGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointShortGetModelWithDefaults() *TestPointShortGetModel {
	this := TestPointShortGetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestPointShortGetModel) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *TestPointShortGetModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *TestPointShortGetModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *TestPointShortGetModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *TestPointShortGetModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *TestPointShortGetModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *TestPointShortGetModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *TestPointShortGetModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *TestPointShortGetModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetTesterId returns the TesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetTesterId() string {
	if o == nil || IsNil(o.TesterId.Get()) {
		var ret string
		return ret
	}
	return *o.TesterId.Get()
}

// GetTesterIdOk returns a tuple with the TesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetTesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TesterId.Get(), o.TesterId.IsSet()
}

// HasTesterId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasTesterId() bool {
	if o != nil && o.TesterId.IsSet() {
		return true
	}

	return false
}

// SetTesterId gets a reference to the given NullableString and assigns it to the TesterId field.
func (o *TestPointShortGetModel) SetTesterId(v string) {
	o.TesterId.Set(&v)
}
// SetTesterIdNil sets the value for TesterId to be an explicit nil
func (o *TestPointShortGetModel) SetTesterIdNil() {
	o.TesterId.Set(nil)
}

// UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
func (o *TestPointShortGetModel) UnsetTesterId() {
	o.TesterId.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *TestPointShortGetModel) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *TestPointShortGetModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TestPointShortGetModel) SetTags(v []string) {
	o.Tags = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetLinks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []string and assigns it to the Links field.
func (o *TestPointShortGetModel) SetLinks(v []string) {
	o.Links = v
}

// GetTestSuiteId returns the TestSuiteId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetTestSuiteId() string {
	if o == nil || IsNil(o.TestSuiteId) {
		var ret string
		return ret
	}
	return *o.TestSuiteId
}

// GetTestSuiteIdOk returns a tuple with the TestSuiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetTestSuiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestSuiteId) {
		return nil, false
	}
	return o.TestSuiteId, true
}

// HasTestSuiteId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasTestSuiteId() bool {
	if o != nil && !IsNil(o.TestSuiteId) {
		return true
	}

	return false
}

// SetTestSuiteId gets a reference to the given string and assigns it to the TestSuiteId field.
func (o *TestPointShortGetModel) SetTestSuiteId(v string) {
	o.TestSuiteId = &v
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId) {
		var ret string
		return ret
	}
	return *o.WorkItemId
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkItemId) {
		return nil, false
	}
	return o.WorkItemId, true
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasWorkItemId() bool {
	if o != nil && !IsNil(o.WorkItemId) {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given string and assigns it to the WorkItemId field.
func (o *TestPointShortGetModel) SetWorkItemId(v string) {
	o.WorkItemId = &v
}

// GetWorkItemGlobalId returns the WorkItemGlobalId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetWorkItemGlobalId() int64 {
	if o == nil || IsNil(o.WorkItemGlobalId) {
		var ret int64
		return ret
	}
	return *o.WorkItemGlobalId
}

// GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetWorkItemGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WorkItemGlobalId) {
		return nil, false
	}
	return o.WorkItemGlobalId, true
}

// HasWorkItemGlobalId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasWorkItemGlobalId() bool {
	if o != nil && !IsNil(o.WorkItemGlobalId) {
		return true
	}

	return false
}

// SetWorkItemGlobalId gets a reference to the given int64 and assigns it to the WorkItemGlobalId field.
func (o *TestPointShortGetModel) SetWorkItemGlobalId(v int64) {
	o.WorkItemGlobalId = &v
}

// GetWorkItemVersionId returns the WorkItemVersionId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetWorkItemVersionId() string {
	if o == nil || IsNil(o.WorkItemVersionId) {
		var ret string
		return ret
	}
	return *o.WorkItemVersionId
}

// GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetWorkItemVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkItemVersionId) {
		return nil, false
	}
	return o.WorkItemVersionId, true
}

// HasWorkItemVersionId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasWorkItemVersionId() bool {
	if o != nil && !IsNil(o.WorkItemVersionId) {
		return true
	}

	return false
}

// SetWorkItemVersionId gets a reference to the given string and assigns it to the WorkItemVersionId field.
func (o *TestPointShortGetModel) SetWorkItemVersionId(v string) {
	o.WorkItemVersionId = &v
}

// GetStatus returns the Status field value
func (o *TestPointShortGetModel) GetStatus() TestPointStatus {
	if o == nil {
		var ret TestPointStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetStatusOk() (*TestPointStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestPointShortGetModel) SetStatus(v TestPointStatus) {
	o.Status = v
}

// GetPriority returns the Priority field value
func (o *TestPointShortGetModel) GetPriority() WorkItemPriorityModel {
	if o == nil {
		var ret WorkItemPriorityModel
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetPriorityOk() (*WorkItemPriorityModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *TestPointShortGetModel) SetPriority(v WorkItemPriorityModel) {
	o.Priority = v
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated) {
		var ret bool
		return ret
	}
	return *o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomated) {
		return nil, false
	}
	return o.IsAutomated, true
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasIsAutomated() bool {
	if o != nil && !IsNil(o.IsAutomated) {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given bool and assigns it to the IsAutomated field.
func (o *TestPointShortGetModel) SetIsAutomated(v bool) {
	o.IsAutomated = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TestPointShortGetModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TestPointShortGetModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TestPointShortGetModel) UnsetName() {
	o.Name.Unset()
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationId) {
		return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasConfigurationId() bool {
	if o != nil && !IsNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *TestPointShortGetModel) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *TestPointShortGetModel) SetDuration(v int32) {
	o.Duration = &v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetSectionId() string {
	if o == nil || IsNil(o.SectionId) {
		var ret string
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetSectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SectionId) {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasSectionId() bool {
	if o != nil && !IsNil(o.SectionId) {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given string and assigns it to the SectionId field.
func (o *TestPointShortGetModel) SetSectionId(v string) {
	o.SectionId = &v
}

// GetSectionName returns the SectionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModel) GetSectionName() string {
	if o == nil || IsNil(o.SectionName.Get()) {
		var ret string
		return ret
	}
	return *o.SectionName.Get()
}

// GetSectionNameOk returns a tuple with the SectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModel) GetSectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionName.Get(), o.SectionName.IsSet()
}

// HasSectionName returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasSectionName() bool {
	if o != nil && o.SectionName.IsSet() {
		return true
	}

	return false
}

// SetSectionName gets a reference to the given NullableString and assigns it to the SectionName field.
func (o *TestPointShortGetModel) SetSectionName(v string) {
	o.SectionName.Set(&v)
}
// SetSectionNameNil sets the value for SectionName to be an explicit nil
func (o *TestPointShortGetModel) SetSectionNameNil() {
	o.SectionName.Set(nil)
}

// UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
func (o *TestPointShortGetModel) UnsetSectionName() {
	o.SectionName.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *TestPointShortGetModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetLastTestResult returns the LastTestResult field value
func (o *TestPointShortGetModel) GetLastTestResult() LastTestResultModel {
	if o == nil {
		var ret LastTestResultModel
		return ret
	}

	return o.LastTestResult
}

// GetLastTestResultOk returns a tuple with the LastTestResult field value
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetLastTestResultOk() (*LastTestResultModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastTestResult, true
}

// SetLastTestResult sets field value
func (o *TestPointShortGetModel) SetLastTestResult(v LastTestResultModel) {
	o.LastTestResult = v
}

// GetIterationId returns the IterationId field value if set, zero value otherwise.
func (o *TestPointShortGetModel) GetIterationId() string {
	if o == nil || IsNil(o.IterationId) {
		var ret string
		return ret
	}
	return *o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModel) GetIterationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IterationId) {
		return nil, false
	}
	return o.IterationId, true
}

// HasIterationId returns a boolean if a field has been set.
func (o *TestPointShortGetModel) HasIterationId() bool {
	if o != nil && !IsNil(o.IterationId) {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given string and assigns it to the IterationId field.
func (o *TestPointShortGetModel) SetIterationId(v string) {
	o.IterationId = &v
}

func (o TestPointShortGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointShortGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.TesterId.IsSet() {
		toSerialize["testerId"] = o.TesterId.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.TestSuiteId) {
		toSerialize["testSuiteId"] = o.TestSuiteId
	}
	if !IsNil(o.WorkItemId) {
		toSerialize["workItemId"] = o.WorkItemId
	}
	if !IsNil(o.WorkItemGlobalId) {
		toSerialize["workItemGlobalId"] = o.WorkItemGlobalId
	}
	if !IsNil(o.WorkItemVersionId) {
		toSerialize["workItemVersionId"] = o.WorkItemVersionId
	}
	toSerialize["status"] = o.Status
	toSerialize["priority"] = o.Priority
	if !IsNil(o.IsAutomated) {
		toSerialize["isAutomated"] = o.IsAutomated
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ConfigurationId) {
		toSerialize["configurationId"] = o.ConfigurationId
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.SectionId) {
		toSerialize["sectionId"] = o.SectionId
	}
	if o.SectionName.IsSet() {
		toSerialize["sectionName"] = o.SectionName.Get()
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	toSerialize["lastTestResult"] = o.LastTestResult
	if !IsNil(o.IterationId) {
		toSerialize["iterationId"] = o.IterationId
	}
	return toSerialize, nil
}

type NullableTestPointShortGetModel struct {
	value *TestPointShortGetModel
	isSet bool
}

func (v NullableTestPointShortGetModel) Get() *TestPointShortGetModel {
	return v.value
}

func (v *NullableTestPointShortGetModel) Set(val *TestPointShortGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointShortGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointShortGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointShortGetModel(val *TestPointShortGetModel) *NullableTestPointShortGetModel {
	return &NullableTestPointShortGetModel{value: val, isSet: true}
}

func (v NullableTestPointShortGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointShortGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

