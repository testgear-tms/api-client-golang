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

// checks if the TestPointByTestSuiteModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointByTestSuiteModel{}

// TestPointByTestSuiteModel struct for TestPointByTestSuiteModel
type TestPointByTestSuiteModel struct {
	Id *string `json:"id,omitempty"`
	TesterId NullableString `json:"testerId,omitempty"`
	WorkItemId NullableString `json:"workItemId,omitempty"`
	ConfigurationId NullableString `json:"configurationId,omitempty"`
	// Applies one of these values: Blocked, NoResults, Failed, Passed
	Status NullableString `json:"status,omitempty"`
	LastTestResultId NullableString `json:"lastTestResultId,omitempty"`
	IterationId *string `json:"iterationId,omitempty"`
}

// NewTestPointByTestSuiteModel instantiates a new TestPointByTestSuiteModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointByTestSuiteModel() *TestPointByTestSuiteModel {
	this := TestPointByTestSuiteModel{}
	return &this
}

// NewTestPointByTestSuiteModelWithDefaults instantiates a new TestPointByTestSuiteModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointByTestSuiteModelWithDefaults() *TestPointByTestSuiteModel {
	this := TestPointByTestSuiteModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestPointByTestSuiteModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointByTestSuiteModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestPointByTestSuiteModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestPointByTestSuiteModel) SetId(v string) {
	o.Id = &v
}

// GetTesterId returns the TesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointByTestSuiteModel) GetTesterId() string {
	if o == nil || IsNil(o.TesterId.Get()) {
		var ret string
		return ret
	}
	return *o.TesterId.Get()
}

// GetTesterIdOk returns a tuple with the TesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointByTestSuiteModel) GetTesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TesterId.Get(), o.TesterId.IsSet()
}

// HasTesterId returns a boolean if a field has been set.
func (o *TestPointByTestSuiteModel) HasTesterId() bool {
	if o != nil && o.TesterId.IsSet() {
		return true
	}

	return false
}

// SetTesterId gets a reference to the given NullableString and assigns it to the TesterId field.
func (o *TestPointByTestSuiteModel) SetTesterId(v string) {
	o.TesterId.Set(&v)
}
// SetTesterIdNil sets the value for TesterId to be an explicit nil
func (o *TestPointByTestSuiteModel) SetTesterIdNil() {
	o.TesterId.Set(nil)
}

// UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
func (o *TestPointByTestSuiteModel) UnsetTesterId() {
	o.TesterId.Unset()
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointByTestSuiteModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemId.Get()
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointByTestSuiteModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemId.Get(), o.WorkItemId.IsSet()
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *TestPointByTestSuiteModel) HasWorkItemId() bool {
	if o != nil && o.WorkItemId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given NullableString and assigns it to the WorkItemId field.
func (o *TestPointByTestSuiteModel) SetWorkItemId(v string) {
	o.WorkItemId.Set(&v)
}
// SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil
func (o *TestPointByTestSuiteModel) SetWorkItemIdNil() {
	o.WorkItemId.Set(nil)
}

// UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
func (o *TestPointByTestSuiteModel) UnsetWorkItemId() {
	o.WorkItemId.Unset()
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointByTestSuiteModel) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigurationId.Get()
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointByTestSuiteModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationId.Get(), o.ConfigurationId.IsSet()
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *TestPointByTestSuiteModel) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given NullableString and assigns it to the ConfigurationId field.
func (o *TestPointByTestSuiteModel) SetConfigurationId(v string) {
	o.ConfigurationId.Set(&v)
}
// SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil
func (o *TestPointByTestSuiteModel) SetConfigurationIdNil() {
	o.ConfigurationId.Set(nil)
}

// UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
func (o *TestPointByTestSuiteModel) UnsetConfigurationId() {
	o.ConfigurationId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointByTestSuiteModel) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointByTestSuiteModel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TestPointByTestSuiteModel) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TestPointByTestSuiteModel) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TestPointByTestSuiteModel) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TestPointByTestSuiteModel) UnsetStatus() {
	o.Status.Unset()
}

// GetLastTestResultId returns the LastTestResultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointByTestSuiteModel) GetLastTestResultId() string {
	if o == nil || IsNil(o.LastTestResultId.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestResultId.Get()
}

// GetLastTestResultIdOk returns a tuple with the LastTestResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointByTestSuiteModel) GetLastTestResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestResultId.Get(), o.LastTestResultId.IsSet()
}

// HasLastTestResultId returns a boolean if a field has been set.
func (o *TestPointByTestSuiteModel) HasLastTestResultId() bool {
	if o != nil && o.LastTestResultId.IsSet() {
		return true
	}

	return false
}

// SetLastTestResultId gets a reference to the given NullableString and assigns it to the LastTestResultId field.
func (o *TestPointByTestSuiteModel) SetLastTestResultId(v string) {
	o.LastTestResultId.Set(&v)
}
// SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil
func (o *TestPointByTestSuiteModel) SetLastTestResultIdNil() {
	o.LastTestResultId.Set(nil)
}

// UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
func (o *TestPointByTestSuiteModel) UnsetLastTestResultId() {
	o.LastTestResultId.Unset()
}

// GetIterationId returns the IterationId field value if set, zero value otherwise.
func (o *TestPointByTestSuiteModel) GetIterationId() string {
	if o == nil || IsNil(o.IterationId) {
		var ret string
		return ret
	}
	return *o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointByTestSuiteModel) GetIterationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IterationId) {
		return nil, false
	}
	return o.IterationId, true
}

// HasIterationId returns a boolean if a field has been set.
func (o *TestPointByTestSuiteModel) HasIterationId() bool {
	if o != nil && !IsNil(o.IterationId) {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given string and assigns it to the IterationId field.
func (o *TestPointByTestSuiteModel) SetIterationId(v string) {
	o.IterationId = &v
}

func (o TestPointByTestSuiteModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointByTestSuiteModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.TesterId.IsSet() {
		toSerialize["testerId"] = o.TesterId.Get()
	}
	if o.WorkItemId.IsSet() {
		toSerialize["workItemId"] = o.WorkItemId.Get()
	}
	if o.ConfigurationId.IsSet() {
		toSerialize["configurationId"] = o.ConfigurationId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.LastTestResultId.IsSet() {
		toSerialize["lastTestResultId"] = o.LastTestResultId.Get()
	}
	if !IsNil(o.IterationId) {
		toSerialize["iterationId"] = o.IterationId
	}
	return toSerialize, nil
}

type NullableTestPointByTestSuiteModel struct {
	value *TestPointByTestSuiteModel
	isSet bool
}

func (v NullableTestPointByTestSuiteModel) Get() *TestPointByTestSuiteModel {
	return v.value
}

func (v *NullableTestPointByTestSuiteModel) Set(val *TestPointByTestSuiteModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointByTestSuiteModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointByTestSuiteModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointByTestSuiteModel(val *TestPointByTestSuiteModel) *NullableTestPointByTestSuiteModel {
	return &NullableTestPointByTestSuiteModel{value: val, isSet: true}
}

func (v NullableTestPointByTestSuiteModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointByTestSuiteModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


