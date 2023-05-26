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

// checks if the TestResultsStatisticsGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultsStatisticsGetModel{}

// TestResultsStatisticsGetModel struct for TestResultsStatisticsGetModel
type TestResultsStatisticsGetModel struct {
	Statuses *TestRunStatisticsStatusesGetModel `json:"statuses,omitempty"`
	FailureCategories *TestRunStatisticsErrorCategoriesGetModel `json:"failureCategories,omitempty"`
}

// NewTestResultsStatisticsGetModel instantiates a new TestResultsStatisticsGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultsStatisticsGetModel() *TestResultsStatisticsGetModel {
	this := TestResultsStatisticsGetModel{}
	return &this
}

// NewTestResultsStatisticsGetModelWithDefaults instantiates a new TestResultsStatisticsGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultsStatisticsGetModelWithDefaults() *TestResultsStatisticsGetModel {
	this := TestResultsStatisticsGetModel{}
	return &this
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *TestResultsStatisticsGetModel) GetStatuses() TestRunStatisticsStatusesGetModel {
	if o == nil || IsNil(o.Statuses) {
		var ret TestRunStatisticsStatusesGetModel
		return ret
	}
	return *o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModel) GetStatusesOk() (*TestRunStatisticsStatusesGetModel, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *TestResultsStatisticsGetModel) HasStatuses() bool {
	if o != nil && !IsNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given TestRunStatisticsStatusesGetModel and assigns it to the Statuses field.
func (o *TestResultsStatisticsGetModel) SetStatuses(v TestRunStatisticsStatusesGetModel) {
	o.Statuses = &v
}

// GetFailureCategories returns the FailureCategories field value if set, zero value otherwise.
func (o *TestResultsStatisticsGetModel) GetFailureCategories() TestRunStatisticsErrorCategoriesGetModel {
	if o == nil || IsNil(o.FailureCategories) {
		var ret TestRunStatisticsErrorCategoriesGetModel
		return ret
	}
	return *o.FailureCategories
}

// GetFailureCategoriesOk returns a tuple with the FailureCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModel) GetFailureCategoriesOk() (*TestRunStatisticsErrorCategoriesGetModel, bool) {
	if o == nil || IsNil(o.FailureCategories) {
		return nil, false
	}
	return o.FailureCategories, true
}

// HasFailureCategories returns a boolean if a field has been set.
func (o *TestResultsStatisticsGetModel) HasFailureCategories() bool {
	if o != nil && !IsNil(o.FailureCategories) {
		return true
	}

	return false
}

// SetFailureCategories gets a reference to the given TestRunStatisticsErrorCategoriesGetModel and assigns it to the FailureCategories field.
func (o *TestResultsStatisticsGetModel) SetFailureCategories(v TestRunStatisticsErrorCategoriesGetModel) {
	o.FailureCategories = &v
}

func (o TestResultsStatisticsGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultsStatisticsGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	if !IsNil(o.FailureCategories) {
		toSerialize["failureCategories"] = o.FailureCategories
	}
	return toSerialize, nil
}

type NullableTestResultsStatisticsGetModel struct {
	value *TestResultsStatisticsGetModel
	isSet bool
}

func (v NullableTestResultsStatisticsGetModel) Get() *TestResultsStatisticsGetModel {
	return v.value
}

func (v *NullableTestResultsStatisticsGetModel) Set(val *TestResultsStatisticsGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultsStatisticsGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultsStatisticsGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultsStatisticsGetModel(val *TestResultsStatisticsGetModel) *NullableTestResultsStatisticsGetModel {
	return &NullableTestResultsStatisticsGetModel{value: val, isSet: true}
}

func (v NullableTestResultsStatisticsGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultsStatisticsGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

