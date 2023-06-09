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

// checks if the TestPointChangeViewModelTestPlanChangedFieldViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointChangeViewModelTestPlanChangedFieldViewModel{}

// TestPointChangeViewModelTestPlanChangedFieldViewModel struct for TestPointChangeViewModelTestPlanChangedFieldViewModel
type TestPointChangeViewModelTestPlanChangedFieldViewModel struct {
	OldValue *TestPointChangeViewModel `json:"oldValue,omitempty"`
	NewValue *TestPointChangeViewModel `json:"newValue,omitempty"`
}

// NewTestPointChangeViewModelTestPlanChangedFieldViewModel instantiates a new TestPointChangeViewModelTestPlanChangedFieldViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointChangeViewModelTestPlanChangedFieldViewModel() *TestPointChangeViewModelTestPlanChangedFieldViewModel {
	this := TestPointChangeViewModelTestPlanChangedFieldViewModel{}
	return &this
}

// NewTestPointChangeViewModelTestPlanChangedFieldViewModelWithDefaults instantiates a new TestPointChangeViewModelTestPlanChangedFieldViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointChangeViewModelTestPlanChangedFieldViewModelWithDefaults() *TestPointChangeViewModelTestPlanChangedFieldViewModel {
	this := TestPointChangeViewModelTestPlanChangedFieldViewModel{}
	return &this
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) GetOldValue() TestPointChangeViewModel {
	if o == nil || IsNil(o.OldValue) {
		var ret TestPointChangeViewModel
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) GetOldValueOk() (*TestPointChangeViewModel, bool) {
	if o == nil || IsNil(o.OldValue) {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) HasOldValue() bool {
	if o != nil && !IsNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given TestPointChangeViewModel and assigns it to the OldValue field.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) SetOldValue(v TestPointChangeViewModel) {
	o.OldValue = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) GetNewValue() TestPointChangeViewModel {
	if o == nil || IsNil(o.NewValue) {
		var ret TestPointChangeViewModel
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) GetNewValueOk() (*TestPointChangeViewModel, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) HasNewValue() bool {
	if o != nil && !IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given TestPointChangeViewModel and assigns it to the NewValue field.
func (o *TestPointChangeViewModelTestPlanChangedFieldViewModel) SetNewValue(v TestPointChangeViewModel) {
	o.NewValue = &v
}

func (o TestPointChangeViewModelTestPlanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointChangeViewModelTestPlanChangedFieldViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OldValue) {
		toSerialize["oldValue"] = o.OldValue
	}
	if !IsNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}
	return toSerialize, nil
}

type NullableTestPointChangeViewModelTestPlanChangedFieldViewModel struct {
	value *TestPointChangeViewModelTestPlanChangedFieldViewModel
	isSet bool
}

func (v NullableTestPointChangeViewModelTestPlanChangedFieldViewModel) Get() *TestPointChangeViewModelTestPlanChangedFieldViewModel {
	return v.value
}

func (v *NullableTestPointChangeViewModelTestPlanChangedFieldViewModel) Set(val *TestPointChangeViewModelTestPlanChangedFieldViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointChangeViewModelTestPlanChangedFieldViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointChangeViewModelTestPlanChangedFieldViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointChangeViewModelTestPlanChangedFieldViewModel(val *TestPointChangeViewModelTestPlanChangedFieldViewModel) *NullableTestPointChangeViewModelTestPlanChangedFieldViewModel {
	return &NullableTestPointChangeViewModelTestPlanChangedFieldViewModel{value: val, isSet: true}
}

func (v NullableTestPointChangeViewModelTestPlanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointChangeViewModelTestPlanChangedFieldViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


