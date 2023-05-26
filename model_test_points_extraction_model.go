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

// checks if the TestPointsExtractionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointsExtractionModel{}

// TestPointsExtractionModel struct for TestPointsExtractionModel
type TestPointsExtractionModel struct {
	Ids *GuidExtractionModel `json:"ids,omitempty"`
}

// NewTestPointsExtractionModel instantiates a new TestPointsExtractionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointsExtractionModel() *TestPointsExtractionModel {
	this := TestPointsExtractionModel{}
	return &this
}

// NewTestPointsExtractionModelWithDefaults instantiates a new TestPointsExtractionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointsExtractionModelWithDefaults() *TestPointsExtractionModel {
	this := TestPointsExtractionModel{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *TestPointsExtractionModel) GetIds() GuidExtractionModel {
	if o == nil || IsNil(o.Ids) {
		var ret GuidExtractionModel
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestPointsExtractionModel) GetIdsOk() (*GuidExtractionModel, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *TestPointsExtractionModel) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given GuidExtractionModel and assigns it to the Ids field.
func (o *TestPointsExtractionModel) SetIds(v GuidExtractionModel) {
	o.Ids = &v
}

func (o TestPointsExtractionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointsExtractionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableTestPointsExtractionModel struct {
	value *TestPointsExtractionModel
	isSet bool
}

func (v NullableTestPointsExtractionModel) Get() *TestPointsExtractionModel {
	return v.value
}

func (v *NullableTestPointsExtractionModel) Set(val *TestPointsExtractionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointsExtractionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointsExtractionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointsExtractionModel(val *TestPointsExtractionModel) *NullableTestPointsExtractionModel {
	return &NullableTestPointsExtractionModel{value: val, isSet: true}
}

func (v NullableTestPointsExtractionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointsExtractionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

