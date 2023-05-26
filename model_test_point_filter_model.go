/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
)

// checks if the TestPointFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointFilterModel{}

// TestPointFilterModel struct for TestPointFilterModel
type TestPointFilterModel struct {
	// Specifies a test point test plan IDS to search for
	TestPlanIds []string `json:"testPlanIds,omitempty"`
	// Specifies a test point test suite IDs to search for
	TestSuiteIds []string `json:"testSuiteIds,omitempty"`
	// Specifies a test point work item global IDs to search for
	WorkItemGlobalIds []int64 `json:"workItemGlobalIds,omitempty"`
	// Specifies a test point statuses to search for
	Statuses []TestPointStatus `json:"statuses,omitempty"`
	// Specifies a test point priorities to search for
	Priorities []WorkItemPriorityModel `json:"priorities,omitempty"`
	// Specifies a test point automation status to search for
	IsAutomated NullableBool `json:"isAutomated,omitempty"`
	// Specifies a test point name to search for
	Name NullableString `json:"name,omitempty"`
	// Specifies a test point configuration IDs to search for
	ConfigurationIds []string `json:"configurationIds,omitempty"`
	// Specifies a test point assigned user IDs to search for
	TesterIds []string `json:"testerIds,omitempty"`
	Duration *Int64RangeSelectorModel `json:"duration,omitempty"`
	// Specifies a test point work item section IDs to search for
	SectionIds []string `json:"sectionIds,omitempty"`
	CreatedDate *DateTimeRangeSelectorModel `json:"createdDate,omitempty"`
	// Specifies a test point creator IDs to search for
	CreatedByIds []string `json:"createdByIds,omitempty"`
	ModifiedDate *DateTimeRangeSelectorModel `json:"modifiedDate,omitempty"`
	// Specifies a test point last editor IDs to search for
	ModifiedByIds []string `json:"modifiedByIds,omitempty"`
	// Specifies a test point tags to search for
	Tags []string `json:"tags,omitempty"`
	// Specifies a test point attributes to search for
	Attributes map[string][]string `json:"attributes,omitempty"`
}

// NewTestPointFilterModel instantiates a new TestPointFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointFilterModel() *TestPointFilterModel {
	this := TestPointFilterModel{}
	return &this
}

// NewTestPointFilterModelWithDefaults instantiates a new TestPointFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointFilterModelWithDefaults() *TestPointFilterModel {
	this := TestPointFilterModel{}
	return &this
}

// GetTestPlanIds returns the TestPlanIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetTestPlanIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestPlanIds
}

// GetTestPlanIdsOk returns a tuple with the TestPlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetTestPlanIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TestPlanIds) {
		return nil, false
	}
	return o.TestPlanIds, true
}

// HasTestPlanIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasTestPlanIds() bool {
	if o != nil && IsNil(o.TestPlanIds) {
		return true
	}

	return false
}

// SetTestPlanIds gets a reference to the given []string and assigns it to the TestPlanIds field.
func (o *TestPointFilterModel) SetTestPlanIds(v []string) {
	o.TestPlanIds = v
}

// GetTestSuiteIds returns the TestSuiteIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetTestSuiteIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestSuiteIds
}

// GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetTestSuiteIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TestSuiteIds) {
		return nil, false
	}
	return o.TestSuiteIds, true
}

// HasTestSuiteIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasTestSuiteIds() bool {
	if o != nil && IsNil(o.TestSuiteIds) {
		return true
	}

	return false
}

// SetTestSuiteIds gets a reference to the given []string and assigns it to the TestSuiteIds field.
func (o *TestPointFilterModel) SetTestSuiteIds(v []string) {
	o.TestSuiteIds = v
}

// GetWorkItemGlobalIds returns the WorkItemGlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetWorkItemGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.WorkItemGlobalIds
}

// GetWorkItemGlobalIdsOk returns a tuple with the WorkItemGlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetWorkItemGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.WorkItemGlobalIds) {
		return nil, false
	}
	return o.WorkItemGlobalIds, true
}

// HasWorkItemGlobalIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasWorkItemGlobalIds() bool {
	if o != nil && IsNil(o.WorkItemGlobalIds) {
		return true
	}

	return false
}

// SetWorkItemGlobalIds gets a reference to the given []int64 and assigns it to the WorkItemGlobalIds field.
func (o *TestPointFilterModel) SetWorkItemGlobalIds(v []int64) {
	o.WorkItemGlobalIds = v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetStatuses() []TestPointStatus {
	if o == nil {
		var ret []TestPointStatus
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetStatusesOk() ([]TestPointStatus, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasStatuses() bool {
	if o != nil && IsNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []TestPointStatus and assigns it to the Statuses field.
func (o *TestPointFilterModel) SetStatuses(v []TestPointStatus) {
	o.Statuses = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetPriorities() []WorkItemPriorityModel {
	if o == nil {
		var ret []WorkItemPriorityModel
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetPrioritiesOk() ([]WorkItemPriorityModel, bool) {
	if o == nil || IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasPriorities() bool {
	if o != nil && IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []WorkItemPriorityModel and assigns it to the Priorities field.
func (o *TestPointFilterModel) SetPriorities(v []WorkItemPriorityModel) {
	o.Priorities = v
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAutomated.Get()
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutomated.Get(), o.IsAutomated.IsSet()
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasIsAutomated() bool {
	if o != nil && o.IsAutomated.IsSet() {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given NullableBool and assigns it to the IsAutomated field.
func (o *TestPointFilterModel) SetIsAutomated(v bool) {
	o.IsAutomated.Set(&v)
}
// SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil
func (o *TestPointFilterModel) SetIsAutomatedNil() {
	o.IsAutomated.Set(nil)
}

// UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
func (o *TestPointFilterModel) UnsetIsAutomated() {
	o.IsAutomated.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TestPointFilterModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TestPointFilterModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TestPointFilterModel) UnsetName() {
	o.Name.Unset()
}

// GetConfigurationIds returns the ConfigurationIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetConfigurationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ConfigurationIds
}

// GetConfigurationIdsOk returns a tuple with the ConfigurationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetConfigurationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConfigurationIds) {
		return nil, false
	}
	return o.ConfigurationIds, true
}

// HasConfigurationIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasConfigurationIds() bool {
	if o != nil && IsNil(o.ConfigurationIds) {
		return true
	}

	return false
}

// SetConfigurationIds gets a reference to the given []string and assigns it to the ConfigurationIds field.
func (o *TestPointFilterModel) SetConfigurationIds(v []string) {
	o.ConfigurationIds = v
}

// GetTesterIds returns the TesterIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetTesterIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TesterIds
}

// GetTesterIdsOk returns a tuple with the TesterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetTesterIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TesterIds) {
		return nil, false
	}
	return o.TesterIds, true
}

// HasTesterIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasTesterIds() bool {
	if o != nil && IsNil(o.TesterIds) {
		return true
	}

	return false
}

// SetTesterIds gets a reference to the given []string and assigns it to the TesterIds field.
func (o *TestPointFilterModel) SetTesterIds(v []string) {
	o.TesterIds = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *TestPointFilterModel) GetDuration() Int64RangeSelectorModel {
	if o == nil || IsNil(o.Duration) {
		var ret Int64RangeSelectorModel
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointFilterModel) GetDurationOk() (*Int64RangeSelectorModel, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given Int64RangeSelectorModel and assigns it to the Duration field.
func (o *TestPointFilterModel) SetDuration(v Int64RangeSelectorModel) {
	o.Duration = &v
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetSectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SectionIds
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetSectionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SectionIds) {
		return nil, false
	}
	return o.SectionIds, true
}

// HasSectionIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasSectionIds() bool {
	if o != nil && IsNil(o.SectionIds) {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given []string and assigns it to the SectionIds field.
func (o *TestPointFilterModel) SetSectionIds(v []string) {
	o.SectionIds = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *TestPointFilterModel) GetCreatedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.CreatedDate) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given DateTimeRangeSelectorModel and assigns it to the CreatedDate field.
func (o *TestPointFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel) {
	o.CreatedDate = &v
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *TestPointFilterModel) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *TestPointFilterModel) GetModifiedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.ModifiedDate) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointFilterModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil || IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasModifiedDate() bool {
	if o != nil && !IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given DateTimeRangeSelectorModel and assigns it to the ModifiedDate field.
func (o *TestPointFilterModel) SetModifiedDate(v DateTimeRangeSelectorModel) {
	o.ModifiedDate = &v
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *TestPointFilterModel) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TestPointFilterModel) SetTags(v []string) {
	o.Tags = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointFilterModel) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointFilterModel) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TestPointFilterModel) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *TestPointFilterModel) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

func (o TestPointFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TestPlanIds != nil {
		toSerialize["testPlanIds"] = o.TestPlanIds
	}
	if o.TestSuiteIds != nil {
		toSerialize["testSuiteIds"] = o.TestSuiteIds
	}
	if o.WorkItemGlobalIds != nil {
		toSerialize["workItemGlobalIds"] = o.WorkItemGlobalIds
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	if o.Priorities != nil {
		toSerialize["priorities"] = o.Priorities
	}
	if o.IsAutomated.IsSet() {
		toSerialize["isAutomated"] = o.IsAutomated.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ConfigurationIds != nil {
		toSerialize["configurationIds"] = o.ConfigurationIds
	}
	if o.TesterIds != nil {
		toSerialize["testerIds"] = o.TesterIds
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if o.SectionIds != nil {
		toSerialize["sectionIds"] = o.SectionIds
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.CreatedByIds != nil {
		toSerialize["createdByIds"] = o.CreatedByIds
	}
	if !IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if o.ModifiedByIds != nil {
		toSerialize["modifiedByIds"] = o.ModifiedByIds
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableTestPointFilterModel struct {
	value *TestPointFilterModel
	isSet bool
}

func (v NullableTestPointFilterModel) Get() *TestPointFilterModel {
	return v.value
}

func (v *NullableTestPointFilterModel) Set(val *TestPointFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointFilterModel(val *TestPointFilterModel) *NullableTestPointFilterModel {
	return &NullableTestPointFilterModel{value: val, isSet: true}
}

func (v NullableTestPointFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

