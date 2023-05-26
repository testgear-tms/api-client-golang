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

// checks if the TestPlanGroupByTester type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanGroupByTester{}

// TestPlanGroupByTester struct for TestPlanGroupByTester
type TestPlanGroupByTester struct {
	UserId NullableString `json:"userId,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewTestPlanGroupByTester instantiates a new TestPlanGroupByTester object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanGroupByTester() *TestPlanGroupByTester {
	this := TestPlanGroupByTester{}
	return &this
}

// NewTestPlanGroupByTesterWithDefaults instantiates a new TestPlanGroupByTester object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanGroupByTesterWithDefaults() *TestPlanGroupByTester {
	this := TestPlanGroupByTester{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanGroupByTester) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanGroupByTester) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *TestPlanGroupByTester) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *TestPlanGroupByTester) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *TestPlanGroupByTester) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *TestPlanGroupByTester) UnsetUserId() {
	o.UserId.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TestPlanGroupByTester) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPlanGroupByTester) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TestPlanGroupByTester) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *TestPlanGroupByTester) SetValue(v int32) {
	o.Value = &v
}

func (o TestPlanGroupByTester) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanGroupByTester) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTestPlanGroupByTester struct {
	value *TestPlanGroupByTester
	isSet bool
}

func (v NullableTestPlanGroupByTester) Get() *TestPlanGroupByTester {
	return v.value
}

func (v *NullableTestPlanGroupByTester) Set(val *TestPlanGroupByTester) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanGroupByTester) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanGroupByTester) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanGroupByTester(val *TestPlanGroupByTester) *NullableTestPlanGroupByTester {
	return &NullableTestPlanGroupByTester{value: val, isSet: true}
}

func (v NullableTestPlanGroupByTester) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanGroupByTester) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


