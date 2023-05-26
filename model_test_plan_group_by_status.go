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

// checks if the TestPlanGroupByStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanGroupByStatus{}

// TestPlanGroupByStatus struct for TestPlanGroupByStatus
type TestPlanGroupByStatus struct {
	Status NullableString `json:"status,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewTestPlanGroupByStatus instantiates a new TestPlanGroupByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanGroupByStatus() *TestPlanGroupByStatus {
	this := TestPlanGroupByStatus{}
	return &this
}

// NewTestPlanGroupByStatusWithDefaults instantiates a new TestPlanGroupByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanGroupByStatusWithDefaults() *TestPlanGroupByStatus {
	this := TestPlanGroupByStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanGroupByStatus) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanGroupByStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TestPlanGroupByStatus) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TestPlanGroupByStatus) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TestPlanGroupByStatus) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TestPlanGroupByStatus) UnsetStatus() {
	o.Status.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TestPlanGroupByStatus) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPlanGroupByStatus) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TestPlanGroupByStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *TestPlanGroupByStatus) SetValue(v int32) {
	o.Value = &v
}

func (o TestPlanGroupByStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanGroupByStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTestPlanGroupByStatus struct {
	value *TestPlanGroupByStatus
	isSet bool
}

func (v NullableTestPlanGroupByStatus) Get() *TestPlanGroupByStatus {
	return v.value
}

func (v *NullableTestPlanGroupByStatus) Set(val *TestPlanGroupByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanGroupByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanGroupByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanGroupByStatus(val *TestPlanGroupByStatus) *NullableTestPlanGroupByStatus {
	return &NullableTestPlanGroupByStatus{value: val, isSet: true}
}

func (v NullableTestPlanGroupByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanGroupByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


