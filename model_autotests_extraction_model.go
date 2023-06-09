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

// checks if the AutotestsExtractionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutotestsExtractionModel{}

// AutotestsExtractionModel struct for AutotestsExtractionModel
type AutotestsExtractionModel struct {
	Ids *GuidExtractionModel `json:"ids,omitempty"`
}

// NewAutotestsExtractionModel instantiates a new AutotestsExtractionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutotestsExtractionModel() *AutotestsExtractionModel {
	this := AutotestsExtractionModel{}
	return &this
}

// NewAutotestsExtractionModelWithDefaults instantiates a new AutotestsExtractionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutotestsExtractionModelWithDefaults() *AutotestsExtractionModel {
	this := AutotestsExtractionModel{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *AutotestsExtractionModel) GetIds() GuidExtractionModel {
	if o == nil || IsNil(o.Ids) {
		var ret GuidExtractionModel
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutotestsExtractionModel) GetIdsOk() (*GuidExtractionModel, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *AutotestsExtractionModel) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given GuidExtractionModel and assigns it to the Ids field.
func (o *AutotestsExtractionModel) SetIds(v GuidExtractionModel) {
	o.Ids = &v
}

func (o AutotestsExtractionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutotestsExtractionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableAutotestsExtractionModel struct {
	value *AutotestsExtractionModel
	isSet bool
}

func (v NullableAutotestsExtractionModel) Get() *AutotestsExtractionModel {
	return v.value
}

func (v *NullableAutotestsExtractionModel) Set(val *AutotestsExtractionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutotestsExtractionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutotestsExtractionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutotestsExtractionModel(val *AutotestsExtractionModel) *NullableAutotestsExtractionModel {
	return &NullableAutotestsExtractionModel{value: val, isSet: true}
}

func (v NullableAutotestsExtractionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutotestsExtractionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


