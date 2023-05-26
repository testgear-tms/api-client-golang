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

// checks if the TestPlanGroupByTesterAndStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanGroupByTesterAndStatus{}

// TestPlanGroupByTesterAndStatus struct for TestPlanGroupByTesterAndStatus
type TestPlanGroupByTesterAndStatus struct {
	UserId NullableString `json:"userId,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewTestPlanGroupByTesterAndStatus instantiates a new TestPlanGroupByTesterAndStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanGroupByTesterAndStatus() *TestPlanGroupByTesterAndStatus {
	this := TestPlanGroupByTesterAndStatus{}
	return &this
}

// NewTestPlanGroupByTesterAndStatusWithDefaults instantiates a new TestPlanGroupByTesterAndStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanGroupByTesterAndStatusWithDefaults() *TestPlanGroupByTesterAndStatus {
	this := TestPlanGroupByTesterAndStatus{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanGroupByTesterAndStatus) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanGroupByTesterAndStatus) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *TestPlanGroupByTesterAndStatus) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *TestPlanGroupByTesterAndStatus) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *TestPlanGroupByTesterAndStatus) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *TestPlanGroupByTesterAndStatus) UnsetUserId() {
	o.UserId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanGroupByTesterAndStatus) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanGroupByTesterAndStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TestPlanGroupByTesterAndStatus) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TestPlanGroupByTesterAndStatus) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TestPlanGroupByTesterAndStatus) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TestPlanGroupByTesterAndStatus) UnsetStatus() {
	o.Status.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TestPlanGroupByTesterAndStatus) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPlanGroupByTesterAndStatus) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TestPlanGroupByTesterAndStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *TestPlanGroupByTesterAndStatus) SetValue(v int32) {
	o.Value = &v
}

func (o TestPlanGroupByTesterAndStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanGroupByTesterAndStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTestPlanGroupByTesterAndStatus struct {
	value *TestPlanGroupByTesterAndStatus
	isSet bool
}

func (v NullableTestPlanGroupByTesterAndStatus) Get() *TestPlanGroupByTesterAndStatus {
	return v.value
}

func (v *NullableTestPlanGroupByTesterAndStatus) Set(val *TestPlanGroupByTesterAndStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanGroupByTesterAndStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanGroupByTesterAndStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanGroupByTesterAndStatus(val *TestPlanGroupByTesterAndStatus) *NullableTestPlanGroupByTesterAndStatus {
	return &NullableTestPlanGroupByTesterAndStatus{value: val, isSet: true}
}

func (v NullableTestPlanGroupByTesterAndStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanGroupByTesterAndStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

