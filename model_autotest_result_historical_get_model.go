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

// checks if the AutotestResultHistoricalGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutotestResultHistoricalGetModel{}

// AutotestResultHistoricalGetModel struct for AutotestResultHistoricalGetModel
type AutotestResultHistoricalGetModel struct {
	Id *string `json:"id,omitempty"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
	TestRunId *string `json:"testRunId,omitempty"`
	TestRunName NullableString `json:"testRunName,omitempty"`
	ConfigurationId *string `json:"configurationId,omitempty"`
	Outcome AutotestResultOutcome `json:"outcome"`
	LaunchSource NullableString `json:"launchSource,omitempty"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	TestPlanId NullableString `json:"testPlanId,omitempty"`
	TestPlanGlobalId NullableInt64 `json:"testPlanGlobalId,omitempty"`
	TestPlanName NullableString `json:"testPlanName,omitempty"`
	Duration NullableInt64 `json:"duration,omitempty"`
}

// NewAutotestResultHistoricalGetModel instantiates a new AutotestResultHistoricalGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutotestResultHistoricalGetModel(outcome AutotestResultOutcome) *AutotestResultHistoricalGetModel {
	this := AutotestResultHistoricalGetModel{}
	this.Outcome = outcome
	return &this
}

// NewAutotestResultHistoricalGetModelWithDefaults instantiates a new AutotestResultHistoricalGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutotestResultHistoricalGetModelWithDefaults() *AutotestResultHistoricalGetModel {
	this := AutotestResultHistoricalGetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutotestResultHistoricalGetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestResultHistoricalGetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutotestResultHistoricalGetModel) SetId(v string) {
	o.Id = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AutotestResultHistoricalGetModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestResultHistoricalGetModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *AutotestResultHistoricalGetModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *AutotestResultHistoricalGetModel) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestResultHistoricalGetModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *AutotestResultHistoricalGetModel) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetTestRunId returns the TestRunId field value if set, zero value otherwise.
func (o *AutotestResultHistoricalGetModel) GetTestRunId() string {
	if o == nil || IsNil(o.TestRunId) {
		var ret string
		return ret
	}
	return *o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestResultHistoricalGetModel) GetTestRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestRunId) {
		return nil, false
	}
	return o.TestRunId, true
}

// HasTestRunId returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasTestRunId() bool {
	if o != nil && !IsNil(o.TestRunId) {
		return true
	}

	return false
}

// SetTestRunId gets a reference to the given string and assigns it to the TestRunId field.
func (o *AutotestResultHistoricalGetModel) SetTestRunId(v string) {
	o.TestRunId = &v
}

// GetTestRunName returns the TestRunName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetTestRunName() string {
	if o == nil || IsNil(o.TestRunName.Get()) {
		var ret string
		return ret
	}
	return *o.TestRunName.Get()
}

// GetTestRunNameOk returns a tuple with the TestRunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetTestRunNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestRunName.Get(), o.TestRunName.IsSet()
}

// HasTestRunName returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasTestRunName() bool {
	if o != nil && o.TestRunName.IsSet() {
		return true
	}

	return false
}

// SetTestRunName gets a reference to the given NullableString and assigns it to the TestRunName field.
func (o *AutotestResultHistoricalGetModel) SetTestRunName(v string) {
	o.TestRunName.Set(&v)
}
// SetTestRunNameNil sets the value for TestRunName to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetTestRunNameNil() {
	o.TestRunName.Set(nil)
}

// UnsetTestRunName ensures that no value is present for TestRunName, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetTestRunName() {
	o.TestRunName.Unset()
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *AutotestResultHistoricalGetModel) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestResultHistoricalGetModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationId) {
		return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasConfigurationId() bool {
	if o != nil && !IsNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *AutotestResultHistoricalGetModel) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetOutcome returns the Outcome field value
func (o *AutotestResultHistoricalGetModel) GetOutcome() AutotestResultOutcome {
	if o == nil {
		var ret AutotestResultOutcome
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *AutotestResultHistoricalGetModel) GetOutcomeOk() (*AutotestResultOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *AutotestResultHistoricalGetModel) SetOutcome(v AutotestResultOutcome) {
	o.Outcome = v
}

// GetLaunchSource returns the LaunchSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetLaunchSource() string {
	if o == nil || IsNil(o.LaunchSource.Get()) {
		var ret string
		return ret
	}
	return *o.LaunchSource.Get()
}

// GetLaunchSourceOk returns a tuple with the LaunchSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetLaunchSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchSource.Get(), o.LaunchSource.IsSet()
}

// HasLaunchSource returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasLaunchSource() bool {
	if o != nil && o.LaunchSource.IsSet() {
		return true
	}

	return false
}

// SetLaunchSource gets a reference to the given NullableString and assigns it to the LaunchSource field.
func (o *AutotestResultHistoricalGetModel) SetLaunchSource(v string) {
	o.LaunchSource.Set(&v)
}
// SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetLaunchSourceNil() {
	o.LaunchSource.Set(nil)
}

// UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetLaunchSource() {
	o.LaunchSource.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *AutotestResultHistoricalGetModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *AutotestResultHistoricalGetModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetTestPlanId returns the TestPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetTestPlanId() string {
	if o == nil || IsNil(o.TestPlanId.Get()) {
		var ret string
		return ret
	}
	return *o.TestPlanId.Get()
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanId.Get(), o.TestPlanId.IsSet()
}

// HasTestPlanId returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasTestPlanId() bool {
	if o != nil && o.TestPlanId.IsSet() {
		return true
	}

	return false
}

// SetTestPlanId gets a reference to the given NullableString and assigns it to the TestPlanId field.
func (o *AutotestResultHistoricalGetModel) SetTestPlanId(v string) {
	o.TestPlanId.Set(&v)
}
// SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetTestPlanIdNil() {
	o.TestPlanId.Set(nil)
}

// UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetTestPlanId() {
	o.TestPlanId.Unset()
}

// GetTestPlanGlobalId returns the TestPlanGlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetTestPlanGlobalId() int64 {
	if o == nil || IsNil(o.TestPlanGlobalId.Get()) {
		var ret int64
		return ret
	}
	return *o.TestPlanGlobalId.Get()
}

// GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetTestPlanGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanGlobalId.Get(), o.TestPlanGlobalId.IsSet()
}

// HasTestPlanGlobalId returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasTestPlanGlobalId() bool {
	if o != nil && o.TestPlanGlobalId.IsSet() {
		return true
	}

	return false
}

// SetTestPlanGlobalId gets a reference to the given NullableInt64 and assigns it to the TestPlanGlobalId field.
func (o *AutotestResultHistoricalGetModel) SetTestPlanGlobalId(v int64) {
	o.TestPlanGlobalId.Set(&v)
}
// SetTestPlanGlobalIdNil sets the value for TestPlanGlobalId to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetTestPlanGlobalIdNil() {
	o.TestPlanGlobalId.Set(nil)
}

// UnsetTestPlanGlobalId ensures that no value is present for TestPlanGlobalId, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetTestPlanGlobalId() {
	o.TestPlanGlobalId.Unset()
}

// GetTestPlanName returns the TestPlanName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetTestPlanName() string {
	if o == nil || IsNil(o.TestPlanName.Get()) {
		var ret string
		return ret
	}
	return *o.TestPlanName.Get()
}

// GetTestPlanNameOk returns a tuple with the TestPlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetTestPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanName.Get(), o.TestPlanName.IsSet()
}

// HasTestPlanName returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasTestPlanName() bool {
	if o != nil && o.TestPlanName.IsSet() {
		return true
	}

	return false
}

// SetTestPlanName gets a reference to the given NullableString and assigns it to the TestPlanName field.
func (o *AutotestResultHistoricalGetModel) SetTestPlanName(v string) {
	o.TestPlanName.Set(&v)
}
// SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetTestPlanNameNil() {
	o.TestPlanName.Set(nil)
}

// UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetTestPlanName() {
	o.TestPlanName.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutotestResultHistoricalGetModel) GetDuration() int64 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutotestResultHistoricalGetModel) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *AutotestResultHistoricalGetModel) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *AutotestResultHistoricalGetModel) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *AutotestResultHistoricalGetModel) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *AutotestResultHistoricalGetModel) UnsetDuration() {
	o.Duration.Unset()
}

func (o AutotestResultHistoricalGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutotestResultHistoricalGetModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TestRunId) {
		toSerialize["testRunId"] = o.TestRunId
	}
	if o.TestRunName.IsSet() {
		toSerialize["testRunName"] = o.TestRunName.Get()
	}
	if !IsNil(o.ConfigurationId) {
		toSerialize["configurationId"] = o.ConfigurationId
	}
	toSerialize["outcome"] = o.Outcome
	if o.LaunchSource.IsSet() {
		toSerialize["launchSource"] = o.LaunchSource.Get()
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.TestPlanId.IsSet() {
		toSerialize["testPlanId"] = o.TestPlanId.Get()
	}
	if o.TestPlanGlobalId.IsSet() {
		toSerialize["testPlanGlobalId"] = o.TestPlanGlobalId.Get()
	}
	if o.TestPlanName.IsSet() {
		toSerialize["testPlanName"] = o.TestPlanName.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	return toSerialize, nil
}

type NullableAutotestResultHistoricalGetModel struct {
	value *AutotestResultHistoricalGetModel
	isSet bool
}

func (v NullableAutotestResultHistoricalGetModel) Get() *AutotestResultHistoricalGetModel {
	return v.value
}

func (v *NullableAutotestResultHistoricalGetModel) Set(val *AutotestResultHistoricalGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutotestResultHistoricalGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutotestResultHistoricalGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutotestResultHistoricalGetModel(val *AutotestResultHistoricalGetModel) *NullableAutotestResultHistoricalGetModel {
	return &NullableAutotestResultHistoricalGetModel{value: val, isSet: true}
}

func (v NullableAutotestResultHistoricalGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutotestResultHistoricalGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


