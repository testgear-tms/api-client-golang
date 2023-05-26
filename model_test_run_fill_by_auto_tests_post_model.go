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

// checks if the TestRunFillByAutoTestsPostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunFillByAutoTestsPostModel{}

// TestRunFillByAutoTestsPostModel struct for TestRunFillByAutoTestsPostModel
type TestRunFillByAutoTestsPostModel struct {
	// Specifies the GUID of the project, in which a test run will be created.
	ProjectId string `json:"projectId"`
	// Specifies the name of the test run.
	Name NullableString `json:"name,omitempty"`
	// Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs.
	ConfigurationIds []string `json:"configurationIds"`
	// Specifies the external ID of the autotest. You can specify several IDs.
	AutoTestExternalIds []string `json:"autoTestExternalIds"`
	// Specifies the test run description.
	Description NullableString `json:"description,omitempty"`
	// Specifies the test run launch source.
	LaunchSource NullableString `json:"launchSource,omitempty"`
}

// NewTestRunFillByAutoTestsPostModel instantiates a new TestRunFillByAutoTestsPostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunFillByAutoTestsPostModel(projectId string, configurationIds []string, autoTestExternalIds []string) *TestRunFillByAutoTestsPostModel {
	this := TestRunFillByAutoTestsPostModel{}
	this.ProjectId = projectId
	this.ConfigurationIds = configurationIds
	this.AutoTestExternalIds = autoTestExternalIds
	return &this
}

// NewTestRunFillByAutoTestsPostModelWithDefaults instantiates a new TestRunFillByAutoTestsPostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunFillByAutoTestsPostModelWithDefaults() *TestRunFillByAutoTestsPostModel {
	this := TestRunFillByAutoTestsPostModel{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *TestRunFillByAutoTestsPostModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TestRunFillByAutoTestsPostModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TestRunFillByAutoTestsPostModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByAutoTestsPostModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByAutoTestsPostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TestRunFillByAutoTestsPostModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TestRunFillByAutoTestsPostModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TestRunFillByAutoTestsPostModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TestRunFillByAutoTestsPostModel) UnsetName() {
	o.Name.Unset()
}

// GetConfigurationIds returns the ConfigurationIds field value
func (o *TestRunFillByAutoTestsPostModel) GetConfigurationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfigurationIds
}

// GetConfigurationIdsOk returns a tuple with the ConfigurationIds field value
// and a boolean to check if the value has been set.
func (o *TestRunFillByAutoTestsPostModel) GetConfigurationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationIds, true
}

// SetConfigurationIds sets field value
func (o *TestRunFillByAutoTestsPostModel) SetConfigurationIds(v []string) {
	o.ConfigurationIds = v
}

// GetAutoTestExternalIds returns the AutoTestExternalIds field value
func (o *TestRunFillByAutoTestsPostModel) GetAutoTestExternalIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AutoTestExternalIds
}

// GetAutoTestExternalIdsOk returns a tuple with the AutoTestExternalIds field value
// and a boolean to check if the value has been set.
func (o *TestRunFillByAutoTestsPostModel) GetAutoTestExternalIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestExternalIds, true
}

// SetAutoTestExternalIds sets field value
func (o *TestRunFillByAutoTestsPostModel) SetAutoTestExternalIds(v []string) {
	o.AutoTestExternalIds = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByAutoTestsPostModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByAutoTestsPostModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestRunFillByAutoTestsPostModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestRunFillByAutoTestsPostModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestRunFillByAutoTestsPostModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestRunFillByAutoTestsPostModel) UnsetDescription() {
	o.Description.Unset()
}

// GetLaunchSource returns the LaunchSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunFillByAutoTestsPostModel) GetLaunchSource() string {
	if o == nil || IsNil(o.LaunchSource.Get()) {
		var ret string
		return ret
	}
	return *o.LaunchSource.Get()
}

// GetLaunchSourceOk returns a tuple with the LaunchSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunFillByAutoTestsPostModel) GetLaunchSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchSource.Get(), o.LaunchSource.IsSet()
}

// HasLaunchSource returns a boolean if a field has been set.
func (o *TestRunFillByAutoTestsPostModel) HasLaunchSource() bool {
	if o != nil && o.LaunchSource.IsSet() {
		return true
	}

	return false
}

// SetLaunchSource gets a reference to the given NullableString and assigns it to the LaunchSource field.
func (o *TestRunFillByAutoTestsPostModel) SetLaunchSource(v string) {
	o.LaunchSource.Set(&v)
}
// SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil
func (o *TestRunFillByAutoTestsPostModel) SetLaunchSourceNil() {
	o.LaunchSource.Set(nil)
}

// UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
func (o *TestRunFillByAutoTestsPostModel) UnsetLaunchSource() {
	o.LaunchSource.Unset()
}

func (o TestRunFillByAutoTestsPostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunFillByAutoTestsPostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["configurationIds"] = o.ConfigurationIds
	toSerialize["autoTestExternalIds"] = o.AutoTestExternalIds
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.LaunchSource.IsSet() {
		toSerialize["launchSource"] = o.LaunchSource.Get()
	}
	return toSerialize, nil
}

type NullableTestRunFillByAutoTestsPostModel struct {
	value *TestRunFillByAutoTestsPostModel
	isSet bool
}

func (v NullableTestRunFillByAutoTestsPostModel) Get() *TestRunFillByAutoTestsPostModel {
	return v.value
}

func (v *NullableTestRunFillByAutoTestsPostModel) Set(val *TestRunFillByAutoTestsPostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunFillByAutoTestsPostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunFillByAutoTestsPostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunFillByAutoTestsPostModel(val *TestRunFillByAutoTestsPostModel) *NullableTestRunFillByAutoTestsPostModel {
	return &NullableTestRunFillByAutoTestsPostModel{value: val, isSet: true}
}

func (v NullableTestRunFillByAutoTestsPostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunFillByAutoTestsPostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


