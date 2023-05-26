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

// checks if the AutotestsSelectModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutotestsSelectModel{}

// AutotestsSelectModel struct for AutotestsSelectModel
type AutotestsSelectModel struct {
	Filter *AutotestFilterModel `json:"filter,omitempty"`
	Includes *SearchAutoTestsQueryIncludesModel `json:"includes,omitempty"`
}

// NewAutotestsSelectModel instantiates a new AutotestsSelectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutotestsSelectModel() *AutotestsSelectModel {
	this := AutotestsSelectModel{}
	return &this
}

// NewAutotestsSelectModelWithDefaults instantiates a new AutotestsSelectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutotestsSelectModelWithDefaults() *AutotestsSelectModel {
	this := AutotestsSelectModel{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AutotestsSelectModel) GetFilter() AutotestFilterModel {
	if o == nil || IsNil(o.Filter) {
		var ret AutotestFilterModel
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestsSelectModel) GetFilterOk() (*AutotestFilterModel, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AutotestsSelectModel) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given AutotestFilterModel and assigns it to the Filter field.
func (o *AutotestsSelectModel) SetFilter(v AutotestFilterModel) {
	o.Filter = &v
}

// GetIncludes returns the Includes field value if set, zero value otherwise.
func (o *AutotestsSelectModel) GetIncludes() SearchAutoTestsQueryIncludesModel {
	if o == nil || IsNil(o.Includes) {
		var ret SearchAutoTestsQueryIncludesModel
		return ret
	}
	return *o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestsSelectModel) GetIncludesOk() (*SearchAutoTestsQueryIncludesModel, bool) {
	if o == nil || IsNil(o.Includes) {
		return nil, false
	}
	return o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *AutotestsSelectModel) HasIncludes() bool {
	if o != nil && !IsNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given SearchAutoTestsQueryIncludesModel and assigns it to the Includes field.
func (o *AutotestsSelectModel) SetIncludes(v SearchAutoTestsQueryIncludesModel) {
	o.Includes = &v
}

func (o AutotestsSelectModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutotestsSelectModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Includes) {
		toSerialize["includes"] = o.Includes
	}
	return toSerialize, nil
}

type NullableAutotestsSelectModel struct {
	value *AutotestsSelectModel
	isSet bool
}

func (v NullableAutotestsSelectModel) Get() *AutotestsSelectModel {
	return v.value
}

func (v *NullableAutotestsSelectModel) Set(val *AutotestsSelectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutotestsSelectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutotestsSelectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutotestsSelectModel(val *AutotestsSelectModel) *NullableAutotestsSelectModel {
	return &NullableAutotestsSelectModel{value: val, isSet: true}
}

func (v NullableAutotestsSelectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutotestsSelectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

