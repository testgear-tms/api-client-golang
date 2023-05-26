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

// checks if the StringTestPlanChangedFieldViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringTestPlanChangedFieldViewModel{}

// StringTestPlanChangedFieldViewModel struct for StringTestPlanChangedFieldViewModel
type StringTestPlanChangedFieldViewModel struct {
	OldValue NullableString `json:"oldValue,omitempty"`
	NewValue NullableString `json:"newValue,omitempty"`
}

// NewStringTestPlanChangedFieldViewModel instantiates a new StringTestPlanChangedFieldViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringTestPlanChangedFieldViewModel() *StringTestPlanChangedFieldViewModel {
	this := StringTestPlanChangedFieldViewModel{}
	return &this
}

// NewStringTestPlanChangedFieldViewModelWithDefaults instantiates a new StringTestPlanChangedFieldViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringTestPlanChangedFieldViewModelWithDefaults() *StringTestPlanChangedFieldViewModel {
	this := StringTestPlanChangedFieldViewModel{}
	return &this
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringTestPlanChangedFieldViewModel) GetOldValue() string {
	if o == nil || IsNil(o.OldValue.Get()) {
		var ret string
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringTestPlanChangedFieldViewModel) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *StringTestPlanChangedFieldViewModel) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableString and assigns it to the OldValue field.
func (o *StringTestPlanChangedFieldViewModel) SetOldValue(v string) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *StringTestPlanChangedFieldViewModel) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *StringTestPlanChangedFieldViewModel) UnsetOldValue() {
	o.OldValue.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringTestPlanChangedFieldViewModel) GetNewValue() string {
	if o == nil || IsNil(o.NewValue.Get()) {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringTestPlanChangedFieldViewModel) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *StringTestPlanChangedFieldViewModel) HasNewValue() bool {
	if o != nil && o.NewValue.IsSet() {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given NullableString and assigns it to the NewValue field.
func (o *StringTestPlanChangedFieldViewModel) SetNewValue(v string) {
	o.NewValue.Set(&v)
}
// SetNewValueNil sets the value for NewValue to be an explicit nil
func (o *StringTestPlanChangedFieldViewModel) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
func (o *StringTestPlanChangedFieldViewModel) UnsetNewValue() {
	o.NewValue.Unset()
}

func (o StringTestPlanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringTestPlanChangedFieldViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OldValue.IsSet() {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	if o.NewValue.IsSet() {
		toSerialize["newValue"] = o.NewValue.Get()
	}
	return toSerialize, nil
}

type NullableStringTestPlanChangedFieldViewModel struct {
	value *StringTestPlanChangedFieldViewModel
	isSet bool
}

func (v NullableStringTestPlanChangedFieldViewModel) Get() *StringTestPlanChangedFieldViewModel {
	return v.value
}

func (v *NullableStringTestPlanChangedFieldViewModel) Set(val *StringTestPlanChangedFieldViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStringTestPlanChangedFieldViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStringTestPlanChangedFieldViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringTestPlanChangedFieldViewModel(val *StringTestPlanChangedFieldViewModel) *NullableStringTestPlanChangedFieldViewModel {
	return &NullableStringTestPlanChangedFieldViewModel{value: val, isSet: true}
}

func (v NullableStringTestPlanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringTestPlanChangedFieldViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


