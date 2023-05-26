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

// checks if the ParameterIterationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterIterationModel{}

// ParameterIterationModel struct for ParameterIterationModel
type ParameterIterationModel struct {
	Id string `json:"id"`
}

// NewParameterIterationModel instantiates a new ParameterIterationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterIterationModel(id string) *ParameterIterationModel {
	this := ParameterIterationModel{}
	this.Id = id
	return &this
}

// NewParameterIterationModelWithDefaults instantiates a new ParameterIterationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterIterationModelWithDefaults() *ParameterIterationModel {
	this := ParameterIterationModel{}
	return &this
}

// GetId returns the Id field value
func (o *ParameterIterationModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterIterationModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterIterationModel) SetId(v string) {
	o.Id = v
}

func (o ParameterIterationModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterIterationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableParameterIterationModel struct {
	value *ParameterIterationModel
	isSet bool
}

func (v NullableParameterIterationModel) Get() *ParameterIterationModel {
	return v.value
}

func (v *NullableParameterIterationModel) Set(val *ParameterIterationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterIterationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterIterationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterIterationModel(val *ParameterIterationModel) *NullableParameterIterationModel {
	return &NullableParameterIterationModel{value: val, isSet: true}
}

func (v NullableParameterIterationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterIterationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

