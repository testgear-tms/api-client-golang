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

// checks if the TestRunStatisticsErrorCategoriesGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunStatisticsErrorCategoriesGetModel{}

// TestRunStatisticsErrorCategoriesGetModel struct for TestRunStatisticsErrorCategoriesGetModel
type TestRunStatisticsErrorCategoriesGetModel struct {
	// Number of test results which outcomes were not analyzed
	NoAnalytics *int32 `json:"noAnalytics,omitempty"`
	// Number of test results which outcomes were not caused by any defect
	NoDefect *int32 `json:"noDefect,omitempty"`
	// Number of test results which outcomes were caused by some infrastructure defect
	InfrastructureDefect *int32 `json:"infrastructureDefect,omitempty"`
	// Number of test results which outcomes were caused by some tested product defect
	ProductDefect *int32 `json:"productDefect,omitempty"`
	// Number of test results which outcomes were caused by test itself
	TestDefect *int32 `json:"testDefect,omitempty"`
}

// NewTestRunStatisticsErrorCategoriesGetModel instantiates a new TestRunStatisticsErrorCategoriesGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunStatisticsErrorCategoriesGetModel() *TestRunStatisticsErrorCategoriesGetModel {
	this := TestRunStatisticsErrorCategoriesGetModel{}
	return &this
}

// NewTestRunStatisticsErrorCategoriesGetModelWithDefaults instantiates a new TestRunStatisticsErrorCategoriesGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunStatisticsErrorCategoriesGetModelWithDefaults() *TestRunStatisticsErrorCategoriesGetModel {
	this := TestRunStatisticsErrorCategoriesGetModel{}
	return &this
}

// GetNoAnalytics returns the NoAnalytics field value if set, zero value otherwise.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetNoAnalytics() int32 {
	if o == nil || IsNil(o.NoAnalytics) {
		var ret int32
		return ret
	}
	return *o.NoAnalytics
}

// GetNoAnalyticsOk returns a tuple with the NoAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetNoAnalyticsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoAnalytics) {
		return nil, false
	}
	return o.NoAnalytics, true
}

// HasNoAnalytics returns a boolean if a field has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) HasNoAnalytics() bool {
	if o != nil && !IsNil(o.NoAnalytics) {
		return true
	}

	return false
}

// SetNoAnalytics gets a reference to the given int32 and assigns it to the NoAnalytics field.
func (o *TestRunStatisticsErrorCategoriesGetModel) SetNoAnalytics(v int32) {
	o.NoAnalytics = &v
}

// GetNoDefect returns the NoDefect field value if set, zero value otherwise.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetNoDefect() int32 {
	if o == nil || IsNil(o.NoDefect) {
		var ret int32
		return ret
	}
	return *o.NoDefect
}

// GetNoDefectOk returns a tuple with the NoDefect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetNoDefectOk() (*int32, bool) {
	if o == nil || IsNil(o.NoDefect) {
		return nil, false
	}
	return o.NoDefect, true
}

// HasNoDefect returns a boolean if a field has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) HasNoDefect() bool {
	if o != nil && !IsNil(o.NoDefect) {
		return true
	}

	return false
}

// SetNoDefect gets a reference to the given int32 and assigns it to the NoDefect field.
func (o *TestRunStatisticsErrorCategoriesGetModel) SetNoDefect(v int32) {
	o.NoDefect = &v
}

// GetInfrastructureDefect returns the InfrastructureDefect field value if set, zero value otherwise.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetInfrastructureDefect() int32 {
	if o == nil || IsNil(o.InfrastructureDefect) {
		var ret int32
		return ret
	}
	return *o.InfrastructureDefect
}

// GetInfrastructureDefectOk returns a tuple with the InfrastructureDefect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetInfrastructureDefectOk() (*int32, bool) {
	if o == nil || IsNil(o.InfrastructureDefect) {
		return nil, false
	}
	return o.InfrastructureDefect, true
}

// HasInfrastructureDefect returns a boolean if a field has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) HasInfrastructureDefect() bool {
	if o != nil && !IsNil(o.InfrastructureDefect) {
		return true
	}

	return false
}

// SetInfrastructureDefect gets a reference to the given int32 and assigns it to the InfrastructureDefect field.
func (o *TestRunStatisticsErrorCategoriesGetModel) SetInfrastructureDefect(v int32) {
	o.InfrastructureDefect = &v
}

// GetProductDefect returns the ProductDefect field value if set, zero value otherwise.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetProductDefect() int32 {
	if o == nil || IsNil(o.ProductDefect) {
		var ret int32
		return ret
	}
	return *o.ProductDefect
}

// GetProductDefectOk returns a tuple with the ProductDefect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetProductDefectOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductDefect) {
		return nil, false
	}
	return o.ProductDefect, true
}

// HasProductDefect returns a boolean if a field has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) HasProductDefect() bool {
	if o != nil && !IsNil(o.ProductDefect) {
		return true
	}

	return false
}

// SetProductDefect gets a reference to the given int32 and assigns it to the ProductDefect field.
func (o *TestRunStatisticsErrorCategoriesGetModel) SetProductDefect(v int32) {
	o.ProductDefect = &v
}

// GetTestDefect returns the TestDefect field value if set, zero value otherwise.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetTestDefect() int32 {
	if o == nil || IsNil(o.TestDefect) {
		var ret int32
		return ret
	}
	return *o.TestDefect
}

// GetTestDefectOk returns a tuple with the TestDefect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) GetTestDefectOk() (*int32, bool) {
	if o == nil || IsNil(o.TestDefect) {
		return nil, false
	}
	return o.TestDefect, true
}

// HasTestDefect returns a boolean if a field has been set.
func (o *TestRunStatisticsErrorCategoriesGetModel) HasTestDefect() bool {
	if o != nil && !IsNil(o.TestDefect) {
		return true
	}

	return false
}

// SetTestDefect gets a reference to the given int32 and assigns it to the TestDefect field.
func (o *TestRunStatisticsErrorCategoriesGetModel) SetTestDefect(v int32) {
	o.TestDefect = &v
}

func (o TestRunStatisticsErrorCategoriesGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunStatisticsErrorCategoriesGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoAnalytics) {
		toSerialize["noAnalytics"] = o.NoAnalytics
	}
	if !IsNil(o.NoDefect) {
		toSerialize["noDefect"] = o.NoDefect
	}
	if !IsNil(o.InfrastructureDefect) {
		toSerialize["infrastructureDefect"] = o.InfrastructureDefect
	}
	if !IsNil(o.ProductDefect) {
		toSerialize["productDefect"] = o.ProductDefect
	}
	if !IsNil(o.TestDefect) {
		toSerialize["testDefect"] = o.TestDefect
	}
	return toSerialize, nil
}

type NullableTestRunStatisticsErrorCategoriesGetModel struct {
	value *TestRunStatisticsErrorCategoriesGetModel
	isSet bool
}

func (v NullableTestRunStatisticsErrorCategoriesGetModel) Get() *TestRunStatisticsErrorCategoriesGetModel {
	return v.value
}

func (v *NullableTestRunStatisticsErrorCategoriesGetModel) Set(val *TestRunStatisticsErrorCategoriesGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunStatisticsErrorCategoriesGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunStatisticsErrorCategoriesGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunStatisticsErrorCategoriesGetModel(val *TestRunStatisticsErrorCategoriesGetModel) *NullableTestRunStatisticsErrorCategoriesGetModel {
	return &NullableTestRunStatisticsErrorCategoriesGetModel{value: val, isSet: true}
}

func (v NullableTestRunStatisticsErrorCategoriesGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunStatisticsErrorCategoriesGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

