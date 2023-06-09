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

// checks if the TestSuiteV2GetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSuiteV2GetModel{}

// TestSuiteV2GetModel struct for TestSuiteV2GetModel
type TestSuiteV2GetModel struct {
	Id *string `json:"id,omitempty"`
	ParentId NullableString `json:"parentId,omitempty"`
	TestPlanId string `json:"testPlanId"`
	Name string `json:"name"`
	Type *TestSuiteType `json:"type,omitempty"`
	SaveStructure NullableBool `json:"saveStructure,omitempty"`
}

// NewTestSuiteV2GetModel instantiates a new TestSuiteV2GetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSuiteV2GetModel(testPlanId string, name string) *TestSuiteV2GetModel {
	this := TestSuiteV2GetModel{}
	this.TestPlanId = testPlanId
	this.Name = name
	return &this
}

// NewTestSuiteV2GetModelWithDefaults instantiates a new TestSuiteV2GetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSuiteV2GetModelWithDefaults() *TestSuiteV2GetModel {
	this := TestSuiteV2GetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestSuiteV2GetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteV2GetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestSuiteV2GetModel) SetId(v string) {
	o.Id = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2GetModel) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2GetModel) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *TestSuiteV2GetModel) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *TestSuiteV2GetModel) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *TestSuiteV2GetModel) UnsetParentId() {
	o.ParentId.Unset()
}

// GetTestPlanId returns the TestPlanId field value
func (o *TestSuiteV2GetModel) GetTestPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestPlanId
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value
// and a boolean to check if the value has been set.
func (o *TestSuiteV2GetModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestPlanId, true
}

// SetTestPlanId sets field value
func (o *TestSuiteV2GetModel) SetTestPlanId(v string) {
	o.TestPlanId = v
}

// GetName returns the Name field value
func (o *TestSuiteV2GetModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestSuiteV2GetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestSuiteV2GetModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TestSuiteV2GetModel) GetType() TestSuiteType {
	if o == nil || IsNil(o.Type) {
		var ret TestSuiteType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteV2GetModel) GetTypeOk() (*TestSuiteType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TestSuiteType and assigns it to the Type field.
func (o *TestSuiteV2GetModel) SetType(v TestSuiteType) {
	o.Type = &v
}

// GetSaveStructure returns the SaveStructure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2GetModel) GetSaveStructure() bool {
	if o == nil || IsNil(o.SaveStructure.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveStructure.Get()
}

// GetSaveStructureOk returns a tuple with the SaveStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2GetModel) GetSaveStructureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveStructure.Get(), o.SaveStructure.IsSet()
}

// HasSaveStructure returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasSaveStructure() bool {
	if o != nil && o.SaveStructure.IsSet() {
		return true
	}

	return false
}

// SetSaveStructure gets a reference to the given NullableBool and assigns it to the SaveStructure field.
func (o *TestSuiteV2GetModel) SetSaveStructure(v bool) {
	o.SaveStructure.Set(&v)
}
// SetSaveStructureNil sets the value for SaveStructure to be an explicit nil
func (o *TestSuiteV2GetModel) SetSaveStructureNil() {
	o.SaveStructure.Set(nil)
}

// UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
func (o *TestSuiteV2GetModel) UnsetSaveStructure() {
	o.SaveStructure.Unset()
}

func (o TestSuiteV2GetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSuiteV2GetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	toSerialize["testPlanId"] = o.TestPlanId
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.SaveStructure.IsSet() {
		toSerialize["saveStructure"] = o.SaveStructure.Get()
	}
	return toSerialize, nil
}

type NullableTestSuiteV2GetModel struct {
	value *TestSuiteV2GetModel
	isSet bool
}

func (v NullableTestSuiteV2GetModel) Get() *TestSuiteV2GetModel {
	return v.value
}

func (v *NullableTestSuiteV2GetModel) Set(val *TestSuiteV2GetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSuiteV2GetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSuiteV2GetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSuiteV2GetModel(val *TestSuiteV2GetModel) *NullableTestSuiteV2GetModel {
	return &NullableTestSuiteV2GetModel{value: val, isSet: true}
}

func (v NullableTestSuiteV2GetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSuiteV2GetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


